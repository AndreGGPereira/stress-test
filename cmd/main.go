package main

import (
	"flag"
	"log"

	stresstest "github.com/andreggpereira/stress-test"
)

func main() {

	var (
		url         = flag.String("url", "localhost:8080", "url  (default is localhost:8080)")
		requests    = flag.Int("requests", 1000, "requests  (default is 1000)")
		concurrency = flag.Int("concurrency", 10, "concurrency  (default is 10)")
	)
	flag.Parse()

	if *url == "" {
		*url = "localhost:8080"
	}

	if *requests == 0 {
		*requests = 1000
	}

	if *concurrency == 0 {
		*concurrency = 10
	}

	log.Println(" [*] Start service")

	s := stresstest.NewStressTest(*url, *requests, *concurrency)
	s.Execute()

}
