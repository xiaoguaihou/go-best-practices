package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// implement of restful api
func (s*CoreService)OnLogin(c *gin.Context) {
	var req LoginRequest
	var rsp LoginResponse

	if paramCheck(&req, c) {
		//...
		c.JSON(http.StatusOK, &rsp)
	}
}

func paramCheck(req interface{}, c *gin.Context) bool {
	var result bool

	if err := c.ShouldBindJSON(req); err != nil {
		var s Status
		s = StatusInvalidParam
		c.JSON(http.StatusOK, gin.H{
			"status":  s,
			"message": s.Message(),
		})
		result = false
	} else {
		result = true
	}
	return result
}