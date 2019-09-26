package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var guess int
	var count int
	num := rand.Intn(100)
	for {
		fmt.Print("Guess number:")

		_,err := fmt.Scanf("%d",&guess)
		count++
		if err == nil {
			if guess > num {
				fmt.Println("Too high")
			}else if guess < num {
				fmt.Println("Too low")
			}else if guess == num {
				fmt.Println("you are right")
				break
			}else{
				fmt.Println("please give a number")
			}
		}
	}
}
