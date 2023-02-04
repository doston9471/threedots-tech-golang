package main
import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
  fmt.Fprint(w, "Hello, ", name)
}
