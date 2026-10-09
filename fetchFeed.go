package main

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		errorS := fmt.Sprintf("Status CODE: %v", res.StatusCode)
		return nil, errors.New(errorS)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var rSSFeed RSSFeed
	if err = xml.Unmarshal(body, &rSSFeed); err != nil {
		return nil, err
	}
	rSSFeed.Channel.Title = html.UnescapeString(rSSFeed.Channel.Title)
	rSSFeed.Channel.Description = html.UnescapeString(rSSFeed.Channel.Description)
	for i := range rSSFeed.Channel.Item {
		item := &rSSFeed.Channel.Item[i]
		item.Description = html.UnescapeString(item.Description)
		item.Title = html.UnescapeString(item.Title)
	}
	return &rSSFeed, nil
}
