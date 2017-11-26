package main

import (
  "net/http"
  "fmt"
  "os"
  "net/http/httputil"
)

type RequestHandler struct{}
/* Inspect requests */
func (handler RequestHandler) ServeHTTP(send http.ResponseWriter, request *http.Request) {
  inspect(request)
}

func inspect(request *http.Request) {
  fmt.Println(">>>  Request-Method", request.Method)
  fmt.Println(">>>  Request-Headers")
  fmt.Println("     User-Agent : ", request.Header["User-Agent"])
  fmt.Println("     Content-Type : ", request.Header["Content-Type"])
  fmt.Println("     Accept-Encoding : ", request.Header["Accept-Encoding"])
  fmt.Println("     Content-Length : ", request.Header["Content-Length"])
  fmt.Println("     Referer : ", request.Header["Referer"])
  fmt.Println(">>>  Request-Body ", request.Body)
}

func helloHandler(handler http.Handler) http.Handler {
  return http.HandlerFunc(
    func(send http.ResponseWriter, request *http.Request) {
      send.Write([]byte("<h1>Ahoy !!</h1>"))
      handler.ServeHTTP(send, request)
    })
}

func infoHandler(handler http.Handler) http.Handler {
  return http.HandlerFunc(
    func(send http.ResponseWriter, request *http.Request) {
    dump, err := httputil.DumpRequest(request, true)
    if err != nil {
      http.Error(send, fmt.Sprint(err), http.StatusInternalServerError)
      os.Exit(1)
    }
    fmt.Fprintf(send, "%q", dump)
  })
}

func main() {
  server := http.NewServeMux()

  /* routes */
  server.Handle("/", helloHandler(http.Handler(&RequestHandler{})))
  server.Handle("/info", infoHandler(http.Handler(&RequestHandler{})))

  /* listen */
  err := http.ListenAndServe(":8080", server)
  if err != nil {
    os.Exit(1)
  }
}







