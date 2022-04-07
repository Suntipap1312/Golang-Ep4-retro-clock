package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	type placeholder [5]string
	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███"} //0

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███"} //1

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███"} //2

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███"} //3

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █"} //4

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███"} //5

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███"} //6

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █"} //7

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███"} //8

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███"} //9

	colon := placeholder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	screen.Clear()
	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		fmt.Println("")
		fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}

// very hard *************
