package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//  go run main.go google.com facebook.com youtube.com twitter.com go.dev yahoo.com instagram.com

func main() {
	start := time.Now()
	for _, site := range os.Args[1:] {
		count("http://" + site)
	}
	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}

func count(url string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s: %s\n", url, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	dt := time.Since(start).Seconds()
	fmt.Printf("%s %d [%.2fs]\n", url, n, dt)
}
