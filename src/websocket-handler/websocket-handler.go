package websocket_handler

import (
	"fmt"
	"gamebackend/src/players-handler"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

var connections []*websocket.Conn

type SocketMessage struct {
	EventType int                    `json:"actionType"`
	Data      map[string]interface{} `json:"data"`
}

type EventHandlers [](func(map[string]interface{}) []byte)

const (
	PLAYER_ACTION      = 0
	PLAYER_STAT_CHANGE = 1
)

var handlers = EventHandlers{
	players.HandlePlayerActionEvent,     // 0[PLAYER_ACTION]
	players.HandlePlayerStatChangeEvent, // 1[PLAYER_STAT_CHANGE]
}

func handler(writter http.ResponseWriter, request *http.Request) {
	connection, err := upgrader.Upgrade(writter, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		connections = append(connections, connection)
		fmt.Println("Connection established, total: " + strconv.Itoa(len(connections)))
	}

	for {
		var event SocketMessage
		err = connection.ReadJSON(&event)
		if err != nil {
			fmt.Println("Invalid JSON in event")
			return
		}
		var response []byte = handlers[event.EventType](event.Data)
		if response != nil {
			for i := 0; i < len(connections); i++ {
				connections[i].WriteMessage(websocket.TextMessage, response)
			}
		}
	}
}

func InitWSService() {
	http.HandleFunc("/ws", handler)
}
