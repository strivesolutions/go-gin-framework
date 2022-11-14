package tests

import (
	"encoding/json"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/stretchr/testify/assert"
	"github.com/strivesolutions/go-gin-framework/pkg/pubsub"
)

func TestUnwrapDataRequestPayload(t *testing.T) {
	type testData map[string]interface{}

	data, _ := json.Marshal(testData{
		"foo": "bar",
	})

	p := pubsub.DataRequestPayload{
		CorrelationId: "123",
		Data:          data,
	}

	data, _ = json.Marshal(p)

	e := event.Event{
		Context:     &event.EventContextV1{},
		DataEncoded: data}

	var result testData
	correlationId, err := pubsub.UnwrapDataRequestPayload(e, &result)

	assert.NoError(t, err)
	assert.Equal(t, "123", correlationId)
	assert.Equal(t, "bar", result["foo"])
}
