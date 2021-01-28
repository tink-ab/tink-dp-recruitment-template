package accounts

import (
	"github.com/tink-ab/digimon-recrutiment-template/template"
	"testing"
	"time"
)

func TestGetAccounts(t *testing.T) {
	apiClient := AccountsApiClient{baseUrl: template.ApiBaseUlr, apiClientId: template.ApiClientId}
	accounts, err := apiClient.GetAccounts()

	if err != nil {
		t.Fatal(err)
	}

	if len(accounts.Accounts) == 0 {
		t.Error("account list is empty")
	}

}

func TestGetTransactions(t *testing.T) {
	apiClient := AccountsApiClient{baseUrl: template.ApiBaseUlr}

	layout := "2006-01-02"
	now := time.Now()
	fromDate := now.AddDate(0, 0, -90)

	_, err := apiClient.GetTransactions("1", fromDate.Format(layout), now.Format(layout))

	if err != nil {
		t.Fatal(err)
	}

}
