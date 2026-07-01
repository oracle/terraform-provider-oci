// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSafeMaskingPolicyApplyDifferenceToMaskingColumnsResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeMaskingPolicyApplyDifferenceToMaskingColumnsWithContext,
		ReadContext:   readDataSafeMaskingPolicyApplyDifferenceToMaskingColumnsWithContext,
		DeleteContext: deleteDataSafeMaskingPolicyApplyDifferenceToMaskingColumnsWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"sdm_masking_policy_difference_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"masking_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createDataSafeMaskingPolicyApplyDifferenceToMaskingColumnsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeMaskingPolicyApplyDifferenceToMaskingColumnsResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDataSafeMaskingPolicyApplyDifferenceToMaskingColumnsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteDataSafeMaskingPolicyApplyDifferenceToMaskingColumnsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type DataSafeMaskingPolicyApplyDifferenceToMaskingColumnsResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.MaskingPolicy
	DisableNotFoundRetries bool
}

func (s *DataSafeMaskingPolicyApplyDifferenceToMaskingColumnsResourceCrud) ID() string {
	return s.D.Id()
}

func (s *DataSafeMaskingPolicyApplyDifferenceToMaskingColumnsResourceCrud) GetWithContext(ctx context.Context) error {
	return nil
}

func (s *DataSafeMaskingPolicyApplyDifferenceToMaskingColumnsResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_data_safe.ApplySdmMaskingPolicyDifferenceRequest{}

	if MaskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := MaskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if SdmMaskingPolicyDifferenceId, ok := s.D.GetOkExists("sdm_masking_policy_difference_id"); ok {
		tmp := SdmMaskingPolicyDifferenceId.(string)
		request.SdmMaskingPolicyDifferenceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ApplySdmMaskingPolicyDifference(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDataSafeMaskingPolicyApplyDifferenceToMaskingColumnsFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeMaskingPolicyApplyDifferenceToMaskingColumnsResourceCrud) getDataSafeMaskingPolicyApplyDifferenceToMaskingColumnsFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	MaskingPolicyId, err := applyDifferenceToMaskingColumnsWaitForWorkRequest(ctx, workId, "masking_policy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*MaskingPolicyId)

	return s.GetWithContext(ctx)
}

func applyDifferenceToMaskingColumnsWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func applyDifferenceToMaskingColumnsWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = applyDifferenceToMaskingColumnsWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_data_safe.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed {
		return nil, getErrorFromDataSafeMaskingPolicyApplyDifferenceToMaskingColumnsWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeMaskingPolicyApplyDifferenceToMaskingColumnsWorkRequest(ctx context.Context, client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_data_safe.ListWorkRequestErrorsRequest{
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

func (s *DataSafeMaskingPolicyApplyDifferenceToMaskingColumnsResourceCrud) SetData() error {
	return nil
}

func (s *DataSafeMaskingPolicyApplyDifferenceToMaskingColumnsResourceCrud) DeleteWithContext(ctx context.Context) error {
	return nil
}
