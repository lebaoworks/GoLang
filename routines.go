package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func send_num(channel chan  int){
	for i:=0; i<20; i++ {
		channel <- i
	}
	//Close channel to make "i:=range channel" stop waiting
	close(channel)
	fmt.Println("Add Done!")
}
func print_out(x int, chanel chan int) {
	defer wg.Done()

	for i:=range chanel {
		fmt.Println(x, " ", i)
	}
	
	fmt.Println(x, "Done!")
}

func main() {
	ch := make(chan int, 20)
	wg.Add(2)	
	go send_num(ch)
	go print_out(1, ch)
	go print_out(2, ch)

	wg.Wait()
	fmt.Println("All Done!")
}
