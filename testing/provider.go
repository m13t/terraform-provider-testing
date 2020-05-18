package testing

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,
		Schema: map[string]*schema.Schema{
			"fail": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"testing_unit": dataSourceUnit(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &config{
		failOnFailure: d.Get("fail").(bool),
	}

	return config, nil
}
