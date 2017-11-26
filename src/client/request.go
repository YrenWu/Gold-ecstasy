package main

import (
    "net/http"
    "fmt"
    "os"
    "io/ioutil"
    "bytes"
)

/* post form */
const inputname = "Foo"
const parameter = "Bar"
/* request headers */
const showheaders = false
const userAgent = "The Good, The Bad & The Ugly"

func main() {
    /* HTTP Client */
    client := &http.Client{}

    response, header := get("http://server:8080", client)
    fmt.Println(response)
    fmt.Printf("%s\n\n", header)

    response, header = post("http://server:8080/info", client)
    fmt.Println(response)
    fmt.Printf("%s\n\n", header)
}

/**
 * GET Method Request
 *
 */
func get(uri string, client *http.Client) (string, http.Header) {

    req, err := http.NewRequest("GET", uri, nil)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
    req.Header.Set("User-Agent", userAgent)
    resp, err := client.Do(req)

    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
    	content, header := handleResponse(resp)
    	if showheaders {
    	    return string(content), header
    	}
        return string(content), nil
    }
    return "", nil
}

/**
 * POST Method Request
 *
 */
func post(uri string, client *http.Client) (string, http.Header) {
    params := []byte(`{`+ inputname +`: `+ parameter +`}`)
    buffer := bytes.NewBuffer(params)

    req, err := http.NewRequest("POST", uri, buffer)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    req.Header.Set("User-Agent", userAgent)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        content, header := handleResponse(resp)
        if showheaders {
            return string(content), header
        }
        return string(content), nil
    }
    return "", nil
}

/**
 * Print errors if there is and read response text from response buffer
 *
 */
func handleResponse(resp *http.Response) (string, http.Header) {
    body := resp.Body
    header := resp.Header
    defer body.Close()
    content, err := ioutil.ReadAll(body)

    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
    return string(content), header
}
