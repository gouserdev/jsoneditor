package main

import (
	gin "github.com/gouserdev/jsoneditor/pkg/gin"
)

func main() {
	gin.GinInit("0.0.0.0", "10889")
}
