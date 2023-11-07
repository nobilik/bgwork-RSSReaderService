package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nobilik/bgwork/RSSReaderService/models"
	"github.com/nobilik/bgwork/RSSReaderService/services"
	"github.com/nobilik/bgwork/RSSReaderService/utils"
)

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req models.Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		result, _ := services.ParseRSSFeeds(req.URLs...)

		utils.ResponseJSON(w, result)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
