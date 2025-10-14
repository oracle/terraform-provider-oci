// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package disaster_recovery

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DisasterRecoveryAutomaticDrConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDisasterRecoveryAutomaticDrConfigurationWithContext,
		ReadContext:   readDisasterRecoveryAutomaticDrConfigurationWithContext,
		UpdateContext: updateDisasterRecoveryAutomaticDrConfigurationWithContext,
		DeleteContext: deleteDisasterRecoveryAutomaticDrConfigurationWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dr_protection_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"members": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"member_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"member_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AUTONOMOUS_CONTAINER_DATABASE",
								"AUTONOMOUS_DATABASE",
								"DATABASE",
							}, true),
						},

						// Optional
						"is_auto_failover_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_auto_switchover_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Optional
			"default_failover_dr_plan_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_switchover_dr_plan_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_automatic_dr_execution_submit_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_automatic_dr_execution_submit_status": {
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
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_automatic_dr_execution_submit_attempt": {
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

func createDisasterRecoveryAutomaticDrConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DisasterRecoveryAutomaticDrConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDisasterRecoveryAutomaticDrConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DisasterRecoveryAutomaticDrConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDisasterRecoveryAutomaticDrConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DisasterRecoveryAutomaticDrConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDisasterRecoveryAutomaticDrConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DisasterRecoveryAutomaticDrConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DisasterRecoveryAutomaticDrConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_disaster_recovery.DisasterRecoveryClient
	Res                    *oci_disaster_recovery.AutomaticDrConfiguration
	DisableNotFoundRetries bool
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_disaster_recovery.AutomaticDrConfigurationLifecycleStateCreating),
	}
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_disaster_recovery.AutomaticDrConfigurationLifecycleStateActive),
		string(oci_disaster_recovery.AutomaticDrConfigurationLifecycleStateNeedsAttention),
		string(oci_disaster_recovery.AutomaticDrConfigurationLifecycleStateFailed),
	}
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_disaster_recovery.AutomaticDrConfigurationLifecycleStateDeleting),
	}
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_disaster_recovery.AutomaticDrConfigurationLifecycleStateDeleted),
	}
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_disaster_recovery.CreateAutomaticDrConfigurationRequest{}

	if defaultFailoverDrPlanId, ok := s.D.GetOkExists("default_failover_dr_plan_id"); ok {
		tmp := defaultFailoverDrPlanId.(string)
		request.DefaultFailoverDrPlanId = &tmp
	}

	if defaultSwitchoverDrPlanId, ok := s.D.GetOkExists("default_switchover_dr_plan_id"); ok {
		tmp := defaultSwitchoverDrPlanId.(string)
		request.DefaultSwitchoverDrPlanId = &tmp
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

	if drProtectionGroupId, ok := s.D.GetOkExists("dr_protection_group_id"); ok {
		tmp := drProtectionGroupId.(string)
		request.DrProtectionGroupId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if members, ok := s.D.GetOkExists("members"); ok {
		interfaces := members.([]interface{})
		tmp := make([]oci_disaster_recovery.CreateAutomaticDrConfigurationMemberDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "members", stateDataIndex)
			converted, err := s.mapToCreateAutomaticDrConfigurationMemberDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("members") {
			request.Members = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.CreateAutomaticDrConfiguration(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAutomaticDrConfigurationFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) getAutomaticDrConfigurationFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_disaster_recovery.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	automaticDrConfigurationId, err := automaticDrConfigurationWaitForWorkRequest(ctx, workId, "automaticdrconfiguration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, automaticDrConfigurationId)
		// Create a retry policy that doesn't retry on IncorrectState (work request already in terminal state)
		cancelRetryPolicy := tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")
		cancelRetryPolicy.ShouldRetryOperation = func(response oci_common.OCIOperationResponse) bool {
			if failure, isServiceError := oci_common.IsServiceError(response.Error); isServiceError {
				// Don't retry if work request is already in a terminal state
				if failure.GetCode() == "IncorrectState" {
					return false
				}
			}
			return false // Don't retry cancel operations by default
		}
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
			oci_disaster_recovery.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: cancelRetryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*automaticDrConfigurationId)

	return s.GetWithContext(ctx)
}

func automaticDrConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "disaster_recovery", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_disaster_recovery.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func automaticDrConfigurationWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_disaster_recovery.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_disaster_recovery.DisasterRecoveryClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "disaster_recovery")
	retryPolicy.ShouldRetryOperation = automaticDrConfigurationWorkRequestShouldRetryFunc(timeout)

	response := oci_disaster_recovery.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_disaster_recovery.OperationStatusInProgress),
			string(oci_disaster_recovery.OperationStatusAccepted),
			string(oci_disaster_recovery.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_disaster_recovery.OperationStatusSucceeded),
			string(oci_disaster_recovery.OperationStatusFailed),
			string(oci_disaster_recovery.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_disaster_recovery.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_disaster_recovery.OperationStatusFailed || response.Status == oci_disaster_recovery.OperationStatusCanceled {
		return nil, getErrorFromDisasterRecoveryAutomaticDrConfigurationWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDisasterRecoveryAutomaticDrConfigurationWorkRequest(ctx context.Context, client *oci_disaster_recovery.DisasterRecoveryClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_disaster_recovery.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_disaster_recovery.ListWorkRequestErrorsRequest{
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

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_disaster_recovery.GetAutomaticDrConfigurationRequest{}

	tmp := s.D.Id()
	request.AutomaticDrConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.GetAutomaticDrConfiguration(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.AutomaticDrConfiguration
	return nil
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_disaster_recovery.UpdateAutomaticDrConfigurationRequest{}

	tmp := s.D.Id()
	request.AutomaticDrConfigurationId = &tmp

	if defaultFailoverDrPlanId, ok := s.D.GetOkExists("default_failover_dr_plan_id"); ok {
		tmp := defaultFailoverDrPlanId.(string)
		request.DefaultFailoverDrPlanId = &tmp
	}

	if defaultSwitchoverDrPlanId, ok := s.D.GetOkExists("default_switchover_dr_plan_id"); ok {
		tmp := defaultSwitchoverDrPlanId.(string)
		request.DefaultSwitchoverDrPlanId = &tmp
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

	if members, ok := s.D.GetOkExists("members"); ok {
		interfaces := members.([]interface{})
		tmp := make([]oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "members", stateDataIndex)
			converted, err := s.mapToUpdateAutomaticDrConfigurationMemberDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("members") {
			request.Members = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.UpdateAutomaticDrConfiguration(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAutomaticDrConfigurationFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_disaster_recovery.DeleteAutomaticDrConfigurationRequest{}

	tmp := s.D.Id()
	request.AutomaticDrConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	_, err := s.Client.DeleteAutomaticDrConfiguration(ctx, request)
	return err
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultFailoverDrPlanId != nil {
		s.D.Set("default_failover_dr_plan_id", *s.Res.DefaultFailoverDrPlanId)
	}

	if s.Res.DefaultSwitchoverDrPlanId != nil {
		s.D.Set("default_switchover_dr_plan_id", *s.Res.DefaultSwitchoverDrPlanId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrProtectionGroupId != nil {
		s.D.Set("dr_protection_group_id", *s.Res.DrProtectionGroupId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LastAutomaticDrExecutionSubmitDetails != nil {
		s.D.Set("last_automatic_dr_execution_submit_details", *s.Res.LastAutomaticDrExecutionSubmitDetails)
	}

	s.D.Set("last_automatic_dr_execution_submit_status", s.Res.LastAutomaticDrExecutionSubmitStatus)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	members := []interface{}{}
	for _, item := range s.Res.Members {
		members = append(members, AutomaticDrConfigurationMemberToMap(item))
	}
	s.D.Set("members", members)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastAutomaticDrExecutionSubmitAttempt != nil {
		s.D.Set("time_last_automatic_dr_execution_submit_attempt", s.Res.TimeLastAutomaticDrExecutionSubmitAttempt.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func AutomaticDrConfigurationSummaryToMap(obj oci_disaster_recovery.AutomaticDrConfigurationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefaultFailoverDrPlanId != nil {
		result["default_failover_dr_plan_id"] = string(*obj.DefaultFailoverDrPlanId)
	}

	if obj.DefaultSwitchoverDrPlanId != nil {
		result["default_switchover_dr_plan_id"] = string(*obj.DefaultSwitchoverDrPlanId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DrProtectionGroupId != nil {
		result["dr_protection_group_id"] = string(*obj.DrProtectionGroupId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) mapToCreateAutomaticDrConfigurationMemberDetails(fieldKeyFormat string) (oci_disaster_recovery.CreateAutomaticDrConfigurationMemberDetails, error) {
	var baseObject oci_disaster_recovery.CreateAutomaticDrConfigurationMemberDetails
	//discriminator
	memberTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_type"))
	var memberType string
	if ok {
		memberType = memberTypeRaw.(string)
	} else {
		memberType = "" // default value
	}
	switch strings.ToLower(memberType) {
	case strings.ToLower("AUTONOMOUS_CONTAINER_DATABASE"):
		details := oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberAutonomousContainerDatabaseDetails{}
		if isAutoFailoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_failover_enabled")); ok {
			tmp := isAutoFailoverEnabled.(bool)
			details.IsAutoFailoverEnabled = &tmp
		}
		if isAutoSwitchoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_switchover_enabled")); ok {
			tmp := isAutoSwitchoverEnabled.(bool)
			details.IsAutoSwitchoverEnabled = &tmp
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("AUTONOMOUS_DATABASE"):
		details := oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberAutonomousDatabaseDetails{}
		if isAutoFailoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_failover_enabled")); ok {
			tmp := isAutoFailoverEnabled.(bool)
			details.IsAutoFailoverEnabled = &tmp
		}
		if isAutoSwitchoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_switchover_enabled")); ok {
			tmp := isAutoSwitchoverEnabled.(bool)
			details.IsAutoSwitchoverEnabled = &tmp
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("DATABASE"):
		details := oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberDatabaseDetails{}
		if isAutoFailoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_failover_enabled")); ok {
			tmp := isAutoFailoverEnabled.(bool)
			details.IsAutoFailoverEnabled = &tmp
		}
		if isAutoSwitchoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_switchover_enabled")); ok {
			tmp := isAutoSwitchoverEnabled.(bool)
			details.IsAutoSwitchoverEnabled = &tmp
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown member_type '%v' was specified", memberType)
	}
	return baseObject, nil
}

func (s *DisasterRecoveryAutomaticDrConfigurationResourceCrud) mapToUpdateAutomaticDrConfigurationMemberDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberDetails, error) {
	var baseObject oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberDetails
	//discriminator
	memberTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_type"))
	var memberType string
	if ok {
		memberType = memberTypeRaw.(string)
	} else {
		memberType = "" // default value
	}
	switch strings.ToLower(memberType) {
	case strings.ToLower("AUTONOMOUS_CONTAINER_DATABASE"):
		details := oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberAutonomousContainerDatabaseDetails{}
		if isAutoFailoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_failover_enabled")); ok {
			tmp := isAutoFailoverEnabled.(bool)
			details.IsAutoFailoverEnabled = &tmp
		}
		if isAutoSwitchoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_switchover_enabled")); ok {
			tmp := isAutoSwitchoverEnabled.(bool)
			details.IsAutoSwitchoverEnabled = &tmp
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("AUTONOMOUS_DATABASE"):
		details := oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberAutonomousDatabaseDetails{}
		if isAutoFailoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_failover_enabled")); ok {
			tmp := isAutoFailoverEnabled.(bool)
			details.IsAutoFailoverEnabled = &tmp
		}
		if isAutoSwitchoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_switchover_enabled")); ok {
			tmp := isAutoSwitchoverEnabled.(bool)
			details.IsAutoSwitchoverEnabled = &tmp
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	case strings.ToLower("DATABASE"):
		details := oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberDatabaseDetails{}
		if isAutoFailoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_failover_enabled")); ok {
			tmp := isAutoFailoverEnabled.(bool)
			details.IsAutoFailoverEnabled = &tmp
		}
		if isAutoSwitchoverEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_switchover_enabled")); ok {
			tmp := isAutoSwitchoverEnabled.(bool)
			details.IsAutoSwitchoverEnabled = &tmp
		}
		if memberId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "member_id")); ok {
			tmp := memberId.(string)
			details.MemberId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown member_type '%v' was specified", memberType)
	}
	return baseObject, nil
}

func AutomaticDrConfigurationMemberToMap(obj oci_disaster_recovery.AutomaticDrConfigurationMember) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberAutonomousContainerDatabaseDetails:
		result["member_type"] = "AUTONOMOUS_CONTAINER_DATABASE"

		if v.IsAutoFailoverEnabled != nil {
			result["is_auto_failover_enabled"] = bool(*v.IsAutoFailoverEnabled)
		}

		if v.IsAutoSwitchoverEnabled != nil {
			result["is_auto_switchover_enabled"] = bool(*v.IsAutoSwitchoverEnabled)
		}

		if v.MemberId != nil {
			result["member_id"] = string(*v.MemberId)
		}
	case oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberAutonomousDatabaseDetails:
		result["member_type"] = "AUTONOMOUS_DATABASE"

		if v.IsAutoFailoverEnabled != nil {
			result["is_auto_failover_enabled"] = bool(*v.IsAutoFailoverEnabled)
		}

		if v.IsAutoSwitchoverEnabled != nil {
			result["is_auto_switchover_enabled"] = bool(*v.IsAutoSwitchoverEnabled)
		}

		if v.MemberId != nil {
			result["member_id"] = string(*v.MemberId)
		}
	case oci_disaster_recovery.UpdateAutomaticDrConfigurationMemberDatabaseDetails:
		result["member_type"] = "DATABASE"

		if v.IsAutoFailoverEnabled != nil {
			result["is_auto_failover_enabled"] = bool(*v.IsAutoFailoverEnabled)
		}

		if v.IsAutoSwitchoverEnabled != nil {
			result["is_auto_switchover_enabled"] = bool(*v.IsAutoSwitchoverEnabled)
		}

		if v.MemberId != nil {
			result["member_id"] = string(*v.MemberId)
		}
	default:
		log.Printf("[WARN] Received 'member_type' of unknown type %v", obj)
		return nil
	}

	return result
}
