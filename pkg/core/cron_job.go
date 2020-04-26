package core

import (
	rest "github.com/xiaoguaihou/go-rest"

	"github.com/xiaoguaihou/go-rest/dingding"
)

// just a sample to break down the core serivce implementation
func (s *CoreService) sendHourReport() {
	var req LoginRequest
	var rsp LoginResponse
	rest.Post("htpp://.../api", &req, &rsp,
		// http header
		map[string]string{
			"Authorization": "APPCODE code...",
		},
		// query parameters
		map[string]string{
			"code":   "code",
			"mobile": "15510165711",
			"name":   "爱直飞",
		})

	dingding.Post2Dingding("https://oapi.dingtalk.com/robot/send?access_token=token", "hello dingding")
}
