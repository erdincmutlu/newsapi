package types

import (
	"errors"
)

const (
	ProviderBBC  = "bbc"
	ProviderSky  = "sky"
	CategoryUK   = "uk"
	CategoryTech = "tech"
)

// Rss is the RSS feed output
// Normally RSS has many other fields, but these fields should suffice
type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []RSSItem `xml:"item" json:"item"`
}

type RSSItem struct {
	ID          string `json:"id"`
	Title       string `xml:"title" json:"title"`
	Description string `xml:"description" json:"description"`
	Link        string `xml:"link" json:"link"`
	Guid        Guid   `xml:"guid" json:"-"`
	PubDate     string `xml:"pubDate" json:"pubDate"`
}

type Guid struct {
	Text string `xml:",chardata"`
}

type ShareRequest struct {
	ID        string `json:"id"`
	Action    string `json:"action"`
	Recipient string `json:"recipient"`
}

const (
	ActionEmail = "email"
	ActionTwit  = "twit"
)

var ErrInvalidInparams = errors.New("invalid input params")
