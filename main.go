package main

import (
	socket "ProtectMyBike/websocket"

	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
//func homePage(w http.ResponseWriter, r *http.Request) {
//	//http.ServeFile(w, r, "websockets.html")
//}
//
//
//// define a reader which will listen for
//// new messages being sent to our WebSocket
//// endpoint
//func reader(conn *websocket.Conn) {
//	for {
//		// read in a message
//		messageType, p, err := conn.ReadMessage()
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		// print out that message for clarity
//		fmt.Println(string(p))
//
//		if err := conn.WriteMessage(messageType, p); err != nil {
//			log.Println(err)
//			return
//		}
//
//	}
//}
//
//func wsEndpoint(w http.ResponseWriter, r *http.Request) {
//	//http.ServeFile(w, r, "websockets.html")
//	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
//
//	//upgrade this connection to a WebSocket
//	//connection
//	ws, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Println(err)
//	}
//	//
//	reader(ws)
//}

func setupRoutes() {
	http.HandleFunc("/", socket.Home)
	http.HandleFunc("/echo", socket.SocketWeb)
}


func main()  {

	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
