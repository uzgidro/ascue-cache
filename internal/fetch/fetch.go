package fetch

import (
	"ascue/internal/redisstore"
	"io"
	"log"
	"net/http"
	"time"
)

func StartFetcher(urls []string, keys []string, interval time.Duration, store redisstore.Store) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			for index, url := range urls {
				key := keys[index]
				FetchAndStore(url, key, store)
			}
			<-ticker.C
		}
	}()
}

func FetchAndStore(url string, key string, store redisstore.Store) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Fetch error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	err = store.Set(key, body)
	if err != nil {
		log.Println("Redis set error:", err)
	}
}
