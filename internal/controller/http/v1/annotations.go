package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ytanne/annotation-manager/pkg/models"
)

func (s server) AddAnnotation(g *gin.Context) {
	var annotation models.Annotation

	if err := g.ShouldBind(&annotation); err != nil {
		log.Println("binding annotation failed", err)
		g.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	id, err := s.u.AddAnnotation(g.Request.Context(), annotation)
	if err != nil {
		log.Println("adding annotation failed", err)
		g.JSON(http.StatusInternalServerError, "adding annotation failed")
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

func (s server) UpdateAnnotation(g *gin.Context) {
	idStr := g.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("obtaining id failed", err)
		g.JSON(http.StatusBadGateway, "failed to update annotation")
		return
	}

	var annotation models.Annotation
	if err := g.ShouldBind(&annotation); err != nil {
		log.Println("binding annotation failed", err)
		g.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	annotation.ID = id

	err = s.u.UpdateAnnotation(g.Request.Context(), annotation)
	if err != nil {
		log.Println("updating annotation failed", err)
		g.JSON(http.StatusBadRequest, "failed to update annotation")
		return
	}

	g.JSON(http.StatusOK, "success")
}

func (s server) GetAnnotation(g *gin.Context) {
	idStr := g.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("obtaining id failed", err)
		g.JSON(http.StatusBadGateway, "failed to get annotation")
		return
	}

	annotation, err := s.u.GetAnnotation(g.Request.Context(), id)
	if err != nil {
		log.Println("obtaining annotation failed", err)
		g.JSON(http.StatusBadGateway, "failed to get annotation")
		return
	}

	g.JSON(http.StatusOK, annotation)
}

func (s server) GetAnnotationsByVideoID(g *gin.Context) {
	idStr := g.Query("video_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("obtaining id failed", err)
		g.JSON(http.StatusBadGateway, "failed to get annotations")
		return
	}

	annotations, err := s.u.GetAnnotationsByVideoID(g.Request.Context(), id)
	if err != nil {
		log.Println("obtaining annotations failed", err)
		g.JSON(http.StatusBadGateway, "failed to get annotations")
		return
	}

	g.JSON(http.StatusOK, annotations)
}

func (s server) DeleteAnnotation(g *gin.Context) {
	idStr := g.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("obtaining id failed", err)
		g.JSON(http.StatusBadGateway, "failed to get annotations")
		return
	}

	err = s.u.DeleteAnnotation(g.Request.Context(), id)
	if err != nil {
		log.Println("deleting annotation failed", err)
		g.JSON(http.StatusBadGateway, "failed to delete annotation")
		return
	}

	g.JSON(http.StatusOK, "deleted")
}
