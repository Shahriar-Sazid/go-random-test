package main

import (
	"fmt"
	"time"

	routine "github.com/Shahriar-Sazid/go-random-test/goroutine"
	"github.com/robfig/cron/v3"
)

func main() {
	// fmt.Println(toCLikeExpression("cr > 2.7 and rating <34 AND abc = -5 Or def <= 8.9"))
	// httpConnectionTest()
	// httpConnectionTestMultiThread()
	// offSet, _ := strconv.ParseInt(os.Args[1], 10, 0)
	// batchSize, _ := strconv.ParseInt(os.Args[2], 10, 0)
	// fmt.Println(offSet, batchSize)
	// res, _ := pack.PageTest(pack.Paging{NextOffset: int(offSet), BatchSize: int(batchSize)})
	// fmt.Println(res)

	cr := cron.New()

	_, err := cr.AddFunc("@every 5s", routine.TestChildDieIfParentDie)
	if err != nil {
		fmt.Println("error occured")
	}

	time.Sleep(time.Second * 2)
}
