package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

		Schema: map[string]*schema.Schema{},
	}
}

func resourceKitchenCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceKitchenRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceKitchenUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceKitchenDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
