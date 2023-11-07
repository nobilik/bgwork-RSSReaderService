package unit_test

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestHandler(t *testing.T) {
// 	requestJSON := []byte(`{"urls": [	"http://rss.cnn.com/rss/cnn_topstories.rss",
// 	"http://feeds.bbci.co.uk/news/rss.xml",
// 	"https://techcrunch.com/rssfeeds/",
// 	"https://arstechnica.com/rss-feeds/",
// 	"http://feeds.bbci.co.uk/news/world/rss.xml",
// 	"http://www.chinaminingmagazine.com/rss/current.xml"]}`)

// 	req, err := http.NewRequest("POST", "/parse", bytes.NewBuffer(requestJSON))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	recorder := httptest.NewRecorder()

// 	handler(recorder, req)

// 	if recorder.Code != http.StatusOK {
// 		t.Errorf("Expected status %d, got %d", http.StatusOK, recorder.Code)
// 	}

// }

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/nobilik/bgwork/RSSReaderService/models"
	"github.com/nobilik/bgwork/RSSReaderService/utils"
)

var mockData = []models.RssItem{
	{
		Title:       "Mock Title 1",
		Source:      "Mock Source 1",
		SourceURL:   "https://example1.com/rss",
		Link:        "https://mock-link-1.com",
		PublishDate: time.Now(),
		Description: "Mock Description 1",
	},
	{
		Title:       "Mock Title 1",
		Source:      "Mock Source 1",
		SourceURL:   "https://example1.com/rss",
		Link:        "https://mock-link-1.com",
		PublishDate: time.Now(),
		Description: "Mock Description 1",
	},
	{
		Title:       "Mock Title 1",
		Source:      "Mock Source 1",
		SourceURL:   "https://example2.com/rss",
		Link:        "https://mock-link-1.com",
		PublishDate: time.Now(),
		Description: "Mock Description 1",
	},
}

func TestResponseJSON(t *testing.T) {
	rr := httptest.NewRecorder()

	utils.ResponseJSON(rr, mockData)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	var jsonResponse []models.RssItem
	err := json.NewDecoder(rr.Body).Decode(&jsonResponse)
	if err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	if len(jsonResponse) != len(mockData) {
		t.Errorf("Expected %d items, but got %d", len(mockData), len(jsonResponse))
	}

	for i := 0; i < len(mockData); i++ {
		if jsonResponse[i].Title != mockData[i].Title {
			t.Errorf("Item %d: Expected title %s, but got %s", i, mockData[i].Title, jsonResponse[i].Title)
		}
	}
}
