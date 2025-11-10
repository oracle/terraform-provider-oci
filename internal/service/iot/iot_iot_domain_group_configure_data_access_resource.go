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

func IotIotDomainGroupConfigureDataAccessResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotIotDomainGroupConfigureDataAccessWithContext,
		ReadContext:   readIotIotDomainGroupConfigureDataAccessWithContext,
		DeleteContext: deleteIotIotDomainGroupConfigureDataAccessWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"db_allow_listed_vcn_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"iot_domain_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createIotIotDomainGroupConfigureDataAccessWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotDomainGroupConfigureDataAccessResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotIotDomainGroupConfigureDataAccessWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteIotIotDomainGroupConfigureDataAccessWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type IotIotDomainGroupConfigureDataAccessResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.IotDomainGroup
	DisableNotFoundRetries bool
}

func (s *IotIotDomainGroupConfigureDataAccessResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IotIotDomainGroupConfigureDataAccessResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.ConfigureIotDomainGroupDataAccessRequest{}

	interfaces := s.D.Get("db_allow_listed_vcn_ids").([]interface{})
	tmp := make([]string, len(interfaces))
	for i := range interfaces {
		if interfaces[i] != nil {
			tmp[i] = interfaces[i].(string)
		}
	}
	request.DbAllowListedVcnIds = tmp

	if iotDomainGroupId, ok := s.D.GetOkExists("iot_domain_group_id"); ok {
		tmp := iotDomainGroupId.(string)
		request.IotDomainGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.ConfigureIotDomainGroupDataAccess(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIotDomainGroupConfigureDataAccessFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot"), oci_iot.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IotIotDomainGroupConfigureDataAccessResourceCrud) getIotDomainGroupConfigureDataAccessFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_iot.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	iotDomainGroupId, err := iotDomainGroupConfigureDataAccessWaitForWorkRequest(ctx, workId, "iotdomaingroup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*iotDomainGroupId)

	return s.GetWithContext(ctx)
}

func (s *IotIotDomainGroupConfigureDataAccessResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetIotDomainGroupRequest{}

	tmp := s.D.Id()
	request.IotDomainGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetIotDomainGroup(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.IotDomainGroup
	return nil
}

func iotDomainGroupConfigureDataAccessWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func iotDomainGroupConfigureDataAccessWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_iot.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_iot.IotClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "iot")
	retryPolicy.ShouldRetryOperation = iotDomainGroupConfigureDataAccessWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromIotIotDomainGroupConfigureDataAccessWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromIotIotDomainGroupConfigureDataAccessWorkRequest(ctx context.Context, client *oci_iot.IotClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_iot.ActionTypeEnum) error {
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

func (s *IotIotDomainGroupConfigureDataAccessResourceCrud) SetData() error {
	return nil
}
