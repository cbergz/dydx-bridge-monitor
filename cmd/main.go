package main

import (
	"context"
	dydxtoken "dydx-bridge-monitor/pkg/dydxv3/contracts/dydxtoken"
	wrapdydxtoken "dydx-bridge-monitor/pkg/dydxv3/contracts/wrapdydxtoken"
	dydxchain "dydx-bridge-monitor/pkg/dydxv4"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/jmoiron/sqlx"
)

var apiCallCounter int

func main() {
	const maxAPICallsBeforePause = 50
	const pauseDurationSeconds = 60

	// Setup the database connection and insert command
	db, err := sqlx.Connect("pgx", "postgresql://doadmin:AVNS_A3mgPzwuR6swL6WYcPs@db-postgresql-nyc1-02376-do-user-3592284-0.b.db.ondigitalocean.com:25061/dydx_user?sslmode=require")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	bridgeInsert := `INSERT INTO bridge_events (id, sender, receiver, amount) VALUES ($1, $2, $3, $4) ON CONFLICT (id) DO NOTHING`
	holdersInsert := `INSERT INTO dydx_holders (dydx_address, ethereum_address) VALUES ($1, $2) ON CONFLICT (dydx_address) DO NOTHING`

	// Load API keys from env file
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error loading env file")
	}
	alchemyKey := os.Getenv("ALCHEMY_KEY")
	dydxApiUrl := os.Getenv("RPC_ENDPOINT")

	// Setup the connection to Ethereum using Infura RPC
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + alchemyKey)
	if err != nil {
		log.Fatal(err)
	}

	// Define the DYDX token contract
	dydxTokenAddress := common.HexToAddress("0x92d6c1e31e14520e676a687f0a93788b716beff5")
	dydxTokenContract, err := dydxtoken.NewDydxtokenCaller(dydxTokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Define the bridge contract
	bridgeAddress := common.HexToAddress("0x46b2deae6eff3011008ea27ea36b7c27255ddfa9")
	bridgeContract, err := wrapdydxtoken.NewBridgeFilterer(bridgeAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	wrapTokenContract, err := wrapdydxtoken.NewBridgeCaller(bridgeAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Pull all the bridge transfer events from Ethereum
	var idValues []*big.Int
	bridgeTransfers, err := bridgeContract.FilterBridge(&bind.FilterOpts{Start: 0, End: nil}, idValues)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through all the events to find any new bridge transfers
	for bridgeTransfers.Next() {
		event := bridgeTransfers.Event

		// Convert the values as needed for reading and inserting
		eventID := event.Id.String()

		dydxAmount, _ := new(big.Rat).SetInt(event.Amount).Float64()
		dydxAmount = dydxAmount / float64(10e17)
		cosmosAddress, _ := bech32.ConvertAndEncode("dydx", event.AccAddress)

		// Add new transfers to the birdge table
		_, errBridge := db.Exec(bridgeInsert, eventID, event.From.String(), cosmosAddress, dydxAmount)
		if errBridge != nil {
			log.Fatal(errBridge)
		}

		// Add new addresses to the holders table
		_, errHolder := db.Exec(holdersInsert, cosmosAddress, event.From.String())
		if errHolder != nil {
			log.Fatal(errHolder)
		}
	}

	dydxAddresses, err := db.Query("SELECT dydx_address FROM dydx_holders")
	if err != nil {
		log.Fatal(err)
	}
	defer dydxAddresses.Close()

	for dydxAddresses.Next() {
		apiCallCounter++
		var dydxAddress string
		var totalDydxBalance float64
		var delegatedDydxBalance float64

		scanErr := dydxAddresses.Scan(&dydxAddress)
		if scanErr != nil {
			log.Fatal(scanErr)
		}

		dydxAvailableBalance := dydxchain.GetBalance(dydxApiUrl, dydxAddress)
		totalDydxBalance += dydxAvailableBalance
		dydxDelegations := dydxchain.GetDelegations(dydxApiUrl, dydxAddress)
		for _, delegation := range dydxDelegations {
			totalDydxBalance += delegation.Amount
			delegatedDydxBalance += delegation.Amount
		}

		balanceInsert := `UPDATE dydx_holders SET adydx_balance = $1, adydx_available = $2, adydx_delegated = $3 WHERE dydx_address = $4`

		_, balanceErr := db.Exec(balanceInsert, totalDydxBalance, dydxAvailableBalance, delegatedDydxBalance, dydxAddress)
		if balanceErr != nil {
			log.Fatal(balanceErr)
		}

		// Check if it's time to pause
		if apiCallCounter%maxAPICallsBeforePause == 0 {
			fmt.Printf("Pausing for %d seconds after %d API calls\n", pauseDurationSeconds, apiCallCounter)
			time.Sleep(time.Second * pauseDurationSeconds)
			apiCallCounter = 0
		}
	}

	ethereumAddresses, err := db.Query("SELECT ethereum_address FROM dydx_holders")
	if err != nil {
		log.Fatal(err)
	}
	defer ethereumAddresses.Close()

	for ethereumAddresses.Next() {
		var ethereumAddress string

		scanErr := ethereumAddresses.Scan(&ethereumAddress)
		if scanErr != nil {
			log.Fatal(scanErr)
		}

		ethBalanceResponse, err := dydxTokenContract.BalanceOf(&bind.CallOpts{Pending: false, Context: context.Background()}, common.HexToAddress(ethereumAddress))
		if err != nil {
			log.Fatal(err)
		}
		balanceFloat, _ := new(big.Rat).SetInt(ethBalanceResponse).Float64()
		ethBalance := balanceFloat / float64(10e17)

		wrapBalanceResponse, err := wrapTokenContract.BalanceOf(&bind.CallOpts{Pending: false, Context: context.Background()}, common.HexToAddress(ethereumAddress))
		if err != nil {
			log.Fatal(err)
		}
		wrapBalanceFloat, _ := new(big.Rat).SetInt(wrapBalanceResponse).Float64()
		wrapBalance := wrapBalanceFloat / float64(10e17)

		ethereumInsert := `UPDATE dydx_holders SET ethdydx_balance = $1, wethdydx_balance = $2 WHERE ethereum_address = $3`

		_, balanceErr := db.Exec(ethereumInsert, ethBalance, wrapBalance, ethereumAddress)
		if balanceErr != nil {
			log.Fatal(balanceErr)
		}
	}

}
