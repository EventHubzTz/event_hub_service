package utils

import (
	"fmt"
	"math/rand"
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
