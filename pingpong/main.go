package main

import ("fmt"
"time")

func main() {
const rodadas = 5

pingCh := make(chan struct{})
pongCh := make(chan struct{})
fim := make(chan struct{})


// goroutine "ping"
go func() {
	for i := 0; i < rodadas; i++ {
		<-pingCh
		fmt.Println("ping")
		time.Sleep(500 * time.Millisecond)
		pongCh <- struct{}{}
	}
} ()

//goroutine "pong"
go func ()  {
	for i := 0; i < rodadas; i++ {
		<-pongCh
		fmt.Println("pong")
		time.Sleep(500 * time.Millisecond)
		if i == rodadas-1 {
			close(fim)
			return
		}
		pingCh <- struct{}{}
	}
} ()


pingCh <- struct{}{}
<-fim
}
