// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagedInstanceRefreshSoftwareManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubManagedInstanceRefreshSoftwareManagement,
		Read:     readOsManagementHubManagedInstanceRefreshSoftwareManagement,
		Delete:   deleteOsManagementHubManagedInstanceRefreshSoftwareManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createOsManagementHubManagedInstanceRefreshSoftwareManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceRefreshSoftwareManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubManagedInstanceRefreshSoftwareManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubManagedInstanceRefreshSoftwareManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubManagedInstanceRefreshSoftwareManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ManagedInstanceClient
	Res                    *oci_os_management_hub.GetManagedInstanceResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_os_management_hub.WorkRequestClient
}

func (s *OsManagementHubManagedInstanceRefreshSoftwareManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubManagedInstanceRefreshSoftwareManagementResourceCrud) Get() error {
	request := oci_os_management_hub.GetManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubManagedInstanceRefreshSoftwareManagementResourceCrud) Create() error {
	request := oci_os_management_hub.RefreshSoftwareOnManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.RefreshSoftwareOnManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getManagedInstanceRefreshSoftwareManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub"), oci_os_management_hub.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OsManagementHubManagedInstanceRefreshSoftwareManagementResourceCrud) getManagedInstanceRefreshSoftwareManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_os_management_hub.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	managedInstanceRefreshSoftwareManagementId, err := managedInstanceRefreshSoftwareManagementWaitForWorkRequest(workId, "instance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*managedInstanceRefreshSoftwareManagementId)

	return s.Get()
}

func managedInstanceRefreshSoftwareManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "os_management_hub", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_os_management_hub.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func managedInstanceRefreshSoftwareManagementWaitForWorkRequest(wId *string, entityType string, action oci_os_management_hub.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_os_management_hub.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "os_management_hub")
	retryPolicy.ShouldRetryOperation = managedInstanceRefreshSoftwareManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_os_management_hub.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_os_management_hub.OperationStatusInProgress),
			string(oci_os_management_hub.OperationStatusAccepted),
			string(oci_os_management_hub.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_os_management_hub.OperationStatusSucceeded),
			string(oci_os_management_hub.OperationStatusFailed),
			string(oci_os_management_hub.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_os_management_hub.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(string(res.EntityType)), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_os_management_hub.OperationStatusFailed || response.Status == oci_os_management_hub.OperationStatusCanceled {
		return nil, getErrorFromOsManagementHubManagedInstanceRefreshSoftwareManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOsManagementHubManagedInstanceRefreshSoftwareManagementWorkRequest(client *oci_os_management_hub.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_os_management_hub.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_os_management_hub.ListWorkRequestErrorsRequest{
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

func (s *OsManagementHubManagedInstanceRefreshSoftwareManagementResourceCrud) SetData() error {
	return nil
}
