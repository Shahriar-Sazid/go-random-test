package pack

import (
	"fmt"
	"os"
	"time"
)

func time_test() {
	fmt.Println(time.Now().Add(time.Hour * 24 * -1).Format("2006-01-02"))
	fmt.Println(time.Now().Add(time.Hour * 9 * -1).Format("2006-01-02"))
	fmt.Println(time.Now().Add(time.Hour * 24 * 1).Format("2006-01-02"))
	os.Setenv("TZ", "UTC")
	utc, _ := time.LoadLocation("UTC")
	time.Local = utc
	fmt.Println(time.Now().Format(time.Layout))

}
