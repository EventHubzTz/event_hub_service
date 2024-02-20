package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	constants2 "github.com/EventHubzTz/event_hub_service/utils/constants"
)

func SuccessPrint(message string) {
	fmt.Println(constants2.ColorGreen, "[✓]", message, constants2.ColorReset)
}

func WarningPrint(message string) {
	fmt.Println(constants2.ColorYellow, "[!]", message, constants2.ColorReset)
}

func InfoPrint(message string) {
	fmt.Println(constants2.ColorBlue, "[ℹ]", message, constants2.ColorReset)
}

func ErrorPrint(message string) {
	fmt.Println(constants2.ColorRed, "[✘]", message, constants2.ColorReset)
}

func FormatString(error string) string {
	format := "======"
	for i := 0; i < len(error); i++ {
		format += "="
	}
	return format
}

func GenerateOrderId() string {
	rand.Seed(time.Now().UnixNano())

	// Generate three random numbers in the range 0-99999
	num1 := rand.Intn(100000)
	num2 := rand.Intn(100000)
	num3 := rand.Intn(100000)

	// Format the numbers with leading zeros and separate them with hyphens
	uniqueString := fmt.Sprintf("%05d-%05d-%05d", num1, num2, num3)

	return uniqueString
}

func CheckMobileNetwork(mobileNumber string) string {
	// Define the prefixes for each network
	networkPrefixes := map[string][]int{
		"Airtel":   {68, 69, 78},
		"Tigo":     {65, 67, 71},
		"Halopesa": {62},
		"Azampesa": {0}, // Considering 0 as prefix for Azampesa, you may adjust as necessary
		"Mpesa":    {74, 75, 76},
	}

	// Extract the first few digits from the mobile number
	prefix, err := strconv.Atoi(mobileNumber[:3])
	if err != nil {
		return toJSONString(map[string]string{"error": "Invalid mobile number"})
	}

	// Check which network the mobile number belongs to
	network := getNetworkByPrefix(prefix, networkPrefixes)

	// Ensure a string is always returned
	if network != "" {
		return network
	} else {
		return toJSONString(map[string]string{"error": "Network not recognized"})
	}
}

func getNetworkByPrefix(prefix int, networkPrefixes map[string][]int) string {
	for network, prefixes := range networkPrefixes {
		for _, p := range prefixes {
			if prefix == p {
				return network
			}
		}
	}
	return "" // If the prefix doesn't match any network
}

func toJSONString(data interface{}) string {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintf(`{"error": "JSON marshalling error: %s"}`, err)
	}
	return string(jsonString)
}
