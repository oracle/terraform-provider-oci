// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeTargetDatabaseGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeTargetDatabaseGroup,
		Read:     readDataSafeTargetDatabaseGroup,
		Update:   updateDataSafeTargetDatabaseGroup,
		Delete:   deleteDataSafeTargetDatabaseGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"matching_criteria": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"include": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"compartments": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"id": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
												"is_include_subtree": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"defined_tags": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"target_database_ids": {
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

						// Optional
						"exclude": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"target_database_ids": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Optional

									// Computed
								},
							},
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"membership_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"membership_update_time": {
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

func createDataSafeTargetDatabaseGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeTargetDatabaseGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeTargetDatabaseGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeTargetDatabaseGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeTargetDatabaseGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.TargetDatabaseGroup
	DisableNotFoundRetries bool
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseGroupLifecycleStateCreating),
	}
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseGroupLifecycleStateActive),
		string(oci_data_safe.TargetDatabaseGroupLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseGroupLifecycleStateDeleting),
	}
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseGroupLifecycleStateDeleted),
	}
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) Create() error {
	request := oci_data_safe.CreateTargetDatabaseGroupRequest{}

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

	if matchingCriteria, ok := s.D.GetOkExists("matching_criteria"); ok {
		if tmpList := matchingCriteria.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "matching_criteria", 0)
			tmp, err := s.mapToMatchingCriteria(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MatchingCriteria = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateTargetDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getTargetDatabaseGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) getTargetDatabaseGroupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	targetDatabaseGroupId, err := targetDatabaseGroupWaitForWorkRequest(workId, "targetdatabasegroup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, targetDatabaseGroupId)
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
	s.D.SetId(*targetDatabaseGroupId)

	return s.Get()
}

func targetDatabaseGroupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func targetDatabaseGroupWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = targetDatabaseGroupWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeTargetDatabaseGroupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeTargetDatabaseGroupWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeTargetDatabaseGroupResourceCrud) Get() error {
	request := oci_data_safe.GetTargetDatabaseGroupRequest{}

	tmp := s.D.Id()
	request.TargetDatabaseGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetTargetDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TargetDatabaseGroup
	return nil
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateTargetDatabaseGroupRequest{}

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

	if matchingCriteria, ok := s.D.GetOkExists("matching_criteria"); ok {
		if tmpList := matchingCriteria.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "matching_criteria", 0)
			tmp, err := s.mapToMatchingCriteria(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MatchingCriteria = &tmp
		}
	}

	tmp := s.D.Id()
	request.TargetDatabaseGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateTargetDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTargetDatabaseGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) Delete() error {
	request := oci_data_safe.DeleteTargetDatabaseGroupRequest{}

	tmp := s.D.Id()
	request.TargetDatabaseGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteTargetDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := targetDatabaseGroupWaitForWorkRequest(workId, "targetdatabasegroup",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) SetData() error {
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MatchingCriteria != nil {
		s.D.Set("matching_criteria", []interface{}{MatchingCriteriaToMap(s.Res.MatchingCriteria)})
	} else {
		s.D.Set("matching_criteria", nil)
	}

	if s.Res.MembershipCount != nil {
		s.D.Set("membership_count", *s.Res.MembershipCount)
	}

	if s.Res.MembershipUpdateTime != nil {
		s.D.Set("membership_update_time", s.Res.MembershipUpdateTime.String())
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) mapToCompartments(fieldKeyFormat string) (oci_data_safe.Compartments, error) {
	result := oci_data_safe.Compartments{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if isIncludeSubtree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_include_subtree")); ok {
		tmp := isIncludeSubtree.(bool)
		result.IsIncludeSubtree = &tmp
	}

	return result, nil
}

func CompartmentsToMap(obj oci_data_safe.Compartments) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsIncludeSubtree != nil {
		result["is_include_subtree"] = bool(*obj.IsIncludeSubtree)
	}

	return result
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) mapToExclude(fieldKeyFormat string) (oci_data_safe.Exclude, error) {
	result := oci_data_safe.Exclude{}

	if targetDatabaseIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_database_ids")); ok {
		interfaces := targetDatabaseIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "target_database_ids")) {
			result.TargetDatabaseIds = tmp
		}
	}

	return result, nil
}

func ExcludeToMap(obj *oci_data_safe.Exclude) map[string]interface{} {
	result := map[string]interface{}{}

	result["target_database_ids"] = obj.TargetDatabaseIds

	return result
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) mapToInclude(fieldKeyFormat string) (oci_data_safe.Include, error) {
	result := oci_data_safe.Include{}

	if compartments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartments")); ok {
		interfaces := compartments.([]interface{})
		tmp := make([]oci_data_safe.Compartments, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "compartments"), stateDataIndex)
			converted, err := s.mapToCompartments(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "compartments")) {
			result.Compartments = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if targetDatabaseIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_database_ids")); ok {
		interfaces := targetDatabaseIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "target_database_ids")) {
			result.TargetDatabaseIds = tmp
		}
	}

	return result, nil
}

func IncludeToMap(obj *oci_data_safe.Include) map[string]interface{} {
	result := map[string]interface{}{}

	compartments := []interface{}{}
	for _, item := range obj.Compartments {
		compartments = append(compartments, CompartmentsToMap(item))
	}
	result["compartments"] = compartments

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	result["target_database_ids"] = obj.TargetDatabaseIds

	return result
}

func (s *DataSafeTargetDatabaseGroupResourceCrud) mapToMatchingCriteria(fieldKeyFormat string) (oci_data_safe.MatchingCriteria, error) {
	result := oci_data_safe.MatchingCriteria{}

	if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
		if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
			tmp, err := s.mapToExclude(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
			}
			result.Exclude = &tmp
		}
	}

	if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
		if tmpList := include.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
			tmp, err := s.mapToInclude(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert include, encountered error: %v", err)
			}
			result.Include = &tmp
		}
	}

	return result, nil
}

func MatchingCriteriaToMap(obj *oci_data_safe.MatchingCriteria) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Exclude != nil {
		result["exclude"] = []interface{}{ExcludeToMap(obj.Exclude)}
	}

	if obj.Include != nil {
		result["include"] = []interface{}{IncludeToMap(obj.Include)}
	}

	return result
}

func TargetDatabaseGroupSummaryToMap(obj oci_data_safe.TargetDatabaseGroupSummary) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MembershipCount != nil {
		result["membership_count"] = int(*obj.MembershipCount)
	}

	if obj.MembershipUpdateTime != nil {
		result["membership_update_time"] = obj.MembershipUpdateTime.String()
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

func (s *DataSafeTargetDatabaseGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeTargetDatabaseGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.TargetDatabaseGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeTargetDatabaseGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTargetDatabaseGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
