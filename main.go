package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicln("upgrade:", err)
		return
	}

	defer func() {
		log.Println("Disconnect!")
		conn.Close()
	}()

	for {
		mtype, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Println("receive:", msg)
		err = conn.WriteMessage(mtype, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
func main() {
	http.HandleFunc("/socket", ChatHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8899", nil))
}
