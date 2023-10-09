package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{23, 23, 1, 5, 58, 8, 8, -2, 44, 3, 0}

	//fmt.Printf("before : %v %d\n\n", a, len(a))

	//blubble slort

	fmt.Println(a)

	for i, _ := range a {

		totalIndex := len(a)
		lastIndex := totalIndex - i

		b := a[:lastIndex]

		for j, v := range b {

			var (
				bigNum   int
				smallNum int
			)

			if j == lastIndex-1 {
				break
			}

			if b[j] > b[j+1] {
				b[j] = b[j+1]
				b[j+1] = v

				bigNum = j
				smallNum = j + 1

			} else {

				bigNum = j + 1
				smallNum = j

			}

			fmt.Printf("\r[")

			for n, _ := range a {

				if n == bigNum {
					fmt.Print("\u001b[1m")
					fmt.Print("\u001b[31m")
				}

				if n == smallNum {
					fmt.Print("\u001b[32m")
				}

				if n > len(b)-1 {
					fmt.Print("\u001b[90m")
				}

				fmt.Printf("%v ", a[n])
				fmt.Print("\u001b[0m")
			}

			fmt.Printf("\b]")

			time.Sleep(time.Millisecond * 100)

		}

	}

	fmt.Printf("\r%v\n", a)

	//fmt.Printf("\nafter  : %v %d\n", a, len(a))

}
