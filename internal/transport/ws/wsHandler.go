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
	Secret    string `json:"secret"`
	ProjectId string `json:"project_id"`
}

func (wssDTO *WSSignInDTO) Validate() error {
	return validation.ValidateStruct(
		wssDTO,
		validation.Field(&wssDTO.Secret, validation.Required),
		validation.Field(&wssDTO.ProjectId, validation.Required),
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
				fmt.Println(msg.Data)
				validationErr := data.Validate()
				_, projectErr := projectStore.FindById(data.ProjectId)
				_, secretErr := accountServiceStore.FindBySecret(data.Secret)

				if validationErr != nil || projectErr != nil || secretErr != nil {
					err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "must be authorization",
					})
					if err != nil {
						fmt.Println("error with write json")
					}

					ClientsMu.Lock()
					delete(clients, conn)
					ClientsMu.Unlock()
					return
				}

				clients[conn].IsAuth = true
				err := conn.WriteJSON(&Message{
					Type: enums.Authorized.String(),
					Data: "Success",
				})
				if err != nil {
					fmt.Println("error with write json")
				}

			default:
				if !clients[conn].IsAuth {
					err := conn.WriteJSON(&Message{
						Type: enums.Unverified.String(),
						Data: "must be authorization",
					})
					if err != nil {
						fmt.Println("error with write json")
					}

					ClientsMu.Lock()
					delete(clients, conn)
					ClientsMu.Unlock()
					return
				}

				WSBroadcast(&Broadcast{
					Type: msg.Type,
					Data: msg.Data,
				}, projectID)
			}

		}
	}
}
