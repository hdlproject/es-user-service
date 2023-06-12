package messaging

import (
	"context"

	"github.com/centrifugal/centrifuge-go"
	"github.com/golang-jwt/jwt"

	"github.com/hdlproject/es-user-service/helper"
)

type CentrifugeClient struct {
	client          *centrifuge.Client
	token           string
	refreshTokenUrl string
	subscriptions   map[string]*centrifuge.Subscription
}

type CentrifugeJWTClaims struct {
	jwt.StandardClaims
	Channel string `json:"channel,omitempty"`
}

var centrifugeClient *CentrifugeClient

func GetCentrifugeClient(serverUrl, token, refreshTokenUrl string) (*CentrifugeClient, error) {
	if centrifugeClient == nil {
		client, err := newCentrifugeClient(serverUrl, token, refreshTokenUrl)
		if err != nil {
			return nil, helper.WrapError(err)
		}

		centrifugeClient = client
	}

	return centrifugeClient, nil
}

func newCentrifugeClient(serverUrl, token, refreshTokenUrl string) (*CentrifugeClient, error) {
	client := centrifuge.NewJsonClient(
		serverUrl,
		centrifuge.Config{
			Token: token,
			GetToken: func(m centrifuge.ConnectionTokenEvent) (string, error) {
				return refreshTokenFunc(token)
			},
		},
	)

	err := client.Connect()
	if err != nil {
		return nil, helper.WrapError(err)
	}

	return &CentrifugeClient{
		client:          client,
		token:           token,
		refreshTokenUrl: refreshTokenUrl,
		subscriptions:   make(map[string]*centrifuge.Subscription),
	}, nil
}

func (instance *CentrifugeClient) Subscribe(channel, token string, receiverFunc func(m string)) (err error) {
	if instance.subscriptions[channel] == nil {
		instance.subscriptions[channel], err = instance.client.NewSubscription(
			channel,
			centrifuge.SubscriptionConfig{
				Token: token,
			},
		)
		if err != nil {
			return helper.WrapError(err)
		}
	}

	instance.subscriptions[channel].OnPublication(func(handler centrifuge.PublicationEvent) {
		receiverFunc(string(handler.Data))
	})

	err = instance.subscriptions[channel].Subscribe()
	if err != nil {
		return helper.WrapError(err)
	}

	return nil
}

func (instance *CentrifugeClient) Publish(ctx context.Context, channel string, m string) (err error) {
	if instance.subscriptions[channel] != nil {
		_, err = instance.subscriptions[channel].Publish(ctx, []byte(m))
		if err != nil {
			return helper.WrapError(err)
		}
	} else {
		_, err = instance.client.Publish(ctx, channel, []byte(m))
		if err != nil {
			return helper.WrapError(err)
		}
	}

	return nil
}

func refreshTokenFunc(currentToken string) (string, error) {
	// TODO: refresh token here
	return currentToken, nil
}
