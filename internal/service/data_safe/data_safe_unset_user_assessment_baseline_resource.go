// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeUnsetUserAssessmentBaselineResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeUnsetUserAssessmentBaseline,
		Read:     readDataSafeUnsetUserAssessmentBaseline,
		Delete:   deleteDataSafeUnsetUserAssessmentBaseline,
		Schema: map[string]*schema.Schema{
			// Required
			"user_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"target_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
		},
	}
}

func createDataSafeUnsetUserAssessmentBaseline(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUnsetUserAssessmentBaselineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeUnsetUserAssessmentBaseline(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDataSafeUnsetUserAssessmentBaseline(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeUnsetUserAssessmentBaselineResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.UserAssessment
	DisableNotFoundRetries bool
}

func (s *DataSafeUnsetUserAssessmentBaselineResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeUnsetUserAssessmentBaselineResourceCrud) Get() error {

	request := oci_data_safe.GetUserAssessmentRequest{}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetUserAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UserAssessment
	return nil
}

func (s *DataSafeUnsetUserAssessmentBaselineResourceCrud) Create() error {
	request := oci_data_safe.UnsetUserAssessmentBaselineRequest{}

	if targetIds, ok := s.D.GetOkExists("target_ids"); ok {
		interfaces := targetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("target_ids") {
			request.TargetIds = tmp
		}
	}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UnsetUserAssessmentBaseline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUnsetUserAssessmentBaselineFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeUnsetUserAssessmentBaselineResourceCrud) getUnsetUserAssessmentBaselineFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	unsetUserAssessmentBaselineId, err := unsetUserAssessmentBaselineWaitForWorkRequest(workId, "userassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*unsetUserAssessmentBaselineId)

	return s.Get()
}

func unsetUserAssessmentBaselineWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func unsetUserAssessmentBaselineWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = unsetUserAssessmentBaselineWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeUnsetUserAssessmentBaselineWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeUnsetUserAssessmentBaselineWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeUnsetUserAssessmentBaselineResourceCrud) SetData() error {
	return nil
}
