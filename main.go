package main

import (
	"github.com/gorilla/schema"
	"github.com/ssk2/whereisbot/robots"
	"io"
	"log"
	"net/http"
	"strconv"
)

var robot = new(robots.WhereIsBot)

func main() {
	http.HandleFunc("/slack", CommandHandler)
	StartServer()
}
func CommandHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		decoder := schema.NewDecoder()
		command := new(robots.SlashCommand)
		err := decoder.Decode(command, r.PostForm)
		if err != nil {
			log.Println("Couldn't parse post request:", err)
		}
		command.Command = command.Command[1:]
		w.WriteHeader(http.StatusOK)
		plainResp(w, robot.Run(command))
	}
}

func plainResp(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, msg)
}

func StartServer() {
	port := robots.Config.Port
	log.Printf("Starting HTTP server on %d", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("Server start error: ", err)
	}
}
