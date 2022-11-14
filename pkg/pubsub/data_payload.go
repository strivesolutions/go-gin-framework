package pubsub

type DataRequestPayload struct {
	CorrelationId string      `json:"correlationId"`
	Data          interface{} `json:"data"`
}
