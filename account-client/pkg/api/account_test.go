package api

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreateFetchDeleteAccount(t *testing.T) {
	accountAPI := NewAccountAPI()
	// account model
	acccount := &Account{
		Data: AccountData{
			ID:             uuid.New().String(),
			Type:           "accounts",
			OrganizationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: AccountAttributes{
				Country:               "GB",
				BaseCurrency:          "GBP",
				BankID:                "000001",
				BankIDCode:            "GBDSC",
				BIC:                   "NWBKGB22",
				Name:                  []string{"Guimarães", "Renato"},
				AccountClassification: "Personal",
			},
		},
	}
	// create an account
	err := accountAPI.Create(acccount)
	if err != nil {
		t.Error(err)
	}
	// fetch an account
	accountFetched, err := accountAPI.Fetch(acccount.Data.ID)
	if err != nil {
		t.Error(err)
	}
	if accountFetched == nil {
		t.Errorf("Account %s not found\n", accountFetched.Data.ID)
	}
	// delete an account
	err = accountAPI.Delete(acccount.Data.ID, accountFetched.Data.Version)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateInvalidAccount(t *testing.T) {
	accountAPI := NewAccountAPI()
	// account model
	acccount := &Account{
		Data: AccountData{},
	}
	// create an account
	err := accountAPI.Create(acccount)
	if err == nil {
		t.Error("An error is expected")
	}
}

func TestFetchNotExistentAccount(t *testing.T) {
	accountAPI := NewAccountAPI()
	// create an account
	accountFetched, err := accountAPI.Fetch(uuid.New().String())
	if accountFetched != nil {
		t.Error("Account shouldn't be found")
	}
	if err == nil {
		t.Error("An error is expected")
	}
}
