package pack

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func valuate_test() {
	expression, err := govaluate.NewEvaluableExpression("(cr_7d < 0.50 && rating_7d < 3.5) || ar_15d == 4")

	fmt.Println(err)
	if err != nil {
		return
	}

	parameters := make(map[string]interface{}, 8)
	parameters["cr_7d"] = 0.6
	parameters["rating_7d"] = 3
	parameters["ar_15d"] = 4

	result, err := expression.Evaluate(parameters)
	fmt.Println(result)
	fmt.Println(err)
	// result is now set to "false", the bool value.
}
