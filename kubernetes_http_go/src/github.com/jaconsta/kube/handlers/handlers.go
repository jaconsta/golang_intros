package handlers

import (
  "log"
  "time"
  "sync/atomic"
  
  "github.com/gorilla/mux"
)

// Router
func Router() *mux.Router {
  isReady := &atomic.Value{}
  isReady.Store(false)

  go func() {
		log.Printf("Readyz probe is negative by default...")
    // not any sense to wait for 10 seconds, but you might want to add here cache warming
		time.Sleep(1 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

  r := mux.NewRouter()
  r.HandleFunc("/home", home).Methods("GET")
  r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))

  return r
}
