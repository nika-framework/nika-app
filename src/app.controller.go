package src

import (
	"NikaApp/src/dto"

	"github.com/gin-gonic/gin"
	"github.com/nika-framework/nika/common/response"
	"github.com/nika-framework/nika/common/validator"
)

type AppController struct {
	service  *AppService
	GetApp func(*gin.Context) `route:"GET:/"`
	SayHello func(*gin.Context) `route:"GET:/:name"`
}

func NewAppController(service *AppService) *AppController {
	c := &AppController{service: service}

	c.SayHello = func(ctx *gin.Context) {
		var findOneDto dto.FindOneDto
		if !validator.BindAndValidateUri(ctx, &findOneDto) {
			return
		}
		message := c.service.GetWelcomeMessage(findOneDto.Name)
		response.Ok(ctx, gin.H{"message": message})
	}
	c.GetApp = func(ctx *gin.Context) {
		 
		response.Ok(ctx, gin.H{"message": "Hi Nika 0.1.1"})
	}

	return c
}
