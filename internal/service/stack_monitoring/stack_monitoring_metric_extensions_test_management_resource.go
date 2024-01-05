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

func StackMonitoringMetricExtensionsTestManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMetricExtensionsTestManagement,
		Read:     readStackMonitoringMetricExtensionsTestManagement,
		Delete:   deleteStackMonitoringMetricExtensionsTestManagement,
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

			// Optional

			// Computed
			"test_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"test_run_metric_suffix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"test_run_namespace_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"test_run_resource_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createStackMonitoringMetricExtensionsTestManagement(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMetricExtensionsTestManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMetricExtensionsTestManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteStackMonitoringMetricExtensionsTestManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type StackMonitoringMetricExtensionsTestManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.TestMetricExtensionResponse
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMetricExtensionsTestManagementResourceCrud) ID() string {

	var metricExtensionIdStr, resourceIdStr string
	if metricExtensionId, ok := s.D.GetOkExists("metric_extension_id"); ok {
		metricExtensionIdStr = metricExtensionId.(string)
	}

	if resourceIds, ok := s.D.GetOkExists("resource_ids"); ok {
		interfaces := resourceIds.([]interface{})
		resourceIdStr = interfaces[0].(string)

	}
	return metricExtensionIdStr + "/" + resourceIdStr + "/" + time.Millisecond.String()

}

func (s *StackMonitoringMetricExtensionsTestManagementResourceCrud) Create() error {
	request := oci_stack_monitoring.TestMetricExtensionRequest{}

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

	response, err := s.Client.TestMetricExtension(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response

	return s.getMetricExtensionsTestManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *StackMonitoringMetricExtensionsTestManagementResourceCrud) getMetricExtensionsTestManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_stack_monitoring.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	metricExtensionsTestManagementId, err := metricExtensionsTestManagementWaitForWorkRequest(workId, "metricextension",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*metricExtensionsTestManagementId)

	//return s.Get()
	return nil
}

func metricExtensionsTestManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func metricExtensionsTestManagementWaitForWorkRequest(wId *string, entityType string, action oci_stack_monitoring.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_stack_monitoring.StackMonitoringClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "stack_monitoring")
	retryPolicy.ShouldRetryOperation = metricExtensionsTestManagementWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromStackMonitoringMetricExtensionsTestManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromStackMonitoringMetricExtensionsTestManagementWorkRequest(client *oci_stack_monitoring.StackMonitoringClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_stack_monitoring.ActionTypeEnum) error {
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

func (s *StackMonitoringMetricExtensionsTestManagementResourceCrud) SetData() error {
	if s.Res.TestRunId != nil {
		s.D.Set("test_run_id", *s.Res.TestRunId)
	}

	if s.Res.TestRunMetricSuffix != nil {
		s.D.Set("test_run_metric_suffix", *s.Res.TestRunMetricSuffix)
	}

	if s.Res.TestRunNamespaceName != nil {
		s.D.Set("test_run_namespace_name", *s.Res.TestRunNamespaceName)
	}

	if s.Res.TestRunResourceGroupName != nil {
		s.D.Set("test_run_resource_group_name", *s.Res.TestRunResourceGroupName)
	}

	return nil
}
