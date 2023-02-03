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

func TestUnwrapDataRequestPayloadFromGo(t *testing.T) {
	type testData map[string]interface{}

	data := []byte("{\"data\": {\"correlationId\": \"a652223a-0c6a-4619-927a-5160365d4b8c\", \"data\": \"eyJpZCI6ImE2NTIyMjNhLTBjNmEtNDYxOS05MjdhLTUxNjAzNjVkNGI4YyIsInN0YXR1cyI6ImZldGNoaW5nX2RhdGEiLCJmYWlsdXJlUmVhc29uIjoiIiwidHJ1c3RGdW5kSWQiOjQ2MCwiY29udGFjdFJlY29yZElkIjoxMjI0OTYsInJldGlyZW1lbnRBZ2UiOjU1LCJyZXRpcmVtZW50RGF0ZSI6bnVsbCwid2l0aFNwb3VzZSI6ZmFsc2UsInNwb3VzZURhdGVPZkJpcnRoIjpudWxsLCJwZW5zaW9uRGF0YSI6bnVsbCwiY3BwTGltaXQiOm51bGwsIm9hc0xpbWl0IjpudWxsLCJyZXN1bHQiOm51bGx9\"}, \"datacontenttype\": \"application/json\", \"id\": \"502c791c-7f54-4cd1-be8f-9c0613e11364\", \"pubsubname\": \"rabbitmq-pubsub\", \"source\": \"pension-estimate\", \"specversion\": \"1.0\", \"time\": \"2022-12-12T21:26:28Z\", \"topic\": \"estimate_needs_data\", \"traceid\": \"00-00000000000000000000000000000000-0000000000000000-00\", \"traceparent\": \"00-00000000000000000000000000000000-0000000000000000-00\", \"tracestate\": \"\", \"type\": \"com.dapr.event.sent\"}")

	var e event.Event
	err := json.Unmarshal(data, &e)
	assert.NoError(t, err)

	var result testData
	correlationId, err := pubsub.UnwrapDataRequestPayload(e, &result)
	assert.NoError(t, err)
	assert.Equal(t, "a652223a-0c6a-4619-927a-5160365d4b8c", correlationId)
}
