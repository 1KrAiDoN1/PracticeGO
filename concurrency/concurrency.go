package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Worker(wg1 *sync.WaitGroup, id int, chwork chan int, chres chan int) {
	for k := range chwork {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("worker: ", id)
		chres <- k * k * k
	}

}
