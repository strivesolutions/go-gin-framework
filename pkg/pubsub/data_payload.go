package pubsub

import (
	"encoding/json"
	"fmt"

	"github.com/cloudevents/sdk-go/v2/event"
)

type DataRequestPayload struct {
	CorrelationId string `json:"correlationId"`
	UserIdToken   string `json:"userIdToken,omitempty"`
	Data          []byte `json:"data"`
}

func UnwrapDataRequestPayload(e event.Event, to interface{}) (string, error) {
	correlationId, _, err := UnwrapDataRequestPayloadForUser(e, to)
	return correlationId, err
}

func UnwrapDataRequestPayloadForUser(e event.Event, to interface{}) (string, string, error) {
	var message DataRequestPayload
	err := e.DataAs(&message)

	if err != nil {
		return "", "", fmt.Errorf("failed to unwrap event data: %v", err)
	}

	err = message.DataAs(&to)

	if err != nil {
		return message.CorrelationId, message.UserIdToken, err
	}

	return message.CorrelationId, message.UserIdToken, nil
}

func (p DataRequestPayload) DataAs(target interface{}) error {
	return json.Unmarshal(p.Data, &target)
}
