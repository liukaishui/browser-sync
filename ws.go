package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	users      = make(map[string]*Client)
	register   = make(chan *Client)
	unregister = make(chan string)
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

func ServerWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer ws.Close()

	id, err := Md5(strconv.FormatInt(time.Now().UnixNano(), 10))
	if err != nil {
		log.Println("md5:", err)
		return
	}
	client := &Client{
		ID:   id,
		Conn: ws,
	}
	register <- client

	for {
		mt, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			unregister <- client.ID
			return
		}
		err = client.Conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			unregister <- client.ID
			return
		}
	}
}

func ServerHome(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, time.Now().Format("2006-01-02 15:04:05"))
}
