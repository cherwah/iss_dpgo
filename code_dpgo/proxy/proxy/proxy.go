package proxy

import (
	"fmt"
	"proxy/db"
	"sync"
	"time"
)

// caches the number of likes for a particular post_id
// also tracks the last sync with data source
type Proxy struct {
	data_src   db.Json_db
	post_likes map[int]int
	last_get   time.Time
	last_put   time.Time
}

var mutex sync.Mutex

func NewProxy() *Proxy {
	return &Proxy{
		data_src:   db.Json_db{Db_name: "likes.dat"},
		post_likes: map[int]int{}, // empty dictionary
		last_get:   time.Now(),
		last_put:   time.Now(),
	}
}

// going through the Proxy to add a Like to a post
func (p *Proxy) AddLike(post_id int) {
	mutex.Lock()         // only one thread can come in at a time
	defer mutex.Unlock() // unlock after end of function

	p.post_likes[post_id] += 1

	if time.Since(p.last_put).Seconds() > 3 {
		p.data_src.Save(&p.post_likes)
		p.last_put = time.Now()

		fmt.Println("Writing to storage.")
	} else {
		fmt.Println("Writing to cache.")
	}
}

// going through the Proxy to get Likes count for a post
func (p *Proxy) GetLikes(post_id int) int {
	mutex.Lock()         // only one thread can come in at a time
	defer mutex.Unlock() // unlock after end of function

	if time.Since(p.last_get).Seconds() > 3 {
		// consider cache is stale; re-fetch from data source
		p.data_src.Retrieve(&p.post_likes)
		p.last_get = time.Now()

		fmt.Println("Reading from storage.")
	} else {
		fmt.Println("Reading from cache.")
	}

	// use cache instead
	nlikes := p.post_likes[post_id]
	return nlikes
}
