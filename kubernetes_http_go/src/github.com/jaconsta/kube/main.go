package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"

  "github.com/jaconsta/kube/handlers"
  "github.com/jaconsta/kube/version"
)


// src: https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/
func main()  {
  log.Print("Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release)

  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal("PORT is not net.")
  }

  router := handlers.Router()

  interrupt := make(chan os.Signal, 1)
  signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

  srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
  log.Print("The service is ready to listen and serve")

  killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}
