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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeMaskingPoliciesMaskingColumnResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeMaskingPoliciesMaskingColumn,
		Read:     readDataSafeMaskingPoliciesMaskingColumn,
		Update:   updateDataSafeMaskingPoliciesMaskingColumn,
		Delete:   deleteDataSafeMaskingPoliciesMaskingColumn,
		Schema: map[string]*schema.Schema{
			// Required
			"column_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"masking_policy_id": {
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

			// Optional
			"is_masking_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"masking_column_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"masking_formats": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"format_entries": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DELETE_ROWS",
											"DETERMINISTIC_ENCRYPTION",
											"DETERMINISTIC_ENCRYPTION_DATE",
											"DETERMINISTIC_SUBSTITUTION",
											"FIXED_NUMBER",
											"FIXED_STRING",
											"LIBRARY_MASKING_FORMAT",
											"NULL_VALUE",
											"PATTERN",
											"POST_PROCESSING_FUNCTION",
											"PRESERVE_ORIGINAL_DATA",
											"RANDOM_DATE",
											"RANDOM_DECIMAL_NUMBER",
											"RANDOM_DIGITS",
											"RANDOM_LIST",
											"RANDOM_NUMBER",
											"RANDOM_STRING",
											"RANDOM_SUBSTITUTION",
											"REGULAR_EXPRESSION",
											"SHUFFLE",
											"SQL_EXPRESSION",
											"SUBSTRING",
											"TRUNCATE_TABLE",
											"USER_DEFINED_FUNCTION",
										}, true),
									},

									// Optional
									"column_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"end_date": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"end_length": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"end_value": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"fixed_number": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"fixed_string": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"grouping_columns": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"length": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"library_masking_format_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pattern": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"post_processing_function": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"random_list": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"regular_expression": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"replace_with": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"schema_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sql_expression": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"start_date": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"start_length": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"start_position": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"start_value": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"table_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"user_defined_function": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Optional
						"condition": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"object_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sensitive_type_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"child_columns": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"data_type": {
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

func createDataSafeMaskingPoliciesMaskingColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPoliciesMaskingColumnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeMaskingPoliciesMaskingColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPoliciesMaskingColumnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeMaskingPoliciesMaskingColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPoliciesMaskingColumnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeMaskingPoliciesMaskingColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPoliciesMaskingColumnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeMaskingPoliciesMaskingColumnResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.MaskingColumn
	DisableNotFoundRetries bool
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) ID() string {
	column := *s.Res
	return GetMaskingPoliciesMaskingColumnCompositeId(*column.Key, s.D.Get("masking_policy_id").(string))
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.MaskingColumnLifecycleStateCreating),
	}
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.MaskingColumnLifecycleStateActive),
		string(oci_data_safe.MaskingColumnLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.MaskingColumnLifecycleStateDeleting),
	}
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) Create() error {
	request := oci_data_safe.CreateMaskingColumnRequest{}

	if columnName, ok := s.D.GetOkExists("column_name"); ok {
		tmp := columnName.(string)
		request.ColumnName = &tmp
	}

	if isMaskingEnabled, ok := s.D.GetOkExists("is_masking_enabled"); ok {
		tmp := isMaskingEnabled.(bool)
		request.IsMaskingEnabled = &tmp
	}

	if maskingColumnGroup, ok := s.D.GetOkExists("masking_column_group"); ok {
		tmp := maskingColumnGroup.(string)
		request.MaskingColumnGroup = &tmp
	}

	if maskingFormats, ok := s.D.GetOkExists("masking_formats"); ok {
		interfaces := maskingFormats.([]interface{})
		tmp := make([]oci_data_safe.MaskingFormat, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "masking_formats", stateDataIndex)
			converted, err := s.mapToMaskingFormat(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("masking_formats") {
			request.MaskingFormats = tmp
		}
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if objectType, ok := s.D.GetOkExists("object_type"); ok {
		request.ObjectType = oci_data_safe.ObjectTypeEnum(objectType.(string))
	}

	if schemaName, ok := s.D.GetOkExists("schema_name"); ok {
		tmp := schemaName.(string)
		request.SchemaName = &tmp
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateMaskingColumn(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.setIdFromWorkRequest(workId)
	return s.getMaskingPoliciesMaskingColumnFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) setIdFromWorkRequest(workId *string) {
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "maskingcolumn") {
				identifier = res.EntityUri
				break
			}
		}
	}
	if identifier != nil {
		s.D.SetId(*identifier)
	}
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) getMaskingPoliciesMaskingColumnFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	maskingPoliciesMaskingColumnId, err := maskingPoliciesMaskingColumnWaitForWorkRequest(workId, "maskingcolumn",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, maskingPoliciesMaskingColumnId)
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

	s.D.SetId(*maskingPoliciesMaskingColumnId)

	return s.Get()
}

func maskingPoliciesMaskingColumnWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func maskingPoliciesMaskingColumnWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = maskingPoliciesMaskingColumnWorkRequestShouldRetryFunc(timeout)

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
			string(oci_data_safe.WorkRequestResourceActionTypeCreated),
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
		return nil, getErrorFromDataSafeMaskingPoliciesMaskingColumnWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeMaskingPoliciesMaskingColumnWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) Get() error {
	request := oci_data_safe.GetMaskingColumnRequest{}

	if maskingColumnKey, ok := s.D.GetOkExists("key"); ok {
		tmp := maskingColumnKey.(string)
		request.MaskingColumnKey = &tmp
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	maskingColumnKey, maskingPolicyId, err := parseMaskingPoliciesMaskingColumnCompositeId(s.D.Id())
	if err == nil {
		request.MaskingColumnKey = &maskingColumnKey
		request.MaskingPolicyId = &maskingPolicyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetMaskingColumn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaskingColumn
	return nil
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) Update() error {
	request := oci_data_safe.UpdateMaskingColumnRequest{}

	if isMaskingEnabled, ok := s.D.GetOkExists("is_masking_enabled"); ok {
		tmp := isMaskingEnabled.(bool)
		request.IsMaskingEnabled = &tmp
	}

	if maskingColumnGroup, ok := s.D.GetOkExists("masking_column_group"); ok {
		tmp := maskingColumnGroup.(string)
		request.MaskingColumnGroup = &tmp
	}

	if maskingColumnKey, ok := s.D.GetOkExists("key"); ok {
		tmp := maskingColumnKey.(string)
		request.MaskingColumnKey = &tmp
	}

	if maskingFormats, ok := s.D.GetOkExists("masking_formats"); ok {
		interfaces := maskingFormats.([]interface{})
		tmp := make([]oci_data_safe.MaskingFormat, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "masking_formats", stateDataIndex)
			converted, err := s.mapToMaskingFormat(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("masking_formats") {
			request.MaskingFormats = tmp
		}
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if objectType, ok := s.D.GetOkExists("object_type"); ok {
		request.ObjectType = oci_data_safe.ObjectTypeEnum(objectType.(string))
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateMaskingColumn(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMaskingPoliciesMaskingColumnFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) Delete() error {
	request := oci_data_safe.DeleteMaskingColumnRequest{}

	if maskingColumnKey, ok := s.D.GetOkExists("key"); ok {
		tmp := maskingColumnKey.(string)
		request.MaskingColumnKey = &tmp
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.DeleteMaskingColumn(context.Background(), request)
	return err
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) SetData() error {

	maskingColumnKey, maskingPolicyId, err := parseMaskingPoliciesMaskingColumnCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", &maskingColumnKey)
		s.D.Set("masking_policy_id", &maskingPolicyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("child_columns", s.Res.ChildColumns)

	if s.Res.ColumnName != nil {
		s.D.Set("column_name", *s.Res.ColumnName)
	}

	if s.Res.DataType != nil {
		s.D.Set("data_type", *s.Res.DataType)
	}

	if s.Res.IsMaskingEnabled != nil {
		s.D.Set("is_masking_enabled", *s.Res.IsMaskingEnabled)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaskingColumnGroup != nil {
		s.D.Set("masking_column_group", *s.Res.MaskingColumnGroup)
	}

	maskingFormats := []interface{}{}
	for _, item := range s.Res.MaskingFormats {
		maskingFormats = append(maskingFormats, MaskingFormatToMap(item))
	}
	s.D.Set("masking_formats", maskingFormats)

	if s.Res.MaskingPolicyId != nil {
		s.D.Set("masking_policy_id", *s.Res.MaskingPolicyId)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	s.D.Set("object_type", s.Res.ObjectType)

	if s.Res.SchemaName != nil {
		s.D.Set("schema_name", *s.Res.SchemaName)
	}

	if s.Res.SensitiveTypeId != nil {
		s.D.Set("sensitive_type_id", *s.Res.SensitiveTypeId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetMaskingPoliciesMaskingColumnCompositeId(maskingColumnKey string, maskingPolicyId string) string {
	maskingColumnKey = url.PathEscape(maskingColumnKey)
	maskingPolicyId = url.PathEscape(maskingPolicyId)
	compositeId := "maskingPolicies/" + maskingPolicyId + "/maskingColumns/" + maskingColumnKey
	return compositeId
}

func parseMaskingPoliciesMaskingColumnCompositeId(compositeId string) (maskingColumnKey string, maskingPolicyId string, err error) {
	firstChar := compositeId[0:1]
	var compositeIdStr string
	if firstChar == "/" {
		compositeIdStr = trimLeftChar(compositeId)
	} else {
		compositeIdStr = compositeId
	}
	parts := strings.Split(compositeIdStr, "/")
	match, _ := regexp.MatchString("maskingPolicies/.*/maskingColumns/.*", compositeIdStr)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeIdStr)
		return
	}
	maskingPolicyId, _ = url.PathUnescape(parts[1])
	maskingColumnKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) mapToFormatEntry(fieldKeyFormat string) (oci_data_safe.FormatEntry, error) {
	var baseObject oci_data_safe.FormatEntry
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DELETE_ROWS"):
		details := oci_data_safe.DeleteRowsFormatEntry{}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("DETERMINISTIC_ENCRYPTION"):
		details := oci_data_safe.DeterministicEncryptionFormatEntry{}
		if regularExpression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "regular_expression")); ok {
			tmp := regularExpression.(string)
			details.RegularExpression = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("DETERMINISTIC_ENCRYPTION_DATE"):
		details := oci_data_safe.DeterministicEncryptionDateFormatEntry{}
		if endDate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "end_date")); ok {
			tmp, err := time.Parse(time.RFC3339, endDate.(string))
			if err != nil {
				return details, err
			}
			details.EndDate = &oci_common.SDKTime{Time: tmp}
		}
		if startDate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_date")); ok {
			tmp, err := time.Parse(time.RFC3339, startDate.(string))
			if err != nil {
				return details, err
			}
			details.StartDate = &oci_common.SDKTime{Time: tmp}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("DETERMINISTIC_SUBSTITUTION"):
		details := oci_data_safe.DeterministicSubstitutionFormatEntry{}
		if columnName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "column_name")); ok {
			tmp := columnName.(string)
			details.ColumnName = &tmp
		}
		if schemaName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_name")); ok {
			tmp := schemaName.(string)
			details.SchemaName = &tmp
		}
		if tableName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "table_name")); ok {
			tmp := tableName.(string)
			details.TableName = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("FIXED_NUMBER"):
		details := oci_data_safe.FixedNumberFormatEntry{}
		if fixedNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fixed_number")); ok {
			tmp := float32(fixedNumber.(float64))
			details.FixedNumber = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("FIXED_STRING"):
		details := oci_data_safe.FixedStringFormatEntry{}
		if fixedString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fixed_string")); ok {
			tmp := fixedString.(string)
			details.FixedString = &tmp
		}
		baseObject = details
	case strings.ToLower("LIBRARY_MASKING_FORMAT"):
		details := oci_data_safe.LibraryMaskingFormatEntry{}
		if libraryMaskingFormatId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "library_masking_format_id")); ok {
			tmp := libraryMaskingFormatId.(string)
			details.LibraryMaskingFormatId = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("NULL_VALUE"):
		details := oci_data_safe.NullValueFormatEntry{}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("PATTERN"):
		details := oci_data_safe.PatternFormatEntry{}
		if pattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pattern")); ok {
			tmp := pattern.(string)
			details.Pattern = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("POST_PROCESSING_FUNCTION"):
		details := oci_data_safe.PpfFormatEntry{}
		if postProcessingFunction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "post_processing_function")); ok {
			tmp := postProcessingFunction.(string)
			details.PostProcessingFunction = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("PRESERVE_ORIGINAL_DATA"):
		details := oci_data_safe.PreserveOriginalDataFormatEntry{}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("RANDOM_DATE"):
		details := oci_data_safe.RandomDateFormatEntry{}
		if endDate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "end_date")); ok {
			tmp, err := time.Parse(time.RFC3339, endDate.(string))
			if err != nil {
				return details, err
			}
			details.EndDate = &oci_common.SDKTime{Time: tmp}
		}
		if startDate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_date")); ok {
			tmp, err := time.Parse(time.RFC3339, startDate.(string))
			if err != nil {
				return details, err
			}
			details.StartDate = &oci_common.SDKTime{Time: tmp}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("RANDOM_DECIMAL_NUMBER"):
		details := oci_data_safe.RandomDecimalNumberFormatEntry{}
		if endValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "end_value")); ok {
			tmp := endValue.(float64)
			details.EndValue = &tmp
		}
		if startValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_value")); ok {
			tmp := startValue.(float64)
			details.StartValue = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("RANDOM_DIGITS"):
		details := oci_data_safe.RandomDigitsFormatEntry{}
		if endLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "end_length")); ok {
			tmp := endLength.(int)
			details.EndLength = &tmp
		}
		if startLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_length")); ok {
			tmp := startLength.(int)
			details.StartLength = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("RANDOM_LIST"):
		details := oci_data_safe.RandomListFormatEntry{}
		if randomList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "random_list")); ok {
			interfaces := randomList.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "random_list")) {
				details.RandomList = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("RANDOM_NUMBER"):
		details := oci_data_safe.RandomNumberFormatEntry{}
		if endValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "end_value")); ok {
			tmp := endValue.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert endValue string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.EndValue = &tmpInt64
		}
		if startValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_value")); ok {
			tmp := startValue.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert startValue string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.StartValue = &tmpInt64
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("RANDOM_STRING"):
		details := oci_data_safe.RandomStringFormatEntry{}
		if endLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "end_length")); ok {
			tmp := endLength.(int)
			details.EndLength = &tmp
		}
		if startLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_length")); ok {
			tmp := startLength.(int)
			details.StartLength = &tmp
		}
		baseObject = details
	case strings.ToLower("RANDOM_SUBSTITUTION"):
		details := oci_data_safe.RandomSubstitutionFormatEntry{}
		if columnName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "column_name")); ok {
			tmp := columnName.(string)
			details.ColumnName = &tmp
		}
		if schemaName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_name")); ok {
			tmp := schemaName.(string)
			details.SchemaName = &tmp
		}
		if tableName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "table_name")); ok {
			tmp := tableName.(string)
			details.TableName = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("REGULAR_EXPRESSION"):
		details := oci_data_safe.RegularExpressionFormatEntry{}
		if regularExpression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "regular_expression")); ok {
			tmp := regularExpression.(string)
			details.RegularExpression = &tmp
		}
		if replaceWith, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "replace_with")); ok {
			tmp := replaceWith.(string)
			details.ReplaceWith = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("SHUFFLE"):
		details := oci_data_safe.ShuffleFormatEntry{}
		if groupingColumns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grouping_columns")); ok {
			interfaces := groupingColumns.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "grouping_columns")) {
				details.GroupingColumns = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("SQL_EXPRESSION"):
		details := oci_data_safe.SqlExpressionFormatEntry{}
		if sqlExpression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_expression")); ok {
			tmp := sqlExpression.(string)
			details.SqlExpression = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("SUBSTRING"):
		details := oci_data_safe.SubstringFormatEntry{}
		if length, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "length")); ok {
			tmp := length.(int)
			details.Length = &tmp
		}
		if startPosition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_position")); ok {
			tmp := startPosition.(int)
			details.StartPosition = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("TRUNCATE_TABLE"):
		details := oci_data_safe.TruncateTableFormatEntry{}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("USER_DEFINED_FUNCTION"):
		details := oci_data_safe.UdfFormatEntry{}
		if userDefinedFunction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_defined_function")); ok {
			tmp := userDefinedFunction.(string)
			details.UserDefinedFunction = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func DataSafeFormatEntryToMap(obj oci_data_safe.FormatEntry) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_data_safe.DeleteRowsFormatEntry:
		result["type"] = "DELETE_ROWS"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.DeterministicEncryptionFormatEntry:
		result["type"] = "DETERMINISTIC_ENCRYPTION"

		if v.RegularExpression != nil {
			result["regular_expression"] = string(*v.RegularExpression)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.DeterministicEncryptionDateFormatEntry:
		result["type"] = "DETERMINISTIC_ENCRYPTION_DATE"

		if v.EndDate != nil {
			result["end_date"] = v.EndDate.Format(time.RFC3339Nano)
		}

		if v.StartDate != nil {
			result["start_date"] = v.StartDate.Format(time.RFC3339Nano)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.DeterministicSubstitutionFormatEntry:
		result["type"] = "DETERMINISTIC_SUBSTITUTION"

		if v.ColumnName != nil {
			result["column_name"] = string(*v.ColumnName)
		}

		if v.SchemaName != nil {
			result["schema_name"] = string(*v.SchemaName)
		}

		if v.TableName != nil {
			result["table_name"] = string(*v.TableName)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.FixedNumberFormatEntry:
		result["type"] = "FIXED_NUMBER"

		if v.FixedNumber != nil {
			result["fixed_number"] = float32(*v.FixedNumber)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.FixedStringFormatEntry:
		result["type"] = "FIXED_STRING"

		if v.FixedString != nil {
			result["fixed_string"] = string(*v.FixedString)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.LibraryMaskingFormatEntry:
		result["type"] = "LIBRARY_MASKING_FORMAT"

		if v.LibraryMaskingFormatId != nil {
			result["library_masking_format_id"] = string(*v.LibraryMaskingFormatId)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.NullValueFormatEntry:
		result["type"] = "NULL_VALUE"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.PatternFormatEntry:
		result["type"] = "PATTERN"

		if v.Pattern != nil {
			result["pattern"] = string(*v.Pattern)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.PpfFormatEntry:
		result["type"] = "POST_PROCESSING_FUNCTION"

		if v.PostProcessingFunction != nil {
			result["post_processing_function"] = string(*v.PostProcessingFunction)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.PreserveOriginalDataFormatEntry:
		result["type"] = "PRESERVE_ORIGINAL_DATA"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.RandomDateFormatEntry:
		result["type"] = "RANDOM_DATE"

		if v.EndDate != nil {
			result["end_date"] = v.EndDate.Format(time.RFC3339Nano)
		}

		if v.StartDate != nil {
			result["start_date"] = v.StartDate.Format(time.RFC3339Nano)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.RandomDecimalNumberFormatEntry:
		result["type"] = "RANDOM_DECIMAL_NUMBER"

		if v.EndValue != nil {
			result["end_value"] = float64(*v.EndValue)
		}

		if v.StartValue != nil {
			result["start_value"] = float64(*v.StartValue)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.RandomDigitsFormatEntry:
		result["type"] = "RANDOM_DIGITS"

		if v.EndLength != nil {
			result["end_length"] = int(*v.EndLength)
		}

		if v.StartLength != nil {
			result["start_length"] = int(*v.StartLength)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.RandomListFormatEntry:
		result["type"] = "RANDOM_LIST"

		result["random_list"] = v.RandomList

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.RandomNumberFormatEntry:
		result["type"] = "RANDOM_NUMBER"

		if v.EndValue != nil {
			result["end_value"] = strconv.FormatInt(*v.EndValue, 10)
		}

		if v.StartValue != nil {
			result["start_value"] = strconv.FormatInt(*v.StartValue, 10)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.RandomStringFormatEntry:
		result["type"] = "RANDOM_STRING"

		if v.EndLength != nil {
			result["end_length"] = int(*v.EndLength)
		}

		if v.StartLength != nil {
			result["start_length"] = int(*v.StartLength)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.RandomSubstitutionFormatEntry:
		result["type"] = "RANDOM_SUBSTITUTION"

		if v.ColumnName != nil {
			result["column_name"] = string(*v.ColumnName)
		}

		if v.SchemaName != nil {
			result["schema_name"] = string(*v.SchemaName)
		}

		if v.TableName != nil {
			result["table_name"] = string(*v.TableName)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.RegularExpressionFormatEntry:
		result["type"] = "REGULAR_EXPRESSION"

		if v.RegularExpression != nil {
			result["regular_expression"] = string(*v.RegularExpression)
		}

		if v.ReplaceWith != nil {
			result["replace_with"] = string(*v.ReplaceWith)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.ShuffleFormatEntry:
		result["type"] = "SHUFFLE"

		result["grouping_columns"] = v.GroupingColumns

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.SqlExpressionFormatEntry:
		result["type"] = "SQL_EXPRESSION"

		if v.SqlExpression != nil {
			result["sql_expression"] = string(*v.SqlExpression)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.SubstringFormatEntry:
		result["type"] = "SUBSTRING"

		if v.Length != nil {
			result["length"] = int(*v.Length)
		}

		if v.StartPosition != nil {
			result["start_position"] = int(*v.StartPosition)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.TruncateTableFormatEntry:
		result["type"] = "TRUNCATE_TABLE"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_data_safe.UdfFormatEntry:
		result["type"] = "USER_DEFINED_FUNCTION"

		if v.UserDefinedFunction != nil {
			result["user_defined_function"] = string(*v.UserDefinedFunction)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func MaskingColumnSummaryToMap(obj oci_data_safe.MaskingColumnSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["child_columns"] = obj.ChildColumns

	if obj.ColumnName != nil {
		result["column_name"] = string(*obj.ColumnName)
	}

	if obj.DataType != nil {
		result["data_type"] = string(*obj.DataType)
	}

	if obj.IsMaskingEnabled != nil {
		result["is_masking_enabled"] = bool(*obj.IsMaskingEnabled)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MaskingColumnGroup != nil {
		result["masking_column_group"] = string(*obj.MaskingColumnGroup)
	}

	maskingFormats := []interface{}{}
	for _, item := range obj.MaskingFormats {
		maskingFormats = append(maskingFormats, MaskingFormatToMap(item))
	}
	result["masking_formats"] = maskingFormats

	if obj.MaskingPolicyId != nil {
		result["masking_policy_id"] = string(*obj.MaskingPolicyId)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	result["object_type"] = string(obj.ObjectType)

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	if obj.SensitiveTypeId != nil {
		result["sensitive_type_id"] = string(*obj.SensitiveTypeId)
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

func (s *DataSafeMaskingPoliciesMaskingColumnResourceCrud) mapToMaskingFormat(fieldKeyFormat string) (oci_data_safe.MaskingFormat, error) {
	result := oci_data_safe.MaskingFormat{}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		tmp := condition.(string)
		result.Condition = &tmp
	}

	if formatEntries, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format_entries")); ok {
		interfaces := formatEntries.([]interface{})
		tmp := make([]oci_data_safe.FormatEntry, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "format_entries"), stateDataIndex)
			converted, err := s.mapToFormatEntry(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "format_entries")) {
			result.FormatEntries = tmp
		}
	}

	return result, nil
}

func MaskingFormatToMap(obj oci_data_safe.MaskingFormat) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Condition != nil {
		result["condition"] = string(*obj.Condition)
	}

	formatEntries := []interface{}{}
	for _, item := range obj.FormatEntries {
		formatEntries = append(formatEntries, FormatEntryToMap(item))
	}
	result["format_entries"] = formatEntries

	return result
}
