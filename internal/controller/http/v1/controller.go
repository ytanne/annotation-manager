package v1

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ytanne/annotation-manager/internal/config"
	"github.com/ytanne/annotation-manager/internal/controller"
)

const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 10 * time.Second
)

type server struct {
	httpPort int
	router   *gin.Engine
	u        controller.Usecase
}

func NewController(cfg config.AuthConfig, u controller.Usecase) (server, error) {
	s := server{
		httpPort: cfg.HttpPort,
		u:        u,
	}

	am, err := newAuthManager(cfg.IdentityKey, cfg.SigningKey)
	if err != nil {
		return server{}, fmt.Errorf("failed to create auth manager: %w", err)
	}

	router := gin.Default()

	authRoutes := router.Group("")
	{
		authRoutes.POST("/login", am.LoginHandler)

		auth := authRoutes.Group("/user")

		// Refresh time can be longer than token timeout
		auth.GET("/refresh_token", am.RefreshHandler)

	}

	r := router.Group("")
	r.Use(am.MiddlewareFunc())

	// video endpoints
	videoRoutes := r.Group("/videos")
	{
		videoRoutes.POST("", s.AddVideo)
		videoRoutes.PATCH("/:id", s.UpdateVideo)
		videoRoutes.GET("/:id", s.GetVideo)
		videoRoutes.DELETE("/:id", s.DeleteVideo)
	}

	annotationRoutes := r.Group("/annotations")
	{
		annotationRoutes.POST("", s.AddAnnotation)
		annotationRoutes.PATCH("/:id", s.UpdateAnnotation)
		annotationRoutes.GET("/:id", s.GetAnnotation)
		annotationRoutes.GET("/", s.GetAnnotationsByVideoID)
		annotationRoutes.DELETE("/:id", s.DeleteAnnotation)
	}

	s.router = router

	return s, nil
}

func (s server) Run() error {
	hs := &http.Server{
		Handler:      s.router,
		Addr:         fmt.Sprintf(":%d", s.httpPort),
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	log.Println("starting server on port", s.httpPort)

	return hs.ListenAndServe()
}
