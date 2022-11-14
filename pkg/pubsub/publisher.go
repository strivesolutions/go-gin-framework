package pubsub

import (
	"context"
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
	event := DataRequestPayload{
		CorrelationId: correlationId,
		Data:          eventData,
	}

	return PublishEvent(ctx, topic, event)
}
