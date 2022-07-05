package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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
	something := d.Get("something").(string)
	anything := d.Get("anything").(string)

	nothing, err := getClient(meta).CreateNothing(something, anything)

	if err != nil {
		return diag.Errorf("Could not create nothing: %#v", err)
	}

	d.SetId(nothing.Id)
	d.Set("something", nothing.Something)
	d.Set("anything", nothing.Anything)

	tflog.Trace(ctx, "created a nothing")

	return nil
}

func resourceNothingRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	nothing, err := getClient(meta).GetNothing(id)

	if err != nil {
		return diag.Errorf("could not fetch nothing details: %#v", err)
	}

	if nothing == nil {
		d.SetId("") // resource has been deleted
		return nil
	}

	d.Set("something", nothing.Something)
	d.Set("anything", nothing.Anything)
	return nil
}

func resourceNothingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	something := d.Get("something").(string)
	anything := d.Get("anything").(string)

	nothing, err := getClient(meta).UpdateNothing(id, something, anything)

	if err != nil {
		return diag.Errorf("could not update nothing details: %#v", err)
	}

	d.Set("something", nothing.Something)
	d.Set("anything", nothing.Anything)

	tflog.Trace(ctx, fmt.Sprintf("Updated nothing: %#v", nothing))

	return nil
}

func resourceNothingDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()

	err := getClient(meta).DeleteNothing(id)

	if err != nil {
		return diag.FromErr(err)
	}

	// resource has been deleted successfully
	d.SetId("")
	return nil
}
