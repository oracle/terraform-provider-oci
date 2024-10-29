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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiAgentKnowledgeBaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("50m"),
			Update: tfresource.GetTimeoutDuration("20m"),
			Delete: tfresource.GetTimeoutDuration("20m"),
		},
		Create: createGenerativeAiAgentKnowledgeBase,
		Read:   readGenerativeAiAgentKnowledgeBase,
		Update: updateGenerativeAiAgentKnowledgeBase,
		Delete: deleteGenerativeAiAgentKnowledgeBase,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index_config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"index_config_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DEFAULT_INDEX_CONFIG",
								"OCI_DATABASE_CONFIG",
								"OCI_OPEN_SEARCH_INDEX_CONFIG",
							}, true),
						},

						// Optional
						"cluster_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"database_functions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"indexes": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"schema": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"body_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"embedding_body_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"title_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"url_key": {
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
						"secret_detail": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"BASIC_AUTH_SECRET",
											"IDCS_SECRET",
										}, true),
									},
									"vault_secret_id": {
										Type:     schema.TypeString,
										Required: true,
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
									"scope_url": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"should_enable_hybrid_search": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
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

func createGenerativeAiAgentKnowledgeBase(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentKnowledgeBaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.CreateResource(d, sync)
}

func readGenerativeAiAgentKnowledgeBase(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentKnowledgeBaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

func updateGenerativeAiAgentKnowledgeBase(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentKnowledgeBaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGenerativeAiAgentKnowledgeBase(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentKnowledgeBaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GenerativeAiAgentKnowledgeBaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai_agent.GenerativeAiAgentClient
	Res                    *oci_generative_ai_agent.KnowledgeBase
	DisableNotFoundRetries bool
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai_agent.KnowledgeBaseLifecycleStateCreating),
	}
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.KnowledgeBaseLifecycleStateActive),
	}
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai_agent.KnowledgeBaseLifecycleStateDeleting),
	}
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.KnowledgeBaseLifecycleStateDeleted),
	}
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) Create() error {
	request := oci_generative_ai_agent.CreateKnowledgeBaseRequest{}

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

	if indexConfig, ok := s.D.GetOkExists("index_config"); ok {
		if tmpList := indexConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "index_config", 0)
			tmp, err := s.mapToIndexConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.IndexConfig = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.CreateKnowledgeBase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getKnowledgeBaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) getKnowledgeBaseFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai_agent.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	knowledgeBaseId, err := knowledgeBaseWaitForWorkRequest(workId, "knowledgebase",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, knowledgeBaseId)
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
	s.D.SetId(*knowledgeBaseId)

	return s.Get()
}

func knowledgeBaseWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func knowledgeBaseWaitForWorkRequest(wId *string, entityType string, action oci_generative_ai_agent.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai_agent.GenerativeAiAgentClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai_agent")
	retryPolicy.ShouldRetryOperation = knowledgeBaseWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromGenerativeAiAgentKnowledgeBaseWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiAgentKnowledgeBaseWorkRequest(client *oci_generative_ai_agent.GenerativeAiAgentClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai_agent.ActionTypeEnum) error {
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

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) Get() error {
	request := oci_generative_ai_agent.GetKnowledgeBaseRequest{}

	tmp := s.D.Id()
	request.KnowledgeBaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.GetKnowledgeBase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KnowledgeBase
	return nil
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai_agent.UpdateKnowledgeBaseRequest{}

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

	if indexConfig, ok := s.D.GetOkExists("index_config"); ok {
		if tmpList := indexConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "index_config", 0)
			indexConfigTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "index_config_type"))
			var indexConfigType string
			if ok {
				indexConfigType = indexConfigTypeRaw.(string)
			} else {
				indexConfigType = "" // default value
			}
			if strings.ToLower(indexConfigType) != strings.ToLower("DEFAULT_INDEX_CONFIG") {
				tmp, err := s.mapToIndexConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.IndexConfig = tmp
			}
		}
	}

	tmp := s.D.Id()
	request.KnowledgeBaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.UpdateKnowledgeBase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getKnowledgeBaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) Delete() error {
	request := oci_generative_ai_agent.DeleteKnowledgeBaseRequest{}

	tmp := s.D.Id()
	request.KnowledgeBaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.DeleteKnowledgeBase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := knowledgeBaseWaitForWorkRequest(workId, "knowledgebase",
		oci_generative_ai_agent.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) SetData() error {
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

	if s.Res.IndexConfig != nil {
		indexConfigArray := []interface{}{}
		if indexConfigMap := IndexConfigToMap(&s.Res.IndexConfig); indexConfigMap != nil {
			indexConfigArray = append(indexConfigArray, indexConfigMap)
		}
		s.D.Set("index_config", indexConfigArray)
	} else {
		s.D.Set("index_config", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) mapToDatabaseConnection(fieldKeyFormat string) (oci_generative_ai_agent.DatabaseConnection, error) {
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

func DatabaseConnectionToMap(obj *oci_generative_ai_agent.DatabaseConnection) map[string]interface{} {
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

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) mapToDatabaseFunction(fieldKeyFormat string) (oci_generative_ai_agent.DatabaseFunction, error) {
	result := oci_generative_ai_agent.DatabaseFunction{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func DatabaseFunctionToMap(obj oci_generative_ai_agent.DatabaseFunction) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) mapToIndex(fieldKeyFormat string) (oci_generative_ai_agent.Index, error) {
	result := oci_generative_ai_agent.Index{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if schema, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema")); ok {
		if tmpList := schema.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "schema"), 0)
			tmp, err := s.mapToIndexSchema(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert schema, encountered error: %v", err)
			}
			result.Schema = &tmp
		}
	}

	return result, nil
}

func IndexToMap(obj oci_generative_ai_agent.Index) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Schema != nil {
		result["schema"] = []interface{}{IndexSchemaToMap(obj.Schema)}
	}

	return result
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) mapToIndexConfig(fieldKeyFormat string) (oci_generative_ai_agent.IndexConfig, error) {
	var baseObject oci_generative_ai_agent.IndexConfig
	//discriminator
	indexConfigTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "index_config_type"))
	var indexConfigType string
	if ok {
		indexConfigType = indexConfigTypeRaw.(string)
	} else {
		indexConfigType = "" // default value
	}
	switch strings.ToLower(indexConfigType) {
	case strings.ToLower("DEFAULT_INDEX_CONFIG"):
		details := oci_generative_ai_agent.DefaultIndexConfig{}
		if shouldEnableHybridSearch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_enable_hybrid_search")); ok {
			tmp := shouldEnableHybridSearch.(bool)
			details.ShouldEnableHybridSearch = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_DATABASE_CONFIG"):
		details := oci_generative_ai_agent.OciDatabaseConfig{}
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
		if databaseFunctions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_functions")); ok {
			interfaces := databaseFunctions.([]interface{})
			tmp := make([]oci_generative_ai_agent.DatabaseFunction, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database_functions"), stateDataIndex)
				converted, err := s.mapToDatabaseFunction(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "database_functions")) {
				details.DatabaseFunctions = tmp
			}
		}
		baseObject = details
	case strings.ToLower("OCI_OPEN_SEARCH_INDEX_CONFIG"):
		details := oci_generative_ai_agent.OciOpenSearchIndexConfig{}
		if clusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cluster_id")); ok {
			tmp := clusterId.(string)
			details.ClusterId = &tmp
		}
		if indexes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "indexes")); ok {
			interfaces := indexes.([]interface{})
			tmp := make([]oci_generative_ai_agent.Index, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "indexes"), stateDataIndex)
				converted, err := s.mapToIndex(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "indexes")) {
				details.Indexes = tmp
			}
		}
		if secretDetail, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_detail")); ok {
			if tmpList := secretDetail.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "secret_detail"), 0)
				tmp, err := s.mapToSecretDetail(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert secret_detail, encountered error: %v", err)
				}
				details.SecretDetail = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown index_config_type '%v' was specified", indexConfigType)
	}
	return baseObject, nil
}

func IndexConfigToMap(obj *oci_generative_ai_agent.IndexConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai_agent.DefaultIndexConfig:
		result["index_config_type"] = "DEFAULT_INDEX_CONFIG"

		if v.ShouldEnableHybridSearch != nil {
			result["should_enable_hybrid_search"] = bool(*v.ShouldEnableHybridSearch)
		}
	case oci_generative_ai_agent.OciDatabaseConfig:
		result["index_config_type"] = "OCI_DATABASE_CONFIG"

		if v.DatabaseConnection != nil {
			databaseConnectionArray := []interface{}{}
			if databaseConnectionMap := DatabaseConnectionToMap(&v.DatabaseConnection); databaseConnectionMap != nil {
				databaseConnectionArray = append(databaseConnectionArray, databaseConnectionMap)
			}
			result["database_connection"] = databaseConnectionArray
		}

		databaseFunctions := []interface{}{}
		for _, item := range v.DatabaseFunctions {
			databaseFunctions = append(databaseFunctions, DatabaseFunctionToMap(item))
		}
		result["database_functions"] = databaseFunctions
	case oci_generative_ai_agent.OciOpenSearchIndexConfig:
		result["index_config_type"] = "OCI_OPEN_SEARCH_INDEX_CONFIG"

		if v.ClusterId != nil {
			result["cluster_id"] = string(*v.ClusterId)
		}

		indexes := []interface{}{}
		for _, item := range v.Indexes {
			indexes = append(indexes, IndexToMap(item))
		}
		result["indexes"] = indexes

		if v.SecretDetail != nil {
			secretDetailArray := []interface{}{}
			if secretDetailMap := SecretDetailToMap(&v.SecretDetail); secretDetailMap != nil {
				secretDetailArray = append(secretDetailArray, secretDetailMap)
			}
			result["secret_detail"] = secretDetailArray
		}
	default:
		log.Printf("[WARN] Received 'index_config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) mapToIndexSchema(fieldKeyFormat string) (oci_generative_ai_agent.IndexSchema, error) {
	result := oci_generative_ai_agent.IndexSchema{}

	if bodyKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "body_key")); ok {
		tmp := bodyKey.(string)
		result.BodyKey = &tmp
	}

	if embeddingBodyKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "embedding_body_key")); ok {
		tmp := embeddingBodyKey.(string)
		result.EmbeddingBodyKey = &tmp
	}

	if titleKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "title_key")); ok {
		tmp := titleKey.(string)
		result.TitleKey = &tmp
	}

	if urlKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url_key")); ok {
		tmp := urlKey.(string)
		result.UrlKey = &tmp
	}

	return result, nil
}

func IndexSchemaToMap(obj *oci_generative_ai_agent.IndexSchema) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BodyKey != nil {
		result["body_key"] = string(*obj.BodyKey)
	}

	if obj.EmbeddingBodyKey != nil {
		result["embedding_body_key"] = string(*obj.EmbeddingBodyKey)
	}

	if obj.TitleKey != nil {
		result["title_key"] = string(*obj.TitleKey)
	}

	if obj.UrlKey != nil {
		result["url_key"] = string(*obj.UrlKey)
	}

	return result
}

func KnowledgeBaseSummaryToMap(obj oci_generative_ai_agent.KnowledgeBaseSummary) map[string]interface{} {
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

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) mapToSecretDetail(fieldKeyFormat string) (oci_generative_ai_agent.SecretDetail, error) {
	var baseObject oci_generative_ai_agent.SecretDetail
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("BASIC_AUTH_SECRET"):
		details := oci_generative_ai_agent.BasicAuthSecret{}
		if vaultSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_secret_id")); ok {
			tmp := vaultSecretId.(string)
			details.VaultSecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("IDCS_SECRET"):
		details := oci_generative_ai_agent.IdcsSecret{}
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
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func SecretDetailToMap(obj *oci_generative_ai_agent.SecretDetail) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai_agent.BasicAuthSecret:
		result["type"] = "BASIC_AUTH_SECRET"

		if v.VaultSecretId != nil {
			result["vault_secret_id"] = string(*v.VaultSecretId)
		}
	case oci_generative_ai_agent.IdcsSecret:
		result["type"] = "IDCS_SECRET"

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
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiAgentKnowledgeBaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai_agent.ChangeKnowledgeBaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.KnowledgeBaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.ChangeKnowledgeBaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getKnowledgeBaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
