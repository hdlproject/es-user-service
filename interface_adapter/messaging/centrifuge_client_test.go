package messaging

import (
	"context"
	"testing"
	"time"
)

func TestCentrifugeClient_Publish(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := GetCentrifugeClient(
		"ws://localhost:8000/connection/websocket",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0IiwiZXhwIjoxNjg2OTczNTgwLCJpYXQiOjE2ODYzNjg3ODB9.B4rHk9pkS5Wqu4Qg2RMNkfG9dg9U0h-LlE72Pa2beDM",
		"")
	if err != nil {
		t.Fatal(err)
	}

	channel := "test_channel"
	message := `{"message": "test"}`

	c := make(chan string, 1)
	receiverFunc := func(m string) {
		c <- m
	}

	err = client.Subscribe(channel, receiverFunc)
	if err != nil {
		t.Fatal(err)
	}

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
