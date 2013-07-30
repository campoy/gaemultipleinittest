package a

import (
	"fmt"
	"net/http"

	"github.com/campoy/bug/b"
)

func init() {
	http.HandleFunc("/", sayA)
}

func sayA(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `hi, A<br>B is in <a href="%v">`, b.Path)
}
