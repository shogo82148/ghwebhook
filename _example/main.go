package main

import (
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/shogo82148/ghwebhook"
)

func main() {
	h := &ghwebhook.Webhook{
		Secret: "very-secret-string",
		Ping: func(e *github.PingEvent) {
			log.Printf("%#v", e)
		},
	}
	http.ListenAndServe(":8080", h)
}
