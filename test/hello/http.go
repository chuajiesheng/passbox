package passbox

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/test/hello", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
