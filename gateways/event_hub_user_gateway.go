package gateways

var EventHubUserGateway = newEventHubUserGateway()

type eventHubUserGateway struct {
}

func newEventHubUserGateway() eventHubUserGateway {
	return eventHubUserGateway{}
}
