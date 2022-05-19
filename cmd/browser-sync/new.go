package main

import (
	"browser-sync/internal/common"
	"browser-sync/internal/file"
	http2 "browser-sync/internal/http"
	"log"
	"path/filepath"
	"time"
)

func newWatcherDir() {
	dir := filepath.Clean(common.Config.GetString("dir"))
	d := common.Config.GetDuration("time") * time.Second
	ignore := common.Config.GetStringSlice("ignore")
	for k, v := range ignore {
		ignore[k] = filepath.Clean(v)
	}

	file.WatcherDir(dir, d, ignore)
}

func newTaskRun() {
	for {
		select {
		case err := <-common.Errors:
			log.Println("err: ", err)
			http2.NoticeAllUsers("err")
			go newWatcherDir()
		case <-common.Change:
			http2.NoticeAllUsers("ok")
		case client := <-common.Register:
			common.Users[client.(*http2.Client).ID] = client
		case id := <-common.Unregister:
			delete(common.Users, id)
		}
	}
}
