// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

const workRequestNetworkRetryWindow = 15 * time.Minute

// A network outage gets its own budget, independent of the time the service
// has already spent provisioning. Successful polls reset that budget. Keep this
// policy local to work-request GETs rather than extending mutation retries.
func distributedDatabaseWorkRequestRetryPolicy(timeout time.Duration, disableNotFoundRetries bool) *oci_common.RetryPolicy {
	policy := tfresource.GetRetryPolicy(disableNotFoundRetries, "distributed_database")
	maxAttempts := policy.MaximumNumberAttempts
	defaultNextDuration := policy.NextDuration
	start := time.Now()
	deadline := start.Add(timeout)
	var outageStart time.Time
	// An SDK attempt cap must not end a recoverable outage before its deadline.
	// Retain the original cap for non-network errors below.
	policy.MaximumNumberAttempts = 0
	policy.ShouldRetryOperation = func(response oci_common.OCIOperationResponse) bool {
		now := time.Now()
		if !now.Before(deadline) {
			return false
		}
		if tfresource.IsNetworkError(response.Error) {
			if outageStart.IsZero() {
				outageStart = now
			}
			return now.Sub(outageStart) < workRequestNetworkRetryWindow
		}
		if response.Error == nil {
			outageStart = time.Time{}
		} else if maxAttempts != 0 && response.AttemptNumber >= maxAttempts {
			return false
		}
		if tfresource.ShouldRetry(response, disableNotFoundRetries, "distributed_database", start) {
			return true
		}
		if response.Error != nil {
			return false
		}
		if work, ok := response.Response.(oci_distributed_database.GetWorkRequestResponse); ok {
			switch work.Status {
			case oci_distributed_database.OperationStatusSucceeded, oci_distributed_database.OperationStatusFailed, oci_distributed_database.OperationStatusCanceled, oci_distributed_database.OperationStatusNeedsAttention:
				return false
			}
			return work.TimeFinished == nil
		}
		return false
	}
	policy.NextDuration = func(response oci_common.OCIOperationResponse) time.Duration {
		delay := defaultNextDuration(response)
		remaining := time.Until(deadline)
		if tfresource.IsNetworkError(response.Error) && !outageStart.IsZero() {
			if outageRemaining := time.Until(outageStart.Add(workRequestNetworkRetryWindow)); outageRemaining < remaining {
				remaining = outageRemaining
			}
		}
		if remaining < 0 {
			return 0
		}
		if delay > remaining {
			return remaining
		}
		return delay
	}
	return policy
}

func distributedDatabaseActionID(prefix string, d *schema.ResourceData, fields ...string) string {
	payload := map[string]interface{}{}
	for _, field := range fields {
		payload[field] = d.Get(field)
	}

	data, err := json.Marshal(payload)
	if err != nil {
		data = []byte(fmt.Sprintf("%v", payload))
	}
	hash := sha1.Sum(data)
	return fmt.Sprintf("%s%x", prefix, hash[:])
}

func waitForDistributedDatabaseActionReady(ctx context.Context, d *schema.ResourceData, client *oci_distributed_database.DistributedDbServiceClient, id string) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_distributed_database.DistributedDatabaseLifecycleStateCreating),
			string(oci_distributed_database.DistributedDatabaseLifecycleStateUpdating),
		},
		Target: []string{"READY"},
		Refresh: func() (interface{}, string, error) {
			response, err := client.GetDistributedDatabase(ctx, oci_distributed_database.GetDistributedDatabaseRequest{
				DistributedDatabaseId: &id,
				RequestMetadata:       oci_common.RequestMetadata{RetryPolicy: tfresource.GetRetryPolicy(false, "distributed_database")},
			})
			if err != nil {
				return nil, "", err
			}
			state := string(response.DistributedDatabase.LifecycleState)
			if state == string(oci_distributed_database.DistributedDatabaseLifecycleStateCreating) || state == string(oci_distributed_database.DistributedDatabaseLifecycleStateUpdating) {
				return &response.DistributedDatabase, state, nil
			}
			if state == "ACTIVE" || state == "INACTIVE" {
				return &response.DistributedDatabase, "READY", nil
			}
			details := ""
			if response.DistributedDatabase.LifecycleDetails != nil {
				details = *response.DistributedDatabase.LifecycleDetails
			}
			return &response.DistributedDatabase, "", fmt.Errorf("database %q ended in lifecycle state %s: %s", id, state, details)
		},
		Timeout: d.Timeout(schema.TimeoutCreate),
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func waitForDistributedAutonomousDatabaseActionReady(ctx context.Context, d *schema.ResourceData, client *oci_distributed_database.DistributedAutonomousDbServiceClient, id string) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateCreating),
			string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateUpdating),
		},
		Target: []string{"READY"},
		Refresh: func() (interface{}, string, error) {
			response, err := client.GetDistributedAutonomousDatabase(ctx, oci_distributed_database.GetDistributedAutonomousDatabaseRequest{
				DistributedAutonomousDatabaseId: &id,
				RequestMetadata:                 oci_common.RequestMetadata{RetryPolicy: tfresource.GetRetryPolicy(false, "distributed_database")},
			})
			if err != nil {
				return nil, "", err
			}
			state := string(response.DistributedAutonomousDatabase.LifecycleState)
			if state == string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateCreating) || state == string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateUpdating) {
				return &response.DistributedAutonomousDatabase, state, nil
			}
			if state == "ACTIVE" || state == "INACTIVE" {
				return &response.DistributedAutonomousDatabase, "READY", nil
			}
			details := ""
			if response.DistributedAutonomousDatabase.LifecycleDetails != nil {
				details = *response.DistributedAutonomousDatabase.LifecycleDetails
			}
			return &response.DistributedAutonomousDatabase, "", fmt.Errorf("database %q ended in lifecycle state %s: %s", id, state, details)
		},
		Timeout: d.Timeout(schema.TimeoutCreate),
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}
