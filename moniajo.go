package main

import "fmt"
import "time"
import "math/rand"

func heiLauri(ender chan bool) {
	fmt.Println("heiLauri() alkoi")
	time.Sleep(3*time.Second)
	fmt.Println("heiLauri() loppui")
	ender <- true
}
func kikkelisKokkelis(ender chan bool) {
	fmt.Println("kikkelisKokkelis() alkoi")
	time.Sleep(3*time.Second)
	fmt.Println("kikkeliskokkelis() loppui")
	ender <- true
}
func tiuTau(ender chan bool) {
        fmt.Println("tiuTau() alkoi")
	s := []int{1,2,3,4,5,6,7}
	fmt.Println(rand.Intn(len(s)))
        ender <- true
}

func main() {
	ch := make(chan bool)
	fmt.Println("main alkoi")
	go heiLauri(ch)
	go kikkelisKokkelis(ch)
	go tiuTau(ch)
	<-ch
	<-ch
	<-ch
	fmt.Println("main() loppui")
}
