package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)
// ab -n 5000 -c 10 http://localhost:1234/   
func main() {
	Handler := func(w http.ResponseWriter, req *http.Request) {
		sleep(5)
		sleep(10)
		fmt.Fprintf(w, "Sleep Profiling Test")
	}
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":1234", nil)
}

func sleep(sleepTime int) {
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println("Slept for ", sleepTime, " Milliseconds")
}
