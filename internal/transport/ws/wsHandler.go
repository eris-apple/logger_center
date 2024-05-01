package ws

import (
	"encoding/json"
	"fmt"
	"github.com/aetherteam/logger_center/internal/enums"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
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
			fmt.Println("Error upgrading connection:", err)
			return
		}
		defer conn.Close()

		ClientsMu.Lock()
		clients[conn] = &Client{
			Conn:      conn,
			ProjectID: projectID,
			IsAuth:    false,
		}
		ClientsMu.Unlock()

		for {
			msg := &Message{}
			if err := conn.ReadJSON(&msg); err != nil {
				fmt.Println("Error reading message:", err)
				ClientsMu.Lock()
				delete(clients, conn)
				ClientsMu.Unlock()
				break
			}

			marshalData, _ := json.Marshal(msg.Data)

			var data WSSignInDTO
			if err := json.Unmarshal(marshalData, &data); err != nil {
				fmt.Println("error with unmarshal")
			}

			switch msg.Type {
			case enums.SignIn.String():
				validationErr := data.Validate()
				if validationErr != nil {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "Must be authorization",
					}); err != nil {
						fmt.Println("error with write json")
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
						fmt.Println("error with write json")
					}

					return
				}

				aSecret, secretErr := accountServiceStore.FindBySecret(data.Secret)
				if aSecret == nil || secretErr != nil {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "service account not found",
					}); err != nil {
						fmt.Println("error with write json")
					}

					return
				}

				if project.IsActive == false {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "Project is disabled",
					}); err != nil {
						fmt.Println("error with write json")
					}

					return
				}

				if aSecret.IsActive == false {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "Service account is disabled",
					}); err != nil {
						fmt.Println("error with write json")
					}

					return
				}

				clients[conn].IsAuth = true
				clients[conn].Secret = data.Secret
				if err := conn.WriteJSON(&Message{
					Type: enums.Authorized.String(),
					Data: "Success",
				}); err != nil {
					fmt.Println("error with write json")
				}

			default:
				if !clients[conn].IsAuth {
					if err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "must be authorization",
					}); err != nil {
						fmt.Println("error with write json")
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
