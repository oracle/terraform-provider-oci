// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package batch

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BatchBatchJobPoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createBatchBatchJobPoolWithContext,
		ReadContext:   readBatchBatchJobPoolWithContext,
		UpdateContext: updateBatchBatchJobPoolWithContext,
		DeleteContext: deleteBatchBatchJobPoolWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"batch_context_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_batch.BatchJobPoolLifecycleStateInactive),
					string(oci_batch.BatchJobPoolLifecycleStateActive),
				}, true),
			},

			// Computed
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBatchBatchJobPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchJobPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_batch.BatchJobPoolLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_batch.BatchJobPoolLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	if powerOff {
		if err := sync.StopBatchJobPool(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_batch.BatchJobPoolLifecycleStateInactive)
	}
	return nil

}

func readBatchBatchJobPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchJobPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateBatchBatchJobPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchJobPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_batch.BatchJobPoolLifecycleStateActive == oci_batch.BatchJobPoolLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_batch.BatchJobPoolLifecycleStateInactive == oci_batch.BatchJobPoolLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartBatchJobPool(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_batch.BatchJobPoolLifecycleStateActive)
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	if powerOff {
		if err := sync.StopBatchJobPool(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_batch.BatchJobPoolLifecycleStateInactive)
	}

	return nil
}

func deleteBatchBatchJobPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchJobPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type BatchBatchJobPoolResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_batch.BatchComputingClient
	Res                    *oci_batch.BatchJobPool
	DisableNotFoundRetries bool
}

func (s *BatchBatchJobPoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BatchBatchJobPoolResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *BatchBatchJobPoolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_batch.BatchJobPoolLifecycleStateActive),
		string(oci_batch.BatchJobPoolLifecycleStateNeedsAttention),
	}
}

func (s *BatchBatchJobPoolResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *BatchBatchJobPoolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_batch.BatchJobPoolLifecycleStateDeleted),
	}
}

func (s *BatchBatchJobPoolResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_batch.CreateBatchJobPoolRequest{}

	if batchContextId, ok := s.D.GetOkExists("batch_context_id"); ok {
		tmp := batchContextId.(string)
		request.BatchContextId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.CreateBatchJobPool(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BatchJobPool
	return nil
}

func (s *BatchBatchJobPoolResourceCrud) getBatchJobPoolFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_batch.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	batchJobPoolId, err := batchJobPoolWaitForWorkRequest(ctx, workId, "batchjobpool",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*batchJobPoolId)

	return s.GetWithContext(ctx)
}

func batchJobPoolWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "batch", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_batch.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func batchJobPoolWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_batch.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_batch.BatchComputingClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "batch")
	retryPolicy.ShouldRetryOperation = batchJobPoolWorkRequestShouldRetryFunc(timeout)

	response := oci_batch.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_batch.OperationStatusInProgress),
			string(oci_batch.OperationStatusAccepted),
			string(oci_batch.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_batch.OperationStatusSucceeded),
			string(oci_batch.OperationStatusFailed),
			string(oci_batch.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_batch.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_batch.OperationStatusFailed || response.Status == oci_batch.OperationStatusCanceled {
		return nil, getErrorFromBatchBatchJobPoolWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBatchBatchJobPoolWorkRequest(ctx context.Context, client *oci_batch.BatchComputingClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_batch.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_batch.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *BatchBatchJobPoolResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.GetBatchJobPoolRequest{}

	tmp := s.D.Id()
	request.BatchJobPoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.GetBatchJobPool(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BatchJobPool
	return nil
}

func (s *BatchBatchJobPoolResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_batch.UpdateBatchJobPoolRequest{}

	tmp := s.D.Id()
	request.BatchJobPoolId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.UpdateBatchJobPool(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBatchJobPoolFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch"), oci_batch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BatchBatchJobPoolResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_batch.DeleteBatchJobPoolRequest{}

	tmp := s.D.Id()
	request.BatchJobPoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	_, err := s.Client.DeleteBatchJobPool(ctx, request)
	return err
}

func (s *BatchBatchJobPoolResourceCrud) SetData() error {
	if s.Res.BatchContextId != nil {
		s.D.Set("batch_context_id", *s.Res.BatchContextId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *BatchBatchJobPoolResourceCrud) StartBatchJobPool(ctx context.Context) error {
	request := oci_batch.StartBatchJobPoolRequest{}

	idTmp := s.D.Id()
	request.BatchJobPoolId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	_, err := s.Client.StartBatchJobPool(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_batch.BatchJobPoolLifecycleStateActive }
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BatchBatchJobPoolResourceCrud) StopBatchJobPool(ctx context.Context) error {
	request := oci_batch.StopBatchJobPoolRequest{}

	idTmp := s.D.Id()
	request.BatchJobPoolId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	_, err := s.Client.StopBatchJobPool(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_batch.BatchJobPoolLifecycleStateInactive }
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func BatchJobPoolSummaryToMap(obj oci_batch.BatchJobPoolSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BatchContextId != nil {
		result["batch_context_id"] = string(*obj.BatchContextId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *BatchBatchJobPoolResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_batch.ChangeBatchJobPoolCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BatchJobPoolId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.ChangeBatchJobPoolCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBatchJobPoolFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch"), oci_batch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
