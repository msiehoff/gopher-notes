package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	urls := []string {
		"http://google.com",
		"http://bing.com",
		"http://yahoo.com",
		"http://mlssoccer.com",
	}

	sizeCh := make(chan string)
	urlCh := make(chan string)

	for i := 0; i < 10; i++ {
		go pageSizeWorker(urlCh, sizeCh)
	}

	for _, url := range urls {
		// Essentially a load balancer
		// Using the queueing of work as a go routine
		// queueing of work will be done concurrently and won't
		// block the calling thread, thus work is given out as
		// workers become available, without blocking the calling process
		go generator(url, urlCh)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("\n%s", <- sizeCh)
	}

	fmt.Printf("\nAll Done!\n")
}

// Load balancing function
func generator(url string, urlCh chan string) {
	urlCh <- url
}

// worker function
// - Wait for input on urlCh
// - Send result on sizeCh
func pageSizeWorker(urlCh, sizeCh chan string) {

	for {
		url := <- urlCh
		pageSize, err := getPageSize(url)

		if err == nil {
			sizeCh <- fmt.Sprintf("%s size: %v", url, pageSize)
		} else {
			sizeCh <- fmt.Sprintf("error getting %s: %v", url, err)
		}
	}
}

// go doc net/http Response | more
// worker function
func getPageSize(url string) (int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, nil
	}

	return len(body), nil
}