package models

import "time"

type RssItem struct {
	Title       string    `json:"title,omitempty"`
	Source      string    `json:"source,omitempty"`
	SourceURL   string    `json:"source_url,omitempty"`
	Link        string    `json:"link,omitempty"`
	PublishDate time.Time `json:"publish_date,omitempty"`
	Description string    `json:"description,omitempty"`
}
