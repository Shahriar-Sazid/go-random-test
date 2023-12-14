package pack

import (
	"encoding/json"
	"fmt"
)

var jsonString = `
{
	"name": "{{.start_date}} KTM BIKE NEW VERIFIED - 1",
	"title": {
		"en": "{{.start_date}} KTM BIKE NEW VERIFIED - 1"
	},
	"body": {
		"en": "Take a look at this quest"
	},
	"start_time": "{{.start_time}}",
	"start_date": "{{.start_date}}",
	"end_time": "{{.end_time}}",
	"end_date": "{{.end_date}}",
	"duration_type": 2,
	"cities": [
		4
	],
	"vehicles": [
		1
	],
	"verticals": [
		"ride"
	],
	"type": "regular",
	"payout_type": "payout",
	"groups": [{
		"name": "A",
		"riders": [
			123, 12, 12, 12
		],
		"steps": [{
			"name": "1",
			"order_id": 1,
			"award": {
				"amount": 100
			},
			"rule": {
				"operator": "and",
				"conditions": [{
						"category": "count",
						"target": 10
					},
					{
						"category": "rate",
						"type": "completion",
						"target": 80
					}
				]
			}
		}]
	}],
	"notifications": [{
			"event": "activate",
			"after": 15,
			"types": [
				"in_app_notification",
				"push_notification"
			],
			"title": {
				"en": "पहिलो रु 200 कमाउने अवसर!!"
			},
			"body": {
				"en": "आदरणीय पठाओ साथिहरू, उत्कृष्ट बोनस Quest सुरु गरिएको छ | जस अन्तर्गत ७ दिन भित्र २५ राइड पुरा गरी थप रु २०० कमाउनुहोस | बोनस Quest मा भाग लिई मौकाको फाइदा उठाउनुहोस |"
			}
		},
		{
			"event": "start",
			"before": 1,
			"types": [
				"in_app_notification",
				"push_notification"
			],
			"title": {
				"en": "पहिलो रु 200 कमाउने अवसर!!"
			},
			"body": {
				"en": "आदरणीय पठाओ राइडर साथीहरु, उत्कृष्ट बोनस Quest सुरु हुदैछ, कृपया राइड लीइ थप् कमाउनुहोस्।"
			}
		},
		{
			"event": "start",
			"after": 1,
			"types": [
				"in_app_notification"
			],
			"title": {
				"en": "पहिलो रु 200 कमाउने अवसर!!"
			},
			"body": {
				"en": "आदरणीय पठाओ राइडर साथीहरु, उत्कृष्ट बोनस Quest सुरु भइसकेकाे छ, कृपया राइड लीइ थप् कमाउनुहोस्।"
			}
		},
		{
			"event": "end",
			"before": 10,
			"types": [
				"in_app_notification",
				"push_notification"
			],
			"title": {
				"en": "Quest End"
			},
			"body": {
				"en": "अबको केही समयएमा Quest समाप्त हुनेछ ।"
			}
		},
		{
			"event": "result_update",
			"after": 1,
			"types": [
				"in_app_notification",
				"push_notification"
			],
			"title": {
				"en": "Quest Result"
			},
			"body": {
				"en": "Quest को विवरण  प्रकाशित भएको छ , कृपया Driver App मा गएर आफ्नो Quest को जानकारी लिनुहोस ।"
			}
		}
	]
}
`

func JSONTest3() {
	var obj map[string]interface{}
	if cErr := json.Unmarshal([]byte(jsonString), &obj); cErr != nil {
		fmt.Println("Error Occurred", cErr)
	}

	fmt.Println(obj)
	nested := obj["groups"].([]interface{})
	fmt.Println("nested: ", nested[0])
	nestedNested := nested[0].(map[string]interface{})
	fmt.Println("nested nested: ", nestedNested)
	nestedNested["riders"] = []int{100, 200, 300}
	fmt.Println(nestedNested)
	fmt.Println("old: ", obj)
}
