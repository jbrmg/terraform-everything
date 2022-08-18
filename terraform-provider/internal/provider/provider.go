package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-everything/internal/ikea"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown
}

func New() func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			ResourcesMap: map[string]*schema.Resource{
				"ikea_kitchen":    resourceKitchen(),
				"ikea_cabinet":    resourceCabinet(),
				"ikea_countertop": resourceCounterTop(),
			},
			DataSourcesMap: map[string]*schema.Resource{},
			Schema:         map[string]*schema.Schema{},
		}

		p.ConfigureContextFunc = providerConfigure
		return p
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return &ikea.ApiClient{}, nil
}

func getClient(meta interface{}) *ikea.ApiClient {
	return meta.(*ikea.ApiClient)
}
