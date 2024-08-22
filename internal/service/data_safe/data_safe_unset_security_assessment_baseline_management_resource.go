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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeUnsetSecurityAssessmentBaselineManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeUnsetSecurityAssessmentBaselineManagement,
		Read:     readDataSafeUnsetSecurityAssessmentBaselineManagement,
		Delete:   deleteDataSafeUnsetSecurityAssessmentBaselineManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"security_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Computed
		},
	}
}

func createDataSafeUnsetSecurityAssessmentBaselineManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUnsetSecurityAssessmentBaselineManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	// Get the compartment id and security assessment id from the resource data
	compartmentId := d.Get("compartment_id").(string)
	SAId := d.Get("security_assessment_id").(string)

	// Wait for set bseline work req to get completed
	_, err := sync.GetBaselineWorkReq(compartmentId, SAId)
	if err != nil {
		return err
	}
	return tfresource.CreateResource(d, sync)
}

func readDataSafeUnsetSecurityAssessmentBaselineManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDataSafeUnsetSecurityAssessmentBaselineManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeUnsetSecurityAssessmentBaselineManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.UnsetSecurityAssessmentBaselineResponse
	DisableNotFoundRetries bool
}

func (s *DataSafeUnsetSecurityAssessmentBaselineManagementResourceCrud) GetBaselineWorkReq(compartmentId string, SAId string) (string, error) {
	listWorkRequestsRequest := oci_data_safe.ListWorkRequestsRequest{SortBy: oci_data_safe.ListWorkRequestsSortByEnum("ACCEPTEDTIME"), SortOrder: oci_data_safe.ListWorkRequestsSortOrderEnum("DESC")}
	var workId *string
	tmp := "SET_SECURITY_ASSESSMENT_BASELINE"
	listWorkRequestsRequest.OperationType = &tmp

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		listWorkRequestsRequest.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("security_assessment_id"); ok {
		tmp := targetId.(string)
		listWorkRequestsRequest.ResourceId = &tmp
	}

	listWorkRequestsResponse, err := s.Client.ListWorkRequests(context.Background(), listWorkRequestsRequest)
	if listWorkRequestsResponse.Items != nil && len(listWorkRequestsResponse.Items) > 0 {
		// Get the latest work request
		var tmp1 = &listWorkRequestsResponse.Items[0]
		workId = tmp1.Id

		for _, res := range tmp1.Resources {
			if strings.Contains(strings.ToLower(*res.EntityType), "securityAssessment") {
				if res.ActionType == oci_data_safe.WorkRequestResourceActionTypeInProgress {
					fmt.Println("IN_PROGRESS Work request found for the given targetId")
					break
				}
			}
		}
	}

	if err != nil {
		return "", err
	}

	if workId != nil {
		return s.getSavedAssessmentIdFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	} else {
		return "No Work request found", nil
	}
}

func (s *DataSafeUnsetSecurityAssessmentBaselineManagementResourceCrud) getSavedAssessmentIdFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) (string, error) {

	// Wait until it finishes
	baselineId, err := unsetSecurityAssessmentBaselineManagementWaitForWorkRequest(workId, "securityassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return "", err
	}
	return *baselineId, nil
}

func (s *DataSafeUnsetSecurityAssessmentBaselineManagementResourceCrud) ID() string {
	return *s.Res.OpcRequestId
}

func (s *DataSafeUnsetSecurityAssessmentBaselineManagementResourceCrud) Create() error {
	request := oci_data_safe.UnsetSecurityAssessmentBaselineRequest{}

	if securityAssessmentId, ok := s.D.GetOkExists("security_assessment_id"); ok {
		tmp := securityAssessmentId.(string)
		request.SecurityAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UnsetSecurityAssessmentBaseline(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response
	workId := response.OpcWorkRequestId
	return s.getUnsetSecurityAssessmentBaselineManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeUnsetSecurityAssessmentBaselineManagementResourceCrud) getUnsetSecurityAssessmentBaselineManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := unsetSecurityAssessmentBaselineManagementWaitForWorkRequest(workId, "securityassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func unsetSecurityAssessmentBaselineManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func unsetSecurityAssessmentBaselineManagementWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = unsetSecurityAssessmentBaselineManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
		return nil, getErrorFromDataSafeUnsetSecurityAssessmentBaselineManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeUnsetSecurityAssessmentBaselineManagementWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeUnsetSecurityAssessmentBaselineManagementResourceCrud) SetData() error {
	return nil
}
