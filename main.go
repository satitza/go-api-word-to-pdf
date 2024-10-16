package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api-word-to-pdf/common"
	"go-api-word-to-pdf/configuration"
	"go-api-word-to-pdf/middleware"
	"go-api-word-to-pdf/router"
)

func main() {

	config, err := configuration.GetConfig()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(middleware.LoggingMiddleware())

	// initial all router
	if len(router.AllAppRouter) > 0 {
		for index := 0; index < len(router.AllAppRouter); index++ {
			switch router.AllAppRouter[index].Method {
			case common.AppHttpMethodGet:
				r.GET(router.AllAppRouter[index].Path, router.AllAppRouter[index].Handler)
			case common.AppHttpMethodPost:
				r.POST(router.AllAppRouter[index].Path, router.AllAppRouter[index].Handler)
			case common.AppHttpMethodPut:
				r.PUT(router.AllAppRouter[index].Path, router.AllAppRouter[index].Handler)
			case common.AppHttpMethodDelete:
				r.DELETE(router.AllAppRouter[index].Path, router.AllAppRouter[index].Handler)
			default:

			}
		}
	}

	err = r.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return
	}
}
