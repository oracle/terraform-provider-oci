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

func GenerativeAiAgentDataSourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("40m"),
			Update: tfresource.GetTimeoutDuration("20m"),
			Delete: tfresource.GetTimeoutDuration("20m"),
		},
		Create: createGenerativeAiAgentDataSource,
		Read:   readGenerativeAiAgentDataSource,
		Update: updateGenerativeAiAgentDataSource,
		Delete: deleteGenerativeAiAgentDataSource,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_source_config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"data_source_config_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"OCI_OBJECT_STORAGE",
							}, true),
						},
						"object_storage_prefixes": {
							Type:     schema.TypeList,
							Required: true,
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

						// Computed
					},
				},
			},
			"knowledge_base_id": {
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

func createGenerativeAiAgentDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.CreateResource(d, sync)
}

func readGenerativeAiAgentDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

func updateGenerativeAiAgentDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGenerativeAiAgentDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GenerativeAiAgentDataSourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai_agent.GenerativeAiAgentClient
	Res                    *oci_generative_ai_agent.DataSource
	DisableNotFoundRetries bool
}

func (s *GenerativeAiAgentDataSourceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiAgentDataSourceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai_agent.DataSourceLifecycleStateCreating),
	}
}

func (s *GenerativeAiAgentDataSourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.DataSourceLifecycleStateActive),
	}
}

func (s *GenerativeAiAgentDataSourceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai_agent.DataSourceLifecycleStateDeleting),
	}
}

func (s *GenerativeAiAgentDataSourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.DataSourceLifecycleStateDeleted),
	}
}

func (s *GenerativeAiAgentDataSourceResourceCrud) Create() error {
	request := oci_generative_ai_agent.CreateDataSourceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataSourceConfig, ok := s.D.GetOkExists("data_source_config"); ok {
		if tmpList := dataSourceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_source_config", 0)
			tmp, err := s.mapToDataSourceConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataSourceConfig = tmp
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

	if knowledgeBaseId, ok := s.D.GetOkExists("knowledge_base_id"); ok {
		tmp := knowledgeBaseId.(string)
		request.KnowledgeBaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.CreateDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDataSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiAgentDataSourceResourceCrud) getDataSourceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai_agent.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	dataSourceId, err := dataSourceWaitForWorkRequest(workId, "datasource",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, dataSourceId)
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
	s.D.SetId(*dataSourceId)

	return s.Get()
}

func dataSourceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func dataSourceWaitForWorkRequest(wId *string, entityType string, action oci_generative_ai_agent.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai_agent.GenerativeAiAgentClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai_agent")
	retryPolicy.ShouldRetryOperation = dataSourceWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromGenerativeAiAgentDataSourceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiAgentDataSourceWorkRequest(client *oci_generative_ai_agent.GenerativeAiAgentClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai_agent.ActionTypeEnum) error {
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

func (s *GenerativeAiAgentDataSourceResourceCrud) Get() error {
	request := oci_generative_ai_agent.GetDataSourceRequest{}

	tmp := s.D.Id()
	request.DataSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.GetDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataSource
	return nil
}

func (s *GenerativeAiAgentDataSourceResourceCrud) Update() error {
	request := oci_generative_ai_agent.UpdateDataSourceRequest{}

	if dataSourceConfig, ok := s.D.GetOkExists("data_source_config"); ok {
		if tmpList := dataSourceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_source_config", 0)
			tmp, err := s.mapToDataSourceConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataSourceConfig = tmp
		}
	}

	tmp := s.D.Id()
	request.DataSourceId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.UpdateDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDataSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GenerativeAiAgentDataSourceResourceCrud) Delete() error {
	request := oci_generative_ai_agent.DeleteDataSourceRequest{}

	tmp := s.D.Id()
	request.DataSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.DeleteDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := dataSourceWaitForWorkRequest(workId, "datasource",
		oci_generative_ai_agent.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiAgentDataSourceResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataSourceConfig != nil {
		dataSourceConfigArray := []interface{}{}
		if dataSourceConfigMap := DataSourceConfigToMap(&s.Res.DataSourceConfig); dataSourceConfigMap != nil {
			dataSourceConfigArray = append(dataSourceConfigArray, dataSourceConfigMap)
		}
		s.D.Set("data_source_config", dataSourceConfigArray)
	} else {
		s.D.Set("data_source_config", nil)
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

	if s.Res.KnowledgeBaseId != nil {
		s.D.Set("knowledge_base_id", *s.Res.KnowledgeBaseId)
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

func (s *GenerativeAiAgentDataSourceResourceCrud) mapToDataSourceConfig(fieldKeyFormat string) (oci_generative_ai_agent.DataSourceConfig, error) {
	var baseObject oci_generative_ai_agent.DataSourceConfig
	//discriminator
	dataSourceConfigTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_source_config_type"))
	var dataSourceConfigType string
	if ok {
		dataSourceConfigType = dataSourceConfigTypeRaw.(string)
	} else {
		dataSourceConfigType = "" // default value
	}
	switch strings.ToLower(dataSourceConfigType) {
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_generative_ai_agent.OciObjectStorageDataSourceConfig{}
		if objectStoragePrefixes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_prefixes")); ok {
			interfaces := objectStoragePrefixes.([]interface{})
			tmp := make([]oci_generative_ai_agent.ObjectStoragePrefix, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_prefixes"), stateDataIndex)
				converted, err := s.mapToObjectStoragePrefix(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "object_storage_prefixes")) {
				details.ObjectStoragePrefixes = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown data_source_config_type '%v' was specified", dataSourceConfigType)
	}
	return baseObject, nil
}

func DataSourceConfigToMap(obj *oci_generative_ai_agent.DataSourceConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai_agent.OciObjectStorageDataSourceConfig:
		result["data_source_config_type"] = "OCI_OBJECT_STORAGE"

		objectStoragePrefixes := []interface{}{}
		for _, item := range v.ObjectStoragePrefixes {
			objectStoragePrefixes = append(objectStoragePrefixes, ObjectStoragePrefixToMap(item))
		}
		result["object_storage_prefixes"] = objectStoragePrefixes
	default:
		log.Printf("[WARN] Received 'data_source_config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DataSourceSummaryToMap(obj oci_generative_ai_agent.DataSourceSummary) map[string]interface{} {
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

	if obj.KnowledgeBaseId != nil {
		result["knowledge_base_id"] = string(*obj.KnowledgeBaseId)
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

func (s *GenerativeAiAgentDataSourceResourceCrud) mapToObjectStoragePrefix(fieldKeyFormat string) (oci_generative_ai_agent.ObjectStoragePrefix, error) {
	result := oci_generative_ai_agent.ObjectStoragePrefix{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.NamespaceName = &tmp
	}

	if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
		tmp := prefix.(string)
		result.Prefix = &tmp
	}

	return result, nil
}

func ObjectStoragePrefixToMap(obj oci_generative_ai_agent.ObjectStoragePrefix) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	if obj.Prefix != nil {
		result["prefix"] = string(*obj.Prefix)
	}

	return result
}
