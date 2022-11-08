package gin

import (
	"github.com/gouserdev/jsoneditor/pkg/gin/middleware"
	"github.com/gouserdev/jsoneditor/pkg/util"
	"net/http"

	_ "github.com/gouserdev/jsoneditor/pkg/gin/dist"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
	"github.com/rakyll/statik/fs"
)

func GetJsonController(c *gin.Context) {
	configpath := "./configs/config.json"
	// c.Writer.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Origin")
	// c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")

	//util.GetFromJson("password", configpath)
	//dump.P(c.PostForm("pass"), util.GetFromJson("password", configpath))
	if c.PostForm("pass") == util.GetFromJson("password", configpath) {
		//dump.P(util.GetFromJson("", configpath))
		c.String(http.StatusOK, util.GetFromJson("", configpath))
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "password is wrong!"})
	}
}

//curl -d 'pass=hello'  http://127.0.0.1:10889/getjson -X POST
func SetJsonController(c *gin.Context) {
	configpath := "./configs/config.json"
	// c.Writer.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Origin")
	// c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
	//dump.P(c.PostForm("pass"), util.GetFromJson("password", configpath), c.PostForm("value"))
	if c.PostForm("pass") == util.GetFromJson("password", configpath) {
		util.SetJsonValue("", c.PostForm("value"), configpath)
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "password is wrong!"})
	}
}

func GinInit(ip string, port string) {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) // 这一行

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

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		method := c.Request.Method
// 		origin := c.Request.Header.Get("Origin") //请求头部
// 		if origin != "" {
// 			//接收客户端发送的origin （重要！）
// 			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
// 			//服务器支持的所有跨域请求的方法
// 			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
// 			//允许跨域设置可以返回其他子段，可以自定义字段
// 			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
// 			// 允许浏览器（客户端）可以解析的头部 （重要）
// 			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
// 			//设置缓存时间
// 			c.Header("Access-Control-Max-Age", "172800")
// 			//允许客户端传递校验信息比如 cookie (重要)
// 			c.Header("Access-Control-Allow-Credentials", "true")
// 		}

// 		//允许类型校验
// 		if method == "OPTIONS" {
// 			c.JSON(http.StatusOK, "ok!")
// 		}

// 		defer func() {
// 			if err := recover(); err != nil {
// 				log.Printf("Panic info is: %v", err)
// 			}
// 		}()

// 		c.Next()
// 	}
// }

func handle_Bug() {
	cors.Default()
	dump.P("handle_Bug")
}
