package http

import (
	"browser-sync/internal/common"
	"github.com/gorilla/websocket"
)

func NoticeAllUsers(message string) {
	for _, client := range common.Users {
		_ = client.(*Client).Conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}
