# gincontroller
gin struct controller

###example

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yansuan/gincontroller"
	"log"
)

type MyController struct {
	gincontroller.Controller
}

func (this *MyController) Prepare(ctx *gin.Context) error {
	log.Println("Prepare")
	return nil
}

func (this *MyController) Finish(ctx *gin.Context) error {
	log.Println("Finish")
	return nil
}

func (this *MyController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println("Get", id)

	ctx.JSON(200, ctx.Params)
}

func (this *MyController) GetAll(ctx *gin.Context) {
	log.Println("GetAll")

	ctx.JSON(200, ctx.FullPath())
}

func (this *MyController) All(ctx *gin.Context) {
	log.Println("All")

	ctx.JSON(200, ctx.FullPath())
}

func main() {
	r := gin.Default()
	myCtl := new(MyController)
	gincontroller.AddRouter(r, "/", myCtl, "get:GetAll")
	gincontroller.AddRouter(r, "/:id", myCtl, "get:Get")
	gincontroller.AddRouter(r, "/:id/*all", myCtl, "get:All")

	r.Run()
}

```