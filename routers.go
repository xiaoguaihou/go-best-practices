package main

import (
	"github.com/gin-gonic/gin"
)

// all external inferfaces are registered here
// all interface definitions are locate in related package with file name api-dto.go
//
// in most case, the interface implementation is singleton
// but we can also make it multi-instancable
func setupApi(router *gin.Engine) {
	// all url path must start with service name!!!
	router.POST("/go-best-practices/login", coreService.OnLogin)
}
