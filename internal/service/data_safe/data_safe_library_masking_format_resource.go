// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeLibraryMaskingFormatResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeLibraryMaskingFormat,
		Read:     readDataSafeLibraryMaskingFormat,
		Update:   updateDataSafeLibraryMaskingFormat,
		Delete:   deleteDataSafeLibraryMaskingFormat,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"sensitive_type_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
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

func createDataSafeLibraryMaskingFormat(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeLibraryMaskingFormatResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeLibraryMaskingFormat(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeLibraryMaskingFormatResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeLibraryMaskingFormat(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeLibraryMaskingFormatResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeLibraryMaskingFormat(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeLibraryMaskingFormatResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeLibraryMaskingFormatResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.LibraryMaskingFormat
	DisableNotFoundRetries bool
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.MaskingLifecycleStateCreating),
	}
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.MaskingLifecycleStateActive),
		string(oci_data_safe.MaskingLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.MaskingLifecycleStateDeleting),
	}
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.MaskingLifecycleStateDeleted),
	}
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) Create() error {
	request := oci_data_safe.CreateLibraryMaskingFormatRequest{}

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

	if formatEntries, ok := s.D.GetOkExists("format_entries"); ok {
		interfaces := formatEntries.([]interface{})
		tmp := make([]oci_data_safe.FormatEntry, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "format_entries", stateDataIndex)
			converted, err := s.mapToFormatEntry(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("format_entries") {
			request.FormatEntries = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if sensitiveTypeIds, ok := s.D.GetOkExists("sensitive_type_ids"); ok {
		interfaces := sensitiveTypeIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("sensitive_type_ids") {
			request.SensitiveTypeIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateLibraryMaskingFormat(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getLibraryMaskingFormatFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, "lmfcreate", s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) getLibraryMaskingFormatFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, entityType string, timeout time.Duration) error {

	// Wait until it finishes
	libraryMaskingFormatId, err := libraryMaskingFormatWaitForWorkRequest(workId, entityType,
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, libraryMaskingFormatId)
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
	s.D.SetId(*libraryMaskingFormatId)

	return s.Get()
}

func libraryMaskingFormatWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func libraryMaskingFormatWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = libraryMaskingFormatWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeLibraryMaskingFormatWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeLibraryMaskingFormatWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeLibraryMaskingFormatResourceCrud) Get() error {
	request := oci_data_safe.GetLibraryMaskingFormatRequest{}

	tmp := s.D.Id()
	request.LibraryMaskingFormatId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetLibraryMaskingFormat(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LibraryMaskingFormat
	return nil
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateLibraryMaskingFormatRequest{}

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

	if formatEntries, ok := s.D.GetOkExists("format_entries"); ok {
		interfaces := formatEntries.([]interface{})
		tmp := make([]oci_data_safe.FormatEntry, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "format_entries", stateDataIndex)
			converted, err := s.mapToFormatEntry(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("format_entries") {
			request.FormatEntries = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.LibraryMaskingFormatId = &tmp

	if sensitiveTypeIds, ok := s.D.GetOkExists("sensitive_type_ids"); ok {
		interfaces := sensitiveTypeIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("sensitive_type_ids") {
			request.SensitiveTypeIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateLibraryMaskingFormat(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getLibraryMaskingFormatFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, "updatelmfwf", s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) Delete() error {
	request := oci_data_safe.DeleteLibraryMaskingFormatRequest{}

	tmp := s.D.Id()
	request.LibraryMaskingFormatId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.DeleteLibraryMaskingFormat(context.Background(), request)
	return err
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) SetData() error {
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

	formatEntries := []interface{}{}
	for _, item := range s.Res.FormatEntries {
		formatEntries = append(formatEntries, FormatEntryToMap(item))
	}
	s.D.Set("format_entries", formatEntries)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("sensitive_type_ids", s.Res.SensitiveTypeIds)

	s.D.Set("source", s.Res.Source)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) mapToFormatEntry(fieldKeyFormat string) (oci_data_safe.FormatEntry, error) {
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
			tmp := fixedNumber.(float32)
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
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
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
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
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

func FormatEntryToMap(obj oci_data_safe.FormatEntry) map[string]interface{} {
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

func LibraryMaskingFormatSummaryToMap(obj oci_data_safe.LibraryMaskingFormatSummary) map[string]interface{} {
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

	result["sensitive_type_ids"] = obj.SensitiveTypeIds

	result["source"] = string(obj.Source)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DataSafeLibraryMaskingFormatResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeLibraryMaskingFormatCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.LibraryMaskingFormatId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeLibraryMaskingFormatCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
