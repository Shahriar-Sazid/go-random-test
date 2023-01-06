package pack

import (
	"fmt"
	"regexp"
	"strings"
)

func regex_test() {

	var htmlRegex *regexp.Regexp = regexp.MustCompile(`(?i)</?[\s\S]*>`)
	var inAppUrlRegex *regexp.Regexp = regexp.MustCompile(`^(pathao://|pathaoqa://|pathao-drive://|pathaoqa-drive//)`)

	testCases := []string{"<abc></abc>", "lsdjflsd", "pathao-drive://sdfl <dd>", "pathaoqa-drive// </df>"}
	for _, cse := range testCases {
		fmt.Printf("%v %v\n", cse, htmlRegex.MatchString(cse))
	}
	fmt.Println()
	testCases = []string{"pathao://dslf?lfd", "lsdjflsd", "pathao-drive://sdfl", "pathaoqa-drive//", "pathao://"}
	for _, cse := range testCases {
		fmt.Printf("%v %v\n", cse, inAppUrlRegex.MatchString(cse))
	}

	// r, _ := regexp.Compile("p([a-z]+)ch")

	// fmt.Println(r.MatchString("peach"))

	// fmt.Println(r.FindString("peach punch"))

	// fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// fmt.Println(r.FindStringSubmatch("peach punch"))

	// fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// fmt.Println(r.FindAllString("peach punch pinch", -1))

	// fmt.Println("all:", r.FindAllStringSubmatchIndex(
	// 	"peach punch pinch", -1))

	// fmt.Println(r.FindAllString("peach punch pinch", 2))

	// fmt.Println(r.Match([]byte("peach")))

	// r = regexp.MustCompile("p([a-z]+)ch")
	// fmt.Println("regexp:", r)

	// fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// in := []byte("a peach")
	// out := r.ReplaceAllFunc(in, bytes.ToUpper)
	// fmt.Println(string(out))
}

func Regex_test2() {
	// var featurePattern = regexp.MustCompile(`(?P<feature>\w+)_(?P<period>\d+\w)`)

	r := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
	fmt.Printf("%#v\n", r.FindStringSubmatch(`2015-05-27`))
	fmt.Printf("%#v\n", r.SubexpNames())
}

func regex_test3() {
	re := regexp.MustCompile(`BBB([^B]|B[^B]|BB[^B])*EEE`)
	fmt.Printf("%#v\n", re.FindAllString("BBB EEE BBB..BBB...EEE", -1))
}

func regex_test4() {
	var myExp = regexp.MustCompile(`(?i)(?P<and>and)|(?P<or>or)|(?P<equal>\b\s*=)`)
	match := myExp.FindAllString("cr > 2.7 and rating <34 AND abc = -5 Or def <= 8.9", -1)
	myExp.ReplaceAllStringFunc("cr > 2.7 and rating <34 AND abc = -5 Or def <= 8.9", nil)
	fmt.Println(match)
}

func toCLikeExpression(exp string) string {
	fmt.Println(exp)
	var re = regexp.MustCompile(`(?i)and|or|\b\s*=`)

	var replacements = map[string]string{
		"and": "&&",
		"or":  "or",
		"=":   "==",
	}

	return re.ReplaceAllStringFunc(exp, func(s string) string {
		return replacements[strings.Trim(strings.ToLower(s), " ")]
	})
}
