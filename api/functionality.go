package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// GetBTCPrice returns the current price of BTC in UAH.
func GetBTCPrice() float64 {
	url := "https://data.binance.com/api/v3/ticker/24hr?symbol=BTCUAH"
	response, _ := http.Get(url)

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	var binanceResponse BinanceResponse
	_ = json.Unmarshal(body, &binanceResponse)

	openPrice := binanceResponse.OpenPrice
	result, _ := strconv.ParseFloat(openPrice, 64)
	return result
}

func LoadEmails(path string) []EmailResponse {
	file, _ := os.ReadFile(path)
	var emails []EmailResponse
	_ = json.Unmarshal(file, &emails)
	return emails
}

func SaveEmails(path string, subscriptions []EmailResponse) bool {
	file, _ := json.Marshal(subscriptions)
	_ = os.WriteFile(path, file, 0644)
	return true
}

// IsUnique add a new email to the file if it is not already there.
// Returns true if the email was added, false if it was already there.
func IsUnique(subscriptions []EmailResponse, response EmailResponse) bool {
	for _, subscription := range subscriptions {
		if subscription.Email == response.Email {
			return false
		}
	}

	return true
}

// SendEmails NOT IMPLEMENTED.
// Sends emails to all the emails in the file.
// Returns true if the emails were sent, false if there was an error.
func SendEmails(to string) bool {
	// TODO: implement
	fmt.Println("Sending emails not implemented")
	return false
}
