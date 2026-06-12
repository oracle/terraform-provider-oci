// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

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
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiProjectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createGenerativeAiProjectWithContext,
		ReadContext:   readGenerativeAiProjectWithContext,
		UpdateContext: updateGenerativeAiProjectWithContext,
		DeleteContext: deleteGenerativeAiProjectWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"conversation_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"conversations_retention_in_hours": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"responses_retention_in_hours": {
							Type:     schema.TypeInt,
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
			"long_term_memory_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"standard_long_term_memory_strategy": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"is_enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},

									// Optional
									"embedding_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"llm_selection": {
													Type:     schema.TypeList,
													Required: true,
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
																	"GEN_AI_MODEL",
																}, true),
															},
															"model_id": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},

												// Optional

												// Computed
											},
										},
									},
									"extraction_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"llm_selection": {
													Type:     schema.TypeList,
													Required: true,
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
																	"GEN_AI_MODEL",
																}, true),
															},
															"model_id": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},

												// Optional

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"short_term_memory_optimization_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"is_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional
						"condenser_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"llm_selection": {
										Type:     schema.TypeList,
										Required: true,
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
														"GEN_AI_MODEL",
													}, true),
												},
												"model_id": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional

												// Computed
											},
										},
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
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

func createGenerativeAiProjectWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readGenerativeAiProjectWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateGenerativeAiProjectWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteGenerativeAiProjectWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type GenerativeAiProjectResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai.GenerativeAiClient
	Res                    *oci_generative_ai.GenerativeAiProject
	DisableNotFoundRetries bool
}

func (s *GenerativeAiProjectResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiProjectResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai.GenerativeAiProjectLifecycleStateCreating),
	}
}

func (s *GenerativeAiProjectResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai.GenerativeAiProjectLifecycleStateActive),
	}
}

func (s *GenerativeAiProjectResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai.GenerativeAiProjectLifecycleStateDeleting),
	}
}

func (s *GenerativeAiProjectResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai.GenerativeAiProjectLifecycleStateDeleted),
	}
}

func (s *GenerativeAiProjectResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_generative_ai.CreateGenerativeAiProjectRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if conversationConfig, ok := s.D.GetOkExists("conversation_config"); ok {
		if tmpList := conversationConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "conversation_config", 0)
			tmp, err := s.mapToConversationConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConversationConfig = &tmp
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

	if longTermMemoryConfig, ok := s.D.GetOkExists("long_term_memory_config"); ok {
		if tmpList := longTermMemoryConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "long_term_memory_config", 0)
			tmp, err := s.mapToLongTermMemoryConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LongTermMemoryConfig = &tmp
		}
	}

	if shortTermMemoryOptimizationConfig, ok := s.D.GetOkExists("short_term_memory_optimization_config"); ok {
		if tmpList := shortTermMemoryOptimizationConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "short_term_memory_optimization_config", 0)
			tmp, err := s.mapToShortTermMemoryOptimizationConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ShortTermMemoryOptimizationConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.CreateGenerativeAiProject(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getProjectFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiProjectResourceCrud) getProjectFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	projectId, err := projectWaitForWorkRequest(ctx, workId, "generativeaiproject",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*projectId)

	return s.GetWithContext(ctx)
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
		if tfresource.ShouldRetry(response, false, "generative_ai", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_generative_ai.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func projectWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_generative_ai.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai.GenerativeAiClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai")
	retryPolicy.ShouldRetryOperation = projectWorkRequestShouldRetryFunc(timeout)

	response := oci_generative_ai.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_generative_ai.OperationStatusInProgress),
			string(oci_generative_ai.OperationStatusAccepted),
			string(oci_generative_ai.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_generative_ai.OperationStatusSucceeded),
			string(oci_generative_ai.OperationStatusFailed),
			string(oci_generative_ai.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_generative_ai.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_generative_ai.OperationStatusFailed || response.Status == oci_generative_ai.OperationStatusCanceled {
		return nil, getErrorFromGenerativeAiProjectWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiProjectWorkRequest(ctx context.Context, client *oci_generative_ai.GenerativeAiClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_generative_ai.ListWorkRequestErrorsRequest{
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

func (s *GenerativeAiProjectResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.GetGenerativeAiProjectRequest{}

	tmp := s.D.Id()
	request.GenerativeAiProjectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.GetGenerativeAiProject(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.GenerativeAiProject
	return nil
}

func (s *GenerativeAiProjectResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai.UpdateGenerativeAiProjectRequest{}

	if conversationConfig, ok := s.D.GetOkExists("conversation_config"); ok {
		if tmpList := conversationConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "conversation_config", 0)
			tmp, err := s.mapToConversationConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConversationConfig = &tmp
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

	if longTermMemoryConfig, ok := s.D.GetOkExists("long_term_memory_config"); ok {
		if tmpList := longTermMemoryConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "long_term_memory_config", 0)
			tmp, err := s.mapToLongTermMemoryConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LongTermMemoryConfig = &tmp
		}
	}

	tmp := s.D.Id()
	request.GenerativeAiProjectId = &tmp

	if shortTermMemoryOptimizationConfig, ok := s.D.GetOkExists("short_term_memory_optimization_config"); ok {
		if tmpList := shortTermMemoryOptimizationConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "short_term_memory_optimization_config", 0)
			tmp, err := s.mapToShortTermMemoryOptimizationConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ShortTermMemoryOptimizationConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.UpdateGenerativeAiProject(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getProjectFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiProjectResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_generative_ai.DeleteGenerativeAiProjectRequest{}

	tmp := s.D.Id()
	request.GenerativeAiProjectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.DeleteGenerativeAiProject(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := projectWaitForWorkRequest(ctx, workId, "generativeaiproject",
		oci_generative_ai.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiProjectResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConversationConfig != nil {
		s.D.Set("conversation_config", []interface{}{ConversationConfigToMap(s.Res.ConversationConfig)})
	} else {
		s.D.Set("conversation_config", nil)
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

	if s.Res.LongTermMemoryConfig != nil {
		s.D.Set("long_term_memory_config", []interface{}{LongTermMemoryConfigToMap(s.Res.LongTermMemoryConfig)})
	} else {
		s.D.Set("long_term_memory_config", nil)
	}

	if s.Res.ShortTermMemoryOptimizationConfig != nil {
		s.D.Set("short_term_memory_optimization_config", []interface{}{ShortTermMemoryOptimizationConfigToMap(s.Res.ShortTermMemoryOptimizationConfig)})
	} else {
		s.D.Set("short_term_memory_optimization_config", nil)
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

func (s *GenerativeAiProjectResourceCrud) mapToCondenserConfig(fieldKeyFormat string) (oci_generative_ai.CondenserConfig, error) {
	result := oci_generative_ai.CondenserConfig{}

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

func CondenserConfigToMap(obj *oci_generative_ai.CondenserConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LlmSelection != nil {
		llmSelectionArray := []interface{}{}
		if llmSelectionMap := LlmSelectionToMap(&obj.LlmSelection); llmSelectionMap != nil {
			llmSelectionArray = append(llmSelectionArray, llmSelectionMap)
		}
		result["llm_selection"] = llmSelectionArray
	}

	return result
}

func (s *GenerativeAiProjectResourceCrud) mapToConversationConfig(fieldKeyFormat string) (oci_generative_ai.ConversationConfig, error) {
	result := oci_generative_ai.ConversationConfig{}

	if conversationsRetentionInHours, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "conversations_retention_in_hours")); ok {
		tmp := conversationsRetentionInHours.(int)
		result.ConversationsRetentionInHours = &tmp
	}

	if responsesRetentionInHours, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "responses_retention_in_hours")); ok {
		tmp := responsesRetentionInHours.(int)
		result.ResponsesRetentionInHours = &tmp
	}

	return result, nil
}

func ConversationConfigToMap(obj *oci_generative_ai.ConversationConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConversationsRetentionInHours != nil {
		result["conversations_retention_in_hours"] = int(*obj.ConversationsRetentionInHours)
	}

	if obj.ResponsesRetentionInHours != nil {
		result["responses_retention_in_hours"] = int(*obj.ResponsesRetentionInHours)
	}

	return result
}

func (s *GenerativeAiProjectResourceCrud) mapToEmbeddingConfig(fieldKeyFormat string) (oci_generative_ai.EmbeddingConfig, error) {
	result := oci_generative_ai.EmbeddingConfig{}

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

func EmbeddingConfigToMap(obj *oci_generative_ai.EmbeddingConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LlmSelection != nil {
		llmSelectionArray := []interface{}{}
		if llmSelectionMap := LlmSelectionToMap(&obj.LlmSelection); llmSelectionMap != nil {
			llmSelectionArray = append(llmSelectionArray, llmSelectionMap)
		}
		result["llm_selection"] = llmSelectionArray
	}

	return result
}

func (s *GenerativeAiProjectResourceCrud) mapToExtractionConfig(fieldKeyFormat string) (oci_generative_ai.ExtractionConfig, error) {
	result := oci_generative_ai.ExtractionConfig{}

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

func ExtractionConfigToMap(obj *oci_generative_ai.ExtractionConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LlmSelection != nil {
		llmSelectionArray := []interface{}{}
		if llmSelectionMap := LlmSelectionToMap(&obj.LlmSelection); llmSelectionMap != nil {
			llmSelectionArray = append(llmSelectionArray, llmSelectionMap)
		}
		result["llm_selection"] = llmSelectionArray
	}

	return result
}

func GenerativeAiProjectSummaryToMap(obj oci_generative_ai.GenerativeAiProjectSummary) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func (s *GenerativeAiProjectResourceCrud) mapToLlmSelection(fieldKeyFormat string) (oci_generative_ai.LlmSelection, error) {
	var baseObject oci_generative_ai.LlmSelection
	//discriminator
	llmSelectionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "llm_selection_type"))
	var llmSelectionType string
	if ok {
		llmSelectionType = llmSelectionTypeRaw.(string)
	} else {
		llmSelectionType = "" // default value
	}
	switch strings.ToLower(llmSelectionType) {
	case strings.ToLower("GEN_AI_MODEL"):
		details := oci_generative_ai.GenAiModelLlmSelection{}
		if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
			tmp := modelId.(string)
			details.ModelId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown llm_selection_type '%v' was specified", llmSelectionType)
	}
	return baseObject, nil
}

func LlmSelectionToMap(obj *oci_generative_ai.LlmSelection) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai.GenAiModelLlmSelection:
		result["llm_selection_type"] = "GEN_AI_MODEL"

		if v.ModelId != nil {
			result["model_id"] = string(*v.ModelId)
		}
	default:
		log.Printf("[WARN] Received 'llm_selection_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiProjectResourceCrud) mapToLongTermMemoryConfig(fieldKeyFormat string) (oci_generative_ai.LongTermMemoryConfig, error) {
	result := oci_generative_ai.LongTermMemoryConfig{}

	if standardLongTermMemoryStrategy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "standard_long_term_memory_strategy")); ok {
		if tmpList := standardLongTermMemoryStrategy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "standard_long_term_memory_strategy"), 0)
			tmp, err := s.mapToStandardLongTermMemoryStrategy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert standard_long_term_memory_strategy, encountered error: %v", err)
			}
			result.StandardLongTermMemoryStrategy = &tmp
		}
	}

	return result, nil
}

func LongTermMemoryConfigToMap(obj *oci_generative_ai.LongTermMemoryConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.StandardLongTermMemoryStrategy != nil {
		result["standard_long_term_memory_strategy"] = []interface{}{StandardLongTermMemoryStrategyToMap(obj.StandardLongTermMemoryStrategy)}
	}

	return result
}

func (s *GenerativeAiProjectResourceCrud) mapToShortTermMemoryOptimizationConfig(fieldKeyFormat string) (oci_generative_ai.ShortTermMemoryOptimizationConfig, error) {
	result := oci_generative_ai.ShortTermMemoryOptimizationConfig{}

	if condenserConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condenser_config")); ok {
		if tmpList := condenserConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "condenser_config"), 0)
			tmp, err := s.mapToCondenserConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert condenser_config, encountered error: %v", err)
			}
			result.CondenserConfig = &tmp
		}
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func ShortTermMemoryOptimizationConfigToMap(obj *oci_generative_ai.ShortTermMemoryOptimizationConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CondenserConfig != nil {
		result["condenser_config"] = []interface{}{CondenserConfigToMap(obj.CondenserConfig)}
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *GenerativeAiProjectResourceCrud) mapToStandardLongTermMemoryStrategy(fieldKeyFormat string) (oci_generative_ai.StandardLongTermMemoryStrategy, error) {
	result := oci_generative_ai.StandardLongTermMemoryStrategy{}

	if embeddingConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "embedding_config")); ok {
		if tmpList := embeddingConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "embedding_config"), 0)
			tmp, err := s.mapToEmbeddingConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert embedding_config, encountered error: %v", err)
			}
			result.EmbeddingConfig = &tmp
		}
	}

	if extractionConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extraction_config")); ok {
		if tmpList := extractionConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "extraction_config"), 0)
			tmp, err := s.mapToExtractionConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert extraction_config, encountered error: %v", err)
			}
			result.ExtractionConfig = &tmp
		}
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func StandardLongTermMemoryStrategyToMap(obj *oci_generative_ai.StandardLongTermMemoryStrategy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EmbeddingConfig != nil {
		result["embedding_config"] = []interface{}{EmbeddingConfigToMap(obj.EmbeddingConfig)}
	}

	if obj.ExtractionConfig != nil {
		result["extraction_config"] = []interface{}{ExtractionConfigToMap(obj.ExtractionConfig)}
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *GenerativeAiProjectResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai.ChangeGenerativeAiProjectCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.GenerativeAiProjectId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	_, err := s.Client.ChangeGenerativeAiProjectCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
