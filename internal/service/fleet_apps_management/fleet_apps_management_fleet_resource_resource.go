// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementFleetResourceResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementFleetResource,
		Read:     readFleetAppsManagementFleetResource,
		Update:   updateFleetAppsManagementFleetResource,
		Delete:   deleteFleetAppsManagementFleetResource,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"resource_region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"application_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compliance_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"environment_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_count": {
				Type:     schema.TypeInt,
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
			"target_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tenancy_name": {
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

func createFleetAppsManagementFleetResource(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetResourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementFleetResource(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetResourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementFleetResource(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetResourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementFleetResource(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetResourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementFleetResourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.FleetResource
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementFleetResourceResourceCrud) ID() string {
	return GetFleetResourceCompositeId(s.D.Get("fleet_id").(string), s.D.Get("resource_id").(string))
}

func (s *FleetAppsManagementFleetResourceResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementFleetResourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.FleetResourceLifecycleStateActive),
		string(oci_fleet_apps_management.FleetResourceLifecycleStateNeedsAttention),
	}
}

func (s *FleetAppsManagementFleetResourceResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementFleetResourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.FleetResourceLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementFleetResourceResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateFleetResourceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if resourceRegion, ok := s.D.GetOkExists("resource_region"); ok {
		tmp := resourceRegion.(string)
		request.ResourceRegion = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateFleetResource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFleetResourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetAppsManagementFleetResourceResourceCrud) getFleetResourceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fleetResourceId, err := fleetResourceWaitForWorkRequest(workId, "fleetresource",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*fleetResourceId)

	return s.Get()
}

func fleetResourceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func fleetResourceWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = fleetResourceWorkRequestShouldRetryFunc(timeout)

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
			if res.ActionType == oci_fleet_apps_management.ActionTypeUpdated {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_fleet_apps_management.OperationStatusFailed || response.Status == oci_fleet_apps_management.OperationStatusCanceled {
		return nil, getErrorFromFleetAppsManagementFleetResourceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementFleetResourceWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
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

func (s *FleetAppsManagementFleetResourceResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetFleetResourceRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.FleetResourceId = &tmp
	}

	fleetId, fleetResourceId, err := parseFleetResourceCompositeId(s.D.Id())
	if err == nil {
		request.FleetId = &fleetId
		request.FleetResourceId = &fleetResourceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetFleetResource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FleetResource
	return nil
}

func (s *FleetAppsManagementFleetResourceResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdateFleetResourceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	tmp := s.D.Id()
	request.FleetResourceId = &tmp

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateFleetResource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFleetResourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementFleetResourceResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteFleetResourceRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	tmp := s.D.Id()
	request.FleetResourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeleteFleetResource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := fleetResourceWaitForWorkRequest(workId, "fleetresource",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *FleetAppsManagementFleetResourceResourceCrud) SetData() error {

	fleetId, fleetResourceId, err := parseFleetResourceCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("fleet_id", &fleetId)
		s.D.SetId(fleetResourceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ApplicationType != nil {
		s.D.Set("application_type", *s.Res.ApplicationType)
	}

	if s.Res.Compartment != nil {
		s.D.Set("compartment", *s.Res.Compartment)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("compliance_state", s.Res.ComplianceState)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EnvironmentType != nil {
		s.D.Set("environment_type", *s.Res.EnvironmentType)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Product != nil {
		s.D.Set("product", *s.Res.Product)
	}

	if s.Res.ProductCount != nil {
		s.D.Set("product_count", *s.Res.ProductCount)
	}

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	} else {
		// FAMS API sometimes returns null rather than {} for empty system_tags.
		systemTags := map[string]interface{}{}
		s.D.Set("system_tags", systemTags)
	}

	if s.Res.TargetCount != nil {
		s.D.Set("target_count", *s.Res.TargetCount)
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TenancyName != nil {
		s.D.Set("tenancy_name", *s.Res.TenancyName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetFleetResourceCompositeId(fleetId string, fleetResourceId string) string {
	fleetId = url.PathEscape(fleetId)
	fleetResourceId = url.PathEscape(fleetResourceId)
	compositeId := "fleets/" + fleetId + "/fleetResources/" + fleetResourceId
	return compositeId
}

func parseFleetResourceCompositeId(compositeId string) (fleetId string, fleetResourceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("fleets/.*/fleetResources/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	fleetId, _ = url.PathUnescape(parts[1])
	fleetResourceId, _ = url.PathUnescape(parts[3])

	return
}

func FleetResourceSummaryToMap(obj oci_fleet_apps_management.FleetResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationType != nil {
		result["application_type"] = string(*obj.ApplicationType)
	}

	if obj.Compartment != nil {
		result["compartment"] = string(*obj.Compartment)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["compliance_state"] = string(obj.ComplianceState)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EnvironmentType != nil {
		result["environment_type"] = string(*obj.EnvironmentType)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Product != nil {
		result["product"] = string(*obj.Product)
	}

	if obj.ProductCount != nil {
		result["product_count"] = int(*obj.ProductCount)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetCount != nil {
		result["target_count"] = int(*obj.TargetCount)
	}

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	if obj.TenancyName != nil {
		result["tenancy_name"] = string(*obj.TenancyName)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
