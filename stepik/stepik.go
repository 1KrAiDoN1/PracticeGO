package stepik

import (
	"sync"
	"time"
)

func RemoveDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
	cnt := ""
	for value := range inputStream {
		cnt1 := value
		if cnt != cnt1 {
			outputStream <- value
		}

		cnt = value

	}
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	outChan := make(chan int)
	go func() {
		defer close(outChan)
		select {
		case msg := <-firstChan:
			outChan <- (msg * msg)

		case msg := <-secondChan:
			outChan <- msg * 3

		case <-stopChan:

		}

	}()
	return outChan
}

func Fn(x int) int {
	time.Sleep(time.Millisecond * 100)
	return x
}
func Merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() { // ← Вся логика в отдельной горутине
		vals := make([]int, n)
		var wg sync.WaitGroup
		wg.Add(2 * n)

		for i := 0; i < n; i++ {
			x1 := <-in1 // Чтение (блокируется только в этой горутине)
			x2 := <-in2 // Чтение
			i := i      // Фиксируем значение i для горутин

			// Запускаем вычисления fn(x1) и fn(x2) параллельно
			go func() {
				vals[i] += fn(x1)
				wg.Done()
			}()
			go func() {
				vals[i] += fn(x2)
				wg.Done()
			}()
		}

		wg.Wait() // Ждём завершения всех вычислений

		// Отправляем результаты
		for _, v := range vals {
			out <- v
		}
		close(out) // Важно: закрываем канал после отправки
	}() // ← Запускаем горутину
}
