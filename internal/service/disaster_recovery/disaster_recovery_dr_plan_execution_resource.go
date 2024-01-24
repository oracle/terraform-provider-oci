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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DisasterRecoveryDrPlanExecutionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDisasterRecoveryDrPlanExecution,
		Read:     readDisasterRecoveryDrPlanExecution,
		Update:   updateDisasterRecoveryDrPlanExecution,
		Delete:   deleteDisasterRecoveryDrPlanExecution,
		Schema: map[string]*schema.Schema{
			// Required
			"execution_options": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"plan_execution_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"FAILOVER",
								"FAILOVER_PRECHECK",
								"START_DRILL",
								"START_DRILL_PRECHECK",
								"STOP_DRILL",
								"STOP_DRILL_PRECHECK",
								"SWITCHOVER",
								"SWITCHOVER_PRECHECK",
							}, true),
						},

						// Optional
						"are_prechecks_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"are_warnings_ignored": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"plan_id": {
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
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dr_protection_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"execution_duration_in_sec": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"group_executions": {
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
						"execution_duration_in_sec": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"step_executions": {
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
									"execution_duration_in_sec": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"group_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"log_location": {
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
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"step_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"life_cycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_location": {
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
			"peer_dr_protection_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plan_execution_type": {
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
			"time_ended": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
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

func createDisasterRecoveryDrPlanExecution(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlanExecutionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.CreateResource(d, sync)
}

func readDisasterRecoveryDrPlanExecution(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlanExecutionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.ReadResource(sync)
}

func updateDisasterRecoveryDrPlanExecution(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlanExecutionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDisasterRecoveryDrPlanExecution(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlanExecutionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DisasterRecoveryDrPlanExecutionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_disaster_recovery.DisasterRecoveryClient
	Res                    *oci_disaster_recovery.DrPlanExecution
	DisableNotFoundRetries bool
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_disaster_recovery.DrPlanExecutionLifecycleStateInProgress),
	}
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_disaster_recovery.DrPlanExecutionLifecycleStateSucceeded),
	}
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_disaster_recovery.DrPlanExecutionLifecycleStateDeleting),
	}
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_disaster_recovery.DrPlanExecutionLifecycleStateDeleted),
	}
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) Create() error {
	request := oci_disaster_recovery.CreateDrPlanExecutionRequest{}

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

	if executionOptions, ok := s.D.GetOkExists("execution_options"); ok {
		if tmpList := executionOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "execution_options", 0)
			tmp, err := s.mapToDrPlanExecutionOptionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ExecutionOptions = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if planId, ok := s.D.GetOkExists("plan_id"); ok {
		tmp := planId.(string)
		request.PlanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.CreateDrPlanExecution(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDrPlanExecutionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) getDrPlanExecutionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_disaster_recovery.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	drPlanExecutionId, err := drPlanExecutionWaitForWorkRequest(workId, "drPlanExecution",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, drPlanExecutionId)
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
	s.D.SetId(*drPlanExecutionId)

	return s.Get()
}

func drPlanExecutionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func drPlanExecutionWaitForWorkRequest(wId *string, entityType string, action oci_disaster_recovery.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_disaster_recovery.DisasterRecoveryClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "disaster_recovery")
	retryPolicy.ShouldRetryOperation = drPlanExecutionWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDisasterRecoveryDrPlanExecutionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDisasterRecoveryDrPlanExecutionWorkRequest(client *oci_disaster_recovery.DisasterRecoveryClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_disaster_recovery.ActionTypeEnum) error {
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

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) Get() error {
	request := oci_disaster_recovery.GetDrPlanExecutionRequest{}

	tmp := s.D.Id()
	request.DrPlanExecutionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.GetDrPlanExecution(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrPlanExecution
	return nil
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) Update() error {
	request := oci_disaster_recovery.UpdateDrPlanExecutionRequest{}

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
	request.DrPlanExecutionId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.UpdateDrPlanExecution(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDrPlanExecutionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery"), oci_disaster_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) Delete() error {
	request := oci_disaster_recovery.DeleteDrPlanExecutionRequest{}

	tmp := s.D.Id()
	request.DrPlanExecutionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "disaster_recovery")

	response, err := s.Client.DeleteDrPlanExecution(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := drPlanExecutionWaitForWorkRequest(workId, "drPlanExecution",
		oci_disaster_recovery.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) SetData() error {
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

	if s.Res.ExecutionDurationInSec != nil {
		s.D.Set("execution_duration_in_sec", *s.Res.ExecutionDurationInSec)
	}

	if s.Res.ExecutionOptions != nil {
		executionOptionsArray := []interface{}{}
		if executionOptionsMap := DrPlanExecutionOptionsToMap(&s.Res.ExecutionOptions); executionOptionsMap != nil {
			executionOptionsArray = append(executionOptionsArray, executionOptionsMap)
		}
		s.D.Set("execution_options", executionOptionsArray)
	} else {
		s.D.Set("execution_options", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	groupExecutions := []interface{}{}
	for _, item := range s.Res.GroupExecutions {
		groupExecutions = append(groupExecutions, DrPlanGroupExecutionToMap(item))
	}
	s.D.Set("group_executions", groupExecutions)

	if s.Res.LifeCycleDetails != nil {
		s.D.Set("life_cycle_details", *s.Res.LifeCycleDetails)
	}

	if s.Res.LogLocation != nil {
		s.D.Set("log_location", []interface{}{ObjectStorageLogLocationToMap(s.Res.LogLocation)})
	} else {
		s.D.Set("log_location", nil)
	}

	if s.Res.PeerDrProtectionGroupId != nil {
		s.D.Set("peer_dr_protection_group_id", *s.Res.PeerDrProtectionGroupId)
	}

	if s.Res.PeerRegion != nil {
		s.D.Set("peer_region", *s.Res.PeerRegion)
	}

	s.D.Set("plan_execution_type", s.Res.PlanExecutionType)

	if s.Res.PlanId != nil {
		s.D.Set("plan_id", *s.Res.PlanId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DisasterRecoveryDrPlanExecutionResourceCrud) mapToDrPlanExecutionOptionDetails(fieldKeyFormat string) (oci_disaster_recovery.DrPlanExecutionOptionDetails, error) {
	var baseObject oci_disaster_recovery.DrPlanExecutionOptionDetails
	//discriminator
	planExecutionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "plan_execution_type"))
	var planExecutionType string
	if ok {
		planExecutionType = planExecutionTypeRaw.(string)
	} else {
		planExecutionType = "" // default value
	}
	switch strings.ToLower(planExecutionType) {
	case strings.ToLower("FAILOVER"):
		details := oci_disaster_recovery.FailoverExecutionOptionDetails{}
		if arePrechecksEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_prechecks_enabled")); ok {
			tmp := arePrechecksEnabled.(bool)
			details.ArePrechecksEnabled = &tmp
		}
		if areWarningsIgnored, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_warnings_ignored")); ok {
			tmp := areWarningsIgnored.(bool)
			details.AreWarningsIgnored = &tmp
		}
		baseObject = details
	case strings.ToLower("FAILOVER_PRECHECK"):
		details := oci_disaster_recovery.FailoverPrecheckExecutionOptionDetails{}
		if areWarningsIgnored, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_warnings_ignored")); ok {
			tmp := areWarningsIgnored.(bool)
			details.AreWarningsIgnored = &tmp
		}
		baseObject = details
	case strings.ToLower("START_DRILL"):
		details := oci_disaster_recovery.StartDrillExecutionOptionDetails{}
		if arePrechecksEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_prechecks_enabled")); ok {
			tmp := arePrechecksEnabled.(bool)
			details.ArePrechecksEnabled = &tmp
		}
		if areWarningsIgnored, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_warnings_ignored")); ok {
			tmp := areWarningsIgnored.(bool)
			details.AreWarningsIgnored = &tmp
		}
		baseObject = details
	case strings.ToLower("START_DRILL_PRECHECK"):
		details := oci_disaster_recovery.StartDrillPrecheckExecutionOptionDetails{}
		if areWarningsIgnored, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_warnings_ignored")); ok {
			tmp := areWarningsIgnored.(bool)
			details.AreWarningsIgnored = &tmp
		}
		baseObject = details
	case strings.ToLower("STOP_DRILL"):
		details := oci_disaster_recovery.StopDrillExecutionOptionDetails{}
		if arePrechecksEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_prechecks_enabled")); ok {
			tmp := arePrechecksEnabled.(bool)
			details.ArePrechecksEnabled = &tmp
		}
		if areWarningsIgnored, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_warnings_ignored")); ok {
			tmp := areWarningsIgnored.(bool)
			details.AreWarningsIgnored = &tmp
		}
		baseObject = details
	case strings.ToLower("STOP_DRILL_PRECHECK"):
		details := oci_disaster_recovery.StopDrillPrecheckExecutionOptionDetails{}
		if areWarningsIgnored, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_warnings_ignored")); ok {
			tmp := areWarningsIgnored.(bool)
			details.AreWarningsIgnored = &tmp
		}
		baseObject = details
	case strings.ToLower("SWITCHOVER"):
		details := oci_disaster_recovery.SwitchoverExecutionOptionDetails{}
		if arePrechecksEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_prechecks_enabled")); ok {
			tmp := arePrechecksEnabled.(bool)
			details.ArePrechecksEnabled = &tmp
		}
		if areWarningsIgnored, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_warnings_ignored")); ok {
			tmp := areWarningsIgnored.(bool)
			details.AreWarningsIgnored = &tmp
		}
		baseObject = details
	case strings.ToLower("SWITCHOVER_PRECHECK"):
		details := oci_disaster_recovery.SwitchoverPrecheckExecutionOptionDetails{}
		if areWarningsIgnored, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_warnings_ignored")); ok {
			tmp := areWarningsIgnored.(bool)
			details.AreWarningsIgnored = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown plan_execution_type '%v' was specified", planExecutionType)
	}
	return baseObject, nil
}

func DrPlanExecutionOptionsToMap(obj *oci_disaster_recovery.DrPlanExecutionOptions) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_disaster_recovery.FailoverExecutionOptions:
		result["plan_execution_type"] = "FAILOVER"

		if v.ArePrechecksEnabled != nil {
			result["are_prechecks_enabled"] = bool(*v.ArePrechecksEnabled)
		}

		if v.AreWarningsIgnored != nil {
			result["are_warnings_ignored"] = bool(*v.AreWarningsIgnored)
		}
	case oci_disaster_recovery.FailoverPrecheckExecutionOptions:
		result["plan_execution_type"] = "FAILOVER_PRECHECK"

		if v.AreWarningsIgnored != nil {
			result["are_warnings_ignored"] = bool(*v.AreWarningsIgnored)
		}
	case oci_disaster_recovery.StartDrillExecutionOptionDetails:
		result["plan_execution_type"] = "START_DRILL"

		if v.ArePrechecksEnabled != nil {
			result["are_prechecks_enabled"] = bool(*v.ArePrechecksEnabled)
		}

		if v.AreWarningsIgnored != nil {
			result["are_warnings_ignored"] = bool(*v.AreWarningsIgnored)
		}
	case oci_disaster_recovery.StartDrillPrecheckExecutionOptionDetails:
		result["plan_execution_type"] = "START_DRILL_PRECHECK"

		if v.AreWarningsIgnored != nil {
			result["are_warnings_ignored"] = bool(*v.AreWarningsIgnored)
		}
	case oci_disaster_recovery.StopDrillExecutionOptionDetails:
		result["plan_execution_type"] = "STOP_DRILL"

		if v.ArePrechecksEnabled != nil {
			result["are_prechecks_enabled"] = bool(*v.ArePrechecksEnabled)
		}

		if v.AreWarningsIgnored != nil {
			result["are_warnings_ignored"] = bool(*v.AreWarningsIgnored)
		}
	case oci_disaster_recovery.StopDrillPrecheckExecutionOptionDetails:
		result["plan_execution_type"] = "STOP_DRILL_PRECHECK"

		if v.AreWarningsIgnored != nil {
			result["are_warnings_ignored"] = bool(*v.AreWarningsIgnored)
		}
	case oci_disaster_recovery.SwitchoverExecutionOptions:
		result["plan_execution_type"] = "SWITCHOVER"

		if v.ArePrechecksEnabled != nil {
			result["are_prechecks_enabled"] = bool(*v.ArePrechecksEnabled)
		}

		if v.AreWarningsIgnored != nil {
			result["are_warnings_ignored"] = bool(*v.AreWarningsIgnored)
		}
	case oci_disaster_recovery.SwitchoverPrecheckExecutionOptions:
		result["plan_execution_type"] = "SWITCHOVER_PRECHECK"

		if v.AreWarningsIgnored != nil {
			result["are_warnings_ignored"] = bool(*v.AreWarningsIgnored)
		}
	default:
		log.Printf("[WARN] Received 'plan_execution_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DrPlanExecutionSummaryToMap(obj oci_disaster_recovery.DrPlanExecutionSummary) map[string]interface{} {
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

	if obj.ExecutionDurationInSec != nil {
		result["execution_duration_in_sec"] = int(*obj.ExecutionDurationInSec)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifeCycleDetails != nil {
		result["life_cycle_details"] = string(*obj.LifeCycleDetails)
	}

	if obj.LogLocation != nil {
		result["log_location"] = []interface{}{ObjectStorageLogLocationToMap(obj.LogLocation)}
	}

	if obj.PeerDrProtectionGroupId != nil {
		result["peer_dr_protection_group_id"] = string(*obj.PeerDrProtectionGroupId)
	}

	if obj.PeerRegion != nil {
		result["peer_region"] = string(*obj.PeerRegion)
	}

	result["plan_execution_type"] = string(obj.PlanExecutionType)

	if obj.PlanId != nil {
		result["plan_id"] = string(*obj.PlanId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func DrPlanGroupExecutionToMap(obj oci_disaster_recovery.DrPlanGroupExecution) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExecutionDurationInSec != nil {
		result["execution_duration_in_sec"] = int(*obj.ExecutionDurationInSec)
	}

	if obj.GroupId != nil {
		result["group_id"] = string(*obj.GroupId)
	}

	result["status"] = string(obj.Status)

	if obj.StatusDetails != nil {
		result["status_details"] = string(*obj.StatusDetails)
	}

	stepExecutions := []interface{}{}
	for _, item := range obj.StepExecutions {
		stepExecutions = append(stepExecutions, DrPlanStepExecutionToMap(item))
	}
	result["step_executions"] = stepExecutions

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func DrPlanStepExecutionToMap(obj oci_disaster_recovery.DrPlanStepExecution) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExecutionDurationInSec != nil {
		result["execution_duration_in_sec"] = int(*obj.ExecutionDurationInSec)
	}

	if obj.GroupId != nil {
		result["group_id"] = string(*obj.GroupId)
	}

	if obj.LogLocation != nil {
		result["log_location"] = []interface{}{ObjectStorageLogLocationToMap(obj.LogLocation)}
	}

	result["status"] = string(obj.Status)

	if obj.StatusDetails != nil {
		result["status_details"] = string(*obj.StatusDetails)
	}

	if obj.StepId != nil {
		result["step_id"] = string(*obj.StepId)
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	result["type"] = string(obj.Type)

	return result
}
