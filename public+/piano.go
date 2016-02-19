package main

import (
	"fmt"
)

func main(){
	buttonPressChan := make(chan int)
	wait := make(chan bool)



	go func(){
		buttonPressChan <- 3
	}()

	go func() {
		for {
	    select {
		    case <-buttonPressChan:
		        go func(){
					fmt.Println("hi")
					return
				}()
		    }
		}
		wait <- false
	}()

	<-wait
}
