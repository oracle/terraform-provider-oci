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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSensitiveTypeGroupGroupedSensitiveTypeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSensitiveTypeGroupGroupedSensitiveType,
		Read:     readDataSafeSensitiveTypeGroupGroupedSensitiveType,
		Update:   updateDataSafeSensitiveTypeGroupGroupedSensitiveType,
		Delete:   deleteDataSafeSensitiveTypeGroupGroupedSensitiveType,
		Schema: map[string]*schema.Schema{
			// Required
			"sensitive_type_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"patch_operations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"operation": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"INSERT",
								"MERGE",
								"REMOVE",
							}, true),
						},
						"selection": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:             schema.TypeMap,
							Required:         true,
							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"sensitive_type_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createDataSafeSensitiveTypeGroupGroupedSensitiveType(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeSensitiveTypeGroupGroupedSensitiveType(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSensitiveTypeGroupGroupedSensitiveType(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeSensitiveTypeGroupGroupedSensitiveType(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.GroupedSensitiveTypeCollection
	PatchResponse          *oci_data_safe.GroupedSensitiveTypeCollection
	DisableNotFoundRetries bool
}

func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud) ID() string {
	return GetSensitiveTypeGroupGroupedSensitiveTypeCompositeId(s.D.Get("sensitive_type_group_id").(string))
}

func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud) Create() error {
	err := s.Patch()
	if err != nil {
		log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
		return err
	}
	return nil
}
func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud) Patch() error {
	request := oci_data_safe.PatchGroupedSensitiveTypesRequest{}

	if patchOperations, ok := s.D.GetOkExists("patch_operations"); ok {
		interfaces := patchOperations.([]interface{})
		tmp := make([]oci_data_safe.PatchInstruction, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_operations", stateDataIndex)
			converted, err := s.mapToPatchInstruction(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("patch_operations") {
			request.Items = tmp
		}
	}

	if sensitiveTypeGroupId, ok := s.D.GetOkExists("sensitive_type_group_id"); ok {
		tmp := sensitiveTypeGroupId.(string)
		request.SensitiveTypeGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")
	response, err := s.Client.PatchGroupedSensitiveTypes(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSensitiveTypeGroupGroupedSensitiveTypeFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud) getSensitiveTypeGroupGroupedSensitiveTypeFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	sensitiveTypeGroupGroupedSensitiveTypeId, err := sensitiveTypeGroupGroupedSensitiveTypeWaitForWorkRequest(workId, "sensitivetypegroup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, sensitiveTypeGroupGroupedSensitiveTypeId)
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
	s.D.SetId(*sensitiveTypeGroupGroupedSensitiveTypeId)

	return s.Get()
}

func sensitiveTypeGroupGroupedSensitiveTypeWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func sensitiveTypeGroupGroupedSensitiveTypeWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = sensitiveTypeGroupGroupedSensitiveTypeWorkRequestShouldRetryFunc(timeout)

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
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeSensitiveTypeGroupGroupedSensitiveTypeWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSensitiveTypeGroupGroupedSensitiveTypeWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud) Get() error {
	request := oci_data_safe.ListGroupedSensitiveTypesRequest{}

	if sensitiveTypeGroupId, ok := s.D.GetOkExists("sensitive_type_group_id"); ok {
		tmp := sensitiveTypeGroupId.(string)
		request.SensitiveTypeGroupId = &tmp
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
	}

	sensitiveTypeGroupId, err := parseSensitiveTypeGroupGroupedSensitiveTypeCompositeId(s.D.Id())
	if err == nil {
		request.SensitiveTypeGroupId = &sensitiveTypeGroupId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListGroupedSensitiveTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.GroupedSensitiveTypeCollection
	return nil
}

func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud) Update() error {
	err := s.Patch()
	if err != nil {
		log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
		return err
	}
	return nil
}

func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud) SetData() error {

	sensitiveTypeGroupId, err := parseSensitiveTypeGroupGroupedSensitiveTypeCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("sensitive_type_group_id", &sensitiveTypeGroupId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, GroupedSensitiveTypeSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func GetSensitiveTypeGroupGroupedSensitiveTypeCompositeId(sensitiveTypeGroupId string) string {
	sensitiveTypeGroupId = url.PathEscape(sensitiveTypeGroupId)
	compositeId := "sensitiveTypeGroups/" + sensitiveTypeGroupId + "/groupedSensitiveTypes"
	return compositeId
}

func parseSensitiveTypeGroupGroupedSensitiveTypeCompositeId(compositeId string) (sensitiveTypeGroupId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("sensitiveTypeGroups/.*/groupedSensitiveTypes", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	sensitiveTypeGroupId, _ = url.PathUnescape(parts[1])

	return
}

func GroupedSensitiveTypeSummaryToMap(obj oci_data_safe.GroupedSensitiveTypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SensitiveTypeId != nil {
		result["sensitive_type_id"] = string(*obj.SensitiveTypeId)
	}

	return result
}

func (s *DataSafeSensitiveTypeGroupGroupedSensitiveTypeResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_data_safe.PatchInstruction, error) {
	var baseObject oci_data_safe.PatchInstruction
	//discriminator
	operationRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation"))
	var operation string
	if ok {
		operation = operationRaw.(string)
	} else {
		operation = "" // default value
	}
	switch strings.ToLower(operation) {
	case strings.ToLower("INSERT"):
		details := oci_data_safe.PatchInsertInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("MERGE"):
		details := oci_data_safe.PatchMergeInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("REMOVE"):
		details := oci_data_safe.PatchRemoveInstruction{}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown operation '%v' was specified", operation)
	}
	return baseObject, nil
}

func PatchInstructionToMap(obj oci_data_safe.PatchInstruction) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_data_safe.PatchInsertInstruction:
		result["operation"] = "INSERT"

		if v.Value != nil {
			result["value"] = []interface{}{v.Value}
		}
	case oci_data_safe.PatchMergeInstruction:
		result["operation"] = "MERGE"

		if v.Value != nil {
			result["value"] = []interface{}{v.Value}
		}
	case oci_data_safe.PatchRemoveInstruction:
		result["operation"] = "REMOVE"
	default:
		log.Printf("[WARN] Received 'operation' of unknown type %v", obj)
		return nil
	}

	return result
}
