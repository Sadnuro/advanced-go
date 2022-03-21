package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {
	fmt.Printf("Calculate Expensive Fibonacci for: %d\n", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {
	s.Lock.RLock() // Solo bloquea escritura
	exists := s.InProgress[job]
	if exists {
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()

		fmt.Printf("Waiting for response job: %d\n", job)
		resp := <-response
		fmt.Printf("Response done: %d\n", resp)

		return
	}
	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calculate Fibonacci for: %d\n", job)
	result := ExpensiveFibonacci(job)

	s.Lock.RLock()
	pendingWorkers, exists := s.IsPending[job]
	s.Lock.RLock()

	if exists {
		for _, perdingWorker := range pendingWorkers {
			perdingWorker <- result
		}
		fmt.Printf("result send - All pendingWorkers Ready! job: %d\n", job)
	}

	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()

}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{2, 3, 5, 5, 4, 8, 8, 8}

	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}

	wg.Wait()
}
