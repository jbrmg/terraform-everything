package everything

import (
	"errors"
	"fmt"
	"github.com/dghubble/sling"
)

type ApiClient struct {
	Username string
	Password string
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

func (a *ApiClient) CreateNothing(something string, anything string) (*Nothing, error) {
	requestBody := nothingRequest{Something: something, Anything: anything}

	nothing := new(Nothing)
	resp, err := a.newClient().
		Post(NothingsApi).
		BodyJSON(requestBody).
		Add("Accept", "application/json").
		ReceiveSuccess(nothing)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return nothing, nil
}

func (a *ApiClient) GetNothing(id string) (*Nothing, error) {
	nothing := new(Nothing)
	resp, err := a.newClient().
		Get(fmt.Sprintf(NothingApi, id)).
		Add("Accept", "application/json").
		ReceiveSuccess(nothing)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not nothing details. %#v", err))
	}

	if resp.StatusCode == 404 {
		return nil, nil
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return nothing, nil
}

func (a *ApiClient) UpdateNothing(id string, something string, anything string) (*Nothing, error) {
	requestBody := nothingRequest{Something: something, Anything: anything}

	nothing := new(Nothing)
	resp, err := a.newClient().
		Put(fmt.Sprintf(NothingApi, id)).
		BodyJSON(requestBody).
		Add("Accept", "application/json").
		ReceiveSuccess(nothing)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return nothing, nil
}

func (a *ApiClient) DeleteNothing(id string) error {
	resp, err := a.newClient().
		Delete(fmt.Sprintf(NothingApi, id)).
		ReceiveSuccess(nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return nil
}

func (a *ApiClient) newClient() *sling.Sling {
	return sling.New().Base("http://localhost:8080/").SetBasicAuth(a.Username, a.Password)
}
