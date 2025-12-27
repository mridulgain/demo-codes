package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func fetchData(url string) string {
	// simulate network call
	time.Sleep(1 * time.Second)
	return "data from " + url
}

func Test_asyncFetchData(t *testing.T) {
	urls := []string{"url1", "url2", "url3"}
	for _, url := range urls {
		go func(url string) {
			t := fetchData(url)
			fmt.Println(t)
		}(url)
	}
	time.Sleep(2 * time.Second) // without this sleep, test exists before the async calls finish and prints nothing
}

// go test -v -run Test_asyncFetchData
// === RUN   Test_asyncFetchData
// data from url1
// data from url2
// data from url3
// --- PASS: Test_asyncFetchData (2.00s)
// PASS
// ok      github.com/mridulgain/demo-codes/go     2.591s

func Test_asyncFetchDataWithChannel(t *testing.T) {
	urls := []string{"url1", "url2", "url3"}
	results := make(chan string, len(urls))
	for _, url := range urls {
		url := url // capture loop variable
		go func() {
			results <- fetchData(url)
		}()
	}
	for i := 0; i < len(urls); i++ {
		data := <-results
		t.Log(data) // all data is printed but in random order
	}
}

// go test -v -run Test_asyncFetchDataWithChannel
// === RUN   Test_asyncFetchDataWithChannel
//     async_func_test.go:52: data from url3
//     async_func_test.go:52: data from url1
//     async_func_test.go:52: data from url2
// --- PASS: Test_asyncFetchDataWithChannel (1.00s)
// PASS
// ok      github.com/mridulgain/demo-codes/go     1.190s

func Test_asyncFetchDataWithWaitGroup(t *testing.T) {
	urls := []string{"url1", "url2", "url3"}
	var wg sync.WaitGroup
	out := make([]string, len(urls))

	for i, url := range urls {
		i, url := i, url
		wg.Add(1)
		go func() {
			defer wg.Done()
			out[i] = fetchData(url)
		}()
	}

	wg.Wait()
	for _, s := range out {
		t.Log(s)
	}
}

// go test -v -run Test_asyncFetchDataWithWaitGroup
// === RUN   Test_asyncFetchDataWithWaitGroup
//     async_func_test.go:76: data from url1
//     async_func_test.go:76: data from url2
//     async_func_test.go:76: data from url3
// --- PASS: Test_asyncFetchDataWithWaitGroup (1.00s)
// PASS
// ok      github.com/mridulgain/demo-codes/go     1.549s
