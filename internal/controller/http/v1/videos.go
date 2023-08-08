package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ytanne/annotation-manager/pkg/models"
)

func (s server) AddVideo(g *gin.Context) {
	var video models.Video

	if err := g.ShouldBind(&video); err != nil {
		log.Println("binding video failed", err)
		g.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	id, err := s.u.AddVideo(g.Request.Context(), video)
	if err != nil {
		log.Println("adding video failed", err)
		g.JSON(http.StatusInternalServerError, "adding video failed")
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

func (s server) UpdateVideo(g *gin.Context) {
	idStr := g.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("obtaining id failed", err)
		g.JSON(http.StatusBadGateway, "failed to update video")
		return
	}

	var video models.Video
	if err := g.ShouldBind(&video); err != nil {
		log.Println("binding video failed", err)
		g.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	video.ID = id

	err = s.u.UpdateVideo(g.Request.Context(), video)
	if err != nil {
		log.Println("updating video failed", err)
		g.JSON(http.StatusBadRequest, "failed to update video")
		return
	}

	g.JSON(http.StatusOK, "success")
}

func (s server) GetVideo(g *gin.Context) {
	idStr := g.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("obtaining id failed", err)
		g.JSON(http.StatusBadGateway, "failed to get video")
		return
	}

	video, err := s.u.GetVideo(g.Request.Context(), id)
	if err != nil {
		log.Println("obtaining video failed", err)
		g.JSON(http.StatusBadGateway, "failed to get video")
		return
	}

	g.JSON(http.StatusOK, video)
}

func (s server) DeleteVideo(g *gin.Context) {
	idStr := g.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("obtaining id failed", err)
		g.JSON(http.StatusBadGateway, "failed to get videos")
		return
	}

	err = s.u.DeleteVideo(g.Request.Context(), id)
	if err != nil {
		log.Println("deleting video failed", err)
		g.JSON(http.StatusBadGateway, "failed to delete video")
		return
	}

	g.JSON(http.StatusOK, "deleted")
}
