package fetch

import (
	"ascue/internal/redisstore"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func Launch(urls []string, keys []string, interval time.Duration, store redisstore.Store) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			var wg sync.WaitGroup

			for index, url := range urls {
				wg.Add(1)
				go func(u, k string) {
					defer wg.Done()
					data, err := GetData(u)
					if err == nil {
						storeError := store.Set(k, data)
						if storeError != nil {
							log.Println("Redis set error:", err)
						}
					}
				}(url, keys[index])
			}
			wg.Wait()
			<-ticker.C
		}
	}()
}

func GetData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Fetch error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error:", err)
		return nil, err
	}

	return body, nil
}
