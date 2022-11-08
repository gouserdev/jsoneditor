package jsoneditor

import (
	gin "github.com/gouserdev/jsoneditor/pkg/gin"
	"github.com/gouserdev/jsoneditor/pkg/util"
)

func Jsoneditor(ip string, port string) {
	configpath := "./configs/config.json"
	//util.Open("http://127.0.0.1:10889/j/index.html?pass=" + util.GetFromJson("password", configpath))
	gin.GinInit(ip, port)
}
