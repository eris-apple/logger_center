package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type Client struct {
	Conn      *websocket.Conn
	ProjectID string `json:"project_id"`
	IsAuth    bool   `json:"is_auth"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients   = make(map[*websocket.Conn]*Client)
	ClientsMu sync.Mutex
)
