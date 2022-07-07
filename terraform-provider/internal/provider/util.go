package provider

import (
	"terraform-provider-everything/internal/everything"
)

func getClient(meta interface{}) *everything.ApiClient {
	return meta.(*everything.ApiClient)
}
