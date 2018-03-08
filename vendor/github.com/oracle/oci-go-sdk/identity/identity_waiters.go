// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"context"
	"fmt"
	oci_common "github.com/oracle/oci-go-sdk/common"
	"time"
)

// PollApiKeyUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollApiKeyUntil(ctx context.Context, request ListApiKeysRequest, predicate func(ListApiKeysResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListApiKeys(deadlineContext, request)

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

// PollCompartmentUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollCompartmentUntil(ctx context.Context, request GetCompartmentRequest, predicate func(GetCompartmentResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetCompartment(deadlineContext, request)

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

// PollCustomerSecretKeySummaryUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollCustomerSecretKeySummaryUntil(ctx context.Context, request ListCustomerSecretKeysRequest, predicate func(ListCustomerSecretKeysResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListCustomerSecretKeys(deadlineContext, request)

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

// PollGroupUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollGroupUntil(ctx context.Context, request GetGroupRequest, predicate func(GetGroupResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetGroup(deadlineContext, request)

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

// PollIdentityProviderUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollIdentityProviderUntil(ctx context.Context, request GetIdentityProviderRequest, predicate func(GetIdentityProviderResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetIdentityProvider(deadlineContext, request)

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

// PollIdpGroupMappingUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollIdpGroupMappingUntil(ctx context.Context, request GetIdpGroupMappingRequest, predicate func(GetIdpGroupMappingResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetIdpGroupMapping(deadlineContext, request)

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

// PollPolicyUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollPolicyUntil(ctx context.Context, request GetPolicyRequest, predicate func(GetPolicyResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetPolicy(deadlineContext, request)

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

// PollRegionSubscriptionUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollRegionSubscriptionUntil(ctx context.Context, request ListRegionSubscriptionsRequest, predicate func(ListRegionSubscriptionsResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListRegionSubscriptions(deadlineContext, request)

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

// PollSwiftPasswordUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollSwiftPasswordUntil(ctx context.Context, request ListSwiftPasswordsRequest, predicate func(ListSwiftPasswordsResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListSwiftPasswords(deadlineContext, request)

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

// PollUserUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollUserUntil(ctx context.Context, request GetUserRequest, predicate func(GetUserResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetUser(deadlineContext, request)

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

// PollUserGroupMembershipUntil polls a resource until the specified predicate returns true
func (client IdentityClient) PollUserGroupMembershipUntil(ctx context.Context, request GetUserGroupMembershipRequest, predicate func(GetUserGroupMembershipResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetUserGroupMembership(deadlineContext, request)

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
