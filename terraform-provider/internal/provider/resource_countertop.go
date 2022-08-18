package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCounterTop() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a kitchen counter-top",

		CreateContext: resourceCounterTopCreate,
		ReadContext:   resourceCounterTopRead,
		UpdateContext: resourceCounterTopUpdate,
		DeleteContext: resourceCounterTopDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The identifier of this counter-top",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description:  "The type of this counter-top",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"EKBACKEN", "KARLBY", "SKARARP"}, false),
			},
			"cabinet_ids": {
				Description: "The type of the cabinetIds",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
		},
	}
}

func resourceCounterTopCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cType := d.Get("type").(string)
	cabinetIds := parseCabinetIds(d)

	countertop, err := getClient(meta).CreateCounterTop(cType, cabinetIds)

	if err != nil {
		return diag.Errorf("Could not create countertop: %s", err)
	}

	d.SetId(countertop.Id)
	d.Set("type", countertop.Type)
	d.Set("cabinet_ids", countertop.CabinetIds)

	tflog.Trace(ctx, "created a countertop")

	return nil
}

func resourceCounterTopRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	counterTop, err := getClient(meta).GetCounterTop(id)

	if err != nil {
		return diag.Errorf("could not fetch counter-top details: %s", err)
	}

	if counterTop == nil {
		d.SetId("") // resource has been deleted
		return nil
	}

	d.Set("type", counterTop.Type)
	d.Set("cabinet_ids", counterTop.CabinetIds)
	return nil
}

func resourceCounterTopUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	cType := d.Get("type").(string)
	cabinetIds := parseCabinetIds(d)

	counterTop, err := getClient(meta).UpdateCounterTop(id, cType, cabinetIds)

	if err != nil {
		return diag.Errorf("could not update counter-top details: %s", err)
	}

	d.Set("type", counterTop.Type)
	d.Set("cabinet_ids", counterTop.CabinetIds)

	tflog.Trace(ctx, fmt.Sprintf("Updated counter-top: %s", counterTop))

	return nil
}

func resourceCounterTopDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()

	err := getClient(meta).DeleteCounterTop(id)

	if err != nil {
		return diag.FromErr(err)
	}

	// resource has been deleted successfully
	d.SetId("")
	return nil
}

func parseCabinetIds(d *schema.ResourceData) []string {
	itemsRaw := d.Get("cabinet_ids").([]interface{})
	items := make([]string, len(itemsRaw))
	for i, raw := range itemsRaw {
		items[i] = raw.(string)
	}
	return items
}
