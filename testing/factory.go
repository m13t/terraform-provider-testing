package testing

import "github.com/hashicorp/terraform/helper/schema"

func assertionFactory(t schema.ValueType) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"statement": {
					Type:     schema.TypeString,
					Required: true,
				},

				"input": {
					Type:     t,
					Required: true,
				},

				"expect": {
					Type:     t,
					Required: true,
				},
			},
		},
	}
}
