// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentRefreshActivityResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createFusionAppsFusionEnvironmentRefreshActivity,
		Read:   readFusionAppsFusionEnvironmentRefreshActivity,
		Delete: deleteFusionAppsFusionEnvironmentRefreshActivity,
		Schema: map[string]*schema.Schema{
			// Required
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Optional
			"is_data_masking_opted": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			// Computed
			"refresh_activity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"refresh_issue_details_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"refresh_issues": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"service_availability": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_accepted": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_expected_finish": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_restoration_point": {
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

func createFusionAppsFusionEnvironmentRefreshActivity(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentRefreshActivityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.CreateResource(d, sync)
}

func readFusionAppsFusionEnvironmentRefreshActivity(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentRefreshActivityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

func deleteFusionAppsFusionEnvironmentRefreshActivity(d *schema.ResourceData, m interface{}) error {
	return nil
}

type FusionAppsFusionEnvironmentRefreshActivityResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fusion_apps.FusionApplicationsClient
	Res                    *oci_fusion_apps.RefreshActivity
	DisableNotFoundRetries bool
}

func (s *FusionAppsFusionEnvironmentRefreshActivityResourceCrud) ID() string {
	return GetFusionEnvironmentRefreshActivityCompositeId(s.D.Get("fusion_environment_id").(string), *(s.Res).Id)
}

func (s *FusionAppsFusionEnvironmentRefreshActivityResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fusion_apps.RefreshActivityLifecycleStateInProgress),
	}
}

func (s *FusionAppsFusionEnvironmentRefreshActivityResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fusion_apps.RefreshActivityLifecycleStateAccepted),
	}
}

func (s *FusionAppsFusionEnvironmentRefreshActivityResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *FusionAppsFusionEnvironmentRefreshActivityResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *FusionAppsFusionEnvironmentRefreshActivityResourceCrud) Create() error {
	request := oci_fusion_apps.CreateRefreshActivityRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if isDataMaskingOpted, ok := s.D.GetOkExists("is_data_masking_opted"); ok {
		tmp := isDataMaskingOpted.(bool)
		request.IsDataMaskingOpted = &tmp
	}

	if sourceFusionEnvironmentId, ok := s.D.GetOkExists("source_fusion_environment_id"); ok {
		tmp := sourceFusionEnvironmentId.(string)
		request.SourceFusionEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.CreateRefreshActivity(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.setIdFromWorkRequest(workId)
	return s.getFusionEnvironmentRefreshActivityFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"), oci_fusion_apps.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FusionAppsFusionEnvironmentRefreshActivityResourceCrud) setIdFromWorkRequest(workId *string) {
	var identifier *string
	var err error

	workRequestResponse := oci_fusion_apps.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_fusion_apps.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "refreshactivity") {
				identifier = res.Identifier
				break
			}
		}
	}
	if identifier != nil {
		compositeId := GetFusionEnvironmentRefreshActivityCompositeId(s.D.Get("fusion_environment_id").(string), *identifier)
		s.D.SetId(compositeId)
	}
}

func (s *FusionAppsFusionEnvironmentRefreshActivityResourceCrud) getFusionEnvironmentRefreshActivityFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fusion_apps.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// We don't wait for the actual refresh complete,
	// the work request will track the actual refresh instead of the refresh resource creation
	// As long as the work request is returned, that means the refresh resource is created
	fusionEnvironmentRefreshActivityId, err := fusionEnvironmentRefreshActivityWaitForWorkRequest(workId, "refreshactivity",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.Set("refresh_activity_id", fusionEnvironmentRefreshActivityId)

	return s.Get()
}

func fusionEnvironmentRefreshActivityWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fusion_apps", startTime) {
			return true
		}

		// This work request will track the actual refresh, for creation,
		//when the work request is accepted then the refresh creation is succeeded
		if workRequestResponse, ok := response.Response.(oci_fusion_apps.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeAccepted == nil
		}
		return false
	}
}

func fusionEnvironmentRefreshActivityWaitForWorkRequest(wId *string, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fusion_apps.FusionApplicationsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fusion_apps")
	retryPolicy.ShouldRetryOperation = fusionEnvironmentRefreshActivityWorkRequestShouldRetryFunc(timeout)

	response := oci_fusion_apps.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fusion_apps.WorkRequestStatusCanceling),
		},
		// We don't wait for the actual refresh complete,
		// the work request will track the actual refresh instead of the refresh resource creation
		// As long as the work request is returned, that means the refresh resource is created
		Target: []string{
			string(oci_fusion_apps.WorkRequestStatusInProgress),
			string(oci_fusion_apps.WorkRequestStatusAccepted),
			string(oci_fusion_apps.WorkRequestStatusSucceeded),
			string(oci_fusion_apps.WorkRequestStatusFailed),
			string(oci_fusion_apps.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_fusion_apps.GetWorkRequestRequest{
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
			identifier = res.Identifier
			break
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_fusion_apps.WorkRequestStatusFailed || response.Status == oci_fusion_apps.WorkRequestStatusCanceled {
		return nil, getErrorFromFusionAppsFusionEnvironmentRefreshActivityWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFusionAppsFusionEnvironmentRefreshActivityWorkRequest(client *oci_fusion_apps.FusionApplicationsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fusion_apps.ListWorkRequestErrorsRequest{
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

func (s *FusionAppsFusionEnvironmentRefreshActivityResourceCrud) Get() error {
	request := oci_fusion_apps.GetRefreshActivityRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if refreshActivityId, ok := s.D.GetOkExists("refresh_activity_id"); ok {
		tmp := refreshActivityId.(string)
		request.RefreshActivityId = &tmp
	}

	fusionEnvironmentId, refreshActivityId, err := parseFusionEnvironmentRefreshActivityCompositeId(s.D.Id())
	if err == nil {
		request.FusionEnvironmentId = &fusionEnvironmentId
		request.RefreshActivityId = &refreshActivityId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.GetRefreshActivity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RefreshActivity
	return nil
}

func (s *FusionAppsFusionEnvironmentRefreshActivityResourceCrud) SetData() error {

	fusionEnvironmentId, refreshActivityId, err := parseFusionEnvironmentRefreshActivityCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("fusion_environment_id", &fusionEnvironmentId)
		s.D.Set("refresh_activity_id", &refreshActivityId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsDataMaskingOpted != nil {
		s.D.Set("is_data_masking_opted", *s.Res.IsDataMaskingOpted)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	refreshIssueDetailsList := []interface{}{}
	for _, item := range s.Res.RefreshIssueDetailsList {
		refreshIssueDetailsList = append(refreshIssueDetailsList, RefreshIssueDetailsToMap(item))
	}
	s.D.Set("refresh_issue_details_list", refreshIssueDetailsList)

	s.D.Set("service_availability", s.Res.ServiceAvailability)

	if s.Res.SourceFusionEnvironmentId != nil {
		s.D.Set("source_fusion_environment_id", *s.Res.SourceFusionEnvironmentId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeExpectedFinish != nil {
		s.D.Set("time_expected_finish", s.Res.TimeExpectedFinish.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeOfRestorationPoint != nil {
		s.D.Set("time_of_restoration_point", s.Res.TimeOfRestorationPoint.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetFusionEnvironmentRefreshActivityCompositeId(fusionEnvironmentId string, refreshActivityId string) string {
	fusionEnvironmentId = url.PathEscape(fusionEnvironmentId)
	refreshActivityId = url.PathEscape(refreshActivityId)
	compositeId := "fusionEnvironments/" + fusionEnvironmentId + "/refreshActivities/" + refreshActivityId
	return compositeId
}

func parseFusionEnvironmentRefreshActivityCompositeId(compositeId string) (fusionEnvironmentId string, refreshActivityId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("fusionEnvironments/.*/refreshActivities/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	fusionEnvironmentId, _ = url.PathUnescape(parts[1])
	refreshActivityId, _ = url.PathUnescape(parts[3])

	return
}

func RefreshActivitySummaryToMap(obj oci_fusion_apps.RefreshActivitySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsDataMaskingOpted != nil {
		result["is_data_masking_opted"] = bool(*obj.IsDataMaskingOpted)
	}

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	refreshIssueDetailsList := []interface{}{}
	for _, item := range obj.RefreshIssueDetailsList {
		refreshIssueDetailsList = append(refreshIssueDetailsList, RefreshIssueDetailsToMap(item))
	}
	result["refresh_issue_details_list"] = refreshIssueDetailsList

	result["service_availability"] = string(obj.ServiceAvailability)

	if obj.SourceFusionEnvironmentId != nil {
		result["source_fusion_environment_id"] = string(*obj.SourceFusionEnvironmentId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeAccepted != nil {
		result["time_accepted"] = obj.TimeAccepted.String()
	}

	if obj.TimeExpectedFinish != nil {
		result["time_expected_finish"] = obj.TimeExpectedFinish.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeOfRestorationPoint != nil {
		result["time_of_restoration_point"] = obj.TimeOfRestorationPoint.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func RefreshIssueDetailsToMap(obj oci_fusion_apps.RefreshIssueDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RefreshIssues != nil {
		result["refresh_issues"] = string(*obj.RefreshIssues)
	}

	return result
}
