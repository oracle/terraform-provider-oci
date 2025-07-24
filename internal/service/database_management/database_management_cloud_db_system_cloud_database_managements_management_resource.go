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

func DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagement,
		Read:     readDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagement,
		Update:   updateDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagement,
		Delete:   deleteDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_cloud_database_management": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				//ForceNew: true,
			},
			"metadata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				//ForceNew: true,
			},

			// Computed
		},
	}
}

func createDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResponse struct {
	enableResponse  *oci_database_management.EnableCloudDbSystemDatabaseManagementResponse
	disableResponse *oci_database_management.DisableCloudDbSystemDatabaseManagementResponse
}

type DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResponse
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResource-", DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResource(), s.D)
}

func (s *DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_cloud_database_management"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database_management.EnableCloudDbSystemDatabaseManagementRequest{}

		if cloudDbSystemId, ok := s.D.GetOkExists("cloud_db_system_id"); ok {
			tmp := cloudDbSystemId.(string)
			request.CloudDbSystemId = &tmp
		}

		if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
			tmp := isEnabled.(bool)
			request.IsEnabled = &tmp
		}

		if metadata, ok := s.D.GetOkExists("metadata"); ok {
			tmp := metadata.(string)
			request.Metadata = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

		response, err := s.Client.EnableCloudDbSystemDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getCloudDbSystemCloudDatabaseManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database_management.DisableCloudDbSystemDatabaseManagementRequest{}

	if cloudDbSystemId, ok := s.D.GetOkExists("cloud_db_system_id"); ok {
		tmp := cloudDbSystemId.(string)
		request.CloudDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableCloudDbSystemDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getCloudDbSystemCloudDatabaseManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResourceCrud) getCloudDbSystemCloudDatabaseManagementsManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := cloudDbSystemCloudDatabaseManagementsManagementWaitForWorkRequest(workId, "database_management",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func cloudDbSystemCloudDatabaseManagementsManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func cloudDbSystemCloudDatabaseManagementsManagementWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = cloudDbSystemCloudDatabaseManagementsManagementWorkRequestShouldRetryFunc(timeout)

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

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	// The work request will not have any resources if the DB system is RAC and
	// no connectors were added to any of its components except DB.
	if response.Status == oci_database_management.WorkRequestStatusFailed || response.Status == oci_database_management.WorkRequestStatusCanceled {
		return nil, getErrorFromDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_cloud_database_management"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database_management.EnableCloudDbSystemDatabaseManagementRequest{}

		if cloudDbSystemId, ok := s.D.GetOkExists("cloud_db_system_id"); ok {
			tmp := cloudDbSystemId.(string)
			request.CloudDbSystemId = &tmp
		}

		if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
			tmp := isEnabled.(bool)
			request.IsEnabled = &tmp
		}

		if metadata, ok := s.D.GetOkExists("metadata"); ok {
			tmp := metadata.(string)
			request.Metadata = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

		response, err := s.Client.EnableCloudDbSystemDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getCloudDbSystemCloudDatabaseManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database_management.DisableCloudDbSystemDatabaseManagementRequest{}

	if cloudDbSystemId, ok := s.D.GetOkExists("cloud_db_system_id"); ok {
		tmp := cloudDbSystemId.(string)
		request.CloudDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableCloudDbSystemDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getCloudDbSystemCloudDatabaseManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_cloud_database_management"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_database_management.DisableCloudDbSystemDatabaseManagementRequest{}

	if cloudDbSystemId, ok := s.D.GetOkExists("cloud_db_system_id"); ok {
		tmp := cloudDbSystemId.(string)
		request.CloudDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableCloudDbSystemDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getCloudDbSystemCloudDatabaseManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResourceCrud) SetData() error {
	return nil
}
