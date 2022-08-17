package provider

import (
	"terraform-provider-everything/internal/ikea"
)

func getClient(meta interface{}) *ikea.ApiClient {
	return meta.(*ikea.ApiClient)
}
