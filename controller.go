package gincontroller

import "github.com/gin-gonic/gin"

type IController interface {
	Prepare(ctx *gin.Context)
	Finish(ctx *gin.Context)
}

type Controller struct {
}

func (this *Controller) Prepare(ctx *gin.Context) {
}

func (this *Controller) Finish(ctx *gin.Context) {
}