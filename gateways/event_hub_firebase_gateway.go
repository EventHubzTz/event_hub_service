package gateways

import (
	"context"
	"errors"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var EventHubFirebaseGateway = newEventHubFirebaseGateway()

type eventHubFirebaseGateway struct{}

func newEventHubFirebaseGateway() eventHubFirebaseGateway {
	return eventHubFirebaseGateway{}
}

var ctx = context.Background()

func (f eventHubFirebaseGateway) InitApp() (*firebase.App, error) {
	firebaseAUthCredentialFile := os.Getenv("AUTH_CREDENTIAL_FILE")
	opt := option.WithCredentialsFile(firebaseAUthCredentialFile)
	/*-------------------------------
	 01. APPLICATION INITIALIZATION
	--------------------------------*/
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, errors.New("error initializing app")
	}
	return app, nil
}

func (f eventHubFirebaseGateway) InitAuth() (*auth.Client, error) {
	app, err := f.InitApp()
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (f eventHubFirebaseGateway) InitFireStore() (*firestore.Client, error) {
	app, err := f.InitApp()
	if err != nil {
		return nil, err
	}
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return firestoreClient, nil
}
