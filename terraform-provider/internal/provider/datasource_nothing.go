package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNothing() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieves data about a Nothing",
		ReadContext: dataSourceNothingRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Identifier of the Nothing",
				Type:        schema.TypeString,
				Required:    true,
			},
			"something": {
				Description: "The attribute something",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"anything": {
				Description: "The attribute anything",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceNothingRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	nothing, err := getClient(meta).GetNothing(id)

	if err != nil {
		return diag.Errorf("could not fetch nothing details: %s", err)
	}

	if nothing == nil {
		return diag.Errorf("Nothing %s does not exist or access is forbidden.", id)
	}

	d.SetId(d.Get("id").(string))
	d.Set("something", nothing.Something)
	d.Set("anything", nothing.Anything)
	return nil
}
