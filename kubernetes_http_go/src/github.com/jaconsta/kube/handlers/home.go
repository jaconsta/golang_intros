package handlers

import (
  "fmt"
  "net/http"
)

// Home: Simple handler
func home(w http.ResponseWriter, _ *http.Request) {
  fmt.Fprint(w, "Hello! Your request was processed.")
}
