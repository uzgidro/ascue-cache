package resource

import (
	"fmt"
	"os"
	"strings"
)

func GetResources() (map[string]string, error) {
	urlsStr := os.Getenv("PING_TARGETS")
	keysStr := os.Getenv("PING_KEYS")

	if urlsStr == "" || keysStr == "" {
		return nil, fmt.Errorf("missing PING_TARGETS or PING_KEYS env")
	}

	urls := strings.Split(urlsStr, ",")
	keys := strings.Split(keysStr, ",")

	if len(urls) != len(keys) {
		return nil, fmt.Errorf("length mismatch between PING_TARGETS and PING_KEYS")
	}

	resources := make(map[string]string)
	for i := range urls {
		resources[keys[i]] = urls[i]
	}

	return resources, nil
}
