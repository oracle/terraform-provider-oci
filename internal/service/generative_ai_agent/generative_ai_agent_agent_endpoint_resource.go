// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiAgentAgentEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("40m"),
			Update: tfresource.GetTimeoutDuration("20m"),
			Delete: tfresource.GetTimeoutDuration("20m"),
		},
		Create: createGenerativeAiAgentAgentEndpoint,
		Read:   readGenerativeAiAgentAgentEndpoint,
		Update: updateGenerativeAiAgentAgentEndpoint,
		Delete: deleteGenerativeAiAgentAgentEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"agent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"content_moderation_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"should_enable_on_input": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"should_enable_on_output": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"session_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"idle_timeout_in_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"should_enable_citation": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"should_enable_session": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"should_enable_trace": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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

func createGenerativeAiAgentAgentEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentAgentEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.CreateResource(d, sync)
}

func readGenerativeAiAgentAgentEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentAgentEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

func updateGenerativeAiAgentAgentEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentAgentEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGenerativeAiAgentAgentEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentAgentEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GenerativeAiAgentAgentEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai_agent.GenerativeAiAgentClient
	Res                    *oci_generative_ai_agent.AgentEndpoint
	DisableNotFoundRetries bool
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai_agent.AgentEndpointLifecycleStateCreating),
	}
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.AgentEndpointLifecycleStateActive),
	}
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai_agent.AgentEndpointLifecycleStateDeleting),
	}
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.AgentEndpointLifecycleStateDeleted),
	}
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) Create() error {
	request := oci_generative_ai_agent.CreateAgentEndpointRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if contentModerationConfig, ok := s.D.GetOkExists("content_moderation_config"); ok {
		if tmpList := contentModerationConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "content_moderation_config", 0)
			tmp, err := s.mapToContentModerationConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ContentModerationConfig = &tmp
		}
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

	if sessionConfig, ok := s.D.GetOkExists("session_config"); ok {
		if tmpList := sessionConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "session_config", 0)
			tmp, err := s.mapToSessionConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SessionConfig = &tmp
		}
	}

	if shouldEnableCitation, ok := s.D.GetOkExists("should_enable_citation"); ok {
		tmp := shouldEnableCitation.(bool)
		request.ShouldEnableCitation = &tmp
	}

	if shouldEnableSession, ok := s.D.GetOkExists("should_enable_session"); ok {
		tmp := shouldEnableSession.(bool)
		request.ShouldEnableSession = &tmp
	}

	if shouldEnableTrace, ok := s.D.GetOkExists("should_enable_trace"); ok {
		tmp := shouldEnableTrace.(bool)
		request.ShouldEnableTrace = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.CreateAgentEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAgentEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) getAgentEndpointFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai_agent.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	agentEndpointId, err := agentEndpointWaitForWorkRequest(workId, "agentendpoint",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, agentEndpointId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_generative_ai_agent.CancelWorkRequestRequest{
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
	s.D.SetId(*agentEndpointId)

	return s.Get()
}

func agentEndpointWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "generative_ai_agent", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_generative_ai_agent.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func agentEndpointWaitForWorkRequest(wId *string, entityType string, action oci_generative_ai_agent.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai_agent.GenerativeAiAgentClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai_agent")
	retryPolicy.ShouldRetryOperation = agentEndpointWorkRequestShouldRetryFunc(timeout)

	response := oci_generative_ai_agent.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_generative_ai_agent.OperationStatusInProgress),
			string(oci_generative_ai_agent.OperationStatusAccepted),
			string(oci_generative_ai_agent.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_generative_ai_agent.OperationStatusSucceeded),
			string(oci_generative_ai_agent.OperationStatusFailed),
			string(oci_generative_ai_agent.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_generative_ai_agent.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_generative_ai_agent.OperationStatusFailed || response.Status == oci_generative_ai_agent.OperationStatusCanceled {
		return nil, getErrorFromGenerativeAiAgentAgentEndpointWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiAgentAgentEndpointWorkRequest(client *oci_generative_ai_agent.GenerativeAiAgentClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai_agent.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_generative_ai_agent.ListWorkRequestErrorsRequest{
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

func (s *GenerativeAiAgentAgentEndpointResourceCrud) Get() error {
	request := oci_generative_ai_agent.GetAgentEndpointRequest{}

	tmp := s.D.Id()
	request.AgentEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.GetAgentEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AgentEndpoint
	return nil
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai_agent.UpdateAgentEndpointRequest{}

	tmp := s.D.Id()
	request.AgentEndpointId = &tmp

	if contentModerationConfig, ok := s.D.GetOkExists("content_moderation_config"); ok {
		if tmpList := contentModerationConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "content_moderation_config", 0)
			tmp, err := s.mapToContentModerationConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ContentModerationConfig = &tmp
		}
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

	if sessionConfig, ok := s.D.GetOkExists("session_config"); ok {
		if tmpList := sessionConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "session_config", 0)
			tmp, err := s.mapToSessionConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SessionConfig = &tmp
		}
	}

	if shouldEnableCitation, ok := s.D.GetOkExists("should_enable_citation"); ok {
		tmp := shouldEnableCitation.(bool)
		request.ShouldEnableCitation = &tmp
	}

	if shouldEnableTrace, ok := s.D.GetOkExists("should_enable_trace"); ok {
		tmp := shouldEnableTrace.(bool)
		request.ShouldEnableTrace = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.UpdateAgentEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAgentEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) Delete() error {
	request := oci_generative_ai_agent.DeleteAgentEndpointRequest{}

	tmp := s.D.Id()
	request.AgentEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.DeleteAgentEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := agentEndpointWaitForWorkRequest(workId, "agentendpoint",
		oci_generative_ai_agent.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) SetData() error {
	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ContentModerationConfig != nil {
		s.D.Set("content_moderation_config", []interface{}{ContentModerationConfigToMap(s.Res.ContentModerationConfig)})
	} else {
		s.D.Set("content_moderation_config", nil)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SessionConfig != nil {
		s.D.Set("session_config", []interface{}{SessionConfigToMap(s.Res.SessionConfig)})
	} else {
		s.D.Set("session_config", nil)
	}

	if s.Res.ShouldEnableCitation != nil {
		s.D.Set("should_enable_citation", *s.Res.ShouldEnableCitation)
	}

	if s.Res.ShouldEnableSession != nil {
		s.D.Set("should_enable_session", *s.Res.ShouldEnableSession)
	}

	if s.Res.ShouldEnableTrace != nil {
		s.D.Set("should_enable_trace", *s.Res.ShouldEnableTrace)
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

func AgentEndpointSummaryToMap(obj oci_generative_ai_agent.AgentEndpointSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ContentModerationConfig != nil {
		result["content_moderation_config"] = []interface{}{ContentModerationConfigToMap(obj.ContentModerationConfig)}
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.SessionConfig != nil {
		result["session_config"] = []interface{}{SessionConfigToMap(obj.SessionConfig)}
	}

	if obj.ShouldEnableCitation != nil {
		result["should_enable_citation"] = bool(*obj.ShouldEnableCitation)
	}

	if obj.ShouldEnableSession != nil {
		result["should_enable_session"] = bool(*obj.ShouldEnableSession)
	}

	if obj.ShouldEnableTrace != nil {
		result["should_enable_trace"] = bool(*obj.ShouldEnableTrace)
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

func (s *GenerativeAiAgentAgentEndpointResourceCrud) mapToContentModerationConfig(fieldKeyFormat string) (oci_generative_ai_agent.ContentModerationConfig, error) {
	result := oci_generative_ai_agent.ContentModerationConfig{}

	if shouldEnableOnInput, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_enable_on_input")); ok {
		tmp := shouldEnableOnInput.(bool)
		result.ShouldEnableOnInput = &tmp
	}

	if shouldEnableOnOutput, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_enable_on_output")); ok {
		tmp := shouldEnableOnOutput.(bool)
		result.ShouldEnableOnOutput = &tmp
	}

	return result, nil
}

func ContentModerationConfigToMap(obj *oci_generative_ai_agent.ContentModerationConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ShouldEnableOnInput != nil {
		result["should_enable_on_input"] = bool(*obj.ShouldEnableOnInput)
	}

	if obj.ShouldEnableOnOutput != nil {
		result["should_enable_on_output"] = bool(*obj.ShouldEnableOnOutput)
	}

	return result
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) mapToSessionConfig(fieldKeyFormat string) (oci_generative_ai_agent.SessionConfig, error) {
	result := oci_generative_ai_agent.SessionConfig{}

	if idleTimeoutInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idle_timeout_in_seconds")); ok {
		tmp := idleTimeoutInSeconds.(int)
		result.IdleTimeoutInSeconds = &tmp
	}

	return result, nil
}

func SessionConfigToMap(obj *oci_generative_ai_agent.SessionConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IdleTimeoutInSeconds != nil {
		result["idle_timeout_in_seconds"] = int(*obj.IdleTimeoutInSeconds)
	}

	return result
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai_agent.ChangeAgentEndpointCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AgentEndpointId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.ChangeAgentEndpointCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAgentEndpointFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
