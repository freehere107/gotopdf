package until

import "net/http"

func ErrorNotify(url string) {
	http.Get(url + "?error=1")
}
