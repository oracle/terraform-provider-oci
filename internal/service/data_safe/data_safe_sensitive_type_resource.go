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

func DataSafeSensitiveTypeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSensitiveType,
		Read:     readDataSafeSensitiveType,
		Update:   updateDataSafeSensitiveType,
		Delete:   deleteDataSafeSensitiveType,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"SENSITIVE_CATEGORY",
					"SENSITIVE_TYPE",
				}, true),
			},

			// Optional
			"comment_pattern": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_pattern": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_masking_format_id": {
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
			"name_pattern": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_category_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"search_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"short_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"is_common": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"source": {
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

func createDataSafeSensitiveType(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeSensitiveType(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSensitiveType(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeSensitiveType(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeSensitiveTypeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SensitiveType
	DisableNotFoundRetries bool
}

func (s *DataSafeSensitiveTypeResourceCrud) ID() string {
	sensitiveType := *s.Res
	return *sensitiveType.GetId()
}

func (s *DataSafeSensitiveTypeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateCreating),
	}
}

func (s *DataSafeSensitiveTypeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateActive),
	}
}

func (s *DataSafeSensitiveTypeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateDeleting),
	}
}

func (s *DataSafeSensitiveTypeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.DiscoveryLifecycleStateDeleted),
	}
}

func (s *DataSafeSensitiveTypeResourceCrud) Create() error {
	request := oci_data_safe.CreateSensitiveTypeRequest{}
	err := s.populateTopLevelPolymorphicCreateSensitiveTypeRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateSensitiveType(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getSensitiveTypeFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSensitiveTypeResourceCrud) getSensitiveTypeFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	sensitiveTypeId, err := sensitiveTypeWaitForWorkRequest(workId, "sensitivetype",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v and enum: %v\n", workId, sensitiveTypeId, actionTypeEnum)
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
	s.D.SetId(*sensitiveTypeId)

	return s.Get()
}

func sensitiveTypeWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func sensitiveTypeWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = sensitiveTypeWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeSensitiveTypeWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSensitiveTypeWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSensitiveTypeResourceCrud) Get() error {
	request := oci_data_safe.GetSensitiveTypeRequest{}

	tmp := s.D.Id()
	request.SensitiveTypeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSensitiveType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SensitiveType
	return nil
}

func (s *DataSafeSensitiveTypeResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateSensitiveTypeRequest{}
	err := s.populateTopLevelPolymorphicUpdateSensitiveTypeRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSensitiveType(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSensitiveTypeFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSensitiveTypeResourceCrud) Delete() error {
	request := oci_data_safe.DeleteSensitiveTypeRequest{}

	tmp := s.D.Id()
	request.SensitiveTypeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.DeleteSensitiveType(context.Background(), request)
	return err
}

func (s *DataSafeSensitiveTypeResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_data_safe.SensitiveCategory:
		s.D.Set("entity_type", "SENSITIVE_CATEGORY")

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.IsCommon != nil {
			s.D.Set("is_common", *v.IsCommon)
		}

		if v.ParentCategoryId != nil {
			s.D.Set("parent_category_id", *v.ParentCategoryId)
		}

		if v.ShortName != nil {
			s.D.Set("short_name", *v.ShortName)
		}

		s.D.Set("sensitive_type_source", v.Source)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_data_safe.SensitiveTypePattern:
		s.D.Set("entity_type", "SENSITIVE_TYPE")

		if v.CommentPattern != nil {
			s.D.Set("comment_pattern", *v.CommentPattern)
		}

		if v.DataPattern != nil {
			s.D.Set("data_pattern", *v.DataPattern)
		}

		if v.DefaultMaskingFormatId != nil {
			s.D.Set("default_masking_format_id", *v.DefaultMaskingFormatId)
		}

		if v.NamePattern != nil {
			s.D.Set("name_pattern", *v.NamePattern)
		}

		s.D.Set("search_type", v.SearchType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefaultMaskingFormatId != nil {
			s.D.Set("default_masking_format_id", *v.DefaultMaskingFormatId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.IsCommon != nil {
			s.D.Set("is_common", *v.IsCommon)
		}

		if v.ParentCategoryId != nil {
			s.D.Set("parent_category_id", *v.ParentCategoryId)
		}

		if v.ShortName != nil {
			s.D.Set("short_name", *v.ShortName)
		}

		s.D.Set("source", v.Source)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'entity_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func SensitiveTypeSummaryToMap(obj oci_data_safe.SensitiveTypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefaultMaskingFormatId != nil {
		result["default_masking_format_id"] = string(*obj.DefaultMaskingFormatId)
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

	result["entity_type"] = string(obj.EntityType)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsCommon != nil {
		result["is_common"] = bool(*obj.IsCommon)
	}

	if obj.ParentCategoryId != nil {
		result["parent_category_id"] = string(*obj.ParentCategoryId)
	}

	if obj.ShortName != nil {
		result["short_name"] = string(*obj.ShortName)
	}

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

func (s *DataSafeSensitiveTypeResourceCrud) populateTopLevelPolymorphicCreateSensitiveTypeRequest(request *oci_data_safe.CreateSensitiveTypeRequest) error {
	//discriminator
	entityTypeRaw, ok := s.D.GetOkExists("entity_type")
	var entityType string
	if ok {
		entityType = entityTypeRaw.(string)
	} else {
		entityType = "" // default value
	}
	switch strings.ToLower(entityType) {
	case strings.ToLower("SENSITIVE_CATEGORY"):
		details := oci_data_safe.CreateSensitiveCategoryDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}

		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if parentCategoryId, ok := s.D.GetOkExists("parent_category_id"); ok {
			tmp := parentCategoryId.(string)
			details.ParentCategoryId = &tmp
		}
		if shortName, ok := s.D.GetOkExists("short_name"); ok {
			tmp := shortName.(string)
			details.ShortName = &tmp
		}
		request.CreateSensitiveTypeDetails = details
	case strings.ToLower("SENSITIVE_TYPE"):
		details := oci_data_safe.CreateSensitiveTypePatternDetails{}
		if commentPattern, ok := s.D.GetOkExists("comment_pattern"); ok {
			tmp := commentPattern.(string)
			details.CommentPattern = &tmp
		}
		if dataPattern, ok := s.D.GetOkExists("data_pattern"); ok {
			tmp := dataPattern.(string)
			details.DataPattern = &tmp
		}
		if defaultMaskingFormatId, ok := s.D.GetOkExists("default_masking_format_id"); ok {
			tmp := defaultMaskingFormatId.(string)
			details.DefaultMaskingFormatId = &tmp
		}
		if namePattern, ok := s.D.GetOkExists("name_pattern"); ok {
			tmp := namePattern.(string)
			details.NamePattern = &tmp
		}
		if searchType, ok := s.D.GetOkExists("search_type"); ok {
			details.SearchType = oci_data_safe.SensitiveTypePatternSearchTypeEnum(searchType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if defaultMaskingFormatId, ok := s.D.GetOkExists("default_masking_format_id"); ok {
			tmp := defaultMaskingFormatId.(string)
			details.DefaultMaskingFormatId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if parentCategoryId, ok := s.D.GetOkExists("parent_category_id"); ok {
			tmp := parentCategoryId.(string)
			details.ParentCategoryId = &tmp
		}
		if shortName, ok := s.D.GetOkExists("short_name"); ok {
			tmp := shortName.(string)
			details.ShortName = &tmp
		}
		request.CreateSensitiveTypeDetails = details
	default:
		return fmt.Errorf("unknown entity_type '%v' was specified", entityType)
	}
	return nil
}

func (s *DataSafeSensitiveTypeResourceCrud) populateTopLevelPolymorphicUpdateSensitiveTypeRequest(request *oci_data_safe.UpdateSensitiveTypeRequest) error {
	//discriminator
	entityTypeRaw, ok := s.D.GetOkExists("entity_type")
	var entityType string
	if ok {
		entityType = entityTypeRaw.(string)
	} else {
		entityType = "" // default value
	}
	switch strings.ToLower(entityType) {
	case strings.ToLower("SENSITIVE_CATEGORY"):
		details := oci_data_safe.UpdateSensitiveCategoryDetails{}

		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if parentCategoryId, ok := s.D.GetOkExists("parent_category_id"); ok {
			tmp := parentCategoryId.(string)
			details.ParentCategoryId = &tmp
		}
		tmp := s.D.Id()
		request.SensitiveTypeId = &tmp
		if shortName, ok := s.D.GetOkExists("short_name"); ok {
			tmp := shortName.(string)
			details.ShortName = &tmp
		}
		request.UpdateSensitiveTypeDetails = details
	case strings.ToLower("SENSITIVE_TYPE"):
		details := oci_data_safe.UpdateSensitiveTypePatternDetails{}
		if commentPattern, ok := s.D.GetOkExists("comment_pattern"); ok {
			tmp := commentPattern.(string)
			details.CommentPattern = &tmp
		}
		if dataPattern, ok := s.D.GetOkExists("data_pattern"); ok {
			tmp := dataPattern.(string)
			details.DataPattern = &tmp
		}
		if defaultMaskingFormatId, ok := s.D.GetOkExists("default_masking_format_id"); ok {
			tmp := defaultMaskingFormatId.(string)
			details.DefaultMaskingFormatId = &tmp
		}
		if namePattern, ok := s.D.GetOkExists("name_pattern"); ok {
			tmp := namePattern.(string)
			details.NamePattern = &tmp
		}
		if searchType, ok := s.D.GetOkExists("search_type"); ok {
			details.SearchType = oci_data_safe.SensitiveTypePatternSearchTypeEnum(searchType.(string))
		}
		if defaultMaskingFormatId, ok := s.D.GetOkExists("default_masking_format_id"); ok {
			tmp := defaultMaskingFormatId.(string)
			details.DefaultMaskingFormatId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if parentCategoryId, ok := s.D.GetOkExists("parent_category_id"); ok {
			tmp := parentCategoryId.(string)
			details.ParentCategoryId = &tmp
		}
		tmp := s.D.Id()
		request.SensitiveTypeId = &tmp
		if shortName, ok := s.D.GetOkExists("short_name"); ok {
			tmp := shortName.(string)
			details.ShortName = &tmp
		}
		request.UpdateSensitiveTypeDetails = details
	default:
		return fmt.Errorf("unknown entity_type '%v' was specified", entityType)
	}
	return nil
}

func (s *DataSafeSensitiveTypeResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeSensitiveTypeCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SensitiveTypeId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeSensitiveTypeCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
