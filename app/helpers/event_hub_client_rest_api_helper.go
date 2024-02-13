package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var EventHubClientRESTAPIHelper = newEventHubClientRESTAPIHelper()

type eventHubClientRESTAPIHelper struct {
}

func newEventHubClientRESTAPIHelper() eventHubClientRESTAPIHelper {
	return eventHubClientRESTAPIHelper{}
}

func (_ eventHubClientRESTAPIHelper) SendOTPMessageToMobileUser(senderID string, messageUrl string,
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
		if res.StatusCode == 200 {
			var responseMap map[string]interface{}
			json.NewDecoder(res.Body).Decode(&responseMap)
			if _, exist := responseMap["messages"].([]interface{})[0].(map[string]interface{})["messageId"]; exist {
				return json.Marshal(responseMap)
			}
			return nil, errors.New("Message not sent something went wrong")
		} else {
			return nil, errors.New("Invalid response from the server")
		}

	} else {
		return nil, errors.New("Invalid request from the server")
	}
}
