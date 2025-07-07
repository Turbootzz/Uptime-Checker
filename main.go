package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := os.Args[1:]

	if len(urls) == 0 {
		fmt.Println("Please use go run main.go <url1> <url2> ...")
		return
	}

	interval := 30 * time.Second

	fmt.Printf("Starting uptime checker with an interval of %s...\n", interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	checkAll(urls)

	for range ticker.C {
		checkAll(urls)
	}
}

func checkURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("URL %s is reachable.\n", url)
	} else {
		fmt.Printf("URL %s returned status code: %d\n", url, resp.StatusCode)
	}
}

func checkAll(urls []string) {
	fmt.Println("----- New check -----")
	for _, url := range urls {
		checkURL(url)
	}
}
