package pack

type Paging struct {
	NextOffset int
	BatchSize  int
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var bucketSize int = 7

func PageTest(paging Paging) ([]int, error) {
	res := make([]int, 0, paging.BatchSize)

	db := make([]int, 112)
	for i := 0; i < 112; i++ {
		db[i] = i + 1
	}

	fromBucket := paging.NextOffset / bucketSize
	toBucket := (paging.NextOffset + paging.BatchSize) / bucketSize
	fromOffset := paging.NextOffset % bucketSize
	toOffset := (paging.NextOffset + paging.BatchSize) % bucketSize
	if toOffset == 0 {
		toOffset = bucketSize
		toBucket--
	}

	for i := fromBucket; i <= toBucket; i++ {
		temp := make([]int, 0, Min(bucketSize, paging.BatchSize))

		if i*bucketSize >= len(db) {
			break
		}
		temp = append(temp, db[i*bucketSize:Min(i*bucketSize+bucketSize, len(db))]...)

		switch i {
		case fromBucket:
			res = append(res, temp[fromOffset:Min(fromOffset+paging.BatchSize, len(temp))]...)
		case toBucket:
			res = append(res, temp[:Min(toOffset, len(temp))]...)
		default:
			res = append(res, temp...)
		}

	}

	return res, nil
}
