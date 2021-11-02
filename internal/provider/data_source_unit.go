package provider

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUnit() *schema.Resource {
	s := map[string]*schema.Schema{
		"subject": {
			Description: "a description of the test being performed",
			Type:        schema.TypeString,
			Required:    true,
		},

		"passed": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		"failed": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	for k, v := range assertions {
		s[k] = v.Schema
	}

	return &schema.Resource{
		ReadContext: dataSourceUnitRead,
		Schema:      s,
	}
}

func dataSourceUnitRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[LL] Subject: %s\n", d.Get("subject"))

	var passed []string
	var failed []string

	for assertionName, assertionValue := range assertions {
		for _, assertionType := range d.Get(assertionName).([]interface{}) {
			assertionsMap := assertionType.(map[string]interface{})

			s := assertionsMap["statement"]
			i := assertionsMap["input"]
			e := assertionsMap["expect"]
			r := assertionValue.evaluate(i, e)

			if r {
				passed = append(passed, fmt.Sprintf(
					"[%s] %s", d.Get("subject"), s))
			} else {
				failed = append(failed, fmt.Sprintf(
					"[%s] %s - expected '%v' but got '%v'", d.Get("subject"), s, e, i))
			}

			// Log result
			log.Printf("[LL] %s (%v => %v) = %t\n", s, i, e, r)
		}
	}

	// Set ID to current time. This allows the data source
	// to be marked as successful but will rerun every time
	d.SetId(time.Now().UTC().String())

	// Set passed and failed
	d.Set("passed", passed)
	d.Set("failed", failed)

	// Get config
	c := m.(*config)

	if len(failed) > 0 && c.failOnFailure {
		return diag.Errorf(strings.Join(failed, "\n"))
	}

	// Return no errors
	return nil
}
