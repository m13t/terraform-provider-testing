package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func assertionFactory(t schema.ValueType, d string) *schema.Schema {
	return &schema.Schema{
		Description: d,
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"statement": {
					Description: "Description of the statement being asserted",
					Type:        schema.TypeString,
					Required:    true,
				},

				"input": {
					Description: "Actual value to be tested",
					Type:        t,
					Optional:    true,
				},

				"expect": {
					Description: "Expected value that `input` should have",
					Type:        t,
					Optional:    true,
				},
			},
		},
	}
}
