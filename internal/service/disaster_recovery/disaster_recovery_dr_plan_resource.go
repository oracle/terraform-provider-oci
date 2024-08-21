// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package disaster_recovery

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DisasterRecoveryDrPlanResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDisasterRecoveryDrPlan,
		Read:     readDisasterRecoveryDrPlan,
		Update:   updateDisasterRecoveryDrPlan,
		Delete:   deleteDisasterRecoveryDrPlan,
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
			"type": {
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
			"life_cycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_dr_protection_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plan_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_pause_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"steps": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"error_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"group_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"member_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"timeout": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_defined_step": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"function_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"function_region": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"object_storage_script_location": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"bucket": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"namespace": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"object": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"request_body": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"run_as_user": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"run_on_instance_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"run_on_instance_region": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"script_command": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"step_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createDisasterRecoveryDrPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.CreateResource(d, sync)
}

func readDisasterRecoveryDrPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.ReadResource(sync)
}

func updateDisasterRecoveryDrPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDisasterRecoveryDrPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DisasterRecoveryDrPlanResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_disaster_recovery.DisasterRecoveryClient
	Res                    *oci_disaster_recovery.DrPlan
	DisableNotFoundRetries bool
}

func (s *DisasterRecoveryDrPlanResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DisasterRecoveryDrPlanResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_disaster_recovery.DrPlanLifecycleStateCreating),
	}
}

func (s *DisasterRecoveryDrPlanResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_disaster_recovery.DrPlanLifecycleStateActive),
		string(oci_disaster_recovery.DrPlanLifecycleStateNeedsAttention),
	}
}

func (s *DisasterRecoveryDrPlanResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_disaster_recovery.DrPlanLifecycleStateDeleting),
	}
}

func (s *DisasterRecoveryDrPlanResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_disaster_recovery.DrPlanLifecycleStateDeleted),
	}
}

func (s *DisasterRecoveryDrPlanResourceCrud) Create() error {
	request := oci_disaster_recovery.CreateDrPlanRequest{}

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

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_disaster_recovery.DrPlanTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.CreateDrPlan(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDrPlanFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DisasterRecoveryDrPlanResourceCrud) getDrPlanFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_disaster_recovery.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	drPlanId, err := drPlanWaitForWorkRequest(workId, "drPlan",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, drPlanId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_disaster_recovery.CancelWorkRequestRequest{
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
	s.D.SetId(*drPlanId)

	return s.Get()
}

func drPlanWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func drPlanWaitForWorkRequest(wId *string, entityType string, action oci_disaster_recovery.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_disaster_recovery.DisasterRecoveryClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "disaster_recovery")
	retryPolicy.ShouldRetryOperation = drPlanWorkRequestShouldRetryFunc(timeout)

	response := oci_disaster_recovery.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
			response, err = client.GetWorkRequest(context.Background(),
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
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_disaster_recovery.OperationStatusFailed || response.Status == oci_disaster_recovery.OperationStatusCanceled {
		return nil, getErrorFromDisasterRecoveryDrPlanWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDisasterRecoveryDrPlanWorkRequest(client *oci_disaster_recovery.DisasterRecoveryClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_disaster_recovery.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
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

func (s *DisasterRecoveryDrPlanResourceCrud) Get() error {
	request := oci_disaster_recovery.GetDrPlanRequest{}

	tmp := s.D.Id()
	request.DrPlanId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.GetDrPlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrPlan
	return nil
}

func (s *DisasterRecoveryDrPlanResourceCrud) Update() error {
	request := oci_disaster_recovery.UpdateDrPlanRequest{}

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

	tmp := s.D.Id()
	request.DrPlanId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if planGroups, ok := s.D.GetOkExists("plan_groups"); ok {
		interfaces := planGroups.([]interface{})
		tmp := make([]oci_disaster_recovery.UpdateDrPlanGroupDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "plan_groups", stateDataIndex)
			converted, err := s.mapToUpdateDrPlanGroupDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("plan_groups") {
			request.PlanGroups = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.UpdateDrPlan(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDrPlanFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DisasterRecoveryDrPlanResourceCrud) Delete() error {
	request := oci_disaster_recovery.DeleteDrPlanRequest{}

	tmp := s.D.Id()
	request.DrPlanId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	_, err := s.Client.DeleteDrPlan(context.Background(), request)
	return err
}

func (s *DisasterRecoveryDrPlanResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.LifeCycleDetails != nil {
		s.D.Set("life_cycle_details", *s.Res.LifeCycleDetails)
	}

	if s.Res.PeerDrProtectionGroupId != nil {
		s.D.Set("peer_dr_protection_group_id", *s.Res.PeerDrProtectionGroupId)
	}

	if s.Res.PeerRegion != nil {
		s.D.Set("peer_region", *s.Res.PeerRegion)
	}

	planGroups := []interface{}{}
	for _, item := range s.Res.PlanGroups {
		planGroups = append(planGroups, DrPlanGroupToMap(item))
	}
	s.D.Set("plan_groups", planGroups)

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

	s.D.Set("type", s.Res.Type)

	return nil
}

func (s *DisasterRecoveryDrPlanResourceCrud) mapToDrPlanGroup(fieldKeyFormat string) (oci_disaster_recovery.DrPlanGroup, error) {
	result := oci_disaster_recovery.DrPlanGroup{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if steps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "steps")); ok {
		interfaces := steps.([]interface{})
		tmp := make([]oci_disaster_recovery.DrPlanStep, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "steps"), stateDataIndex)
			converted, err := s.mapToDrPlanStep(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "steps")) {
			result.Steps = tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_disaster_recovery.DrPlanGroupTypeEnum(type_.(string))
	}

	return result, nil
}

func (s *DisasterRecoveryDrPlanResourceCrud) mapToUpdateDrPlanGroupDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateDrPlanGroupDetails, error) {
	result := oci_disaster_recovery.UpdateDrPlanGroupDetails{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	/*
		if isPauseEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pause_enabled")); ok {
			tmp := isPauseEnabled.(bool)
			result.IsPauseEnabled = &tmp
		}*/

	if steps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "steps")); ok {
		interfaces := steps.([]interface{})
		tmp := make([]oci_disaster_recovery.UpdateDrPlanStepDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "steps"), stateDataIndex)
			converted, err := s.mapToUpdateDrPlanStepDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "steps")) {
			result.Steps = tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_disaster_recovery.DrPlanGroupTypeEnum(type_.(string))
	}

	return result, nil
}

func (s *DisasterRecoveryDrPlanResourceCrud) mapToUpdateDrPlanStepDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateDrPlanStepDetails, error) {
	result := oci_disaster_recovery.UpdateDrPlanStepDetails{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if errorMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "error_mode")); ok {
		result.ErrorMode = oci_disaster_recovery.DrPlanStepErrorModeEnum(errorMode.(string))
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if timeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout")); ok {
		tmp := timeout.(int)
		result.Timeout = &tmp
	}

	if userDefinedStep, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_defined_step")); ok {
		if tmpList := userDefinedStep.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "user_defined_step"), 0)
			tmp, err := s.mapToUpdateDrPlanUserDefinedStepDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert user_defined_step, encountered error: %v", err)
			}
			result.UserDefinedStep = tmp
		}
	}

	return result, nil
}

func DrPlanGroupToMap(obj oci_disaster_recovery.DrPlanGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsPauseEnabled != nil {
		result["is_pause_enabled"] = bool(*obj.IsPauseEnabled)
	}

	steps := []interface{}{}
	for _, item := range obj.Steps {
		steps = append(steps, DrPlanStepToMap(item))
	}
	result["steps"] = steps

	result["type"] = string(obj.Type)

	return result
}

func (s *DisasterRecoveryDrPlanResourceCrud) mapToDrPlanStep(fieldKeyFormat string) (oci_disaster_recovery.DrPlanStep, error) {
	result := oci_disaster_recovery.DrPlanStep{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if errorMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "error_mode")); ok {
		result.ErrorMode = oci_disaster_recovery.DrPlanStepErrorModeEnum(errorMode.(string))
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if timeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout")); ok {
		tmp := timeout.(int)
		result.Timeout = &tmp
	}

	if userDefinedStep, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_defined_step")); ok {
		if tmpList := userDefinedStep.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "user_defined_step"), 0)
			tmp, err := s.mapToDrPlanUserDefinedStep(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert user_defined_step, encountered error: %v", err)
			}
			result.UserDefinedStep = tmp
		}
	}

	return result, nil
}

func DrPlanStepToMap(obj oci_disaster_recovery.DrPlanStep) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["error_mode"] = string(obj.ErrorMode)

	if obj.GroupId != nil {
		result["group_id"] = string(*obj.GroupId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.MemberId != nil {
		result["member_id"] = string(*obj.MemberId)
	}

	if obj.Timeout != nil {
		result["timeout"] = int(*obj.Timeout)
	}

	result["type"] = string(obj.Type)

	if obj.UserDefinedStep != nil {
		userDefinedStepArray := []interface{}{}
		if userDefinedStepMap := DrPlanUserDefinedStepToMap(&obj.UserDefinedStep); userDefinedStepMap != nil {
			userDefinedStepArray = append(userDefinedStepArray, userDefinedStepMap)
		}
		result["user_defined_step"] = userDefinedStepArray
	}

	return result
}

func DrPlanSummaryToMap(obj oci_disaster_recovery.DrPlanSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
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

	if obj.LifeCycleDetails != nil {
		result["life_cycle_details"] = string(*obj.LifeCycleDetails)
	}

	if obj.PeerDrProtectionGroupId != nil {
		result["peer_dr_protection_group_id"] = string(*obj.PeerDrProtectionGroupId)
	}

	if obj.PeerRegion != nil {
		result["peer_region"] = string(*obj.PeerRegion)
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

	result["type"] = string(obj.Type)

	return result
}

func (s *DisasterRecoveryDrPlanResourceCrud) mapToDrPlanUserDefinedStep(fieldKeyFormat string) (oci_disaster_recovery.DrPlanUserDefinedStep, error) {
	var baseObject oci_disaster_recovery.DrPlanUserDefinedStep
	//discriminator
	stepTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_type"))
	var stepType string
	if ok {
		stepType = stepTypeRaw.(string)
	} else {
		stepType = "" // default value
	}
	switch strings.ToLower(stepType) {
	case strings.ToLower("INVOKE_FUNCTION"):
		details := oci_disaster_recovery.UpdateInvokeFunctionUserDefinedStepDetails{}
		if functionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "function_id")); ok {
			tmp := functionId.(string)
			details.FunctionId = &tmp
		}
		if requestBody, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_body")); ok {
			tmp := requestBody.(string)
			details.RequestBody = &tmp
		}
		baseObject = details
	case strings.ToLower("INVOKE_FUNCTION_PRECHECK"):
		details := oci_disaster_recovery.UpdateInvokeFunctionPrecheckStepDetails{}
		baseObject = details
	case strings.ToLower("RUN_LOCAL_SCRIPT"):
		details := oci_disaster_recovery.UpdateRunLocalScriptUserDefinedStepDetails{}
		if runAsUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_as_user")); ok {
			tmp := runAsUser.(string)
			details.RunAsUser = &tmp
		}
		if runOnInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_on_instance_id")); ok {
			tmp := runOnInstanceId.(string)
			details.RunOnInstanceId = &tmp
		}
		if scriptCommand, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "script_command")); ok {
			tmp := scriptCommand.(string)
			details.ScriptCommand = &tmp
		}
		baseObject = details
	case strings.ToLower("RUN_LOCAL_SCRIPT_PRECHECK"):
		details := oci_disaster_recovery.UpdateLocalScriptPrecheckStepDetails{}
		baseObject = details
	case strings.ToLower("RUN_OBJECTSTORE_SCRIPT"):
		details := oci_disaster_recovery.UpdateRunObjectStoreScriptUserDefinedStepDetails{}
		if objectStorageScriptLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_script_location")); ok {
			if tmpList := objectStorageScriptLocation.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_script_location"), 0)
				tmp, err := s.mapToUpdateObjectStorageScriptLocationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_script_location, encountered error: %v", err)
				}
				details.ObjectStorageScriptLocation = &tmp
			}
		}
		if runOnInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_on_instance_id")); ok {
			tmp := runOnInstanceId.(string)
			details.RunOnInstanceId = &tmp
		}
		baseObject = details
	case strings.ToLower("RUN_OBJECTSTORE_SCRIPT_PRECHECK"):
		details := oci_disaster_recovery.UpdateObjectStoreScriptPrecheckStepDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown step_type '%v' was specified", stepType)
	}
	return baseObject, nil
}

func (s *DisasterRecoveryDrPlanResourceCrud) mapToUpdateDrPlanUserDefinedStepDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateDrPlanUserDefinedStepDetails, error) {
	var baseObject oci_disaster_recovery.UpdateDrPlanUserDefinedStepDetails
	//discriminator
	stepTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_type"))
	var stepType string
	if ok {
		stepType = stepTypeRaw.(string)
	} else {
		stepType = "" // default value
	}
	switch strings.ToLower(stepType) {
	case strings.ToLower("INVOKE_FUNCTION"):
		details := oci_disaster_recovery.UpdateInvokeFunctionUserDefinedStepDetails{}
		if functionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "function_id")); ok {
			tmp := functionId.(string)
			details.FunctionId = &tmp
		}
		if requestBody, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_body")); ok {
			tmp := requestBody.(string)
			details.RequestBody = &tmp
		}
		baseObject = details
	case strings.ToLower("INVOKE_FUNCTION_PRECHECK"):
		details := oci_disaster_recovery.UpdateInvokeFunctionPrecheckStepDetails{}
		baseObject = details
	case strings.ToLower("RUN_LOCAL_SCRIPT"):
		details := oci_disaster_recovery.UpdateRunLocalScriptUserDefinedStepDetails{}
		if runAsUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_as_user")); ok {
			tmp := runAsUser.(string)
			details.RunAsUser = &tmp
		}
		if runOnInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_on_instance_id")); ok {
			tmp := runOnInstanceId.(string)
			details.RunOnInstanceId = &tmp
		}
		if scriptCommand, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "script_command")); ok {
			tmp := scriptCommand.(string)
			details.ScriptCommand = &tmp
		}
		baseObject = details
	case strings.ToLower("RUN_LOCAL_SCRIPT_PRECHECK"):
		details := oci_disaster_recovery.UpdateLocalScriptPrecheckStepDetails{}
		baseObject = details
	case strings.ToLower("RUN_OBJECTSTORE_SCRIPT"):
		details := oci_disaster_recovery.UpdateRunObjectStoreScriptUserDefinedStepDetails{}
		if objectStorageScriptLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage_script_location")); ok {
			if tmpList := objectStorageScriptLocation.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage_script_location"), 0)
				tmp, err := s.mapToUpdateObjectStorageScriptLocationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert object_storage_script_location, encountered error: %v", err)
				}
				details.ObjectStorageScriptLocation = &tmp
			}
		}
		if runOnInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_on_instance_id")); ok {
			tmp := runOnInstanceId.(string)
			details.RunOnInstanceId = &tmp
		}
		baseObject = details
	case strings.ToLower("RUN_OBJECTSTORE_SCRIPT_PRECHECK"):
		details := oci_disaster_recovery.UpdateObjectStoreScriptPrecheckStepDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown step_type '%v' was specified", stepType)
	}
	return baseObject, nil
}

func DrPlanUserDefinedStepToMap(obj *oci_disaster_recovery.DrPlanUserDefinedStep) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_disaster_recovery.UpdateInvokeFunctionUserDefinedStepDetails:
		result["step_type"] = "INVOKE_FUNCTION"

		if v.FunctionId != nil {
			result["function_id"] = string(*v.FunctionId)
		}

		if v.RequestBody != nil {
			result["request_body"] = string(*v.RequestBody)
		}
	case oci_disaster_recovery.UpdateInvokeFunctionPrecheckStepDetails:
		result["step_type"] = "INVOKE_FUNCTION_PRECHECK"
	case oci_disaster_recovery.UpdateRunLocalScriptUserDefinedStepDetails:
		result["step_type"] = "RUN_LOCAL_SCRIPT"

		if v.RunAsUser != nil {
			result["run_as_user"] = string(*v.RunAsUser)
		}

		if v.RunOnInstanceId != nil {
			result["run_on_instance_id"] = string(*v.RunOnInstanceId)
		}

		if v.ScriptCommand != nil {
			result["script_command"] = string(*v.ScriptCommand)
		}
	case oci_disaster_recovery.UpdateLocalScriptPrecheckStepDetails:
		result["step_type"] = "RUN_LOCAL_SCRIPT_PRECHECK"
	case oci_disaster_recovery.UpdateRunObjectStoreScriptUserDefinedStepDetails:
		result["step_type"] = "RUN_OBJECTSTORE_SCRIPT"

		if v.ObjectStorageScriptLocation != nil {
			result["object_storage_script_location"] = []interface{}{UpdateObjectStorageScriptLocationDetailsToMap(v.ObjectStorageScriptLocation)}
		}

		if v.RunOnInstanceId != nil {
			result["run_on_instance_id"] = string(*v.RunOnInstanceId)
		}
	case oci_disaster_recovery.UpdateObjectStoreScriptPrecheckStepDetails:
		result["step_type"] = "RUN_OBJECTSTORE_SCRIPT_PRECHECK"
	default:
		log.Printf("[WARN] Received 'step_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DisasterRecoveryDrPlanResourceCrud) mapToObjectStorageScriptLocation(fieldKeyFormat string) (oci_disaster_recovery.ObjectStorageScriptLocation, error) {
	result := oci_disaster_recovery.ObjectStorageScriptLocation{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.Bucket = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
		tmp := object.(string)
		result.Object = &tmp
	}

	return result, nil
}

func ObjectStorageScriptLocationToMap(obj *oci_disaster_recovery.ObjectStorageScriptLocation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Bucket != nil {
		result["bucket"] = string(*obj.Bucket)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.Object != nil {
		result["object"] = string(*obj.Object)
	}

	return result
}

func (s *DisasterRecoveryDrPlanResourceCrud) mapToUpdateObjectStorageScriptLocationDetails(fieldKeyFormat string) (oci_disaster_recovery.UpdateObjectStorageScriptLocationDetails, error) {
	result := oci_disaster_recovery.UpdateObjectStorageScriptLocationDetails{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.Bucket = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
		tmp := object.(string)
		result.Object = &tmp
	}

	return result, nil
}

func UpdateObjectStorageScriptLocationDetailsToMap(obj *oci_disaster_recovery.UpdateObjectStorageScriptLocationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Bucket != nil {
		result["bucket"] = string(*obj.Bucket)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.Object != nil {
		result["object"] = string(*obj.Object)
	}

	return result
}
