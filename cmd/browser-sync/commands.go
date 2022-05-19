package main

import (
	"browser-sync/internal/common"
	http2 "browser-sync/internal/http"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func initialization(cmd *cobra.Command, args []string) {
	viper.SetConfigFile("./config.yaml")

	viper.SetDefault("server", ":2345")
	viper.SetDefault("time", 1)
	viper.SetDefault("dir", "./")
	viper.SetDefault("ignore", []string{".idea", ".git", "config.yaml"})

	err := viper.SafeWriteConfigAs("./config.yaml")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}

func run(cmd *cobra.Command, args []string) {
	v := viper.New()
	v.SetConfigFile(config)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	common.Config = v

	go newWatcherDir()
	go newTaskRun()

	http.HandleFunc("/", http2.ServeHTTP)
	http.HandleFunc("/ws", http2.ServeWS)
	err := http.ListenAndServe(common.Config.GetString("server"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
