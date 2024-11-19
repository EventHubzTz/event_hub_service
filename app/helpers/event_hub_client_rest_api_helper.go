package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var EventHubClientRESTAPIHelper = newEventHubClientRESTAPIHelper()

type eventHubClientRESTAPIHelper struct {
}

func newEventHubClientRESTAPIHelper() eventHubClientRESTAPIHelper {
	return eventHubClientRESTAPIHelper{}
}

func (q eventHubClientRESTAPIHelper) SendOTPMessageToMobileUser(senderID string, messageUrl string,
	authorizationToken string, phoneNo string, message string) ([]byte, error) {
	type RequestBody struct {
		From string `json:"from"`
		To   string `json:"to"`
		Text string `json:"text"`
	}
	var requestBody RequestBody
	requestBody.From = senderID
	requestBody.To = phoneNo
	requestBody.Text = message
	requestByte, _ := json.Marshal(requestBody)
	client := http.Client{}
	req, err := http.NewRequest("POST", messageUrl, bytes.NewBuffer(requestByte))
	if err != nil {
		return nil, err
	}
	if req != nil {
		req.Header = http.Header{
			"Accept":        []string{"application/json"},
			"Content-Type":  []string{"application/json"},
			"Authorization": []string{"Basic " + authorizationToken},
		}
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if res.StatusCode == 200 {
			var responseMap map[string]interface{}
			if err := json.Unmarshal(body, &responseMap); err != nil {
				return nil, err
			}
			if _, exist := responseMap["messages"].([]interface{})[0].(map[string]interface{})["messageId"]; exist {
				return json.Marshal(responseMap)
			}
			return nil, fmt.Errorf("status: %d, Response: %s", res.StatusCode, string(body))
		} else {
			return nil, fmt.Errorf("status: %d, Response: %s", res.StatusCode, string(body))
		}

	} else {
		return nil, errors.New("invalid request from the server")
	}
}

func MobiSMSApi(senderID string, messageUrl string, authorizationToken string, phoneNo string, message string) ([]byte, string, error) {
	encodedWord := url.QueryEscape(message)
	url := messageUrl + "?user=ALECOtr&pwd=" + authorizationToken + "&senderid=" + senderID + "&mobileno=" + phoneNo[len(phoneNo)-9:] + "&msgtext=" + encodedWord + "&priority=High&CountryCode=255"

	response, err := http.Get(url)
	if err != nil {
		return nil, url, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
	if err != nil {
		return nil, url, err
	}

	return body, url, nil
}

type GenerateAzamPayTokenResponse struct {
	Data struct {
		AccessToken string    `json:"accessToken"`
		Expire      time.Time `json:"expire"`
	} `json:"data"`
	Message    string `json:"message"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
}

func GenerateAzamPayToken(url, appName, clientId, clientSecret, apiKey string) (*GenerateAzamPayTokenResponse, error) {

	type Request struct {
		AppName      string `json:"appName"`
		ClientID     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
	}

	var request Request

	request.AppName = appName
	request.ClientID = clientId
	request.ClientSecret = clientSecret

	method := "POST"

	requestByte, _ := json.Marshal(request)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestByte))

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var results GenerateAzamPayTokenResponse
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

type AzamPayPushUSSDResponse struct {
	Errors        map[string][]string `json:"errors"`
	Type          string              `json:"type"`
	Title         string              `json:"title"`
	Status        int                 `json:"status"`
	TraceID       string              `json:"traceId"`
	TransactionID string              `json:"transactionId"`
	Message       string              `json:"message"`
	Success       bool                `json:"success"`
	Results       string              `json:"results"`
}

func AzamPayPushUSSD(url, accountNumber, amount, currency, externalId, provider, bearerToken, apiKey string) (*AzamPayPushUSSDResponse, error) {

	type Request struct {
		AccountNumber string `json:"accountNumber"`
		Amount        string `json:"amount"`
		Currency      string `json:"currency"`
		ExternalID    string `json:"externalId"`
		Provider      string `json:"provider"`
	}

	var request Request

	request.AccountNumber = accountNumber
	request.Amount = amount
	request.Currency = currency
	request.ExternalID = externalId
	request.Provider = provider

	method := "POST"

	requestByte, _ := json.Marshal(request)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestByte))

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearerToken)
	req.Header.Add("X-API-Key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var results AzamPayPushUSSDResponse
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, err
	}

	results.Results = string(body)

	return &results, nil
}

type PushUSSDResponse struct {
	Error bool `json:"error"`
	Data  Data `json:"data"`
}

type Data struct {
	Results       string `json:"results"`
	Message       string `json:"message"`
	OrderID       string `json:"order_id"`
	TransactionID string `json:"transaction_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
}

func PushUSSD(url, PhoneNumber string, amount float32) (*PushUSSDResponse, error) {

	type Request struct {
		PhoneNumber string  `json:"phone_number"`
		Amount      float32 `json:"amount"`
	}

	var request Request

	request.PhoneNumber = PhoneNumber
	request.Amount = amount

	method := "POST"

	requestByte, _ := json.Marshal(request)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestByte))

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("event-hub-sign-auth", "XhoO2yoeISBAJja8AGuul0hYomoEkXKK")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var results PushUSSDResponse
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, err
	}

	results.Data.Results = string(body)

	return &results, nil
}
