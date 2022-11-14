package pubsub

import (
	"context"
	"encoding/json"
	"errors"

	dapr "github.com/dapr/go-sdk/client"
)

func PublishEvent(ctx context.Context, topic string, event interface{}) error {
	if PubsubName == "" {
		return errors.New("pubsub name was not configured on server initialization")
	}

	client, err := dapr.NewClient()
	if err != nil {
		return err
	}

	return client.PublishEvent(ctx, PubsubName, topic, event)
}

func PublishDataPayload(ctx context.Context, topic string, correlationId string, eventData interface{}) error {
	data, err := json.Marshal(eventData)

	if err != nil {
		return err
	}

	event := DataRequestPayload{
		CorrelationId: correlationId,
		Data:          data,
	}

	return PublishEvent(ctx, topic, event)
}
