package main

import "os"

func main() {

    _, err := os.Create("/tmp/hello")
    
    // if we have an unexpected error
    if err != nil {
        panic(err)  // fast fail and dump the error to the screen
    }
}

