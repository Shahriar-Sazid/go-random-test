package main

import "github.com/Shahriar-Sazid/go-random-test/storage"

func main() {
	// fmt.Println(toCLikeExpression("cr > 2.7 and rating <34 AND abc = -5 Or def <= 8.9"))
	// httpConnectionTest()
	// httpConnectionTestMultiThread()
	// offSet, _ := strconv.ParseInt(os.Args[1], 10, 0)
	// batchSize, _ := strconv.ParseInt(os.Args[2], 10, 0)
	// fmt.Println(offSet, batchSize)
	// res, _ := pack.PageTest(pack.Paging{NextOffset: int(offSet), BatchSize: int(batchSize)})
	// fmt.Println(res)

	storage.GormTest2()
	// template.TemplateTest()
}
