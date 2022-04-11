package handlerhttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MapRoute struct {
	HttpMethod   string
	RelativePath string
	HandlerFunc  func(a interface{}) gin.HandlerFunc
	IoC          func() interface{}
}

var r *gin.Engine

func ConfigureMapRoute(maps func() []MapRoute) {
	for _, mr := range maps() {
		if mr.HttpMethod == http.MethodGet {
			r.GET(mr.RelativePath, mr.HandlerFunc(mr.IoC()))
		}

		if mr.HttpMethod == http.MethodPost {
			r.GET(mr.RelativePath, mr.HandlerFunc(mr.IoC()))
		}

		if mr.HttpMethod == http.MethodPut {
			r.GET(mr.RelativePath, mr.HandlerFunc(mr.IoC()))
		}

		if mr.HttpMethod == http.MethodDelete {
			r.GET(mr.RelativePath, mr.HandlerFunc(mr.IoC()))
		}

		if mr.HttpMethod == http.MethodHead {
			r.GET(mr.RelativePath, mr.HandlerFunc(mr.IoC()))
		}

		if mr.HttpMethod == http.MethodOptions {
			r.GET(mr.RelativePath, mr.HandlerFunc(mr.IoC()))
		}

		if mr.HttpMethod == http.MethodTrace {
			r.GET(mr.RelativePath, mr.HandlerFunc(mr.IoC()))
		}
	}
}

func New() *gin.Engine {
	r = gin.Default()
	return r
}
