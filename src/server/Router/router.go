package Router 

import (
	"net/http"
	"os"
	"fmt"
	"net/http/httputil"
	"html/template"
)

const TEMPLATEPATH = "/home/ouombette/go/src/gold-ecstasy/src/server/Template/"

func HomeHandler(handler http.Handler) http.Handler {
  return http.HandlerFunc(
    func(send http.ResponseWriter, request *http.Request) {
      tmpl, err := template.ParseFiles(TEMPLATEPATH + "hello.html")
      if err != nil {
        http.Error(send, err.Error(), http.StatusInternalServerError)
        return
      }
      if err := tmpl.Execute(send, nil); err != nil {
        http.Error(send, err.Error(), http.StatusInternalServerError)
      }
      handler.ServeHTTP(send, request) // call internal handler
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
