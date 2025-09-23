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

func ResourceAnalyticsTenancyAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &tfresource.TwoHours,
			Update: &tfresource.TwoHours,
			Delete: &tfresource.TwoHours,
			Read:   &tfresource.TwoHours,
		},
		Create: createResourceAnalyticsTenancyAttachment,
		Read:   readResourceAnalyticsTenancyAttachment,
		Update: updateResourceAnalyticsTenancyAttachment,
		Delete: deleteResourceAnalyticsTenancyAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"resource_analytics_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"is_reporting_tenancy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
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

func createResourceAnalyticsTenancyAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsTenancyAttachmentResourceCrud{}
	sync.D = d
	clients := m.(*client.OracleClients)
	sync.Client = clients.TenancyAttachmentClient()
	sync.WorkReqClient = clients.ResourceAnalyticsInstanceClient()

	return tfresource.CreateResource(d, sync)
}

func readResourceAnalyticsTenancyAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsTenancyAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).TenancyAttachmentClient()

	return tfresource.ReadResource(sync)
}

func updateResourceAnalyticsTenancyAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsTenancyAttachmentResourceCrud{}
	sync.D = d
	clients := m.(*client.OracleClients)
	sync.Client = clients.TenancyAttachmentClient()
	sync.WorkReqClient = clients.ResourceAnalyticsInstanceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteResourceAnalyticsTenancyAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsTenancyAttachmentResourceCrud{}
	sync.D = d
	clients := m.(*client.OracleClients)
	sync.Client = clients.TenancyAttachmentClient()
	sync.WorkReqClient = clients.ResourceAnalyticsInstanceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ResourceAnalyticsTenancyAttachmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_resource_analytics.TenancyAttachmentClient
	WorkReqClient          *oci_resource_analytics.ResourceAnalyticsInstanceClient
	Res                    *oci_resource_analytics.TenancyAttachment
	DisableNotFoundRetries bool
}

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_resource_analytics.TenancyAttachmentLifecycleStateCreating),
	}
}

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_resource_analytics.TenancyAttachmentLifecycleStateActive),
		string(oci_resource_analytics.TenancyAttachmentLifecycleStateNeedsAttention),
	}
}

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_resource_analytics.TenancyAttachmentLifecycleStateDeleting),
	}
}

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_resource_analytics.TenancyAttachmentLifecycleStateDeleted),
	}
}

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) Create() error {
	request := oci_resource_analytics.CreateTenancyAttachmentRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
		tmp := resourceAnalyticsInstanceId.(string)
		request.ResourceAnalyticsInstanceId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.CreateTenancyAttachment(context.Background(), request)
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "tenancyattachment") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getTenancyAttachmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) getTenancyAttachmentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_resource_analytics.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	tenancyAttachmentId, err := tenancyAttachmentWaitForWorkRequest(workId, "tenancyattachment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkReqClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] getTenancyAttachmentFromWorkRequest: creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, tenancyAttachmentId)
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
	s.D.SetId(*tenancyAttachmentId)

	return s.Get()
}

func tenancyAttachmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func tenancyAttachmentWaitForWorkRequest(wId *string, entityType string, action oci_resource_analytics.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, workReqClient *oci_resource_analytics.ResourceAnalyticsInstanceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "resource_analytics")
	retryPolicy.ShouldRetryOperation = tenancyAttachmentWorkRequestShouldRetryFunc(timeout)

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

	fmt.Printf("[DEBUG] Searching in WR.resources for [%v] to become [%v] \n", entityType, action)

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
		return nil, getErrorFromResourceAnalyticsTenancyAttachmentWorkRequest(workReqClient, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromResourceAnalyticsTenancyAttachmentWorkRequest(workReqClient *oci_resource_analytics.ResourceAnalyticsInstanceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_resource_analytics.ActionTypeEnum) error {
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

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) Get() error {
	request := oci_resource_analytics.GetTenancyAttachmentRequest{}

	tmp := s.D.Id()
	request.TenancyAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.GetTenancyAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TenancyAttachment
	return nil
}

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) Update() error {
	request := oci_resource_analytics.UpdateTenancyAttachmentRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.TenancyAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.UpdateTenancyAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTenancyAttachmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) Delete() error {
	request := oci_resource_analytics.DeleteTenancyAttachmentRequest{}

	tmp := s.D.Id()
	request.TenancyAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.DeleteTenancyAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := tenancyAttachmentWaitForWorkRequest(workId, "tenancyattachment",
		oci_resource_analytics.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkReqClient)
	return delWorkRequestErr
}

func (s *ResourceAnalyticsTenancyAttachmentResourceCrud) SetData() error {
	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.IsReportingTenancy != nil {
		s.D.Set("is_reporting_tenancy", *s.Res.IsReportingTenancy)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ResourceAnalyticsInstanceId != nil {
		s.D.Set("resource_analytics_instance_id", *s.Res.ResourceAnalyticsInstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func TenancyAttachmentSummaryToMap(obj oci_resource_analytics.TenancyAttachmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsReportingTenancy != nil {
		result["is_reporting_tenancy"] = bool(*obj.IsReportingTenancy)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ResourceAnalyticsInstanceId != nil {
		result["resource_analytics_instance_id"] = string(*obj.ResourceAnalyticsInstanceId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
