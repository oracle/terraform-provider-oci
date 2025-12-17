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
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BatchBatchTaskProfileResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createBatchBatchTaskProfileWithContext,
		ReadContext:   readBatchBatchTaskProfileWithContext,
		UpdateContext: updateBatchBatchTaskProfileWithContext,
		DeleteContext: deleteBatchBatchTaskProfileWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"min_memory_in_gbs": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"min_ocpus": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
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

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
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

func createBatchBatchTaskProfileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readBatchBatchTaskProfileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateBatchBatchTaskProfileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteBatchBatchTaskProfileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type BatchBatchTaskProfileResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_batch.BatchComputingClient
	Res                    *oci_batch.BatchTaskProfile
	DisableNotFoundRetries bool
}

func (s *BatchBatchTaskProfileResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BatchBatchTaskProfileResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *BatchBatchTaskProfileResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_batch.BatchTaskProfileLifecycleStateActive),
	}
}

func (s *BatchBatchTaskProfileResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *BatchBatchTaskProfileResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_batch.BatchTaskProfileLifecycleStateDeleted),
	}
}

func (s *BatchBatchTaskProfileResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_batch.CreateBatchTaskProfileRequest{}

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

	if minMemoryInGBs, ok := s.D.GetOkExists("min_memory_in_gbs"); ok {
		tmp := minMemoryInGBs.(int)
		request.MinMemoryInGBs = &tmp
	}

	if minOcpus, ok := s.D.GetOkExists("min_ocpus"); ok {
		tmp := minOcpus.(int)
		request.MinOcpus = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.CreateBatchTaskProfile(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BatchTaskProfile
	return nil
}

func (s *BatchBatchTaskProfileResourceCrud) getBatchTaskProfileFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_batch.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	batchTaskProfileId, err := batchTaskProfileWaitForWorkRequest(ctx, workId, "batchtaskprofile",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*batchTaskProfileId)

	return s.GetWithContext(ctx)
}

func batchTaskProfileWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func batchTaskProfileWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_batch.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_batch.BatchComputingClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "batch")
	retryPolicy.ShouldRetryOperation = batchTaskProfileWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromBatchBatchTaskProfileWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBatchBatchTaskProfileWorkRequest(ctx context.Context, client *oci_batch.BatchComputingClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_batch.ActionTypeEnum) error {
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

func (s *BatchBatchTaskProfileResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.GetBatchTaskProfileRequest{}

	tmp := s.D.Id()
	request.BatchTaskProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.GetBatchTaskProfile(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BatchTaskProfile
	return nil
}

func (s *BatchBatchTaskProfileResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_batch.UpdateBatchTaskProfileRequest{}

	tmp := s.D.Id()
	request.BatchTaskProfileId = &tmp

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

	response, err := s.Client.UpdateBatchTaskProfile(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BatchTaskProfile
	return nil
}

func (s *BatchBatchTaskProfileResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_batch.DeleteBatchTaskProfileRequest{}

	tmp := s.D.Id()
	request.BatchTaskProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	_, err := s.Client.DeleteBatchTaskProfile(ctx, request)
	return err
}

func (s *BatchBatchTaskProfileResourceCrud) SetData() error {
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

	if s.Res.MinMemoryInGBs != nil {
		s.D.Set("min_memory_in_gbs", *s.Res.MinMemoryInGBs)
	}

	if s.Res.MinOcpus != nil {
		s.D.Set("min_ocpus", *s.Res.MinOcpus)
	}

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

func BatchTaskProfileSummaryToMap(obj oci_batch.BatchTaskProfileSummary) map[string]interface{} {
	result := map[string]interface{}{}

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

	if obj.MinMemoryInGBs != nil {
		result["min_memory_in_gbs"] = int(*obj.MinMemoryInGBs)
	}

	if obj.MinOcpus != nil {
		result["min_ocpus"] = int(*obj.MinOcpus)
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

func (s *BatchBatchTaskProfileResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_batch.ChangeBatchTaskProfileCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BatchTaskProfileId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.ChangeBatchTaskProfileCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBatchTaskProfileFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch"), oci_batch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
