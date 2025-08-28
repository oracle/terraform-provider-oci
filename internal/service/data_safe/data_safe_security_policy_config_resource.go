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

func DataSafeSecurityPolicyConfigResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSecurityPolicyConfig,
		Read:     readDataSafeSecurityPolicyConfig,
		Update:   updateDataSafeSecurityPolicyConfig,
		Delete:   deleteDataSafeSecurityPolicyConfig,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_policy_id": {
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
			"firewall_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"exclude_job": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"violation_log_auto_purge": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"time_status_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"unified_audit_policy_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"exclude_datasafe_user": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
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

func createDataSafeSecurityPolicyConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeSecurityPolicyConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSecurityPolicyConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeSecurityPolicyConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeSecurityPolicyConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SecurityPolicyConfig
	DisableNotFoundRetries bool
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.SecurityPolicyConfigLifecycleStateCreating),
	}
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.SecurityPolicyConfigLifecycleStateActive),
		string(oci_data_safe.SecurityPolicyConfigLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.SecurityPolicyConfigLifecycleStateDeleting),
	}
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.SecurityPolicyConfigLifecycleStateDeleted),
	}
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) Create() error {
	request := oci_data_safe.CreateSecurityPolicyConfigRequest{}

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

	if firewallConfig, ok := s.D.GetOkExists("firewall_config"); ok {
		if tmpList := firewallConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "firewall_config", 0)
			tmp, err := s.mapToFirewallConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FirewallConfig = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if securityPolicyId, ok := s.D.GetOkExists("security_policy_id"); ok {
		tmp := securityPolicyId.(string)
		request.SecurityPolicyId = &tmp
	}

	if unifiedAuditPolicyConfig, ok := s.D.GetOkExists("unified_audit_policy_config"); ok {
		if tmpList := unifiedAuditPolicyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "unified_audit_policy_config", 0)
			tmp, err := s.mapToUnifiedAuditPolicyConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UnifiedAuditPolicyConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateSecurityPolicyConfig(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getSecurityPolicyConfigFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) getSecurityPolicyConfigFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	securityPolicyConfigId, err := securityPolicyConfigWaitForWorkRequest(workId, "securitypolicyconfig",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, securityPolicyConfigId)
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
	s.D.SetId(*securityPolicyConfigId)

	return s.Get()
}

func securityPolicyConfigWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func securityPolicyConfigWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = securityPolicyConfigWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeSecurityPolicyConfigWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSecurityPolicyConfigWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSecurityPolicyConfigResourceCrud) Get() error {
	request := oci_data_safe.GetSecurityPolicyConfigRequest{}

	tmp := s.D.Id()
	request.SecurityPolicyConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSecurityPolicyConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityPolicyConfig
	return nil
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateSecurityPolicyConfigRequest{}

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

	if firewallConfig, ok := s.D.GetOkExists("firewall_config"); ok {
		if tmpList := firewallConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "firewall_config", 0)
			tmp, err := s.mapToFirewallConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FirewallConfig = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SecurityPolicyConfigId = &tmp

	if unifiedAuditPolicyConfig, ok := s.D.GetOkExists("unified_audit_policy_config"); ok {
		if tmpList := unifiedAuditPolicyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "unified_audit_policy_config", 0)
			tmp, err := s.mapToUnifiedAuditPolicyConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UnifiedAuditPolicyConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSecurityPolicyConfig(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSecurityPolicyConfigFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) Delete() error {
	request := oci_data_safe.DeleteSecurityPolicyConfigRequest{}

	tmp := s.D.Id()
	request.SecurityPolicyConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteSecurityPolicyConfig(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := securityPolicyConfigWaitForWorkRequest(workId, "securitypolicyconfig",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) SetData() error {
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

	if s.Res.FirewallConfig != nil {
		s.D.Set("firewall_config", []interface{}{FirewallConfigToMap(s.Res.FirewallConfig)})
	} else {
		s.D.Set("firewall_config", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SecurityPolicyId != nil {
		s.D.Set("security_policy_id", *s.Res.SecurityPolicyId)
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

	if s.Res.UnifiedAuditPolicyConfig != nil {
		s.D.Set("unified_audit_policy_config", []interface{}{UnifiedAuditPolicyConfigToMap(s.Res.UnifiedAuditPolicyConfig)})
	} else {
		s.D.Set("unified_audit_policy_config", nil)
	}

	return nil
}

func FirewallConfigToMap(obj *oci_data_safe.FirewallConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["exclude_job"] = string(obj.ExcludeJob)

	result["status"] = string(obj.Status)

	if obj.TimeStatusUpdated != nil {
		result["time_status_updated"] = obj.TimeStatusUpdated.String()
	}

	result["violation_log_auto_purge"] = string(obj.ViolationLogAutoPurge)

	return result
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) mapToFirewallConfigDetails(fieldKeyFormat string) (oci_data_safe.FirewallConfigDetails, error) {
	result := oci_data_safe.FirewallConfigDetails{}

	if excludeJob, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_job")); ok {
		result.ExcludeJob = oci_data_safe.FirewallConfigDetailsExcludeJobEnum(excludeJob.(string))
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_data_safe.FirewallConfigDetailsStatusEnum(status.(string))
	}

	if violationLogAutoPurge, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "violation_log_auto_purge")); ok {
		result.ViolationLogAutoPurge = oci_data_safe.FirewallConfigDetailsViolationLogAutoPurgeEnum(violationLogAutoPurge.(string))
	}

	return result, nil
}

func SecurityPolicyConfigSummaryToMap(obj oci_data_safe.SecurityPolicyConfigSummary) map[string]interface{} {
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

	if obj.FirewallConfig != nil {
		result["firewall_config"] = []interface{}{FirewallConfigToMap(obj.FirewallConfig)}
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.SecurityPolicyId != nil {
		result["security_policy_id"] = string(*obj.SecurityPolicyId)
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

	if obj.UnifiedAuditPolicyConfig != nil {
		result["unified_audit_policy_config"] = []interface{}{UnifiedAuditPolicyConfigToMap(obj.UnifiedAuditPolicyConfig)}
	}

	return result
}

func UnifiedAuditPolicyConfigToMap(obj *oci_data_safe.UnifiedAuditPolicyConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["exclude_datasafe_user"] = string(obj.ExcludeDatasafeUser)

	return result
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) mapToUnifiedAuditPolicyConfigDetails(fieldKeyFormat string) (oci_data_safe.UnifiedAuditPolicyConfigDetails, error) {
	result := oci_data_safe.UnifiedAuditPolicyConfigDetails{}

	if excludeDatasafeUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_datasafe_user")); ok {
		result.ExcludeDatasafeUser = oci_data_safe.UnifiedAuditPolicyConfigDetailsExcludeDatasafeUserEnum(excludeDatasafeUser.(string))
	}

	return result, nil
}

func (s *DataSafeSecurityPolicyConfigResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeSecurityPolicyConfigCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SecurityPolicyConfigId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeSecurityPolicyConfigCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSecurityPolicyConfigFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
