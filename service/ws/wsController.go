package ws

import (
	"gee"
	"log"
	"net/http"
)

//var wsUpgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}

func httpHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	//conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("wsUpgrader.Upgrad error!", err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func WsController(c *gee.Context, hub *Hub) {
	c.Status(http.StatusOK)
	httpHandler(hub, c.Writer, c.Req)
}
