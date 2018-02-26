// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"context"
	"fmt"
	oci_common "github.com/oracle/oci-go-sdk/common"
	"time"
)

// PollBackendHealthUntil polls a resource until the specified predicate returns true
func (client LoadBalancerClient) PollBackendHealthUntil(ctx context.Context, request GetBackendHealthRequest, predicate func(GetBackendHealthResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetBackendHealth(deadlineContext, request)

		select {
		case <-deadlineContext.Done():
			return ctx.Err()
		default:
			// non-blocking select
		}

		if predicate(response, err) {
			return nil
		}

		if policy.ShouldRetryOperation(response.RawResponse, err, currentOperationAttempt) {
			if currentOperationAttempt != policy.MaximumNumberAttempts {
				// sleep before retrying the operation
				duration := policy.GetNextDuration(currentOperationAttempt)
				if deadline, ok := ctx.Deadline(); ok && time.Now().Add(duration).After(deadline) {
					return oci_common.DurationExceedsDeadline
				}
				time.Sleep(duration)
			}
		} else {
			// we should NOT retry operation based on response and/or error => return
			return err
		}
	}
	return fmt.Errorf("maximum number of attempts exceeded (%v)", policy.MaximumNumberAttempts)
}

// PollBackendSetHealthUntil polls a resource until the specified predicate returns true
func (client LoadBalancerClient) PollBackendSetHealthUntil(ctx context.Context, request GetBackendSetHealthRequest, predicate func(GetBackendSetHealthResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetBackendSetHealth(deadlineContext, request)

		select {
		case <-deadlineContext.Done():
			return ctx.Err()
		default:
			// non-blocking select
		}

		if predicate(response, err) {
			return nil
		}

		if policy.ShouldRetryOperation(response.RawResponse, err, currentOperationAttempt) {
			if currentOperationAttempt != policy.MaximumNumberAttempts {
				// sleep before retrying the operation
				duration := policy.GetNextDuration(currentOperationAttempt)
				if deadline, ok := ctx.Deadline(); ok && time.Now().Add(duration).After(deadline) {
					return oci_common.DurationExceedsDeadline
				}
				time.Sleep(duration)
			}
		} else {
			// we should NOT retry operation based on response and/or error => return
			return err
		}
	}
	return fmt.Errorf("maximum number of attempts exceeded (%v)", policy.MaximumNumberAttempts)
}

// PollLoadBalancerUntil polls a resource until the specified predicate returns true
func (client LoadBalancerClient) PollLoadBalancerUntil(ctx context.Context, request GetLoadBalancerRequest, predicate func(GetLoadBalancerResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetLoadBalancer(deadlineContext, request)

		select {
		case <-deadlineContext.Done():
			return ctx.Err()
		default:
			// non-blocking select
		}

		if predicate(response, err) {
			return nil
		}

		if policy.ShouldRetryOperation(response.RawResponse, err, currentOperationAttempt) {
			if currentOperationAttempt != policy.MaximumNumberAttempts {
				// sleep before retrying the operation
				duration := policy.GetNextDuration(currentOperationAttempt)
				if deadline, ok := ctx.Deadline(); ok && time.Now().Add(duration).After(deadline) {
					return oci_common.DurationExceedsDeadline
				}
				time.Sleep(duration)
			}
		} else {
			// we should NOT retry operation based on response and/or error => return
			return err
		}
	}
	return fmt.Errorf("maximum number of attempts exceeded (%v)", policy.MaximumNumberAttempts)
}

// PollLoadBalancerHealthUntil polls a resource until the specified predicate returns true
func (client LoadBalancerClient) PollLoadBalancerHealthUntil(ctx context.Context, request GetLoadBalancerHealthRequest, predicate func(GetLoadBalancerHealthResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetLoadBalancerHealth(deadlineContext, request)

		select {
		case <-deadlineContext.Done():
			return ctx.Err()
		default:
			// non-blocking select
		}

		if predicate(response, err) {
			return nil
		}

		if policy.ShouldRetryOperation(response.RawResponse, err, currentOperationAttempt) {
			if currentOperationAttempt != policy.MaximumNumberAttempts {
				// sleep before retrying the operation
				duration := policy.GetNextDuration(currentOperationAttempt)
				if deadline, ok := ctx.Deadline(); ok && time.Now().Add(duration).After(deadline) {
					return oci_common.DurationExceedsDeadline
				}
				time.Sleep(duration)
			}
		} else {
			// we should NOT retry operation based on response and/or error => return
			return err
		}
	}
	return fmt.Errorf("maximum number of attempts exceeded (%v)", policy.MaximumNumberAttempts)
}

// PollLoadBalancerHealthSummaryUntil polls a resource until the specified predicate returns true
func (client LoadBalancerClient) PollLoadBalancerHealthSummaryUntil(ctx context.Context, request ListLoadBalancerHealthsRequest, predicate func(ListLoadBalancerHealthsResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListLoadBalancerHealths(deadlineContext, request)

		select {
		case <-deadlineContext.Done():
			return ctx.Err()
		default:
			// non-blocking select
		}

		if predicate(response, err) {
			return nil
		}

		if policy.ShouldRetryOperation(response.RawResponse, err, currentOperationAttempt) {
			if currentOperationAttempt != policy.MaximumNumberAttempts {
				// sleep before retrying the operation
				duration := policy.GetNextDuration(currentOperationAttempt)
				if deadline, ok := ctx.Deadline(); ok && time.Now().Add(duration).After(deadline) {
					return oci_common.DurationExceedsDeadline
				}
				time.Sleep(duration)
			}
		} else {
			// we should NOT retry operation based on response and/or error => return
			return err
		}
	}
	return fmt.Errorf("maximum number of attempts exceeded (%v)", policy.MaximumNumberAttempts)
}

// PollWorkRequestUntil polls a resource until the specified predicate returns true
func (client LoadBalancerClient) PollWorkRequestUntil(ctx context.Context, request GetWorkRequestRequest, predicate func(GetWorkRequestResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetWorkRequest(deadlineContext, request)

		select {
		case <-deadlineContext.Done():
			return ctx.Err()
		default:
			// non-blocking select
		}

		if predicate(response, err) {
			return nil
		}

		if policy.ShouldRetryOperation(response.RawResponse, err, currentOperationAttempt) {
			if currentOperationAttempt != policy.MaximumNumberAttempts {
				// sleep before retrying the operation
				duration := policy.GetNextDuration(currentOperationAttempt)
				if deadline, ok := ctx.Deadline(); ok && time.Now().Add(duration).After(deadline) {
					return oci_common.DurationExceedsDeadline
				}
				time.Sleep(duration)
			}
		} else {
			// we should NOT retry operation based on response and/or error => return
			return err
		}
	}
	return fmt.Errorf("maximum number of attempts exceeded (%v)", policy.MaximumNumberAttempts)
}
