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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSensitiveDataModelsApplyDiscoveryJobResults,
		Read:     readDataSafeSensitiveDataModelsApplyDiscoveryJobResults,
		Delete:   deleteDataSafeSensitiveDataModelsApplyDiscoveryJobResults,
		Schema: map[string]*schema.Schema{
			// Required
			"discovery_job_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createDataSafeSensitiveDataModelsApplyDiscoveryJobResults(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeSensitiveDataModelsApplyDiscoveryJobResults(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDataSafeSensitiveDataModelsApplyDiscoveryJobResults(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SensitiveDataModel
	DisableNotFoundRetries bool
}

func (s *DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResourceCrud) ID() string {
	return s.D.Id()
}

func (s *DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResourceCrud) Get() error {
	return nil
}

func (s *DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResourceCrud) Create() error {
	request := oci_data_safe.ApplyDiscoveryJobResultsRequest{}

	if SensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := SensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	if DiscoveryJobId, ok := s.D.GetOkExists("discovery_job_id"); ok {
		tmp := DiscoveryJobId.(string)
		request.DiscoveryJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ApplyDiscoveryJobResults(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDataSafeSensitiveDataModelsApplyDiscoveryJobResultsFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResourceCrud) getDataSafeSensitiveDataModelsApplyDiscoveryJobResultsFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	sensitiveDataModelId, err := applyDiscoveryJobResultsWaitForWorkRequest(workId, "sensitivedatamodel",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*sensitiveDataModelId)

	return s.Get()
}

func applyDiscoveryJobResultsWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func applyDiscoveryJobResultsWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = applyDiscoveryJobResultsWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeSensitiveDataModelsApplyDiscoveryJobResultsWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSensitiveDataModelsApplyDiscoveryJobResultsWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResourceCrud) SetData() error {
	return nil
}

func (s *DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResourceCrud) Delete() error {
	return nil
}
