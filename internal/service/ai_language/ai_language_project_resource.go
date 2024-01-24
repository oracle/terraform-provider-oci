// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_language

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiLanguageProjectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiLanguageProject,
		Read:     readAiLanguageProject,
		Update:   updateAiLanguageProject,
		Delete:   deleteAiLanguageProject,
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

func createAiLanguageProject(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.CreateResource(d, sync)
}

func readAiLanguageProject(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.ReadResource(sync)
}

func updateAiLanguageProject(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiLanguageProject(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiLanguageProjectResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_language.AIServiceLanguageClient
	Res                    *oci_ai_language.Project
	DisableNotFoundRetries bool
}

func (s *AiLanguageProjectResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiLanguageProjectResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_language.ProjectLifecycleStateCreating),
	}
}

func (s *AiLanguageProjectResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_language.ProjectLifecycleStateActive),
	}
}

func (s *AiLanguageProjectResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_ai_language.ProjectLifecycleStateUpdating),
	}
}

func (s *AiLanguageProjectResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_ai_language.ProjectLifecycleStateActive),
	}
}

func (s *AiLanguageProjectResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_language.ProjectLifecycleStateDeleting),
	}
}

func (s *AiLanguageProjectResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_language.ProjectLifecycleStateDeleted),
	}
}

func (s *AiLanguageProjectResourceCrud) Create() error {
	request := oci_ai_language.CreateProjectRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.CreateProject(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getProjectFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language"), "CREATED", s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiLanguageProjectResourceCrud) getProjectFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_language.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	projectId, err := projectWaitForWorkRequest(workId, "project",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*projectId)

	return s.Get()
}

func projectWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ai_language", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ai_language.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func projectWaitForWorkRequest(wId *string, entityType string, action oci_ai_language.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_language.AIServiceLanguageClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_language")
	retryPolicy.ShouldRetryOperation = projectWorkRequestShouldRetryFunc(timeout)

	response := oci_ai_language.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_ai_language.ActionTypeInProgress),
		},
		Target: []string{
			string("ACTIVE"),
			string("SUCCEEDED"),
			string("FAILED"),
			string(oci_ai_language.ActionTypeDeleted),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ai_language.GetWorkRequestRequest{
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
	// if identifier == nil || response.Status == oci_ai_language.WorkRequestStatusFailed || response.Status == oci_ai_language.WorkRequestStatusCanceled {
	// 	return nil, getErrorFromAiLanguageProjectWorkRequest(client, wId, retryPolicy, entityType, action)
	// }

	return identifier, nil
}

func getErrorFromAiLanguageProjectWorkRequest(client *oci_ai_language.AIServiceLanguageClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_language.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ai_language.ListWorkRequestErrorsRequest{
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

func (s *AiLanguageProjectResourceCrud) Get() error {
	request := oci_ai_language.GetProjectRequest{}

	tmp := s.D.Id()
	request.ProjectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.GetProject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Project
	return nil
}

func (s *AiLanguageProjectResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_language.UpdateProjectRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	_, err := s.Client.UpdateProject(context.Background(), request)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		fmt.Printf("waitErr: %v\n", waitErr)
		return waitErr
	}

	return nil

	// workId := response.OpcWorkRequestId
	// return s.getProjectFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language"), "CREATED", s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiLanguageProjectResourceCrud) Delete() error {
	request := oci_ai_language.DeleteProjectRequest{}

	tmp := s.D.Id()
	request.ProjectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.DeleteProject(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := projectWaitForWorkRequest(workId, "project",
		oci_ai_language.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiLanguageProjectResourceCrud) SetData() error {
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

	if s.Res.FreeformTags != nil {
		s.D.Set("freeform_tags", s.Res.FreeformTags)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", s.Res.SystemTags)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ProjectSummaryToMap(obj oci_ai_language.ProjectSummary) map[string]interface{} {
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

	if obj.FreeformTags != nil {
		result["freeform_tags"] = obj.FreeformTags
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = obj.SystemTags
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *AiLanguageProjectResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_language.ChangeProjectCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ProjectId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	_, err := s.Client.ChangeProjectCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
