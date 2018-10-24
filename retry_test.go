package retry

import (
	"testing"
	"time"

	"github.com/ansel1/merry"
	"github.com/stretchr/testify/assert"
)

var (
	executionDelta = float64(10 * time.Millisecond)
	dummyError     = merry.New("expected error")
)

func TestTryNoError(t *testing.T) {
	err := Re{}.Try(func() error {
		return nil
	})

	assert.NoError(t, err)
}

func TestTryUnexpectedError(t *testing.T) {
	err := Re{Max: 2, RetryableErrors: []merry.Error{dummyError}}.Try(func() error {
		return merry.New("unexpected error")
	})

	assert.Error(t, err)
}

func TestTryExpectedError(t *testing.T) {
	err := Re{Max: 2, RetryableErrors: []merry.Error{dummyError}}.Try(func() error {
		return dummyError.Here()
	})

	assert.Error(t, err)
}

func TestTryCalled5Times(t *testing.T) {
	i := uint(0)
	max := uint(5)
	err := Re{Max: max}.Try(func() error {
		i++
		return merry.New("unexpected error")
	})

	assert.Error(t, err)
	assert.Equal(t, max, i)
}

func TestTryCalled1TimesWhenMaxIs0(t *testing.T) {
	i := 0
	max := uint(0)
	err := Re{Max: max}.Try(func() error {
		i++
		return merry.New("unexpected error")
	})

	assert.Error(t, err)
	assert.Equal(t, 1, i)
}

func TestTryDelay(t *testing.T) {
	i := 0
	delay := time.Duration(100 * time.Millisecond)

	var called, calledPrevious time.Time
	err := Re{Delay: delay}.Try(func() error {
		switch i {
		case 1:
			called = time.Now()
			return nil
		default:
			i++
			calledPrevious = time.Now()
			return merry.New("unexpected error")
		}
	})

	assert.NoError(t, err)
	assert.InDelta(t, called.Sub(calledPrevious), delay, executionDelta)
}

func TestTryBackoff(t *testing.T) {
	i := 0
	delay := time.Duration(100 * time.Millisecond)
	backoff := float64(1.2)

	var called, calledPrevious time.Time
	err := Re{Delay: delay, Backoff: backoff}.Try(func() error {
		switch i {
		case 2:
			called = time.Now()
			return nil
		default:
			i++
			calledPrevious = time.Now()
			return merry.New("unexpected error")
		}
	})

	assert.NoError(t, err)
	assert.InDelta(t, called.Sub(calledPrevious), time.Duration(float64(delay)*backoff), executionDelta)
}

func TestTryMaxDelay(t *testing.T) {
	i := 0
	delay := time.Duration(100 * time.Millisecond)
	maxDelay := time.Duration(200 * time.Millisecond)
	backoff := float64(5)

	var called, calledPrevious time.Time
	err := Re{Delay: delay, Backoff: backoff, MaxDelay: maxDelay}.Try(func() error {
		switch i {
		case 2:
			called = time.Now()
			return nil
		default:
			i++
			calledPrevious = time.Now()
			return merry.New("unexpected error")
		}
	})

	assert.NoError(t, err)
	assert.InDelta(t, called.Sub(calledPrevious), maxDelay, executionDelta)
}
