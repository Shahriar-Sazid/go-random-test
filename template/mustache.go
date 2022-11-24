package template

import (
	"fmt"

	"github.com/cbroglie/mustache"
)

func MustacheTest() {
	data, err := mustache.Render(`hello 
	{{#cr>5}}
	Well, {{taxed_value}} dollars, after taxes.
	{{/cr>5}}
	{{c}}`, map[string]string{"c": "world", "cr": "6"})
	if err != nil {
		fmt.Println("error rendering mustache", err)
	}

	fmt.Println("print data: ", data)
}
