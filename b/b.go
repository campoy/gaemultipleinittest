package b

import (
	"fmt"
	"net/http"
)

const Path = "/b"

func init() {
	http.HandleFunc(Path, sayB)
}

func sayB(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi, B")
}
