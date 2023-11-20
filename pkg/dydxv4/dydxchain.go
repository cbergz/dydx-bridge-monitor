package dydxchain

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Struct from the API to query user balances
type Holder struct {
	Balance struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"balance"`
}

// Struct to be used for populating user delegations
type Delegation struct {
	Validator string
	Amount    float64
}

// Struct from the API to query user delegations
type Delegations struct {
	DelegationResponse []struct {
		Delegation struct {
			ValidatorAddress string `json:"validator_address"`
			Shares           string `json:"shares"`
		} `json:"delegation"`
	} `json:"delegation_responses"`
}

// Struct from the API to query validator moniker
type Validators struct {
	Validator struct {
		Description struct {
			Moniker string `json:"moniker"`
		} `json:"description"`
	} `json:"validator"`
}

func GetBalance(api_url string, address string) float64 {
	url := fmt.Sprintf(api_url+"/cosmos/bank/v1beta1/balances/%s/by_denom?denom=adydx", address)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var dydxHolder Holder

	if decodeErr := json.NewDecoder(res.Body).Decode(&dydxHolder); err != nil {
		log.Fatal(decodeErr)
	}

	if dydxHolder.Balance.Amount == "" {
		dydxHolder.Balance.Amount = "0.00"
	}

	dydxBalance, err := strconv.ParseFloat(dydxHolder.Balance.Amount, 64)
	if err != nil {
		log.Fatal(err, dydxHolder, address)
	}

	return dydxBalance / float64(10e17)

}

func GetDelegations(api_url string, address string) []Delegation {
	url := fmt.Sprintf(api_url+"/cosmos/staking/v1beta1/delegations/%s", address)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var userDelegation []Delegation
	var resDelegations Delegations

	if resErr := json.NewDecoder(res.Body).Decode(&resDelegations); resErr != nil {
		log.Fatal(resErr)
	}

	// Skip processing if DelegationResponse is empty
	if len(resDelegations.DelegationResponse) == 0 {
		newDelegation := Delegation{"", 0}
		userDelegation = append(userDelegation, newDelegation)
		return userDelegation
	}

	for _, validator := range resDelegations.DelegationResponse {

		moniker := GetValidatorMoniker(api_url, validator.Delegation.ValidatorAddress)
		amount, err := strconv.ParseFloat(validator.Delegation.Shares, 64)
		if err != nil {
			log.Fatal(err, validator, address)
		}
		amount = amount / float64(10e17)
		newDelegation := Delegation{moniker, amount}
		userDelegation = append(userDelegation, newDelegation)
	}

	return userDelegation
}

func GetValidatorMoniker(api_url string, address string) string {
	url := fmt.Sprintf(api_url+"/cosmos/staking/v1beta1/validators/%s", address)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var resValidator Validators

	if resErr := json.NewDecoder(res.Body).Decode(&resValidator); resErr != nil {
		log.Fatal(resErr)
	}

	return resValidator.Validator.Description.Moniker

}
