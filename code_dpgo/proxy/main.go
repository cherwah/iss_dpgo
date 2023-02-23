package main

import (
	"fmt"
	"math/rand"
	"proxy/proxy"
	"sync"
	"time"
)

type simulate struct {
	mitm *proxy.Proxy
}

func main() {
	sim := simulate{
		mitm: proxy.NewProxy(),
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	sim.start(&wg)

	wg.Wait()
}

func (sim *simulate) start(wg *sync.WaitGroup) {
	fmt.Println("Simulation started.")

	go sim.add_likes(wg)
	go sim.get_likes(wg)
}

func (sim *simulate) add_likes(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		// randomly generating a post_id of 1 to 10
		post_id := rand.Intn(11) + 1

		// adding likes via our proxy
		sim.mitm.AddLike(post_id)

		time.Sleep(200 * time.Millisecond)
	}
}

func (sim *simulate) get_likes(wg *sync.WaitGroup) {
	defer wg.Done()

	var nlikes int

	for i := 0; i < 100; i++ {
		// randomly generating a post_id of 1 to 10
		post_id := rand.Intn(11) + 1

		// getting likes via our proxy
		nlikes = sim.mitm.GetLikes(post_id)

		fmt.Printf("Post %d has %d Likes\n", post_id, nlikes)

		time.Sleep(200 * time.Millisecond)
	}
}
