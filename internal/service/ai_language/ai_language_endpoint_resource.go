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

func AiLanguageEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiLanguageEndpoint,
		Read:     readAiLanguageEndpoint,
		Update:   updateAiLanguageEndpoint,
		Delete:   deleteAiLanguageEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_id": {
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
			"inference_units": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_id": {
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

func createAiLanguageEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.CreateResource(d, sync)
}

func readAiLanguageEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.ReadResource(sync)
}

func updateAiLanguageEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiLanguageEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiLanguageEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_language.AIServiceLanguageClient
	Res                    *oci_ai_language.Endpoint
	DisableNotFoundRetries bool
}

func (s *AiLanguageEndpointResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiLanguageEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_language.EndpointLifecycleStateCreating),
	}
}

func (s *AiLanguageEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_language.EndpointLifecycleStateActive),
	}
}

func (s *AiLanguageEndpointResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_ai_language.EndpointLifecycleStateUpdating),
	}
}

func (s *AiLanguageEndpointResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_ai_language.EndpointLifecycleStateActive),
	}
}

func (s *AiLanguageEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_language.EndpointLifecycleStateDeleting),
	}
}

func (s *AiLanguageEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_language.EndpointLifecycleStateDeleted),
	}
}

func (s *AiLanguageEndpointResourceCrud) Create() error {
	request := oci_ai_language.CreateEndpointRequest{}

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

	if inferenceUnits, ok := s.D.GetOkExists("inference_units"); ok {
		tmp := inferenceUnits.(int)
		request.InferenceUnits = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.CreateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language"), oci_ai_language.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiLanguageEndpointResourceCrud) getEndpointFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_language.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	endpointId, err := endpointWaitForWorkRequest(workId, "endpoint",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*endpointId)

	return s.Get()
}

func endpointWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func endpointWaitForWorkRequest(wId *string, entityType string, action oci_ai_language.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_language.AIServiceLanguageClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_language")
	retryPolicy.ShouldRetryOperation = endpointWorkRequestShouldRetryFunc(timeout)

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
	// 	return nil, getErrorFromAiLanguageEndpointWorkRequest(client, wId, retryPolicy, entityType, action)
	// }

	return identifier, nil
}

func getErrorFromAiLanguageEndpointWorkRequest(client *oci_ai_language.AIServiceLanguageClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_language.ActionTypeEnum) error {
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

func (s *AiLanguageEndpointResourceCrud) Get() error {
	request := oci_ai_language.GetEndpointRequest{}

	tmp := s.D.Id()
	request.EndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.GetEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Endpoint
	return nil
}

func (s *AiLanguageEndpointResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_language.UpdateEndpointRequest{}

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

	tmp := s.D.Id()
	request.EndpointId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if inferenceUnits, ok := s.D.GetOkExists("inference_units"); ok {
		tmp := inferenceUnits.(int)
		request.InferenceUnits = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	// response, err := s.Client.UpdateEndpoint(context.Background(), request)
	_, err := s.Client.UpdateEndpoint(context.Background(), request)

	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		fmt.Printf("waitErr: %v\n", waitErr)
		return waitErr
	}

	return nil

	// workId := response.OpcWorkRequestId
	// return s.getEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language"), oci_ai_language.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiLanguageEndpointResourceCrud) Delete() error {
	request := oci_ai_language.DeleteEndpointRequest{}

	tmp := s.D.Id()
	request.EndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.DeleteEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := endpointWaitForWorkRequest(workId, "endpoint",
		oci_ai_language.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiLanguageEndpointResourceCrud) SetData() error {
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

	// s.D.Set("freeform_tags", s.Res.FreeformTags)
	// s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InferenceUnits != nil {
		s.D.Set("inference_units", *s.Res.InferenceUnits)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ModelId != nil {
		s.D.Set("model_id", *s.Res.ModelId)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
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

func EndpointSummaryToMap(obj oci_ai_language.EndpointSummary) map[string]interface{} {
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

	// result["freeform_tags"] = obj.FreeformTags
	// result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InferenceUnits != nil {
		result["inference_units"] = int(*obj.InferenceUnits)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ModelId != nil {
		result["model_id"] = string(*obj.ModelId)
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = obj.SystemTags
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *AiLanguageEndpointResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_language.ChangeEndpointCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.EndpointId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	_, err := s.Client.ChangeEndpointCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
