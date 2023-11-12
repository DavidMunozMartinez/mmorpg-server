package players

import "fmt"

// Player action events
const (
	MOVE_UP    = 0
	MOVE_DOWN  = 1
	MOVE_LEFT  = 2
	MOVE_RIGHT = 3
)

func HandlePlayerActionEvent(data map[string]interface{}) []byte {
	if data["playerId"] == nil || data["direction"] == nil {
		fmt.Println("Invalid request")
		return []byte("Invalid request")
	}
	playerId := data["playerId"].(string)
	response := data["direction"].(string)

	fmt.Println(response)
	return []byte("{ playerId: \"" + playerId + "\", action: " + response + "}")
}

func HandlePlayerStatChangeEvent(data map[string]interface{}) []byte {
	return []byte("Not implemented yet")
}
