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

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeAuditPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeAuditPolicy,
		Read:     readDataSafeAuditPolicy,
		Update:   updateDataSafeAuditPolicy,
		Delete:   deleteDataSafeAuditPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"audit_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
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
			"provision_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"retrieve_from_target_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"audit_conditions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"audit_policy_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_conditions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"entity_names": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"entity_selection": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operation_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_data_safe_service_account_audited": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_priv_users_managed_by_data_safe": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"audit_specifications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"audit_policy_category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"audit_policy_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_policy_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"enable_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled_entities": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_created": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_enabled_for_all_users": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_seeded_in_data_safe": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_seeded_in_target": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_view_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"partially_enabled_msg": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_data_safe_service_account_excluded": {
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
			"target_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_provisioned": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_retrieved": {
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

func createDataSafeAuditPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := tfresource.CreateResource(d, sync)
	if err != nil {
		return err
	}

	if _, ok := sync.D.GetOkExists("provision_trigger"); ok {
		err := sync.ProvisionAuditPolicy()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("retrieve_from_target_trigger"); ok {
		err := sync.RetrieveAuditPolicies()
		if err != nil {
			return err
		}
	}

	if ok && compartment != *sync.Res.CompartmentId {
		err = sync.updateCompartment(compartment)
		if err != nil {
			return err
		}
		tmp := compartment.(string)
		sync.Res.CompartmentId = &tmp
		err := sync.Get()
		if err != nil {
			log.Printf("error doing a Get() after compartment update: %v", err)
		}
		err = sync.SetData()
		if err != nil {
			log.Printf("error doing a SetData() after compartment update: %v", err)
		}
	}
	return nil
}

func readDataSafeAuditPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeAuditPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if _, ok := sync.D.GetOkExists("provision_trigger"); ok && sync.D.HasChange("provision_trigger") {
		oldRaw, newRaw := sync.D.GetChange("provision_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ProvisionAuditPolicy()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("provision_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("retrieve_from_target_trigger"); ok && sync.D.HasChange("retrieve_from_target_trigger") {
		oldRaw, newRaw := sync.D.GetChange("retrieve_from_target_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.RetrieveAuditPolicies()
			if err != nil {
				return err
			}
		} else {
			sync.D.Set("retrieve_from_target_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteDataSafeAuditPolicy(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeAuditPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.AuditPolicy
	DisableNotFoundRetries bool
}

func (s *DataSafeAuditPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeAuditPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.AuditPolicyLifecycleStateCreating),
	}
}

func (s *DataSafeAuditPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.AuditPolicyLifecycleStateActive),
		string(oci_data_safe.AuditPolicyLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeAuditPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.AuditPolicyLifecycleStateDeleting),
	}
}

func (s *DataSafeAuditPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.AuditPolicyLifecycleStateDeleted),
	}
}

func (s *DataSafeAuditPolicyResourceCrud) Create() error {
	request := oci_data_safe.UpdateAuditPolicyRequest{}

	if auditPolicyId, ok := s.D.GetOkExists("audit_policy_id"); ok {
		tmp := auditPolicyId.(string)
		request.AuditPolicyId = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAuditPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "auditpolicy") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getAuditPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeAuditPolicyResourceCrud) getAuditPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	auditPolicyId, err := auditPolicyWaitForWorkRequest(workId, "auditpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*auditPolicyId)

	return s.Get()
}

func auditPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func auditPolicyWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = auditPolicyWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeAuditPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeAuditPolicyWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeAuditPolicyResourceCrud) Get() error {
	request := oci_data_safe.GetAuditPolicyRequest{}

	tmp := s.D.Id()
	request.AuditPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetAuditPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuditPolicy
	return nil
}

func (s *DataSafeAuditPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateAuditPolicyRequest{}

	tmp := s.D.Id()
	request.AuditPolicyId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAuditPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditPolicyResourceCrud) SetData() error {
	auditConditions := []interface{}{}
	for _, item1 := range s.Res.AuditConditions {
		for _, item2 := range s.Res.AuditSpecifications {
			if item1.AuditPolicyName != nil && item2.AuditPolicyName != nil &&
				*item1.AuditPolicyName == *item2.AuditPolicyName &&
				item2.IsViewOnly != nil && *item2.IsViewOnly == false {
				auditConditions = append(auditConditions, AuditConditionsToMap(item1))
			}
		}
	}
	s.D.Set("audit_conditions", auditConditions)

	auditSpecifications := []interface{}{}
	for _, item := range s.Res.AuditSpecifications {
		auditSpecifications = append(auditSpecifications, AuditSpecificationToMap(item))
	}
	s.D.Set("audit_specifications", auditSpecifications)

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

	if s.Res.IsDataSafeServiceAccountExcluded != nil {
		s.D.Set("is_data_safe_service_account_excluded", *s.Res.IsDataSafeServiceAccountExcluded)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastProvisioned != nil {
		s.D.Set("time_last_provisioned", s.Res.TimeLastProvisioned.String())
	}

	if s.Res.TimeLastRetrieved != nil {
		s.D.Set("time_last_retrieved", s.Res.TimeLastRetrieved.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataSafeAuditPolicyResourceCrud) ProvisionAuditPolicy() error {
	request := oci_data_safe.ProvisionAuditPolicyRequest{}

	idTmp := s.D.Id()
	request.AuditPolicyId = &idTmp

	if isDataSafeServiceAccountExcluded, ok := s.D.GetOkExists("is_data_safe_service_account_excluded"); ok {
		tmp := isDataSafeServiceAccountExcluded.(bool)
		request.IsDataSafeServiceAccountExcluded = &tmp
	}
	if provisionAuditConditions, ok := s.D.GetOkExists("audit_conditions"); ok {
		interfaces := provisionAuditConditions.([]interface{})
		tmp := make([]oci_data_safe.ProvisionAuditConditions, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "audit_conditions", stateDataIndex)
			converted, err := s.mapToProvisionAuditConditions(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("audit_conditions") {
			request.ProvisionAuditConditions = tmp
		}
	}
	fmt.Printf(" request = %+v\n", request)
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ProvisionAuditPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("provision_trigger")
	s.D.Set("provision_trigger", val)

	return nil
}

func (s *DataSafeAuditPolicyResourceCrud) RetrieveAuditPolicies() error {
	request := oci_data_safe.RetrieveAuditPoliciesRequest{}

	idTmp := s.D.Id()
	request.AuditPolicyId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.RetrieveAuditPolicies(context.Background(), request)

	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("retrieve_from_target_trigger")
	s.D.Set("retrieve_from_target_trigger", val)

	//s.Res = &response.AuditPolicy
	return nil
}

func AuditConditionsToMap(obj oci_data_safe.AuditConditions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuditPolicyName != nil {
		result["audit_policy_name"] = string(*obj.AuditPolicyName)
	}

	enableConditions := []interface{}{}
	for _, item := range obj.EnableConditions {
		enableConditions = append(enableConditions, EnableConditionsToMap(item))
	}
	result["enable_conditions"] = enableConditions

	if obj.IsDataSafeServiceAccountAudited != nil {
		result["is_data_safe_service_account_audited"] = bool(*obj.IsDataSafeServiceAccountAudited)
	}

	if obj.IsPrivUsersManagedByDataSafe != nil {
		result["is_priv_users_managed_by_data_safe"] = bool(*obj.IsPrivUsersManagedByDataSafe)
	}

	return result
}

func AuditPolicySummaryToMap(obj oci_data_safe.AuditPolicySummary) map[string]interface{} {
	result := map[string]interface{}{}

	auditSpecifications := []interface{}{}
	for _, item := range obj.AuditSpecifications {
		auditSpecifications = append(auditSpecifications, AuditSpecificationToMap(item))
	}
	result["audit_specifications"] = auditSpecifications

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

	if obj.IsDataSafeServiceAccountExcluded != nil {
		result["is_data_safe_service_account_excluded"] = bool(*obj.IsDataSafeServiceAccountExcluded)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastProvisioned != nil {
		result["time_last_provisioned"] = obj.TimeLastProvisioned.String()
	}

	if obj.TimeLastRetrieved != nil {
		result["time_last_retrieved"] = obj.TimeLastRetrieved.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func AuditSpecificationToMap(obj oci_data_safe.AuditSpecification) map[string]interface{} {
	result := map[string]interface{}{}

	result["audit_policy_category"] = string(obj.AuditPolicyCategory)

	if obj.AuditPolicyName != nil {
		result["audit_policy_name"] = string(*obj.AuditPolicyName)
	}

	result["database_policy_names"] = obj.DatabasePolicyNames

	result["enable_status"] = string(obj.EnableStatus)

	result["enabled_entities"] = string(obj.EnabledEntities)

	if obj.IsCreated != nil {
		result["is_created"] = bool(*obj.IsCreated)
	}

	if obj.IsEnabledForAllUsers != nil {
		result["is_enabled_for_all_users"] = bool(*obj.IsEnabledForAllUsers)
	}

	if obj.IsSeededInDataSafe != nil {
		result["is_seeded_in_data_safe"] = bool(*obj.IsSeededInDataSafe)
	}

	if obj.IsSeededInTarget != nil {
		result["is_seeded_in_target"] = bool(*obj.IsSeededInTarget)
	}

	if obj.IsViewOnly != nil {
		result["is_view_only"] = bool(*obj.IsViewOnly)
	}

	if obj.PartiallyEnabledMsg != nil {
		result["partially_enabled_msg"] = string(*obj.PartiallyEnabledMsg)
	}

	return result
}

func EnableConditionsToMap(obj oci_data_safe.EnableConditions) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_names"] = obj.EntityNames

	result["entity_selection"] = string(obj.EntitySelection)

	result["entity_type"] = string(obj.EntityType)

	result["operation_status"] = string(obj.OperationStatus)

	return result
}

func (s *DataSafeAuditPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeAuditPolicyCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AuditPolicyId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeAuditPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditPolicyResourceCrud) mapToProvisionAuditConditions(fieldKeyFormat string) (oci_data_safe.ProvisionAuditConditions, error) {
	result := oci_data_safe.ProvisionAuditConditions{}

	if auditPolicyName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "audit_policy_name")); ok {
		tmp := auditPolicyName.(string)
		result.AuditPolicyName = &tmp
	}

	if enableConditions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_conditions")); ok {
		interfaces := enableConditions.([]interface{})
		tmp := make([]oci_data_safe.EnableConditions, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "enable_conditions"), stateDataIndex)
			converted, err := s.mapToEnableConditions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "enable_conditions")) {
			result.EnableConditions = tmp
		}
	}

	//if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
	//tmp := isEnabled.(bool)
	b := false
	result.IsEnabled = &b
	//}

	if isPrivUsersManagedByDataSafe, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_priv_users_managed_by_data_safe")); ok {
		tmp := isPrivUsersManagedByDataSafe.(bool)
		result.IsPrivUsersManagedByDataSafe = &tmp
	}

	return result, nil
}
func (s *DataSafeAuditPolicyResourceCrud) mapToEnableConditions(fieldKeyFormat string) (oci_data_safe.EnableConditions, error) {
	result := oci_data_safe.EnableConditions{}

	if entityNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_names")); ok {
		interfaces := entityNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "entity_names")) {
			result.EntityNames = tmp
		}
	}

	if entitySelection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_selection")); ok {
		result.EntitySelection = oci_data_safe.EnableConditionsEntitySelectionEnum(entitySelection.(string))
	}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_data_safe.EnableConditionsEntityTypeEnum(entityType.(string))
	}

	if operationStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_status")); ok {
		result.OperationStatus = oci_data_safe.EnableConditionsOperationStatusEnum(operationStatus.(string))
	}

	return result, nil
}
