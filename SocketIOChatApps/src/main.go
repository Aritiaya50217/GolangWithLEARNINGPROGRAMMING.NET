package main

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {

	server, err := socketio.NewServer(nil)
	if err != nil {
		fmt.Println(err)
	} else {
		server.On("connection", func(so socketio.Socket) {
			fmt.Println("on connection")
			// Create room
			so.Join("room1")
			so.On("chat message", func(msg string) {
				fmt.Println("emit: ", so.Emit("chat message", msg))
				so.BroadcastTo("room1", "chat message", msg)
			})
			so.On("disconnection", func() {
				fmt.Println("On Disconnection")
			})
		})
		server.On("error", func(so socketio.Socket, err error) {
			fmt.Println("error: ", err)
		})

		http.Handle("/socket.io/", server)

		// สร้าง func เพื่อเรียกดูหน้า index.html
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "views/chat/index.html")
		})
		err2 := http.ListenAndServe(":8080", nil)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Serving at http://localhost:8080...")
		}
	}
}
