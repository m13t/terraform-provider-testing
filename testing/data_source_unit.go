package testing

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceUnit() *schema.Resource {
	s := map[string]*schema.Schema{
		"subject": {
			Type:     schema.TypeString,
			Required: true,
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
		Read:   dataSourceUnitRead,
		Schema: s,
	}
}

func dataSourceUnitRead(d *schema.ResourceData, m interface{}) error {
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
					"[%s] %s", d.Get("subject"), s))
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

	/*
		 * If we want to fail on errors. I'll make this configurable later
		 *
		if len(failed) > 0 {
			return errors.New(strings.Join(failed, "\n"))
		}
	*/

	// Return no errors
	return nil
}
