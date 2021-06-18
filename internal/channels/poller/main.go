package poller

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Resource struct {
	Url        string
	isPolling  bool
	lastPolled time.Time
}

type Resources struct {
	Data []*Resource
	Lock *sync.Mutex
}

func (r *Resource) poll(lock *sync.Mutex) string {
	lock.Lock()
	r.isPolling = true
	lock.Unlock()

	resp, err := http.Head(r.Url)

	lock.Lock()
	fmt.Println(r.Url + " - " + resp.Status)
	lock.Unlock()

	if err != nil {
		panic(err)
	}

	lock.Lock()
	r.lastPolled = time.Now()
	r.isPolling = false
	lock.Unlock()

	return resp.Status
}

func (resources *Resources) getLastPolledResource() *Resource {
	var targetResource *Resource
	for _, res := range resources.Data {
		resources.Lock.Lock()
		if res.isPolling == false {
			if targetResource == nil || targetResource.lastPolled.Before(res.lastPolled) {
				targetResource = res
			}
		}
		resources.Lock.Unlock()

	}
	return targetResource
}
func (res *Resources) poller() {
	for {
		targetResource := res.getLastPolledResource()
		if targetResource != nil {
			targetResource.poll(res.Lock)
		}
	}
}

func (res *Resources) Poll(theadNum int) {
	var wg sync.WaitGroup
	for i := 0; i < theadNum; i++ {
		wg.Add(1)
		go res.poller()
	}
	wg.Wait()
}
