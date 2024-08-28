// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

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
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMaintenanceWindowResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMaintenanceWindow,
		Read:     readStackMonitoringMaintenanceWindow,
		Update:   updateStackMonitoringMaintenanceWindow,
		Delete:   deleteStackMonitoringMaintenanceWindow,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resources": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"resource_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"are_members_included": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"schedule": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"schedule_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ONE_TIME",
								"RECURRENT",
							}, true),
						},

						// Optional
						"maintenance_window_duration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maintenance_window_recurrences": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_maintenance_window_end": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"time_maintenance_window_start": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resources_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"number_of_members": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"resource_id": {
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

func StackMonitoringMaintenanceWindowSummaryResponse() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMaintenanceWindow,
		Read:     readStackMonitoringMaintenanceWindow,
		Update:   updateStackMonitoringMaintenanceWindow,
		Delete:   deleteStackMonitoringMaintenanceWindow,
		Schema: map[string]*schema.Schema{
			// Required
			"id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"number_of_resources": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"lifecycle_state": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"operation_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"operation_status": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schedule": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"schedule_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ONE_TIME",
								"RECURRENT",
							}, true),
						},

						// Optional
						"maintenance_window_duration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maintenance_window_recurrences": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_maintenance_window_end": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"time_maintenance_window_start": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			// remove
			"state": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createStackMonitoringMaintenanceWindow(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMaintenanceWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMaintenanceWindow(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMaintenanceWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateStackMonitoringMaintenanceWindow(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMaintenanceWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringMaintenanceWindow(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMaintenanceWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringMaintenanceWindowResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.MaintenanceWindow
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_stack_monitoring.MaintenanceWindowLifecycleStateCreating),
	}
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.MaintenanceWindowLifecycleStateActive),
		string(oci_stack_monitoring.MaintenanceWindowLifecycleStateNeedsAttention),
	}
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_stack_monitoring.MaintenanceWindowLifecycleStateDeleting),
	}
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_stack_monitoring.MaintenanceWindowLifecycleStateDeleted),
	}
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateMaintenanceWindowRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if resources, ok := s.D.GetOkExists("resources"); ok {
		interfaces := resources.([]interface{})
		tmp := make([]oci_stack_monitoring.CreateMaintenanceWindowResourceDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resources", stateDataIndex)
			converted, err := s.mapToCreateMaintenanceWindowResourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("resources") {
			request.Resources = tmp
		}
	}

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		if tmpList := schedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedule", 0)
			tmp, err := s.mapToMaintenanceWindowSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Schedule = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateMaintenanceWindow(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getMaintenanceWindowFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) getMaintenanceWindowFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_stack_monitoring.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	maintenanceWindowId, err := maintenanceWindowWaitForWorkRequest(workId, "maintenancewindow",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*maintenanceWindowId)

	return s.Get()
}

func maintenanceWindowWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "stack_monitoring", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_stack_monitoring.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func maintenanceWindowWaitForWorkRequest(wId *string, entityType string, action oci_stack_monitoring.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_stack_monitoring.StackMonitoringClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "stack_monitoring")
	retryPolicy.ShouldRetryOperation = maintenanceWindowWorkRequestShouldRetryFunc(timeout)

	response := oci_stack_monitoring.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_stack_monitoring.OperationStatusInProgress),
			string(oci_stack_monitoring.OperationStatusAccepted),
			string(oci_stack_monitoring.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_stack_monitoring.OperationStatusSucceeded),
			string(oci_stack_monitoring.OperationStatusFailed),
			string(oci_stack_monitoring.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_stack_monitoring.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_stack_monitoring.OperationStatusFailed || response.Status == oci_stack_monitoring.OperationStatusCanceled {
		return nil, getErrorFromStackMonitoringMaintenanceWindowWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromStackMonitoringMaintenanceWindowWorkRequest(client *oci_stack_monitoring.StackMonitoringClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_stack_monitoring.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_stack_monitoring.ListWorkRequestErrorsRequest{
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

func (s *StackMonitoringMaintenanceWindowResourceCrud) Get() error {
	request := oci_stack_monitoring.GetMaintenanceWindowRequest{}

	tmp := s.D.Id()
	request.MaintenanceWindowId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetMaintenanceWindow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaintenanceWindow
	return nil
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) Update() error {
	request := oci_stack_monitoring.UpdateMaintenanceWindowRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.MaintenanceWindowId = &tmp

	if resources, ok := s.D.GetOkExists("resources"); ok {
		interfaces := resources.([]interface{})
		tmp := make([]oci_stack_monitoring.CreateMaintenanceWindowResourceDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resources", stateDataIndex)
			converted, err := s.mapToCreateMaintenanceWindowResourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("resources") {
			request.Resources = tmp
		}
	}

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		if tmpList := schedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedule", 0)
			tmp, err := s.mapToMaintenanceWindowSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Schedule = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UpdateMaintenanceWindow(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMaintenanceWindowFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) Delete() error {
	request := oci_stack_monitoring.DeleteMaintenanceWindowRequest{}

	tmp := s.D.Id()
	request.MaintenanceWindowId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.DeleteMaintenanceWindow(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := maintenanceWindowWaitForWorkRequest(workId, "maintenancewindow",
		oci_stack_monitoring.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, CreateMaintenanceWindowResourceDetailsToMap(item))
	}
	s.D.Set("resources", resources)

	resourcesDetails := []interface{}{}
	for _, item := range s.Res.ResourcesDetails {
		resourcesDetails = append(resourcesDetails, MonitoredResourceDetailsToMap(item))
	}
	s.D.Set("resources_details", resourcesDetails)

	if s.Res.Schedule != nil {
		scheduleArray := []interface{}{}
		if scheduleMap := MaintenanceWindowScheduleToMap(&s.Res.Schedule); scheduleMap != nil {
			scheduleArray = append(scheduleArray, scheduleMap)
		}
		s.D.Set("schedule", scheduleArray)
	} else {
		s.D.Set("schedule", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) mapToCreateMaintenanceWindowResourceDetails(fieldKeyFormat string) (oci_stack_monitoring.CreateMaintenanceWindowResourceDetails, error) {
	result := oci_stack_monitoring.CreateMaintenanceWindowResourceDetails{}

	if areMembersIncluded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_members_included")); ok {
		tmp := areMembersIncluded.(bool)
		result.AreMembersIncluded = &tmp
	}

	if resourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_id")); ok {
		tmp := resourceId.(string)
		result.ResourceId = &tmp
	}

	return result, nil
}

func CreateMaintenanceWindowResourceDetailsToMap(obj oci_stack_monitoring.CreateMaintenanceWindowResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreMembersIncluded != nil {
		result["are_members_included"] = bool(*obj.AreMembersIncluded)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	return result
}

func (s *StackMonitoringMaintenanceWindowResourceCrud) mapToMaintenanceWindowSchedule(fieldKeyFormat string) (oci_stack_monitoring.MaintenanceWindowSchedule, error) {
	var baseObject oci_stack_monitoring.MaintenanceWindowSchedule
	//discriminator
	scheduleTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_type"))
	var scheduleType string
	if ok {
		scheduleType = scheduleTypeRaw.(string)
	} else {
		scheduleType = "" // default value
	}
	switch strings.ToLower(scheduleType) {
	case strings.ToLower("ONE_TIME"):
		details := oci_stack_monitoring.OneTimeMaintenanceWindowSchedule{}
		if timeMaintenanceWindowEnd, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_maintenance_window_end")); ok {
			tmp, err := time.Parse(time.RFC3339, timeMaintenanceWindowEnd.(string))
			if err != nil {
				return details, err
			}
			details.TimeMaintenanceWindowEnd = &oci_common.SDKTime{Time: tmp}
		}
		if timeMaintenanceWindowStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_maintenance_window_start")); ok {
			tmp, err := time.Parse(time.RFC3339, timeMaintenanceWindowStart.(string))
			if err != nil {
				return details, err
			}
			details.TimeMaintenanceWindowStart = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	case strings.ToLower("RECURRENT"):
		details := oci_stack_monitoring.RecurrentMaintenanceWindowSchedule{}
		if maintenanceWindowDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maintenance_window_duration")); ok {
			tmp := maintenanceWindowDuration.(string)
			details.MaintenanceWindowDuration = &tmp
		}
		if maintenanceWindowRecurrences, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maintenance_window_recurrences")); ok {
			tmp := maintenanceWindowRecurrences.(string)
			details.MaintenanceWindowRecurrences = &tmp
		}
		if timeMaintenanceWindowEnd, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_maintenance_window_end")); ok {
			tmp, err := time.Parse(time.RFC3339, timeMaintenanceWindowEnd.(string))
			if err != nil {
				return details, err
			}
			details.TimeMaintenanceWindowEnd = &oci_common.SDKTime{Time: tmp}
		}
		if timeMaintenanceWindowStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_maintenance_window_start")); ok {
			tmp, err := time.Parse(time.RFC3339, timeMaintenanceWindowStart.(string))
			if err != nil {
				return details, err
			}
			details.TimeMaintenanceWindowStart = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown schedule_type '%v' was specified", scheduleType)
	}
	return baseObject, nil
}

func MaintenanceWindowScheduleToMap(obj *oci_stack_monitoring.MaintenanceWindowSchedule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_stack_monitoring.OneTimeMaintenanceWindowSchedule:
		result["schedule_type"] = "ONE_TIME"

		if v.TimeMaintenanceWindowEnd != nil {
			result["time_maintenance_window_end"] = v.TimeMaintenanceWindowEnd.Format(time.RFC3339Nano)
		}

		if v.TimeMaintenanceWindowStart != nil {
			result["time_maintenance_window_start"] = v.TimeMaintenanceWindowStart.Format(time.RFC3339Nano)
		}
	case oci_stack_monitoring.RecurrentMaintenanceWindowSchedule:
		result["schedule_type"] = "RECURRENT"

		if v.MaintenanceWindowDuration != nil {
			result["maintenance_window_duration"] = string(*v.MaintenanceWindowDuration)
		}

		if v.MaintenanceWindowRecurrences != nil {
			result["maintenance_window_recurrences"] = string(*v.MaintenanceWindowRecurrences)
		}

		if v.TimeMaintenanceWindowEnd != nil {
			result["time_maintenance_window_end"] = v.TimeMaintenanceWindowEnd.Format(time.RFC3339Nano)
		}

		if v.TimeMaintenanceWindowStart != nil {
			result["time_maintenance_window_start"] = v.TimeMaintenanceWindowStart.Format(time.RFC3339Nano)
		}
	default:
		log.Printf("[WARN] Received 'schedule_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func MaintenanceWindowSummaryToMap(obj oci_stack_monitoring.MaintenanceWindowSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NumberOfResources != nil {
		result["number_of_resources"] = int(*obj.NumberOfResources)
	}

	result["operation_status"] = string(obj.OperationStatus)

	result["operation_type"] = string(obj.OperationType)

	if obj.Schedule != nil {
		scheduleArray := []interface{}{}
		if scheduleMap := MaintenanceWindowScheduleToMap(&obj.Schedule); scheduleMap != nil {
			scheduleArray = append(scheduleArray, scheduleMap)
		}
		result["schedule"] = scheduleArray
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}

func MonitoredResourceDetailsToMap(obj oci_stack_monitoring.MonitoredResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NumberOfMembers != nil {
		result["number_of_members"] = int(*obj.NumberOfMembers)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}
