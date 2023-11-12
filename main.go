package main

import (
	"fmt"
	websocket_handler "gamebackend/src/websocket-handler"
	"net/http"
)

func main() {
	websocket_handler.InitWSService()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
