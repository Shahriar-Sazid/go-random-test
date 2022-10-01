package pack

import (
	"fmt"
	"path"
)

func path_test() {
	url := "http://www.pathao.com/"
	paths := "api/v1/emails"

	fullPath := path.Join(url, paths)

	fmt.Println(fullPath)
}
