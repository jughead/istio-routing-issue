package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ServiceB struct {
	Url string
}

func (b *ServiceB) Call() (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", b.Url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		serviceB := &ServiceB{os.Getenv("SERVICE_B_URL")}
		res, err := serviceB.Call()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		if res != "2" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The result from service B is: " + res + "\n"))
	})

	fmt.Println("Starting server on port", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)

	if err != nil {
		fmt.Println(err)
	}
}
