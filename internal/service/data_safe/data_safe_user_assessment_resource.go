// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeUserAssessmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("45m"),
			Update: tfresource.GetTimeoutDuration("45m"),
			Delete: tfresource.GetTimeoutDuration("45m"),
		},
		Create: createDataSafeUserAssessment,
		Read:   readDataSafeUserAssessment,
		Update: updateDataSafeUserAssessment,
		Delete: deleteDataSafeUserAssessment,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"schedule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"ignored_assessment_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignored_targets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lifecycle_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_assessment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_baseline": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_deviated_from_baseline": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_compared_baseline_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_assessment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"statistics": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"target_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_assessed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"triggered_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeUserAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeUserAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeUserAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeUserAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeUserAssessmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.UserAssessment
	DisableNotFoundRetries bool
}

func (s *DataSafeUserAssessmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeUserAssessmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.UserAssessmentLifecycleStateCreating),
	}
}

func (s *DataSafeUserAssessmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.UserAssessmentLifecycleStateSucceeded),
	}
}

func (s *DataSafeUserAssessmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.UserAssessmentLifecycleStateDeleting),
	}
}

func (s *DataSafeUserAssessmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.UserAssessmentLifecycleStateDeleted),
	}
}

func (s *DataSafeUserAssessmentResourceCrud) Create() error {
	request := oci_data_safe.CreateUserAssessmentRequest{}

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

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		tmp := schedule.(string)
		request.Schedule = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateUserAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getUserAssessmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeUserAssessmentResourceCrud) getUserAssessmentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	userAssessmentId, err := userAssessmentWaitForWorkRequest(workId, "userassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*userAssessmentId)

	return s.Get()
}

func userAssessmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func userAssessmentWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = userAssessmentWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed {
		return nil, getErrorFromDataSafeUserAssessmentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeUserAssessmentWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeUserAssessmentResourceCrud) Get() error {
	request := oci_data_safe.GetUserAssessmentRequest{}

	tmp := s.D.Id()
	request.UserAssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetUserAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UserAssessment
	return nil
}

func (s *DataSafeUserAssessmentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateUserAssessmentRequest{}

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

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		tmp := schedule.(string)
		request.Schedule = &tmp
	}

	tmp := s.D.Id()
	request.UserAssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateUserAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUserAssessmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeUserAssessmentResourceCrud) Delete() error {
	request := oci_data_safe.DeleteUserAssessmentRequest{}

	tmp := s.D.Id()
	request.UserAssessmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteUserAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := userAssessmentWaitForWorkRequest(workId, "userassessment",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeUserAssessmentResourceCrud) SetData() error {
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

	ignoredAssessmentIds := []interface{}{}
	for _, item := range s.Res.IgnoredAssessmentIds {
		ignoredAssessmentIds = append(ignoredAssessmentIds, item)
	}
	s.D.Set("ignored_assessment_ids", ignoredAssessmentIds)

	ignoredTargets := []interface{}{}
	for _, item := range s.Res.IgnoredTargets {
		ignoredTargets = append(ignoredTargets, item)
	}
	s.D.Set("ignored_targets", ignoredTargets)

	if s.Res.IsBaseline != nil {
		s.D.Set("is_baseline", *s.Res.IsBaseline)
	}

	if s.Res.IsDeviatedFromBaseline != nil {
		s.D.Set("is_deviated_from_baseline", *s.Res.IsDeviatedFromBaseline)
	}

	if s.Res.LastComparedBaselineId != nil {
		s.D.Set("last_compared_baseline_id", *s.Res.LastComparedBaselineId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Schedule != nil {
		s.D.Set("schedule", *s.Res.Schedule)
	}

	if s.Res.ScheduleAssessmentId != nil {
		s.D.Set("schedule_assessment_id", *s.Res.ScheduleAssessmentId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	var statsBytes, err = json.Marshal(s.Res.Statistics)
	if err != nil {
		return fmt.Errorf("unable to Marshal Statistics, encountered error: %v", err)
	}
	statDetails := string(statsBytes)

	s.D.Set("statistics", statDetails)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("target_ids", s.Res.TargetIds)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastAssessed != nil {
		s.D.Set("time_last_assessed", s.Res.TimeLastAssessed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("triggered_by", s.Res.TriggeredBy)

	s.D.Set("type", s.Res.Type)

	return nil
}

func (s *DataSafeUserAssessmentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeUserAssessmentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.UserAssessmentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeUserAssessmentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
