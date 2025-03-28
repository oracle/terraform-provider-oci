// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v65/aianomalydetection"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func AiAnomalyDetectionProjectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiAnomalyDetectionProject,
		Read:     readAiAnomalyDetectionProject,
		Update:   updateAiAnomalyDetectionProject,
		Delete:   deleteAiAnomalyDetectionProject,
		Schema: map[string]*schema.Schema{
			// Required
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

func createAiAnomalyDetectionProject(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.CreateResource(d, sync)
}

func readAiAnomalyDetectionProject(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

func updateAiAnomalyDetectionProject(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiAnomalyDetectionProject(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiAnomalyDetectionProjectResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res                    *oci_ai_anomaly_detection.Project
	DisableNotFoundRetries bool
}

func (s *AiAnomalyDetectionProjectResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiAnomalyDetectionProjectResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_anomaly_detection.ProjectLifecycleStateCreating),
	}
}

func (s *AiAnomalyDetectionProjectResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_anomaly_detection.ProjectLifecycleStateActive),
	}
}

func (s *AiAnomalyDetectionProjectResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_anomaly_detection.ProjectLifecycleStateDeleting),
	}
}

func (s *AiAnomalyDetectionProjectResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_anomaly_detection.ProjectLifecycleStateDeleted),
	}
}

func (s *AiAnomalyDetectionProjectResourceCrud) Create() error {
	request := oci_ai_anomaly_detection.CreateProjectRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.CreateProject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Project
	return nil
}

func (s *AiAnomalyDetectionProjectResourceCrud) getProjectFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_anomaly_detection.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	projectId, err := aiAnomalyDetectionProjectWaitForWorkRequest(workId, "ai_anomaly_detection",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, projectId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_ai_anomaly_detection.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*projectId)

	return s.Get()
}

func aiAnomalyDetectionProjectWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ai_anomaly_detection", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ai_anomaly_detection.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func aiAnomalyDetectionProjectWaitForWorkRequest(wId *string, entityType string, action oci_ai_anomaly_detection.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_anomaly_detection.AnomalyDetectionClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_anomaly_detection")
	retryPolicy.ShouldRetryOperation = aiAnomalyDetectionProjectWorkRequestShouldRetryFunc(timeout)

	response := oci_ai_anomaly_detection.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_ai_anomaly_detection.OperationStatusInProgress),
			string(oci_ai_anomaly_detection.OperationStatusAccepted),
			string(oci_ai_anomaly_detection.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ai_anomaly_detection.OperationStatusSucceeded),
			string(oci_ai_anomaly_detection.OperationStatusFailed),
			string(oci_ai_anomaly_detection.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ai_anomaly_detection.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ai_anomaly_detection.OperationStatusFailed || response.Status == oci_ai_anomaly_detection.OperationStatusCanceled {
		return nil, getErrorFromAiAnomalyDetectionProjectWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAiAnomalyDetectionProjectWorkRequest(client *oci_ai_anomaly_detection.AnomalyDetectionClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_anomaly_detection.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ai_anomaly_detection.ListWorkRequestErrorsRequest{
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

func (s *AiAnomalyDetectionProjectResourceCrud) Get() error {
	request := oci_ai_anomaly_detection.GetProjectRequest{}

	tmp := s.D.Id()
	request.ProjectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.GetProject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Project
	return nil
}

func (s *AiAnomalyDetectionProjectResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_anomaly_detection.UpdateProjectRequest{}

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

	tmp := s.D.Id()
	request.ProjectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.UpdateProject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Project
	return nil
}

func (s *AiAnomalyDetectionProjectResourceCrud) Delete() error {
	request := oci_ai_anomaly_detection.DeleteProjectRequest{}

	tmp := s.D.Id()
	request.ProjectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.DeleteProject(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := aiAnomalyDetectionProjectWaitForWorkRequest(workId, "projects",
		oci_ai_anomaly_detection.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiAnomalyDetectionProjectResourceCrud) SetData() error {
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

	// if s.Res.SystemTags != nil {
	// 	s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	// }

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func aiAnomalyDetectionProjectSummaryToMap(obj oci_ai_anomaly_detection.ProjectSummary) map[string]interface{} {
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

	result["state"] = string(obj.LifecycleState)

	// if obj.SystemTags != nil {
	// 	result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	// }

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *AiAnomalyDetectionProjectResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_anomaly_detection.ChangeProjectCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ProjectId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	_, err := s.Client.ChangeProjectCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
