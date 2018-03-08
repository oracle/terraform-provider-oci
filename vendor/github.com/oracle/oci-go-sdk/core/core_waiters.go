// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"context"
	"fmt"
	oci_common "github.com/oracle/oci-go-sdk/common"
	"time"
)

// PollBootVolumeUntil polls a resource until the specified predicate returns true
func (client BlockstorageClient) PollBootVolumeUntil(ctx context.Context, request GetBootVolumeRequest, predicate func(GetBootVolumeResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetBootVolume(deadlineContext, request)

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

// PollBootVolumeAttachmentUntil polls a resource until the specified predicate returns true
func (client ComputeClient) PollBootVolumeAttachmentUntil(ctx context.Context, request GetBootVolumeAttachmentRequest, predicate func(GetBootVolumeAttachmentResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetBootVolumeAttachment(deadlineContext, request)

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

// PollConsoleHistoryUntil polls a resource until the specified predicate returns true
func (client ComputeClient) PollConsoleHistoryUntil(ctx context.Context, request GetConsoleHistoryRequest, predicate func(GetConsoleHistoryResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetConsoleHistory(deadlineContext, request)

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

// PollCrossConnectUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollCrossConnectUntil(ctx context.Context, request GetCrossConnectRequest, predicate func(GetCrossConnectResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetCrossConnect(deadlineContext, request)

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

// PollCrossConnectGroupUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollCrossConnectGroupUntil(ctx context.Context, request GetCrossConnectGroupRequest, predicate func(GetCrossConnectGroupResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetCrossConnectGroup(deadlineContext, request)

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

// PollDhcpOptionsUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollDhcpOptionsUntil(ctx context.Context, request GetDhcpOptionsRequest, predicate func(GetDhcpOptionsResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetDhcpOptions(deadlineContext, request)

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

// PollDrgUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollDrgUntil(ctx context.Context, request GetDrgRequest, predicate func(GetDrgResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetDrg(deadlineContext, request)

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

// PollDrgAttachmentUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollDrgAttachmentUntil(ctx context.Context, request GetDrgAttachmentRequest, predicate func(GetDrgAttachmentResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetDrgAttachment(deadlineContext, request)

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

// PollImageUntil polls a resource until the specified predicate returns true
func (client ComputeClient) PollImageUntil(ctx context.Context, request GetImageRequest, predicate func(GetImageResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetImage(deadlineContext, request)

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

// PollInstanceUntil polls a resource until the specified predicate returns true
func (client ComputeClient) PollInstanceUntil(ctx context.Context, request GetInstanceRequest, predicate func(GetInstanceResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetInstance(deadlineContext, request)

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

// PollInstanceConsoleConnectionUntil polls a resource until the specified predicate returns true
func (client ComputeClient) PollInstanceConsoleConnectionUntil(ctx context.Context, request GetInstanceConsoleConnectionRequest, predicate func(GetInstanceConsoleConnectionResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetInstanceConsoleConnection(deadlineContext, request)

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

// PollInternetGatewayUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollInternetGatewayUntil(ctx context.Context, request GetInternetGatewayRequest, predicate func(GetInternetGatewayResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetInternetGateway(deadlineContext, request)

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

// PollIPSecConnectionUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollIPSecConnectionUntil(ctx context.Context, request GetIPSecConnectionRequest, predicate func(GetIPSecConnectionResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetIPSecConnection(deadlineContext, request)

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

// PollLocalPeeringGatewayUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollLocalPeeringGatewayUntil(ctx context.Context, request GetLocalPeeringGatewayRequest, predicate func(GetLocalPeeringGatewayResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetLocalPeeringGateway(deadlineContext, request)

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

// PollRouteTableUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollRouteTableUntil(ctx context.Context, request GetRouteTableRequest, predicate func(GetRouteTableResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetRouteTable(deadlineContext, request)

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

// PollSecurityListUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollSecurityListUntil(ctx context.Context, request GetSecurityListRequest, predicate func(GetSecurityListResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetSecurityList(deadlineContext, request)

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

// PollSubnetUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollSubnetUntil(ctx context.Context, request GetSubnetRequest, predicate func(GetSubnetResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetSubnet(deadlineContext, request)

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

// PollVcnUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollVcnUntil(ctx context.Context, request GetVcnRequest, predicate func(GetVcnResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetVcn(deadlineContext, request)

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

// PollVirtualCircuitUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollVirtualCircuitUntil(ctx context.Context, request GetVirtualCircuitRequest, predicate func(GetVirtualCircuitResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetVirtualCircuit(deadlineContext, request)

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

// PollVnicUntil polls a resource until the specified predicate returns true
func (client VirtualNetworkClient) PollVnicUntil(ctx context.Context, request GetVnicRequest, predicate func(GetVnicResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetVnic(deadlineContext, request)

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

// PollVnicAttachmentUntil polls a resource until the specified predicate returns true
func (client ComputeClient) PollVnicAttachmentUntil(ctx context.Context, request GetVnicAttachmentRequest, predicate func(GetVnicAttachmentResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetVnicAttachment(deadlineContext, request)

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

// PollVolumeUntil polls a resource until the specified predicate returns true
func (client BlockstorageClient) PollVolumeUntil(ctx context.Context, request GetVolumeRequest, predicate func(GetVolumeResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetVolume(deadlineContext, request)

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

// PollVolumeAttachmentUntil polls a resource until the specified predicate returns true
func (client ComputeClient) PollVolumeAttachmentUntil(ctx context.Context, request GetVolumeAttachmentRequest, predicate func(GetVolumeAttachmentResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetVolumeAttachment(deadlineContext, request)

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

// PollVolumeBackupUntil polls a resource until the specified predicate returns true
func (client BlockstorageClient) PollVolumeBackupUntil(ctx context.Context, request GetVolumeBackupRequest, predicate func(GetVolumeBackupResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetVolumeBackup(deadlineContext, request)

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
