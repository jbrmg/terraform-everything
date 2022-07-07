package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-everything/internal/everything"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown
}

func New() func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				"everything_nothing": dataSourceNothing(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"everything_nothing": resourceNothing(),
			},
			Schema: map[string]*schema.Schema{
				"username": {
					Description: "Username for basic auth authentication",
					Required:    true,
					Type:        schema.TypeString,
					DefaultFunc: schema.EnvDefaultFunc("EVERYTHING_USERNAME", nil),
				},
				"password": {
					Description: "Password for basic auth authentication",
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("EVERYTHING_PASSWORD", nil),
				},
			},
		}

		p.ConfigureContextFunc = providerConfigure
		return p
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	return &everything.ApiClient{
		Username: username,
		Password: password,
	}, nil
}
