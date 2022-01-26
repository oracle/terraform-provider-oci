// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package appmgmt_control

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_appmgmt_control "github.com/oracle/oci-go-sdk/v56/appmgmtcontrol"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func AppmgmtControlMonitorPluginManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAppmgmtControlMonitorPluginManagement,
		Read:     readAppmgmtControlMonitorPluginManagement,
		Delete:   deleteAppmgmtControlMonitorPluginManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"monitored_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"monitored_instance_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"monitored_instance_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"monitored_instance_management_agent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAppmgmtControlMonitorPluginManagement(d *schema.ResourceData, m interface{}) error {
	sync := &AppmgmtControlMonitorPluginManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AppmgmtControlClient()

	return tfresource.CreateResource(d, sync)
}

func readAppmgmtControlMonitorPluginManagement(d *schema.ResourceData, m interface{}) error {
	sync := &AppmgmtControlMonitorPluginManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AppmgmtControlClient()

	return tfresource.ReadResource(sync)
}

func deleteAppmgmtControlMonitorPluginManagement(d *schema.ResourceData, m interface{}) error {
	//N/A - once deactivate endpoint will be implemented/public, use it here
	return nil
}

type AppmgmtControlMonitorPluginManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_appmgmt_control.AppmgmtControlClient
	Res                    *oci_appmgmt_control.MonitoredInstance
	DisableNotFoundRetries bool
}

func (s *AppmgmtControlMonitorPluginManagementResourceCrud) ID() string {
	return *s.Res.InstanceId
}

func (s *AppmgmtControlMonitorPluginManagementResourceCrud) Create() error {
	request := oci_appmgmt_control.ActivateMonitoringPluginRequest{}

	if monitoredInstanceId, ok := s.D.GetOkExists("monitored_instance_id"); ok {
		tmp := monitoredInstanceId.(string)
		request.MonitoredInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "appmgmt_control")

	response, err := s.Client.ActivateMonitoringPlugin(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMonitorPluginManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "appmgmt_control"), oci_appmgmt_control.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AppmgmtControlMonitorPluginManagementResourceCrud) getMonitorPluginManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_appmgmt_control.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	monitorPluginManagementId, err := monitorPluginManagementWaitForWorkRequest(workId, "monitoredInstance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*monitorPluginManagementId)

	return s.Get()
}

func monitorPluginManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "appmgmt_control", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_appmgmt_control.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func monitorPluginManagementWaitForWorkRequest(wId *string, entityType string, action oci_appmgmt_control.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_appmgmt_control.AppmgmtControlClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "appmgmt_control")
	retryPolicy.ShouldRetryOperation = monitorPluginManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_appmgmt_control.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_appmgmt_control.OperationStatusInProgress),
			string(oci_appmgmt_control.OperationStatusAccepted),
			string(oci_appmgmt_control.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_appmgmt_control.OperationStatusSucceeded),
			string(oci_appmgmt_control.OperationStatusFailed),
			string(oci_appmgmt_control.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_appmgmt_control.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_appmgmt_control.OperationStatusFailed || response.Status == oci_appmgmt_control.OperationStatusCanceled {
		return nil, getErrorFromAppmgmtControlMonitorPluginManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAppmgmtControlMonitorPluginManagementWorkRequest(client *oci_appmgmt_control.AppmgmtControlClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_appmgmt_control.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_appmgmt_control.ListWorkRequestErrorsRequest{
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

func (s *AppmgmtControlMonitorPluginManagementResourceCrud) Get() error {
	request := oci_appmgmt_control.GetMonitoredInstanceRequest{}

	tmp := s.D.Id()
	request.MonitoredInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "appmgmt_control")

	response, err := s.Client.GetMonitoredInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredInstance
	return nil
}

func (s *AppmgmtControlMonitorPluginManagementResourceCrud) SetData() error {
	if s.Res.InstanceId != nil {
		s.D.Set("monitored_instance_id", *s.Res.InstanceId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ManagementAgentId != nil {
		s.D.Set("monitored_instance_management_agent_id", *s.Res.ManagementAgentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("monitored_instance_display_name", *s.Res.DisplayName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
