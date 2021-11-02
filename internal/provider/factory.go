package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
					Optional: true,
				},

				"expect": {
					Type:     t,
					Optional: true,
				},
			},
		},
	}
}
