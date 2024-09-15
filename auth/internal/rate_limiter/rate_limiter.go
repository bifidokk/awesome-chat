package rate_limiter

import (
	"context"
	"time"
)

// TokenBucketLimiter implements a token bucket rate-limiting algorithm.
type TokenBucketLimiter struct {
	tokenBucketChannel chan struct{}
}

// NewTokenBucketLimiter creates a new token bucket limiter with a specified limit and refill period
func NewTokenBucketLimiter(ctx context.Context, limit int, period time.Duration) *TokenBucketLimiter {
	limiter := &TokenBucketLimiter{
		tokenBucketChannel: make(chan struct{}, limit),
	}

	for i := 0; i < limit; i++ {
		limiter.tokenBucketChannel <- struct{}{}
	}

	refillInterval := period.Nanoseconds() / int64(limit)
	go limiter.startPeriodicRefilling(ctx, time.Duration(refillInterval))

	return limiter
}

// Allow checks if a token is available in the token bucket; if so, it allows the request, otherwise it denies it.
func (limiter *TokenBucketLimiter) Allow() bool {
	select {
	case <-limiter.tokenBucketChannel:
		return true
	default:
		return false
	}
}

func (limiter *TokenBucketLimiter) startPeriodicRefilling(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			limiter.tokenBucketChannel <- struct{}{}
		}
	}
}
