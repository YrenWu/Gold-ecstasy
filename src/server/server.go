package main

import (
  "net/http"
  "fmt"
  "os"
  Daemon "gold-ecstasy/src/server/Daemon"
  Router "gold-ecstasy/src/server/Router"
)

func main() {
  server := http.NewServeMux()

  /* routes */
  server.Handle("/", Router.HomeHandler(http.Handler(&Daemon.RequestHandler{})))
  server.Handle("/info", Router.InfoHandler(http.Handler(&Daemon.RequestHandler{})))

  /* listen */
  err := http.ListenAndServe(":8080", server)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

