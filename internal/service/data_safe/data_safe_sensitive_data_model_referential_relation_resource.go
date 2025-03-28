// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSensitiveDataModelReferentialRelationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSensitiveDataModelReferentialRelation,
		Read:     readDataSafeSensitiveDataModelReferentialRelation,
		Delete:   deleteDataSafeSensitiveDataModelReferentialRelation,
		Schema: map[string]*schema.Schema{
			// Required
			"child": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"app_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"column_group": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"object": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"object_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"schema_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"sensitive_type_ids": {
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
			"parent": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"app_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"column_group": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"object": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"object_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"schema_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"sensitive_type_ids": {
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
			"relation_type": {
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
			"is_sensitive": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeSensitiveDataModelReferentialRelation(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelReferentialRelationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeSensitiveDataModelReferentialRelation(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelReferentialRelationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func deleteDataSafeSensitiveDataModelReferentialRelation(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelReferentialRelationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeSensitiveDataModelReferentialRelationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.ReferentialRelation
	DisableNotFoundRetries bool
}

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) ID() string {
	refRelation := *s.Res
	return GetSensitiveDataModelReferentialRelationCompositeId(*refRelation.Key, s.D.Get("sensitive_data_model_id").(string))
}

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.ReferentialRelationLifecycleStateCreating),
	}
}

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.ReferentialRelationLifecycleStateActive),
	}
}

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.ReferentialRelationLifecycleStateActive),
	}
}

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) DeletedTarget() []string {
	return []string{}
}
func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) setIdFromWorkRequest(workId *string) {
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "referentialrelation") {
				identifier = res.EntityUri
				break
			}
		}
	}
	if identifier != nil {
		s.D.SetId(*identifier)
	}
}

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) Create() error {
	request := oci_data_safe.CreateReferentialRelationRequest{}

	if child, ok := s.D.GetOkExists("child"); ok {
		if tmpList := child.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "child", 0)
			tmp, err := s.mapToColumnsInfo(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Child = &tmp
		}
	}

	if isSensitive, ok := s.D.GetOkExists("is_sensitive"); ok {
		tmp := isSensitive.(bool)
		request.IsSensitive = &tmp
	}

	if parent, ok := s.D.GetOkExists("parent"); ok {
		if tmpList := parent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent", 0)
			tmp, err := s.mapToColumnsInfo(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Parent = &tmp
		}
	}

	if relationType, ok := s.D.GetOkExists("relation_type"); ok {
		request.RelationType = oci_data_safe.CreateReferentialRelationDetailsRelationTypeEnum(relationType.(string))
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateReferentialRelation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.setIdFromWorkRequest(workId)
	return s.getSensitiveDataModelReferentialRelationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) getSensitiveDataModelReferentialRelationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	sensitiveDataModelReferentialRelationId, err := sensitiveDataModelReferentialRelationWaitForWorkRequest(workId, "referentialrelation",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, sensitiveDataModelReferentialRelationId)
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
	s.D.SetId(*sensitiveDataModelReferentialRelationId)

	return s.Get()
}

func sensitiveDataModelReferentialRelationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func sensitiveDataModelReferentialRelationWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = sensitiveDataModelReferentialRelationWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
				identifier = res.EntityUri
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeSensitiveDataModelReferentialRelationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSensitiveDataModelReferentialRelationWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) Get() error {
	request := oci_data_safe.GetReferentialRelationRequest{}

	referentialRelationKey, sensitiveDataModelId, err := parseSensitiveDataModelReferentialRelationCompositeId(s.D.Id())
	if err == nil {
		request.ReferentialRelationKey = &referentialRelationKey
		request.SensitiveDataModelId = &sensitiveDataModelId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetReferentialRelation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ReferentialRelation
	return nil
}

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) Delete() error {
	request := oci_data_safe.DeleteReferentialRelationRequest{}

	if referentialRelationKey, ok := s.D.GetOkExists("key"); ok {
		tmp := referentialRelationKey.(string)
		request.ReferentialRelationKey = &tmp
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.DeleteReferentialRelation(context.Background(), request)
	return err
}

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) SetData() error {

	referentialRelationKey, sensitiveDataModelId, err := parseSensitiveDataModelReferentialRelationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", &referentialRelationKey)
		s.D.Set("sensitive_data_model_id", &sensitiveDataModelId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Child != nil {
		s.D.Set("child", []interface{}{ColumnsInfoToMap(s.Res.Child)})
	} else {
		s.D.Set("child", nil)
	}

	if s.Res.IsSensitive != nil {
		s.D.Set("is_sensitive", *s.Res.IsSensitive)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Parent != nil {
		s.D.Set("parent", []interface{}{ColumnsInfoToMap(s.Res.Parent)})
	} else {
		s.D.Set("parent", nil)
	}

	s.D.Set("relation_type", s.Res.RelationType)

	if s.Res.SensitiveDataModelId != nil {
		s.D.Set("sensitive_data_model_id", *s.Res.SensitiveDataModelId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}

func GetSensitiveDataModelReferentialRelationCompositeId(referentialRelationKey string, sensitiveDataModelId string) string {
	referentialRelationKey = url.PathEscape(referentialRelationKey)
	sensitiveDataModelId = url.PathEscape(sensitiveDataModelId)
	compositeId := "sensitiveDataModels/" + sensitiveDataModelId + "/referentialRelations/" + referentialRelationKey
	return compositeId
}

func parseSensitiveDataModelReferentialRelationCompositeId(compositeId string) (referentialRelationKey string, sensitiveDataModelId string, err error) {
	firstChar := compositeId[0:1]
	var compositeIdStr string
	if firstChar == "/" {
		compositeIdStr = trimLeftChar(compositeId)
	} else {
		compositeIdStr = compositeId
	}
	parts := strings.Split(compositeIdStr, "/")
	match, _ := regexp.MatchString("sensitiveDataModels/.*/referentialRelations/.*", compositeIdStr)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeIdStr)
		return
	}
	sensitiveDataModelId, _ = url.PathUnescape(parts[1])
	referentialRelationKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *DataSafeSensitiveDataModelReferentialRelationResourceCrud) mapToColumnsInfo(fieldKeyFormat string) (oci_data_safe.ColumnsInfo, error) {
	result := oci_data_safe.ColumnsInfo{}

	if appName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "app_name")); ok {
		tmp := appName.(string)
		result.AppName = &tmp
	}

	if columnGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "column_group")); ok {
		interfaces := columnGroup.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "column_group")) {
			result.ColumnGroup = tmp
		}
	}

	if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
		tmp := object.(string)
		result.ObjectName = &tmp
	}

	if objectType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_type")); ok {
		result.ObjectType = oci_data_safe.ColumnsInfoObjectTypeEnum(objectType.(string))
	}

	if schemaName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_name")); ok {
		tmp := schemaName.(string)
		result.SchemaName = &tmp
	}

	if sensitiveTypeIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sensitive_type_ids")); ok {
		interfaces := sensitiveTypeIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sensitive_type_ids")) {
			result.SensitiveTypeIds = tmp
		}
	}

	return result, nil
}

func ColumnsInfoToMap(obj *oci_data_safe.ColumnsInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AppName != nil {
		result["app_name"] = string(*obj.AppName)
	}

	result["column_group"] = obj.ColumnGroup

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	result["object_type"] = string(obj.ObjectType)

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	result["sensitive_type_ids"] = obj.SensitiveTypeIds

	return result
}

func ReferentialRelationSummaryToMap(obj oci_data_safe.ReferentialRelationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Child != nil {
		result["child"] = []interface{}{ColumnsInfoToMap(obj.Child)}
	}

	if obj.IsSensitive != nil {
		result["is_sensitive"] = bool(*obj.IsSensitive)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Parent != nil {
		result["parent"] = []interface{}{ColumnsInfoToMap(obj.Parent)}
	}

	result["relation_type"] = string(obj.RelationType)

	if obj.SensitiveDataModelId != nil {
		result["sensitive_data_model_id"] = string(*obj.SensitiveDataModelId)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}
