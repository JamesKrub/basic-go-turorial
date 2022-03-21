package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//  go run guide.go google.com facebook.com youtube.com twitter.com go.dev yahoo.com instagram.com

// ***** without WaitGroups
//func main() {
//	start := time.Now()
//	n := 0
//	for _, site := range os.Args[1:] {
//		n++
//		go count("http://" + site)
//	}
//
//	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
//}

// ***** using WaitGroups
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(len(os.Args[1:]))
//
//	start := time.Now()
//	for _, site := range os.Args[1:] {
//		//wg.Add(1)
//		go func(s string) {
//			defer wg.Done()
//			count("http://" + s)
//		}(site)
//	}
//	wg.Wait()
//
//	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
//}

// ***** why must send data to anonymous function
//func main() {
//	var wg sync.WaitGroup
//
//	start := time.Now()
//	for _, site := range os.Args[1:] {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			count("http://" + site)
//		}()
//	}
//	wg.Wait()
//
//	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
//}

func count(url string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("curl got err, ", err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	defer r.Body.Close()

	dt := time.Since(start).Seconds()
	fmt.Printf("%s %d [%.2fs]\n", url, n, dt)
}
