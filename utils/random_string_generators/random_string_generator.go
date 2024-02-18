package random_string_generators

import (
	"math/rand"
	"time"
)

/**
 * Created by GoLand.
* Project : event_hub_service
* User: LAZARO WILLY
* Date: 01/02/2022
* Time: 10:54
*/
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func GeneratePromoCode(n int) string {
	var letters = []rune("0123456789")
	rand.Seed(time.Now().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
