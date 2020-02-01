package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const baseURL = "https://api.clashroyale.com"

func main() {
	log.Println("Started...")

	// allow query string

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s%s", baseURL, r.URL.EscapedPath())
		log.Printf("GET: %s \n", url)

		token := os.Getenv("TOKEN")

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal("error creating new request")
			return
		}

		req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "api unavailable")
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("processing body")
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		fmt.Fprintf(w, "%q", body)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
