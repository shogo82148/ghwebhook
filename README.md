# ghwebhook

## USAGE

```go
package main

import (
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/shogo82148/ghwebhook"
)

func main() {
	h := &ghwebhook.Webhook{
		// recommend to set secret
		Secret:       "very-secret-string",

		// Restrict IP address
		RestrictAddr: true,
		TrustAddrs:   []string{"::1/128", "127.0.0.0/8"},

		Ping: func(e *github.PingEvent) {
			log.Printf("%#v", e)
		},
	}
	http.ListenAndServe(":8080", h)
}
```

## Related Projects

- [Konboi/ghooks](https://github.com/Konboi/ghooks)
  - [ghooks というGithubのWeb Hook Receiver をgolangで書いた](http://konboi.hatenablog.com/entry/2014/11/11/100000)
- [tkuchiki/ghooks-cmd-runner](https://github.com/tkuchiki/ghooks-cmd-runner)
  - [Github Webhook を受けて任意のスクリプトを実行するツール](http://tkuchiki.hatenablog.com/entry/2016/05/13/112151)
