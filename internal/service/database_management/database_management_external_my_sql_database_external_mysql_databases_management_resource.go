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

func DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagement,
		Read:     readDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagement,
		Update:   updateDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagement,
		Delete:   deleteDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"external_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_external_mysql_database": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"connector_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResponse struct {
	enableResponse  *oci_database_management.EnableExternalMySqlDatabaseManagementResponse
	disableResponse *oci_database_management.DisableExternalMySqlDatabaseManagementResponse
}

type DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResponse
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResource-", DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResource(), s.D)
}

func (s *DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_external_mysql_database"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database_management.EnableExternalMySqlDatabaseManagementRequest{}

		if connectorId, ok := s.D.GetOkExists("connector_id"); ok {
			tmp := connectorId.(string)
			request.ConnectorId = &tmp
		}

		if externalMySqlDatabaseId, ok := s.D.GetOkExists("external_my_sql_database_id"); ok {
			tmp := externalMySqlDatabaseId.(string)
			request.ExternalMySqlDatabaseId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "connector")

		response, err := s.Client.EnableExternalMySqlDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getExternalMySqlDatabaseExternalMysqlDatabasesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "connector"), oci_database_management.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database_management.DisableExternalMySqlDatabaseManagementRequest{}

	if externalMySqlDatabaseId, ok := s.D.GetOkExists("external_my_sql_database_id"); ok {
		tmp := externalMySqlDatabaseId.(string)
		request.ExternalMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "connector")

	response, err := s.Client.DisableExternalMySqlDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalMySqlDatabaseExternalMysqlDatabasesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "connector"), oci_database_management.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResourceCrud) getExternalMySqlDatabaseExternalMysqlDatabasesManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := externalMySqlDatabaseExternalMysqlDatabasesManagementWaitForWorkRequest(workId, "connector",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func externalMySqlDatabaseExternalMysqlDatabasesManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "connector", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_database_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func externalMySqlDatabaseExternalMysqlDatabasesManagementWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "connector")
	retryPolicy.ShouldRetryOperation = externalMySqlDatabaseExternalMysqlDatabasesManagementWorkRequestShouldRetryFunc(timeout)

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
	if identifier == nil || response.Status == oci_database_management.WorkRequestStatusFailed || response.Status == oci_database_management.WorkRequestStatusCanceled {
		return nil, getErrorFromDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_external_mysql_database"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database_management.EnableExternalMySqlDatabaseManagementRequest{}

		if connectorId, ok := s.D.GetOkExists("connector_id"); ok {
			tmp := connectorId.(string)
			request.ConnectorId = &tmp
		}

		if externalMySqlDatabaseId, ok := s.D.GetOkExists("external_my_sql_database_id"); ok {
			tmp := externalMySqlDatabaseId.(string)
			request.ExternalMySqlDatabaseId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "connector")

		response, err := s.Client.EnableExternalMySqlDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getExternalMySqlDatabaseExternalMysqlDatabasesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "connector"), oci_database_management.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database_management.DisableExternalMySqlDatabaseManagementRequest{}

	if externalMySqlDatabaseId, ok := s.D.GetOkExists("external_my_sql_database_id"); ok {
		tmp := externalMySqlDatabaseId.(string)
		request.ExternalMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "connector")

	response, err := s.Client.DisableExternalMySqlDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalMySqlDatabaseExternalMysqlDatabasesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "connector"), oci_database_management.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_external_mysql_database"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_database_management.DisableExternalMySqlDatabaseManagementRequest{}

	if externalMySqlDatabaseId, ok := s.D.GetOkExists("external_my_sql_database_id"); ok {
		tmp := externalMySqlDatabaseId.(string)
		request.ExternalMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "connector")

	response, err := s.Client.DisableExternalMySqlDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalMySqlDatabaseExternalMysqlDatabasesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "connector"), oci_database_management.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResourceCrud) SetData() error {
	return nil
}
