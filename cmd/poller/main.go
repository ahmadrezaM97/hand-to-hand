package main

import (
	"sync"

	"github.com/ahmadrezam97/hand-to-hand/internal/channels/poller"
)

func main() {
	var resources poller.Resources
	var lock sync.Mutex
	resources.Lock = &lock

	urls := []string{
		"https://www.alef.ir/",
		"https://www.varzesh3.com/",
		"https://www.aparat.com/",
	}
	for _, url := range urls {
		resources.Data = append(resources.Data, &poller.Resource{Url: url})
	}
	resources.Poll(4)
}
