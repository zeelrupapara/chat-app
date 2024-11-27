package websocket

type Room struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

var hub = Room{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

func ManageRoom() {
    for {
        select {
        case client := <-hub.register:
            hub.clients[client] = true
        case client := <-hub.unregister:
            if _, ok := hub.clients[client]; ok {
                delete(hub.clients, client)
                close(client.send)
            }
        case message := <-hub.broadcast:
            for client := range hub.clients {
                select {
                case client.send <- message:
                default:
                    delete(hub.clients, client)
                    close(client.send)
                }
            }
        }
    }
}