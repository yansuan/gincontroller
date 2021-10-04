package gincontroller

import "github.com/gin-gonic/gin"

type IController interface {
	Prepare(ctx *gin.Context) error
	Finish(ctx *gin.Context) error
}

type Controller struct {
}

func (this *Controller) Prepare(ctx *gin.Context) error {
	return nil
}

func (this *Controller) Finish(ctx *gin.Context) error {
	return nil
}
