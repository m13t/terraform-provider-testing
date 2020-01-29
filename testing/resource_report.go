package testing

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceReport() *schema.Resource {
	s := map[string]*schema.Schema{
		"format": {
			Type:     schema.TypeString,
			Default:  "html",
			Optional: true,
		},

		"path": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	return &schema.Resource{
		Schema: s,
		Create: resourceReportCreate,
		Read:   resourceReportRead,
		Update: resourceReportUpdate,
		Delete: resourceReportDelete,
	}
}

func resourceReportCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("test")

	return resourceReportRead(d, m)
}

func resourceReportRead(d *schema.ResourceData, m interface{}) error {
	time.Sleep(100000 * time.Millisecond)

	return nil
}

func resourceReportUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceReportRead(d, m)
}

func resourceReportDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
