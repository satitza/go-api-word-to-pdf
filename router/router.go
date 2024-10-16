package router

import (
	"github.com/gin-gonic/gin"
	"go-api-word-to-pdf/common"
	"go-api-word-to-pdf/handler"
)

type AppRouter struct {
	Method  common.AppHttpMethod
	Path    string
	Handler gin.HandlerFunc
}

var AllAppRouter = []AppRouter{
	{
		Method:  common.AppHttpMethodPost,
		Path:    "word-to-pdf",
		Handler: handler.AppConvertWordToPdfHandler,
	},
}
