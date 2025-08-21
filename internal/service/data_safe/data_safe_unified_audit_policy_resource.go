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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeUnifiedAuditPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeUnifiedAuditPolicy,
		Read:     readDataSafeUnifiedAuditPolicy,
		Update:   updateDataSafeUnifiedAuditPolicy,
		Delete:   deleteDataSafeUnifiedAuditPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"conditions": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"entity_selection": {
							Type:     schema.TypeString,
							Required: true,
						},
						"entity_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ALL_USERS",
								"ATTRIBUTE_SET",
								"ROLE",
								"USER",
							}, true),
						},
						"operation_status": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"attribute_set_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"role_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"user_names": {
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
			"security_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Required: true,
			},
			"unified_audit_policy_definition_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

			// Computed
			"enabled_entities": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_seeded": {
				Type:     schema.TypeBool,
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

func createDataSafeUnifiedAuditPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUnifiedAuditPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeUnifiedAuditPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUnifiedAuditPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeUnifiedAuditPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUnifiedAuditPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeUnifiedAuditPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUnifiedAuditPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeUnifiedAuditPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.UnifiedAuditPolicy
	DisableNotFoundRetries bool
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.UnifiedAuditPolicyLifecycleStateCreating),
	}
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.UnifiedAuditPolicyLifecycleStateActive),
		string(oci_data_safe.UnifiedAuditPolicyLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.UnifiedAuditPolicyLifecycleStateDeleting),
	}
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.UnifiedAuditPolicyLifecycleStateDeleted),
	}
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) Create() error {
	request := oci_data_safe.CreateUnifiedAuditPolicyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if conditions, ok := s.D.GetOkExists("conditions"); ok {
		interfaces := conditions.([]interface{})
		tmp := make([]oci_data_safe.PolicyCondition, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "conditions", stateDataIndex)
			converted, err := s.mapToPolicyCondition(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("conditions") {
			request.Conditions = tmp
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

	if securityPolicyId, ok := s.D.GetOkExists("security_policy_id"); ok {
		tmp := securityPolicyId.(string)
		request.SecurityPolicyId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_data_safe.UnifiedAuditPolicyStatusEnum(status.(string))
	}

	if unifiedAuditPolicyDefinitionId, ok := s.D.GetOkExists("unified_audit_policy_definition_id"); ok {
		tmp := unifiedAuditPolicyDefinitionId.(string)
		request.UnifiedAuditPolicyDefinitionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateUnifiedAuditPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getUnifiedAuditPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) getUnifiedAuditPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	unifiedAuditPolicyId, err := unifiedAuditPolicyWaitForWorkRequest(workId, "unifiedauditpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, unifiedAuditPolicyId)
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
	s.D.SetId(*unifiedAuditPolicyId)

	return s.Get()
}

func unifiedAuditPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func unifiedAuditPolicyWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = unifiedAuditPolicyWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeUnifiedAuditPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeUnifiedAuditPolicyWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeUnifiedAuditPolicyResourceCrud) Get() error {
	request := oci_data_safe.GetUnifiedAuditPolicyRequest{}

	tmp := s.D.Id()
	request.UnifiedAuditPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetUnifiedAuditPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UnifiedAuditPolicy
	return nil
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateUnifiedAuditPolicyRequest{}

	if conditions, ok := s.D.GetOkExists("conditions"); ok {
		interfaces := conditions.([]interface{})
		tmp := make([]oci_data_safe.PolicyCondition, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "conditions", stateDataIndex)
			converted, err := s.mapToPolicyCondition(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("conditions") {
			request.Conditions = tmp
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

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_data_safe.UnifiedAuditPolicyStatusEnum(status.(string))
	}

	tmp := s.D.Id()
	request.UnifiedAuditPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateUnifiedAuditPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUnifiedAuditPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) Delete() error {
	request := oci_data_safe.DeleteUnifiedAuditPolicyRequest{}

	tmp := s.D.Id()
	request.UnifiedAuditPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteUnifiedAuditPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := unifiedAuditPolicyWaitForWorkRequest(workId, "unifiedauditpolicy",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	conditions := []interface{}{}
	for _, item := range s.Res.Conditions {
		conditions = append(conditions, PolicyConditionToMap(item))
	}
	s.D.Set("conditions", conditions)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("enabled_entities", s.Res.EnabledEntities)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSeeded != nil {
		s.D.Set("is_seeded", *s.Res.IsSeeded)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SecurityPolicyId != nil {
		s.D.Set("security_policy_id", *s.Res.SecurityPolicyId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UnifiedAuditPolicyDefinitionId != nil {
		s.D.Set("unified_audit_policy_definition_id", *s.Res.UnifiedAuditPolicyDefinitionId)
	}

	return nil
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) mapToPolicyCondition(fieldKeyFormat string) (oci_data_safe.PolicyCondition, error) {
	var baseObject oci_data_safe.PolicyCondition
	//discriminator
	entityTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type"))
	var entityType string
	if ok {
		entityType = entityTypeRaw.(string)
	} else {
		entityType = "" // default value
	}
	switch strings.ToLower(entityType) {
	case strings.ToLower("ALL_USERS"):
		details := oci_data_safe.AllUserCondition{}
		if entitySelection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_selection")); ok {
			details.EntitySelection = oci_data_safe.PolicyConditionEntitySelectionEnum(entitySelection.(string))
		}
		if operationStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_status")); ok {
			details.OperationStatus = oci_data_safe.PolicyConditionOperationStatusEnum(operationStatus.(string))
		}
		baseObject = details
	case strings.ToLower("ATTRIBUTE_SET"):
		details := oci_data_safe.AttributeSetCondition{}
		if attributeSetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_set_id")); ok {
			tmp := attributeSetId.(string)
			details.AttributeSetId = &tmp
		}
		if entitySelection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_selection")); ok {
			details.EntitySelection = oci_data_safe.PolicyConditionEntitySelectionEnum(entitySelection.(string))
		}
		if operationStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_status")); ok {
			details.OperationStatus = oci_data_safe.PolicyConditionOperationStatusEnum(operationStatus.(string))
		}
		baseObject = details
	case strings.ToLower("ROLE"):
		details := oci_data_safe.RoleCondition{}
		if roleNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role_names")); ok {
			interfaces := roleNames.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "role_names")) {
				details.RoleNames = tmp
			}
		}
		if entitySelection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_selection")); ok {
			details.EntitySelection = oci_data_safe.PolicyConditionEntitySelectionEnum(entitySelection.(string))
		}
		if operationStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_status")); ok {
			details.OperationStatus = oci_data_safe.PolicyConditionOperationStatusEnum(operationStatus.(string))
		}
		baseObject = details
	case strings.ToLower("USER"):
		details := oci_data_safe.UserCondition{}
		if userNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_names")); ok {
			interfaces := userNames.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "user_names")) {
				details.UserNames = tmp
			}
		}
		if entitySelection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_selection")); ok {
			details.EntitySelection = oci_data_safe.PolicyConditionEntitySelectionEnum(entitySelection.(string))
		}
		if operationStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_status")); ok {
			details.OperationStatus = oci_data_safe.PolicyConditionOperationStatusEnum(operationStatus.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown entity_type '%v' was specified", entityType)
	}
	return baseObject, nil
}

func PolicyConditionToMap(obj oci_data_safe.PolicyCondition) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_data_safe.AllUserCondition:
		result["entity_type"] = "ALL_USERS"

		result["entity_selection"] = string(v.EntitySelection)

		result["operation_status"] = string(v.OperationStatus)
	case oci_data_safe.AttributeSetCondition:
		result["entity_type"] = "ATTRIBUTE_SET"

		if v.AttributeSetId != nil {
			result["attribute_set_id"] = string(*v.AttributeSetId)
		}

		result["entity_selection"] = string(v.EntitySelection)

		result["operation_status"] = string(v.OperationStatus)
	case oci_data_safe.RoleCondition:
		result["entity_type"] = "ROLE"

		result["role_names"] = v.RoleNames

		result["entity_selection"] = string(v.EntitySelection)

		result["operation_status"] = string(v.OperationStatus)
	case oci_data_safe.UserCondition:
		result["entity_type"] = "USER"

		result["user_names"] = v.UserNames

		result["entity_selection"] = string(v.EntitySelection)

		result["operation_status"] = string(v.OperationStatus)
	default:
		log.Printf("[WARN] Received 'entity_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func UnifiedAuditPolicySummaryToMap(obj oci_data_safe.UnifiedAuditPolicySummary) map[string]interface{} {
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

	result["enabled_entities"] = string(obj.EnabledEntities)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsSeeded != nil {
		result["is_seeded"] = bool(*obj.IsSeeded)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.SecurityPolicyId != nil {
		result["security_policy_id"] = string(*obj.SecurityPolicyId)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.UnifiedAuditPolicyDefinitionId != nil {
		result["unified_audit_policy_definition_id"] = string(*obj.UnifiedAuditPolicyDefinitionId)
	}

	return result
}

func (s *DataSafeUnifiedAuditPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeUnifiedAuditPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.UnifiedAuditPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeUnifiedAuditPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUnifiedAuditPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
