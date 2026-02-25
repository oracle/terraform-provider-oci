// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiAgentAgentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createGenerativeAiAgentAgentWithContext,
		ReadContext:   readGenerativeAiAgentAgentWithContext,
		UpdateContext: updateGenerativeAiAgentAgentWithContext,
		DeleteContext: deleteGenerativeAiAgentAgentWithContext,
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
			"knowledge_base_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"llm_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"routing_llm_customization": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"instruction": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"llm_hyper_parameters": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"llm_selection": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"llm_selection_type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"CUSTOM_GEN_AI_ENDPOINT",
														"CUSTOM_GEN_AI_MODEL",
														"DEFAULT",
													}, true),
												},

												// Optional
												"endpoint_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"model_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"runtime_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"welcome_message": {
				Type:     schema.TypeString,
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

func createGenerativeAiAgentAgentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiAgentAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readGenerativeAiAgentAgentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiAgentAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateGenerativeAiAgentAgentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiAgentAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteGenerativeAiAgentAgentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiAgentAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type GenerativeAiAgentAgentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai_agent.GenerativeAiAgentClient
	Res                    *oci_generative_ai_agent.Agent
	DisableNotFoundRetries bool
}

func (s *GenerativeAiAgentAgentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiAgentAgentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai_agent.AgentLifecycleStateCreating),
	}
}

func (s *GenerativeAiAgentAgentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.AgentLifecycleStateActive),
	}
}

func (s *GenerativeAiAgentAgentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai_agent.AgentLifecycleStateDeleting),
	}
}

func (s *GenerativeAiAgentAgentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.AgentLifecycleStateDeleted),
	}
}

func (s *GenerativeAiAgentAgentResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_generative_ai_agent.CreateAgentRequest{}

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

	if knowledgeBaseIds, ok := s.D.GetOkExists("knowledge_base_ids"); ok {
		interfaces := knowledgeBaseIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("knowledge_base_ids") {
			request.KnowledgeBaseIds = tmp
		}
	}

	if llmConfig, ok := s.D.GetOkExists("llm_config"); ok {
		if tmpList := llmConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "llm_config", 0)
			tmp, err := s.mapToLlmConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LlmConfig = &tmp
		}
	}

	if welcomeMessage, ok := s.D.GetOkExists("welcome_message"); ok {
		tmp := welcomeMessage.(string)
		request.WelcomeMessage = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.CreateAgent(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAgentFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiAgentAgentResourceCrud) getAgentFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai_agent.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	agentId, err := agentWaitForWorkRequest(ctx, workId, "agent",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, agentId)
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
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
	s.D.SetId(*agentId)

	return s.GetWithContext(ctx)
}

func agentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func agentWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_generative_ai_agent.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai_agent.GenerativeAiAgentClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai_agent")
	retryPolicy.ShouldRetryOperation = agentWorkRequestShouldRetryFunc(timeout)

	response := oci_generative_ai_agent.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
			response, err = client.GetWorkRequest(ctx,
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
		return nil, getErrorFromGenerativeAiAgentAgentWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiAgentAgentWorkRequest(ctx context.Context, client *oci_generative_ai_agent.GenerativeAiAgentClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai_agent.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
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

func (s *GenerativeAiAgentAgentResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai_agent.GetAgentRequest{}

	tmp := s.D.Id()
	request.AgentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.GetAgent(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Agent
	return nil
}

func (s *GenerativeAiAgentAgentResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai_agent.UpdateAgentRequest{}

	tmp := s.D.Id()
	request.AgentId = &tmp

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

	if knowledgeBaseIds, ok := s.D.GetOkExists("knowledge_base_ids"); ok {
		interfaces := knowledgeBaseIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("knowledge_base_ids") {
			request.KnowledgeBaseIds = tmp
		}
	}

	if llmConfig, ok := s.D.GetOkExists("llm_config"); ok {
		if tmpList := llmConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "llm_config", 0)
			tmp, err := s.mapToLlmConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LlmConfig = &tmp
		}
	}

	if welcomeMessage, ok := s.D.GetOkExists("welcome_message"); ok {
		tmp := welcomeMessage.(string)
		request.WelcomeMessage = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.UpdateAgent(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAgentFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiAgentAgentResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_generative_ai_agent.DeleteAgentRequest{}

	tmp := s.D.Id()
	request.AgentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.DeleteAgent(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := agentWaitForWorkRequest(ctx, workId, "agent",
		oci_generative_ai_agent.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiAgentAgentResourceCrud) SetData() error {
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

	s.D.Set("knowledge_base_ids", s.Res.KnowledgeBaseIds)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LlmConfig != nil {
		s.D.Set("llm_config", []interface{}{LlmConfigToMap(s.Res.LlmConfig)})
	} else {
		s.D.Set("llm_config", nil)
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

	if s.Res.WelcomeMessage != nil {
		s.D.Set("welcome_message", *s.Res.WelcomeMessage)
	}

	return nil
}

func AgentSummaryToMap(obj oci_generative_ai_agent.AgentSummary) map[string]interface{} {
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

	result["knowledge_base_ids"] = obj.KnowledgeBaseIds

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.LlmConfig != nil {
		result["llm_config"] = []interface{}{LlmConfigToMap(obj.LlmConfig)}
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

	if obj.WelcomeMessage != nil {
		result["welcome_message"] = string(*obj.WelcomeMessage)
	}

	return result
}

func (s *GenerativeAiAgentAgentResourceCrud) mapToLlmConfig(fieldKeyFormat string) (oci_generative_ai_agent.LlmConfig, error) {
	result := oci_generative_ai_agent.LlmConfig{}

	if routingLlmCustomization, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "routing_llm_customization")); ok {
		if tmpList := routingLlmCustomization.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "routing_llm_customization"), 0)
			tmp, err := s.mapToLlmCustomization(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert routing_llm_customization, encountered error: %v", err)
			}
			result.RoutingLlmCustomization = &tmp
		}
	}

	if runtimeVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runtime_version")); ok {
		tmp := runtimeVersion.(string)
		result.RuntimeVersion = &tmp
	}

	return result, nil
}

func LlmConfigToMap(obj *oci_generative_ai_agent.LlmConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RoutingLlmCustomization != nil {
		result["routing_llm_customization"] = []interface{}{LlmCustomizationToMap(obj.RoutingLlmCustomization)}
	}

	if obj.RuntimeVersion != nil {
		result["runtime_version"] = string(*obj.RuntimeVersion)
	}

	return result
}

func (s *GenerativeAiAgentAgentResourceCrud) mapToLlmCustomization(fieldKeyFormat string) (oci_generative_ai_agent.LlmCustomization, error) {
	result := oci_generative_ai_agent.LlmCustomization{}

	if instruction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instruction")); ok {
		tmp := instruction.(string)
		result.Instruction = &tmp
	}

	if llmHyperParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "llm_hyper_parameters")); ok {
		result.LlmHyperParameters = llmHyperParameters.(map[string]interface{})
	}

	if llmSelection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "llm_selection")); ok {
		if tmpList := llmSelection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "llm_selection"), 0)
			tmp, err := s.mapToLlmSelection(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert llm_selection, encountered error: %v", err)
			}
			result.LlmSelection = tmp
		}
	}

	return result, nil
}

func LlmCustomizationToMap(obj *oci_generative_ai_agent.LlmCustomization) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Instruction != nil {
		result["instruction"] = string(*obj.Instruction)
	}

	result["llm_hyper_parameters"] = obj.LlmHyperParameters

	if obj.LlmSelection != nil {
		llmSelectionArray := []interface{}{}
		if llmSelectionMap := LlmSelectionToMap(&obj.LlmSelection); llmSelectionMap != nil {
			llmSelectionArray = append(llmSelectionArray, llmSelectionMap)
		}
		result["llm_selection"] = llmSelectionArray
	}

	return result
}

func (s *GenerativeAiAgentAgentResourceCrud) mapToLlmSelection(fieldKeyFormat string) (oci_generative_ai_agent.LlmSelection, error) {
	var baseObject oci_generative_ai_agent.LlmSelection
	//discriminator
	llmSelectionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "llm_selection_type"))
	var llmSelectionType string
	if ok {
		llmSelectionType = llmSelectionTypeRaw.(string)
	} else {
		llmSelectionType = "" // default value
	}
	switch strings.ToLower(llmSelectionType) {
	case strings.ToLower("CUSTOM_GEN_AI_ENDPOINT"):
		details := oci_generative_ai_agent.CustomGenAiEndpointLlmSelection{}
		if endpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "endpoint_id")); ok {
			tmp := endpointId.(string)
			details.EndpointId = &tmp
		}
		baseObject = details
	case strings.ToLower("CUSTOM_GEN_AI_MODEL"):
		details := oci_generative_ai_agent.CustomGenAiModelLlmSelection{}
		if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
			tmp := modelId.(string)
			details.ModelId = &tmp
		}
		baseObject = details
	case strings.ToLower("DEFAULT"):
		details := oci_generative_ai_agent.DefaultLlmSelection{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown llm_selection_type '%v' was specified", llmSelectionType)
	}
	return baseObject, nil
}

func LlmSelectionToMap(obj *oci_generative_ai_agent.LlmSelection) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai_agent.CustomGenAiEndpointLlmSelection:
		result["llm_selection_type"] = "CUSTOM_GEN_AI_ENDPOINT"

		if v.EndpointId != nil {
			result["endpoint_id"] = string(*v.EndpointId)
		}
	case oci_generative_ai_agent.CustomGenAiModelLlmSelection:
		result["llm_selection_type"] = "CUSTOM_GEN_AI_MODEL"

		if v.ModelId != nil {
			result["model_id"] = string(*v.ModelId)
		}
	case oci_generative_ai_agent.DefaultLlmSelection:
		result["llm_selection_type"] = "DEFAULT"
	default:
		log.Printf("[WARN] Received 'llm_selection_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiAgentAgentResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai_agent.ChangeAgentCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AgentId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.ChangeAgentCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAgentFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
