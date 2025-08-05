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

func GenerativeAiAgentToolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createGenerativeAiAgentTool,
		Read:     readGenerativeAiAgentTool,
		Update:   updateGenerativeAiAgentTool,
		Delete:   deleteGenerativeAiAgentTool,
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
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tool_config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"tool_config_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AGENT_TOOL_CONFIG",
								"FUNCTION_CALLING_TOOL_CONFIG",
								"HTTP_ENDPOINT_TOOL_CONFIG",
								"RAG_TOOL_CONFIG",
								"SQL_TOOL_CONFIG",
							}, true),
						},

						// Optional
						"agent_endpoint_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"api_schema": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"api_schema_input_location_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"INLINE",
											"OBJECT_STORAGE_LOCATION",
										}, true),
									},

									// Optional
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"content": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"object": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"database_connection": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"connection_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"connection_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DATABASE_TOOL_CONNECTION",
										}, true),
									},

									// Optional

									// Computed
								},
							},
						},
						"database_schema": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"input_location_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"INLINE",
											"OBJECT_STORAGE_PREFIX",
										}, true),
									},

									// Optional
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"content": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"prefix": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"dialect": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"function": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"parameters": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},

									// Computed
								},
							},
						},
						"generation_llm_customization": {
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

									// Computed
								},
							},
						},
						"http_endpoint_auth_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"http_endpoint_auth_sources": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"http_endpoint_auth_scope": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"http_endpoint_auth_scope_config": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"http_endpoint_auth_scope_config_type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG",
																	"HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG",
																	"HTTP_ENDPOINT_BEARER_AUTH_SCOPE_CONFIG",
																	"HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG",
																	"HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG",
																	"HTTP_ENDPOINT_OCI_AUTH_SCOPE_CONFIG",
																}, true),
															},

															// Optional
															"client_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"idcs_url": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"key_location": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"key_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"scope_url": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"vault_secret_id": {
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

									// Computed
								},
							},
						},
						"icl_examples": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"input_location_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"INLINE",
											"OBJECT_STORAGE_PREFIX",
										}, true),
									},

									// Optional
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"content": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"prefix": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"knowledge_base_configs": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"knowledge_base_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"model_size": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"should_enable_self_correction": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"should_enable_sql_execution": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"table_and_column_description": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"input_location_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"INLINE",
											"OBJECT_STORAGE_PREFIX",
										}, true),
									},

									// Optional
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"content": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"prefix": {
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

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
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
			"metadata": {
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

func createGenerativeAiAgentTool(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentToolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.CreateResource(d, sync)
}

func readGenerativeAiAgentTool(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentToolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

func updateGenerativeAiAgentTool(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentToolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGenerativeAiAgentTool(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentToolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GenerativeAiAgentToolResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai_agent.GenerativeAiAgentClient
	Res                    *oci_generative_ai_agent.Tool
	DisableNotFoundRetries bool
}

func (s *GenerativeAiAgentToolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiAgentToolResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai_agent.ToolLifecycleStateCreating),
	}
}

func (s *GenerativeAiAgentToolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.ToolLifecycleStateActive),
	}
}

func (s *GenerativeAiAgentToolResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai_agent.ToolLifecycleStateDeleting),
	}
}

func (s *GenerativeAiAgentToolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.ToolLifecycleStateDeleted),
	}
}

func (s *GenerativeAiAgentToolResourceCrud) Create() error {
	request := oci_generative_ai_agent.CreateToolRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

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

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	if toolConfig, ok := s.D.GetOkExists("tool_config"); ok {
		if tmpList := toolConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tool_config", 0)
			tmp, err := s.mapToToolConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ToolConfig = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.CreateTool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getToolFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiAgentToolResourceCrud) getToolFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai_agent.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	toolId, err := toolWaitForWorkRequest(workId, "tool",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, toolId)
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
	s.D.SetId(*toolId)

	return s.Get()
}

func toolWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func toolWaitForWorkRequest(wId *string, entityType string, action oci_generative_ai_agent.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai_agent.GenerativeAiAgentClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai_agent")
	retryPolicy.ShouldRetryOperation = toolWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromGenerativeAiAgentToolWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiAgentToolWorkRequest(client *oci_generative_ai_agent.GenerativeAiAgentClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai_agent.ActionTypeEnum) error {
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

func (s *GenerativeAiAgentToolResourceCrud) Get() error {
	request := oci_generative_ai_agent.GetToolRequest{}

	tmp := s.D.Id()
	request.ToolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.GetTool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Tool
	return nil
}

func (s *GenerativeAiAgentToolResourceCrud) Update() error {
	request := oci_generative_ai_agent.UpdateToolRequest{}

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

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	if toolConfig, ok := s.D.GetOkExists("tool_config"); ok {
		if tmpList := toolConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tool_config", 0)
			tmp, err := s.mapToToolConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ToolConfig = tmp
		}
	}

	tmp := s.D.Id()
	request.ToolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.UpdateTool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getToolFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiAgentToolResourceCrud) Delete() error {
	request := oci_generative_ai_agent.DeleteToolRequest{}

	tmp := s.D.Id()
	request.ToolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.DeleteTool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := toolWaitForWorkRequest(workId, "tool",
		oci_generative_ai_agent.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiAgentToolResourceCrud) SetData() error {
	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
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

	s.D.Set("metadata", s.Res.Metadata)

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

	if s.Res.ToolConfig != nil {
		toolConfigArray := []interface{}{}
		if toolConfigMap := ToolConfigToMap(&s.Res.ToolConfig); toolConfigMap != nil {
			toolConfigArray = append(toolConfigArray, toolConfigMap)
		}
		s.D.Set("tool_config", toolConfigArray)
	} else {
		s.D.Set("tool_config", nil)
	}

	return nil
}

func (s *GenerativeAiAgentToolResourceCrud) mapToApiSchemaInputLocation(fieldKeyFormat string) (oci_generative_ai_agent.ApiSchemaInputLocation, error) {
	var baseObject oci_generative_ai_agent.ApiSchemaInputLocation
	//discriminator
	apiSchemaInputLocationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "api_schema_input_location_type"))
	var apiSchemaInputLocationType string
	if ok {
		apiSchemaInputLocationType = apiSchemaInputLocationTypeRaw.(string)
	} else {
		apiSchemaInputLocationType = "" // default value
	}
	switch strings.ToLower(apiSchemaInputLocationType) {
	case strings.ToLower("INLINE"):
		details := oci_generative_ai_agent.ApiSchemaInlineInputLocation{}
		if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
			tmp := content.(string)
			details.Content = &tmp
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE_LOCATION"):
		details := oci_generative_ai_agent.ApiSchemaObjectStorageInputLocation{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := object.(string)
			details.ObjectName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown api_schema_input_location_type '%v' was specified", apiSchemaInputLocationType)
	}
	return baseObject, nil
}

func ApiSchemaInputLocationToMap(obj *oci_generative_ai_agent.ApiSchemaInputLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai_agent.ApiSchemaInlineInputLocation:
		result["api_schema_input_location_type"] = "INLINE"

		if v.Content != nil {
			result["content"] = string(*v.Content)
		}
	case oci_generative_ai_agent.ApiSchemaObjectStorageInputLocation:
		result["api_schema_input_location_type"] = "OBJECT_STORAGE_LOCATION"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}
	default:
		log.Printf("[WARN] Received 'api_schema_input_location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiAgentToolResourceCrud) mapToDatabaseConnection(fieldKeyFormat string) (oci_generative_ai_agent.DatabaseConnection, error) {
	var baseObject oci_generative_ai_agent.DatabaseConnection
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_type"))
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("DATABASE_TOOL_CONNECTION"):
		details := oci_generative_ai_agent.DatabaseToolConnection{}
		if connectionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_id")); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return baseObject, nil
}

// We should use the same function as the KB one or keep them separate
func GenAiAgentToolsDatabaseConnectionToMap(obj *oci_generative_ai_agent.DatabaseConnection) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai_agent.DatabaseToolConnection:
		result["connection_type"] = "DATABASE_TOOL_CONNECTION"

		if v.ConnectionId != nil {
			result["connection_id"] = string(*v.ConnectionId)
		}
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiAgentToolResourceCrud) mapToFunction(fieldKeyFormat string) (oci_generative_ai_agent.Function, error) {
	result := oci_generative_ai_agent.Function{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if parameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parameters")); ok {
		result.Parameters = tfresource.ObjectMapToStringMap(parameters.(map[string]interface{}))
	}

	return result, nil
}

func FunctionToMap(obj *oci_generative_ai_agent.Function) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["parameters"] = obj.Parameters

	return result
}

func (s *GenerativeAiAgentToolResourceCrud) mapToHttpEndpointAuthConfig(fieldKeyFormat string) (oci_generative_ai_agent.HttpEndpointAuthConfig, error) {
	result := oci_generative_ai_agent.HttpEndpointAuthConfig{}

	if httpEndpointAuthSources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "http_endpoint_auth_sources")); ok {
		interfaces := httpEndpointAuthSources.([]interface{})
		tmp := make([]oci_generative_ai_agent.HttpEndpointAuthSource, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "http_endpoint_auth_sources"), stateDataIndex)
			converted, err := s.mapToHttpEndpointAuthSource(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "http_endpoint_auth_sources")) {
			result.HttpEndpointAuthSources = tmp
		}
	}

	return result, nil
}

func HttpEndpointAuthConfigToMap(obj *oci_generative_ai_agent.HttpEndpointAuthConfig) map[string]interface{} {
	result := map[string]interface{}{}

	httpEndpointAuthSources := []interface{}{}
	for _, item := range obj.HttpEndpointAuthSources {
		httpEndpointAuthSources = append(httpEndpointAuthSources, HttpEndpointAuthSourceToMap(item))
	}
	result["http_endpoint_auth_sources"] = httpEndpointAuthSources

	return result
}

func (s *GenerativeAiAgentToolResourceCrud) mapToHttpEndpointAuthScopeConfig(fieldKeyFormat string) (oci_generative_ai_agent.HttpEndpointAuthScopeConfig, error) {
	var baseObject oci_generative_ai_agent.HttpEndpointAuthScopeConfig
	//discriminator
	httpEndpointAuthScopeConfigTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "http_endpoint_auth_scope_config_type"))
	var httpEndpointAuthScopeConfigType string
	if ok {
		httpEndpointAuthScopeConfigType = httpEndpointAuthScopeConfigTypeRaw.(string)
	} else {
		httpEndpointAuthScopeConfigType = "" // default value
	}
	switch strings.ToLower(httpEndpointAuthScopeConfigType) {
	case strings.ToLower("HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG"):
		details := oci_generative_ai_agent.HttpEndpointApiKeyAuthScopeConfig{}
		if keyLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_location")); ok {
			details.KeyLocation = oci_generative_ai_agent.HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum(keyLocation.(string))
		}
		if keyName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_name")); ok {
			tmp := keyName.(string)
			details.KeyName = &tmp
		}
		if vaultSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_secret_id")); ok {
			tmp := vaultSecretId.(string)
			details.VaultSecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG"):
		details := oci_generative_ai_agent.HttpEndpointBasicAuthScopeConfig{}
		if vaultSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_secret_id")); ok {
			tmp := vaultSecretId.(string)
			details.VaultSecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("HTTP_ENDPOINT_BEARER_AUTH_SCOPE_CONFIG"):
		details := oci_generative_ai_agent.HttpEndpointBearerAuthScopeConfig{}
		if vaultSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_secret_id")); ok {
			tmp := vaultSecretId.(string)
			details.VaultSecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG"):
		details := oci_generative_ai_agent.HttpEndpointIdcsAuthScopeConfig{}
		if clientId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_id")); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if idcsUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idcs_url")); ok {
			tmp := idcsUrl.(string)
			details.IdcsUrl = &tmp
		}
		if scopeUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope_url")); ok {
			tmp := scopeUrl.(string)
			details.ScopeUrl = &tmp
		}
		if vaultSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_secret_id")); ok {
			tmp := vaultSecretId.(string)
			details.VaultSecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG"):
		details := oci_generative_ai_agent.HttpEndpointNoAuthScopeConfig{}
		baseObject = details
	case strings.ToLower("HTTP_ENDPOINT_OCI_AUTH_SCOPE_CONFIG"):
		details := oci_generative_ai_agent.HttpEndpointOciAuthScopeConfig{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown http_endpoint_auth_scope_config_type '%v' was specified", httpEndpointAuthScopeConfigType)
	}
	return baseObject, nil
}

func HttpEndpointAuthScopeConfigToMap(obj *oci_generative_ai_agent.HttpEndpointAuthScopeConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai_agent.HttpEndpointApiKeyAuthScopeConfig:
		result["http_endpoint_auth_scope_config_type"] = "HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG"

		result["key_location"] = string(v.KeyLocation)

		if v.KeyName != nil {
			result["key_name"] = string(*v.KeyName)
		}

		if v.VaultSecretId != nil {
			result["vault_secret_id"] = string(*v.VaultSecretId)
		}
	case oci_generative_ai_agent.HttpEndpointBasicAuthScopeConfig:
		result["http_endpoint_auth_scope_config_type"] = "HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG"

		if v.VaultSecretId != nil {
			result["vault_secret_id"] = string(*v.VaultSecretId)
		}
	case oci_generative_ai_agent.HttpEndpointBearerAuthScopeConfig:
		result["http_endpoint_auth_scope_config_type"] = "HTTP_ENDPOINT_BEARER_AUTH_SCOPE_CONFIG"

		if v.VaultSecretId != nil {
			result["vault_secret_id"] = string(*v.VaultSecretId)
		}
	case oci_generative_ai_agent.HttpEndpointIdcsAuthScopeConfig:
		result["http_endpoint_auth_scope_config_type"] = "HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG"

		if v.ClientId != nil {
			result["client_id"] = string(*v.ClientId)
		}

		if v.IdcsUrl != nil {
			result["idcs_url"] = string(*v.IdcsUrl)
		}

		if v.ScopeUrl != nil {
			result["scope_url"] = string(*v.ScopeUrl)
		}

		if v.VaultSecretId != nil {
			result["vault_secret_id"] = string(*v.VaultSecretId)
		}
	case oci_generative_ai_agent.HttpEndpointNoAuthScopeConfig:
		result["http_endpoint_auth_scope_config_type"] = "HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG"
	case oci_generative_ai_agent.HttpEndpointOciAuthScopeConfig:
		result["http_endpoint_auth_scope_config_type"] = "HTTP_ENDPOINT_OCI_AUTH_SCOPE_CONFIG"
	default:
		log.Printf("[WARN] Received 'http_endpoint_auth_scope_config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiAgentToolResourceCrud) mapToHttpEndpointAuthSource(fieldKeyFormat string) (oci_generative_ai_agent.HttpEndpointAuthSource, error) {
	result := oci_generative_ai_agent.HttpEndpointAuthSource{}

	if httpEndpointAuthScope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "http_endpoint_auth_scope")); ok {
		result.HttpEndpointAuthScope = oci_generative_ai_agent.HttpEndpointAuthSourceHttpEndpointAuthScopeEnum(httpEndpointAuthScope.(string))
	}

	if httpEndpointAuthScopeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "http_endpoint_auth_scope_config")); ok {
		if tmpList := httpEndpointAuthScopeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "http_endpoint_auth_scope_config"), 0)
			tmp, err := s.mapToHttpEndpointAuthScopeConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert http_endpoint_auth_scope_config, encountered error: %v", err)
			}
			result.HttpEndpointAuthScopeConfig = tmp
		}
	}

	return result, nil
}

func HttpEndpointAuthSourceToMap(obj oci_generative_ai_agent.HttpEndpointAuthSource) map[string]interface{} {
	result := map[string]interface{}{}

	result["http_endpoint_auth_scope"] = string(obj.HttpEndpointAuthScope)

	if obj.HttpEndpointAuthScopeConfig != nil {
		httpEndpointAuthScopeConfigArray := []interface{}{}
		if httpEndpointAuthScopeConfigMap := HttpEndpointAuthScopeConfigToMap(&obj.HttpEndpointAuthScopeConfig); httpEndpointAuthScopeConfigMap != nil {
			httpEndpointAuthScopeConfigArray = append(httpEndpointAuthScopeConfigArray, httpEndpointAuthScopeConfigMap)
		}
		result["http_endpoint_auth_scope_config"] = httpEndpointAuthScopeConfigArray
	}

	return result
}

func (s *GenerativeAiAgentToolResourceCrud) mapToInputLocation(fieldKeyFormat string) (oci_generative_ai_agent.InputLocation, error) {
	var baseObject oci_generative_ai_agent.InputLocation
	//discriminator
	inputLocationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "input_location_type"))
	var inputLocationType string
	if ok {
		inputLocationType = inputLocationTypeRaw.(string)
	} else {
		inputLocationType = "" // default value
	}
	switch strings.ToLower(inputLocationType) {
	case strings.ToLower("INLINE"):
		details := oci_generative_ai_agent.InlineInputLocation{}
		if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
			tmp := content.(string)
			details.Content = &tmp
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE_PREFIX"):
		details := oci_generative_ai_agent.ObjectStorageInputLocation{}
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
		return nil, fmt.Errorf("unknown input_location_type '%v' was specified", inputLocationType)
	}
	return baseObject, nil
}

func InputLocationToMap(obj *oci_generative_ai_agent.InputLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai_agent.InlineInputLocation:
		result["input_location_type"] = "INLINE"

		if v.Content != nil {
			result["content"] = string(*v.Content)
		}
	case oci_generative_ai_agent.ObjectStorageInputLocation:
		result["input_location_type"] = "OBJECT_STORAGE_PREFIX"

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
		log.Printf("[WARN] Received 'input_location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiAgentToolResourceCrud) mapToKnowledgeBaseConfig(fieldKeyFormat string) (oci_generative_ai_agent.KnowledgeBaseConfig, error) {
	result := oci_generative_ai_agent.KnowledgeBaseConfig{}

	if knowledgeBaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "knowledge_base_id")); ok {
		tmp := knowledgeBaseId.(string)
		result.KnowledgeBaseId = &tmp
	}

	return result, nil
}

func KnowledgeBaseConfigToMap(obj oci_generative_ai_agent.KnowledgeBaseConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KnowledgeBaseId != nil {
		result["knowledge_base_id"] = string(*obj.KnowledgeBaseId)
	}

	return result
}

func (s *GenerativeAiAgentToolResourceCrud) mapToLlmCustomization(fieldKeyFormat string) (oci_generative_ai_agent.LlmCustomization, error) {
	result := oci_generative_ai_agent.LlmCustomization{}

	if instruction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instruction")); ok {
		tmp := instruction.(string)
		result.Instruction = &tmp
	}

	return result, nil
}

func (s *GenerativeAiAgentToolResourceCrud) mapToToolConfig(fieldKeyFormat string) (oci_generative_ai_agent.ToolConfig, error) {
	var baseObject oci_generative_ai_agent.ToolConfig
	//discriminator
	toolConfigTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tool_config_type"))
	var toolConfigType string
	if ok {
		toolConfigType = toolConfigTypeRaw.(string)
	} else {
		toolConfigType = "" // default value
	}
	switch strings.ToLower(toolConfigType) {
	case strings.ToLower("AGENT_TOOL_CONFIG"):
		details := oci_generative_ai_agent.AgentToolConfig{}
		if agentEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "agent_endpoint_id")); ok {
			tmp := agentEndpointId.(string)
			details.AgentEndpointId = &tmp
		}
		baseObject = details
	case strings.ToLower("FUNCTION_CALLING_TOOL_CONFIG"):
		details := oci_generative_ai_agent.FunctionCallingToolConfig{}
		if function, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "function")); ok {
			if tmpList := function.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "function"), 0)
				tmp, err := s.mapToFunction(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert function, encountered error: %v", err)
				}
				details.Function = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("HTTP_ENDPOINT_TOOL_CONFIG"):
		details := oci_generative_ai_agent.HttpEndpointToolConfig{}
		if apiSchema, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "api_schema")); ok {
			if tmpList := apiSchema.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "api_schema"), 0)
				tmp, err := s.mapToApiSchemaInputLocation(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert api_schema, encountered error: %v", err)
				}
				details.ApiSchema = tmp
			}
		}
		if httpEndpointAuthConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "http_endpoint_auth_config")); ok {
			if tmpList := httpEndpointAuthConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "http_endpoint_auth_config"), 0)
				tmp, err := s.mapToHttpEndpointAuthConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert http_endpoint_auth_config, encountered error: %v", err)
				}
				details.HttpEndpointAuthConfig = &tmp
			}
		}
		if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		baseObject = details
	case strings.ToLower("RAG_TOOL_CONFIG"):
		details := oci_generative_ai_agent.RagToolConfig{}
		if generationLlmCustomization, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "generation_llm_customization")); ok {
			if tmpList := generationLlmCustomization.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "generation_llm_customization"), 0)
				tmp, err := s.mapToLlmCustomization(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert generation_llm_customization, encountered error: %v", err)
				}
				details.GenerationLlmCustomization = &tmp
			}
		}
		if knowledgeBaseConfigs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "knowledge_base_configs")); ok {
			interfaces := knowledgeBaseConfigs.([]interface{})
			tmp := make([]oci_generative_ai_agent.KnowledgeBaseConfig, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "knowledge_base_configs"), stateDataIndex)
				converted, err := s.mapToKnowledgeBaseConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "knowledge_base_configs")) {
				details.KnowledgeBaseConfigs = tmp
			}
		}
		baseObject = details
	case strings.ToLower("SQL_TOOL_CONFIG"):
		details := oci_generative_ai_agent.SqlToolConfig{}
		if databaseConnection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_connection")); ok {
			if tmpList := databaseConnection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database_connection"), 0)
				tmp, err := s.mapToDatabaseConnection(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert database_connection, encountered error: %v", err)
				}
				details.DatabaseConnection = tmp
			}
		}
		if databaseSchema, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_schema")); ok {
			if tmpList := databaseSchema.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database_schema"), 0)
				tmp, err := s.mapToInputLocation(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert database_schema, encountered error: %v", err)
				}
				details.DatabaseSchema = tmp
			}
		}
		if dialect, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dialect")); ok {
			details.Dialect = oci_generative_ai_agent.SqlToolConfigDialectEnum(dialect.(string))
		}
		if generationLlmCustomization, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "generation_llm_customization")); ok {
			if tmpList := generationLlmCustomization.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "generation_llm_customization"), 0)
				tmp, err := s.mapToLlmCustomization(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert generation_llm_customization, encountered error: %v", err)
				}
				details.GenerationLlmCustomization = &tmp
			}
		}
		if iclExamples, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "icl_examples")); ok {
			if tmpList := iclExamples.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "icl_examples"), 0)
				tmp, err := s.mapToInputLocation(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert icl_examples, encountered error: %v", err)
				}
				details.IclExamples = tmp
			}
		}
		if modelSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_size")); ok {
			details.ModelSize = oci_generative_ai_agent.SqlToolConfigModelSizeEnum(modelSize.(string))
		}
		if shouldEnableSelfCorrection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_enable_self_correction")); ok {
			tmp := shouldEnableSelfCorrection.(bool)
			details.ShouldEnableSelfCorrection = &tmp
		}
		if shouldEnableSqlExecution, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_enable_sql_execution")); ok {
			tmp := shouldEnableSqlExecution.(bool)
			details.ShouldEnableSqlExecution = &tmp
		}
		if tableAndColumnDescription, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "table_and_column_description")); ok {
			if tmpList := tableAndColumnDescription.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "table_and_column_description"), 0)
				tmp, err := s.mapToInputLocation(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert table_and_column_description, encountered error: %v", err)
				}
				details.TableAndColumnDescription = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown tool_config_type '%v' was specified", toolConfigType)
	}
	return baseObject, nil
}

func ToolConfigToMap(obj *oci_generative_ai_agent.ToolConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai_agent.AgentToolConfig:
		result["tool_config_type"] = "AGENT_TOOL_CONFIG"

		if v.AgentEndpointId != nil {
			result["agent_endpoint_id"] = string(*v.AgentEndpointId)
		}
	case oci_generative_ai_agent.FunctionCallingToolConfig:
		result["tool_config_type"] = "FUNCTION_CALLING_TOOL_CONFIG"

		if v.Function != nil {
			result["function"] = []interface{}{FunctionToMap(v.Function)}
		}
	case oci_generative_ai_agent.HttpEndpointToolConfig:
		result["tool_config_type"] = "HTTP_ENDPOINT_TOOL_CONFIG"

		if v.ApiSchema != nil {
			apiSchemaArray := []interface{}{}
			if apiSchemaMap := ApiSchemaInputLocationToMap(&v.ApiSchema); apiSchemaMap != nil {
				apiSchemaArray = append(apiSchemaArray, apiSchemaMap)
			}
			result["api_schema"] = apiSchemaArray
		}

		if v.HttpEndpointAuthConfig != nil {
			result["http_endpoint_auth_config"] = []interface{}{HttpEndpointAuthConfigToMap(v.HttpEndpointAuthConfig)}
		}

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}
	case oci_generative_ai_agent.RagToolConfig:
		result["tool_config_type"] = "RAG_TOOL_CONFIG"

		if v.GenerationLlmCustomization != nil {
			result["generation_llm_customization"] = []interface{}{LlmCustomizationToMap(v.GenerationLlmCustomization)}
		}

		knowledgeBaseConfigs := []interface{}{}
		for _, item := range v.KnowledgeBaseConfigs {
			knowledgeBaseConfigs = append(knowledgeBaseConfigs, KnowledgeBaseConfigToMap(item))
		}
		result["knowledge_base_configs"] = knowledgeBaseConfigs
	case oci_generative_ai_agent.SqlToolConfig:
		result["tool_config_type"] = "SQL_TOOL_CONFIG"

		if v.DatabaseConnection != nil {
			databaseConnectionArray := []interface{}{}
			if databaseConnectionMap := GenAiAgentToolsDatabaseConnectionToMap(&v.DatabaseConnection); databaseConnectionMap != nil {
				databaseConnectionArray = append(databaseConnectionArray, databaseConnectionMap)
			}
			result["database_connection"] = databaseConnectionArray
		}

		if v.DatabaseSchema != nil {
			databaseSchemaArray := []interface{}{}
			if databaseSchemaMap := InputLocationToMap(&v.DatabaseSchema); databaseSchemaMap != nil {
				databaseSchemaArray = append(databaseSchemaArray, databaseSchemaMap)
			}
			result["database_schema"] = databaseSchemaArray
		}

		result["dialect"] = string(v.Dialect)

		if v.GenerationLlmCustomization != nil {
			result["generation_llm_customization"] = []interface{}{LlmCustomizationToMap(v.GenerationLlmCustomization)}
		}

		if v.IclExamples != nil {
			iclExamplesArray := []interface{}{}
			if iclExamplesMap := InputLocationToMap(&v.IclExamples); iclExamplesMap != nil {
				iclExamplesArray = append(iclExamplesArray, iclExamplesMap)
			}
			result["icl_examples"] = iclExamplesArray
		}

		result["model_size"] = string(v.ModelSize)

		if v.ShouldEnableSelfCorrection != nil {
			result["should_enable_self_correction"] = bool(*v.ShouldEnableSelfCorrection)
		}

		if v.ShouldEnableSqlExecution != nil {
			result["should_enable_sql_execution"] = bool(*v.ShouldEnableSqlExecution)
		}

		if v.TableAndColumnDescription != nil {
			tableAndColumnDescriptionArray := []interface{}{}
			if tableAndColumnDescriptionMap := InputLocationToMap(&v.TableAndColumnDescription); tableAndColumnDescriptionMap != nil {
				tableAndColumnDescriptionArray = append(tableAndColumnDescriptionArray, tableAndColumnDescriptionMap)
			}
			result["table_and_column_description"] = tableAndColumnDescriptionArray
		}
	default:
		log.Printf("[WARN] Received 'tool_config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func ToolSummaryToMap(obj oci_generative_ai_agent.ToolSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
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

	result["metadata"] = obj.Metadata

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

	if obj.ToolConfig != nil {
		toolConfigArray := []interface{}{}
		if toolConfigMap := ToolConfigToMap(&obj.ToolConfig); toolConfigMap != nil {
			toolConfigArray = append(toolConfigArray, toolConfigMap)
		}
		result["tool_config"] = toolConfigArray
	}

	return result
}
