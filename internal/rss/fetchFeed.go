package internal

import (
	"context"
	"encoding/xml"
	"html"

	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Set("User-Agent", "gator")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	new_feed := RSSFeed{}
	err = xml.Unmarshal(dat, &new_feed)
	if err != nil {
		return &RSSFeed{}, err
	}
	new_feed.Channel.Title = html.UnescapeString(new_feed.Channel.Title)
	new_feed.Channel.Description = html.UnescapeString(new_feed.Channel.Description)
	for i, item := range new_feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		new_feed.Channel.Item[i] = item
	}

	return &new_feed, nil
}
