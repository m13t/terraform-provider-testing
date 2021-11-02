package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type assertionEvaluator func(a, b interface{}) bool

type assertion struct {
	Schema    *schema.Schema
	Evaluator assertionEvaluator
}

// Shorthand for running Evaluator
func (as assertion) evaluate(a, b interface{}) bool {
	return as.Evaluator(a, b)
}
