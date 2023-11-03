package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nobilik/rssreader"
)

type request struct {
	URLs []string `json:"urls"`
}

func server() {
	http.HandleFunc("/parse", handler)

	fmt.Println("Server listening on " + os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		result := []rssreader.RssItem{}
		var errArr []error

		if len(req.URLs) == 0 {
			w.Header().Set("Content-Type", "application/json")
			fmt.Printf("%+v\n", result)
			json.NewEncoder(w).Encode(result)
			return
		}

		result, errArr = rssreader.Parse(req.URLs...)
		if len(errArr) > 0 {
			log.Printf("parsing errors: %v", errArr)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
