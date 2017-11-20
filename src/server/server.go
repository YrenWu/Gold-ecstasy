package main

import (
  "net/http"
  "fmt"
  "os"
  "net/http/httputil"
)

func listen(send http.ResponseWriter, request *http.Request) {
  dump, err := httputil.DumpRequest(request, true)
  if err != nil {
    http.Error(send, fmt.Sprint(err), http.StatusInternalServerError)
    os.Exit(1)
  }
  inspect(request)
  fmt.Fprintf(send, "%q", dump)
}

func main() {
  http.HandleFunc("/", listen)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    os.Exit(1)
  }
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