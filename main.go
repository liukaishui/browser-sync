package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

func main() {
	go WatcherDir(
		filepath.Clean("D:\\Go_WorkSpace\\src\\shigo"),
		1*time.Second,
		[]string{
			filepath.Clean(".idea"),
			filepath.Clean("go.mod"),
			filepath.Clean("go.sum"),
		},
	)

	go func() {
		for {
			select {
			case err := <-errors:
				fmt.Println("err: ", err)
				for _, client := range users {
					_ = client.Conn.WriteMessage(websocket.TextMessage, []byte("err"))
				}
				return
			case notice := <-change:
				fmt.Println("notice: ", notice)
				for _, client := range users {
					_ = client.Conn.WriteMessage(websocket.TextMessage, []byte("ok"))
				}
			case client := <-register:
				users[client.ID] = client
			case id := <-unregister:
				delete(users, id)
			}
		}
	}()

	http.HandleFunc("/", ServerHome)
	http.HandleFunc("/ws", ServerWS)
	err := http.ListenAndServe(":2048", nil)
	if err != nil {
		log.Fatal(err)
	}
}
