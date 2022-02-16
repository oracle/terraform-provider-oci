// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v58/datasafe"
)

func DataSafeCompareUserAssessmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeCompareUserAssessment,
		Read:     readDataSafeCompareUserAssessment,
		Delete:   deleteDataSafeCompareUserAssessment,
		Schema: map[string]*schema.Schema{
			// Required
			"comparison_user_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createDataSafeCompareUserAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeCompareUserAssessmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeCompareUserAssessment(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDataSafeCompareUserAssessment(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeCompareUserAssessmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.GetUserAssessmentComparisonResponse
	DisableNotFoundRetries bool
}

func (s *DataSafeCompareUserAssessmentResourceCrud) ID() string {
	return *s.Res.OpcRequestId
}

func (s *DataSafeCompareUserAssessmentResourceCrud) Get() error {
	request := oci_data_safe.GetUserAssessmentComparisonRequest{}

	if comparisonUserAssessmentId, ok := s.D.GetOkExists("comparison_user_assessment_id"); ok {
		tmp := comparisonUserAssessmentId.(string)
		request.ComparisonUserAssessmentId = &tmp
	}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetUserAssessmentComparison(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeCompareUserAssessmentResourceCrud) Create() error {
	request := oci_data_safe.CompareUserAssessmentRequest{}

	if comparisonUserAssessmentId, ok := s.D.GetOkExists("comparison_user_assessment_id"); ok {
		tmp := comparisonUserAssessmentId.(string)
		request.ComparisonUserAssessmentId = &tmp
	}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CompareUserAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCompareUserAssessmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeCompareUserAssessmentResourceCrud) getCompareUserAssessmentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	compareUserAssessmentId, err := compareUserAssessmentWaitForWorkRequest(workId, "userassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*compareUserAssessmentId)

	return s.Get()
}

func compareUserAssessmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func compareUserAssessmentWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = compareUserAssessmentWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
			// string(oci_data_safe.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			// string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed {
		return nil, getErrorFromDataSafeCompareUserAssessmentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeCompareUserAssessmentWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
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

func (s *DataSafeCompareUserAssessmentResourceCrud) SetData() error {
	return nil
}

func (s *DataSafeCompareUserAssessmentResourceCrud) Delete() error {
	return nil
}
