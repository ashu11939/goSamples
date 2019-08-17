package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

type logWritter struct {

}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// fmt.Println(resp)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWritter{}
	io.Copy(lw, resp.Body)
}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote len : ", len(bs))
	return len(bs), nil
} 