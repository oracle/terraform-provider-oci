// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_scheduler

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
	oci_resource_scheduler "github.com/oracle/oci-go-sdk/v65/resourcescheduler"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourceSchedulerScheduleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createResourceSchedulerSchedule,
		Read:     readResourceSchedulerSchedule,
		Update:   updateResourceSchedulerSchedule,
		Delete:   deleteResourceSchedulerSchedule,
		Schema: map[string]*schema.Schema{
			// Required
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"recurrence_details": {
				Type:     schema.TypeString,
				Required: true,
			},
			"recurrence_type": {
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
			"resource_filters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"attribute": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COMPARTMENT_ID",
								"DEFINED_TAGS",
								"LIFECYCLE_STATE",
								"RESOURCE_TYPE",
								"TIME_CREATED",
							}, true),
						},

						// Optional
						"condition": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"should_include_child_compartments": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tag_key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"resources": {
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
						"metadata": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},

						// Computed
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_resource_scheduler.ScheduleLifecycleStateInactive),
					string(oci_resource_scheduler.ScheduleLifecycleStateActive),
				}, true),
			},
			"time_ends": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"time_starts": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Computed
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_run": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_next_run": {
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

func createResourceSchedulerSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceSchedulerScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduleClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_resource_scheduler.ScheduleLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_resource_scheduler.ScheduleLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopSchedule(); err != nil {
			return err
		}
		sync.D.Set("state", oci_resource_scheduler.ScheduleLifecycleStateInactive)
	}
	return nil

}

func readResourceSchedulerSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceSchedulerScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduleClient()

	return tfresource.ReadResource(sync)
}

func updateResourceSchedulerSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceSchedulerScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduleClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_resource_scheduler.ScheduleLifecycleStateActive == oci_resource_scheduler.ScheduleLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_resource_scheduler.ScheduleLifecycleStateInactive == oci_resource_scheduler.ScheduleLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartSchedule(); err != nil {
			return err
		}
		sync.D.Set("state", oci_resource_scheduler.ScheduleLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopSchedule(); err != nil {
			return err
		}
		sync.D.Set("state", oci_resource_scheduler.ScheduleLifecycleStateInactive)
	}

	return nil
}

func deleteResourceSchedulerSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceSchedulerScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduleClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ResourceSchedulerScheduleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_resource_scheduler.ScheduleClient
	Res                    *oci_resource_scheduler.Schedule
	DisableNotFoundRetries bool
}

func (s *ResourceSchedulerScheduleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ResourceSchedulerScheduleResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_resource_scheduler.ScheduleLifecycleStateCreating),
	}
}

func (s *ResourceSchedulerScheduleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_resource_scheduler.ScheduleLifecycleStateActive),
	}
}

func (s *ResourceSchedulerScheduleResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_resource_scheduler.ScheduleLifecycleStateDeleting),
	}
}

func (s *ResourceSchedulerScheduleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_resource_scheduler.ScheduleLifecycleStateDeleted),
	}
}

func (s *ResourceSchedulerScheduleResourceCrud) Create() error {
	request := oci_resource_scheduler.CreateScheduleRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		request.Action = oci_resource_scheduler.CreateScheduleDetailsActionEnum(action.(string))
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

	if recurrenceDetails, ok := s.D.GetOkExists("recurrence_details"); ok {
		tmp := recurrenceDetails.(string)
		request.RecurrenceDetails = &tmp
	}

	if recurrenceType, ok := s.D.GetOkExists("recurrence_type"); ok {
		request.RecurrenceType = oci_resource_scheduler.CreateScheduleDetailsRecurrenceTypeEnum(recurrenceType.(string))
	}

	if resourceFilters, ok := s.D.GetOkExists("resource_filters"); ok {
		interfaces := resourceFilters.([]interface{})
		tmp := make([]oci_resource_scheduler.ResourceFilter, len(interfaces))
		for i := range interfaces {
			//interfaces[i] map[attribute:RESOURCE_TYPE condition: should_include_child_compartments:%!s(bool=false) value:[map[namespace: tag_key: value:instance]]]
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_filters", stateDataIndex)
			converted, err := s.mapToResourceFilter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
			//{ Value=[{ Namespace=ResourceSchedulerCanary TagKey=ScheduleTagFilterTestKey Value=foo }] }
		}
		if len(tmp) != 0 || s.D.HasChange("resource_filters") {
			request.ResourceFilters = tmp
		}
	}

	if resources, ok := s.D.GetOkExists("resources"); ok {
		interfaces := resources.([]interface{})
		tmp := make([]oci_resource_scheduler.Resource, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resources", stateDataIndex)
			converted, err := s.mapToResource(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("resources") {
			request.Resources = tmp
		}
	}

	if timeEnds, ok := s.D.GetOkExists("time_ends"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnds.(string))
		if err != nil {
			return err
		}
		request.TimeEnds = &oci_common.SDKTime{Time: tmp}
	}

	if timeStarts, ok := s.D.GetOkExists("time_starts"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStarts.(string))
		if err != nil {
			return err
		}
		request.TimeStarts = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_scheduler")

	response, err := s.Client.CreateSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getScheduleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_scheduler"), oci_resource_scheduler.ActionTypeRelated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ResourceSchedulerScheduleResourceCrud) getScheduleFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_resource_scheduler.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	scheduleId, err := scheduleWaitForWorkRequest(workId, "resourceschedule",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, scheduleId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_resource_scheduler.CancelWorkRequestRequest{
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
	s.D.SetId(*scheduleId)

	return s.Get()
}

func scheduleWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "resource_scheduler", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_resource_scheduler.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func scheduleWaitForWorkRequest(wId *string, entityType string, action oci_resource_scheduler.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_resource_scheduler.ScheduleClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "resource_scheduler")
	retryPolicy.ShouldRetryOperation = scheduleWorkRequestShouldRetryFunc(timeout)

	response := oci_resource_scheduler.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_resource_scheduler.OperationStatusInProgress),
			string(oci_resource_scheduler.OperationStatusAccepted),
			string(oci_resource_scheduler.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_resource_scheduler.OperationStatusSucceeded),
			string(oci_resource_scheduler.OperationStatusFailed),
			string(oci_resource_scheduler.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_resource_scheduler.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_resource_scheduler.OperationStatusFailed || response.Status == oci_resource_scheduler.OperationStatusCanceled {
		return nil, getErrorFromResourceSchedulerScheduleWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromResourceSchedulerScheduleWorkRequest(client *oci_resource_scheduler.ScheduleClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_resource_scheduler.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_resource_scheduler.ListWorkRequestErrorsRequest{
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

func (s *ResourceSchedulerScheduleResourceCrud) Get() error {
	request := oci_resource_scheduler.GetScheduleRequest{}

	tmp := s.D.Id()
	request.ScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_scheduler")

	response, err := s.Client.GetSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Schedule
	return nil
}

func (s *ResourceSchedulerScheduleResourceCrud) Update() error {
	request := oci_resource_scheduler.UpdateScheduleRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		request.Action = oci_resource_scheduler.UpdateScheduleDetailsActionEnum(action.(string))
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

	if recurrenceDetails, ok := s.D.GetOkExists("recurrence_details"); ok {
		tmp := recurrenceDetails.(string)
		request.RecurrenceDetails = &tmp
	}

	if recurrenceType, ok := s.D.GetOkExists("recurrence_type"); ok {
		request.RecurrenceType = oci_resource_scheduler.UpdateScheduleDetailsRecurrenceTypeEnum(recurrenceType.(string))
	}

	if resourceFilters, ok := s.D.GetOkExists("resource_filters"); ok {
		interfaces := resourceFilters.([]interface{})
		tmp := make([]oci_resource_scheduler.ResourceFilter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_filters", stateDataIndex)
			converted, err := s.mapToResourceFilter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("resource_filters") {
			request.ResourceFilters = tmp
		}
	}

	if resources, ok := s.D.GetOkExists("resources"); ok {
		interfaces := resources.([]interface{})
		tmp := make([]oci_resource_scheduler.Resource, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resources", stateDataIndex)
			converted, err := s.mapToResource(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("resources") {
			request.Resources = tmp
		}
	}

	tmp := s.D.Id()
	request.ScheduleId = &tmp

	if timeEnds, ok := s.D.GetOkExists("time_ends"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnds.(string))
		if err != nil {
			return err
		}
		request.TimeEnds = &oci_common.SDKTime{Time: tmp}
	}

	if timeStarts, ok := s.D.GetOkExists("time_starts"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStarts.(string))
		if err != nil {
			return err
		}
		request.TimeStarts = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_scheduler")

	response, err := s.Client.UpdateSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getScheduleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_scheduler"), oci_resource_scheduler.ActionTypeRelated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ResourceSchedulerScheduleResourceCrud) Delete() error {
	request := oci_resource_scheduler.DeleteScheduleRequest{}

	tmp := s.D.Id()
	request.ScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_scheduler")

	_, err := s.Client.DeleteSchedule(context.Background(), request)
	return err
}

func (s *ResourceSchedulerScheduleResourceCrud) SetData() error {
	s.D.Set("action", s.Res.Action)

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

	if s.Res.RecurrenceDetails != nil {
		s.D.Set("recurrence_details", *s.Res.RecurrenceDetails)
	}

	s.D.Set("recurrence_type", s.Res.RecurrenceType)

	resourceFilters := []interface{}{}
	for _, item := range s.Res.ResourceFilters {
		resourceFilters = append(resourceFilters, ResourceFilterToMap(item))
	}
	s.D.Set("resource_filters", resourceFilters)

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, ResourceToMap(item))
	}
	s.D.Set("resources", resources)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeEnds != nil {
		s.D.Set("time_ends", s.Res.TimeEnds.Format(time.RFC3339Nano))
	}

	if s.Res.TimeLastRun != nil {
		s.D.Set("time_last_run", s.Res.TimeLastRun.String())
	}

	if s.Res.TimeNextRun != nil {
		s.D.Set("time_next_run", s.Res.TimeNextRun.String())
	}

	if s.Res.TimeStarts != nil {
		s.D.Set("time_starts", s.Res.TimeStarts.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *ResourceSchedulerScheduleResourceCrud) StartSchedule() error {
	request := oci_resource_scheduler.ActivateScheduleRequest{}

	idTmp := s.D.Id()
	request.ScheduleId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_scheduler")

	_, err := s.Client.ActivateSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_resource_scheduler.ScheduleLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ResourceSchedulerScheduleResourceCrud) StopSchedule() error {
	request := oci_resource_scheduler.DeactivateScheduleRequest{}

	idTmp := s.D.Id()
	request.ScheduleId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_scheduler")

	_, err := s.Client.DeactivateSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_resource_scheduler.ScheduleLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ResourceSchedulerScheduleResourceCrud) mapToDefinedTagFilterValue(fieldKeyFormat string) (oci_resource_scheduler.DefinedTagFilterValue, error) {
	result := oci_resource_scheduler.DefinedTagFilterValue{}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if tagKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag_key")); ok {
		tmp := tagKey.(string)
		result.TagKey = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func DefinedTagFilterValueToMap(obj oci_resource_scheduler.DefinedTagFilterValue) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.TagKey != nil {
		result["tag_key"] = string(*obj.TagKey)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *ResourceSchedulerScheduleResourceCrud) mapToResource(fieldKeyFormat string) (oci_resource_scheduler.Resource, error) {
	result := oci_resource_scheduler.Resource{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		result.Metadata = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	return result, nil
}

func ResourceToMap(obj oci_resource_scheduler.Resource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["metadata"] = obj.Metadata

	return result
}

func (s *ResourceSchedulerScheduleResourceCrud) mapToResourceFilter(fieldKeyFormat string) (oci_resource_scheduler.ResourceFilter, error) {
	var baseObject oci_resource_scheduler.ResourceFilter
	//discriminator
	attributeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute"))
	var attribute string
	if ok {
		attribute = attributeRaw.(string)
	} else {
		attribute = "" // default value
	}
	switch strings.ToLower(attribute) {
	case strings.ToLower("COMPARTMENT_ID"):
		details := oci_resource_scheduler.CompartmentIdResourceFilter{}
		if shouldIncludeChildCompartments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_include_child_compartments")); ok {
			tmp := shouldIncludeChildCompartments.(bool)
			details.ShouldIncludeChildCompartments = &tmp
		}
		details.Value = mapToStringFilter(s, fieldKeyFormat)
		baseObject = details
	case strings.ToLower("TIME_CREATED"):
		details := oci_resource_scheduler.TimeCreatedResourceFilter{}
		if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
			details.Condition = oci_resource_scheduler.TimeCreatedResourceFilterConditionEnum(condition.(string))
		}
		details.Value = mapToStringFilter(s, fieldKeyFormat)
		baseObject = details
	case strings.ToLower("LIFECYCLE_STATE"):
		details := oci_resource_scheduler.LifecycleStateResourceFilter{}
		details.Value = mapToStringArrayFilter(s, fieldKeyFormat)
		baseObject = details
	case strings.ToLower("RESOURCE_TYPE"):
		details := oci_resource_scheduler.ResourceTypeResourceFilter{}
		details.Value = mapToStringArrayFilter(s, fieldKeyFormat)
		baseObject = details
	case strings.ToLower("DEFINED_TAGS"):
		details := oci_resource_scheduler.DefinedTagsResourceFilter{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			interfaces := value.([]interface{})
			tmp := make([]oci_resource_scheduler.DefinedTagFilterValue, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "value"), stateDataIndex)
				converted, err := s.mapToDefinedTagFilterValue(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "value")) {
				details.Value = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown attribute '%v' was specified", attribute)
	}
	return baseObject, nil
}

func mapToStringFilter(s *ResourceSchedulerScheduleResourceCrud, fieldKeyFormat string) *string {
	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		interfaces := value.([]interface{})
		if len(interfaces) > 0 {
			if valMap, ok := interfaces[0].(map[string]interface{}); ok {
				tmp := valMap["value"].(string)
				return &tmp
			}
		}
	}
	return nil
}

// mapToStringArrayFilter extracts a slice of strings from a map stored in a ResourceSchedulerScheduleResourceCrud object.
// It looks for the key "value" within the maps contained in the interfaces slice.
//
// Example:
//
//	Input:  interfaces slice with maps containing "value" keys
//	Output: []string{"instance", "anotherValue"}
func mapToStringArrayFilter(s *ResourceSchedulerScheduleResourceCrud, fieldKeyFormat string) []string {
	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		interfaces := value.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			//interfaces[i] map[namespace: tag_key: value:instance]
			if valMap, ok := interfaces[i].(map[string]interface{}); ok {
				tmp[i] = valMap["value"].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "value")) {
			return tmp
		}
	}
	return nil
}

// stringArrayFilterToMap converts a slice of strings into a slice of interfaces,
// where each string is wrapped in a map with a single key "value".
// It logs the resulting slice for debugging purposes.
//
// Example:
//
//	Input:  []string{"item1", "item2"}
//	Output: []interface{}{
//	           map[string]interface{}{"value": "item1"},
//	           map[string]interface{}{"value": "item2"},
//	        }
func stringArrayFilterToMap(values []string) []interface{} {
	result := make([]interface{}, len(values))
	for i, item := range values {
		result[i] = map[string]interface{}{"value": item}
	}
	return result
}

// stringFilterToMap converts a pointer to a string into a slice of interfaces containing a map.
// The map has a single key "value" with the string's dereferenced value.
// If the input pointer is nil, the function returns nil.
//
// Examples:
//   - Non-nil string pointer ("exampleString") -> [map[value:exampleString]]
//   - Nil string pointer -> []
//   - Empty string ("") -> [map[value:]]
func stringFilterToMap(value *string) []interface{} {
	if value == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{"value": *value}}
}

func ResourceFilterToMap(obj oci_resource_scheduler.ResourceFilter) map[string]interface{} {

	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_resource_scheduler.CompartmentIdResourceFilter:
		result["attribute"] = "COMPARTMENT_ID"
		if v.ShouldIncludeChildCompartments != nil {
			result["should_include_child_compartments"] = bool(*v.ShouldIncludeChildCompartments)
		}
		result["value"] = stringFilterToMap(v.Value)
	case oci_resource_scheduler.TimeCreatedResourceFilter:
		result["attribute"] = "TIME_CREATED"
		result["condition"] = string(v.Condition)
		result["value"] = stringFilterToMap(v.Value)
	case oci_resource_scheduler.LifecycleStateResourceFilter:
		result["attribute"] = "LIFECYCLE_STATE"
		// v: { Value=[active creating] }
		result["value"] = stringArrayFilterToMap(v.Value)
	case oci_resource_scheduler.ResourceTypeResourceFilter:
		result["attribute"] = "RESOURCE_TYPE"
		// v: { Value=[instance autonomousdatabase] }
		result["value"] = stringArrayFilterToMap(v.Value)
	case oci_resource_scheduler.DefinedTagsResourceFilter:
		result["attribute"] = "DEFINED_TAGS"

		value := []interface{}{}
		for _, item := range v.Value {
			value = append(value, DefinedTagFilterValueToMap(item))
		}
		result["value"] = value
	default:
		log.Printf("[WARN] Received 'attribute' of unknown type %v", obj)
		return nil
	}
	return result
}

func ScheduleSummaryToMap(obj oci_resource_scheduler.ScheduleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

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

	//result["last_run_status"] = string(obj.LastRunStatus)

	if obj.RecurrenceDetails != nil {
		result["recurrence_details"] = string(*obj.RecurrenceDetails)
	}

	result["recurrence_type"] = string(obj.RecurrenceType)

	resourceFilters := []interface{}{}
	for _, item := range obj.ResourceFilters {
		resourceFilters = append(resourceFilters, ResourceFilterToMap(item))
	}
	result["resource_filters"] = resourceFilters

	resources := []interface{}{}
	for _, item := range obj.Resources {
		resources = append(resources, ResourceToMap(item))
	}
	result["resources"] = resources

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeEnds != nil {
		result["time_ends"] = obj.TimeEnds.String()
	}

	if obj.TimeLastRun != nil {
		result["time_last_run"] = obj.TimeLastRun.String()
	}

	if obj.TimeNextRun != nil {
		result["time_next_run"] = obj.TimeNextRun.String()
	}

	if obj.TimeStarts != nil {
		result["time_starts"] = obj.TimeStarts.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
