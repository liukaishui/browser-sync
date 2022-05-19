package http

import (
	"browser-sync/internal/common"
	"browser-sync/pkg/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

func ServeWS(w http.ResponseWriter, r *http.Request) {
	ws, err := common.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer ws.Close()

	addrHash, err := utils.Md5(ws.RemoteAddr().String())
	if err != nil {
		log.Println("md5:", err)
		return
	}
	client := &Client{
		ID:   addrHash,
		Conn: ws,
	}
	common.Register <- client

	for {
		mt, message, err := client.Conn.ReadMessage()
		if err != nil {
			common.Unregister <- client.ID
			return
		}
		err = client.Conn.WriteMessage(mt, message)
		if err != nil {
			common.Unregister <- client.ID
			return
		}
	}
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "I'm ok")
}
