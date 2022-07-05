package main

import (
	"errors"
	"fmt"
	"github.com/dghubble/sling"
)

type apiClient struct {
	username string
	password string
}

const NothingApi = "nothings/%s"
const NothingsApi = "nothings"

type Nothing struct {
	Id        string `json:"id"`
	Something string `json:"something"`
	Anything  string `json:"anything"`
}

type nothingRequest struct {
	Something string `json:"something"`
	Anything  string `json:"anything"`
}

func (a *apiClient) CreateNothing(something string, anything string) (*Nothing, error) {
	requestBody := nothingRequest{Something: something, Anything: anything}

	nothing := new(Nothing)
	_, err := a.newClient().
		Post(NothingsApi).
		BodyJSON(requestBody).
		Add("Accept", "application/json").
		ReceiveSuccess(nothing)

	if err != nil {
		return nil, err
	}

	return nothing, nil
}

func (a *apiClient) GetNothing(id string) (*Nothing, error) {
	nothing := new(Nothing)
	resp, err := a.newClient().
		Get(fmt.Sprintf(NothingApi, id)).
		Add("Accept", "application/json").
		ReceiveSuccess(nothing)

	if resp.StatusCode == 404 {
		return nil, nil
	}

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not nothing details. %#v", err))
	}

	return nothing, nil
}

func (a *apiClient) UpdateNothing(id string, something string, anything string) (*Nothing, error) {
	requestBody := nothingRequest{Something: something, Anything: anything}

	nothing := new(Nothing)
	_, err := a.newClient().
		Put(fmt.Sprintf(NothingApi, id)).
		BodyJSON(requestBody).
		Add("Accept", "application/json").
		ReceiveSuccess(nothing)

	if err != nil {
		return nil, err
	}

	return nothing, nil
}

func (a *apiClient) DeleteNothing(id string) error {
	_, err := a.newClient().
		Delete(fmt.Sprintf(NothingApi, id)).
		ReceiveSuccess(nil)

	if err != nil {
		return err
	}

	return nil
}

func (a *apiClient) newClient() *sling.Sling {
	return sling.New().Base("http://localhost:8080/").SetBasicAuth(a.username, a.password)
}
