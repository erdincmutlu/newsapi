package internal

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ReneKroon/ttlcache/v2"

	"github.com/erdincmutlu/newsapi/internal/db"
	"github.com/erdincmutlu/newsapi/types"
)

func getNews(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()
	provider := urlParams.Get("provider")
	category := urlParams.Get("category")

	w.Header().Add("Content-Type", "application/json")

	feed, err := getFeed(provider, category)
	if err != nil {
		log.Printf("getFeed error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := getNewsList(feed)
	if err != nil {
		log.Printf("getNewsList error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("json encode error: %s", err.Error())
	}
}

func getFeed(provider string, category string) (string, error) {
	cacheKey := provider + "-" + category
	url, err := cache.Get(cacheKey)
	if err == nil {
		// Found in cache
		return url.(string), nil
	}

	if err != ttlcache.ErrNotFound {
		return "", err
	}

	// Not found in the cache
	feed, err := db.GetFeed(dbClient, provider, category)
	if err != nil {
		return "", err
	}
	cache.Set(cacheKey, feed)

	return feed, nil
}

func getNewsList(url string) (*types.Rss, error) {
	rssCached, err := cache.Get(url)
	if err == nil {
		// Found in cache
		cached := rssCached.(types.Rss)
		return &cached, nil
	}

	if err != ttlcache.ErrNotFound {
		return nil, err
	}

	// Not found in the cache. Let's request via url
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss types.Rss
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return nil, err
	}

	// Set IDs of articles
	for i := range rss.Channel.Items {
		index := strings.LastIndex(rss.Channel.Items[i].Guid.Text, "-")
		if index == -1 {
			log.Printf("The article does not have id\n")
		} else {
			rss.Channel.Items[i].ID = rss.Channel.Items[i].Guid.Text[index+1:]
		}
	}
	cache.Set(url, rss)
	storeItems(rss)

	return &rss, err
}

// Store each item to be able to access quickly for sharing
func storeItems(rss types.Rss) {
	for _, item := range rss.Channel.Items {
		// Keep the item in the cache a bit longer
		cache.SetWithTTL(item.ID, item, time.Hour)
	}
}

func getItem(id string) (*types.RSSItem, error) {
	itemCached, err := cache.Get(id)
	if err != nil {
		return nil, err
	}

	cached := itemCached.(types.RSSItem)
	return &cached, nil
}
