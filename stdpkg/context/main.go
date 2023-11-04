package main

import (
	"context"
	"fmt"
	"time"
)

/*
	c := context.Background()
	c = context.TODO()

	context.WithValue()
	context.WithCancel()
	context.WithDeadline()
	context.WithTimeout()

if parent context is cancelled => child is cancelled

	c.Done()
		this channel will be closed when (context is timedout, manually cancelled)
	c.Err()
		nil 	-> c.Done() is not closed
		context.DeadlineExceeded -> c.Done() is closed bcz of timeout/deadline
		context.Canceled -> c.Done() is closed due to manual cancel

*/

func main() {

	cSimple := context.WithValue(context.Background(), "name", "shiva")
	fmt.Println(cSimple.Value("key").(string)) //values are available even after the context is cancelled

	cParent, cancelParent := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancelParent()
	cChild, cancelChild := context.WithTimeout(cParent, 10*time.Second)
	defer cancelChild()

	t0 := time.Now()
	<-cChild.Done()
	fmt.Println("time taken : ", time.Since(t0).Seconds())
}
