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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeMaskingPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeMaskingPolicy,
		Read:     readDataSafeMaskingPolicy,
		Update:   updateDataSafeMaskingPolicy,
		Delete:   deleteDataSafeMaskingPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"column_source": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"column_source": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SENSITIVE_DATA_MODEL",
								"TARGET",
							}, true),
						},

						// Optional
						"sensitive_data_model_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"is_drop_temp_tables_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_redo_logging_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_refresh_stats_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"parallel_degree": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"post_masking_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pre_masking_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recompile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_masking_columns_from_sdm_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
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

func createDataSafeMaskingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("add_masking_columns_from_sdm_trigger"); ok {
		err := sync.AddMaskingColumnsFromSdm()
		if err != nil {
			return err
		}
	}
	return nil

}

func readDataSafeMaskingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeMaskingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if _, ok := sync.D.GetOkExists("add_masking_columns_from_sdm_trigger"); ok && sync.D.HasChange("add_masking_columns_from_sdm_trigger") {
		oldRaw, newRaw := sync.D.GetChange("add_masking_columns_from_sdm_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.AddMaskingColumnsFromSdm()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("add_masking_columns_from_sdm_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteDataSafeMaskingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeMaskingPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.MaskingPolicy
	DisableNotFoundRetries bool
}

func (s *DataSafeMaskingPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeMaskingPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.MaskingLifecycleStateCreating),
	}
}

func (s *DataSafeMaskingPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.MaskingLifecycleStateActive),
		string(oci_data_safe.MaskingLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeMaskingPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.MaskingLifecycleStateDeleting),
	}
}

func (s *DataSafeMaskingPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.MaskingLifecycleStateDeleted),
	}
}

func (s *DataSafeMaskingPolicyResourceCrud) Create() error {
	request := oci_data_safe.CreateMaskingPolicyRequest{}

	if columnSource, ok := s.D.GetOkExists("column_source"); ok {
		if tmpList := columnSource.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "column_source", 0)
			tmp, err := s.mapToCreateColumnSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ColumnSource = tmp
		}
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

	if isDropTempTablesEnabled, ok := s.D.GetOkExists("is_drop_temp_tables_enabled"); ok {
		tmp := isDropTempTablesEnabled.(bool)
		request.IsDropTempTablesEnabled = &tmp
	}

	if isRedoLoggingEnabled, ok := s.D.GetOkExists("is_redo_logging_enabled"); ok {
		tmp := isRedoLoggingEnabled.(bool)
		request.IsRedoLoggingEnabled = &tmp
	}

	if isRefreshStatsEnabled, ok := s.D.GetOkExists("is_refresh_stats_enabled"); ok {
		tmp := isRefreshStatsEnabled.(bool)
		request.IsRefreshStatsEnabled = &tmp
	}

	if parallelDegree, ok := s.D.GetOkExists("parallel_degree"); ok {
		tmp := parallelDegree.(string)
		request.ParallelDegree = &tmp
	}

	if postMaskingScript, ok := s.D.GetOkExists("post_masking_script"); ok {
		tmp := postMaskingScript.(string)
		request.PostMaskingScript = &tmp
	}

	if preMaskingScript, ok := s.D.GetOkExists("pre_masking_script"); ok {
		tmp := preMaskingScript.(string)
		request.PreMaskingScript = &tmp
	}

	if recompile, ok := s.D.GetOkExists("recompile"); ok {
		request.Recompile = oci_data_safe.MaskingPolicyRecompileEnum(recompile.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateMaskingPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getMaskingPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeMaskingPolicyResourceCrud) getMaskingPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	maskingPolicyId, err := maskingPolicyWaitForWorkRequest(workId, "masking_policy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, maskingPolicyId)
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
	s.D.SetId(*maskingPolicyId)

	return s.Get()
}

func maskingPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func maskingPolicyWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = maskingPolicyWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeMaskingPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeMaskingPolicyWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeMaskingPolicyResourceCrud) Get() error {
	request := oci_data_safe.GetMaskingPolicyRequest{}

	tmp := s.D.Id()
	request.MaskingPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetMaskingPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaskingPolicy
	return nil
}

func (s *DataSafeMaskingPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateMaskingPolicyRequest{}

	if columnSource, ok := s.D.GetOkExists("column_source"); ok {
		if tmpList := columnSource.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "column_source", 0)
			tmp, err := s.mapToUpdateColumnSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ColumnSource = tmp
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

	if isDropTempTablesEnabled, ok := s.D.GetOkExists("is_drop_temp_tables_enabled"); ok {
		tmp := isDropTempTablesEnabled.(bool)
		request.IsDropTempTablesEnabled = &tmp
	}

	if isRedoLoggingEnabled, ok := s.D.GetOkExists("is_redo_logging_enabled"); ok {
		tmp := isRedoLoggingEnabled.(bool)
		request.IsRedoLoggingEnabled = &tmp
	}

	if isRefreshStatsEnabled, ok := s.D.GetOkExists("is_refresh_stats_enabled"); ok {
		tmp := isRefreshStatsEnabled.(bool)
		request.IsRefreshStatsEnabled = &tmp
	}

	tmp := s.D.Id()
	request.MaskingPolicyId = &tmp

	if parallelDegree, ok := s.D.GetOkExists("parallel_degree"); ok {
		tmp := parallelDegree.(string)
		request.ParallelDegree = &tmp
	}

	if postMaskingScript, ok := s.D.GetOkExists("post_masking_script"); ok {
		tmp := postMaskingScript.(string)
		request.PostMaskingScript = &tmp
	}

	if preMaskingScript, ok := s.D.GetOkExists("pre_masking_script"); ok {
		tmp := preMaskingScript.(string)
		request.PreMaskingScript = &tmp
	}

	if recompile, ok := s.D.GetOkExists("recompile"); ok {
		request.Recompile = oci_data_safe.MaskingPolicyRecompileEnum(recompile.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateMaskingPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMaskingPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeMaskingPolicyResourceCrud) Delete() error {
	request := oci_data_safe.DeleteMaskingPolicyRequest{}

	tmp := s.D.Id()
	request.MaskingPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteMaskingPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := maskingPolicyWaitForWorkRequest(workId, "masking_policy",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeMaskingPolicyResourceCrud) SetData() error {
	if s.Res.ColumnSource != nil {
		columnSourceArray := []interface{}{}
		if columnSourceMap := ColumnSourceDetailsToMap(&s.Res.ColumnSource); columnSourceMap != nil {
			columnSourceArray = append(columnSourceArray, columnSourceMap)
		}
		s.D.Set("column_source", columnSourceArray)
	} else {
		s.D.Set("column_source", nil)
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

	if s.Res.IsDropTempTablesEnabled != nil {
		s.D.Set("is_drop_temp_tables_enabled", *s.Res.IsDropTempTablesEnabled)
	}

	if s.Res.IsRedoLoggingEnabled != nil {
		s.D.Set("is_redo_logging_enabled", *s.Res.IsRedoLoggingEnabled)
	}

	if s.Res.IsRefreshStatsEnabled != nil {
		s.D.Set("is_refresh_stats_enabled", *s.Res.IsRefreshStatsEnabled)
	}

	if s.Res.ParallelDegree != nil {
		s.D.Set("parallel_degree", *s.Res.ParallelDegree)
	}

	if s.Res.PostMaskingScript != nil {
		s.D.Set("post_masking_script", *s.Res.PostMaskingScript)
	}

	if s.Res.PreMaskingScript != nil {
		s.D.Set("pre_masking_script", *s.Res.PreMaskingScript)
	}

	s.D.Set("recompile", s.Res.Recompile)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataSafeMaskingPolicyResourceCrud) AddMaskingColumnsFromSdm() error {
	return nil
}

func ColumnSourceDetailsToMap(obj *oci_data_safe.ColumnSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_safe.ColumnSourceFromSdmDetails:
		result["column_source"] = "SENSITIVE_DATA_MODEL"

		if v.SensitiveDataModelId != nil {
			result["sensitive_data_model_id"] = string(*v.SensitiveDataModelId)
		}
	case oci_data_safe.ColumnSourceFromTargetDetails:
		result["column_source"] = "TARGET"

		if v.TargetId != nil {
			result["target_id"] = string(*v.TargetId)
		}
	default:
		log.Printf("[WARN] Received 'column_source' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DataSafeMaskingPolicyResourceCrud) mapToCreateColumnSourceDetails(fieldKeyFormat string) (oci_data_safe.CreateColumnSourceDetails, error) {
	var baseObject oci_data_safe.CreateColumnSourceDetails
	//discriminator
	columnSourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "column_source"))
	var columnSource string
	if ok {
		columnSource = columnSourceRaw.(string)
	} else {
		columnSource = "" // default value
	}
	switch strings.ToLower(columnSource) {
	case strings.ToLower("SENSITIVE_DATA_MODEL"):
		details := oci_data_safe.CreateColumnSourceFromSdmDetails{}
		if sensitiveDataModelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sensitive_data_model_id")); ok {
			tmp := sensitiveDataModelId.(string)
			details.SensitiveDataModelId = &tmp
		}
		baseObject = details
	case strings.ToLower("TARGET"):
		details := oci_data_safe.CreateColumnSourceFromTargetDetails{}
		if targetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_id")); ok {
			tmp := targetId.(string)
			details.TargetId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown column_source '%v' was specified", columnSource)
	}
	return baseObject, nil
}

func (s *DataSafeMaskingPolicyResourceCrud) mapToUpdateColumnSourceDetails(fieldKeyFormat string) (oci_data_safe.CreateColumnSourceDetails, error) {
	var baseObject oci_data_safe.CreateColumnSourceDetails
	//discriminator
	columnSourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "column_source"))
	var columnSource string
	if ok {
		columnSource = columnSourceRaw.(string)
	} else {
		columnSource = "" // default value
	}
	switch strings.ToLower(columnSource) {
	case strings.ToLower("SENSITIVE_DATA_MODEL"):
		details := oci_data_safe.UpdateColumnSourceSdmDetails{}
		if sensitiveDataModelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sensitive_data_model_id")); ok {
			tmp := sensitiveDataModelId.(string)
			details.SensitiveDataModelId = &tmp
		}
		baseObject = details
	case strings.ToLower("TARGET"):
		details := oci_data_safe.UpdateColumnSourceTargetDetails{}
		if targetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_id")); ok {
			tmp := targetId.(string)
			details.TargetId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown column_source '%v' was specified", columnSource)
	}
	return baseObject, nil
}

func DataSafeColumnSourceDetailsToMap(obj *oci_data_safe.ColumnSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_safe.ColumnSourceFromSdmDetails:
		result["column_source"] = "SENSITIVE_DATA_MODEL"

		if v.SensitiveDataModelId != nil {
			result["sensitive_data_model_id"] = string(*v.SensitiveDataModelId)
		}
	case oci_data_safe.ColumnSourceFromTargetDetails:
		result["column_source"] = "TARGET"

		if v.TargetId != nil {
			result["target_id"] = string(*v.TargetId)
		}
	default:
		log.Printf("[WARN] Received 'column_source' of unknown type %v", *obj)
		return nil
	}

	return result
}

func MaskingPolicySummaryToMap(obj oci_data_safe.MaskingPolicySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ColumnSource != nil {
		columnSourceArray := []interface{}{}
		if columnSourceMap := ColumnSourceDetailsToMap(&obj.ColumnSource); columnSourceMap != nil {
			columnSourceArray = append(columnSourceArray, columnSourceMap)
		}
		result["column_source"] = columnSourceArray
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

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DataSafeMaskingPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeMaskingPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MaskingPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeMaskingPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
