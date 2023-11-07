package services

import (
	"log"

	"github.com/nobilik/bgwork/RSSReaderService/models"
	"github.com/nobilik/rssreader"
)

func ParseRSSFeeds(feedURLs ...string) ([]models.RssItem, []error) {
	result, errs := rssreader.Parse(feedURLs...)
	if len(errs) > 0 {
		log.Printf("parsing errors: %v", errs)
	}
	var rssItems []models.RssItem
	for _, r := range result {
		rssItems = append(rssItems, models.RssItem(r))
	}
	return rssItems, errs
}
