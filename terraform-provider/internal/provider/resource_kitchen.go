package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceKitchen() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a kitchen",

		CreateContext: resourceKitchenCreate,
		ReadContext:   resourceKitchenRead,
		UpdateContext: resourceKitchenUpdate,
		DeleteContext: resourceKitchenDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The identifier of this cabinet",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description:  "Name of the kitchen",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func resourceKitchenCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	kitchen, err := getClient(meta).CreateKitchen(name)

	if err != nil {
		return diag.Errorf("Could not create kitchen: %s", err)
	}

	d.SetId(kitchen.Id)
	d.Set("name", kitchen.Name)

	tflog.Trace(ctx, "created a kitchen")

	return nil
}

func resourceKitchenRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	kitchen, err := getClient(meta).GetKitchen(id)

	if err != nil {
		return diag.Errorf("could not fetch kitchen details: %s", err)
	}

	if kitchen == nil {
		d.SetId("") // resource has been deleted
		return nil
	}

	d.Set("name", kitchen.Name)
	return nil
}

func resourceKitchenUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	name := d.Get("name").(string)

	kitchen, err := getClient(meta).UpdateKitchen(id, name)

	if err != nil {
		return diag.Errorf("could not update kitchen details: %s", err)
	}

	d.Set("name", kitchen.Name)

	tflog.Trace(ctx, fmt.Sprintf("Updated kitchen: %s", kitchen))

	return nil
}

func resourceKitchenDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()

	err := getClient(meta).DeleteKitchen(id)

	if err != nil {
		return diag.FromErr(err)
	}

	// resource has been deleted successfully
	d.SetId("")
	return nil
}
