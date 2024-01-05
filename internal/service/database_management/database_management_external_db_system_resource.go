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

func DatabaseManagementExternalDbSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalDbSystem,
		Read:     readDatabaseManagementExternalDbSystem,
		Update:   updateDatabaseManagementExternalDbSystem,
		Delete:   deleteDatabaseManagementExternalDbSystem,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_discovery_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"database_management_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"license_model": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stack_monitoring_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"is_enabled": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"metadata": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"discovery_agent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"home_directory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cluster": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
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

func createDatabaseManagementExternalDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementExternalDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementExternalDbSystemResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ExternalDbSystem
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.ExternalDbSystemLifecycleStateCreating),
	}
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.ExternalDbSystemLifecycleStateActive),
	}
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.ExternalDbSystemLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.ExternalDbSystemLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) Create() error {
	request := oci_database_management.CreateExternalDbSystemRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseManagementConfig, ok := s.D.GetOkExists("database_management_config"); ok {
		if tmpList := databaseManagementConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database_management_config", 0)
			tmp, err := s.mapToExternalDbSystemDatabaseManagementConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatabaseManagementConfig = &tmp
		}
	}

	if dbSystemDiscoveryId, ok := s.D.GetOkExists("db_system_discovery_id"); ok {
		tmp := dbSystemDiscoveryId.(string)
		request.DbSystemDiscoveryId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if stackMonitoringConfig, ok := s.D.GetOkExists("stack_monitoring_config"); ok {
		if tmpList := stackMonitoringConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "stack_monitoring_config", 0)
			tmp, err := s.mapToAssociatedServiceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.StackMonitoringConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.CreateExternalDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getExternalDbSystemFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) getExternalDbSystemFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	externalDbSystemId, err := externalDbSystemWaitForWorkRequest(workId, "dbsystem",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*externalDbSystemId)

	return s.Get()
}

func externalDbSystemWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func externalDbSystemWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = externalDbSystemWorkRequestShouldRetryFunc(timeout)

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
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_database_management.WorkRequestStatusFailed || response.Status == oci_database_management.WorkRequestStatusCanceled {
		return nil, getErrorFromDatabaseManagementExternalDbSystemWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementExternalDbSystemWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatabaseManagementExternalDbSystemResourceCrud) Get() error {
	request := oci_database_management.GetExternalDbSystemRequest{}

	tmp := s.D.Id()
	request.ExternalDbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetExternalDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalDbSystem
	return nil
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_management.UpdateExternalDbSystemRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.ExternalDbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalDbSystem
	return nil
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) Delete() error {
	request := oci_database_management.DeleteExternalDbSystemRequest{}

	tmp := s.D.Id()
	request.ExternalDbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DeleteExternalDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := externalDbSystemWaitForWorkRequest(workId, "dbsystem",
		oci_database_management.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseManagementConfig != nil {
		s.D.Set("database_management_config", []interface{}{ExternalDbSystemDatabaseManagementConfigDetailsToMap(s.Res.DatabaseManagementConfig)})
	} else {
		s.D.Set("database_management_config", nil)
	}

	if s.Res.DbSystemDiscoveryId != nil {
		s.D.Set("db_system_discovery_id", *s.Res.DbSystemDiscoveryId)
	}

	if s.Res.DiscoveryAgentId != nil {
		s.D.Set("discovery_agent_id", *s.Res.DiscoveryAgentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.HomeDirectory != nil {
		s.D.Set("home_directory", *s.Res.HomeDirectory)
	}

	if s.Res.IsCluster != nil {
		s.D.Set("is_cluster", *s.Res.IsCluster)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.StackMonitoringConfig != nil {
		s.D.Set("stack_monitoring_config", []interface{}{ExternalDbSystemStackMonitoringConfigDetailsToMap(s.Res.StackMonitoringConfig)})
	} else {
		s.D.Set("stack_monitoring_config", nil)
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

func (s *DatabaseManagementExternalDbSystemResourceCrud) mapToAssociatedServiceDetails(fieldKeyFormat string) (oci_database_management.AssociatedServiceDetails, error) {
	result := oci_database_management.AssociatedServiceDetails{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		tmp := metadata.(string)
		result.Metadata = &tmp
	}

	return result, nil
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) mapToExternalDbSystemDatabaseManagementConfigDetails(fieldKeyFormat string) (oci_database_management.ExternalDbSystemDatabaseManagementConfigDetails, error) {
	result := oci_database_management.ExternalDbSystemDatabaseManagementConfigDetails{}

	if licenseModel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "license_model")); ok {
		result.LicenseModel = oci_database_management.ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum(licenseModel.(string))
	}

	return result, nil
}

func ExternalDbSystemStackMonitoringConfigDetailsToMap(obj *oci_database_management.ExternalDbSystemStackMonitoringConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.Metadata != nil {
		result["metadata"] = string(*obj.Metadata)
	}

	return result
}

func ExternalDbSystemDatabaseManagementConfigDetailsToMap(obj *oci_database_management.ExternalDbSystemDatabaseManagementConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["license_model"] = string(obj.LicenseModel)

	return result
}

func ExternalDbSystemSummaryToMap(obj oci_database_management.ExternalDbSystemSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabaseManagementConfig != nil {
		result["database_management_config"] = []interface{}{ExternalDbSystemDatabaseManagementConfigDetailsToMap(obj.DatabaseManagementConfig)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.HomeDirectory != nil {
		result["home_directory"] = string(*obj.HomeDirectory)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DatabaseManagementExternalDbSystemResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database_management.ChangeExternalDbSystemCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ExternalDbSystemId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.ChangeExternalDbSystemCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getExternalDbSystemFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
