package stressTest

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type StressTest struct {
	URL        string
	NumRequest int
	Jobs       int
}

type ResultTest struct {
	TimeTotal         time.Time
	NumRequest        int
	NumRequestSuccess int
	NumRequestFailed  int
	Mu                sync.Mutex
}

func (s *StressTest) Execute() {

	result := NewResultTest()

	requestsPerWorker := s.NumRequest / s.Jobs

	wg := sync.WaitGroup{}
	for i := 0; i < s.Jobs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerWorker; j++ {
				if err := s.doRequests(); err != nil {
					result.failed()
				} else {
					result.success()
				}
			}
		}()
	}

	wg.Wait()

	result.printResults(s.NumRequest, s.URL)

}

func NewStressTest(url string, numRequest, jobs int) StressTest {
	return StressTest{
		URL:        url,
		NumRequest: numRequest,
		Jobs:       jobs,
	}
}

func (s *StressTest) doRequests() error {

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(s.URL)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func NewResultTest() ResultTest {
	return ResultTest{
		TimeTotal:         time.Now(),
		NumRequest:        0,
		NumRequestSuccess: 0,
		NumRequestFailed:  0,
		Mu:                sync.Mutex{},
	}
}

func (s *ResultTest) success() {
	defer s.Mu.Unlock()
	s.Mu.Lock()
	s.NumRequestSuccess++
	s.NumRequest++
}

func (s *ResultTest) failed() {
	defer s.Mu.Unlock()
	s.Mu.Lock()
	s.NumRequestFailed++
	s.NumRequest++

}

func (s *ResultTest) printResults(numRequestTotal int, url string) {
	fmt.Printf("RELATÓRIO FINAL:\n")
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Total de requests: %d\n", numRequestTotal)
	fmt.Printf("Total de requests com sucesso: %d\n", s.NumRequestSuccess)
	fmt.Printf("Total de requests com falha: %d\n", s.NumRequestFailed)
	fmt.Printf("Tempo total de processamento: %v\n", time.Since(s.TimeTotal).Round(time.Second))

}
