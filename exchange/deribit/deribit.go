package deribit

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type DERIBIT struct {
	Client      *http.Client
	key         string
	secret      string
	UsedAccount string
	Subaccounts []string
	conn        *websocket.Conn
}
