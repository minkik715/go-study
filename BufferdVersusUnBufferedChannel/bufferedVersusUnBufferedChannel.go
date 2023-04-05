package BufferdVersusUnBufferedChannel

import (
	"fmt"
	"math/rand"
	"time"
)

var hamburgerMaterial = [...]string{"소세지", "치킨", "불고기", "새우", "계란", "야채", "민트초코"}

func orderHanBurger(c chan<- string) {
	for i := range [5]int{} {
		material := hamburgerMaterial[rand.Intn(7)]
		orderNumber := i + 1
		fmt.Printf("order %s buger. order number is %d\n", material, orderNumber)
		c <- fmt.Sprintf("make %s burger. order number is %d\n", material, orderNumber)
		fmt.Printf("%s buger is served. order number is %d\n", material, orderNumber)
	}
	close(c)
}

func makeHamburger(c <-chan string) {
	for {
		time.Sleep(time.Second * 5)
		order, ok := <-c
		println(order)
		if !ok {
			break
		}
	}
}

func Start() {
	unBufferChannel := make(chan string)
	bufferChannelSize5 := make(chan string, 5)
	fmt.Println("start unBuffer order Hamburger")
	go orderHanBurger(unBufferChannel)
	makeHamburger(unBufferChannel)
	fmt.Println("finish unBuffer order Hamburger")
	fmt.Println()
	fmt.Println("start Buffer order Hamburger")
	go orderHanBurger(bufferChannelSize5)
	makeHamburger(bufferChannelSize5)
	fmt.Println("finish Buffer order Hamburger")
}
