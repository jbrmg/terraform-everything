package ikea

import (
	"errors"
	"fmt"
	"github.com/dghubble/sling"
)

type ApiClient struct {
	Username string
	Password string
}

const KitchensApi = "kitchens"
const KitchenApi = KitchensApi + "/%s"
const CabinetsApi = "cabinets"
const CabinetApi = CabinetsApi + "/%s"
const CounterTopsApi = "countertops"
const CounterTopApi = CounterTopsApi + "/%s"

type Kitchen struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type kitchenRequest struct {
	Name string `json:"name"`
}

type Cabinet struct {
	Id        string `json:"id"`
	Front     string `json:"front"`
	Color     string `json:"color"`
	KitchenId string `json:"kitchenId"`
}

type cabinetRequest struct {
	Front     string `json:"front"`
	Color     string `json:"color"`
	KitchenId string `json:"kitchenId"`
}

type CounterTop struct {
	Id         string   `json:"id"`
	Type       string   `json:"type"`
	CabinetIds []string `json:"cabinetIds"`
}

type counterTopRequest struct {
	Type       string   `json:"type"`
	CabinetIds []string `json:"cabinetIds"`
}

func (a *ApiClient) CreateKitchen(name string) (*Kitchen, error) {
	requestBody := kitchenRequest{Name: name}

	kitchen := new(Kitchen)
	resp, err := a.newClient().
		Post(KitchensApi).
		BodyJSON(requestBody).
		Add("Accept", "application/json").
		ReceiveSuccess(kitchen)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return kitchen, nil
}

func (a *ApiClient) GetKitchen(id string) (*Kitchen, error) {
	kitchen := new(Kitchen)
	resp, err := a.newClient().
		Get(fmt.Sprintf(KitchenApi, id)).
		Add("Accept", "application/json").
		ReceiveSuccess(kitchen)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not kitchen details. %#v", err))
	}

	if resp.StatusCode == 404 {
		return nil, nil
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return kitchen, nil
}

func (a *ApiClient) UpdateKitchen(id string, name string) (*Kitchen, error) {
	requestBody := kitchenRequest{Name: name}

	kitchen := new(Kitchen)
	resp, err := a.newClient().
		Put(fmt.Sprintf(KitchenApi, id)).
		BodyJSON(requestBody).
		Add("Accept", "application/json").
		ReceiveSuccess(kitchen)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return kitchen, nil
}

func (a *ApiClient) DeleteKitchen(id string) error {
	resp, err := a.newClient().
		Delete(fmt.Sprintf(KitchenApi, id)).
		ReceiveSuccess(nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return nil
}

func (a *ApiClient) CreateCabinet(color string, front string, kitchenId string) (*Cabinet, error) {
	requestBody := cabinetRequest{Color: color, Front: front, KitchenId: kitchenId}

	cabinet := new(Cabinet)
	resp, err := a.newClient().
		Post(CabinetsApi).
		BodyJSON(requestBody).
		Add("Accept", "application/json").
		ReceiveSuccess(cabinet)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return cabinet, nil
}

func (a *ApiClient) GetCabinet(id string) (*Cabinet, error) {
	cabinet := new(Cabinet)
	resp, err := a.newClient().
		Get(fmt.Sprintf(CabinetApi, id)).
		Add("Accept", "application/json").
		ReceiveSuccess(cabinet)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not cabinet details. %#v", err))
	}

	if resp.StatusCode == 404 {
		return nil, nil
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return cabinet, nil
}

func (a *ApiClient) UpdateCabinet(id string, color string, front string, kitchenId string) (*Cabinet, error) {
	requestBody := cabinetRequest{Color: color, Front: front, KitchenId: kitchenId}

	cabinet := new(Cabinet)
	resp, err := a.newClient().
		Put(fmt.Sprintf(CabinetApi, id)).
		BodyJSON(requestBody).
		Add("Accept", "application/json").
		ReceiveSuccess(cabinet)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return cabinet, nil
}

func (a *ApiClient) DeleteCabinet(id string) error {
	resp, err := a.newClient().
		Delete(fmt.Sprintf(CabinetApi, id)).
		ReceiveSuccess(nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return nil
}

func (a *ApiClient) CreateCounterTop(cType string, cabinetIds []string) (*CounterTop, error) {
	requestBody := counterTopRequest{Type: cType, CabinetIds: cabinetIds}

	counterTop := new(CounterTop)
	resp, err := a.newClient().
		Post(CounterTopsApi).
		BodyJSON(requestBody).
		Add("Accept", "application/json").
		ReceiveSuccess(counterTop)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return counterTop, nil
}

func (a *ApiClient) GetCounterTop(id string) (*CounterTop, error) {
	counterTop := new(CounterTop)
	resp, err := a.newClient().
		Get(fmt.Sprintf(CounterTopApi, id)).
		Add("Accept", "application/json").
		ReceiveSuccess(counterTop)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not counterTop details. %#v", err))
	}

	if resp.StatusCode == 404 {
		return nil, nil
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return counterTop, nil
}

func (a *ApiClient) UpdateCounterTop(id string, cType string, cabinetIds []string) (*CounterTop, error) {
	requestBody := counterTopRequest{Type: cType, CabinetIds: cabinetIds}

	counterTop := new(CounterTop)
	resp, err := a.newClient().
		Put(fmt.Sprintf(CounterTopApi, id)).
		BodyJSON(requestBody).
		Add("Accept", "application/json").
		ReceiveSuccess(counterTop)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got invalid response code: %d", resp.StatusCode)
	}

	return counterTop, nil
}

func (a *ApiClient) DeleteCounterTop(id string) error {
	resp, err := a.newClient().
		Delete(fmt.Sprintf(CounterTopApi, id)).
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
