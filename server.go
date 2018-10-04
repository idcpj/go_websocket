package main

import (
	"net/http"

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
		conn *websocket.Conn
		err  error
		data []byte
	)

	///简历 websocket 请求
	if conn, err = upgrader.Upgrade(writer, request, nil); err != nil {
		return
	}
	//循环链接
	for {
		//data 是接受到的数据
		if _, data, err = conn.ReadMessage(); err != nil {
			//错误就关闭连接
			goto ERROR
		}
		//第二参数是发送给client 数据
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERROR
		}
	}

ERROR:
	conn.Close()
}
