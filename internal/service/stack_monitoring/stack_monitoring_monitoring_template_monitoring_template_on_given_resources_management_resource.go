// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResource() *schema.Resource {

	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagement,
		Read:     readStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagement,
		Update:   updateStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagement,
		Delete:   deleteStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"monitoring_template_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_monitoring_template_on_given_resources": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional

			// Computed
		},
	}
}

func createStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.Res = &StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.Res = &StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.Res = &StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResponse struct {
	enableResponse  *oci_stack_monitoring.ApplyMonitoringTemplateResponse
	disableResponse *oci_stack_monitoring.UnapplyMonitoringTemplateResponse
}

type StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResponse
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResource-", StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResource(), s.D)
}

func (s *StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_monitoring_template_on_given_resources"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_stack_monitoring.ApplyMonitoringTemplateRequest{}

		if monitoringTemplateId, ok := s.D.GetOkExists("monitoring_template_id"); ok {
			tmp := monitoringTemplateId.(string)
			request.MonitoringTemplateId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

		response, err := s.Client.ApplyMonitoringTemplate(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_stack_monitoring.UnapplyMonitoringTemplateRequest{}

	if monitoringTemplateId, ok := s.D.GetOkExists("monitoring_template_id"); ok {
		tmp := monitoringTemplateId.(string)
		request.MonitoringTemplateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UnapplyMonitoringTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceCrud) getMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_stack_monitoring.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := monitoringTemplateMonitoringTemplateOnGivenResourcesManagementWaitForWorkRequest(workId, "stack_monitoring",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func monitoringTemplateMonitoringTemplateOnGivenResourcesManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func monitoringTemplateMonitoringTemplateOnGivenResourcesManagementWaitForWorkRequest(wId *string, entityType string, action oci_stack_monitoring.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_stack_monitoring.StackMonitoringClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "stack_monitoring")
	retryPolicy.ShouldRetryOperation = monitoringTemplateMonitoringTemplateOnGivenResourcesManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_stack_monitoring.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if response.Status == oci_stack_monitoring.OperationStatusFailed || response.Status == oci_stack_monitoring.OperationStatusCanceled {
		return nil, getErrorFromStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromStackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementWorkRequest(client *oci_stack_monitoring.StackMonitoringClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_stack_monitoring.ActionTypeEnum) error {
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

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_monitoring_template_on_given_resources"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_stack_monitoring.ApplyMonitoringTemplateRequest{}

		if monitoringTemplateId, ok := s.D.GetOkExists("monitoring_template_id"); ok {
			tmp := monitoringTemplateId.(string)
			request.MonitoringTemplateId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

		response, err := s.Client.ApplyMonitoringTemplate(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_stack_monitoring.UnapplyMonitoringTemplateRequest{}

	if monitoringTemplateId, ok := s.D.GetOkExists("monitoring_template_id"); ok {
		tmp := monitoringTemplateId.(string)
		request.MonitoringTemplateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UnapplyMonitoringTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_monitoring_template_on_given_resources"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_stack_monitoring.UnapplyMonitoringTemplateRequest{}

	if monitoringTemplateId, ok := s.D.GetOkExists("monitoring_template_id"); ok {
		tmp := monitoringTemplateId.(string)
		request.MonitoringTemplateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UnapplyMonitoringTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResourceCrud) SetData() error {
	return nil
}
