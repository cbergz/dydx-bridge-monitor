package main

import (
	bridge "dydx-bridge-monitor"
	"log"
	"math/big"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/jmoiron/sqlx"
)

func main() {
	// Setup the database connection and insert command
	db, err := sqlx.Connect("pgx", "postgresql://doadmin:AVNS_A3mgPzwuR6swL6WYcPs@db-postgresql-nyc1-02376-do-user-3592284-0.b.db.ondigitalocean.com:25060/dydx?sslmode=require")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	insertStmt := `INSERT INTO bridge_events (id, sender, receiver, amount) VALUES ($1, $2, $3, $4) ON CONFLICT (id) DO NOTHING`

	// Setup the connection to Ethereum using Infura RPC
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/f61acb8dac8b4fa5bd201f129cc37a67")
	if err != nil {
		log.Fatal(err)
	}

	// Define the bridge contract we're looking for
	bridgeAddress := common.HexToAddress("0x46b2deae6eff3011008ea27ea36b7c27255ddfa9")
	bridgeContract, err := bridge.NewBridgeFilterer(bridgeAddress, client)
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
		eventId := event.Id.Int64()
		dydxAmount, _ := new(big.Rat).SetInt(event.Amount).Float64()
		dydxAmount = dydxAmount / float64(10e17)
		cosmosAddress, _ := bech32.ConvertAndEncode("dydx", event.AccAddress)

		// Add new transfers to the database
		_, err := db.Exec(insertStmt, eventId, event.From.String(), cosmosAddress, dydxAmount)
		if err != nil {
			log.Fatal(err)
		}
	}
}
