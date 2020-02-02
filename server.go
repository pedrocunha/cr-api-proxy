package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const baseURL = "https://api.clashroyale.com"

func withBasicAuth(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, pass, _ := r.BasicAuth()

		masterPass := os.Getenv("PASSWORD")
		if pass != masterPass {
			http.Error(w, "Unauthorized.", 401)
			return
		}
		fn(w, r)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	var (
		token    = os.Getenv("TOKEN")
		fixieURL = os.Getenv("FIXIE_URL")
		client   = &http.Client{}
		endpoint = fmt.Sprintf("%s%s", baseURL, r.URL.EscapedPath())
	)

	log.Printf("GET %s \n", endpoint)

	if len(fixieURL) > 0 {
		proxyURL, _ := url.Parse(fixieURL)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal("error creating new request")
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func main() {
	log.Println("Started...")

	// TODO: allow query string
	http.HandleFunc("/", withBasicAuth(handle))

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
