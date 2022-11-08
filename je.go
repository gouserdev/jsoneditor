package jsoneditor

import (
	gin "github.com/gouserdev/jsoneditor/pkg/gin"
)

func Jsoneditor(ip string, port string) {
	gin.GinInit(ip, port)
}
