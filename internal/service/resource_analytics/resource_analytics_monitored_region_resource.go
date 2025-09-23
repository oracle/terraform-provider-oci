// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_analytics

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_resource_analytics "github.com/oracle/oci-go-sdk/v65/resourceanalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourceAnalyticsMonitoredRegionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Read:   &tfresource.TwoHours,
			Create: &tfresource.TwoHours,
			Update: &tfresource.TwoHours,
			Delete: &tfresource.TwoHours,
		},
		Create: createResourceAnalyticsMonitoredRegion,
		Read:   readResourceAnalyticsMonitoredRegion,
		Delete: deleteResourceAnalyticsMonitoredRegion,
		Schema: map[string]*schema.Schema{
			// Required
			"region_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_analytics_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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

func createResourceAnalyticsMonitoredRegion(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsMonitoredRegionResourceCrud{}
	sync.D = d
	clients := m.(*client.OracleClients)
	sync.Client = clients.MonitoredRegionClient()
	sync.WorkReqClient = clients.ResourceAnalyticsInstanceClient()

	return tfresource.CreateResource(d, sync)
}

func readResourceAnalyticsMonitoredRegion(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsMonitoredRegionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MonitoredRegionClient()

	return tfresource.ReadResource(sync)
}

func deleteResourceAnalyticsMonitoredRegion(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsMonitoredRegionResourceCrud{}
	sync.D = d
	clients := m.(*client.OracleClients)
	sync.Client = clients.MonitoredRegionClient()
	sync.WorkReqClient = clients.ResourceAnalyticsInstanceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ResourceAnalyticsMonitoredRegionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_resource_analytics.MonitoredRegionClient
	WorkReqClient          *oci_resource_analytics.ResourceAnalyticsInstanceClient
	Res                    *oci_resource_analytics.MonitoredRegion
	DisableNotFoundRetries bool
}

func (s *ResourceAnalyticsMonitoredRegionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ResourceAnalyticsMonitoredRegionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_resource_analytics.MonitoredRegionLifecycleStateCreating),
	}
}

func (s *ResourceAnalyticsMonitoredRegionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_resource_analytics.MonitoredRegionLifecycleStateActive),
	}
}

func (s *ResourceAnalyticsMonitoredRegionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_resource_analytics.MonitoredRegionLifecycleStateDeleting),
	}
}

func (s *ResourceAnalyticsMonitoredRegionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_resource_analytics.MonitoredRegionLifecycleStateDeleted),
	}
}

func (s *ResourceAnalyticsMonitoredRegionResourceCrud) Create() error {
	request := oci_resource_analytics.CreateMonitoredRegionRequest{}

	if regionId, ok := s.D.GetOkExists("region_id"); ok {
		tmp := regionId.(string)
		request.RegionId = &tmp
	}

	if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
		tmp := resourceAnalyticsInstanceId.(string)
		request.ResourceAnalyticsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.CreateMonitoredRegion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_resource_analytics.GetWorkRequestResponse{}
	workRequestResponse, err = s.WorkReqClient.GetWorkRequest(context.Background(),
		oci_resource_analytics.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "monitoredregion") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getMonitoredRegionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ResourceAnalyticsMonitoredRegionResourceCrud) getMonitoredRegionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_resource_analytics.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	monitoredRegionId, err := monitoredRegionWaitForWorkRequest(workId, "monitoredregion",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkReqClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] getMonitoredRegionFromWorkRequest: creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, monitoredRegionId)
		_, cancelErr := s.WorkReqClient.CancelWorkRequest(context.Background(),
			oci_resource_analytics.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*monitoredRegionId)

	return s.Get()
}

func monitoredRegionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "resource_analytics", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_resource_analytics.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func monitoredRegionWaitForWorkRequest(wId *string, entityType string, action oci_resource_analytics.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, workReqClient *oci_resource_analytics.ResourceAnalyticsInstanceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "resource_analytics")
	retryPolicy.ShouldRetryOperation = monitoredRegionWorkRequestShouldRetryFunc(timeout)

	response := oci_resource_analytics.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_resource_analytics.OperationStatusInProgress),
			string(oci_resource_analytics.OperationStatusAccepted),
			string(oci_resource_analytics.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_resource_analytics.OperationStatusSucceeded),
			string(oci_resource_analytics.OperationStatusFailed),
			string(oci_resource_analytics.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = workReqClient.GetWorkRequest(context.Background(),
				oci_resource_analytics.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_resource_analytics.OperationStatusFailed || response.Status == oci_resource_analytics.OperationStatusCanceled {
		return nil, getErrorFromResourceAnalyticsMonitoredRegionWorkRequest(workReqClient, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromResourceAnalyticsMonitoredRegionWorkRequest(workReqClient *oci_resource_analytics.ResourceAnalyticsInstanceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_resource_analytics.ActionTypeEnum) error {
	response, err := workReqClient.ListWorkRequestErrors(context.Background(),
		oci_resource_analytics.ListWorkRequestErrorsRequest{
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

func (s *ResourceAnalyticsMonitoredRegionResourceCrud) Get() error {
	request := oci_resource_analytics.GetMonitoredRegionRequest{}

	tmp := s.D.Id()
	request.MonitoredRegionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.GetMonitoredRegion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredRegion
	return nil
}

func (s *ResourceAnalyticsMonitoredRegionResourceCrud) Delete() error {
	request := oci_resource_analytics.DeleteMonitoredRegionRequest{}

	tmp := s.D.Id()
	request.MonitoredRegionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.DeleteMonitoredRegion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := monitoredRegionWaitForWorkRequest(workId, "monitoredregion",
		oci_resource_analytics.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkReqClient)
	return delWorkRequestErr
}

func (s *ResourceAnalyticsMonitoredRegionResourceCrud) SetData() error {
	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RegionId != nil {
		s.D.Set("region_id", *s.Res.RegionId)
	}

	if s.Res.ResourceAnalyticsInstanceId != nil {
		s.D.Set("resource_analytics_instance_id", *s.Res.ResourceAnalyticsInstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func MonitoredRegionSummaryToMap(obj oci_resource_analytics.MonitoredRegionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.RegionId != nil {
		result["region_id"] = string(*obj.RegionId)
	}

	if obj.ResourceAnalyticsInstanceId != nil {
		result["resource_analytics_instance_id"] = string(*obj.ResourceAnalyticsInstanceId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
