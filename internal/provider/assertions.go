package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var (
	assertions = map[string]assertion{
		"bool_equal": {
			Schema: assertionFactory(schema.TypeBool),
			Evaluator: func(a, b interface{}) bool {
				return a.(bool) == b.(bool)
			},
		},

		"bool_not_equal": {
			Schema: assertionFactory(schema.TypeBool),
			Evaluator: func(a, b interface{}) bool {
				return a.(bool) != b.(bool)
			},
		},

		"string_equal": {
			Schema: assertionFactory(schema.TypeString),
			Evaluator: func(a, b interface{}) bool {
				return a.(string) == b.(string)
			},
		},

		"string_not_equal": {
			Schema: assertionFactory(schema.TypeString),
			Evaluator: func(a, b interface{}) bool {
				return a.(string) != b.(string)
			},
		},

		"integer_equal": {
			Schema: assertionFactory(schema.TypeInt),
			Evaluator: func(a, b interface{}) bool {
				return a.(int) == b.(int)
			},
		},

		"integer_not_equal": {
			Schema: assertionFactory(schema.TypeInt),
			Evaluator: func(a, b interface{}) bool {
				return a.(int) != b.(int)
			},
		},

		"integer_less_than": {
			Schema: assertionFactory(schema.TypeInt),
			Evaluator: func(a, b interface{}) bool {
				return a.(int) < b.(int)
			},
		},

		"integer_greater_than": {
			Schema: assertionFactory(schema.TypeInt),
			Evaluator: func(a, b interface{}) bool {
				return a.(int) > b.(int)
			},
		},

		"float_equal": {
			Schema: assertionFactory(schema.TypeFloat),
			Evaluator: func(a, b interface{}) bool {
				return a.(float64) == b.(float64)
			},
		},

		"float_not_equal": {
			Schema: assertionFactory(schema.TypeFloat),
			Evaluator: func(a, b interface{}) bool {
				return a.(float64) != b.(float64)
			},
		},

		"float_less_than": {
			Schema: assertionFactory(schema.TypeFloat),
			Evaluator: func(a, b interface{}) bool {
				return a.(float64) < b.(float64)
			},
		},

		"float_greater_than": {
			Schema: assertionFactory(schema.TypeFloat),
			Evaluator: func(a, b interface{}) bool {
				return a.(float64) > b.(float64)
			},
		},
	}
)
