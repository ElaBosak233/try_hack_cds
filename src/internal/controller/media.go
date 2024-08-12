package controller

import (
	"github.com/elabosak233/cloudsdale/internal/service"
	"github.com/elabosak233/cloudsdale/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

type IMediaController interface {
	GetFile(ctx *gin.Context)
}

type MediaController struct {
	mediaService service.IMediaService
}

func NewMediaController(s *service.Service) IMediaController {
	return &MediaController{
		mediaService: s.MediaService,
	}
}

func (m *MediaController) GetFile(ctx *gin.Context) {
	a := ctx.Param("path")
	p := path.Join(utils.MediaPath, a)
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.File(p)
}
