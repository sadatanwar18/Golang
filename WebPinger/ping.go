package main

import(
	"fmt"
	"net/http"
	"sync"
)

func checkWebsite(url string, wg *sync.WaitGroup, result chan string){
	defer wg.Done()

	// send get request
	resp, err := http.Get(url)

	if err != nil{
		result <- fmt.Sprintf("%s is Down (%v)", url, err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		result <- fmt.Sprintf("%s is Up", url)
	}else {
		result <- fmt.Sprintf("%s returned status %d", url, resp.StatusCode)
	}
}


func main(){
	var wg sync.WaitGroup

	results := make(chan string, 10)

	websites :=[] string {
		"https://google.com",
		"https://golang.com",
		"https://github.com",
		"https://notawebsite.bhag",
	}

	for _, site := range websites{
		wg.Add(1)
		go checkWebsite(site, &wg, results)
	}

	go func(){
		wg.Wait()
		close(results)
	}()

	for result := range results{
		fmt.Println(result)
	}

	fmt.Println("All websites checked")
}


/*
Main goroutine
│
├─> for loop → starts multiple "checkWebsite" goroutines
│     ├─> each goroutine sends result into channel
│     └─> each goroutine calls wg.Done()
│
├─> anonymous goroutine (go func)
│     ├─> waits for wg.Wait() (counter = 0)
│     └─> closes results channel
│
└─> for res := range results → keeps receiving
      until channel is closed


This is a classic Go concurrency pattern called:

“Fan-out, Fan-in pattern”

Fan-out:

You launch multiple goroutines to perform independent tasks (like checking multiple websites).

Fan-in:

You collect results into a single channel — and once all tasks are done, you close the channel to signal completion.

| Step                | What happens                          |
| ------------------- | ------------------------------------- |
| `wg.Add(1)`         | Count how many goroutines are running |
| `go checkWebsite()` | Launch goroutine to check site        |
| `wg.Done()`         | Signal that one goroutine is done     |
| `wg.Wait()`         | Wait for all goroutines to finish     |
| `close(results)`    | Signal that no more results will come |
| `range results`     | Receive results until channel closes  |

*/