// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RecoveryProtectionPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRecoveryProtectionPolicy,
		Read:     readRecoveryProtectionPolicy,
		Update:   updateRecoveryProtectionPolicy,
		Delete:   deleteRecoveryProtectionPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"backup_retention_period_in_days": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"must_enforce_cloud_locality": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policy_locked_date_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"is_predefined_policy": {
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

func createRecoveryProtectionPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectionPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.CreateResource(d, sync)
}

func readRecoveryProtectionPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectionPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.ReadResource(sync)
}

func updateRecoveryProtectionPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectionPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteRecoveryProtectionPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectionPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type RecoveryProtectionPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_recovery.DatabaseRecoveryClient
	Res                    *oci_recovery.ProtectionPolicy
	DisableNotFoundRetries bool
}

func (s *RecoveryProtectionPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *RecoveryProtectionPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_recovery.LifecycleStateCreating),
	}
}

func (s *RecoveryProtectionPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_recovery.LifecycleStateActive),
	}
}

func (s *RecoveryProtectionPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_recovery.LifecycleStateDeleting),
	}
}

func (s *RecoveryProtectionPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_recovery.LifecycleStateDeleted),
	}
}

func (s *RecoveryProtectionPolicyResourceCrud) Create() error {
	request := oci_recovery.CreateProtectionPolicyRequest{}

	if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok {
		tmp := backupRetentionPeriodInDays.(int)
		request.BackupRetentionPeriodInDays = &tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if mustEnforceCloudLocality, ok := s.D.GetOkExists("must_enforce_cloud_locality"); ok {
		tmp := mustEnforceCloudLocality.(bool)
		request.MustEnforceCloudLocality = &tmp
	}

	if policyLockedDateTime, ok := s.D.GetOkExists("policy_locked_date_time"); ok {
		tmp := policyLockedDateTime.(string)
		request.PolicyLockedDateTime = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.CreateProtectionPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getProtectionPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *RecoveryProtectionPolicyResourceCrud) getProtectionPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_recovery.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	protectionPolicyId, err := protectionPolicyWaitForWorkRequest(workId, "protectionpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*protectionPolicyId)

	return s.Get()
}

func protectionPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "recovery", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_recovery.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func protectionPolicyWaitForWorkRequest(wId *string, entityType string, action oci_recovery.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_recovery.DatabaseRecoveryClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "recovery")
	retryPolicy.ShouldRetryOperation = protectionPolicyWorkRequestShouldRetryFunc(timeout)

	response := oci_recovery.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_recovery.OperationStatusInProgress),
			string(oci_recovery.OperationStatusAccepted),
			string(oci_recovery.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_recovery.OperationStatusSucceeded),
			string(oci_recovery.OperationStatusFailed),
			string(oci_recovery.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_recovery.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_recovery.OperationStatusFailed || response.Status == oci_recovery.OperationStatusCanceled {
		return nil, getErrorFromRecoveryProtectionPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromRecoveryProtectionPolicyWorkRequest(client *oci_recovery.DatabaseRecoveryClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_recovery.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_recovery.ListWorkRequestErrorsRequest{
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

func (s *RecoveryProtectionPolicyResourceCrud) Get() error {
	request := oci_recovery.GetProtectionPolicyRequest{}

	tmp := s.D.Id()
	request.ProtectionPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.GetProtectionPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProtectionPolicy
	return nil
}

func (s *RecoveryProtectionPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_recovery.UpdateProtectionPolicyRequest{}

	if backupRetentionPeriodInDays, ok := s.D.GetOkExists("backup_retention_period_in_days"); ok {
		tmp := backupRetentionPeriodInDays.(int)
		request.BackupRetentionPeriodInDays = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if policyLockedDateTime, ok := s.D.GetOkExists("policy_locked_date_time"); ok {
		tmp := policyLockedDateTime.(string)
		request.PolicyLockedDateTime = &tmp
	}

	tmp := s.D.Id()
	request.ProtectionPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.UpdateProtectionPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getProtectionPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *RecoveryProtectionPolicyResourceCrud) Delete() error {
	request := oci_recovery.DeleteProtectionPolicyRequest{}

	tmp := s.D.Id()
	request.ProtectionPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.DeleteProtectionPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := protectionPolicyWaitForWorkRequest(workId, "protectionpolicy",
		oci_recovery.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *RecoveryProtectionPolicyResourceCrud) SetData() error {
	if s.Res.BackupRetentionPeriodInDays != nil {
		s.D.Set("backup_retention_period_in_days", *s.Res.BackupRetentionPeriodInDays)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsPredefinedPolicy != nil {
		s.D.Set("is_predefined_policy", *s.Res.IsPredefinedPolicy)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MustEnforceCloudLocality != nil {
		s.D.Set("must_enforce_cloud_locality", *s.Res.MustEnforceCloudLocality)
	}

	if s.Res.PolicyLockedDateTime != nil {
		s.D.Set("policy_locked_date_time", *s.Res.PolicyLockedDateTime)
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

func ProtectionPolicySummaryToMap(obj oci_recovery.ProtectionPolicySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupRetentionPeriodInDays != nil {
		result["backup_retention_period_in_days"] = int(*obj.BackupRetentionPeriodInDays)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsPredefinedPolicy != nil {
		result["is_predefined_policy"] = bool(*obj.IsPredefinedPolicy)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MustEnforceCloudLocality != nil {
		result["must_enforce_cloud_locality"] = bool(*obj.MustEnforceCloudLocality)
	}

	if obj.PolicyLockedDateTime != nil {
		result["policy_locked_date_time"] = string(*obj.PolicyLockedDateTime)
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

func (s *RecoveryProtectionPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_recovery.ChangeProtectionPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ProtectionPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.ChangeProtectionPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getProtectionPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
