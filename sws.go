package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// For syncing a slew of go-routines, if needed.
var wg sync.WaitGroup
var renderWaitGroup sync.WaitGroup

// Tracking web socket connections
var out1Clients = make(map[*websocket.Conn]io.Writer)
var out2Clients = make(map[*websocket.Conn]io.Writer)

var LOG_FILE_NAME string = "sws.log"
var BIND_ADDRESS string = ":8888"


func Output() {
	for {
    	time.Sleep(2 * time.Second)

		// Iterate over websocket out1Clients
		for clientSocket, _ := range out1Clients {
			write_err := clientSocket.WriteMessage(
				websocket.TextMessage,
				[]byte(fmt.Sprintf("Current Unix Time: %v\n", time.Now().Unix())))
			if write_err != nil {
			    log.Printf("Websocket error: %s", write_err)
			    clientSocket.Close()
			    delete(out1Clients, clientSocket)
			}
		}
	}
}


func Output2() {
	var runningTime int = 0

	for {
    	time.Sleep(10 * time.Second)
    	runningTime += 10
		// Iterate over websocket clients
		for clientSocket, _ := range out2Clients {
			write_err := clientSocket.WriteMessage(
				websocket.TextMessage,
				[]byte(fmt.Sprintf("run time: %d", runningTime)))
			if write_err != nil {
			    log.Printf("Websocket error: %s", write_err)
			    clientSocket.Close()
			    delete(out2Clients, clientSocket)
			}
		}
	}
}

// "index view" handler
func basePathHandler(w http.ResponseWriter, r *http.Request) {
	// startTime := time.Now().Unix()

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Failed processing index page template!")
		log.Printf("%q", err)
		return
	}
	t.Execute(w, nil)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}
	out1Clients[ws] = w
	log.Printf("WebSocket connection created: %s", r.RemoteAddr)
}

func webSocket2Handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}
	out2Clients[ws] = w
	log.Printf("WebSocket - Type 2 - connection created: %s", r.RemoteAddr)
}

func setupLogFileOrDie() {
	file, err := os.OpenFile(
		LOG_FILE_NAME,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.SetOutput(file)
	log.Print("-- sws started --")
}

func main() {
	defer fmt.Print("Exiting")
	setupLogFileOrDie()

//  We want the file server to handle paths for, well, static files
//  Think: javascript, CSS, images.
	fs := http.FileServer(http.Dir("."))
	http.Handle("/static/", fs)

	http.HandleFunc("/", basePathHandler)
	http.HandleFunc("/ws", webSocketHandler)
	http.HandleFunc("/notifications", webSocket2Handler)
	log.Printf("Starting web service, listening on %s\n", BIND_ADDRESS)

	go Output()
	go Output2()

	log.Printf("%q", http.ListenAndServe(BIND_ADDRESS, nil))
}
