package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCabinet() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieves data about a cabinet",
		ReadContext: dataSourceCabinetRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Identifier of the cabinet",
				Type:        schema.TypeString,
				Required:    true,
			},
			"color": {
				Description: "The color",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"front": {
				Description: "The type of the front.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceCabinetRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	cabinet, err := getClient(meta).GetCabinet(id)

	if err != nil {
		return diag.Errorf("could not fetch cabinet details: %s", err)
	}

	if cabinet == nil {
		return diag.Errorf("cabinet %s does not exist or access is forbidden.", id)
	}

	d.SetId(d.Get("id").(string))
	d.Set("color", cabinet.Color)
	d.Set("front", cabinet.Front)
	return nil
}
