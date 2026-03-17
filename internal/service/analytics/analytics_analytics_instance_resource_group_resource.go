// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package analytics

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AnalyticsAnalyticsInstanceResourceGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createAnalyticsAnalyticsInstanceResourceGroupWithContext,
		ReadContext:   readAnalyticsAnalyticsInstanceResourceGroupWithContext,
		UpdateContext: updateAnalyticsAnalyticsInstanceResourceGroupWithContext,
		DeleteContext: deleteAnalyticsAnalyticsInstanceResourceGroupWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"analytics_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"capacity": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
		},
	}
}

func createAnalyticsAnalyticsInstanceResourceGroupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AnalyticsAnalyticsInstanceResourceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readAnalyticsAnalyticsInstanceResourceGroupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AnalyticsAnalyticsInstanceResourceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateAnalyticsAnalyticsInstanceResourceGroupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AnalyticsAnalyticsInstanceResourceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteAnalyticsAnalyticsInstanceResourceGroupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AnalyticsAnalyticsInstanceResourceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type AnalyticsAnalyticsInstanceResourceGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_analytics.AnalyticsClient
	Res                    *oci_analytics.InstanceResourceGroup
	DisableNotFoundRetries bool
}

func (s *AnalyticsAnalyticsInstanceResourceGroupResourceCrud) ID() string {
	return GetAnalyticsInstanceResourceGroupCompositeId(s.D.Get("analytics_instance_id").(string), *s.Res.Id)
}

func (s *AnalyticsAnalyticsInstanceResourceGroupResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_analytics.CreateResourceGroupRequest{}

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	if capacity, ok := s.D.GetOkExists("capacity"); ok {
		tmp := capacity.(int)
		request.Capacity = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.CreateResourceGroup(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsInstanceResourceGroupFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultResourceGroupCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AnalyticsAnalyticsInstanceResourceGroupResourceCrud) getAnalyticsInstanceResourceGroupFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_analytics.WorkRequestActionResultEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := analyticsInstanceResourceGroupWaitForWorkRequest(ctx, workId, "resourcegroup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v\n", workId)
		_, cancelErr := s.Client.DeleteWorkRequest(ctx,
			oci_analytics.DeleteWorkRequestRequest{
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

	analyticsInstanceId := s.D.Get("analytics_instance_id").(string)
	resourceName := s.D.Get("resource_name").(string)

	getRequest := oci_analytics.GetAnalyticsInstanceRequest{}
	getRequest.AnalyticsInstanceId = &analyticsInstanceId
	getRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	getResponse, err := s.Client.GetAnalyticsInstance(ctx, getRequest)
	if err != nil {
		return err
	}

	resourceNameToFind := resourceName
	for _, resourceGroup := range getResponse.AnalyticsInstance.ResourceGroups {
		if resourceGroup.ResourceName == nil || *resourceGroup.ResourceName != resourceNameToFind {
			continue
		}

		if resourceGroup.Id == nil {
			return fmt.Errorf("[ERROR] analytics instance resource group %q was returned without an id", resourceNameToFind)
		}

		s.D.SetId(GetAnalyticsInstanceResourceGroupCompositeId(analyticsInstanceId, *resourceGroup.Id))
		return s.GetWithContext(ctx)
	}

	return fmt.Errorf("[ERROR] unable to find analytics instance resource group %q in analytics instance %s", resourceNameToFind, analyticsInstanceId)
}

func analyticsInstanceResourceGroupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "analytics", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_analytics.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func analyticsInstanceResourceGroupWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_analytics.WorkRequestActionResultEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_analytics.AnalyticsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "analytics")
	retryPolicy.ShouldRetryOperation = analyticsInstanceResourceGroupWorkRequestShouldRetryFunc(timeout)

	response := oci_analytics.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_analytics.WorkRequestStatusInProgress),
			string(oci_analytics.WorkRequestStatusAccepted),
			string(oci_analytics.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_analytics.WorkRequestStatusSucceeded),
			string(oci_analytics.WorkRequestStatusFailed),
			string(oci_analytics.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_analytics.GetWorkRequestRequest{
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
	for _, res := range response.WorkRequest.Resources {
		if res.ResourceType == oci_analytics.WorkRequestResourceTypeAnalyticsInstance {
			if res.ActionResult == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_analytics.WorkRequestStatusFailed || response.Status == oci_analytics.WorkRequestStatusCanceled {
		return nil, getErrorFromAnalyticsAnalyticsInstanceResourceGroupWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAnalyticsAnalyticsInstanceResourceGroupWorkRequest(ctx context.Context, client *oci_analytics.AnalyticsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_analytics.WorkRequestActionResultEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_analytics.ListWorkRequestErrorsRequest{
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

func (s *AnalyticsAnalyticsInstanceResourceGroupResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_analytics.GetResourceGroupRequest{}

	analyticsInstanceId, analyticsInstanceResourceGroupId, err := parseAnalyticsInstanceResourceGroupCompositeId(s.D.Id())
	if err == nil {
		request.AnalyticsInstanceId = &analyticsInstanceId
		request.AnalyticsInstanceResourceGroupId = &analyticsInstanceResourceGroupId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.GetResourceGroup(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceResourceGroup
	return nil
}

func (s *AnalyticsAnalyticsInstanceResourceGroupResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_analytics.UpdateResourceGroupRequest{}

	analyticsInstanceId, analyticsInstanceResourceGroupId, err := parseAnalyticsInstanceResourceGroupCompositeId(s.D.Id())
	if err == nil {
		request.AnalyticsInstanceId = &analyticsInstanceId
		request.AnalyticsInstanceResourceGroupId = &analyticsInstanceResourceGroupId
	} else {
		log.Printf("[WARN] UpdateWithContext() unable to parse current ID: %s", s.D.Id())
		if analyticsInstanceIdFromConfig, ok := s.D.GetOkExists("analytics_instance_id"); ok {
			tmp := analyticsInstanceIdFromConfig.(string)
			request.AnalyticsInstanceId = &tmp
		}

		tmp := s.D.Id()
		request.AnalyticsInstanceResourceGroupId = &tmp
	}

	if capacity, ok := s.D.GetOkExists("capacity"); ok {
		tmp := capacity.(int)
		request.Capacity = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.UpdateResourceGroup(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsInstanceResourceGroupFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultResourceGroupUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AnalyticsAnalyticsInstanceResourceGroupResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_analytics.DeleteResourceGroupRequest{}

	analyticsInstanceId, analyticsInstanceResourceGroupId, err := parseAnalyticsInstanceResourceGroupCompositeId(s.D.Id())
	if err == nil {
		request.AnalyticsInstanceId = &analyticsInstanceId
		request.AnalyticsInstanceResourceGroupId = &analyticsInstanceResourceGroupId
	} else {
		log.Printf("[WARN] DeleteWithContext() unable to parse current ID: %s", s.D.Id())
		if analyticsInstanceIdFromConfig, ok := s.D.GetOkExists("analytics_instance_id"); ok {
			tmp := analyticsInstanceIdFromConfig.(string)
			request.AnalyticsInstanceId = &tmp
		}

		tmp := s.D.Id()
		request.AnalyticsInstanceResourceGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.DeleteResourceGroup(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := analyticsInstanceResourceGroupWaitForWorkRequest(ctx, workId, "resourcegroup",
		oci_analytics.WorkRequestActionResultResourceGroupDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AnalyticsAnalyticsInstanceResourceGroupResourceCrud) SetData() error {

	analyticsInstanceId, _, err := parseAnalyticsInstanceResourceGroupCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("analytics_instance_id", analyticsInstanceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Capacity != nil {
		s.D.Set("capacity", *s.Res.Capacity)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	return nil
}

func GetAnalyticsInstanceResourceGroupCompositeId(analyticsInstanceId string, analyticsInstanceResourceGroupId string) string {
	analyticsInstanceId = url.PathEscape(analyticsInstanceId)
	analyticsInstanceResourceGroupId = url.PathEscape(analyticsInstanceResourceGroupId)
	compositeId := "analyticsInstances/" + analyticsInstanceId + "/resourceGroups/" + analyticsInstanceResourceGroupId
	return compositeId
}

func parseAnalyticsInstanceResourceGroupCompositeId(compositeId string) (analyticsInstanceId string, analyticsInstanceResourceGroupId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("analyticsInstances/.*/resourceGroups/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	analyticsInstanceId, _ = url.PathUnescape(parts[1])
	analyticsInstanceResourceGroupId, _ = url.PathUnescape(parts[3])

	return
}
