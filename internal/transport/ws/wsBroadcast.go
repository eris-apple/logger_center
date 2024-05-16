package ws

import (
	"github.com/eris-apple/logger_center/internal/enums"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/services"
	"github.com/eris-apple/logger_center/internal/store"
	"github.com/mitchellh/mapstructure"
	"log"
	"time"
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
				log.Println("error with decode data")
				return
			}

			if body.Timestamp <= 0 {
				now := time.Now()
				body.Timestamp = now.Unix()
			}

			l := &models.Log{
				ProjectID: c.ProjectID,
				ChainID:   body.ChainID,
				Title:     body.Title,
				Error:     body.Error,
				Params:    body.Params,
				Content:   body.Content,
				Level:     body.Level,
				Timestamp: body.Timestamp,
			}

			if _, err := ls.Create(l); err != nil {
				log.Println("error with create log", err)
				return
			}

			response.Type = enums.LogSend.String()
			response.Data = l
		default:
			response.Type = "unexpected type"
			response.Data = nil
		}

		for _, c := range clients {
			if err := c.Conn.WriteJSON(response); err != nil {
				log.Println("Error sending message:", err)
				_ = c.Conn.Close()
				delete(clients, c.Conn)
			}
		}
	}
}
