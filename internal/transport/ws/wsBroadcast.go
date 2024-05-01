package ws

import (
	"fmt"
	"github.com/aetherteam/logger_center/internal/enums"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/services"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/mitchellh/mapstructure"
)

type Broadcast struct {
	Type string
	Data interface{}
}

func WSBroadcast(b *Broadcast, store store.Store, c *Client) {
	for client := range clients {
		if clients[client] == nil {
			delete(clients, client)
			continue
		}

		if clients[client].ProjectID != c.ProjectID {
			return
		}

		if clients[client].Secret != c.Secret {
			return
		}

		pr := store.Project()
		lr := store.Log()

		ls := services.NewLogService(lr, pr)

		response := &Message{}

		switch b.Type {
		case enums.LogCreate.String():
			body := &models.Log{}
			if err := mapstructure.Decode(b.Data, &body); err != nil {
				fmt.Println("error with decode data")
			}

			log := &models.Log{
				ProjectID: c.ProjectID,
				ChainID:   body.ChainID,
				Content:   body.Content,
				Level:     body.Level,
				Timestamp: body.Timestamp,
			}

			if _, err := ls.Create(log); err != nil {
				fmt.Println("error with create log", err)
			}

			response.Type = enums.LogSend.String()
			response.Data = log
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
