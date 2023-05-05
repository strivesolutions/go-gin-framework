package state

import (
	"context"
	"fmt"
	"strconv"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

var client dapr.Client
var stateStoreName string

func Setup(name string) error {
	stateStoreName = name

	return createClient()
}

func Dispose() {
	if client != nil {
		client.Close()
	}
}

type StoreRead func(ctx context.Context, key string) ([]byte, error)
type StoreWrite func(ctx context.Context, key string, value []byte, ttlInSeconds int) error

func createClient() error {
	c, err := dapr.NewClient()

	if err != nil {
		logging.Error("Error creating Dapr client: %s", err)
		return err
	}
	client = c

	return nil
}

func getClient() dapr.Client {
	if client == nil {
		createClient()
	}

	return client
}

func Read(ctx context.Context, key string) ([]byte, error) {
	result, err := getClient().GetState(ctx, stateStoreName, key, nil)

	if err != nil {
		logging.Error(fmt.Sprintf("Error reading document: %s", err))
		return nil, err
	}

	if result == nil || result.Value == nil {
		return nil, nil
	}

	return result.Value, nil
}

func Write(ctx context.Context, key string, value []byte, ttlInSeconds int) error {
	meta := map[string]string{
		"ttlInSeconds": strconv.Itoa(ttlInSeconds),
	}

	return getClient().SaveState(ctx, stateStoreName, key, value, meta)
}
