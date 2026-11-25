// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
package distributed_database

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdk "github.com/oracle/oci-go-sdk/v65/distributeddatabase"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

type DistributedDatabaseDistributedAutonomousDatabasesPatchResourceCrud struct {
	tfresource.BaseCrud
	Client            *sdk.DistributedAutonomousDbServiceClient
	WorkRequestClient *sdk.DistributedDbWorkRequestServiceClient
	Kind              databaseActionKind
	Res               *sdk.DistributedAutonomousDatabase
}

func (s *DistributedDatabaseDistributedAutonomousDatabasesPatchResourceCrud) CreateWithContext(ctx context.Context) error {
	return s.executeAction(ctx, schema.TimeoutCreate)
}
func (s *DistributedDatabaseDistributedAutonomousDatabasesPatchResourceCrud) UpdateWithContext(ctx context.Context) error {
	return s.executeAction(ctx, schema.TimeoutUpdate)
}
func (s *DistributedDatabaseDistributedAutonomousDatabasesPatchResourceCrud) executeAction(ctx context.Context, timeoutKey string) error {
	if s.D.Get("action_pending").(bool) {
		expectedID := distributedDatabaseActionID("DistributedDatabase"+actionOperationName(s.Kind)+"Action-", s.D, actionTargetField(true), string(s.Kind))
		if s.D.Id() != expectedID {
			return fmt.Errorf("cannot change operations while work request %q is pending; resume the submitted configuration first", s.D.Get("work_request_id"))
		}
		return s.resumeAction(ctx, timeoutKey)
	}
	id := s.D.Get("distributed_autonomous_database_id").(string)
	instructions, err := operationActionInstructions(s.D, s.Kind, true)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(ctx, s.D.Timeout(timeoutKey))
	defer cancel()
	if err := waitForDistributedAutonomousDatabaseActionReady(ctx, s.D, s.Client, id); err != nil {
		return err
	}
	if s.Kind == databaseActionInsert {
		if err := s.GetWithContext(ctx); err != nil {
			return err
		}
		if err := validateInsertActionTargets(s.D, s.Res); err != nil {
			return err
		}
	}
	token := sdkRetryToken()
	if err := s.D.Set("retry_token", token); err != nil {
		return err
	}
	request := sdk.PatchDistributedAutonomousDatabaseRequest{DistributedAutonomousDatabaseId: &id, OpcRetryToken: &token, PatchDistributedAutonomousDatabaseDetails: sdk.PatchDistributedAutonomousDatabaseDetails{Items: instructions}}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "distributed_database")
	response, err := s.Client.PatchDistributedAutonomousDatabase(ctx, request)
	if err != nil {
		return err
	}
	workID := response.OpcWorkRequestId
	if workID == nil || strings.TrimSpace(*workID) == "" {
		return fmt.Errorf("missing opc-work-request-id")
	}
	if err := recordDistributedDatabaseActionSubmission(s.D, s.Kind, true, *workID); err != nil {
		return err
	}
	return s.resumeAction(ctx, timeoutKey)
}

func (s *DistributedDatabaseDistributedAutonomousDatabasesPatchResourceCrud) resumeAction(ctx context.Context, timeoutKey string) error {
	id := s.D.Get("distributed_autonomous_database_id").(string)
	work := strings.TrimSpace(s.D.Get("work_request_id").(string))
	if work == "" {
		return fmt.Errorf("pending action has no work_request_id; refusing to resubmit")
	}
	ctx, cancel := context.WithTimeout(ctx, s.D.Timeout(timeoutKey))
	defer cancel()
	workID := &work
	// The target is known; action completion does not require resource discovery.
	if _, err := distributedAutonomousDatabaseWaitForWorkRequest(ctx, workID, "", sdk.ActionTypeUpdated, s.D.Timeout(timeoutKey), false, s.WorkRequestClient); err != nil {
		return err
	}
	if err := waitForDistributedAutonomousDatabaseActionReady(ctx, s.D, s.Client, id); err != nil {
		return err
	}
	if err := s.GetWithContext(ctx); err != nil {
		return err
	}
	return s.D.Set("action_pending", false)
}
func (s *DistributedDatabaseDistributedAutonomousDatabasesPatchResourceCrud) GetWithContext(ctx context.Context) error {
	id := s.D.Get("distributed_autonomous_database_id").(string)
	response, err := s.Client.GetDistributedAutonomousDatabase(ctx, sdk.GetDistributedAutonomousDatabaseRequest{DistributedAutonomousDatabaseId: &id})
	if err != nil {
		return err
	}
	s.Res = &response.DistributedAutonomousDatabase
	if err := s.D.Set("state", string(response.DistributedAutonomousDatabase.LifecycleState)); err != nil {
		return err
	}
	details := ""
	if response.DistributedAutonomousDatabase.LifecycleDetails != nil {
		details = *response.DistributedAutonomousDatabase.LifecycleDetails
	}
	return s.D.Set("lifecycle_details", details)
}
