package ws

import (
	"encoding/json"
	"github.com/aetherteam/logger_center/internal/enums"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/websocket"
	"log"
)

type Message struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type WSSignInWrapperDTO struct {
	Data WSSignInDTO `json:"data"`
}

type WSSignInDTO struct {
	Secret string `json:"secret"`
}

func (wssDTO *WSSignInDTO) Validate() error {
	return validation.ValidateStruct(
		wssDTO,
		validation.Field(&wssDTO.Secret, validation.Required),
	)
}

func WSHandler(store store.Store) gin.HandlerFunc {
	accountServiceStore := store.ServiceAccount()
	projectStore := store.Project()

	return func(ctx *gin.Context) {
		projectID := ctx.Query("project_id")

		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Fatal("Error upgrading connection:", err)
			return
		}
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				log.Println("Error close connection:", err)
				return
			}
		}(conn)

		ClientsMu.Lock()
		clients[conn] = &Client{
			Conn:      conn,
			ProjectID: projectID,
			IsAuth:    false,
		}
		ClientsMu.Unlock()

		for {
			msg := &Message{}
			_, p, err := conn.ReadMessage()
			if err != nil {
				return
			}

			if err := json.Unmarshal(p, &msg); err != nil {
				log.Println("error with unmarshal:", err)
				return
			}

			switch msg.Type {
			case enums.SignIn.String():
				data := WSSignInWrapperDTO{}
				if err := json.Unmarshal(p, &data); err != nil {
					log.Println("error with unmarshal:", err)
					return
				}

				validationErr := data.Data.Validate()
				if validationErr != nil {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "Must be authorization",
					}); err != nil {
						log.Println("error with write json:", err)
						return
					}

					ClientsMu.Lock()
					delete(clients, conn)
					ClientsMu.Unlock()
					return
				}

				project, projectErr := projectStore.FindById(projectID)
				if project == nil || projectErr != nil {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "project not found",
					}); err != nil {
						log.Println("error with write json:", err)
						return
					}

					return
				}

				aSecret, secretErr := accountServiceStore.FindBySecret(data.Data.Secret)
				if aSecret == nil || secretErr != nil {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "service account not found",
					}); err != nil {
						log.Println("error with write json:", err)
						return
					}
					return
				}

				if project.IsActive == false {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "Project is disabled",
					}); err != nil {
						log.Println("error with write json:", err)
						return
					}

					return
				}

				if aSecret.IsActive == false {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "Service account is disabled",
					}); err != nil {
						log.Println("error with write json:", err)
						return
					}
					return
				}

				clients[conn].IsAuth = true
				clients[conn].Secret = data.Data.Secret
				if err := conn.WriteJSON(&Message{
					Type: enums.Authorized.String(),
					Data: "Success",
				}); err != nil {
					log.Println("error with write json:", err)
					return
				}
			default:
				if !clients[conn].IsAuth {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "must be authorization",
					}); err != nil {
						log.Println("error with write json:", err)
						return
					}

					ClientsMu.Lock()
					delete(clients, conn)
					ClientsMu.Unlock()
					return
				}

				WSBroadcast(
					&Broadcast{
						Type: msg.Type,
						Data: msg.Data,
					},
					store,
					clients[conn],
				)
			}

		}
	}
}
