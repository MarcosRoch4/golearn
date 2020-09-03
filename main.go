package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

//var mu sync.Mutex

func main() {

	fmt.Println("CPUS:", runtime.NumCPU())
	//	fmt.Println("Goroutine:", runtime.NumGoroutine())
	var contador int64
	contador = 0
	totalGoroutine := 15

	wg.Add(totalGoroutine)

	for i := 0; i < totalGoroutine; i++ {
		go func() {
			//			mu.Lock()
			//	v := contador
			runtime.Gosched()
			atomic.AddInt64(&contador, 1)
			//		v++
			//		contador = v
			fmt.Println("Contador: \t", atomic.LoadInt64(&contador))
			//			mu.Unlock()
			wg.Done()
		}() // por ser uma função anônima é necessário incluir os () no final da função

	}

	wg.Wait()

	fmt.Println(contador)

}
