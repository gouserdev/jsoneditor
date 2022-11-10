package gin

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
	"github.com/gorilla/websocket"
)

var (
	wsUpgrader = websocket.Upgrader{}
	rs         = 0
)

func Status(c *gin.Context) {
	// defer ws.Close()
	mainProcess(c)
}

// 主程序，负责循环读取客户端消息跟消息的发送
//ws *websocket.Conn
func mainProcess(c *gin.Context) {
	// @see https://github.com/gorilla/websocket/issues/523
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, _ := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	defer ws.Close()
	rs = 0
	for {
		_, message, err := ws.ReadMessage()
		_ = err
		serveMsgStr := message
		dump.P(string(serveMsgStr))

		if string(serveMsgStr) == `isonline` {
			ws.WriteMessage(websocket.TextMessage, []byte(`{"status":0,"data":"online"}`))
			log.Println("is online,waitting")
			time.Sleep(time.Duration(1) * time.Second)
			continue
		} else if string(serveMsgStr) == `close` {
			os.Exit(0)
		}

		if err != nil { // 离线通知
			log.Println("ReadMessage error")
			log.Println("waiting a few moment")
			time.Sleep(time.Duration(2) * time.Second)
			rs = rs + 1
			dump.P(rs)
			if rs > 4 {
				ws.Close()
				os.Exit(0)
				break
			}
		}

	}
}
