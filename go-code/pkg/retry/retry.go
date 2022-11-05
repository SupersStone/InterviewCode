package retry

import (
	"time"

	"github.com/cenkalti/backoff/v4"
)

func newBackoffPolicy() *backoff.ExponentialBackOff {
	policy := backoff.NewExponentialBackOff()
	policy.InitialInterval = 20 * time.Millisecond
	policy.MaxElapsedTime = 15 * time.Second
	return policy
}

// BackoffRetry 初始20ms间隔，最大持续15s
func BackoffRetry(f func() error) error {
	return backoff.Retry(f, newBackoffPolicy())
}

func BackoffRetryWithPolicy(f func() error, p *backoff.ExponentialBackOff) error {
	return backoff.Retry(f, p)
}

// Permanent 包裹错误，返回此错误不再重试
func Permanent(err error) error {
	return backoff.Permanent(err)
}
