package fetch

import (
	"ascue/internal/redisstore"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func Launch(urls []string, keys []string, interval time.Duration, store redisstore.Store, client *http.Client) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			var wg sync.WaitGroup

			for index, url := range urls {
				wg.Add(1)
				go func(u, k string) {
					defer wg.Done()
					data, err := GetData(client, u)
					if err != nil {
						return
					}

					if !IsJSON(data) {
						log.Printf("Invalid JSON response from %s", u)
						return
					}

					storeError := store.Set(k, data)
					if storeError != nil {
						log.Println("Redis set error:", storeError)
					}
				}(url, keys[index])
			}

			wg.Wait()
			<-ticker.C
		}
	}()
}

func GetData(client *http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Fetch error for URL %s: %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err := fmt.Errorf("bad status code from %s: %d", url, resp.StatusCode)
		log.Println(err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error for URL %s: %v", url, err)
		return nil, err
	}

	return body, nil
}

func IsJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}
