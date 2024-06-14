// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeMaskingPolicyHealthReportManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeMaskingPolicyHealthReportManagement,
		Read:     readDataSafeMaskingPolicyHealthReportManagement,
		Update:   updateDataSafeMaskingPolicyHealthReportManagement,
		Delete:   deleteDataSafeMaskingPolicyHealthReportManagement,
		Schema: map[string]*schema.Schema{
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"masking_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeMaskingPolicyHealthReportManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyHealthReportManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	err := sync.getMaskingPolicyHealthReportIdFromFilter()
	if err != nil {
		return err
	}

	err1 := sync.Get()
	if err1 != nil {
		return err1
	}

	err = sync.SetData()
	if err != nil {
		return err
	}

	return nil
}

func readDataSafeMaskingPolicyHealthReportManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyHealthReportManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeMaskingPolicyHealthReportManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDataSafeMaskingPolicyHealthReportManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyHealthReportManagementResourceCrud{}
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.D = d

	return tfresource.DeleteResource(d, sync)
}

type DataSafeMaskingPolicyHealthReportManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.MaskingPolicyHealthReport
	DisableNotFoundRetries bool
}

func (s *DataSafeMaskingPolicyHealthReportManagementResourceCrud) Delete() error {
	request := oci_data_safe.DeleteMaskingPolicyHealthReportRequest{}

	tmp := s.D.Id()
	request.MaskingPolicyHealthReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.DeleteMaskingPolicyHealthReport(context.Background(), request)
	if err != nil {
		return err
	}

	if err != nil {
		fmt.Printf("Error deleting MaskingPolicyHealthReport %s, It is possible that the resource is already deleted. Please verify manually \n", err)
	}
	return nil
}

func (s *DataSafeMaskingPolicyHealthReportManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func generateHealthReportWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = maskDataWorkRequestShouldRetryFunc(timeout)

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
		if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), entityType) && res.Identifier != nil {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}
	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed {
		return nil, getErrorFromDataSafeMaskDataWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}
func (s *DataSafeMaskingPolicyHealthReportManagementResourceCrud) getMaskingPolicyHealthReportIdFromFilter() error {
	// Masking health report will be in same compartment as of policy
	getMaskingPolicyRequest := oci_data_safe.GetMaskingPolicyRequest{}
	var compartmentId *string

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		getMaskingPolicyRequest.MaskingPolicyId = &tmp
	}

	getMaskingPolicyResponse, err := s.Client.GetMaskingPolicy(context.Background(), getMaskingPolicyRequest)
	if err != nil {
		return err
	}

	compartmentId = getMaskingPolicyResponse.CompartmentId

	err = s.D.Set("compartment_id", compartmentId)
	if err != nil {
		return err
	}

	// Get the masking health report for given target and masking policy ID
	err = s.GetMaskingPolicyHealthReportList()
	if err != nil {
		return err
	}

	// check if masking health report id is set and masking health report already exists
	if s.D.Id() != "" {
		return nil
	}

	// generate health report
	generateHealthReportRequest := oci_data_safe.GenerateHealthReportRequest{}
	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		generateHealthReportRequest.MaskingPolicyId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		generateHealthReportRequest.TargetId = &tmp
	}

	generateHealthReportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GenerateHealthReport(context.Background(), generateHealthReportRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	maskingPolicyHealthReportId, err := generateHealthReportWaitForWorkRequest(workId, "policy_check_health",
		oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)

	if err == nil {
		s.D.SetId(*maskingPolicyHealthReportId)
		return nil
	}
	return err
}

func (s *DataSafeMaskingPolicyHealthReportManagementResourceCrud) GetMaskingPolicyHealthReportList() error {
	request := oci_data_safe.ListMaskingPolicyHealthReportsRequest{}
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
		request.SortOrder = oci_data_safe.ListMaskingPolicyHealthReportsSortOrderDesc
		request.SortBy = oci_data_safe.ListMaskingPolicyHealthReportsSortByTimecreated
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")
	var maskingPolicyHealthReport = new(oci_data_safe.MaskingPolicyHealthReport)
	var activeStatus bool
	response, err := s.Client.ListMaskingPolicyHealthReports(context.Background(), request)
	if err != nil {
		return err
	}
	if response.MaskingPolicyHealthReportCollection.Items != nil && len(response.MaskingPolicyHealthReportCollection.Items) > 0 {
		temp1 := response.MaskingPolicyHealthReportCollection.Items[0]
		if temp1.LifecycleState == "CREATING" {
			activeStatus, err = WaitForHealthReportStatusSuccess(s, temp1.Id, "MASK_POLICY_GENERATE_HEALTH_REPORT", request.CompartmentId)
			if s.D.SetId(*temp1.Id); activeStatus == true {
				return nil
			}
		}
		maskingPolicyHealthReport.Id = temp1.Id
	}

	if maskingPolicyHealthReport.Id == nil {
		return nil
	}

	s.D.SetId(*maskingPolicyHealthReport.Id)
	return nil
}
func WaitForHealthReportStatusSuccess(s *DataSafeMaskingPolicyHealthReportManagementResourceCrud, resourceId *string, operationType string, compartmentId *string) (bool, error) {
	listWorkRequestsRequest := oci_data_safe.ListWorkRequestsRequest{SortBy: oci_data_safe.ListWorkRequestsSortByEnum("ACCEPTEDTIME"), SortOrder: oci_data_safe.ListWorkRequestsSortOrderEnum("DESC")}
	var workId *string
	tmp := operationType
	listWorkRequestsRequest.OperationType = &tmp
	listWorkRequestsRequest.CompartmentId = compartmentId
	listWorkRequestsRequest.ResourceId = resourceId
	listWorkRequestsResponse, err := s.Client.ListWorkRequests(context.Background(), listWorkRequestsRequest)
	if listWorkRequestsResponse.Items != nil && len(listWorkRequestsResponse.Items) > 0 {
		var tmp1 = &listWorkRequestsResponse.Items[0]
		workId = tmp1.Id
	}

	if err != nil {
		return false, err
	}

	if workId != nil {
		_, err = generateHealthReportWaitForWorkRequest(workId, "policy_check_health",
			oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return true, err
		}
		return true, nil
	} else {
		return false, err
	}
}
func (s *DataSafeMaskingPolicyHealthReportManagementResourceCrud) Get() error {
	request := oci_data_safe.GetMaskingPolicyHealthReportRequest{}

	tmp := s.D.Id()
	request.MaskingPolicyHealthReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetMaskingPolicyHealthReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaskingPolicyHealthReport
	return nil
}

func (s *DataSafeMaskingPolicyHealthReportManagementResourceCrud) SetData() error {

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.MaskingPolicyId != nil {
		s.D.Set("masking_policy_id", *s.Res.MaskingPolicyId)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	return nil
}
