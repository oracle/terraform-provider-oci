// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeSensitiveDataModelsSensitiveColumnResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSensitiveDataModelsSensitiveColumn,
		Read:     readDataSafeSensitiveDataModelsSensitiveColumn,
		Update:   updateDataSafeSensitiveDataModelsSensitiveColumn,
		Delete:   deleteDataSafeSensitiveDataModelsSensitiveColumn,
		Schema: map[string]*schema.Schema{
			// Required
			"column_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"object": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schema_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"app_defined_child_column_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"app_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"data_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_defined_child_column_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"object_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"parent_column_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"relation_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sensitive_type_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"column_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"estimated_data_value_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sample_data_values": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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

func createDataSafeSensitiveDataModelsSensitiveColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelsSensitiveColumnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeSensitiveDataModelsSensitiveColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelsSensitiveColumnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSensitiveDataModelsSensitiveColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelsSensitiveColumnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeSensitiveDataModelsSensitiveColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelsSensitiveColumnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeSensitiveDataModelsSensitiveColumnResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SensitiveColumn
	DisableNotFoundRetries bool
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) ID() string {
	column := *s.Res
	return GetSensitiveDataModelsSensitiveColumnCompositeId(*column.Key, s.D.Get("sensitive_data_model_id").(string))
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.SensitiveColumnLifecycleStateCreating),
	}
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.SensitiveColumnLifecycleStateActive),
	}
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.SensitiveColumnLifecycleStateDeleting),
	}
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) Create() error {
	request := oci_data_safe.CreateSensitiveColumnRequest{}
	if appDefinedChildColumnKeys, ok := s.D.GetOkExists("app_defined_child_column_keys"); ok {
		interfaces := appDefinedChildColumnKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("app_defined_child_column_keys") {
			request.AppDefinedChildColumnKeys = tmp
		}
	}

	if appName, ok := s.D.GetOkExists("app_name"); ok {
		tmp := appName.(string)
		request.AppName = &tmp
	}

	if columnName, ok := s.D.GetOkExists("column_name"); ok {
		tmp := columnName.(string)
		request.ColumnName = &tmp
	}

	if dataType, ok := s.D.GetOkExists("data_type"); ok {
		tmp := dataType.(string)
		request.DataType = &tmp
	}

	if dbDefinedChildColumnKeys, ok := s.D.GetOkExists("db_defined_child_column_keys"); ok {
		interfaces := dbDefinedChildColumnKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("db_defined_child_column_keys") {
			request.DbDefinedChildColumnKeys = tmp
		}
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if objectType, ok := s.D.GetOkExists("object_type"); ok {
		request.ObjectType = oci_data_safe.CreateSensitiveColumnDetailsObjectTypeEnum(objectType.(string))
	}

	if parentColumnKeys, ok := s.D.GetOkExists("parent_column_keys"); ok {
		interfaces := parentColumnKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("parent_column_keys") {
			request.ParentColumnKeys = tmp
		}
	}

	if relationType, ok := s.D.GetOkExists("relation_type"); ok {
		request.RelationType = oci_data_safe.CreateSensitiveColumnDetailsRelationTypeEnum(relationType.(string))
	}

	if schemaName, ok := s.D.GetOkExists("schema_name"); ok {
		tmp := schemaName.(string)
		request.SchemaName = &tmp
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_data_safe.CreateSensitiveColumnDetailsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateSensitiveColumn(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.setIdFromWorkRequest(workId)
	return s.getSensitiveDataModelsSensitiveColumnFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) setIdFromWorkRequest(workId *string) {
	var identifier *string
	var err error

	workRequestResponse := oci_data_safe.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_data_safe.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "sensitivecolumn") {
				identifier = res.EntityUri
				break
			}
		}
	}
	if identifier != nil {
		s.D.SetId(*identifier)
	}
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) getSensitiveDataModelsSensitiveColumnFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	sensitiveDataModelsSensitiveColumnId, err := sensitiveDataModelsSensitiveColumnWaitForWorkRequest(workId, "sensitivecolumn",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, sensitiveDataModelsSensitiveColumnId)
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

	s.D.SetId(*sensitiveDataModelsSensitiveColumnId)

	return s.Get()
}

func sensitiveDataModelsSensitiveColumnWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func sensitiveDataModelsSensitiveColumnWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = sensitiveDataModelsSensitiveColumnWorkRequestShouldRetryFunc(timeout)
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
			string(oci_data_safe.WorkRequestResourceActionTypeUpdated),
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
				identifier = res.EntityUri
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeSensitiveDataModelsSensitiveColumnWorkRequest(client, wId, retryPolicy, entityType, action)
	}
	return identifier, nil
}

func getErrorFromDataSafeSensitiveDataModelsSensitiveColumnWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) Get() error {
	request := oci_data_safe.GetSensitiveColumnRequest{}

	sensitiveColumnKey, sensitiveDataModelId, err := parseSensitiveDataModelsSensitiveColumnCompositeId(s.D.Id())
	if err == nil {
		request.SensitiveColumnKey = &sensitiveColumnKey
		request.SensitiveDataModelId = &sensitiveDataModelId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSensitiveColumn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SensitiveColumn
	return nil
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) Update() error {
	request := oci_data_safe.UpdateSensitiveColumnRequest{}

	if appDefinedChildColumnKeys, ok := s.D.GetOkExists("app_defined_child_column_keys"); ok {
		interfaces := appDefinedChildColumnKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("app_defined_child_column_keys") {
			request.AppDefinedChildColumnKeys = tmp
		}
	}

	if dataType, ok := s.D.GetOkExists("data_type"); ok {
		tmp := dataType.(string)
		request.DataType = &tmp
	}

	if dbDefinedChildColumnKeys, ok := s.D.GetOkExists("db_defined_child_column_keys"); ok {
		interfaces := dbDefinedChildColumnKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("db_defined_child_column_keys") {
			request.DbDefinedChildColumnKeys = tmp
		}
	}

	if parentColumnKeys, ok := s.D.GetOkExists("parent_column_keys"); ok {
		interfaces := parentColumnKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("parent_column_keys") {
			request.ParentColumnKeys = tmp
		}
	}

	if relationType, ok := s.D.GetOkExists("relation_type"); ok {
		request.RelationType = oci_data_safe.UpdateSensitiveColumnDetailsRelationTypeEnum(relationType.(string))
	}

	if sensitiveColumnKey, ok := s.D.GetOkExists("key"); ok {
		tmp := sensitiveColumnKey.(string)
		request.SensitiveColumnKey = &tmp
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_data_safe.UpdateSensitiveColumnDetailsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSensitiveColumn(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSensitiveDataModelsSensitiveColumnFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) Delete() error {
	request := oci_data_safe.DeleteSensitiveColumnRequest{}

	if sensitiveColumnKey, ok := s.D.GetOkExists("key"); ok {
		tmp := sensitiveColumnKey.(string)
		request.SensitiveColumnKey = &tmp
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.DeleteSensitiveColumn(context.Background(), request)
	return err
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnResourceCrud) SetData() error {

	sensitiveColumnKey, sensitiveDataModelId, err := parseSensitiveDataModelsSensitiveColumnCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", &sensitiveColumnKey)
		s.D.Set("sensitive_data_model_id", &sensitiveDataModelId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
		if s.Res.Key != nil {
			s.D.Set("key", *s.Res.Key)
		}
		if s.Res.SensitiveDataModelId != nil {
			s.D.Set("sensitive_data_model_id", *s.Res.SensitiveDataModelId)
		}
	}

	s.D.Set("app_defined_child_column_keys", s.Res.AppDefinedChildColumnKeys)

	if s.Res.AppName != nil {
		s.D.Set("app_name", *s.Res.AppName)
	}

	s.D.Set("column_groups", s.Res.ColumnGroups)

	if s.Res.ColumnName != nil {
		s.D.Set("column_name", *s.Res.ColumnName)
	}

	if s.Res.DataType != nil {
		s.D.Set("data_type", *s.Res.DataType)
	}

	s.D.Set("db_defined_child_column_keys", s.Res.DbDefinedChildColumnKeys)

	if s.Res.EstimatedDataValueCount != nil {
		s.D.Set("estimated_data_value_count", strconv.FormatInt(*s.Res.EstimatedDataValueCount, 10))
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	s.D.Set("object_type", s.Res.ObjectType)

	s.D.Set("parent_column_keys", s.Res.ParentColumnKeys)

	s.D.Set("relation_type", s.Res.RelationType)

	s.D.Set("sample_data_values", s.Res.SampleDataValues)

	if s.Res.SchemaName != nil {
		s.D.Set("schema_name", *s.Res.SchemaName)
	}

	if s.Res.SensitiveDataModelId != nil {
		s.D.Set("sensitive_data_model_id", *s.Res.SensitiveDataModelId)
	}

	if s.Res.SensitiveTypeId != nil {
		s.D.Set("sensitive_type_id", *s.Res.SensitiveTypeId)
	}

	s.D.Set("source", s.Res.Source)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetSensitiveDataModelsSensitiveColumnCompositeId(sensitiveColumnKey string, sensitiveDataModelId string) string {
	sensitiveColumnKey = url.PathEscape(sensitiveColumnKey)
	sensitiveDataModelId = url.PathEscape(sensitiveDataModelId)
	compositeId := "sensitiveDataModels/" + sensitiveDataModelId + "/sensitiveColumns/" + sensitiveColumnKey
	return compositeId
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func parseSensitiveDataModelsSensitiveColumnCompositeId(compositeId string) (sensitiveColumnKey string, sensitiveDataModelId string, err error) {
	firstChar := compositeId[0:1]
	var compositeIdStr string
	if firstChar == "/" {
		compositeIdStr = trimLeftChar(compositeId)
	} else {
		compositeIdStr = compositeId
	}
	parts := strings.Split(compositeIdStr, "/")
	match, _ := regexp.MatchString("sensitiveDataModels/.*/sensitiveColumns/.*", compositeIdStr)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeIdStr)
		return
	}
	sensitiveDataModelId, _ = url.PathUnescape(parts[1])
	sensitiveColumnKey, _ = url.PathUnescape(parts[3])
	return
}

func SensitiveColumnSummaryToMap(obj oci_data_safe.SensitiveColumnSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AppName != nil {
		result["app_name"] = string(*obj.AppName)
	}

	if obj.ColumnName != nil {
		result["column_name"] = string(*obj.ColumnName)
	}

	if obj.DataType != nil {
		result["data_type"] = string(*obj.DataType)
	}

	if obj.EstimatedDataValueCount != nil {
		result["estimated_data_value_count"] = strconv.FormatInt(*obj.EstimatedDataValueCount, 10)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	result["object_type"] = string(obj.ObjectType)

	result["parent_column_keys"] = obj.ParentColumnKeys

	result["relation_type"] = string(obj.RelationType)

	result["sample_data_values"] = obj.SampleDataValues

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	if obj.SensitiveDataModelId != nil {
		result["sensitive_data_model_id"] = string(*obj.SensitiveDataModelId)
	}

	if obj.SensitiveTypeId != nil {
		result["sensitive_type_id"] = string(*obj.SensitiveTypeId)
	}

	result["source"] = string(obj.Source)

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
