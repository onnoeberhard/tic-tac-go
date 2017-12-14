package main

import (
	"fmt"
	"time"
	"math/rand"
)

var board [9]int
var turn = 0
var done = make(chan bool)

func play(player int) {
	for {
		if turn == player {
			var fields []int
			for i, p := range board {
				var s string
				switch p {
				case 1:
					s = "x"
				case 2:
					s = "o"
				default:
					s = " "
				}
				if i == 0 {
					fmt.Printf("\n %s ", s)
				} else if i != 0 && i%3 == 0 {
					fmt.Printf("\n-----------\n %s ", s)
				} else {
					fmt.Printf("| %s ", s)
				}
				if p == 0 {
					fields = append(fields, i)
				}
			}
			fmt.Println("\n\n***********")
			if len(fields) > 0 {
				rand.Seed(time.Now().UnixNano())
				board[fields[rand.Intn(len(fields))]] = player
				time.Sleep(300 * time.Millisecond)
				if player == 1 {
					turn = 2
				} else {
					turn = 1
				}
			} else {
				done <- true
				return
			}
		}
	}
}

func main() {
	go play(1)
	go play(2)
	turn = 1
	<-done
}
