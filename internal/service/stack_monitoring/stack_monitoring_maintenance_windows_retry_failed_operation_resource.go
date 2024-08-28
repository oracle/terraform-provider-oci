// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMaintenanceWindowsRetryFailedOperationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMaintenanceWindowsRetryFailedOperation,
		Read:     readStackMonitoringMaintenanceWindowsRetryFailedOperation,
		Delete:   deleteStackMonitoringMaintenanceWindowsRetryFailedOperation,
		Schema: map[string]*schema.Schema{
			// Required
			"maintenance_window_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createStackMonitoringMaintenanceWindowsRetryFailedOperation(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMaintenanceWindowsRetryFailedOperationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMaintenanceWindowsRetryFailedOperation(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteStackMonitoringMaintenanceWindowsRetryFailedOperation(d *schema.ResourceData, m interface{}) error {
	return nil
}

type StackMonitoringMaintenanceWindowsRetryFailedOperationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.RetryFailedMaintenanceWindowOperationResponse
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMaintenanceWindowsRetryFailedOperationResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("StackMonitoringMaintenanceWindowsRetryFailedOperationResource-", StackMonitoringMaintenanceWindowsRetryFailedOperationResource(), s.D)
}

func (s *StackMonitoringMaintenanceWindowsRetryFailedOperationResourceCrud) Create() error {
	request := oci_stack_monitoring.RetryFailedMaintenanceWindowOperationRequest{}

	if maintenanceWindowId, ok := s.D.GetOkExists("maintenance_window_id"); ok {
		tmp := maintenanceWindowId.(string)
		request.MaintenanceWindowId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.RetryFailedMaintenanceWindowOperation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	expectedActionTypes := []oci_stack_monitoring.ActionTypeEnum{
		oci_stack_monitoring.ActionTypeCreated,
		oci_stack_monitoring.ActionTypeDeleted,
		oci_stack_monitoring.ActionTypeUpdated,
		oci_stack_monitoring.ActionTypeFailed,
	}

	return s.getMaintenanceWindowsRetryFailedOperationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), expectedActionTypes, s.D.Timeout(schema.TimeoutCreate))
}

func (s *StackMonitoringMaintenanceWindowsRetryFailedOperationResourceCrud) getMaintenanceWindowsRetryFailedOperationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum []oci_stack_monitoring.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	maintenanceWindowsRetryFailedOperationId, err := maintenanceWindowsRetryFailedOperationWaitForWorkRequest(workId, "maintenancewindow",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*maintenanceWindowsRetryFailedOperationId)

	return nil
}

func maintenanceWindowsRetryFailedOperationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "stack_monitoring", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_stack_monitoring.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func maintenanceWindowsRetryFailedOperationWaitForWorkRequest(wId *string, entityType string, actions []oci_stack_monitoring.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_stack_monitoring.StackMonitoringClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "stack_monitoring")
	retryPolicy.ShouldRetryOperation = maintenanceWindowsRetryFailedOperationWorkRequestShouldRetryFunc(timeout)

	response := oci_stack_monitoring.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_stack_monitoring.OperationStatusInProgress),
			string(oci_stack_monitoring.OperationStatusAccepted),
			string(oci_stack_monitoring.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_stack_monitoring.OperationStatusSucceeded),
			string(oci_stack_monitoring.OperationStatusFailed),
			string(oci_stack_monitoring.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_stack_monitoring.GetWorkRequestRequest{
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
			for _, action := range actions {
				if res.ActionType == action {
					identifier = res.Identifier
					break
				}
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or got cancelled
	if identifier == nil || response.Status == oci_stack_monitoring.OperationStatusCanceled {
		return nil, getErrorFromStackMonitoringMaintenanceWindowsRetryFailedOperationWorkRequest(client, wId, retryPolicy, entityType, actions)
	}

	return identifier, nil
}

func getErrorFromStackMonitoringMaintenanceWindowsRetryFailedOperationWorkRequest(client *oci_stack_monitoring.StackMonitoringClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, actions []oci_stack_monitoring.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_stack_monitoring.ListWorkRequestErrorsRequest{
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

	var actionsStrings []string

	for _, action := range actions {
		actionsStrings = append(actionsStrings, string(action))
	}

	actionsString := strings.Join(actionsStrings, "|")
	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, actionsString, errorMessage)

	return workRequestErr
}

func (s *StackMonitoringMaintenanceWindowsRetryFailedOperationResourceCrud) SetData() error {
	return nil
}
