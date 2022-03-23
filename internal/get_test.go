package internal

import (
	"testing"
	"time"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/stretchr/testify/require"

	"github.com/erdincmutlu/newsapi/types"
)

func TestStoreGetItems(t *testing.T) {
	fillCache()

	tests := []struct {
		name    string
		rss     types.Rss
		itemID  string
		rssItem *types.RSSItem
		err     error
	}{
		{
			name:    "ok",
			itemID:  "2",
			rssItem: &types.RSSItem{ID: "2", Title: "Title 2", Description: "This is description 2"},
		},
		{
			name:    "not found",
			itemID:  "3",
			rssItem: nil,
			err:     ttlcache.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fillCache()
			storeItems(test.rss)
			rssItem, err := getItem(test.itemID)
			require.Equal(t, test.err, err)
			require.Equal(t, test.rssItem, rssItem)
			cache.Close()
		})
	}
}

// Store some data to cache which will be used in tests
func fillCache() {
	setupCache()
	rss := types.Rss{
		Channel: types.Channel{
			Items: []types.RSSItem{
				{ID: "1", Title: "Title 1", Description: "This is description 1"},
				{ID: "2", Title: "Title 2", Description: "This is description 2"},
			},
		},
	}
	storeItems(rss)
}

// Set up cache for tests
func setupCache() {
	cache = ttlcache.NewCache()
	cache.SetTTL(time.Duration(10 * time.Second))
}
