package accounts

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type AccountsResponse struct {
	Accounts []struct {
		AvailableBalance string   `json:"availableBalance"`
		Currency         []string `json:"currency"`
		ID               string   `json:"id"`
		Links            struct {
			Balances     string `json:"balances"`
			Transactions string `json:"transactions"`
		} `json:"links"`
		Number string `json:"number"`
		Owner  string `json:"owner"`
	} `json:"accounts"`
}

type TransactionsResponse struct {
	//TODO map resposne
}

type AccountApi interface {
	GetAccounts() (AccountsResponse, error)
	GetTransactions(accountId string, dateFrom string, dateTo string) (*TransactionsResponse, error)
}

type AccountsApiClient struct {
	baseUrl     string
	apiClientId string
}

func (c *AccountsApiClient) GetAccounts() (*AccountsResponse, error) {
	httpClient := http.Client{}

	req, _ := http.NewRequest("GET", c.baseUrl+"/api/accounts"+"?api_client_id="+c.apiClientId, nil)

	httpResp, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "problem with api connection")
	}

	checkHttpStatus(*httpResp)

	accJson, _ := ioutil.ReadAll(httpResp.Body)
	response := AccountsResponse{}
	if err := json.Unmarshal(accJson, &response); err != nil {
		return nil, errors.Wrap(err, "problem with response deserialization")
	}

	return &response, nil
}

func (c *AccountsApiClient) GetTransactions(accountId string, dateFrom string, dateTo string) (*TransactionsResponse, error) {
	if err := checkRangeBelow90days(dateFrom, dateTo); err != nil {
		return nil, err
	}

	_, err := http.Get(fmt.Sprintf("%s/api/accounts/%s/transactions?from=%s&to=%s", c.baseUrl, accountId, dateFrom, dateTo))
	if err != nil {
		return nil, errors.Wrap(err, "problem with api connection")
	}

	//TODO finish implementation
	return nil, nil
}

func checkRangeBelow90days(from string, to string) error {
	layout := "2006-01-02"
	f, _ := time.Parse(layout, from)
	t, _ := time.Parse(layout, to)

	if t.Sub(f).Hours() > 24*90 { //90 days
		return errors.New("date range must not exceed 90 days")
	}
	return nil
}

func checkHttpStatus(resp http.Response) error {
	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with http status %d", resp.StatusCode)
	}
	return nil
}
