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

func GenerativeAiSemanticStoreResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createGenerativeAiSemanticStoreWithContext,
		ReadContext:   readGenerativeAiSemanticStoreWithContext,
		UpdateContext: updateGenerativeAiSemanticStoreWithContext,
		DeleteContext: deleteGenerativeAiSemanticStoreWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_source": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"connection_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DATABASE_TOOLS_CONNECTION",
							}, true),
						},
						"enrichment_connection_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"querying_connection_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"connection_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DATABASE_TOOLS_CONNECTION",
							}, true),
						},
						"schemas": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"refresh_schedule": {
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
								"INTERVAL",
								"NONE",
								"ON_CREATE",
							}, true),
						},

						// Optional
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func createGenerativeAiSemanticStoreWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiSemanticStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readGenerativeAiSemanticStoreWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiSemanticStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateGenerativeAiSemanticStoreWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiSemanticStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteGenerativeAiSemanticStoreWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiSemanticStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type GenerativeAiSemanticStoreResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai.GenerativeAiClient
	Res                    *oci_generative_ai.SemanticStore
	DisableNotFoundRetries bool
}

func (s *GenerativeAiSemanticStoreResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiSemanticStoreResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai.SemanticStoreLifecycleStateCreating),
	}
}

func (s *GenerativeAiSemanticStoreResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai.SemanticStoreLifecycleStateActive),
	}
}

func (s *GenerativeAiSemanticStoreResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai.SemanticStoreLifecycleStateDeleting),
	}
}

func (s *GenerativeAiSemanticStoreResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai.SemanticStoreLifecycleStateDeleted),
	}
}

func (s *GenerativeAiSemanticStoreResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_generative_ai.CreateSemanticStoreRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataSource, ok := s.D.GetOkExists("data_source"); ok {
		if tmpList := dataSource.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_source", 0)
			tmp, err := s.mapToCreateDataSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataSource = tmp
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

	if refreshSchedule, ok := s.D.GetOkExists("refresh_schedule"); ok {
		if tmpList := refreshSchedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "refresh_schedule", 0)
			tmp, err := s.mapToRefreshScheduleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RefreshSchedule = tmp
		}
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		if tmpList := schemas.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schemas", 0)
			tmp, err := s.mapToCreateSchemasDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Schemas = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.CreateSemanticStore(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getSemanticStoreFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiSemanticStoreResourceCrud) getSemanticStoreFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	semanticStoreId, err := semanticStoreWaitForWorkRequest(ctx, workId, "generativeaisemanticstore",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*semanticStoreId)

	return s.GetWithContext(ctx)
}

func semanticStoreWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func semanticStoreWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_generative_ai.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai.GenerativeAiClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai")
	retryPolicy.ShouldRetryOperation = semanticStoreWorkRequestShouldRetryFunc(timeout)

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

			// If any resource reports ActionType == IN_PROGRESS, keep waiting
			for _, res := range wr.Resources {
				log.Printf("[DEBUG]in progress wait res %v\n", res)
				if res.ActionType == oci_generative_ai.ActionTypeInProgress {
					log.Printf("[DEBUG]action type %v\n", res.ActionType)
					return wr, string(oci_generative_ai.OperationStatusInProgress), nil
				}
			}
			log.Printf("[DEBUG]in refresh wr %v\n", wr)
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		log.Printf("[DEBUG]in stateconf %v\n", e)
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action || res.ActionType == oci_generative_ai.ActionTypeCreated {
				identifier = res.Identifier
				break
			}
		}
	}
	log.Printf("[DEBUG]444 identifier %v\n", identifier)

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_generative_ai.OperationStatusFailed || response.Status == oci_generative_ai.OperationStatusCanceled {
		return nil, getErrorFromGenerativeAiSemanticStoreWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiSemanticStoreWorkRequest(ctx context.Context, client *oci_generative_ai.GenerativeAiClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai.ActionTypeEnum) error {
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

func (s *GenerativeAiSemanticStoreResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.GetSemanticStoreRequest{}

	tmp := s.D.Id()
	request.SemanticStoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.GetSemanticStore(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.SemanticStore
	return nil
}

func (s *GenerativeAiSemanticStoreResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai.UpdateSemanticStoreRequest{}

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

	if refreshSchedule, ok := s.D.GetOkExists("refresh_schedule"); ok {
		if tmpList := refreshSchedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "refresh_schedule", 0)
			tmp, err := s.mapToRefreshScheduleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RefreshSchedule = tmp
		}
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		if tmpList := schemas.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schemas", 0)
			tmp, err := s.mapToCreateSchemasDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Schemas = tmp
		}
	}

	tmp := s.D.Id()
	request.SemanticStoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.UpdateSemanticStore(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.SemanticStore
	return nil
}

func (s *GenerativeAiSemanticStoreResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_generative_ai.DeleteSemanticStoreRequest{}

	tmp := s.D.Id()
	request.SemanticStoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.DeleteSemanticStore(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := semanticStoreWaitForWorkRequest(ctx, workId, "generativeaisemanticstore",
		oci_generative_ai.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiSemanticStoreResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataSource != nil {
		dataSourceArray := []interface{}{}
		if dataSourceMap := DataSourceDetailsToMap(&s.Res.DataSource); dataSourceMap != nil {
			dataSourceArray = append(dataSourceArray, dataSourceMap)
		}
		s.D.Set("data_source", dataSourceArray)
	} else {
		s.D.Set("data_source", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	} else {
		s.D.Set("defined_tags", nil)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	} else {
		s.D.Set("description", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RefreshSchedule != nil {
		refreshScheduleArray := []interface{}{}
		if refreshScheduleMap := RefreshScheduleDetailsToMap(&s.Res.RefreshSchedule); refreshScheduleMap != nil {
			refreshScheduleArray = append(refreshScheduleArray, refreshScheduleMap)
		}
		s.D.Set("refresh_schedule", refreshScheduleArray)
	} else {
		s.D.Set("refresh_schedule", nil)
	}

	if s.Res.Schemas != nil {
		schemasArray := []interface{}{}
		if schemasMap := SchemasDetailsToMap(&s.Res.Schemas); schemasMap != nil {
			schemasArray = append(schemasArray, schemasMap)
		}
		s.D.Set("schemas", schemasArray)
	} else {
		s.D.Set("schemas", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	} else {
		s.D.Set("system_tags", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *GenerativeAiSemanticStoreResourceCrud) mapToCreateDataSourceDetails(fieldKeyFormat string) (oci_generative_ai.CreateDataSourceDetails, error) {
	var baseObject oci_generative_ai.CreateDataSourceDetails
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_type"))
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("DATABASE_TOOLS_CONNECTION"):
		details := oci_generative_ai.CreateDataSourceDatabaseToolsConnectionDetails{}
		if enrichmentConnectionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enrichment_connection_id")); ok {
			tmp := enrichmentConnectionId.(string)
			details.EnrichmentConnectionId = &tmp
		}
		if queryingConnectionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "querying_connection_id")); ok {
			tmp := queryingConnectionId.(string)
			details.QueryingConnectionId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return baseObject, nil
}

func DataSourceDetailsToMap(obj *oci_generative_ai.DataSourceDetails) map[string]interface{} {
	switch v := (*obj).(type) {
	case oci_generative_ai.CreateDataSourceDatabaseToolsConnectionDetails:
		return dataSourceDatabaseToolsConnectionDetailsToMap(v.EnrichmentConnectionId, v.QueryingConnectionId)
	case oci_generative_ai.DataSourceDatabaseToolsConnectionDetails:
		return dataSourceDatabaseToolsConnectionDetailsToMap(v.EnrichmentConnectionId, v.QueryingConnectionId)
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", *obj)
		return nil
	}
}

func dataSourceDatabaseToolsConnectionDetailsToMap(enrichmentConnectionId, queryingConnectionId *string) map[string]interface{} {
	result := map[string]interface{}{
		"connection_type": "DATABASE_TOOLS_CONNECTION",
	}

	if enrichmentConnectionId != nil {
		result["enrichment_connection_id"] = string(*enrichmentConnectionId)
	}

	if queryingConnectionId != nil {
		result["querying_connection_id"] = string(*queryingConnectionId)
	}

	return result
}

func (s *GenerativeAiSemanticStoreResourceCrud) mapToCreateSchemasDetails(fieldKeyFormat string) (oci_generative_ai.CreateSchemasDetails, error) {
	var baseObject oci_generative_ai.CreateSchemasDetails
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_type"))
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("DATABASE_TOOLS_CONNECTION"):
		details := oci_generative_ai.CreateSchemasDatabaseToolsConnectionDetails{}
		if schemas, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schemas")); ok {
			interfaces := schemas.([]interface{})
			tmp := make([]oci_generative_ai.SchemaItem, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "schemas"), stateDataIndex)
				converted, err := s.mapToSchemaItem(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "schemas")) {
				details.Schemas = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return baseObject, nil
}

func SchemasDetailsToMap(obj *oci_generative_ai.SchemasDetails) map[string]interface{} {
	switch v := (*obj).(type) {
	case oci_generative_ai.CreateSchemasDatabaseToolsConnectionDetails:
		return schemasDatabaseToolsConnectionDetailsToMap(v.Schemas)
	case oci_generative_ai.SchemasDatabaseToolsConnectionDetails:
		return schemasDatabaseToolsConnectionDetailsToMap(v.Schemas)
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", *obj)
		return nil
	}
}

func schemasDatabaseToolsConnectionDetailsToMap(items []oci_generative_ai.SchemaItem) map[string]interface{} {
	result := map[string]interface{}{
		"connection_type": "DATABASE_TOOLS_CONNECTION",
	}

	schemas := make([]interface{}, 0, len(items))
	for _, item := range items {
		schemas = append(schemas, SchemaItemToMap(item))
	}
	result["schemas"] = schemas

	return result
}

func (s *GenerativeAiSemanticStoreResourceCrud) mapToRefreshScheduleDetails(fieldKeyFormat string) (oci_generative_ai.RefreshScheduleDetails, error) {
	var baseObject oci_generative_ai.RefreshScheduleDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("INTERVAL"):
		details := oci_generative_ai.RefreshScheduleIntervalDetails{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_generative_ai.RefreshScheduleNoneDetails{}
		baseObject = details
	case strings.ToLower("ON_CREATE"):
		details := oci_generative_ai.RefreshScheduleOnCreateDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func RefreshScheduleDetailsToMap(obj *oci_generative_ai.RefreshScheduleDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai.RefreshScheduleIntervalDetails:
		result["type"] = "INTERVAL"

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}
	case oci_generative_ai.RefreshScheduleNoneDetails:
		result["type"] = "NONE"
	case oci_generative_ai.RefreshScheduleOnCreateDetails:
		result["type"] = "ON_CREATE"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiSemanticStoreResourceCrud) mapToSchemaItem(fieldKeyFormat string) (oci_generative_ai.SchemaItem, error) {
	result := oci_generative_ai.SchemaItem{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func SchemaItemToMap(obj oci_generative_ai.SchemaItem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func SemanticStoreSummaryToMap(obj oci_generative_ai.SemanticStoreSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DataSource != nil {
		dataSourceArray := []interface{}{}
		if dataSourceMap := DataSourceDetailsToMap(&obj.DataSource); dataSourceMap != nil {
			dataSourceArray = append(dataSourceArray, dataSourceMap)
		}
		result["data_source"] = dataSourceArray
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

	if obj.RefreshSchedule != nil {
		refreshScheduleArray := []interface{}{}
		if refreshScheduleMap := RefreshScheduleDetailsToMap(&obj.RefreshSchedule); refreshScheduleMap != nil {
			refreshScheduleArray = append(refreshScheduleArray, refreshScheduleMap)
		}
		result["refresh_schedule"] = refreshScheduleArray
	}

	if obj.Schemas != nil {
		schemasArray := []interface{}{}
		if schemasMap := SchemasDetailsToMap(&obj.Schemas); schemasMap != nil {
			schemasArray = append(schemasArray, schemasMap)
		}
		result["schemas"] = schemasArray
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

func (s *GenerativeAiSemanticStoreResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai.ChangeSemanticStoreCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SemanticStoreId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.ChangeSemanticStoreCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSemanticStoreFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
