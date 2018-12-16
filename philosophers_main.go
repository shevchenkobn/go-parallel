package main

import (
	"./philosophers"
	"time"
)

func main() {
	philosophers.RunPhilosophers(50, time.Second * 4)
}
