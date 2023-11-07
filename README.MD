# RSSReaderService

RSSReaderService is a service for parsing RSS feeds from multiple URLs and sending the parsed RSS items in response to a receiver. It utilizes the [rssreader](https://github.com/nobilik/rssreader) package for efficient RSS parsing.

## Usage

RSSReaderService listens to [http://localhost:{ENV[API_PORT]}/parse](http://localhost:{ENV[API_PORT]}/parse) for incoming requests with an array of RSS feed URLs in the following format:

```json
{
  "urls": ["url1", "url2", , , "urlN"]
}
```

The service then parses the provided URLs and sends a response containing the parsed RSS items.
As a result you'll have json array of RSS items

```go
type RssItem struct {
	Title       string    `json:"title,omitempty"`
	Source      string    `json:"source,omitempty"`
	SourceURL   string    `json:"source_url,omitempty"`
	Link        string    `json:"link,omitempty"`
	PublishDate time.Time `json:"publish_date,omitempty"`
	Description string    `json:"description,omitempty"`
}
```

API_PORT env var can also be set in .env file next to binary.

You can run it in docker:

```bash
docker-compose build
docker-compose up
```

### Package Structure

- handlers: Contains the HTTP request handlers for the service.
- utils: Provides utility functions for handling HTTP responses.
- services: Implements the ParseRSSFeeds function, which uses the rssreader package to parse RSS feeds.
- models: Defines the data structures used for requests and responses.

### License

This package is licensed under the MIT License.

## This package is created for testing purposes and not maintained
