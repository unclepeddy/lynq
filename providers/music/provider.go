package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

// Provider returns the actual provider instance.
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"music_concert": dataSourceConcert(),
		},
		ConfigureFunc: providerConfigure,
		Schema: map[string]*schema.Schema{
			"spotify_user": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var c Config
	spotifyUserId := d.Get("spotify_user").(string)
	if err := c.loadAndValidate(spotifyUserId); err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}
	return &c, nil
	}
