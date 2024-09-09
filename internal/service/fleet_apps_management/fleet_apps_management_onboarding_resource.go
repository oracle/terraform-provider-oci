// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"fmt"

	//"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementOnboardingResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementOnboarding,
		Read:     readFleetAppsManagementOnboarding,
		Delete:   deleteFleetAppsManagementOnboarding,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"is_cost_tracking_tag_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_fams_tag_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFleetAppsManagementOnboarding(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementOnboardingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementOnboarding(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementOnboardingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func deleteFleetAppsManagementOnboarding(d *schema.ResourceData, m interface{}) error {
	return nil
}

type FleetAppsManagementOnboardingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementAdminClient
	FleetClient            *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.Onboarding
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementOnboardingResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementOnboardingResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_apps_management.OnboardingLifecycleStateCreating),
	}
}

func (s *FleetAppsManagementOnboardingResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.OnboardingLifecycleStateActive),
		string(oci_fleet_apps_management.OnboardingLifecycleStateNeedsAttention),
	}
}

func (s *FleetAppsManagementOnboardingResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.OnboardingLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementOnboardingResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.OnboardingLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementOnboardingResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateOnboardingRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isCostTrackingTagEnabled, ok := s.D.GetOkExists("is_cost_tracking_tag_enabled"); ok {
		tmp := isCostTrackingTagEnabled.(bool)
		request.IsCostTrackingTagEnabled = &tmp
	}

	if isFamsTagEnabled, ok := s.D.GetOkExists("is_fams_tag_enabled"); ok {
		tmp := isFamsTagEnabled.(bool)
		request.IsFamsTagEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateOnboarding(context.Background(), request)
	if err != nil {
		return err
	}

	s.D.SetId(*response.Id)

	return s.Get()
	// This does not return a work request.
}

func (s *FleetAppsManagementOnboardingResourceCrud) getOnboardingFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	onboardingId, err := onboardingWaitForWorkRequest(workId, "onboarding",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.FleetClient)

	if err != nil {
		return err
	}
	s.D.SetId(*onboardingId)

	return s.Get()
}

func onboardingWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func onboardingWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = onboardingWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromFleetAppsManagementOnboardingWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementOnboardingWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
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

func (s *FleetAppsManagementOnboardingResourceCrud) Get() error {

	// FAMS 1.0 API does not have a GET for Onboarding. We need to do a List with the ID
	//  and convert the Summary.  FAMS 1.2 API will add the appropriate GET call.

	request := oci_fleet_apps_management.ListOnboardingsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	tmp := s.D.Id()
	request.Id = &tmp

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_apps_management.OnboardingLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.ListOnboardings(context.Background(), request)
	if err != nil {
		return err
	}

	if len(response.OnboardingCollection.Items) > 0 {
		realResponse := onboardingSummaryToOnboarding(response.OnboardingCollection.Items[0])
		s.Res = &realResponse
	} else {
		return fmt.Errorf("onboarding id %s not found", s.D.Id())
	}

	return nil
}

func onboardingSummaryToOnboarding(obj oci_fleet_apps_management.OnboardingSummary) oci_fleet_apps_management.Onboarding {
	result := oci_fleet_apps_management.Onboarding{}
	result.Id = obj.Id
	result.CompartmentId = obj.CompartmentId
	result.IsFamsTagEnabled = obj.IsFamsTagEnabled
	result.IsCostTrackingTagEnabled = obj.IsCostTrackingTagEnabled
	result.ResourceRegion = obj.ResourceRegion
	result.TimeCreated = obj.TimeCreated
	result.TimeUpdated = obj.TimeUpdated
	result.LifecycleState = obj.LifecycleState
	result.SystemTags = obj.SystemTags
	result.Version = obj.Version

	return result
}

func (s *FleetAppsManagementOnboardingResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	if s.Res.IsCostTrackingTagEnabled != nil {
		s.D.Set("is_cost_tracking_tag_enabled", *s.Res.IsCostTrackingTagEnabled)
	}

	if s.Res.IsFamsTagEnabled != nil {
		s.D.Set("is_fams_tag_enabled", *s.Res.IsFamsTagEnabled)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	} else {
		// FAMS API sometimes returns null rather than {} for empty system_tags.
		systemTags := map[string]interface{}{}
		s.D.Set("system_tags", systemTags)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func OnboardingSummaryToMap(obj oci_fleet_apps_management.OnboardingSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsCostTrackingTagEnabled != nil {
		result["is_cost_tracking_tag_enabled"] = bool(*obj.IsCostTrackingTagEnabled)
	}

	if obj.IsFamsTagEnabled != nil {
		result["is_fams_tag_enabled"] = bool(*obj.IsFamsTagEnabled)
	}

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
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

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
