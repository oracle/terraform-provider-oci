// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementMaintenanceWindowResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementMaintenanceWindow,
		Read:     readFleetAppsManagementMaintenanceWindow,
		Update:   updateFleetAppsManagementMaintenanceWindow,
		Delete:   deleteFleetAppsManagementMaintenanceWindow,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"duration": {
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
			"is_outage": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_recurring": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"maintenance_window_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recurrences": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"task_initiation_cutoff": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"time_schedule_start": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_region": {
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

func createFleetAppsManagementMaintenanceWindow(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementMaintenanceWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementMaintenanceWindowClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementMaintenanceWindow(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementMaintenanceWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementMaintenanceWindowClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementMaintenanceWindow(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementMaintenanceWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementMaintenanceWindowClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementMaintenanceWindow(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementMaintenanceWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementMaintenanceWindowClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementMaintenanceWindowResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementMaintenanceWindowClient
	FleetClient            *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.MaintenanceWindow
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.MaintenanceWindowLifecycleStateActive),
		string(oci_fleet_apps_management.MaintenanceWindowLifecycleStateNeedsAttention),
	}
}

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.MaintenanceWindowLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.MaintenanceWindowLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateMaintenanceWindowRequest{}

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

	if duration, ok := s.D.GetOkExists("duration"); ok {
		tmp := duration.(string)
		request.Duration = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isOutage, ok := s.D.GetOkExists("is_outage"); ok {
		tmp := isOutage.(bool)
		request.IsOutage = &tmp
	}

	if isRecurring, ok := s.D.GetOkExists("is_recurring"); ok {
		tmp := isRecurring.(bool)
		request.IsRecurring = &tmp
	}

	if maintenanceWindowType, ok := s.D.GetOkExists("maintenance_window_type"); ok {
		request.MaintenanceWindowType = oci_fleet_apps_management.MaintenanceWindowTypeEnum(maintenanceWindowType.(string))
	}

	if recurrences, ok := s.D.GetOkExists("recurrences"); ok {
		tmp := recurrences.(string)
		request.Recurrences = &tmp
	}

	if taskInitiationCutoff, ok := s.D.GetOkExists("task_initiation_cutoff"); ok {
		tmp := taskInitiationCutoff.(int)
		request.TaskInitiationCutoff = &tmp
	}

	if timeScheduleStart, ok := s.D.GetOkExists("time_schedule_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduleStart.(string))
		if err != nil {
			return err
		}
		request.TimeScheduleStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateMaintenanceWindow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaintenanceWindow
	return nil
}

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) getMaintenanceWindowFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	maintenanceWindowId, err := maintenanceWindowWaitForWorkRequest(workId, "maintenance-window",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.FleetClient)

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
		if tfresource.ShouldRetry(response, false, "fleet_apps_management", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fleet_apps_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func maintenanceWindowWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = maintenanceWindowWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_apps_management.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fleet_apps_management.OperationStatusInProgress),
			string(oci_fleet_apps_management.OperationStatusAccepted),
			string(oci_fleet_apps_management.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_fleet_apps_management.OperationStatusSucceeded),
			string(oci_fleet_apps_management.OperationStatusFailed),
			string(oci_fleet_apps_management.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_fleet_apps_management.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_fleet_apps_management.OperationStatusFailed || response.Status == oci_fleet_apps_management.OperationStatusCanceled {
		return nil, getErrorFromFleetAppsManagementMaintenanceWindowWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementMaintenanceWindowWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fleet_apps_management.ListWorkRequestErrorsRequest{
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

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetMaintenanceWindowRequest{}

	tmp := s.D.Id()
	request.MaintenanceWindowId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetMaintenanceWindow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaintenanceWindow
	return nil
}

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdateMaintenanceWindowRequest{}

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

	if duration, ok := s.D.GetOkExists("duration"); ok {
		tmp := duration.(string)
		request.Duration = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isOutage, ok := s.D.GetOkExists("is_outage"); ok {
		tmp := isOutage.(bool)
		request.IsOutage = &tmp
	}

	if isRecurring, ok := s.D.GetOkExists("is_recurring"); ok {
		tmp := isRecurring.(bool)
		request.IsRecurring = &tmp
	}

	tmp := s.D.Id()
	request.MaintenanceWindowId = &tmp

	if maintenanceWindowType, ok := s.D.GetOkExists("maintenance_window_type"); ok {
		request.MaintenanceWindowType = oci_fleet_apps_management.MaintenanceWindowTypeEnum(maintenanceWindowType.(string))
	}

	if recurrences, ok := s.D.GetOkExists("recurrences"); ok {
		tmp := recurrences.(string)
		request.Recurrences = &tmp
	}

	if taskInitiationCutoff, ok := s.D.GetOkExists("task_initiation_cutoff"); ok {
		tmp := taskInitiationCutoff.(int)
		request.TaskInitiationCutoff = &tmp
	}

	if timeScheduleStart, ok := s.D.GetOkExists("time_schedule_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduleStart.(string))
		if err != nil {
			return err
		}
		request.TimeScheduleStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateMaintenanceWindow(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMaintenanceWindowFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteMaintenanceWindowRequest{}

	tmp := s.D.Id()
	request.MaintenanceWindowId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeleteMaintenanceWindow(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := maintenanceWindowWaitForWorkRequest(workId, "maintenance-window",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.FleetClient)
	return delWorkRequestErr
}

func (s *FleetAppsManagementMaintenanceWindowResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
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

	if s.Res.Duration != nil {
		s.D.Set("duration", *s.Res.Duration)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsOutage != nil {
		s.D.Set("is_outage", *s.Res.IsOutage)
	}

	if s.Res.IsRecurring != nil {
		s.D.Set("is_recurring", *s.Res.IsRecurring)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("maintenance_window_type", s.Res.MaintenanceWindowType)

	if s.Res.Recurrences != nil {
		s.D.Set("recurrences", *s.Res.Recurrences)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TaskInitiationCutoff != nil {
		s.D.Set("task_initiation_cutoff", *s.Res.TaskInitiationCutoff)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeScheduleStart != nil {
		s.D.Set("time_schedule_start", s.Res.TimeScheduleStart.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func MaintenanceWindowSummaryToMap(obj oci_fleet_apps_management.MaintenanceWindowSummary) map[string]interface{} {
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

	if obj.Duration != nil {
		result["duration"] = string(*obj.Duration)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsOutage != nil {
		result["is_outage"] = bool(*obj.IsOutage)
	}

	if obj.IsRecurring != nil {
		result["is_recurring"] = bool(*obj.IsRecurring)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["maintenance_window_type"] = string(obj.MaintenanceWindowType)

	if obj.Recurrences != nil {
		result["recurrences"] = string(*obj.Recurrences)
	}

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TaskInitiationCutoff != nil {
		result["task_initiation_cutoff"] = int(*obj.TaskInitiationCutoff)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeScheduleStart != nil {
		result["time_schedule_start"] = obj.TimeScheduleStart.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
