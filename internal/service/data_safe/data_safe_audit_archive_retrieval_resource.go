// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditArchiveRetrievalResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeAuditArchiveRetrieval,
		Read:     readDataSafeAuditArchiveRetrieval,
		Update:   updateDataSafeAuditArchiveRetrieval,
		Delete:   deleteDataSafeAuditArchiveRetrieval,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"end_date": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"start_date": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"target_id": {
				Type:     schema.TypeString,
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
			"audit_event_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_completed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_expiry": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_requested": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeAuditArchiveRetrieval(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditArchiveRetrievalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeAuditArchiveRetrieval(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditArchiveRetrievalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeAuditArchiveRetrieval(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditArchiveRetrievalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeAuditArchiveRetrieval(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditArchiveRetrievalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeAuditArchiveRetrievalResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.AuditArchiveRetrieval
	DisableNotFoundRetries bool
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.AuditArchiveRetrievalLifecycleStateCreating),
	}
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.AuditArchiveRetrievalLifecycleStateActive),
		string(oci_data_safe.AuditArchiveRetrievalLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.AuditArchiveRetrievalLifecycleStateDeleting),
	}
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.AuditArchiveRetrievalLifecycleStateDeleted),
	}
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) Create() error {
	request := oci_data_safe.CreateAuditArchiveRetrievalRequest{}

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

	if endDate, ok := s.D.GetOkExists("end_date"); ok {
		tmp, err := time.Parse(time.RFC3339, endDate.(string))
		if err != nil {
			return err
		}
		request.EndDate = &oci_common.SDKTime{Time: tmp}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if startDate, ok := s.D.GetOkExists("start_date"); ok {
		tmp, err := time.Parse(time.RFC3339, startDate.(string))
		if err != nil {
			return err
		}
		request.StartDate = &oci_common.SDKTime{Time: tmp}
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateAuditArchiveRetrieval(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAuditArchiveRetrievalFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) getAuditArchiveRetrievalFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	auditArchiveRetrievalId, err := auditArchiveRetrievalWaitForWorkRequest(workId, "data_safe",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*auditArchiveRetrievalId)

	return s.Get()
}

func auditArchiveRetrievalWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func auditArchiveRetrievalWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = auditArchiveRetrievalWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
			response, err = client.GetWorkRequest(context.Background(),
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeAuditArchiveRetrievalWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeAuditArchiveRetrievalWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
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

func (s *DataSafeAuditArchiveRetrievalResourceCrud) Get() error {
	request := oci_data_safe.GetAuditArchiveRetrievalRequest{}

	tmp := s.D.Id()
	request.AuditArchiveRetrievalId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetAuditArchiveRetrieval(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuditArchiveRetrieval
	return nil
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateAuditArchiveRetrievalRequest{}

	tmp := s.D.Id()
	request.AuditArchiveRetrievalId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAuditArchiveRetrieval(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditArchiveRetrievalFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) Delete() error {
	request := oci_data_safe.DeleteAuditArchiveRetrievalRequest{}

	tmp := s.D.Id()
	request.AuditArchiveRetrievalId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteAuditArchiveRetrieval(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := auditArchiveRetrievalWaitForWorkRequest(workId, "data_safe",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) SetData() error {
	if s.Res.AuditEventCount != nil {
		s.D.Set("audit_event_count", strconv.FormatInt(*s.Res.AuditEventCount, 10))
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

	if s.Res.EndDate != nil {
		s.D.Set("end_date", s.Res.EndDate.Format(time.RFC3339Nano))
	}

	if s.Res.ErrorInfo != nil {
		s.D.Set("error_info", *s.Res.ErrorInfo)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.StartDate != nil {
		s.D.Set("start_date", s.Res.StartDate.Format(time.RFC3339Nano))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCompleted != nil {
		s.D.Set("time_completed", s.Res.TimeCompleted.String())
	}

	if s.Res.TimeOfExpiry != nil {
		s.D.Set("time_of_expiry", s.Res.TimeOfExpiry.String())
	}

	if s.Res.TimeRequested != nil {
		s.D.Set("time_requested", s.Res.TimeRequested.String())
	}

	return nil
}

func AuditArchiveRetrievalSummaryToMap(obj oci_data_safe.AuditArchiveRetrievalSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuditEventCount != nil {
		result["audit_event_count"] = strconv.FormatInt(*obj.AuditEventCount, 10)
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

	if obj.EndDate != nil {
		result["end_date"] = obj.EndDate.String()
	}

	if obj.ErrorInfo != nil {
		result["error_info"] = string(*obj.ErrorInfo)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.StartDate != nil {
		result["start_date"] = obj.StartDate.String()
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCompleted != nil {
		result["time_completed"] = obj.TimeCompleted.String()
	}

	if obj.TimeOfExpiry != nil {
		result["time_of_expiry"] = obj.TimeOfExpiry.String()
	}

	if obj.TimeRequested != nil {
		result["time_requested"] = obj.TimeRequested.String()
	}

	return result
}

func (s *DataSafeAuditArchiveRetrievalResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeAuditArchiveRetrievalCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AuditArchiveRetrievalId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeAuditArchiveRetrievalCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditArchiveRetrievalFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
