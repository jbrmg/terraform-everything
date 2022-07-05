package main

func getClient(meta interface{}) *apiClient {
	return meta.(*apiClient)
}
