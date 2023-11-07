package main

import (
	bridge "dydx-bridge-monitor/pkg/contracts"
	"encoding/csv"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
)

func main() {
	// Setup the database connection and insert command
	// db, err := sqlx.Connect("pgx", "postgresql://doadmin:AVNS_A3mgPzwuR6swL6WYcPs@db-postgresql-nyc1-02376-do-user-3592284-0.b.db.ondigitalocean.com:25060/dydx?sslmode=require")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// insertStmt := `INSERT INTO bridge_events (id, sender, receiver, amount) VALUES ($1, $2, $3, $4) ON CONFLICT (id) DO NOTHING`

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

	// Open or create a CSV file for reading and writing
	csvFilePath := "bridge_events.csv"
	csvFile, err := os.OpenFile(csvFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Read existing data to avoid duplicates
	existingData := make(map[string]bool)
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		// Assuming the event ID is in the first column
		existingData[record[0]] = true
	}

	// Create a CSV writer
	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = ',' // Set the CSV delimiter

	// Loop through all the events to find any new bridge transfers
	for bridgeTransfers.Next() {
		event := bridgeTransfers.Event

		// Convert the values as needed for reading and inserting
		eventID := event.Id.Int64()

		// Check for duplicates based on event ID
		if existingData[strconv.FormatInt(eventID, 10)] {
			continue // Skip duplicates
		}

		dydxAmount, _ := new(big.Rat).SetInt(event.Amount).Float64()
		dydxAmount = dydxAmount / float64(10e17)
		cosmosAddress, _ := bech32.ConvertAndEncode("dydx", event.AccAddress)

		// Add new transfers to the CSV file
		data := []string{
			strconv.FormatInt(eventID, 10),
			event.From.String(),
			cosmosAddress,
			strconv.FormatFloat(dydxAmount, 'f', -1, 64),
		}
		err = csvWriter.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Flush any buffered data to the CSV file
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Fatal(err)
	}
}
