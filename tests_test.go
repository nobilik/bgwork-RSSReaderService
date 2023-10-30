package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	requestJSON := []byte(`{"urls": [	"http://rss.cnn.com/rss/cnn_topstories.rss",
	"http://feeds.bbci.co.uk/news/rss.xml",
	"https://techcrunch.com/rssfeeds/",
	"https://arstechnica.com/rss-feeds/",
	"http://feeds.bbci.co.uk/news/world/rss.xml",
	"http://www.chinaminingmagazine.com/rss/current.xml"]}`)

	req, err := http.NewRequest("POST", "/parse", bytes.NewBuffer(requestJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, recorder.Code)
	}

}
