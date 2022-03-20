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
	c := make(chan string)
	n := 0
	for _, site := range os.Args[1:] {
		n++
		go count("http://"+site, c)
	}
	for i := 0; i < n; i++ {
		fmt.Print(<-c)
	}
	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}
func count(url string, c chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s: %s\n", url, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	dt := time.Since(start).Seconds()
	c <- fmt.Sprintf("%s %d [%.2fs]\n", url, n, dt)
}
