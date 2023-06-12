package messaging

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestCentrifugeClient_Publish(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	channel := "test_channel"
	message := `{"message": "test"}`
	userId := "test_user"
	secret := "token_secret"

	clientToken := generateClientJWT(t, userId, secret)
	client, err := GetCentrifugeClient(
		"ws://localhost:8000/connection/websocket",
		clientToken,
		"")
	if err != nil {
		t.Fatal(err)
	}

	c := make(chan string, 1)
	receiverFunc := func(m string) {
		c <- m
	}

	subsToken := generateSubscriptionJWT(t, userId, channel, secret)
	err = client.Subscribe(channel, subsToken, receiverFunc)
	if err != nil {
		t.Fatal(err)
	}

	// let the subscription created on Centrifugo server
	time.Sleep(1 * time.Second)

	err = client.Publish(ctx, channel, message)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case m := <-c:
		if m != message {
			t.Fatalf("expect message %s but got %s", message, m)
		}
	case <-ctx.Done():
		t.Fatalf("test timeout")
	}
}

func generateClientJWT(t *testing.T, userId, secret string) string {
	claims := CentrifugeJWTClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userId,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	return generateJWT(t, claims, secret)
}

func generateSubscriptionJWT(t *testing.T, userId, channel, secret string) string {
	claims := CentrifugeJWTClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userId,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		Channel: channel,
	}

	return generateJWT(t, claims, secret)
}

func generateJWT(t *testing.T, claims jwt.Claims, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatal(err)
	}

	return signedToken
}
