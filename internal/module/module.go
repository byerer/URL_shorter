package module

import (
	"URL_shorter/internal/module/urlShorter"
	"github.com/gin-gonic/gin"
)

type Module interface {
	InitRouter(r *gin.RouterGroup)
}

var Modules []Module

func RegisterModule(m Module) {
	Modules = append(Modules, m)
}

func init() {
	RegisterModule(&urlShorter.ModelURL{})
}
