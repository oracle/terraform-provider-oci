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

func IotIotDomainChangeDataRetentionPeriodResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotIotDomainChangeDataRetentionPeriodWithContext,
		ReadContext:   readIotIotDomainChangeDataRetentionPeriodWithContext,
		DeleteContext: deleteIotIotDomainChangeDataRetentionPeriodWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"data_retention_period_in_days": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"iot_domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createIotIotDomainChangeDataRetentionPeriodWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotDomainChangeDataRetentionPeriodResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotIotDomainChangeDataRetentionPeriodWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteIotIotDomainChangeDataRetentionPeriodWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type IotIotDomainChangeDataRetentionPeriodResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.IotDomain
	DisableNotFoundRetries bool
}

func (s *IotIotDomainChangeDataRetentionPeriodResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IotIotDomainChangeDataRetentionPeriodResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.ChangeIotDomainDataRetentionPeriodRequest{}

	if dataRetentionPeriodInDays, ok := s.D.GetOkExists("data_retention_period_in_days"); ok {
		tmp := dataRetentionPeriodInDays.(int)
		request.DataRetentionPeriodInDays = &tmp
	}

	if iotDomainId, ok := s.D.GetOkExists("iot_domain_id"); ok {
		tmp := iotDomainId.(string)
		request.IotDomainId = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_iot.ChangeIotDomainDataRetentionPeriodDetailsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.ChangeIotDomainDataRetentionPeriod(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIotDomainChangeDataRetentionPeriodFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot"), oci_iot.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IotIotDomainChangeDataRetentionPeriodResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetIotDomainRequest{}

	tmp := s.D.Id()
	request.IotDomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetIotDomain(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.IotDomain
	return nil
}

func (s *IotIotDomainChangeDataRetentionPeriodResourceCrud) getIotDomainChangeDataRetentionPeriodFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_iot.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	iotDomainId, err := iotDomainChangeDataRetentionPeriodWaitForWorkRequest(ctx, workId, "iotdomain",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*iotDomainId)

	return s.GetWithContext(ctx)
}

func iotDomainChangeDataRetentionPeriodWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func iotDomainChangeDataRetentionPeriodWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_iot.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_iot.IotClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "iot")
	retryPolicy.ShouldRetryOperation = iotDomainChangeDataRetentionPeriodWorkRequestShouldRetryFunc(timeout)

	response := oci_iot.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_iot.OperationStatusInProgress),
			string(oci_iot.OperationStatusAccepted),
			string(oci_iot.OperationStatusWaiting),
		},
		Target: []string{
			string(oci_iot.OperationStatusSucceeded),
			string(oci_iot.OperationStatusFailed),
			string(oci_iot.OperationStatusNeedsAttention),
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
	if identifier == nil || response.Status == oci_iot.OperationStatusFailed || response.Status == oci_iot.OperationStatusNeedsAttention {
		return nil, getErrorFromIotIotDomainChangeDataRetentionPeriodWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromIotIotDomainChangeDataRetentionPeriodWorkRequest(ctx context.Context, client *oci_iot.IotClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_iot.ActionTypeEnum) error {
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

func (s *IotIotDomainChangeDataRetentionPeriodResourceCrud) SetData() error {
	return nil
}
