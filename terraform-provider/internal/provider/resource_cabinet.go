package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"regexp"
)

func resourceCabinet() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a kitchen cabinet",

		CreateContext: resourceCabinetCreate,
		ReadContext:   resourceCabinetRead,
		UpdateContext: resourceCabinetUpdate,
		DeleteContext: resourceCabinetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The identifier of this cabinet",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"color": {
				Description: "The color of this cabinet. Should be a valid hex color code. Example: #FFFFFF",
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^#(?:[0-9a-fA-F]{3}){1,2}$"),
					"color must be a valid hex color code"),
			},
			"front": {
				Description:  "The type of the front",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"RINGHULT", "VOXTORP", "VEDDINGE"}, false),
			},
			"kitchen_id": {
				Description:  "Identifier of the kitchen this cabinet belongs to",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func resourceCabinetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	color := d.Get("color").(string)
	front := d.Get("front").(string)
	kitchenId := d.Get("kitchen_id").(string)

	cabinet, err := getClient(meta).CreateCabinet(color, front, kitchenId)

	if err != nil {
		return diag.Errorf("Could not create cabinet: %s", err)
	}

	d.SetId(cabinet.Id)
	d.Set("color", cabinet.Color)
	d.Set("front", cabinet.Front)

	tflog.Trace(ctx, "created a cabinet")

	return nil
}

func resourceCabinetRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	cabinet, err := getClient(meta).GetCabinet(id)

	if err != nil {
		return diag.Errorf("could not fetch cabinet details: %s", err)
	}

	if cabinet == nil {
		d.SetId("") // resource has been deleted
		return nil
	}

	d.Set("color", cabinet.Color)
	d.Set("front", cabinet.Front)
	return nil
}

func resourceCabinetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	color := d.Get("color").(string)
	front := d.Get("front").(string)
	kitchenId := d.Get("kitchen_id").(string)

	cabinet, err := getClient(meta).UpdateCabinet(id, color, front, kitchenId)

	if err != nil {
		return diag.Errorf("could not update cabinet details: %s", err)
	}

	d.Set("color", cabinet.Color)
	d.Set("front", cabinet.Front)

	tflog.Trace(ctx, fmt.Sprintf("Updated cabinet: %s", cabinet))

	return nil
}

func resourceCabinetDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()

	err := getClient(meta).DeleteCabinet(id)

	if err != nil {
		return diag.FromErr(err)
	}

	// resource has been deleted successfully
	d.SetId("")
	return nil
}
