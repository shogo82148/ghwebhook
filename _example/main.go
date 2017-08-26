package main

import (
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/shogo82148/ghwebhook"
)

func main() {
	h := &ghwebhook.Webhook{
		Secret:       "very-secret-string",
		RestrictAddr: true,
		TrustAddrs:   []string{"::1/128", "127.0.0.0/8"},
		Ping: func(e *github.PingEvent) {
			log.Printf("%#v", e)
		},
	}
	http.ListenAndServe(":8080", h)
}
