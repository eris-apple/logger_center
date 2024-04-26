package ws

import (
	"fmt"
	"github.com/aetherteam/logger_center/internal/enums"
)

type Broadcast struct {
	Type string
	Data interface{}
}

func WSBroadcast(b *Broadcast, projectID string) {
	for client := range clients {
		if clients[client] == nil {
			delete(clients, client)
			continue
		}

		fmt.Println("clients[client].ProjectID")
		fmt.Println(clients[client].ProjectID)

		if clients[client].ProjectID != projectID {
			return
		}

		response := &Message{}

		switch b.Type {
		case enums.LogCreate.String():
			response.Type = enums.LogSend.String()
			response.Data = b.Data
		default:
			response.Type = "unexpected type"
			response.Data = nil
		}

		err := client.WriteJSON(response)
		if err != nil {
			fmt.Println("Error sending message:", err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}
