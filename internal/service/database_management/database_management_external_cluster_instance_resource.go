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

func DatabaseManagementExternalClusterInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalClusterInstance,
		Read:     readDatabaseManagementExternalClusterInstance,
		Update:   updateDatabaseManagementExternalClusterInstance,
		Delete:   deleteDatabaseManagementExternalClusterInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"external_cluster_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"external_connector_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"adr_home_directory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"component_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crs_base_directory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_db_node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_role": {
				Type:     schema.TypeString,
				Computed: true,
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

func createDatabaseManagementExternalClusterInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalClusterInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalClusterInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalClusterInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementExternalClusterInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalClusterInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalClusterInstance(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementExternalClusterInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ExternalClusterInstance
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalClusterInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementExternalClusterInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.ExternalClusterInstanceLifecycleStateCreating),
	}
}

func (s *DatabaseManagementExternalClusterInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.ExternalClusterInstanceLifecycleStateNotConnected),
		string(oci_database_management.ExternalClusterInstanceLifecycleStateActive),
	}
}

func (s *DatabaseManagementExternalClusterInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.ExternalClusterInstanceLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementExternalClusterInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.ExternalClusterInstanceLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementExternalClusterInstanceResourceCrud) Create() error {
	request := oci_database_management.UpdateExternalClusterInstanceRequest{}

	if externalClusterInstanceId, ok := s.D.GetOkExists("external_cluster_instance_id"); ok {
		tmp := externalClusterInstanceId.(string)
		request.ExternalClusterInstanceId = &tmp
	}

	if externalConnectorId, ok := s.D.GetOkExists("external_connector_id"); ok {
		tmp := externalConnectorId.(string)
		request.ExternalConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalClusterInstance(context.Background(), request)
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
	return s.getExternalClusterInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseManagementExternalClusterInstanceResourceCrud) getExternalClusterInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	externalClusterInstanceId, err := externalClusterInstanceWaitForWorkRequest(workId, "cluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*externalClusterInstanceId)

	return s.Get()
}

func externalClusterInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func externalClusterInstanceWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = externalClusterInstanceWorkRequestShouldRetryFunc(timeout)

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
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_database_management.WorkRequestStatusFailed || response.Status == oci_database_management.WorkRequestStatusCanceled {
		return nil, getErrorFromDatabaseManagementExternalClusterInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementExternalClusterInstanceWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatabaseManagementExternalClusterInstanceResourceCrud) Get() error {
	request := oci_database_management.GetExternalClusterInstanceRequest{}

	tmp := s.D.Id()
	request.ExternalClusterInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetExternalClusterInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalClusterInstance
	return nil
}

func (s *DatabaseManagementExternalClusterInstanceResourceCrud) Update() error {
	request := oci_database_management.UpdateExternalClusterInstanceRequest{}

	tmp := s.D.Id()
	request.ExternalClusterInstanceId = &tmp

	if externalConnectorId, ok := s.D.GetOkExists("external_connector_id"); ok {
		tmp := externalConnectorId.(string)
		request.ExternalConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalClusterInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getExternalClusterInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseManagementExternalClusterInstanceResourceCrud) SetData() error {
	if s.Res.AdrHomeDirectory != nil {
		s.D.Set("adr_home_directory", *s.Res.AdrHomeDirectory)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.CrsBaseDirectory != nil {
		s.D.Set("crs_base_directory", *s.Res.CrsBaseDirectory)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalClusterId != nil {
		s.D.Set("external_cluster_id", *s.Res.ExternalClusterId)
	}

	if s.Res.ExternalConnectorId != nil {
		s.D.Set("external_connector_id", *s.Res.ExternalConnectorId)
	}

	if s.Res.ExternalDbNodeId != nil {
		s.D.Set("external_db_node_id", *s.Res.ExternalDbNodeId)
	}

	if s.Res.ExternalDbSystemId != nil {
		s.D.Set("external_db_system_id", *s.Res.ExternalDbSystemId)
	}

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("node_role", s.Res.NodeRole)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ExternalClusterInstanceSummaryToMap(obj oci_database_management.ExternalClusterInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdrHomeDirectory != nil {
		result["adr_home_directory"] = string(*obj.AdrHomeDirectory)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComponentName != nil {
		result["component_name"] = string(*obj.ComponentName)
	}

	if obj.CrsBaseDirectory != nil {
		result["crs_base_directory"] = string(*obj.CrsBaseDirectory)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalClusterId != nil {
		result["external_cluster_id"] = string(*obj.ExternalClusterId)
	}

	if obj.ExternalConnectorId != nil {
		result["external_connector_id"] = string(*obj.ExternalConnectorId)
	}

	if obj.ExternalDbNodeId != nil {
		result["external_db_node_id"] = string(*obj.ExternalDbNodeId)
	}

	if obj.ExternalDbSystemId != nil {
		result["external_db_system_id"] = string(*obj.ExternalDbSystemId)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["node_role"] = string(obj.NodeRole)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
