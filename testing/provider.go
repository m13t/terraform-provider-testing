package testing

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"testing_unit": dataSourceUnit(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"testing_report": resourceReport(),
		},
	}
}
