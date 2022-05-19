package common

import (
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"net/http"
)

var (
	Config *viper.Viper
)

var (
	Errors = make(chan interface{})
	Change = make(chan bool)
)

var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Users      = make(map[string]interface{})
	Register   = make(chan interface{})
	Unregister = make(chan string)
)
