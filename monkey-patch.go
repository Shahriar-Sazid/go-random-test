package main

import (
	"fmt"
	"time"

	"bou.ke/monkey"
)

func monkey_patch1() {
	monkey.Patch(time.Now, func() (t time.Time) {
		return time.Date(2022, 9, 4, 10, 0, 0, 0, time.Local)
	})

	fmt.Println(time.Now())
}
