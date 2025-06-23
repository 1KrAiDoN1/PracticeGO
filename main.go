package main

import (
	"fmt"
	"helloapp/db"
	"helloapp/stepik"

	"time"

	//"io"

	//"helloapp/handler"
	"helloapp/concurrency"
	//"helloapp/httpserver"
	"helloapp/inter"
	"helloapp/payment"

	//"helloapp/stepik"
	"log"

	//"net/http"
	"strings"
	"sync"
)

func main() {

	in1 := make(chan int, 10)
	in2 := make(chan int, 10)
	out := make(chan int, 10)
	n := 10
	for i := 1; i < 11; i++ {
		in1 <- i
		in2 <- i
	}
	stepik.Merge2Channels(stepik.Fn, in1, in2, out, n)

	for v := range out {
		fmt.Println(v)
	}

	chanal := make(chan struct{})
	go func() {
		func() {
			fmt.Println("выполнилось")

		}()
		close(chanal) // в Горутине выполняем функцию и закрываем канал

	}()
	<-chanal // ждем пока канал закроется

	//httpserver.HandleFunc()

	// router := handler.NewRouter()
	// router.AddHandler(&handler.JSONHandler{})
	// log.Println("Server starting on :8080")
	// if err := http.ListenAndServe(":8080", router); err != nil {
	// 	log.Fatal(err)
	// }

	// req, err := http.Get("https://google.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// body, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer req.Body.Close()
	// fmt.Println(string(body))

	// client := http.Client{}
	// req1, err := http.NewRequest("GET", "https://google.com", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// body1, err := client.Do(req1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer body1.Body.Close()
	// q, _ := io.ReadAll(body1.Body)
	// fmt.Println(string(q))
	wg1 := &sync.WaitGroup{}

	ccc := make(chan int)
	go func() {
		ccc <- 1
	}()

	// time.Sleep(time.Second)
	go func() {
		ccc <- 2
	}()
	fmt.Println(<-ccc)
	fmt.Println(<-ccc)

	t := time.Now()
	cntwork := make(chan int, 15)
	cnres := make(chan int, 15)
	for i := 0; i < 15; i++ {
		cntwork <- i + 1
	}
	for i := 0; i < 5; i++ {

		go concurrency.Worker(wg1, i+1, cntwork, cnres)
	}
	close(cntwork)
	for i := 0; i < 15; i++ {
		fmt.Println("result: ", <-cnres)
	}
	fmt.Println(time.Since(t))

	ch1 := make(chan int)
	for i := 0; i < 10; i++ {
		wg1.Add(2)
		go func() {
			defer wg1.Done()
			ch1 <- 1
			fmt.Println("Запись в канал")
		}()

		go func() {
			defer wg1.Done()
			fmt.Println("Вывод из канала", <-ch1)
		}()
		wg1.Wait()
	}

	chData := make(chan int)
	chSync := make(chan struct{}) // Канал для синхронизации

	// Горутина для записи
	go func() {
		for i := 0; i < 10; i++ {
			chData <- i
			fmt.Println("Записано:", i)
			chSync <- struct{}{} // Сигнал, что запись завершена
		}
		close(chData)
	}()

	// Горутина для чтения
	go func() {
		for val := range chData {
			<-chSync // Ждём сигнала, что можно читать
			fmt.Println("Прочитано:", val)
		}
	}()

	user := &db.UserDB{
		Id:    1,
		Email: "test@gmail.com",
	}
	db.GetUserInfo(user)
	fmt.Println(user.Password)
	fmt.Println(user)

	// Способы задания анонимных функций
	fn := func(a int) int {
		return 1 + a
	}
	fmt.Println(fn(1))
	// объявляется анонимная функция
	func() {
		fmt.Println(111)
	}() // тут запускается

	var wg sync.WaitGroup
	ch := make(chan int, 1)
	wg.Add(1)
	go AppendData(ch, &wg, 100)
	wg.Wait()

	fmt.Println(<-ch)
	a := Pointer()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	var stringa = "777"
	ChangeStr(&stringa)
	fmt.Println(stringa)
	user1 := NewUser("pavel", 18, "aaaa")
	user1.age = 22
	fmt.Println(*user1)
	batteryForTest := &Battery{Bat: "110001101"}
	fmt.Println(batteryForTest.String())
	world := Some{99}
	world.ChangeS()
	fmt.Println(world.slovo)
	lamp := inter.NewLamp()
	fmt.Println("Яркость лампы:", lamp.Brightness)
	fmt.Println(inter.ControlDevice(lamp, "on"))
	fmt.Println(inter.ControlDevice(lamp, "status"))
	fmt.Println(inter.ControlDevice(lamp, "off"))
	fmt.Println(inter.ControlDevice(lamp, "status"))

	stripe := &payment.StripeProcessor{
		APIKey: "qwerty",
	}

	transactionID, err := payment.ProcessPayment(stripe, 100) // помещается структура, которая реализовывает интерфейс
	if err != nil {
		log.Fatalf("Payment failed: %v", err)
	}
	details, err := stripe.GetPaymentDetails(transactionID)
	if err != nil {
		log.Fatalf("Failed to get payment details: %v", err)
	}

	fmt.Printf("Payment details: %+v\n", details)

}

func AppendData(ch chan int, wg *sync.WaitGroup, cnt int) {
	ch <- cnt
	defer wg.Done()

}

type Battery struct {
	Bat string
}

type Some struct {
	slovo int
}

func (s *Some) ChangeS() {
	s.slovo++
	fmt.Println(s.slovo)
}
func (b *Battery) String() string {
	countpro := strings.Count(b.Bat, "0")
	countx := strings.Count(b.Bat, "1")
	return "[" + strings.Repeat(" ", countpro) + strings.Repeat("X", countx) + "]"
}

type BatteryEx interface {
	String()
}

func ChangeStr(str *string) {
	*str += "111"
	fmt.Println(*str)
}

func Pointer() func() (int, int) {
	cnt := 0
	b := 1
	return func() (int, int) {
		cnt++
		return cnt, b + 1
	}
}

func NewUser(name string, age int, email string) *User {
	return &User{
		name:  name,
		age:   age,
		email: email}
}

type User struct {
	name  string
	age   int
	email string
}
