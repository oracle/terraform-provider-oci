// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_software_update

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
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetSoftwareUpdateFsuCycleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetSoftwareUpdateFsuCycle,
		Read:     readFleetSoftwareUpdateFsuCycle,
		Update:   updateFleetSoftwareUpdateFsuCycle,
		Delete:   deleteFleetSoftwareUpdateFsuCycle,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fsu_collection_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"goal_version_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"IMAGE_ID",
								"VERSION",
							}, true),
						},

						// Optional
						"home_policy": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"new_home_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"software_image_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"PATCH",
				}, true),
			},

			// Optional
			"apply_action_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"time_to_start": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"START_TIME",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},
			"batching_strategy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_force_rolling": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_wait_for_batch_resume": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"percentage": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"FIFTY_FIFTY",
								"NONE",
								"NON_ROLLING",
								"SEQUENTIAL",
								"SERVICE_AVAILABILITY_FACTOR",
							}, true),
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
			"diagnostics_collection": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"log_collection_mode": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"is_ignore_missing_patches": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_ignore_patches": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_keep_placement": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"max_drain_timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"stage_action_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"time_to_start": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"START_TIME",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"collection_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"executing_fsu_action_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_completed_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"next_action_to_execute": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"time_to_start": {
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
			"time_finished": {
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

func createFleetSoftwareUpdateFsuCycle(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCycleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetSoftwareUpdateFsuCycle(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCycleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.ReadResource(sync)
}

func updateFleetSoftwareUpdateFsuCycle(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCycleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetSoftwareUpdateFsuCycle(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCycleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetSoftwareUpdateFsuCycleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_software_update.FleetSoftwareUpdateClient
	Res                    *oci_fleet_software_update.FsuCycle
	DisableNotFoundRetries bool
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) ID() string {
	fsuCycle := *s.Res
	return *fsuCycle.GetId()
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_software_update.CycleLifecycleStatesCreating),
		string(oci_fleet_software_update.CycleLifecycleStatesInProgress),
	}
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_software_update.CycleLifecycleStatesActive),
		string(oci_fleet_software_update.CycleLifecycleStatesNeedsAttention),
		string(oci_fleet_software_update.CycleLifecycleStatesSucceeded),
	}
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_software_update.CycleLifecycleStatesDeleting),
	}
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_software_update.CycleLifecycleStatesDeleted),
	}
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) Create() error {
	request := oci_fleet_software_update.CreateFsuCycleRequest{}
	err := s.populateTopLevelPolymorphicCreateFsuCycleRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.CreateFsuCycle(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getFsuCycleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update"), oci_fleet_software_update.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) getFsuCycleFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_software_update.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fsuCycleId, err := fsuCycleWaitForWorkRequest(workId, "cycle",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*fsuCycleId)

	return s.Get()
}

func fsuCycleWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fleet_software_update", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fleet_software_update.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func fsuCycleWaitForWorkRequest(wId *string, entityType string, action oci_fleet_software_update.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_software_update.FleetSoftwareUpdateClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_software_update")
	retryPolicy.ShouldRetryOperation = fsuCycleWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_software_update.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fleet_software_update.OperationStatusInProgress),
			string(oci_fleet_software_update.OperationStatusAccepted),
			string(oci_fleet_software_update.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_fleet_software_update.OperationStatusSucceeded),
			string(oci_fleet_software_update.OperationStatusFailed),
			string(oci_fleet_software_update.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_fleet_software_update.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_fleet_software_update.OperationStatusFailed || response.Status == oci_fleet_software_update.OperationStatusCanceled {
		return nil, getErrorFromFleetSoftwareUpdateFsuCycleWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetSoftwareUpdateFsuCycleWorkRequest(client *oci_fleet_software_update.FleetSoftwareUpdateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_software_update.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fleet_software_update.ListWorkRequestErrorsRequest{
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

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) Get() error {
	request := oci_fleet_software_update.GetFsuCycleRequest{}

	tmp := s.D.Id()
	request.FsuCycleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.GetFsuCycle(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response.FsuCycle
	return nil
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_fleet_software_update.UpdateFsuCycleRequest{}
	log.Printf("[FSULOG] Update request %v", request)
	err := s.populateTopLevelPolymorphicUpdateFsuCycleRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.UpdateFsuCycle(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFsuCycleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update"), oci_fleet_software_update.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) Delete() error {
	request := oci_fleet_software_update.DeleteFsuCycleRequest{}

	tmp := s.D.Id()
	request.FsuCycleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")
	response, err := s.Client.DeleteFsuCycle(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := fsuCycleWaitForWorkRequest(workId, "cycle",
		oci_fleet_software_update.ActionTypeRelated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_fleet_software_update.PatchFsuCycle:
		s.D.Set("type", "PATCH")

		s.D.Set("is_ignore_missing_patches", v.IsIgnoreMissingPatches)

		if v.IsIgnorePatches != nil {
			s.D.Set("is_ignore_patches", *v.IsIgnorePatches)
		}

		if v.IsKeepPlacement != nil {
			s.D.Set("is_keep_placement", *v.IsKeepPlacement)
		}

		if v.MaxDrainTimeoutInSeconds != nil {
			s.D.Set("max_drain_timeout_in_seconds", *v.MaxDrainTimeoutInSeconds)
		}

		if v.ApplyActionSchedule != nil {
			applyActionScheduleArray := []interface{}{}
			if applyActionScheduleMap := ScheduleDetailsToMap(&v.ApplyActionSchedule); applyActionScheduleMap != nil {
				applyActionScheduleArray = append(applyActionScheduleArray, applyActionScheduleMap)
			}
			s.D.Set("apply_action_schedule", applyActionScheduleArray)
		} else {
			s.D.Set("apply_action_schedule", nil)
		}

		if v.StageActionSchedule != nil {
			stageActionScheduleArray := []interface{}{}
			if stageActionScheduleMap := ScheduleDetailsToMap(&v.StageActionSchedule); stageActionScheduleMap != nil {
				stageActionScheduleArray = append(stageActionScheduleArray, stageActionScheduleMap)
			}
			s.D.Set("stage_action_schedule", stageActionScheduleArray)

		} else {
			s.D.Set("stage_action_schedule", nil)
		}

		if v.BatchingStrategy != nil {
			batchingStrategyArray := []interface{}{}
			if batchingStrategyMap := BatchingStrategyDetailsToMap(&v.BatchingStrategy); batchingStrategyMap != nil {
				batchingStrategyArray = append(batchingStrategyArray, batchingStrategyMap)
			}
			s.D.Set("batching_strategy", batchingStrategyArray)
		} else {
			s.D.Set("batching_strategy", nil)
		}
		log.Printf("[FSULOG] collection_type Set: %v", v.CollectionType)
		s.D.Set("collection_type", v.CollectionType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DiagnosticsCollection != nil {
			s.D.Set("diagnostics_collection", []interface{}{DiagnosticsCollectionDetailsToMap(v.DiagnosticsCollection)})
		} else {
			s.D.Set("diagnostics_collection", nil)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ExecutingFsuActionId != nil {
			s.D.Set("executing_fsu_action_id", *v.ExecutingFsuActionId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.FsuCollectionId != nil {
			s.D.Set("fsu_collection_id", *v.FsuCollectionId)
		}
		if v.GoalVersionDetails != nil {
			goalVersionDetailsArray := []interface{}{}
			if goalVersionDetailsMap := FsuGoalVersionDetailsToMap(&v.GoalVersionDetails); goalVersionDetailsMap != nil {
				goalVersionDetailsArray = append(goalVersionDetailsArray, goalVersionDetailsMap)
			}
			s.D.Set("goal_version_details", goalVersionDetailsArray)
		} else {
			s.D.Set("goal_version_details", nil)
		}

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		s.D.Set("last_completed_action", v.LastCompletedAction)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		nextActionToExecute := []interface{}{}
		for _, item := range v.NextActionToExecute {
			nextActionToExecute = append(nextActionToExecute, NextActionToExecuteDetailsToMap(item))
		}
		s.D.Set("next_action_to_execute", nextActionToExecute)

		if v.StageActionSchedule != nil {
			stageActionScheduleArray := []interface{}{}
			if stageActionScheduleMap := ScheduleDetailsToMap(&v.StageActionSchedule); stageActionScheduleMap != nil {
				stageActionScheduleArray = append(stageActionScheduleArray, stageActionScheduleMap)
			}
			s.D.Set("stage_action_schedule", stageActionScheduleArray)
		} else {
			s.D.Set("stage_action_schedule", nil)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeFinished != nil {
			s.D.Set("time_finished", v.TimeFinished.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) mapToCreateBatchingStrategyDetails(fieldKeyFormat string) (oci_fleet_software_update.CreateBatchingStrategyDetails, error) {
	var baseObject oci_fleet_software_update.CreateBatchingStrategyDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("FIFTY_FIFTY"):
		details := oci_fleet_software_update.UpdateFiftyFiftyBatchingStrategyDetails{}
		if isForceRolling, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_force_rolling")); ok {
			tmp := isForceRolling.(bool)
			details.IsForceRolling = &tmp
		}
		if isWaitForBatchResume, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_wait_for_batch_resume")); ok {
			tmp := isWaitForBatchResume.(bool)
			details.IsWaitForBatchResume = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_fleet_software_update.NoneBatchingStrategyDetails{}
		baseObject = details
	case strings.ToLower("NON_ROLLING"):
		details := oci_fleet_software_update.UpdateNonRollingBatchingStrategyDetails{}
		baseObject = details
	case strings.ToLower("SEQUENTIAL"):
		details := oci_fleet_software_update.UpdateSequentialBatchingStrategyDetails{}
		if isForceRolling, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_force_rolling")); ok {
			tmp := isForceRolling.(bool)
			details.IsForceRolling = &tmp
		}
		baseObject = details
	case strings.ToLower("SERVICE_AVAILABILITY_FACTOR"):
		details := oci_fleet_software_update.UpdateServiceAvailabilityFactorBatchingStrategyDetails{}
		if isForceRolling, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_force_rolling")); ok {
			tmp := isForceRolling.(bool)
			details.IsForceRolling = &tmp
		}
		if percentage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "percentage")); ok {
			tmp := percentage.(int)
			details.Percentage = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) mapToUpdateBatchingStrategyDetails(fieldKeyFormat string) (oci_fleet_software_update.UpdateBatchingStrategyDetails, error) {
	var baseObject oci_fleet_software_update.UpdateBatchingStrategyDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("FIFTY_FIFTY"):
		details := oci_fleet_software_update.UpdateFiftyFiftyBatchingStrategyDetails{}
		if isForceRolling, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_force_rolling")); ok {
			tmp := isForceRolling.(bool)
			details.IsForceRolling = &tmp
		}
		if isWaitForBatchResume, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_wait_for_batch_resume")); ok {
			tmp := isWaitForBatchResume.(bool)
			details.IsWaitForBatchResume = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_fleet_software_update.NoneBatchingStrategyDetails{}
		baseObject = details
	case strings.ToLower("NON_ROLLING"):
		details := oci_fleet_software_update.UpdateNonRollingBatchingStrategyDetails{}
		baseObject = details
	case strings.ToLower("SEQUENTIAL"):
		details := oci_fleet_software_update.UpdateSequentialBatchingStrategyDetails{}
		if isForceRolling, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_force_rolling")); ok {
			tmp := isForceRolling.(bool)
			details.IsForceRolling = &tmp
		}
		baseObject = details
	case strings.ToLower("SERVICE_AVAILABILITY_FACTOR"):
		details := oci_fleet_software_update.UpdateServiceAvailabilityFactorBatchingStrategyDetails{}
		if isForceRolling, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_force_rolling")); ok {
			tmp := isForceRolling.(bool)
			details.IsForceRolling = &tmp
		}
		if percentage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "percentage")); ok {
			tmp := percentage.(int)
			details.Percentage = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func BatchingStrategyDetailsToMap(obj *oci_fleet_software_update.BatchingStrategyDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_software_update.FiftyFiftyBatchingStrategyDetails:
		result["type"] = "FIFTY_FIFTY"

		if v.IsForceRolling != nil {
			result["is_force_rolling"] = bool(*v.IsForceRolling)
		}

		if v.IsWaitForBatchResume != nil {
			result["is_wait_for_batch_resume"] = bool(*v.IsWaitForBatchResume)
		}
	case oci_fleet_software_update.NoneBatchingStrategyDetails:
		result["type"] = "NONE"
	case oci_fleet_software_update.NonRollingBatchingStrategyDetails:
		result["type"] = "NON_ROLLING"
	case oci_fleet_software_update.SequentialBatchingStrategyDetails:
		result["type"] = "SEQUENTIAL"

		if v.IsForceRolling != nil {
			result["is_force_rolling"] = bool(*v.IsForceRolling)
		}
	case oci_fleet_software_update.ServiceAvailabilityFactorBatchingStrategyDetails:
		result["type"] = "SERVICE_AVAILABILITY_FACTOR"

		if v.IsForceRolling != nil {
			result["is_force_rolling"] = bool(*v.IsForceRolling)
		}

		if v.Percentage != nil {
			result["percentage"] = int(*v.Percentage)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %T", v)
		return nil
	}

	return result
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) mapToCreateScheduleDetails(fieldKeyFormat string) (oci_fleet_software_update.CreateScheduleDetails, error) {
	var baseObject oci_fleet_software_update.CreateScheduleDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("START_TIME"):
		details := oci_fleet_software_update.CreateStartTimeScheduleDetails{}
		if timeToStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_to_start")); ok {
			tmp, err := time.Parse(time.RFC3339, timeToStart.(string))
			if err != nil {
				return details, err
			}
			details.TimeToStart = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ScheduleDetailsToMap(obj *oci_fleet_software_update.ScheduleDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_software_update.StartTimeScheduleDetails:
		result["type"] = "START_TIME"

		if v.TimeToStart != nil {
			result["time_to_start"] = v.TimeToStart.Format(time.RFC3339Nano)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %T", *obj)
		return nil
	}

	return result
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) mapToDiagnosticsCollectionDetails(fieldKeyFormat string) (oci_fleet_software_update.DiagnosticsCollectionDetails, error) {
	result := oci_fleet_software_update.DiagnosticsCollectionDetails{}

	if logCollectionMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_collection_mode")); ok {
		result.LogCollectionMode = oci_fleet_software_update.DataCollectionModesEnum(logCollectionMode.(string))
	}

	return result, nil
}

func DiagnosticsCollectionDetailsToMap(obj *oci_fleet_software_update.DiagnosticsCollectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["log_collection_mode"] = string(obj.LogCollectionMode)

	return result
}

func FsuCycleSummaryToMap(obj oci_fleet_software_update.FsuCycleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["collection_type"] = string(obj.CollectionType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DiagnosticsCollection != nil {
		result["diagnostics_collection"] = []interface{}{DiagnosticsCollectionDetailsToMap(obj.DiagnosticsCollection)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExecutingFsuActionId != nil {
		result["executing_fsu_action_id"] = string(*obj.ExecutingFsuActionId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.FsuCollectionId != nil {
		result["fsu_collection_id"] = string(*obj.FsuCollectionId)
	}

	if obj.GoalVersionDetails != nil {
		goalVersionDetailsArray := []interface{}{}
		if goalVersionDetailsMap := FsuGoalVersionDetailsToMap(&obj.GoalVersionDetails); goalVersionDetailsMap != nil {
			goalVersionDetailsArray = append(goalVersionDetailsArray, goalVersionDetailsMap)
		}
		result["goal_version_details"] = goalVersionDetailsArray
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["last_completed_action"] = string(obj.LastCompletedAction)

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	nextActionToExecute := []interface{}{}
	for _, item := range obj.NextActionToExecute {
		nextActionToExecute = append(nextActionToExecute, NextActionToExecuteDetailsToMap(item))
	}
	result["next_action_to_execute"] = nextActionToExecute

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) mapToFsuGoalVersionDetails(fieldKeyFormat string) (oci_fleet_software_update.FsuGoalVersionDetails, error) {
	var baseObject oci_fleet_software_update.FsuGoalVersionDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("IMAGE_ID"):
		details := oci_fleet_software_update.ImageIdFsuTargetDetails{}
		if softwareImageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "software_image_id")); ok {
			tmp := softwareImageId.(string)
			details.SoftwareImageId = &tmp
		}
		if homePolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "home_policy")); ok && homePolicy != "" {
			details.HomePolicy = oci_fleet_software_update.FsuGoalVersionDetailsHomePolicyEnum(homePolicy.(string))
		}
		if newHomePrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "new_home_prefix")); ok && newHomePrefix != "" {
			tmp := newHomePrefix.(string)
			details.NewHomePrefix = &tmp
		}
		baseObject = details
	case strings.ToLower("VERSION"):
		details := oci_fleet_software_update.VersionFsuTargetDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if homePolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "home_policy")); ok {
			details.HomePolicy = oci_fleet_software_update.FsuGoalVersionDetailsHomePolicyEnum(homePolicy.(string))
		}
		if newHomePrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "new_home_prefix")); ok {
			tmp := newHomePrefix.(string)
			details.NewHomePrefix = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func FsuGoalVersionDetailsToMap(obj *oci_fleet_software_update.FsuGoalVersionDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_software_update.ImageIdFsuTargetDetails:
		result["type"] = "IMAGE_ID"

		if v.SoftwareImageId != nil {
			result["software_image_id"] = string(*v.SoftwareImageId)
		}
		if v.NewHomePrefix != nil {
			result["new_home_prefix"] = string(*v.NewHomePrefix)
		}
		if homePolicy := oci_fleet_software_update.FsuGoalVersionDetailsHomePolicyEnum(v.HomePolicy); homePolicy != "" {
			result["home_policy"] = homePolicy
		}
	case oci_fleet_software_update.VersionFsuTargetDetails:
		result["type"] = "VERSION"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}
		if v.NewHomePrefix != nil {
			result["new_home_prefix"] = string(*v.NewHomePrefix)
		}
		if homePolicy := oci_fleet_software_update.FsuGoalVersionDetailsHomePolicyEnum(v.HomePolicy); homePolicy != "" {
			result["home_policy"] = homePolicy
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %T", v)
		return nil
	}
	return result
}

func NextActionToExecuteDetailsToMap(obj oci_fleet_software_update.NextActionToExecuteDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TimeToStart != nil {
		result["time_to_start"] = obj.TimeToStart.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) populateTopLevelPolymorphicCreateFsuCycleRequest(request *oci_fleet_software_update.CreateFsuCycleRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("PATCH"):
		details := oci_fleet_software_update.CreatePatchFsuCycle{}
		if isIgnoreMissingPatches, ok := s.D.GetOkExists("is_ignore_missing_patches"); ok {
			interfaces := isIgnoreMissingPatches.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("is_ignore_missing_patches") {
				details.IsIgnoreMissingPatches = tmp
			}
		}
		if isIgnorePatches, ok := s.D.GetOkExists("is_ignore_patches"); ok {
			tmp := isIgnorePatches.(bool)
			details.IsIgnorePatches = &tmp
		}
		if isKeepPlacement, ok := s.D.GetOkExists("is_keep_placement"); ok {
			tmp := isKeepPlacement.(bool)
			details.IsKeepPlacement = &tmp
		}
		if maxDrainTimeoutInSeconds, ok := s.D.GetOkExists("max_drain_timeout_in_seconds"); ok {
			tmp := maxDrainTimeoutInSeconds.(int)
			details.MaxDrainTimeoutInSeconds = &tmp
		}
		if applyActionSchedule, ok := s.D.GetOkExists("apply_action_schedule"); ok {
			if tmpList := applyActionSchedule.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "apply_action_schedule", 0)
				tmp, err := s.mapToCreateScheduleDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ApplyActionSchedule = tmp
			}
		}
		if batchingStrategy, ok := s.D.GetOkExists("batching_strategy"); ok {
			if tmpList := batchingStrategy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "batching_strategy", 0)
				tmp, err := s.mapToCreateBatchingStrategyDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BatchingStrategy = tmp
			}
		}
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
		if diagnosticsCollection, ok := s.D.GetOkExists("diagnostics_collection"); ok {
			if tmpList := diagnosticsCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "diagnostics_collection", 0)
				tmp, err := s.mapToDiagnosticsCollectionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DiagnosticsCollection = &tmp
			}
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if fsuCollectionId, ok := s.D.GetOkExists("fsu_collection_id"); ok {
			tmp := fsuCollectionId.(string)
			details.FsuCollectionId = &tmp
		}
		if goalVersionDetails, ok := s.D.GetOkExists("goal_version_details"); ok {
			if tmpList := goalVersionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "goal_version_details", 0)
				tmp, err := s.mapToFsuGoalVersionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.GoalVersionDetails = tmp
			}
		}
		if stageActionSchedule, ok := s.D.GetOkExists("stage_action_schedule"); ok {
			if tmpList := stageActionSchedule.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "stage_action_schedule", 0)
				tmp, err := s.mapToCreateScheduleDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.StageActionSchedule = tmp
			}
		}
		request.CreateFsuCycleDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) populateTopLevelPolymorphicUpdateFsuCycleRequest(request *oci_fleet_software_update.UpdateFsuCycleRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("PATCH"):
		details := oci_fleet_software_update.UpdatePatchFsuCycle{}
		if isIgnoreMissingPatches, ok := s.D.GetOkExists("is_ignore_missing_patches"); ok {
			interfaces := isIgnoreMissingPatches.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("is_ignore_missing_patches") {
				details.IsIgnoreMissingPatches = tmp
			}
		}
		if isIgnorePatches, ok := s.D.GetOkExists("is_ignore_patches"); ok {
			tmp := isIgnorePatches.(bool)
			details.IsIgnorePatches = &tmp
		}
		if isKeepPlacement, ok := s.D.GetOkExists("is_keep_placement"); ok {
			tmp := isKeepPlacement.(bool)
			details.IsKeepPlacement = &tmp
		}
		if maxDrainTimeoutInSeconds, ok := s.D.GetOkExists("max_drain_timeout_in_seconds"); ok {
			tmp := maxDrainTimeoutInSeconds.(int)
			details.MaxDrainTimeoutInSeconds = &tmp
		}
		if batchingStrategy, ok := s.D.GetOkExists("batching_strategy"); ok {
			if tmpList := batchingStrategy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "batching_strategy", 0)
				tmp, err := s.mapToUpdateBatchingStrategyDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BatchingStrategy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if diagnosticsCollection, ok := s.D.GetOkExists("diagnostics_collection"); ok {
			if tmpList := diagnosticsCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "diagnostics_collection", 0)
				tmp, err := s.mapToDiagnosticsCollectionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DiagnosticsCollection = &tmp
			}
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.FsuCycleId = &tmp
		if goalVersionDetails, ok := s.D.GetOkExists("goal_version_details"); ok {
			log.Printf("[FSULOG] goalVersionDetails-Update %v", goalVersionDetails)
			if tmpList := goalVersionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "goal_version_details", 0)
				tmp, err := s.mapToFsuGoalVersionDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.GoalVersionDetails = tmp
			}
		}
		request.UpdateFsuCycleDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *FleetSoftwareUpdateFsuCycleResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_fleet_software_update.ChangeFsuCycleCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.FsuCycleId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.ChangeFsuCycleCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFsuCycleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update"), oci_fleet_software_update.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
