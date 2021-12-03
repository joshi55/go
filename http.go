/* RZFeeser | Alata3 Research
   HTTP Server                   */

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    // The HandleFunc registers the handler function for the given URL pattern
    http.HandleFunc("/", HielloHandler)
    fmt.Println("Server started at port 8089")
    log.Fatal(http.ListenAndServe(":8089", nil))      // listens on TCP network address for incoming HTTP requests
}

// describes how to respond to the HTTP request
func HielloHandler(abc http.ResponseWriter, m *http.Request) {
	fmt.Println("printing header %v",m)
    fmt.Fprintf(abc, "Hello, there\n")
}

