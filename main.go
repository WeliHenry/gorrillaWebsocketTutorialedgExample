package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func setupRoutes()  {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
	http.ListenAndServe(":8080", nil )
}


func main()  {
	fmt.Println("web socket started successfully")
	setupRoutes()
}

func reader(conn *websocket.Conn)  {
	for  {
		messageType, p, err:= conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(p))
		if err:= conn.WriteMessage(messageType, p); err!=nil{
			return
		}

	}
}

func homePage(w http.ResponseWriter, r * http.Request)  {
	http.ServeFile(w,r, "index.html")
}


func wsEndpoint(w http.ResponseWriter, r * http.Request)  {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err:= upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("client connected successfully")
	reader(ws)
}


