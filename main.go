package main

import (
	"fmt"

	semioptimizedtrieed "github.com/Shahriar-Sazid/go-random-test/semioptimized_trie_ed"
	"github.com/Shahriar-Sazid/go-random-test/trie"
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

	// ed.TestEDIndividual()
	// triefuzz.TestTrieFuzzz()
	// for i := 0; i < 10; i++ {
	// 	trie.TestTrieFuzz()
	// triefuzz.TestTrieFuzz()
	// }
	word := "moymensing"
	for i := 0; i < 10; i++ {
		trie.TestTrieFuzz(word)
		semioptimizedtrieed.TestTrieFuzz(word)
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
	// ed.SanityCheck()
}
