package Daemon

import (
	"net/http"
	"fmt"
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

