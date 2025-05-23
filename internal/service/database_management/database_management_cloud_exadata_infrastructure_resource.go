// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementCloudExadataInfrastructureResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementCloudExadataInfrastructure,
		Read:     readDatabaseManagementCloudExadataInfrastructure,
		Update:   updateDatabaseManagementCloudExadataInfrastructure,
		Delete:   deleteDatabaseManagementCloudExadataInfrastructure,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vm_cluster_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"discovery_key": {
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
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storage_server_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"infrastructure_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"additional_details": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"database_compartments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deployment_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"internal_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rack_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_grid": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"additional_details": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"internal_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
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
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"vm_clusters": {
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
						"additional_details": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"deployment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"home_directory": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"internal_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license_model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
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
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createDatabaseManagementCloudExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementCloudExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementCloudExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementCloudExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.DisableNotFoundRetries = true

	return sync.Delete()
}

type DatabaseManagementCloudExadataInfrastructureResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.CloudExadataInfrastructure
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.LifecycleStatesCreating),
	}
}

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.LifecycleStatesActive),
	}
}

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.LifecycleStatesDeleted),
	}
}

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.LifecycleStatesDeleted),
	}
}

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) Create() error {
	request := oci_database_management.CreateCloudExadataInfrastructureRequest{}

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

	if discoveryKey, ok := s.D.GetOkExists("discovery_key"); ok {
		tmp := discoveryKey.(string)
		request.DiscoveryKey = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database_management.CreateCloudExadataInfrastructureDetailsLicenseModelEnum(licenseModel.(string))
	}

	if storageServerNames, ok := s.D.GetOkExists("storage_server_names"); ok {
		interfaces := storageServerNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("storage_server_names") {
			request.StorageServerNames = tmp
		}
	}

	if vmClusterIds, ok := s.D.GetOkExists("vm_cluster_ids"); ok {
		interfaces := vmClusterIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("vm_cluster_ids") {
			request.VmClusterIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.CreateCloudExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudExadataInfrastructure
	return nil
}

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) getCloudExadataInfrastructureFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	cloudExadataInfrastructureId, err := cloudExadataInfrastructureWaitForWorkRequest(workId, "cloudexadatainfrastructure",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	if cloudExadataInfrastructureId != nil {
		s.D.SetId(*cloudExadataInfrastructureId)
	}

	return nil
}

func cloudExadataInfrastructureWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "database_management", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_database_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func cloudExadataInfrastructureWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = cloudExadataInfrastructureWorkRequestShouldRetryFunc(timeout)

	response := oci_database_management.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_database_management.WorkRequestStatusInProgress),
			string(oci_database_management.WorkRequestStatusAccepted),
			string(oci_database_management.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_database_management.WorkRequestStatusSucceeded),
			string(oci_database_management.WorkRequestStatusFailed),
			string(oci_database_management.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_database_management.GetWorkRequestRequest{
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

	return identifier, nil
}

func getErrorFromDatabaseManagementCloudExadataInfrastructureWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_database_management.ListWorkRequestErrorsRequest{
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

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) Get() error {
	request := oci_database_management.GetCloudExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.CloudExadataInfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
	response, err := s.Client.GetCloudExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudExadataInfrastructure
	return nil
}

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_management.UpdateCloudExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.CloudExadataInfrastructureId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if discoveryKey, ok := s.D.GetOkExists("discovery_key"); ok {
		tmp := discoveryKey.(string)
		request.DiscoveryKey = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database_management.UpdateCloudExadataInfrastructureDetailsLicenseModelEnum(licenseModel.(string))
	}

	if storageServerNames, ok := s.D.GetOkExists("storage_server_names"); ok {
		interfaces := storageServerNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("storage_server_names") {
			request.StorageServerNames = tmp
		}
	}

	if vmClusterIds, ok := s.D.GetOkExists("vm_cluster_ids"); ok {
		interfaces := vmClusterIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("vm_cluster_ids") {
			request.VmClusterIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateCloudExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudExadataInfrastructure
	return nil
}

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) Delete() error {
	request := oci_database_management.DeleteCloudExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.CloudExadataInfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")
	response, err := s.Client.DeleteCloudExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := cloudExadataInfrastructureWaitForWorkRequest(workId, "cloudexadatainfrastructure",
		oci_database_management.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) SetData() error {
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_compartments", s.Res.DatabaseCompartments)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("deployment_type", s.Res.DeploymentType)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InternalId != nil {
		s.D.Set("internal_id", *s.Res.InternalId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("rack_size", s.Res.RackSize)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	if s.Res.StorageGrid != nil {
		s.D.Set("storage_grid", []interface{}{CloudExadataStorageGridSummaryToMap(s.Res.StorageGrid)})
	} else {
		s.D.Set("storage_grid", nil)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
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

	vmClusters := []interface{}{}
	for _, item := range s.Res.VmClusters {
		vmClusters = append(vmClusters, ExadataVmClusterSummaryToMap(item))
	}
	s.D.Set("vm_clusters", vmClusters)

	return nil
}

func CloudExadataInfrastructureSummaryToMap(obj oci_database_management.CloudExadataInfrastructureSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["additional_details"] = obj.AdditionalDetails

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.GridHomePath != nil {
		result["grid_home_path"] = string(*obj.GridHomePath)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["infrastructure_type"] = string(obj.InfrastructureType)

	if obj.InternalId != nil {
		result["internal_id"] = string(*obj.InternalId)
	}

	result["license_model"] = string(obj.LicenseModel)

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["rack_size"] = string(obj.RackSize)

	result["state"] = string(obj.LifecycleState)

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

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

func CloudExadataStorageGridSummaryToMap(obj *oci_database_management.CloudExadataStorageGridSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["additional_details"] = obj.AdditionalDetails

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InternalId != nil {
		result["internal_id"] = string(*obj.InternalId)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ServerCount != nil {
		result["server_count"] = int(*obj.ServerCount)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
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

func ExadataVmClusterSummaryToMap(obj oci_database_management.ExadataVmClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["additional_details"] = obj.AdditionalDetails

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["deployment_type"] = string(obj.DeploymentType)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.HomeDirectory != nil {
		result["home_directory"] = string(*obj.HomeDirectory)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InternalId != nil {
		result["internal_id"] = string(*obj.InternalId)
	}

	result["license_model"] = string(obj.LicenseModel)

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
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

func (s *DatabaseManagementCloudExadataInfrastructureResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database_management.ChangeCloudExadataInfrastructureCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CloudExadataInfrastructureId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.ChangeCloudExadataInfrastructureCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCloudExadataInfrastructureFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
