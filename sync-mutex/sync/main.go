package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()

	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock() // Bloquea lecturas
	b := balance
	lock.RUnlock()
	return b
}

func main() {
	// mutex de lectura y escritura
	// Solo 1 escribe
	// Muchos pueden leer

	var wg sync.WaitGroup
	var lock sync.RWMutex // Bloqueo lectura escritura

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance(&lock))

	// con el manejo de semaforo para controlar la escritura de balance
	// Se eliminan las advertencias de riesgo de carrera sobre los datos
}
