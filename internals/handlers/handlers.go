package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var wsChan = make(chan WsPayload)

var clients = make(map[WebSocketConnection]string)
var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

var upgradeCopnnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebSocketConnection struct {
	*websocket.Conn
}

type wsJsonResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

type WsPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {

	ws, err := upgradeCopnnection.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Client Connected to endpoint")

	var response wsJsonResponse
	response.Message = `<h1>Connected to Client</h1?`
	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""
	err = ws.WriteJSON(response)

	if err != nil {
		log.Fatal(err)
	}

	go ListenForWS(&conn)

}

func ListenForWS(conn *WebSocketConnection) {

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	var payload WsPayload

	for {

		err := conn.ReadJSON(&payload)
		if err != nil {

		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}

}

func ListenToWsChannel() {

	var response wsJsonResponse

	for {
		e := <-wsChan

		response.Action = "Got Here"
		response.Message = fmt.Sprintf("SOme meesage and action was %s", e.Action)
		boradcastToAll(response)

		switch e.Action {

		case "username":
			fmt.Print(e.Username, "value of the username")
			clients[e.Conn] = e.Username
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			boradcastToAll(response)

		case "left":
			response.Action = "list_users"
			delete(clients, e.Conn)
			users := getUserList()
			response.ConnectedUsers = users
			boradcastToAll(response)

		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", e.Username, e.Message)
			boradcastToAll(response)

		}

	}

}

func getUserList() []string {

	var userList []string
	for _, x := range clients {
		if x != "" {

			userList = append(userList, x)
		}
	}

	sort.Strings(userList)
	return userList
}

func boradcastToAll(response wsJsonResponse) {

	for client := range clients {

		err := client.WriteJSON(response)
		if err != nil {
			log.Println("err", err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {

	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Fatal("Something inside home handler", err)

	}
}

func renderPage(w http.ResponseWriter, html string, data jet.VarMap) error {

	view, err := views.GetTemplate(html)
	if err != nil {
		log.Fatal("Something wrong while rendering", err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Fatal("Something went wrong while executing the page")
		return err
	}
	return nil
}
