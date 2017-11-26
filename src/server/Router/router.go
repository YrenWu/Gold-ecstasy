package Router 

import (
	"net/http"
	"os"
	"fmt"
	"net/http/httputil"
)

func HomeHandler(handler http.Handler) http.Handler {
  return http.HandlerFunc(
    func(send http.ResponseWriter, request *http.Request) {
      send.Write([]byte("<h1>Ahoy !!</h1>"))
      handler.ServeHTTP(send, request)
    })
}

func InfoHandler(handler http.Handler) http.Handler {
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