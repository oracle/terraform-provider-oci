// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"context"
	"fmt"
	oci_common "github.com/oracle/oci-go-sdk/common"
	"time"
)

// PollBackupUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollBackupUntil(ctx context.Context, request GetBackupRequest, predicate func(GetBackupResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetBackup(deadlineContext, request)

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

// PollBackupSummaryUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollBackupSummaryUntil(ctx context.Context, request ListBackupsRequest, predicate func(ListBackupsResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListBackups(deadlineContext, request)

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

// PollDataGuardAssociationUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollDataGuardAssociationUntil(ctx context.Context, request GetDataGuardAssociationRequest, predicate func(GetDataGuardAssociationResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetDataGuardAssociation(deadlineContext, request)

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

// PollDataGuardAssociationSummaryUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollDataGuardAssociationSummaryUntil(ctx context.Context, request ListDataGuardAssociationsRequest, predicate func(ListDataGuardAssociationsResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListDataGuardAssociations(deadlineContext, request)

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

// PollDatabaseUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollDatabaseUntil(ctx context.Context, request GetDatabaseRequest, predicate func(GetDatabaseResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetDatabase(deadlineContext, request)

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

// PollDatabaseSummaryUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollDatabaseSummaryUntil(ctx context.Context, request ListDatabasesRequest, predicate func(ListDatabasesResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListDatabases(deadlineContext, request)

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

// PollDbHomeUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollDbHomeUntil(ctx context.Context, request GetDbHomeRequest, predicate func(GetDbHomeResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetDbHome(deadlineContext, request)

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

// PollDbHomeSummaryUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollDbHomeSummaryUntil(ctx context.Context, request ListDbHomesRequest, predicate func(ListDbHomesResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListDbHomes(deadlineContext, request)

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

// PollDbNodeUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollDbNodeUntil(ctx context.Context, request GetDbNodeRequest, predicate func(GetDbNodeResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetDbNode(deadlineContext, request)

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

// PollDbNodeSummaryUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollDbNodeSummaryUntil(ctx context.Context, request ListDbNodesRequest, predicate func(ListDbNodesResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListDbNodes(deadlineContext, request)

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

// PollDbSystemUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollDbSystemUntil(ctx context.Context, request GetDbSystemRequest, predicate func(GetDbSystemResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetDbSystem(deadlineContext, request)

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

// PollDbSystemSummaryUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollDbSystemSummaryUntil(ctx context.Context, request ListDbSystemsRequest, predicate func(ListDbSystemsResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListDbSystems(deadlineContext, request)

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

// PollPatchUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollPatchUntil(ctx context.Context, request GetDbHomePatchRequest, predicate func(GetDbHomePatchResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetDbHomePatch(deadlineContext, request)

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

// PollPatchHistoryEntryUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollPatchHistoryEntryUntil(ctx context.Context, request GetDbHomePatchHistoryEntryRequest, predicate func(GetDbHomePatchHistoryEntryResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.GetDbHomePatchHistoryEntry(deadlineContext, request)

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

// PollPatchHistoryEntrySummaryUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollPatchHistoryEntrySummaryUntil(ctx context.Context, request ListDbHomePatchHistoryEntriesRequest, predicate func(ListDbHomePatchHistoryEntriesResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListDbHomePatchHistoryEntries(deadlineContext, request)

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

// PollPatchSummaryUntil polls a resource until the specified predicate returns true
func (client DatabaseClient) PollPatchSummaryUntil(ctx context.Context, request ListDbHomePatchesRequest, predicate func(ListDbHomePatchesResponse, error) bool, options ...oci_common.RetryPolicyOption) error {
	policy := oci_common.BuildRetryPolicy(options...)
	deadlineContext, deadlineCancel := context.WithTimeout(ctx, oci_common.GetMaximumTimeout(policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); oci_common.ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		response, err := client.ListDbHomePatches(deadlineContext, request)

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
