package validators

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Delegations struct {
	DelegationResponse []struct {
		Delegation struct {
			ValidatorAddress string `json:"validator_address"`
			Shares           string `json:"shares"`
		} `json:"delegation"`
	} `json:"delegation_responses"`
}

type Validators struct {
	Validator struct {
		Description struct {
			Moniker string `json:"moniker"`
		} `json:"description"`
	} `json:"validator"`
}

type Validations struct {
	Validator string
	Amount    float64
}

func GetValidator(url string, address string) []Validations {
	validationData := []Validations{}

	// Collect delegation data
	delegationReq, err := http.NewRequest("GET", url+"cosmos/staking/v1beta1/delegations/"+address, nil)
	if err != nil {
		log.Fatal(err)
	}
	delegationReq.Header.Add("Content-Type", "application/json")

	var delegationData Delegations

	delegationResponse, err := http.DefaultClient.Do(delegationReq)
	if err != nil {
		log.Fatal(err)
	}
	defer delegationResponse.Body.Close()

	// Populate data to delegationData struct
	if err := json.NewDecoder(delegationResponse.Body).Decode(&delegationData); err != nil {
		log.Fatal(err)
	}

	// Loop over delegations to map data to validator monikers
	for _, delegation := range delegationData.DelegationResponse {
		// Convert the delegation amount to float for db insert
		delegationAmount, err := strconv.ParseFloat(delegation.Delegation.Shares, 64)
		if err != nil {
			log.Fatal(err)
		}
		delegationAmount = delegationAmount / float64(10e17)

		delegatedValidator := delegation.Delegation.ValidatorAddress
		validatorReq, err := http.NewRequest("GET", url+"cosmos/staking/v1beta1/validators/"+delegatedValidator, nil)
		if err != nil {
			log.Fatal(err)
		}
		validatorReq.Header.Add("Content-Type", "application/json")

		var validatorData Validators

		validatorResponse, err := http.DefaultClient.Do(validatorReq)
		if err != nil {
			log.Fatal(err)
		}
		defer validatorResponse.Body.Close()

		// Populate data to validatorData struct
		if err := json.NewDecoder(validatorResponse.Body).Decode(&validatorData); err != nil {
			log.Fatal(err)
		}

		delegatedData := Validations{
			Validator: validatorData.Validator.Description.Moniker,
			Amount:    float64(delegationAmount),
		}

		validationData = append(validationData, delegatedData)
	}

	return validationData
}
