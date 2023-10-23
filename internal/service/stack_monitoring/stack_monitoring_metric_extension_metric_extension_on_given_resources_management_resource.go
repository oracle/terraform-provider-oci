// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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

func StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagement,
		Read:     readStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagement,
		Update:   updateStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagement,
		Delete:   deleteStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"metric_extension_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MinItems: 1,
				MaxItems: 1,
			},
			"enable_metric_extension_on_given_resources": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional

			// Computed
		},
	}
}

func createStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.Res = &StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.Res = &StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.Res = &StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResponse struct {
	enableResponse  *oci_stack_monitoring.EnableMetricExtensionResponse
	disableResponse *oci_stack_monitoring.DisableMetricExtensionResponse
}

type StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResponse
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResource-", StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResource(), s.D)
}

func (s *StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_metric_extension_on_given_resources"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_stack_monitoring.EnableMetricExtensionRequest{}

		if metricExtensionId, ok := s.D.GetOkExists("metric_extension_id"); ok {
			tmp := metricExtensionId.(string)
			request.MetricExtensionId = &tmp
		}

		if resourceIds, ok := s.D.GetOkExists("resource_ids"); ok {
			interfaces := resourceIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("resource_ids") {
				request.ResourceIds = tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

		response, err := s.Client.EnableMetricExtension(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getMetricExtensionMetricExtensionOnGivenResourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_stack_monitoring.DisableMetricExtensionRequest{}

	if metricExtensionId, ok := s.D.GetOkExists("metric_extension_id"); ok {
		tmp := metricExtensionId.(string)
		request.MetricExtensionId = &tmp
	}

	if resourceIds, ok := s.D.GetOkExists("resource_ids"); ok {
		interfaces := resourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resource_ids") {
			request.ResourceIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.DisableMetricExtension(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getMetricExtensionMetricExtensionOnGivenResourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResourceCrud) getMetricExtensionMetricExtensionOnGivenResourcesManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_stack_monitoring.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := metricExtensionMetricExtensionOnGivenResourcesManagementWaitForWorkRequest(workId, "metricextension",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func metricExtensionMetricExtensionOnGivenResourcesManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func metricExtensionMetricExtensionOnGivenResourcesManagementWaitForWorkRequest(wId *string, entityType string, action oci_stack_monitoring.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_stack_monitoring.StackMonitoringClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "stack_monitoring")
	retryPolicy.ShouldRetryOperation = metricExtensionMetricExtensionOnGivenResourcesManagementWorkRequestShouldRetryFunc(timeout)

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
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_stack_monitoring.OperationStatusFailed || response.Status == oci_stack_monitoring.OperationStatusCanceled {
		return nil, getErrorFromStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromStackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementWorkRequest(client *oci_stack_monitoring.StackMonitoringClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_stack_monitoring.ActionTypeEnum) error {
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

func (s *StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_metric_extension_on_given_resources"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_stack_monitoring.EnableMetricExtensionRequest{}

		if metricExtensionId, ok := s.D.GetOkExists("metric_extension_id"); ok {
			tmp := metricExtensionId.(string)
			request.MetricExtensionId = &tmp
		}

		if resourceIds, ok := s.D.GetOkExists("resource_ids"); ok {
			interfaces := resourceIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("resource_ids") {
				request.ResourceIds = tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

		response, err := s.Client.EnableMetricExtension(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getMetricExtensionMetricExtensionOnGivenResourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	request := oci_stack_monitoring.DisableMetricExtensionRequest{}

	if metricExtensionId, ok := s.D.GetOkExists("metric_extension_id"); ok {
		tmp := metricExtensionId.(string)
		request.MetricExtensionId = &tmp
	}

	if resourceIds, ok := s.D.GetOkExists("resource_ids"); ok {
		interfaces := resourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resource_ids") {
			request.ResourceIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.DisableMetricExtension(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getMetricExtensionMetricExtensionOnGivenResourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_metric_extension_on_given_resources"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_stack_monitoring.DisableMetricExtensionRequest{}

	if metricExtensionId, ok := s.D.GetOkExists("metric_extension_id"); ok {
		tmp := metricExtensionId.(string)
		request.MetricExtensionId = &tmp
	}

	if resourceIds, ok := s.D.GetOkExists("resource_ids"); ok {
		interfaces := resourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resource_ids") {
			request.ResourceIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.DisableMetricExtension(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getMetricExtensionMetricExtensionOnGivenResourcesManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResourceCrud) SetData() error {
	return nil
}
