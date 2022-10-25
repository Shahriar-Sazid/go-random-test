package pack

import "fmt"

func ReplaceEqual() {
	str1 := "cr_7d <= 0.45 and rating_5d > 4.2 and ar_3d = 0"
	str2 := "cr_7d > 0.45 and rating_5d >= 4.2 and ar_3d == 0"
	str3 := "cr_7d <= 0.7 AND rating_5d <= 3.5 Or ar_15d = -3="

	fmt.Println(replaceEqual(str1))
	fmt.Println(replaceEqual(str2))
	fmt.Println(replaceEqual(str3))

}

func replaceEqual(str string) string {
	for i := 1; i < len(str); i++ {
		if str[i] == '=' {
			switch str[i-1] {
			case '<', '>', '=':
				continue
			}
			if i+1 < len(str) && str[i+1] == '=' {
				continue
			}

			str = str[0:i] + "==" + str[i+1:]
			i++
		}
	}
	return str
}
