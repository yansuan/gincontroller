package gincontroller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"strings"
)

var (
	routers []*Router
)

type Router struct {
	Path           string
	Controller     IController
	ControllerType reflect.Type
	Method         string
	Name           string
}

/**
mappingMethods:get:Get;delete:Delete;put:Put
*/
func AddRouter(r *gin.Engine, path string, ctl IController, mappingMethods string) {
	for _, mapping := range strings.Split(mappingMethods, ";") {
		values := strings.Split(mapping, ":")
		if len(values) != 2 {
			log.Println(path, mappingMethods, "mapping format error")
			break
		}

		method := strings.ToUpper(values[0])
		name := values[1]

		ctlVal := reflect.ValueOf(ctl)
		ctlType := reflect.Indirect(ctlVal).Type()

		router := &Router{}
		router.Controller = ctl
		router.ControllerType = ctlType
		router.Path = path
		router.Method = method
		router.Name = name
		routers = append(routers, router)

		if method == http.MethodGet {
			r.GET(path, RouterDefault)
		} else if method == http.MethodPost {
			r.POST(path, RouterDefault)
		} else if method == http.MethodPut {
			r.PUT(path, RouterDefault)
		} else if method == http.MethodDelete {
			r.DELETE(path, RouterDefault)
		} else if method == http.MethodHead {
			r.HEAD(path, RouterDefault)
		} else if method == http.MethodPatch {
			r.PATCH(path, RouterDefault)
		} else if method == http.MethodOptions {
			r.OPTIONS(path, RouterDefault)
		} else if method == "*" {
			r.Any(path, RouterDefault)
		}
	}
}

func getRouter(ctx *gin.Context) *Router {

	//
	path := ctx.FullPath()
	method := ctx.Request.Method
	for _, router := range routers {
		if router.Path == path && (router.Method == method || router.Method == "*") {
			return router
		}
	}
	return nil
}

func RouterDefault(ctx *gin.Context) {
	router := getRouter(ctx)
	if router == nil {
		panic("controller not found")
		return
	}

	//
	vc := reflect.New(router.ControllerType)
	execController, ok := vc.Interface().(IController)
	if !ok {
		panic("controller is not ControllerInterface")
	}

	elemVal := reflect.ValueOf(router.Controller).Elem()
	elemType := reflect.TypeOf(router.Controller).Elem()
	execElem := reflect.ValueOf(execController).Elem()

	numOfFields := elemVal.NumField()
	for i := 0; i < numOfFields; i++ {
		fieldType := elemType.Field(i)
		elemField := execElem.FieldByName(fieldType.Name)
		if elemField.CanSet() {
			fieldVal := elemVal.Field(i)
			elemField.Set(fieldVal)
		}
	}

	//
	execController.Prepare(ctx)
	//

	ctlValue := reflect.ValueOf(execController)
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(ctx)
	callMethod := ctlValue.MethodByName(router.Name)
	callMethod.Call(in)

	//
	execController.Finish(ctx)
}
