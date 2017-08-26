package ghwebhook

import (
	"io/ioutil"
	"mime"
	"net/http"

	"github.com/google/go-github/github"
)

type Webhook struct {
	Secret string
	Ping   func(e *github.PingEvent)
	Push   func(e *github.PushEvent)
}

func (h *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	t, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	var payload []byte
	switch t {
	case "application/x-www-form-urlencoded":
		payload = []byte(r.PostFormValue("payload"))
	case "application/json":
		if h.Secret != "" {
			payload, err = github.ValidatePayload(r, []byte(h.Secret))
		} else {
			payload, err = ioutil.ReadAll(r.Body)
		}
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	e, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	go h.handle(e)
	w.WriteHeader(http.StatusOK)
}

func (h *Webhook) handle(e interface{}) {
	switch e := e.(type) {
	case *github.PingEvent:
		if h.Ping != nil {
			h.Ping(e)
		}
	case *github.PushEvent:
		if h.Push != nil {
			h.Push(e)
		}
	}
}
