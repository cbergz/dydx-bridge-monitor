package main

import (
	bridge "dydx-bridge-monitor/pkg/contracts"
	"dydx-bridge-monitor/pkg/dune"
	"encoding/csv"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// Load API keys from env file
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error loading env file")
	}
	infuraKey := os.Getenv("INFURA_KEY")
	duneKey := os.Getenv("DUNE_KEY")
	filePath := os.Getenv("FILEPATH")

	// Setup the connection to Ethereum using Infura RPC
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + infuraKey)
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
	csvFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
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
		existingData[record[0]] = true
	}

	// Create a CSV writer
	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = ',' // Set the CSV delimiter

	// Check if the file is empty, and if so, write the header
	fileInfo, err := csvFile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if fileInfo.Size() == 0 {
		headers := []string{"evt_id", "evt_tx_hash", "from", "accAddress", "amount"}
		err = csvWriter.Write(headers)
		if err != nil {
			log.Fatal(err)
		}
		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			log.Fatal(err)
		}
	}

	// Loop through all the events to find any new bridge transfers
	for bridgeTransfers.Next() {
		event := bridgeTransfers.Event

		// Convert the values as needed for reading and inserting
		eventID := event.Id.String()

		// Check for duplicates based on event ID
		if existingData[eventID] {
			continue // Skip duplicates
		}

		dydxAmount, _ := new(big.Rat).SetInt(event.Amount).Float64()
		dydxAmount = dydxAmount / float64(10e17)
		cosmosAddress, _ := bech32.ConvertAndEncode("dydx", event.AccAddress)

		// Add new transfers to the CSV file
		data := []string{
			eventID,
			event.Raw.TxHash.Hex(),
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

	// Upload the final csv file to dune
	dune.UploadToDune(filePath, duneKey)
}
