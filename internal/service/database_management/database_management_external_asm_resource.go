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

func DatabaseManagementExternalAsmResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalAsm,
		Read:     readDatabaseManagementExternalAsm,
		Update:   updateDatabaseManagementExternalAsm,
		Delete:   deleteDatabaseManagementExternalAsm,
		Schema: map[string]*schema.Schema{
			// Required
			"external_asm_id": {
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
			"is_cluster": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_flex_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serviced_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_sub_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_unique_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"disk_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_managed": {
							Type:     schema.TypeBool,
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
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseManagementExternalAsm(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalAsmResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalAsm(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalAsmResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementExternalAsm(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalAsmResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalAsm(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementExternalAsmResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ExternalAsm
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalAsmResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementExternalAsmResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.ExternalAsmLifecycleStateCreating),
	}
}

func (s *DatabaseManagementExternalAsmResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.ExternalAsmLifecycleStateNotConnected),
		string(oci_database_management.ExternalAsmLifecycleStateActive),
	}
}

func (s *DatabaseManagementExternalAsmResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.ExternalAsmLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementExternalAsmResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.ExternalAsmLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementExternalAsmResourceCrud) Create() error {
	request := oci_database_management.UpdateExternalAsmRequest{}

	if externalAsmId, ok := s.D.GetOkExists("external_asm_id"); ok {
		tmp := externalAsmId.(string)
		request.ExternalAsmId = &tmp
	}

	if externalConnectorId, ok := s.D.GetOkExists("external_connector_id"); ok {
		tmp := externalConnectorId.(string)
		request.ExternalConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalAsm(context.Background(), request)
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "asm") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getExternalAsmFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseManagementExternalAsmResourceCrud) getExternalAsmFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	externalAsmId, err := externalAsmWaitForWorkRequest(workId, "asm",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*externalAsmId)

	return s.Get()
}

func externalAsmWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func externalAsmWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = externalAsmWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDatabaseManagementExternalAsmWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementExternalAsmWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatabaseManagementExternalAsmResourceCrud) Get() error {
	request := oci_database_management.GetExternalAsmRequest{}

	tmp := s.D.Id()
	request.ExternalAsmId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetExternalAsm(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalAsm
	return nil
}

func (s *DatabaseManagementExternalAsmResourceCrud) Update() error {
	request := oci_database_management.UpdateExternalAsmRequest{}

	tmp := s.D.Id()
	request.ExternalAsmId = &tmp

	if externalConnectorId, ok := s.D.GetOkExists("external_connector_id"); ok {
		tmp := externalConnectorId.(string)
		request.ExternalConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalAsm(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getExternalAsmFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseManagementExternalAsmResourceCrud) SetData() error {
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
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

	if s.Res.GridHome != nil {
		s.D.Set("grid_home", *s.Res.GridHome)
	}

	if s.Res.IsCluster != nil {
		s.D.Set("is_cluster", *s.Res.IsCluster)
	}

	if s.Res.IsFlexEnabled != nil {
		s.D.Set("is_flex_enabled", *s.Res.IsFlexEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	servicedDatabases := []interface{}{}
	for _, item := range s.Res.ServicedDatabases {
		servicedDatabases = append(servicedDatabases, ExternalAsmServicedDatabaseToMap(item))
	}
	s.D.Set("serviced_databases", servicedDatabases)

	s.D.Set("state", s.Res.LifecycleState)

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

func ExternalAsmServicedDatabaseToMap(obj oci_database_management.ExternalAsmServicedDatabase) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["database_sub_type"] = string(obj.DatabaseSubType)

	result["database_type"] = string(obj.DatabaseType)

	if obj.DbUniqueName != nil {
		result["db_unique_name"] = string(*obj.DbUniqueName)
	}

	result["disk_groups"] = obj.DiskGroups

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsManaged != nil {
		result["is_managed"] = bool(*obj.IsManaged)
	}

	return result
}

func ExternalAsmSummaryToMap(obj oci_database_management.ExternalAsmSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComponentName != nil {
		result["component_name"] = string(*obj.ComponentName)
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
