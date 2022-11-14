package pubsub

import (
	"encoding/json"
	"fmt"

	"github.com/cloudevents/sdk-go/v2/event"
)

type DataRequestPayload struct {
	CorrelationId string `json:"correlationId"`
	Data          []byte `json:"data"`
}

func UnwrapDataRequestPayload[T interface{}](e event.Event, to *T) (string, error) {
	var message DataRequestPayload
	err := e.DataAs(&message)

	if err != nil {
		return "", fmt.Errorf("failed to unwrap event data: %v", err)
	}

	err = message.DataAs(&to)

	if err != nil {
		return message.CorrelationId, err
	}

	return message.CorrelationId, nil
}

func (p DataRequestPayload) DataAs(target interface{}) error {
	return json.Unmarshal(p.Data, &target)
}
