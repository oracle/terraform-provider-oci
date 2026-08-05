// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeAuditTrailResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeAuditTrailWithContext,
		ReadContext:   readDataSafeAuditTrailWithContext,
		UpdateContext: updateDataSafeAuditTrailWithContext,
		DeleteContext: deleteDataSafeAuditTrailWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"audit_trail_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"can_update_last_archive_time_on_target": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"is_auto_purge_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_data_safe.AuditTrailLifecycleStateInactive),
					string(oci_data_safe.AuditTrailLifecycleStateActive),
				}, true),
			},
			"resume_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"audit_collection_start_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"audit_profile_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_target_database_key": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"purge_job_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"purge_job_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"purge_job_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"target_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_collected": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trail_location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trail_source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeAuditTrailWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeAuditTrailResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	var powerOff = false
	var powerOn = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_data_safe.AuditTrailLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_data_safe.AuditTrailLifecycleStateActive {
			powerOn = true
		}
		if wantedPowerState == oci_data_safe.AuditTrailLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync)); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("resume_trigger"); ok {
		err := sync.ResumeAuditTrail(ctx)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	oldRaw, newRaw := sync.D.GetChange("resume_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if powerOn && !(oldValue < newValue) {
		if err := sync.StartAuditTrail(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_data_safe.AuditTrailLifecycleStateActive)
	}

	if powerOff {
		if err := sync.StopAuditTrail(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_data_safe.AuditTrailLifecycleStateInactive)
	}
	return nil

}

func readDataSafeAuditTrailWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeAuditTrailResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataSafeAuditTrailWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeAuditTrailResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	powerOn, powerOff := false, false
	wantedState := strings.ToUpper(sync.D.Get("state").(string))
	if oci_data_safe.AuditTrailLifecycleStateActive == oci_data_safe.AuditTrailLifecycleStateEnum(wantedState) {
		powerOn = true
	} else if oci_data_safe.AuditTrailLifecycleStateInactive == oci_data_safe.AuditTrailLifecycleStateEnum(wantedState) {
		powerOff = true
	}

	oldRaw, newRaw := sync.D.GetChange("resume_trigger")
	oldValue := oldRaw.(int)
	newValue := newRaw.(int)
	if powerOn && !(oldValue < newValue) {
		if err := sync.StartAuditTrail(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_data_safe.AuditTrailLifecycleStateActive)
	}

	if _, ok := sync.D.GetOkExists("resume_trigger"); ok && sync.D.HasChange("resume_trigger") {
		if oldValue < newValue {
			err := sync.ResumeAuditTrail(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("resume_trigger", oldRaw)
			return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
		}
	}

	if err := tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync)); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopAuditTrail(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_data_safe.AuditTrailLifecycleStateInactive)
	}

	return nil
}

func deleteDataSafeAuditTrailWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeAuditTrailResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataSafeAuditTrailResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.AuditTrail
	DisableNotFoundRetries bool
}

func (s *DataSafeAuditTrailResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeAuditTrailResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DataSafeAuditTrailResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.AuditTrailLifecycleStateActive),
		string(oci_data_safe.AuditTrailLifecycleStateNeedsAttention),
		string(oci_data_safe.AuditTrailLifecycleStateInactive),
	}
}

func (s *DataSafeAuditTrailResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.AuditTrailLifecycleStateDeleting),
	}
}

func (s *DataSafeAuditTrailResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DataSafeAuditTrailResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_data_safe.UpdateAuditTrailRequest{}

	if auditTrailId, ok := s.D.GetOkExists("audit_trail_id"); ok {
		tmp := auditTrailId.(string)
		request.AuditTrailId = &tmp
	}

	if canUpdateLastArchiveTimeOnTarget, ok := s.D.GetOkExists("can_update_last_archive_time_on_target"); ok {
		tmp := canUpdateLastArchiveTimeOnTarget.(bool)
		request.CanUpdateLastArchiveTimeOnTarget = &tmp
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

	if isAutoPurgeEnabled, ok := s.D.GetOkExists("is_auto_purge_enabled"); ok {
		tmp := isAutoPurgeEnabled.(bool)
		request.IsAutoPurgeEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAuditTrail(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_data_safe.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(ctx,
		oci_data_safe.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "audittrail") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getAuditTrailFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeAuditTrailResourceCrud) getAuditTrailFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	auditTrailId, err := auditTrailWaitForWorkRequest(ctx, workId, "audittrail",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*auditTrailId)

	return s.GetWithContext(ctx)
}

func auditTrailWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func auditTrailWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = auditTrailWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
			string(oci_data_safe.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_data_safe.GetWorkRequestRequest{
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
	if _, e := stateConf.WaitForStateContext(ctx); e != nil {
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeAuditTrailWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeAuditTrailWorkRequest(ctx context.Context, client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_data_safe.ListWorkRequestErrorsRequest{
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

func (s *DataSafeAuditTrailResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetAuditTrailRequest{}

	tmp := s.D.Id()
	request.AuditTrailId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetAuditTrail(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.AuditTrail
	return nil
}

func (s *DataSafeAuditTrailResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_data_safe.UpdateAuditTrailRequest{}

	tmp := s.D.Id()
	request.AuditTrailId = &tmp

	if canUpdateLastArchiveTimeOnTarget, ok := s.D.GetOkExists("can_update_last_archive_time_on_target"); ok {
		tmp := canUpdateLastArchiveTimeOnTarget.(bool)
		request.CanUpdateLastArchiveTimeOnTarget = &tmp
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

	if isAutoPurgeEnabled, ok := s.D.GetOkExists("is_auto_purge_enabled"); ok {
		tmp := isAutoPurgeEnabled.(bool)
		request.IsAutoPurgeEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAuditTrail(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditTrailFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditTrailResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_data_safe.DeleteAuditTrailRequest{}

	tmp := s.D.Id()
	request.AuditTrailId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteAuditTrail(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := auditTrailWaitForWorkRequest(ctx, workId, "audittrail",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeAuditTrailResourceCrud) SetData() error {
	if s.Res.AuditCollectionStartTime != nil {
		s.D.Set("audit_collection_start_time", s.Res.AuditCollectionStartTime.String())
	}

	if s.Res.AuditProfileId != nil {
		s.D.Set("audit_profile_id", *s.Res.AuditProfileId)
	}

	if s.Res.CanUpdateLastArchiveTimeOnTarget != nil {
		s.D.Set("can_update_last_archive_time_on_target", *s.Res.CanUpdateLastArchiveTimeOnTarget)
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

	if s.Res.IsAutoPurgeEnabled != nil {
		s.D.Set("is_auto_purge_enabled", *s.Res.IsAutoPurgeEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PeerTargetDatabaseKey != nil {
		s.D.Set("peer_target_database_key", *s.Res.PeerTargetDatabaseKey)
	}

	if s.Res.PurgeJobDetails != nil {
		s.D.Set("purge_job_details", *s.Res.PurgeJobDetails)
	}

	s.D.Set("purge_job_status", s.Res.PurgeJobStatus)

	if s.Res.PurgeJobTime != nil {
		s.D.Set("purge_job_time", s.Res.PurgeJobTime.String())
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastCollected != nil {
		s.D.Set("time_last_collected", s.Res.TimeLastCollected.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TrailLocation != nil {
		s.D.Set("trail_location", *s.Res.TrailLocation)
	}

	s.D.Set("trail_source", s.Res.TrailSource)

	if s.Res.WorkRequestId != nil {
		s.D.Set("work_request_id", *s.Res.WorkRequestId)
	}

	return nil
}

func (s *DataSafeAuditTrailResourceCrud) StartAuditTrail(ctx context.Context) error {
	request := oci_data_safe.StartAuditTrailRequest{}

	tmp, err := time.Parse(time.RFC3339, "2025-06-01T00:00:00.000Z")
	if err != nil {
		return err
	}
	request.AuditCollectionStartTime = &oci_common.SDKTime{Time: tmp}

	idTmp := s.D.Id()
	request.AuditTrailId = &idTmp

	if canUpdateLastArchiveTimeOnTarget, ok := s.D.GetOkExists("can_update_last_archive_time_on_target"); ok {
		tmp := canUpdateLastArchiveTimeOnTarget.(bool)
		request.CanUpdateLastArchiveTimeOnTarget = &tmp
	}

	if isAutoPurgeEnabled, ok := s.D.GetOkExists("is_auto_purge_enabled"); ok {
		tmp := isAutoPurgeEnabled.(bool)
		request.IsAutoPurgeEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err = s.Client.StartAuditTrail(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_data_safe.AuditTrailLifecycleStateActive }
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditTrailResourceCrud) StopAuditTrail(ctx context.Context) error {
	request := oci_data_safe.StopAuditTrailRequest{}

	idTmp := s.D.Id()
	request.AuditTrailId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.StopAuditTrail(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_data_safe.AuditTrailLifecycleStateInactive }
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditTrailResourceCrud) ResumeAuditTrail(ctx context.Context) error {
	request := oci_data_safe.ResumeAuditTrailRequest{}

	idTmp := s.D.Id()
	request.AuditTrailId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ResumeAuditTrail(ctx, request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("resume_trigger")
	s.D.Set("resume_trigger", val)

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_data_safe.AuditTrailLifecycleStateActive }
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func AuditTrailSummaryToMap(obj oci_data_safe.AuditTrailSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuditCollectionStartTime != nil {
		result["audit_collection_start_time"] = obj.AuditCollectionStartTime.String()
	}

	if obj.AuditProfileId != nil {
		result["audit_profile_id"] = string(*obj.AuditProfileId)
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

	if obj.IsAutoPurgeEnabled != nil {
		result["is_auto_purge_enabled"] = bool(*obj.IsAutoPurgeEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.PeerTargetDatabaseKey != nil {
		result["peer_target_database_key"] = int(*obj.PeerTargetDatabaseKey)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TrailLocation != nil {
		result["trail_location"] = string(*obj.TrailLocation)
	}

	return result
}
