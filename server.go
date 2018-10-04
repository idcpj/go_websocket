package main

import (
	"net/http"

	"./impl"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	http.HandleFunc("/ws", wsHandle)

	http.ListenAndServe("0.0.0.0:7777", nil)
}

///ws api
func wsHandle(writer http.ResponseWriter, request *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *impl.Connection
		data   []byte
	)

	///简历 websocket 请求
	if wsConn, err = upgrader.Upgrade(writer, request, nil); err != nil {
		return
	}

	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	//测试心跳
	/*go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat ")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}

	}()*/

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}
