package philosophers

import (
	"log"
	"os"
	"time"
)

func RunPhilosophers(count int, dinnerTime time.Duration) {
	stopCh := make(chan bool)
	go startPhilosophers(count, stopCh)
	<-time.After(dinnerTime)
	close(stopCh)
}

var logger = log.New(os.Stdout, "", 0)
func startPhilosophers(count int, stopCh <-chan bool) {
	forks := make([]chan int, count)

	hasFood := true
	philosopher := func (i int) {
		for hasFood {
			var left int
			var right int
			if i == 0 {
				logger.Printf("%v: taking left\n", i)
				left = <-forks[i]
				logger.Printf("%v: taking right\n", i)
				right = <-forks[(i+1)%count]
			} else {
				logger.Printf("%v: taking right\n", i)
				right = <-forks[(i+1)%count]
				logger.Printf("%v: taking left\n", i)
				left = <-forks[i]
			}
			logger.Printf("%v: eating, left used: %v, right used: %v\n", i, left, right)
			<-time.After(100 * time.Microsecond)
			if i == 0 {
				logger.Printf("%v: putting right\n", i)
				forks[(i+1)%count] <- (right + 1)
				logger.Printf("%v: putting left\n", i)
				forks[i] <- (left + 1)
			} else {
				logger.Printf("%v: putting left\n", i)
				forks[i] <- (left + 1)
				logger.Printf("%v: putting right\n", i)
				forks[(i+1)%count] <- (right + 1)
			}
		}
	}

	for i := range forks {
		forks[i] = make(chan int, 1)
		forks[i]<-0
		go philosopher(i)
	}
	for <-stopCh {}
	hasFood = false
	for _, fork := range forks {
		<-fork
		<-fork
		close(fork)
	}
	return
}