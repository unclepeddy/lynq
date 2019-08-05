package main

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func dataSourceConcert() *schema.Resource {
	return &schema.Resource{
        Read:   dataSourceConcertRead,
        Schema: map[string]*schema.Schema{
			"max_concerts": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"concerts": {
				Type: schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"concert_id": {
							Type: schema.TypeString,
							Computed: true,
						},
						"title": {
							Type: schema.TypeString,
							Computed: true,
						},
						"location": {
							Type: schema.TypeString,
							Computed: true,
						},
						"date_start": {
							Type: schema.TypeString,
							Computed: true,
						},
						"date_end": {
							Type: schema.TypeString,
							Computed: true,
						},
					},
				},
			},
        },
    }
}

func dataSourceConcertRead(d *schema.ResourceData, meta interface{}) error {
	config:= meta.(*Config)

	con, err := config.music.GetConcerts()
	if err != nil {
		return errors.Wrap(err, "Failed to get concerts for user")
	}

	var k int
	if kMeta, ok := d.GetOk("max_concerts"); !ok {
		k = len(con)
	} else {
		k = kMeta.(int)
		if k > len(con) {
			k = len(con)
		}
	}

	a := make([]interface{}, k)

	for i, c := range con[:k] {
		// Attach Z suffix to denote UTC and parse date time
		t, err := time.Parse(time.RFC3339, c.Datetime + "Z")
		if err != nil {
			return fmt.Errorf("Error while parsing date `%s`", c.Datetime)
		}

		a[i] = map[string]interface{}{
				"concert_id": c.Id,
				"title": c.Title,
				"location": c.City,
				"date_start": t.Format(time.RFC3339),
				"date_end": t.Add(time.Hour * 3).Format(time.RFC3339),
		}
	}

	if err  := d.Set("concerts", a); err != nil {
		return fmt.Errorf("Error setting concerts")
	}

	d.SetId(time.Now().String())

    return nil
}
