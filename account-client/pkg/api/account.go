package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// AccountAttributes model
type AccountAttributes struct {
	Country               string   `json:"country"`
	BaseCurrency          string   `json:"base_currency"`
	BankID                string   `json:"bank_id"`
	BankIDCode            string   `json:"bank_id_code"`
	AccountNumber         string   `json:"account_number"`
	BIC                   string   `json:"bic"`
	IBAN                  string   `json:"iban"`
	CustomerID            string   `json:"customer_id"`
	Name                  []string `json:"name"`
	AccountClassification string   `json:"account_classification"`
}

// AccountData model
type AccountData struct {
	ID             string            `json:"id"`
	Type           string            `json:"type"`
	OrganizationID string            `json:"organisation_id"`
	Version        int               `json:"version"`
	Attributes     AccountAttributes `json:"attributes"`
}

// Account model
type Account struct {
	Data AccountData `json:"data"`
}

// AccountAPI client interact with Account API
type AccountAPI struct {
	Client  *http.Client
	BaseURL string
}

// NewAccountAPI create AccountAPI instance
func NewAccountAPI() *AccountAPI {
	baseURL := os.Getenv("ACCOUNT_API_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return &AccountAPI{Client: http.DefaultClient, BaseURL: baseURL}
}

// Create creat an Account
func (api *AccountAPI) Create(account *Account) error {
	request, err := json.Marshal(*account)
	if err != nil {
		return err
	}
	resp, err := api.Client.Post(api.BaseURL+"/v1/organisation/accounts", "application/json", bytes.NewBuffer(request))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return errors.New(string(body))
}

// Fetch an Account by id
func (api *AccountAPI) Fetch(ID string) (*Account, error) {
	resp, err := api.Client.Get(api.BaseURL + "/v1/organisation/accounts/" + ID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
	}
	if resp.StatusCode == 400 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}
	if resp.StatusCode == 404 {
		return nil, errors.New("Account not found")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	account := Account{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// Delete an Account by id and version
func (api *AccountAPI) Delete(ID string, version int) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/organisation/accounts/%s?version=%d", api.BaseURL, ID, version), nil)
	resp, err := api.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		return nil
	}
	if resp.StatusCode == 404 {
		return errors.New("Account does not exist")
	}
	if resp.StatusCode == 409 {
		return errors.New("Account version incorrect")
	}
	return nil
}

func getEnv(env, fallback string) string {
	value := os.Getenv("ACCOUNT_API_URL")
	if value != "" {
		return value
	}
	return fallback
}
