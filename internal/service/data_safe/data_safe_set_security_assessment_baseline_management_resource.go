// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"errors"

	//"errors"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeSetSecurityAssessmentBaselineManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSetSecurityAssessmentBaselineManagement,
		Read:     readDataSafeSetSecurityAssessmentBaselineManagement,
		Delete:   deleteDataSafeSetSecurityAssessmentBaselineManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"target_id": {
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
			"security_assessment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assessment_ids": {
				Type:     schema.TypeList,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createDataSafeSetSecurityAssessmentBaselineManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSetSecurityAssessmentBaselineManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	// Get the target ID and compartment id from the resource data
	targetId := d.Get("target_id").(string)
	compartmentId := d.Get("compartment_id").(string)

	// Get Id of Recent Saved SA which can be set as baseline
	securityAssessmentId, err := sync.GetAssessmentWorkReq(targetId, compartmentId)
	if err != nil {
		return err
	}

	// Set the security assessment ID and assessment_ids in the resource data for setting baseline
	d.Set("security_assessment_id", securityAssessmentId)
	d.Set("assessment_ids", []interface{}{securityAssessmentId})
	return tfresource.CreateResource(d, sync)
}

func readDataSafeSetSecurityAssessmentBaselineManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDataSafeSetSecurityAssessmentBaselineManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeSetSecurityAssessmentBaselineManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SetSecurityAssessmentBaselineResponse
	DisableNotFoundRetries bool
}

func (s *DataSafeSetSecurityAssessmentBaselineManagementResourceCrud) GetAssessmentWorkReq(targetId string, compartmentId string) (string, error) {
	listWorkRequestsRequest := oci_data_safe.ListWorkRequestsRequest{SortBy: oci_data_safe.ListWorkRequestsSortByEnum("ACCEPTEDTIME"), SortOrder: oci_data_safe.ListWorkRequestsSortOrderEnum("DESC")}
	var workId *string
	tmp := "CREATE_SECURITY_ASSESSMENT"
	listWorkRequestsRequest.OperationType = &tmp

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		listWorkRequestsRequest.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
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
		return s.getSavedAssessmentIdFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate), listWorkRequestsRequest.ResourceId, listWorkRequestsRequest.CompartmentId)
	} else {
		return s.GetSavedAssessmentList(targetId, compartmentId)
	}
}
func (s *DataSafeSetSecurityAssessmentBaselineManagementResourceCrud) getSavedAssessmentIdFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration, targetId *string, compartmentId *string) (string, error) {

	// Wait until it finishes
	_, err := securityAssessmentWaitForWorkRequest(workId, "securityassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return "", err
	}

	return s.GetSavedAssessmentList(*targetId, *compartmentId)
}

func (s *DataSafeSetSecurityAssessmentBaselineManagementResourceCrud) GetSavedAssessmentList(targetId string, compartmentId string) (string, error) {
	// Set up the request
	request := oci_data_safe.ListSecurityAssessmentsRequest{
		CompartmentId: &compartmentId,
		TargetId:      &targetId,
		// Set the desired assessment type to ListSecurityAssessmentsTypeSaved
		Type:      oci_data_safe.ListSecurityAssessmentsTypeSaved,
		SortOrder: oci_data_safe.ListSecurityAssessmentsSortOrderDesc,
		SortBy:    oci_data_safe.ListSecurityAssessmentsSortByTimecreated,
	}

	// Call the API to list security assessments
	response, err := s.Client.ListSecurityAssessments(context.Background(), request)
	if err != nil {
		return "", err
	}

	// Check if there are security assessments in the response
	for _, assessment := range response.Items {
		// Find the first assessment that is Saved and not set to baseline

		if assessment.IsBaseline != nil && !*assessment.IsBaseline && assessment.LifecycleState == oci_data_safe.SecurityAssessmentLifecycleStateSucceeded {
			fmt.Printf("Found relevant assessment: %+v\n", assessment)
			return *assessment.Id, nil
		}
	}
	// If no ID is found, return an error
	return "", errors.New("no security assessment ID found with baseline set to false")
}

func (s *DataSafeSetSecurityAssessmentBaselineManagementResourceCrud) ID() string {
	return *s.Res.OpcRequestId
}

func (s *DataSafeSetSecurityAssessmentBaselineManagementResourceCrud) Create() error {
	request := oci_data_safe.SetSecurityAssessmentBaselineRequest{}
	baseLineStruct := oci_data_safe.SecurityAssessmentBaseLineDetails{}
	if assessmentIds, ok := s.D.GetOkExists("assessment_ids"); ok {
		interfaces := assessmentIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("assessment_ids") {
			baseLineStruct.AssessmentIds = tmp
		}
	}

	request.BaseLineDetails = baseLineStruct

	if securityAssessmentId, ok := s.D.GetOkExists("security_assessment_id"); ok {
		tmp := securityAssessmentId.(string)
		request.SecurityAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.SetSecurityAssessmentBaseline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSetSecurityAssessmentBaselineFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSetSecurityAssessmentBaselineManagementResourceCrud) Get() error {
	request := oci_data_safe.SetSecurityAssessmentBaselineRequest{}
	baseLineStruct := oci_data_safe.SecurityAssessmentBaseLineDetails{}
	if assessmentIds, ok := s.D.GetOkExists("assessment_ids"); ok {
		interfaces := assessmentIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("assessment_ids") {
			baseLineStruct.AssessmentIds = tmp
		}
	}
	request.BaseLineDetails = baseLineStruct

	if securityAssessmentId, ok := s.D.GetOkExists("security_assessment_id"); ok {
		tmp := securityAssessmentId.(string)
		request.SecurityAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.SetSecurityAssessmentBaseline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSetSecurityAssessmentBaselineManagementResourceCrud) getSetSecurityAssessmentBaselineFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	setSecurityAssessmentBaselineId, err := setSecurityAssessmentBaselineManagementWaitForWorkRequest(workId, "securityassessment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*setSecurityAssessmentBaselineId)

	return s.Get()
}

func setSecurityAssessmentBaselineManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func setSecurityAssessmentBaselineManagementWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = setSecurityAssessmentBaselineManagementWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataSafeSetSecurityAssessmentBaselineManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSetSecurityAssessmentBaselineManagementWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
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

func (s *DataSafeSetSecurityAssessmentBaselineManagementResourceCrud) SetData() error {
	return nil
}
