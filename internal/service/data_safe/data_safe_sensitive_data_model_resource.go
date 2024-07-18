// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSensitiveDataModelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSensitiveDataModel,
		Read:     readDataSafeSensitiveDataModel,
		Update:   updateDataSafeSensitiveDataModel,
		Delete:   deleteDataSafeSensitiveDataModel,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"app_suite_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"is_app_defined_relation_discovery_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_include_all_schemas": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_include_all_sensitive_types": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_sample_data_collection_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"schemas_for_discovery": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sensitive_type_ids_for_discovery": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tables_for_discovery": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"schema_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"table_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
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

func createDataSafeSensitiveDataModel(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeSensitiveDataModel(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSensitiveDataModel(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeSensitiveDataModel(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeSensitiveDataModelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SensitiveDataModel
	DisableNotFoundRetries bool
}

func (s *DataSafeSensitiveDataModelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeSensitiveDataModelResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateCreating),
	}
}

func (s *DataSafeSensitiveDataModelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateActive),
	}
}

func (s *DataSafeSensitiveDataModelResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateDeleting),
	}
}

func (s *DataSafeSensitiveDataModelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateDeleted),
	}
}

func (s *DataSafeSensitiveDataModelResourceCrud) Create() error {
	request := oci_data_safe.CreateSensitiveDataModelRequest{}

	if appSuiteName, ok := s.D.GetOkExists("app_suite_name"); ok {
		tmp := appSuiteName.(string)
		request.AppSuiteName = &tmp
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

	if isAppDefinedRelationDiscoveryEnabled, ok := s.D.GetOkExists("is_app_defined_relation_discovery_enabled"); ok {
		tmp := isAppDefinedRelationDiscoveryEnabled.(bool)
		request.IsAppDefinedRelationDiscoveryEnabled = &tmp
	}

	if isIncludeAllSchemas, ok := s.D.GetOkExists("is_include_all_schemas"); ok {
		tmp := isIncludeAllSchemas.(bool)
		request.IsIncludeAllSchemas = &tmp
	}

	if isIncludeAllSensitiveTypes, ok := s.D.GetOkExists("is_include_all_sensitive_types"); ok {
		tmp := isIncludeAllSensitiveTypes.(bool)
		request.IsIncludeAllSensitiveTypes = &tmp
	}

	if isSampleDataCollectionEnabled, ok := s.D.GetOkExists("is_sample_data_collection_enabled"); ok {
		tmp := isSampleDataCollectionEnabled.(bool)
		request.IsSampleDataCollectionEnabled = &tmp
	}

	if schemasForDiscovery, ok := s.D.GetOkExists("schemas_for_discovery"); ok {
		interfaces := schemasForDiscovery.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas_for_discovery") {
			request.SchemasForDiscovery = tmp
		}
	}

	if sensitiveTypeIdsForDiscovery, ok := s.D.GetOkExists("sensitive_type_ids_for_discovery"); ok {
		interfaces := sensitiveTypeIdsForDiscovery.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("sensitive_type_ids_for_discovery") {
			request.SensitiveTypeIdsForDiscovery = tmp
		}
	}

	if tablesForDiscovery, ok := s.D.GetOkExists("tables_for_discovery"); ok {
		interfaces := tablesForDiscovery.([]interface{})
		tmp := make([]oci_data_safe.TablesForDiscovery, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tables_for_discovery", stateDataIndex)
			converted, err := s.mapToTablesForDiscovery(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tables_for_discovery") {
			request.TablesForDiscovery = tmp
		}
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateSensitiveDataModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getSensitiveDataModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSensitiveDataModelResourceCrud) getSensitiveDataModelFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	sensitiveDataModelId, err := sensitiveDataModelWaitForWorkRequest(workId, "sensitivedatamodel",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)
	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, sensitiveDataModelId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_data_safe.CancelWorkRequestRequest{
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
	s.D.SetId(*sensitiveDataModelId)

	return s.Get()
}

func sensitiveDataModelWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func sensitiveDataModelWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = sensitiveDataModelWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
			string(oci_data_safe.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_data_safe.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeSensitiveDataModelWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSensitiveDataModelWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_data_safe.ListWorkRequestErrorsRequest{
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

func (s *DataSafeSensitiveDataModelResourceCrud) Get() error {
	request := oci_data_safe.GetSensitiveDataModelRequest{}

	tmp := s.D.Id()
	request.SensitiveDataModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSensitiveDataModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SensitiveDataModel
	return nil
}

func (s *DataSafeSensitiveDataModelResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateSensitiveDataModelRequest{}

	if appSuiteName, ok := s.D.GetOkExists("app_suite_name"); ok {
		tmp := appSuiteName.(string)
		request.AppSuiteName = &tmp
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

	if isAppDefinedRelationDiscoveryEnabled, ok := s.D.GetOkExists("is_app_defined_relation_discovery_enabled"); ok {
		tmp := isAppDefinedRelationDiscoveryEnabled.(bool)
		request.IsAppDefinedRelationDiscoveryEnabled = &tmp
	}

	if isSampleDataCollectionEnabled, ok := s.D.GetOkExists("is_sample_data_collection_enabled"); ok {
		tmp := isSampleDataCollectionEnabled.(bool)
		request.IsSampleDataCollectionEnabled = &tmp
	}

	if schemasForDiscovery, ok := s.D.GetOkExists("schemas_for_discovery"); ok {
		interfaces := schemasForDiscovery.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas_for_discovery") {
			request.SchemasForDiscovery = tmp
		}
	}

	tmp := s.D.Id()
	request.SensitiveDataModelId = &tmp

	if sensitiveTypeIdsForDiscovery, ok := s.D.GetOkExists("sensitive_type_ids_for_discovery"); ok {
		interfaces := sensitiveTypeIdsForDiscovery.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("sensitive_type_ids_for_discovery") {
			request.SensitiveTypeIdsForDiscovery = tmp
		}
	}

	if tablesForDiscovery, ok := s.D.GetOkExists("tables_for_discovery"); ok {
		interfaces := tablesForDiscovery.([]interface{})
		tmp := make([]oci_data_safe.TablesForDiscovery, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tables_for_discovery", stateDataIndex)
			converted, err := s.mapToTablesForDiscovery(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tables_for_discovery") {
			request.TablesForDiscovery = tmp
		}
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSensitiveDataModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSensitiveDataModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSensitiveDataModelResourceCrud) Delete() error {
	request := oci_data_safe.DeleteSensitiveDataModelRequest{}

	tmp := s.D.Id()
	request.SensitiveDataModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteSensitiveDataModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := sensitiveDataModelWaitForWorkRequest(workId, "sensitivedatamodel",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeSensitiveDataModelResourceCrud) SetData() error {
	if s.Res.AppSuiteName != nil {
		s.D.Set("app_suite_name", *s.Res.AppSuiteName)
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

	if s.Res.IsAppDefinedRelationDiscoveryEnabled != nil {
		s.D.Set("is_app_defined_relation_discovery_enabled", *s.Res.IsAppDefinedRelationDiscoveryEnabled)
	}

	if s.Res.IsIncludeAllSchemas != nil {
		s.D.Set("is_include_all_schemas", *s.Res.IsIncludeAllSchemas)
	}

	if s.Res.IsIncludeAllSensitiveTypes != nil {
		s.D.Set("is_include_all_sensitive_types", *s.Res.IsIncludeAllSensitiveTypes)
	}

	if s.Res.IsSampleDataCollectionEnabled != nil {
		s.D.Set("is_sample_data_collection_enabled", *s.Res.IsSampleDataCollectionEnabled)
	}

	s.D.Set("schemas_for_discovery", s.Res.SchemasForDiscovery)

	s.D.Set("sensitive_type_ids_for_discovery", s.Res.SensitiveTypeIdsForDiscovery)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	tablesForDiscovery := []interface{}{}
	for _, item := range s.Res.TablesForDiscovery {
		tablesForDiscovery = append(tablesForDiscovery, TablesForDiscoveryToMap(item))
	}
	s.D.Set("tables_for_discovery", tablesForDiscovery)

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func SensitiveDataModelSummaryToMap(obj oci_data_safe.SensitiveDataModelSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AppSuiteName != nil {
		result["app_suite_name"] = string(*obj.AppSuiteName)
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

	result["state"] = string(obj.LifecycleState)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DataSafeSensitiveDataModelResourceCrud) mapToTablesForDiscovery(fieldKeyFormat string) (oci_data_safe.TablesForDiscovery, error) {
	result := oci_data_safe.TablesForDiscovery{}

	if schemaName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_name")); ok {
		tmp := schemaName.(string)
		result.SchemaName = &tmp
	}

	if tableNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "table_names")); ok {
		interfaces := tableNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "table_names")) {
			result.TableNames = tmp
		}
	}

	return result, nil
}

func TablesForDiscoveryToMap(obj oci_data_safe.TablesForDiscovery) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	result["table_names"] = obj.TableNames

	return result
}

func (s *DataSafeSensitiveDataModelResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeSensitiveDataModelCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SensitiveDataModelId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeSensitiveDataModelCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
