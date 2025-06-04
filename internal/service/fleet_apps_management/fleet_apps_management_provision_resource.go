// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementProvisionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementProvision,
		Read:     readFleetAppsManagementProvision,
		Update:   updateFleetAppsManagementProvision,
		Delete:   deleteFleetAppsManagementProvision,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config_catalog_item_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"package_catalog_item_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tf_variable_region_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tf_variable_tenancy_id": {
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
			"provision_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tf_variable_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tf_variable_current_user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"config_catalog_item_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_catalog_item_listing_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_catalog_item_listing_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deployed_resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_instance_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"compartment_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"resource_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_provider": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"package_catalog_item_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"package_catalog_item_listing_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"package_catalog_item_listing_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rms_apply_job_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stack_id": {
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
			"tf_outputs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_sensitive": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"output_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"output_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"output_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"output_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createFleetAppsManagementProvision(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementProvisionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementProvisionClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementProvision(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementProvisionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementProvisionClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementProvision(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementProvisionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementProvisionClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementProvision(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementProvisionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementProvisionClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementProvisionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementProvisionClient
	Res                    *oci_fleet_apps_management.Provision
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient
}

func (s *FleetAppsManagementProvisionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementProvisionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_apps_management.ProvisionLifecycleStateCreating),
	}
}

func (s *FleetAppsManagementProvisionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.ProvisionLifecycleStateActive),
	}
}

func (s *FleetAppsManagementProvisionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.ProvisionLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementProvisionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.ProvisionLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementProvisionResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateProvisionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configCatalogItemId, ok := s.D.GetOkExists("config_catalog_item_id"); ok {
		tmp := configCatalogItemId.(string)
		request.ConfigCatalogItemId = &tmp
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

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if packageCatalogItemId, ok := s.D.GetOkExists("package_catalog_item_id"); ok {
		tmp := packageCatalogItemId.(string)
		request.PackageCatalogItemId = &tmp
	}

	if provisionDescription, ok := s.D.GetOkExists("provision_description"); ok {
		tmp := provisionDescription.(string)
		request.ProvisionDescription = &tmp
	}

	if tfVariableCompartmentId, ok := s.D.GetOkExists("tf_variable_compartment_id"); ok {
		tmp := tfVariableCompartmentId.(string)
		request.TfVariableCompartmentId = &tmp
	}

	if tfVariableCurrentUserId, ok := s.D.GetOkExists("tf_variable_current_user_id"); ok {
		tmp := tfVariableCurrentUserId.(string)
		request.TfVariableCurrentUserId = &tmp
	}

	if tfVariableRegionId, ok := s.D.GetOkExists("tf_variable_region_id"); ok {
		tmp := tfVariableRegionId.(string)
		request.TfVariableRegionId = &tmp
	}

	if tfVariableTenancyId, ok := s.D.GetOkExists("tf_variable_tenancy_id"); ok {
		tmp := tfVariableTenancyId.(string)
		request.TfVariableTenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateProvision(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getProvisionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetAppsManagementProvisionResourceCrud) getProvisionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	provisionId, err := provisionWaitForWorkRequest(workId, "provision",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*provisionId)

	return s.Get()
}

func provisionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func provisionWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = provisionWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_apps_management.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
		return nil, getErrorFromFleetAppsManagementProvisionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementProvisionWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
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

func (s *FleetAppsManagementProvisionResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetProvisionRequest{}

	tmp := s.D.Id()
	request.ProvisionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetProvision(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Provision
	return nil
}

func (s *FleetAppsManagementProvisionResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("compartmentId"); ok && s.D.HasChange("compartmentId") {
		err := s.ChangeProvisionCompartment()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_fleet_apps_management.UpdateProvisionRequest{}

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

	if provisionDescription, ok := s.D.GetOkExists("provision_description"); ok {
		tmp := provisionDescription.(string)
		request.ProvisionDescription = &tmp
	}

	tmp := s.D.Id()
	request.ProvisionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateProvision(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getProvisionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementProvisionResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteProvisionRequest{}

	tmp := s.D.Id()
	request.ProvisionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeleteProvision(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := provisionWaitForWorkRequest(workId, "provision",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *FleetAppsManagementProvisionResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigCatalogItemDisplayName != nil {
		s.D.Set("config_catalog_item_display_name", *s.Res.ConfigCatalogItemDisplayName)
	}

	if s.Res.ConfigCatalogItemId != nil {
		s.D.Set("config_catalog_item_id", *s.Res.ConfigCatalogItemId)
	}

	if s.Res.ConfigCatalogItemListingId != nil {
		s.D.Set("config_catalog_item_listing_id", *s.Res.ConfigCatalogItemListingId)
	}

	if s.Res.ConfigCatalogItemListingVersion != nil {
		s.D.Set("config_catalog_item_listing_version", *s.Res.ConfigCatalogItemListingVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	deployedResources := []interface{}{}
	for _, item := range s.Res.DeployedResources {
		deployedResources = append(deployedResources, DeployedResourceDetailsToMap(item))
	}
	s.D.Set("deployed_resources", deployedResources)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FleetId != nil {
		s.D.Set("fleet_id", *s.Res.FleetId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PackageCatalogItemDisplayName != nil {
		s.D.Set("package_catalog_item_display_name", *s.Res.PackageCatalogItemDisplayName)
	}

	if s.Res.PackageCatalogItemId != nil {
		s.D.Set("package_catalog_item_id", *s.Res.PackageCatalogItemId)
	}

	if s.Res.PackageCatalogItemListingId != nil {
		s.D.Set("package_catalog_item_listing_id", *s.Res.PackageCatalogItemListingId)
	}

	if s.Res.PackageCatalogItemListingVersion != nil {
		s.D.Set("package_catalog_item_listing_version", *s.Res.PackageCatalogItemListingVersion)
	}

	if s.Res.ProvisionDescription != nil {
		s.D.Set("provision_description", *s.Res.ProvisionDescription)
	}

	if s.Res.RmsApplyJobId != nil {
		s.D.Set("rms_apply_job_id", *s.Res.RmsApplyJobId)
	}

	if s.Res.StackId != nil {
		s.D.Set("stack_id", *s.Res.StackId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	tfOutputs := []interface{}{}
	for _, item := range s.Res.TfOutputs {
		tfOutputs = append(tfOutputs, JobExecutionDetailsToMap(item))
	}
	s.D.Set("tf_outputs", tfOutputs)

	if s.Res.TfVariableCompartmentId != nil {
		s.D.Set("tf_variable_compartment_id", *s.Res.TfVariableCompartmentId)
	}

	if s.Res.TfVariableCurrentUserId != nil {
		s.D.Set("tf_variable_current_user_id", *s.Res.TfVariableCurrentUserId)
	}

	if s.Res.TfVariableRegionId != nil {
		s.D.Set("tf_variable_region_id", *s.Res.TfVariableRegionId)
	}

	if s.Res.TfVariableTenancyId != nil {
		s.D.Set("tf_variable_tenancy_id", *s.Res.TfVariableTenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *FleetAppsManagementProvisionResourceCrud) ChangeProvisionCompartment() error {
	request := oci_fleet_apps_management.ChangeProvisionCompartmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	idTmp := s.D.Id()
	request.ProvisionId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	_, err := s.Client.ChangeProvisionCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func DeployedResourceDetailsToMap(obj oci_fleet_apps_management.DeployedResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Mode != nil {
		result["mode"] = string(*obj.Mode)
	}

	resourceInstanceList := []interface{}{}
	for _, item := range obj.ResourceInstanceList {
		resourceInstanceList = append(resourceInstanceList, InstanceSummaryToMap(item))
	}
	result["resource_instance_list"] = resourceInstanceList

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.ResourceProvider != nil {
		result["resource_provider"] = string(*obj.ResourceProvider)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	return result
}

func InstanceSummaryToMap(obj oci_fleet_apps_management.InstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.State != nil {
		result["state"] = string(*obj.State)
	}

	return result
}

func JobExecutionDetailsToMap(obj oci_fleet_apps_management.JobExecutionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsSensitive != nil {
		result["is_sensitive"] = bool(*obj.IsSensitive)
	}

	if obj.OutputDescription != nil {
		result["output_description"] = string(*obj.OutputDescription)
	}

	if obj.OutputName != nil {
		result["output_name"] = string(*obj.OutputName)
	}

	if obj.OutputType != nil {
		result["output_type"] = string(*obj.OutputType)
	}

	if obj.OutputValue != nil {
		result["output_value"] = string(*obj.OutputValue)
	}

	return result
}

func ProvisionSummaryToMap(obj oci_fleet_apps_management.ProvisionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConfigCatalogItemDisplayName != nil {
		result["config_catalog_item_display_name"] = string(*obj.ConfigCatalogItemDisplayName)
	}

	if obj.ConfigCatalogItemId != nil {
		result["config_catalog_item_id"] = string(*obj.ConfigCatalogItemId)
	}

	if obj.ConfigCatalogItemListingId != nil {
		result["config_catalog_item_listing_id"] = string(*obj.ConfigCatalogItemListingId)
	}

	if obj.ConfigCatalogItemListingVersion != nil {
		result["config_catalog_item_listing_version"] = string(*obj.ConfigCatalogItemListingVersion)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FleetId != nil {
		result["fleet_id"] = string(*obj.FleetId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.PackageCatalogItemDisplayName != nil {
		result["package_catalog_item_display_name"] = string(*obj.PackageCatalogItemDisplayName)
	}

	if obj.PackageCatalogItemId != nil {
		result["package_catalog_item_id"] = string(*obj.PackageCatalogItemId)
	}

	if obj.PackageCatalogItemListingId != nil {
		result["package_catalog_item_listing_id"] = string(*obj.PackageCatalogItemListingId)
	}

	if obj.PackageCatalogItemListingVersion != nil {
		result["package_catalog_item_listing_version"] = string(*obj.PackageCatalogItemListingVersion)
	}

	if obj.ProvisionDescription != nil {
		result["provision_description"] = string(*obj.ProvisionDescription)
	}

	if obj.StackId != nil {
		result["stack_id"] = string(*obj.StackId)
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

func (s *FleetAppsManagementProvisionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_fleet_apps_management.ChangeProvisionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ProvisionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.ChangeProvisionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getProvisionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
