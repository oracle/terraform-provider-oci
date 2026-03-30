// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalDbSystemDatabaseManagementsManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseManagementExternalDbSystemDatabaseManagementsManagement,
		ReadContext:   readDatabaseManagementExternalDbSystemDatabaseManagementsManagement,
		UpdateContext: updateDatabaseManagementExternalDbSystemDatabaseManagementsManagement,
		DeleteContext: deleteDatabaseManagementExternalDbSystemDatabaseManagementsManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"external_db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_database_management": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createDatabaseManagementExternalDbSystemDatabaseManagementsManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseManagementExternalDbSystemDatabaseManagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalDbSystemDatabaseManagementsManagementResponse{}

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseManagementExternalDbSystemDatabaseManagementsManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func updateDatabaseManagementExternalDbSystemDatabaseManagementsManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseManagementExternalDbSystemDatabaseManagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalDbSystemDatabaseManagementsManagementResponse{}

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseManagementExternalDbSystemDatabaseManagementsManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseManagementExternalDbSystemDatabaseManagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.Res = &DatabaseManagementExternalDbSystemDatabaseManagementsManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseManagementExternalDbSystemDatabaseManagementsManagementResponse struct {
	enableResponse  *oci_database_management.EnableExternalDbSystemDatabaseManagementResponse
	disableResponse *oci_database_management.DisableExternalDbSystemDatabaseManagementResponse
}

type DatabaseManagementExternalDbSystemDatabaseManagementsManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *DatabaseManagementExternalDbSystemDatabaseManagementsManagementResponse
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalDbSystemDatabaseManagementsManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseManagementExternalDbSystemDatabaseManagementsManagementResource-", DatabaseManagementExternalDbSystemDatabaseManagementsManagementResource(), s.D)
}

func (s *DatabaseManagementExternalDbSystemDatabaseManagementsManagementResourceCrud) CreateWithContext(ctx context.Context) error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_database_management"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database_management.EnableExternalDbSystemDatabaseManagementRequest{}

		if externalDbSystemId, ok := s.D.GetOkExists("external_db_system_id"); ok {
			tmp := externalDbSystemId.(string)
			request.ExternalDbSystemId = &tmp
		}

		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			request.LicenseModel = oci_database_management.EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum(licenseModel.(string))
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

		response, err := s.Client.EnableExternalDbSystemDatabaseManagement(ctx, request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getExternalDbSystemDatabaseManagementsManagementFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database_management.DisableExternalDbSystemDatabaseManagementRequest{}

	if externalDbSystemId, ok := s.D.GetOkExists("external_db_system_id"); ok {
		tmp := externalDbSystemId.(string)
		request.ExternalDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableExternalDbSystemDatabaseManagement(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalDbSystemDatabaseManagementsManagementFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalDbSystemDatabaseManagementsManagementResourceCrud) getExternalDbSystemDatabaseManagementsManagementFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := externalDbSystemDatabaseManagementsManagementWaitForWorkRequest(ctx, workId, "database_management",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func externalDbSystemDatabaseManagementsManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func externalDbSystemDatabaseManagementsManagementWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = externalDbSystemDatabaseManagementsManagementWorkRequestShouldRetryFunc(timeout)

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
			response, err = client.GetWorkRequest(ctx,
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
		return nil, getErrorFromDatabaseManagementExternalDbSystemDatabaseManagementsManagementWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementExternalDbSystemDatabaseManagementsManagementWorkRequest(ctx context.Context, client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
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

func (s *DatabaseManagementExternalDbSystemDatabaseManagementsManagementResourceCrud) UpdateWithContext(ctx context.Context) error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_database_management"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database_management.EnableExternalDbSystemDatabaseManagementRequest{}

		if externalDbSystemId, ok := s.D.GetOkExists("external_db_system_id"); ok {
			tmp := externalDbSystemId.(string)
			request.ExternalDbSystemId = &tmp
		}

		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			request.LicenseModel = oci_database_management.EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum(licenseModel.(string))
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

		response, err := s.Client.EnableExternalDbSystemDatabaseManagement(ctx, request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getExternalDbSystemDatabaseManagementsManagementFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database_management.DisableExternalDbSystemDatabaseManagementRequest{}

	if externalDbSystemId, ok := s.D.GetOkExists("external_db_system_id"); ok {
		tmp := externalDbSystemId.(string)
		request.ExternalDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableExternalDbSystemDatabaseManagement(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalDbSystemDatabaseManagementsManagementFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalDbSystemDatabaseManagementsManagementResourceCrud) DeleteWithContext(ctx context.Context) error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_database_management"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_database_management.DisableExternalDbSystemDatabaseManagementRequest{}

	if externalDbSystemId, ok := s.D.GetOkExists("external_db_system_id"); ok {
		tmp := externalDbSystemId.(string)
		request.ExternalDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.DisableExternalDbSystemDatabaseManagement(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getExternalDbSystemDatabaseManagementsManagementFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseManagementExternalDbSystemDatabaseManagementsManagementResourceCrud) SetData() error {
	return nil
}
