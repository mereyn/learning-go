package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	// fmt.Println(response)
	// real world development do not set 99999
	// bs := make([]byte, 99999)
	// response.Body.Read(bs)
	// fmt.Println(string(bs))

	// simplify and advance implementation
	// io.Copy(os.Stdout, response.Body)
	
	lw := logWriter{}
}

func (logWriter) Writer(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
