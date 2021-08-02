package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	r1 := http.NewServeMux()
	r1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to /")
		for name, values := range r.Header {
			log.Println(name, values)
		}
		w.Write([]byte("1"))
	})

	r2 := http.NewServeMux()
	r2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to /")
		for name, values := range r.Header {
			log.Println(name, values)
		}
		w.Write([]byte("2"))
	})

	go func() {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r1))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT2"), r2))
	}()

	select {}
}
