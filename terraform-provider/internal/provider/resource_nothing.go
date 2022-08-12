package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"regexp"
)

func resourceNothing() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a Nothing",

		CreateContext: resourceNothingCreate,
		ReadContext:   resourceNothingRead,
		UpdateContext: resourceNothingUpdate,
		DeleteContext: resourceNothingDelete,

		Schema: map[string]*schema.Schema{
			"something": {
				Description: "something",
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[a-z0-9-]+$"),
					"Something is in invalid format"),
			},
			"anything": {
				Description: "anything",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceNothingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceNothingRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceNothingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceNothingDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
