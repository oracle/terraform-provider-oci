// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalCluster,
		Read:     readDatabaseManagementExternalCluster,
		Update:   updateDatabaseManagementExternalCluster,
		Delete:   deleteDatabaseManagementExternalCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"external_cluster_id": {
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
			"external_connector_id": {
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
			"additional_details": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"component_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"grid_home": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_flex_cluster": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"network_number": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ocr_file_location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scan_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"network_number": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scan_port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_protocol": {
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
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vip_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_number": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createDatabaseManagementExternalCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementExternalCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalCluster(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementExternalClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ExternalCluster
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementExternalClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.ExternalClusterLifecycleStateCreating),
	}
}

func (s *DatabaseManagementExternalClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.ExternalClusterLifecycleStateNotConnected),
		string(oci_database_management.ExternalClusterLifecycleStateActive),
	}
}

func (s *DatabaseManagementExternalClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.ExternalClusterLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementExternalClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.ExternalClusterLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementExternalClusterResourceCrud) Create() error {
	request := oci_database_management.UpdateExternalClusterRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if externalClusterId, ok := s.D.GetOkExists("external_cluster_id"); ok {
		tmp := externalClusterId.(string)
		request.ExternalClusterId = &tmp
	}

	if externalConnectorId, ok := s.D.GetOkExists("external_connector_id"); ok {
		tmp := externalConnectorId.(string)
		request.ExternalConnectorId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_database_management.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_database_management.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "cluster") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getExternalClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseManagementExternalClusterResourceCrud) getExternalClusterFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	externalClusterId, err := externalClusterWaitForWorkRequest(workId, "cluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*externalClusterId)

	return s.Get()
}

func externalClusterWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func externalClusterWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = externalClusterWorkRequestShouldRetryFunc(timeout)

	response := oci_database_management.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
			identifier = res.Identifier
			break
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_database_management.WorkRequestStatusFailed || response.Status == oci_database_management.WorkRequestStatusCanceled {
		return nil, getErrorFromDatabaseManagementExternalClusterWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementExternalClusterWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatabaseManagementExternalClusterResourceCrud) Get() error {
	request := oci_database_management.GetExternalClusterRequest{}

	tmp := s.D.Id()
	request.ExternalClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetExternalCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalCluster
	return nil
}

func (s *DatabaseManagementExternalClusterResourceCrud) Update() error {
	request := oci_database_management.UpdateExternalClusterRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.ExternalClusterId = &tmp

	if externalConnectorId, ok := s.D.GetOkExists("external_connector_id"); ok {
		tmp := externalConnectorId.(string)
		request.ExternalConnectorId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getExternalClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseManagementExternalClusterResourceCrud) SetData() error {
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalConnectorId != nil {
		s.D.Set("external_connector_id", *s.Res.ExternalConnectorId)
	}

	if s.Res.ExternalDbSystemId != nil {
		s.D.Set("external_db_system_id", *s.Res.ExternalDbSystemId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GridHome != nil {
		s.D.Set("grid_home", *s.Res.GridHome)
	}

	if s.Res.IsFlexCluster != nil {
		s.D.Set("is_flex_cluster", *s.Res.IsFlexCluster)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	networkConfigurations := []interface{}{}
	for _, item := range s.Res.NetworkConfigurations {
		networkConfigurations = append(networkConfigurations, ExternalClusterNetworkConfigurationToMap(item))
	}
	s.D.Set("network_configurations", networkConfigurations)

	if s.Res.OcrFileLocation != nil {
		s.D.Set("ocr_file_location", *s.Res.OcrFileLocation)
	}

	scanConfigurations := []interface{}{}
	for _, item := range s.Res.ScanConfigurations {
		scanConfigurations = append(scanConfigurations, ExternalClusterScanListenerConfigurationToMap(item))
	}
	s.D.Set("scan_configurations", scanConfigurations)

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

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	vipConfigurations := []interface{}{}
	for _, item := range s.Res.VipConfigurations {
		vipConfigurations = append(vipConfigurations, ExternalClusterVipConfigurationToMap(item))
	}
	s.D.Set("vip_configurations", vipConfigurations)

	return nil
}

func ExternalClusterNetworkConfigurationToMap(obj oci_database_management.ExternalClusterNetworkConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NetworkNumber != nil {
		result["network_number"] = int(*obj.NetworkNumber)
	}

	result["network_type"] = string(obj.NetworkType)

	if obj.Subnet != nil {
		result["subnet"] = string(*obj.Subnet)
	}

	return result
}

func ExternalClusterScanListenerConfigurationToMap(obj oci_database_management.ExternalClusterScanListenerConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NetworkNumber != nil {
		result["network_number"] = int(*obj.NetworkNumber)
	}

	if obj.ScanName != nil {
		result["scan_name"] = string(*obj.ScanName)
	}

	if obj.ScanPort != nil {
		result["scan_port"] = int(*obj.ScanPort)
	}

	result["scan_protocol"] = string(obj.ScanProtocol)

	return result
}

func ExternalClusterSummaryToMap(obj oci_database_management.ExternalClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComponentName != nil {
		result["component_name"] = string(*obj.ComponentName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalConnectorId != nil {
		result["external_connector_id"] = string(*obj.ExternalConnectorId)
	}

	if obj.ExternalDbSystemId != nil {
		result["external_db_system_id"] = string(*obj.ExternalDbSystemId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func ExternalClusterVipConfigurationToMap(obj oci_database_management.ExternalClusterVipConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Address != nil {
		result["address"] = string(*obj.Address)
	}

	if obj.NetworkNumber != nil {
		result["network_number"] = int(*obj.NetworkNumber)
	}

	if obj.NodeName != nil {
		result["node_name"] = string(*obj.NodeName)
	}

	return result
}
