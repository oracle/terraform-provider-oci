// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotIotFlowRuntimeActivateResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotIotFlowRuntimeActivateWithContext,
		ReadContext:   readIotIotFlowRuntimeActivateWithContext,
		DeleteContext: deleteIotIotFlowRuntimeActivateWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"iot_flow_runtime_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createIotIotFlowRuntimeActivateWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotFlowRuntimeActivateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotIotFlowRuntimeActivateWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteIotIotFlowRuntimeActivateWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type IotIotFlowRuntimeActivateResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.IotFlowRuntime
	DisableNotFoundRetries bool
}

func (s *IotIotFlowRuntimeActivateResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IotIotFlowRuntimeActivateResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.ActivateIotFlowRuntimeRequest{}

	if iotFlowRuntimeId, ok := s.D.GetOkExists("iot_flow_runtime_id"); ok {
		tmp := iotFlowRuntimeId.(string)
		request.IotFlowRuntimeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.ActivateIotFlowRuntime(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIotFlowRuntimeActivateFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot"), oci_iot.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IotIotFlowRuntimeActivateResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetIotFlowRuntimeRequest{}

	tmp := s.D.Id()
	request.IotFlowRuntimeId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetIotFlowRuntime(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.IotFlowRuntime
	return nil
}

func (s *IotIotFlowRuntimeActivateResourceCrud) getIotFlowRuntimeActivateFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_iot.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	iotFlowRuntimeActivateId, err := iotFlowRuntimeActivateWaitForWorkRequest(ctx, workId, "iotflowruntime",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*iotFlowRuntimeActivateId)

	return s.GetWithContext(ctx)
}

func iotFlowRuntimeActivateWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "iot", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_iot.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func iotFlowRuntimeActivateWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_iot.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_iot.IotClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "iot")
	retryPolicy.ShouldRetryOperation = iotFlowRuntimeActivateWorkRequestShouldRetryFunc(timeout)

	response := oci_iot.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_iot.OperationStatusAccepted),
			string(oci_iot.OperationStatusInProgress),
			string(oci_iot.OperationStatusWaiting),
			string(oci_iot.OperationStatusNeedsAttention),
		},
		Target: []string{
			string(oci_iot.OperationStatusFailed),
			string(oci_iot.OperationStatusSucceeded),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_iot.GetWorkRequestRequest{
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
	if _, e := stateConf.WaitForStateContext(ctx); e != nil {
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
	if identifier == nil || response.Status == oci_iot.OperationStatusFailed || response.Status == oci_iot.OperationStatusNeedsAttention {
		return nil, getErrorFromIotIotFlowRuntimeActivateWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromIotIotFlowRuntimeActivateWorkRequest(ctx context.Context, client *oci_iot.IotClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_iot.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_iot.ListWorkRequestErrorsRequest{
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

func (s *IotIotFlowRuntimeActivateResourceCrud) SetData() error {
	return nil
}
