package gin

import (
	"net/http"

	"github.com/gouserdev/jsoneditor/pkg/gin/middleware"
	"github.com/gouserdev/jsoneditor/pkg/util"

	_ "github.com/gouserdev/jsoneditor/pkg/gin/dist"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
	"github.com/rakyll/statik/fs"
)

func GetJsonController(c *gin.Context) {
	configpath := "./configs/config.json"
	util.CheckJsonFile(configpath)
	if c.PostForm("pass") == util.GetFromJson("password", configpath) {
		c.String(http.StatusOK, util.GetFromJson("", configpath))
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "password is wrong!"})
	}
}

//curl -d 'pass=hello'  http://127.0.0.1:10889/getjson -X POST
func SetJsonController(c *gin.Context) {
	configpath := "./configs/config.json"
	if c.PostForm("pass") == util.GetFromJson("password", configpath) {
		util.SetJsonValue("", c.PostForm("value"), configpath)
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "password is wrong!"})
	}
}

func GinInit(ip string, port string) {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	//r.StaticFS("/j/", http.Dir("./pkg/gin/dist"))
	statikFS, err := fs.New()
	if err != nil {
		r.StaticFS("/j/", http.Dir("./pkg/gin/view"))
	} else {
		r.StaticFS("/j/", statikFS)
	}

	r.POST("/getjson", GetJsonController)
	r.POST("/setjson", SetJsonController)
	r.Run(ip + ":" + port)
}

func handle_Bug() {
	cors.Default()
	dump.P("handle_Bug")
}
