// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// OCIRetry represents a base interface we use in the Retry function.
type OCIRetry interface {
	// This validates whether it's a valid retry
	validate() (bool, error)

	// True for enable retry, false for no retry
	Enabled() bool

	// used in retry loop, this determines the retry count
	ShouldContinueRetryV2(uint) bool

	// This is to determine whether we should retry on specific status code & error code
	ShouldRetryOperationV2(OCIOperationResponse) bool

	// wait time
	NextDurationV2(OCIOperationResponse) time.Duration

	// this is to maintain backward compatibility, we will need to implement for composingRetryPolicy, for other simple retryPolicy (ie. current retryPolicy), we'll use returnSamePolicy
	DeterminePolicyToUseV2() (OCIRetry, *time.Time, float64)

	// GetMaximumCumulativeBackoffWithoutJitter returns the maximum cumulative backoff the retry policy would do,
	// taking into account whether eventually consistency is considered or not.
	GetMaximumCumulativeBackoffWithoutJitter() time.Duration

	// SetRetry sets the ShouldRetryOperation
	SetRetry(func(OCIOperationResponse) bool)
}

// RetryPolicyV2 is the class that holds all relevant information for retrying operations.
type RetryPolicyV2 struct {
	// MaximumNumberAttempts is the maximum number of times to retry a request. Zero indicates an unlimited
	// number of attempts.
	MaximumNumberAttempts uint

	// minimum sleep between attempts in seconds
	MinSleepBetween float64

	// maximum sleep between attempts in seconds
	MaxSleepBetween float64

	// the base for the exponential backoff
	ExponentialBackoffBase float64

	// Stores the maximum cumulative backoff in seconds. This can usually be calculated using
	// MaximumNumberAttempts, MinSleepBetween, MaxSleepBetween, and ExponentialBackoffBase,
	// but if MaximumNumberAttempts is 0 (unlimited attempts), then this needs to be set explicitly
	// for Eventual Consistency retries to work.
	MaximumCumulativeBackoffWithoutJitter float64

	// ShouldRetryOperation inspects the http response, error, and operation attempt number, and
	// - returns true if we should retry the operation
	// - returns false otherwise
	ShouldRetryOperationFunc func(OCIOperationResponse) bool

	// used in retry loop, this determines the retry count
	ShouldContinueRetryFunc func(uint, uint) bool

	// wait time
	NextDurationFunc func(OCIOperationResponse, float64, float64, float64, uint) time.Duration

	// GetMaximumCumulativeBackoffWithoutJitterFunc returns the maximum cumulative backoff the retry policy would do,
	// taking into account whether eventually consistency is considered or not.
	GetMaximumCumulativeBackoffWithoutJitterFunc func(float64, float64, float64, uint, float64) time.Duration

	// True for enable retry, false for no retry
	Enable bool
}

// ShouldRetryOperationV2 is the function that should be used for RetryPolicy.ShouldRetryOperation when
// taking eventual consistency into account
func (rp *RetryPolicyV2) ShouldRetryOperationV2(r OCIOperationResponse) bool {
	return rp.ShouldRetryOperationFunc(r)
}

// NextDurationV2 computes the duration to pause between operation retries.
func (rp *RetryPolicyV2) NextDurationV2(r OCIOperationResponse) time.Duration {
	return rp.NextDurationFunc(r, rp.MinSleepBetween, rp.MaxSleepBetween, rp.ExponentialBackoffBase, rp.MaximumNumberAttempts)
}

// Enabled is set to true for retry policy and set to false for NoRetryPolicy
func (rp *RetryPolicyV2) Enabled() bool {
	return rp.Enable
}

//ShouldContinueRetryV2 is used in retry loop, this determines the retry count
func (rp *RetryPolicyV2) ShouldContinueRetryV2(current uint) bool {
	return rp.ShouldContinueRetryFunc(current, rp.MaximumNumberAttempts)
}

// DeterminePolicyToUseV2 may modify the policy to handle eventual consistency; the return values are
// the retry policy to use, the end of the eventually consistent time window, and the backoff scaling factor
// If eventual consistency is not considered, this function should return the unmodified policy that was
// provided as input, along with (*time.Time)(nil) (no time window), and 1.0 (unscaled backoff).
func (rp *RetryPolicyV2) DeterminePolicyToUseV2() (OCIRetry, *time.Time, float64) {
	eowt := EcContext.GetEndOfWindow()
	var backoffScalingFactor = 1.0
	return rp, eowt, backoffScalingFactor
}

// Validate returns true if the RetryPolicy is valid; if not, it also returns an error.
func (rp *RetryPolicyV2) validate() (success bool, err error) {
	var errorStrings []string

	if rp.ShouldRetryOperationFunc == nil {
		errorStrings = append(errorStrings, "ShouldRetryOperation may not be nil")
	}
	if rp.NextDurationFunc == nil {
		errorStrings = append(errorStrings, "NextDuration may not be nil")
	}
	if rp.MaximumNumberAttempts == 0 && rp.MaximumCumulativeBackoffWithoutJitter <= 0 {
		errorStrings = append(errorStrings, "If eventual consistency is handled, and the MaximumNumberAttempts of the EC retry policy is 0 (unlimited attempts), then the MaximumCumulativeBackoffWithoutJitter of the EC retry policy must be positive; used WithUnlimitedAttempts instead")
	}
	if len(errorStrings) > 0 {
		return false, errors.New(strings.Join(errorStrings, ", "))
	}
	return true, nil
}

// ComplexRetryPolicyV2 is the class that holds all relevant information for retrying operations in eventual consistency.
type ComplexRetryPolicyV2 struct {
	// This always represents the non-EC retryPolicy
	RetryPolicy OCIRetry
	// This always represents the EC retryPolicy, to simplify the implementation, we can continue using the RetryPolicy here
	EcRetryPolicy OCIRetry
}

// Validate returns true if the RetryPolicy is valid; if not, it also returns an error.
func (crp *ComplexRetryPolicyV2) validate() (success bool, err error) {
	policyToUse, _, _ := crp.DeterminePolicyToUseV2()
	return policyToUse.validate()
}

//ShouldContinueRetryV2 is used in retry loop, this determines the retry count
func (crp *ComplexRetryPolicyV2) ShouldContinueRetryV2(current uint) bool {
	policyToUse, _, _ := crp.DeterminePolicyToUseV2()
	return policyToUse.ShouldContinueRetryV2(current)
}

// Enabled is set to true for retry policy and set to false for NoRetryPolicy
func (crp *ComplexRetryPolicyV2) Enabled() bool {
	policyToUse, _, _ := crp.DeterminePolicyToUseV2()
	return policyToUse.Enabled()
}

// ShouldRetryOperationV2 is the function that should be used for RetryPolicy.ShouldRetryOperation when
// taking eventual consistency into account
func (crp *ComplexRetryPolicyV2) ShouldRetryOperationV2(r OCIOperationResponse) bool {
	policyToUse, _, _ := crp.DeterminePolicyToUseV2()
	return policyToUse.ShouldRetryOperationV2(r)
}

// NextDurationV2 computes the duration to pause between operation retries.
func (crp *ComplexRetryPolicyV2) NextDurationV2(r OCIOperationResponse) time.Duration {
	policyToUse, _, _ := crp.DeterminePolicyToUseV2()
	return policyToUse.NextDurationV2(r)
}

// GetMaximumCumulativeBackoffWithoutJitter returns the maximum cumulative backoff the retry policy would do,
// taking into account whether eventually consistency is considered or not.
// This function uses either GetMaximumCumulativeBackoffWithoutJitter or GetMaximumCumulativeEventuallyConsistentBackoffWithoutJitter,
// whichever is appropriate
func (crp *ComplexRetryPolicyV2) GetMaximumCumulativeBackoffWithoutJitter() time.Duration {
	policyToUse, _, _ := crp.DeterminePolicyToUseV2()
	return policyToUse.GetMaximumCumulativeBackoffWithoutJitter()
}

// GetBackoffWithoutJitterV2 calculates the backoff without jitter for the attempt, given the retry policy.
func GetBackoffWithoutJitterV2(policy RetryPolicyV2, attempt uint) time.Duration {
	return time.Duration(getBackoffWithoutJitterHelper(policy.MinSleepBetween, policy.MaxSleepBetween, policy.ExponentialBackoffBase, attempt)) * time.Second
}

// GetMaximumCumulativeBackoffWithoutJitter calculates the maximum backoff without jitter, according to the retry
// policy, if every retry attempt is made.
func (rp *RetryPolicyV2) GetMaximumCumulativeBackoffWithoutJitter() time.Duration {
	return rp.GetMaximumCumulativeBackoffWithoutJitterFunc(rp.MinSleepBetween, rp.MaxSleepBetween, rp.ExponentialBackoffBase, rp.MaximumNumberAttempts, rp.MaximumCumulativeBackoffWithoutJitter)
}

// SetRetry sets the ShouldRetryOperation for RetryPolicy
func (rp *RetryPolicyV2) SetRetry(retryFunc func(r OCIOperationResponse) bool) {
	rp.ShouldRetryOperationFunc = retryFunc
}

// SetRetry sets the ShouldRetryOperation for ComplexRetryPolicy
func (crp *ComplexRetryPolicyV2) SetRetry(retryFunc func(r OCIOperationResponse) bool) {
}

func returnSamePolicyV2(policy OCIRetry) (OCIRetry, *time.Time, float64) {
	// we're returning the end of window time nonetheless, even though the default non-eventual consistency (EC)
	// retry policy doesn't use it; this is useful in case developers wants to write an EC-aware retry policy
	// on their own
	eowt := EcContext.GetEndOfWindow()
	return policy, eowt, 1.0
}

// NoRetryPolicyV2 is a helper method that assembles and returns a return policy that indicates an operation should
// never be retried (the operation is performed exactly once).
func NoRetryPolicyV2() OCIRetry {
	retryPolicy := DefaultRetryPolicyV2()
	retryPolicy.MaxSleepBetween = 0.0
	retryPolicy.MaximumNumberAttempts = 1
	retryPolicy.Enable = false
	retryPolicy.ExponentialBackoffBase = 0.0
	retryPolicy.ShouldRetryOperationFunc = func(OCIOperationResponse) bool { return false }
	retryPolicy.NextDurationFunc = func(OCIOperationResponse, float64, float64, float64, uint) time.Duration { return 0 * time.Second }
	return retryPolicy
}

// DefaultComplexRetryPolicyV2 is a helper method that assembles and returns a return policy that is defined to be a default one
// The default retry policy will retry on (409, IncorrectState), (429, TooManyRequests) and any 5XX errors except (501, MethodNotImplemented)
// The default retry behavior is using exponential backoff with jitter, the maximum wait time is 30s plus 1s jitter
// The maximum cumulative backoff after all 8 attempts have been made is about 1.5 minutes.
// It will also retry on errors affected by eventual consistency.
// The eventual consistency retry behavior is using exponential backoff with jitter, the maximum wait time is 45s plus 1s jitter
// Under eventual consistency, the maximum cumulative backoff after all 9 attempts have been made is about 4 minutes.
func DefaultComplexRetryPolicyV2() *ComplexRetryPolicyV2 {
	return &ComplexRetryPolicyV2{
		RetryPolicy:   DefaultRetryPolicyV2(),
		EcRetryPolicy: DefaultECRetryPolicyV2(),
	}
}

// DefaultRetryPolicyV2 is a helper method that assembles and returns a return policy that is defined to be a default one
// The default retry policy will retry on (409, IncorrectState), (429, TooManyRequests) and any 5XX errors except (501, MethodNotImplemented)
// The default retry behavior is using exponential backoff with jitter, the maximum wait time is 30s plus 1s jitter
// The maximum cumulative backoff after all 8 attempts have been made is about 1.5 minutes.
// It will also retry on errors affected by eventual consistency.
// The eventual consistency retry behavior is using exponential backoff with jitter, the maximum wait time is 45s plus 1s jitter
// Under eventual consistency, the maximum cumulative backoff after all 9 attempts have been made is about 4 minutes.
func DefaultRetryPolicyV2() *RetryPolicyV2 {
	return &RetryPolicyV2{
		MaximumNumberAttempts:  defaultMaximumNumberAttempts,
		MinSleepBetween:        defaultMinSleepBetween,
		MaxSleepBetween:        defaultMaxSleepBetween,
		ExponentialBackoffBase: defaultExponentialBackoffBase,
		Enable:                 true,
		ShouldRetryOperationFunc: func(r OCIOperationResponse) bool {
			if r.Error == nil && 199 < r.Response.HTTPResponse().StatusCode && r.Response.HTTPResponse().StatusCode < 300 {
				// success
				return false
			}
			return IsErrorRetryableByDefault(r.Error)
		},
		NextDurationFunc: func(r OCIOperationResponse, minSleepBetween float64, maxSleepBetween float64, exponentialBackoffBase float64, attempt uint) time.Duration {
			sleepTime := getBackoffWithoutJitterHelper(minSleepBetween, maxSleepBetween, exponentialBackoffBase, r.AttemptNumber)
			nextDuration := time.Duration(1000.0*(sleepTime+rand.Float64())) * time.Millisecond
			return nextDuration
		},
		ShouldContinueRetryFunc: func(current uint, attempt uint) bool {
			return attempt == UnlimitedNumAttemptsValue || current <= attempt
		},
		GetMaximumCumulativeBackoffWithoutJitterFunc: func(minSleepBetween float64, maxSleepBetween float64, exponentialBackoffBase float64, attempt uint, maxCum float64) time.Duration {
			return getMaximumCumulativeBackoffWithoutJitterHelper(minSleepBetween, maxSleepBetween, exponentialBackoffBase, attempt, maxCum)
		},
	}
}

// DefaultECRetryPolicyV2 is a helper method that assembles and returns a return policy that is defined to be a default one
// The default retry policy will retry on (409, IncorrectState), (429, TooManyRequests) and any 5XX errors except (501, MethodNotImplemented)
// The default retry behavior is using exponential backoff with jitter, the maximum wait time is 30s plus 1s jitter
// The maximum cumulative backoff after all 8 attempts have been made is about 1.5 minutes.
// It will also retry on errors affected by eventual consistency.
// The eventual consistency retry behavior is using exponential backoff with jitter, the maximum wait time is 45s plus 1s jitter
// Under eventual consistency, the maximum cumulative backoff after all 9 attempts have been made is about 4 minutes.
func DefaultECRetryPolicyV2() *RetryPolicyV2 {
	return &RetryPolicyV2{
		MaximumNumberAttempts:  ecMaximumNumberAttempts,
		MinSleepBetween:        ecMinSleepBetween,
		MaxSleepBetween:        ecMaxSleepBetween,
		ExponentialBackoffBase: ecExponentialBackoffBase,
		Enable:                 true,
		ShouldRetryOperationFunc: func(r OCIOperationResponse) bool {
			if r.Error == nil && 199 < r.Response.HTTPResponse().StatusCode && r.Response.HTTPResponse().StatusCode < 300 {
				// success
				Debugln(fmt.Sprintf("EC.ShouldRetryOperation, status = %v, 2xx, returning false", r.Response.HTTPResponse().StatusCode))
				return false
			}
			if IsErrorRetryableByDefault(r.Error) {
				return true
			}
			// not retryable by default
			if _, ok := IsServiceError(r.Error); ok {
				now := EcContext.timeNowProvider()
				if r.EndOfWindowTime == nil || r.EndOfWindowTime.Before(now) {
					// either no eventually consistent effects, or they have disappeared by now
					Debugln(fmt.Sprintf("EC.ShouldRetryOperation, no EC or in the past, returning false: endOfWindowTime = %v, now = %v", r.EndOfWindowTime, now))
					return false
				}
				// there were eventually consistent effects present at the time of the first request
				// and they could still affect the retries
				if IsErrorAffectedByEventualConsistency(r.Error) {
					// and it's one of the three affected error codes
					Debugln(fmt.Sprintf("EC.ShouldRetryOperation, affected by EC, EC is present: endOfWindowTime = %v, now = %v", r.EndOfWindowTime, now))
					return true
				}
				return false
			}
			return false
		},
		NextDurationFunc: func(r OCIOperationResponse, minSleepBetween float64, maxSleepBetween float64, exponentialBackoffBase float64, attempt uint) time.Duration {
			sleepTime := getEventuallyConsistentBackoffWithoutJitterHelper(ecMinSleepBetween, ecMaxSleepBetween, ecExponentialBackoffBase, r.AttemptNumber, r.BackoffScalingFactor,
				func(minSleepBetween float64, maxSleepBetween float64, exponentialBackoffBase float64, attempt uint) float64 {
					return getBackoffWithoutJitterHelper(minSleepBetween, maxSleepBetween, exponentialBackoffBase, attempt)
				})
			nextDuration := time.Duration(1000.0*(sleepTime+rand.Float64())) * time.Millisecond
			Debugln(fmt.Sprintf("EventuallyConsistentRetryPolicy.NextDuration for attempt %v: sleepTime = %.1fs, nextDuration = %v", r.AttemptNumber, sleepTime, nextDuration))
			return nextDuration
		},
		ShouldContinueRetryFunc: func(current uint, attempt uint) bool {
			return attempt == UnlimitedNumAttemptsValue || current <= attempt
		},
		GetMaximumCumulativeBackoffWithoutJitterFunc: func(minSleepBetween float64, maxSleepBetween float64, exponentialBackoffBase float64, attempt uint, maxCum float64) time.Duration {
			return getMaximumCumulativeEventuallyConsistentBackoffWithoutJitterHelper(minSleepBetween, maxSleepBetween, exponentialBackoffBase,
				attempt, maxCum,
				func(minSleepBetween float64, maxSleepBetween float64, exponentialBackoffBase float64, attempt uint) float64 {
					return getBackoffWithoutJitterHelper(minSleepBetween, maxSleepBetween, exponentialBackoffBase, attempt)
				})
		},
	}
}

// DeterminePolicyToUseV2 may modify the policy to handle eventual consistency; the return values are
// the retry policy to use, the end of the eventually consistent time window, and the backoff scaling factor
// If eventual consistency is not considered, this function should return the unmodified policy that was
// provided as input, along with (*time.Time)(nil) (no time window), and 1.0 (unscaled backoff).
func (crp *ComplexRetryPolicyV2) DeterminePolicyToUseV2() (OCIRetry, *time.Time, float64) {
	initialAttemptTime := EcContext.timeNowProvider()
	var endOfWindowTime = (*time.Time)(nil)
	var backoffScalingFactor = 1.0
	var policyToUse OCIRetry = crp.RetryPolicy

	eowt := EcContext.GetEndOfWindow()
	if eowt != nil {
		// there was an eventually consistent request
		if eowt.After(initialAttemptTime) {
			// and the eventually consistent effects may still be present
			endOfWindowTime = eowt
			// if the time between now and the end of the window is less than the time we normally would retry, use the default timing
			durationToEndOfWindow := endOfWindowTime.Sub(initialAttemptTime)
			maxCumulativeBackoffWithoutJitter := crp.RetryPolicy.GetMaximumCumulativeBackoffWithoutJitter()
			Debugln(fmt.Sprintf("durationToEndOfWindow = %v, maxCumulativeBackoffWithoutJitter = %v", durationToEndOfWindow, maxCumulativeBackoffWithoutJitter))
			if durationToEndOfWindow > maxCumulativeBackoffWithoutJitter {
				// the end of the eventually consistent window is later than when default retries would end
				// do not use default timing
				policyToUse = crp.EcRetryPolicy
				maximumCumulativeBackoffWithoutJitter := crp.EcRetryPolicy.GetMaximumCumulativeBackoffWithoutJitter()
				backoffScalingFactor = float64(durationToEndOfWindow) / float64(maximumCumulativeBackoffWithoutJitter)
				Debugln(fmt.Sprintf("Use eventually consistent timing, durationToEndOfWindow = %v, maximumCumulativeBackoffWithoutJitter = %v, backoffScalingFactor = %.2f",
					durationToEndOfWindow, maximumCumulativeBackoffWithoutJitter, backoffScalingFactor))
			} else {
				policyToUse.SetRetry(EventuallyConsistentShouldRetryOperation)
				Debugln(fmt.Sprintf("Use default timing, end of EC window is sooner than default retries"))
			}
		} else {
			Debugln(fmt.Sprintf("Use default timing and strategy, end of EC window is in the past"))
		}
	} else {
		Debugln(fmt.Sprintf("Use default timing and strategy, no EC window set"))
	}
	return policyToUse, endOfWindowTime, backoffScalingFactor
}
