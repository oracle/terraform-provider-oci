// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeDiscoveryJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeDiscoveryJob,
		Read:     readDataSafeDiscoveryJob,
		Update:   updateDataSafeDiscoveryJob,
		Delete:   deleteDataSafeDiscoveryJob,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"discovery_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"is_app_defined_relation_discovery_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"schemas_for_discovery": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sensitive_type_ids_for_discovery": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tables_for_discovery": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"schema_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"table_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
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
			"target_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_columns_scanned": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_deleted_sensitive_columns": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_modified_sensitive_columns": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_new_sensitive_columns": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_objects_scanned": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_schemas_scanned": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeDiscoveryJob(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeDiscoveryJob(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeDiscoveryJob(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeDiscoveryJob(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeDiscoveryJobResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.DiscoveryJob
	DisableNotFoundRetries bool
}

func (s *DataSafeDiscoveryJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeDiscoveryJobResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateCreating),
	}
}

func (s *DataSafeDiscoveryJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateActive),
	}
}

func (s *DataSafeDiscoveryJobResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateDeleting),
	}
}

func (s *DataSafeDiscoveryJobResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateDeleted),
	}
}

func (s *DataSafeDiscoveryJobResourceCrud) Create() error {
	request := oci_data_safe.CreateDiscoveryJobRequest{}

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

	if discoveryType, ok := s.D.GetOkExists("discovery_type"); ok {
		request.DiscoveryType = oci_data_safe.DiscoveryJobDiscoveryTypeEnum(discoveryType.(string))
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

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateDiscoveryJob(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDiscoveryJobFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeDiscoveryJobResourceCrud) getDiscoveryJobFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	discoveryJobId, err := discoveryJobWaitForWorkRequest(workId, "discoveryjob",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, discoveryJobId)
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
	s.D.SetId(*discoveryJobId)

	return s.Get()
}

func discoveryJobWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func discoveryJobWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = discoveryJobWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeDiscoveryJobWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeDiscoveryJobWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeDiscoveryJobResourceCrud) Get() error {
	request := oci_data_safe.GetDiscoveryJobRequest{}

	tmp := s.D.Id()
	request.DiscoveryJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetDiscoveryJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DiscoveryJob
	return nil
}

func (s *DataSafeDiscoveryJobResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	//s.D.SetId(*discoveryJobId)
	return s.Get()
}

func (s *DataSafeDiscoveryJobResourceCrud) Delete() error {
	request := oci_data_safe.DeleteDiscoveryJobRequest{}

	tmp := s.D.Id()
	request.DiscoveryJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteDiscoveryJob(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := discoveryJobWaitForWorkRequest(workId, "discoveryjob",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeDiscoveryJobResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("discovery_type", s.Res.DiscoveryType)

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

	if s.Res.SensitiveDataModelId != nil {
		s.D.Set("sensitive_data_model_id", *s.Res.SensitiveDataModelId)
	}

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

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TotalColumnsScanned != nil {
		s.D.Set("total_columns_scanned", strconv.FormatInt(*s.Res.TotalColumnsScanned, 10))
	}

	if s.Res.TotalDeletedSensitiveColumns != nil {
		s.D.Set("total_deleted_sensitive_columns", strconv.FormatInt(*s.Res.TotalDeletedSensitiveColumns, 10))
	}

	if s.Res.TotalModifiedSensitiveColumns != nil {
		s.D.Set("total_modified_sensitive_columns", strconv.FormatInt(*s.Res.TotalModifiedSensitiveColumns, 10))
	}

	if s.Res.TotalNewSensitiveColumns != nil {
		s.D.Set("total_new_sensitive_columns", strconv.FormatInt(*s.Res.TotalNewSensitiveColumns, 10))
	}

	if s.Res.TotalObjectsScanned != nil {
		s.D.Set("total_objects_scanned", strconv.FormatInt(*s.Res.TotalObjectsScanned, 10))
	}

	if s.Res.TotalSchemasScanned != nil {
		s.D.Set("total_schemas_scanned", strconv.FormatInt(*s.Res.TotalSchemasScanned, 10))
	}

	return nil
}

func DiscoveryJobSummaryToMap(obj oci_data_safe.DiscoveryJobSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["discovery_type"] = string(obj.DiscoveryType)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.SensitiveDataModelId != nil {
		result["sensitive_data_model_id"] = string(*obj.SensitiveDataModelId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}

func (s *DataSafeDiscoveryJobResourceCrud) mapToTablesForDiscovery(fieldKeyFormat string) (oci_data_safe.TablesForDiscovery, error) {
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

func (s *DataSafeDiscoveryJobResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeDiscoveryJobCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DiscoveryJobId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeDiscoveryJobCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	/*if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}*/

	return nil
}
