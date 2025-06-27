// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

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
			"guardrail_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

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
									"input_guardrail_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"output_guardrail_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"personally_identifiable_information_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"input_guardrail_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"output_guardrail_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"prompt_injection_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"input_guardrail_mode": {
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
			"human_input_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"should_enable_human_input": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"output_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"output_location": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"bucket": {
										Type:     schema.TypeString,
										Required: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Required: true,
									},
									"output_location_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"OBJECT_STORAGE_PREFIX",
										}, true),
									},

									// Optional
									"prefix": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Optional
						"retention_period_in_minutes": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"should_enable_multi_language": {
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

	if guardrailConfig, ok := s.D.GetOkExists("guardrail_config"); ok {
		if tmpList := guardrailConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "guardrail_config", 0)
			tmp, err := s.mapToGuardrailConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.GuardrailConfig = &tmp
		}
	}

	if humanInputConfig, ok := s.D.GetOkExists("human_input_config"); ok {
		if tmpList := humanInputConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "human_input_config", 0)
			tmp, err := s.mapToHumanInputConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.HumanInputConfig = &tmp
		}
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	if outputConfig, ok := s.D.GetOkExists("output_config"); ok {
		if tmpList := outputConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_config", 0)
			tmp, err := s.mapToOutputConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.OutputConfig = &tmp
		}
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

	if shouldEnableMultiLanguage, ok := s.D.GetOkExists("should_enable_multi_language"); ok {
		tmp := shouldEnableMultiLanguage.(bool)
		request.ShouldEnableMultiLanguage = &tmp
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

	if guardrailConfig, ok := s.D.GetOkExists("guardrail_config"); ok {
		if tmpList := guardrailConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "guardrail_config", 0)
			tmp, err := s.mapToGuardrailConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.GuardrailConfig = &tmp
		}
	}

	if humanInputConfig, ok := s.D.GetOkExists("human_input_config"); ok {
		if tmpList := humanInputConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "human_input_config", 0)
			tmp, err := s.mapToHumanInputConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.HumanInputConfig = &tmp
		}
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	if outputConfig, ok := s.D.GetOkExists("output_config"); ok {
		if tmpList := outputConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_config", 0)
			tmp, err := s.mapToOutputConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.OutputConfig = &tmp
		}
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

	if shouldEnableMultiLanguage, ok := s.D.GetOkExists("should_enable_multi_language"); ok {
		tmp := shouldEnableMultiLanguage.(bool)
		request.ShouldEnableMultiLanguage = &tmp
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

	if s.Res.GuardrailConfig != nil {
		s.D.Set("guardrail_config", []interface{}{GuardrailConfigToMap(s.Res.GuardrailConfig)})
	} else {
		s.D.Set("guardrail_config", nil)
	}

	if s.Res.HumanInputConfig != nil {
		s.D.Set("human_input_config", []interface{}{HumanInputConfigToMap(s.Res.HumanInputConfig)})
	} else {
		s.D.Set("human_input_config", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("metadata", s.Res.Metadata)

	if s.Res.OutputConfig != nil {
		s.D.Set("output_config", []interface{}{OutputConfigToMap(s.Res.OutputConfig)})
	} else {
		s.D.Set("output_config", nil)
	}

	if s.Res.SessionConfig != nil {
		s.D.Set("session_config", []interface{}{SessionConfigToMap(s.Res.SessionConfig)})
	} else {
		s.D.Set("session_config", nil)
	}

	if s.Res.ShouldEnableCitation != nil {
		s.D.Set("should_enable_citation", *s.Res.ShouldEnableCitation)
	}

	if s.Res.ShouldEnableMultiLanguage != nil {
		s.D.Set("should_enable_multi_language", *s.Res.ShouldEnableMultiLanguage)
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

	if obj.GuardrailConfig != nil {
		result["guardrail_config"] = []interface{}{GuardrailConfigToMap(obj.GuardrailConfig)}
	}

	if obj.HumanInputConfig != nil {
		result["human_input_config"] = []interface{}{HumanInputConfigToMap(obj.HumanInputConfig)}
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["metadata"] = obj.Metadata

	if obj.OutputConfig != nil {
		result["output_config"] = []interface{}{OutputConfigToMap(obj.OutputConfig)}
	}

	if obj.SessionConfig != nil {
		result["session_config"] = []interface{}{SessionConfigToMap(obj.SessionConfig)}
	}

	if obj.ShouldEnableCitation != nil {
		result["should_enable_citation"] = bool(*obj.ShouldEnableCitation)
	}

	if obj.ShouldEnableMultiLanguage != nil {
		result["should_enable_multi_language"] = bool(*obj.ShouldEnableMultiLanguage)
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

func (s *GenerativeAiAgentAgentEndpointResourceCrud) mapToContentModerationGuardrailConfig(fieldKeyFormat string) (oci_generative_ai_agent.ContentModerationGuardrailConfig, error) {
	result := oci_generative_ai_agent.ContentModerationGuardrailConfig{}

	if inputGuardrailMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "input_guardrail_mode")); ok {
		result.InputGuardrailMode = oci_generative_ai_agent.GuardrailModeEnum(inputGuardrailMode.(string))
	}

	if outputGuardrailMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_guardrail_mode")); ok {
		result.OutputGuardrailMode = oci_generative_ai_agent.GuardrailModeEnum(outputGuardrailMode.(string))
	}

	return result, nil
}

func ContentModerationGuardrailConfigToMap(obj *oci_generative_ai_agent.ContentModerationGuardrailConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["input_guardrail_mode"] = string(obj.InputGuardrailMode)

	result["output_guardrail_mode"] = string(obj.OutputGuardrailMode)

	return result
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) mapToGuardrailConfig(fieldKeyFormat string) (oci_generative_ai_agent.GuardrailConfig, error) {
	result := oci_generative_ai_agent.GuardrailConfig{}

	if contentModerationConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content_moderation_config")); ok {
		if tmpList := contentModerationConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "content_moderation_config"), 0)
			tmp, err := s.mapToContentModerationGuardrailConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert content_moderation_config, encountered error: %v", err)
			}
			result.ContentModerationConfig = &tmp
		}
	}

	if personallyIdentifiableInformationConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "personally_identifiable_information_config")); ok {
		if tmpList := personallyIdentifiableInformationConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "personally_identifiable_information_config"), 0)
			tmp, err := s.mapToPersonallyIdentifiableInformationGuardrailConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert personally_identifiable_information_config, encountered error: %v", err)
			}
			result.PersonallyIdentifiableInformationConfig = &tmp
		}
	}

	if promptInjectionConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prompt_injection_config")); ok {
		if tmpList := promptInjectionConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "prompt_injection_config"), 0)
			tmp, err := s.mapToPromptInjectionGuardrailConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert prompt_injection_config, encountered error: %v", err)
			}
			result.PromptInjectionConfig = &tmp
		}
	}

	return result, nil
}

func GuardrailConfigToMap(obj *oci_generative_ai_agent.GuardrailConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ContentModerationConfig != nil {
		result["content_moderation_config"] = []interface{}{ContentModerationGuardrailConfigToMap(obj.ContentModerationConfig)}
	}

	if obj.PersonallyIdentifiableInformationConfig != nil {
		result["personally_identifiable_information_config"] = []interface{}{PersonallyIdentifiableInformationGuardrailConfigToMap(obj.PersonallyIdentifiableInformationConfig)}
	}

	if obj.PromptInjectionConfig != nil {
		result["prompt_injection_config"] = []interface{}{PromptInjectionGuardrailConfigToMap(obj.PromptInjectionConfig)}
	}

	return result
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) mapToHumanInputConfig(fieldKeyFormat string) (oci_generative_ai_agent.HumanInputConfig, error) {
	result := oci_generative_ai_agent.HumanInputConfig{}

	if shouldEnableHumanInput, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_enable_human_input")); ok {
		tmp := shouldEnableHumanInput.(bool)
		result.ShouldEnableHumanInput = &tmp
	}

	return result, nil
}

func HumanInputConfigToMap(obj *oci_generative_ai_agent.HumanInputConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ShouldEnableHumanInput != nil {
		result["should_enable_human_input"] = bool(*obj.ShouldEnableHumanInput)
	}

	return result
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) mapToOutputConfig(fieldKeyFormat string) (oci_generative_ai_agent.OutputConfig, error) {
	result := oci_generative_ai_agent.OutputConfig{}

	if outputLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_location")); ok {
		if tmpList := outputLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "output_location"), 0)
			tmp, err := s.mapToOutputLocation(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert output_location, encountered error: %v", err)
			}
			result.OutputLocation = tmp
		}
	}

	if retentionPeriodInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_period_in_minutes")); ok {
		tmp := retentionPeriodInMinutes.(int)
		result.RetentionPeriodInMinutes = &tmp
	}

	return result, nil
}

func OutputConfigToMap(obj *oci_generative_ai_agent.OutputConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.OutputLocation != nil {
		outputLocationArray := []interface{}{}
		if outputLocationMap := OutputLocationToMap(&obj.OutputLocation); outputLocationMap != nil {
			outputLocationArray = append(outputLocationArray, outputLocationMap)
		}
		result["output_location"] = outputLocationArray
	}

	if obj.RetentionPeriodInMinutes != nil {
		result["retention_period_in_minutes"] = int(*obj.RetentionPeriodInMinutes)
	}

	return result
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) mapToOutputLocation(fieldKeyFormat string) (oci_generative_ai_agent.OutputLocation, error) {
	var baseObject oci_generative_ai_agent.OutputLocation
	//discriminator
	outputLocationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_location_type"))
	var outputLocationType string
	if ok {
		outputLocationType = outputLocationTypeRaw.(string)
	} else {
		outputLocationType = "OBJECT_STORAGE_PREFIX" // default value
	}
	switch strings.ToLower(outputLocationType) {
	case strings.ToLower("OBJECT_STORAGE_PREFIX"):
		details := oci_generative_ai_agent.ObjectStoragePrefixOutputLocation{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
			tmp := prefix.(string)
			details.Prefix = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown output_location_type '%v' was specified", outputLocationType)
	}
	return baseObject, nil
}

func OutputLocationToMap(obj *oci_generative_ai_agent.OutputLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai_agent.ObjectStoragePrefixOutputLocation:
		result["output_location_type"] = "OBJECT_STORAGE_PREFIX"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		if v.Prefix != nil {
			result["prefix"] = string(*v.Prefix)
		}
	default:
		log.Printf("[WARN] Received 'output_location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) mapToPersonallyIdentifiableInformationGuardrailConfig(fieldKeyFormat string) (oci_generative_ai_agent.PersonallyIdentifiableInformationGuardrailConfig, error) {
	result := oci_generative_ai_agent.PersonallyIdentifiableInformationGuardrailConfig{}

	if inputGuardrailMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "input_guardrail_mode")); ok {
		result.InputGuardrailMode = oci_generative_ai_agent.GuardrailModeEnum(inputGuardrailMode.(string))
	}

	if outputGuardrailMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_guardrail_mode")); ok {
		result.OutputGuardrailMode = oci_generative_ai_agent.GuardrailModeEnum(outputGuardrailMode.(string))
	}

	return result, nil
}

func PersonallyIdentifiableInformationGuardrailConfigToMap(obj *oci_generative_ai_agent.PersonallyIdentifiableInformationGuardrailConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["input_guardrail_mode"] = string(obj.InputGuardrailMode)

	result["output_guardrail_mode"] = string(obj.OutputGuardrailMode)

	return result
}

func (s *GenerativeAiAgentAgentEndpointResourceCrud) mapToPromptInjectionGuardrailConfig(fieldKeyFormat string) (oci_generative_ai_agent.PromptInjectionGuardrailConfig, error) {
	result := oci_generative_ai_agent.PromptInjectionGuardrailConfig{}

	if inputGuardrailMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "input_guardrail_mode")); ok {
		result.InputGuardrailMode = oci_generative_ai_agent.GuardrailModeEnum(inputGuardrailMode.(string))
	}

	return result, nil
}

func PromptInjectionGuardrailConfigToMap(obj *oci_generative_ai_agent.PromptInjectionGuardrailConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["input_guardrail_mode"] = string(obj.InputGuardrailMode)

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
