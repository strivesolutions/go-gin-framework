package lock

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

type StoreLock func(lockKey string) (LockDefinition, error)

var lockOwner string
var LockExpirationSeconds int32 = 1
var MaxAquisitionAttempts int = 32

var client dapr.Client
var lockStoreName string

const (
	minRetryDelayMilliSec = 50
	maxRetryDelayMilliSec = 250
)

func Setup(name, owner string) error {
	lockStoreName = name
	lockOwner = owner

	return createClient()
}

type LockDefinition interface {
	AcquireLock() error
	Unlock() error
}

type distributedLock struct {
	lockKey string
}

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

func delay(tries int) time.Duration {
	return time.Duration(rand.Intn(maxRetryDelayMilliSec-minRetryDelayMilliSec)+minRetryDelayMilliSec) * time.Millisecond
}

func CreateLock(lockKey string) LockDefinition {
	return distributedLock{
		lockKey: lockKey,
	}
}

func (l distributedLock) AcquireLock() error {
	req := dapr.LockRequest{
		LockOwner:       lockOwner,
		ResourceID:      l.lockKey,
		ExpiryInSeconds: LockExpirationSeconds,
	}

	ctx := context.Background()

	var resp *dapr.LockResponse
	var err error

	for i := 0; i < MaxAquisitionAttempts; i++ {
		if i != 0 {
			<-time.After(delay(i))
		}

		resp, err = getClient().TryLockAlpha1(ctx, lockStoreName, &req)

		if err != nil {
			break
		}

		if resp.Success {
			break
		}
	}

	if err != nil {
		logging.ErrorObject(fmt.Errorf("failed to aquire lock: %v", err))
		return err
	}

	if !resp.Success {
		logging.Warn(fmt.Sprintf("lock response not successful: %v", resp.Success))
		return errors.New("failed to aquire lock")
	}

	return nil
}

func (l distributedLock) Unlock() error {
	req := dapr.UnlockRequest{
		LockOwner:  lockOwner,
		ResourceID: l.lockKey,
	}

	ctx := context.Background()
	_, err := getClient().UnlockAlpha1(ctx, lockStoreName, &req)

	return err
}
