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

func FusionAppsFusionEnvironmentDataMaskingActivityResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFusionAppsFusionEnvironmentDataMaskingActivity,
		Read:     readFusionAppsFusionEnvironmentDataMaskingActivity,
		Delete:   deleteFusionAppsFusionEnvironmentDataMaskingActivity,
		Schema: map[string]*schema.Schema{
			// Required
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"is_resume_data_masking": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_masking_finish": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_masking_start": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFusionAppsFusionEnvironmentDataMaskingActivity(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.CreateResource(d, sync)
}

func readFusionAppsFusionEnvironmentDataMaskingActivity(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

func deleteFusionAppsFusionEnvironmentDataMaskingActivity(d *schema.ResourceData, m interface{}) error {
	return nil
}

type FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fusion_apps.FusionApplicationsClient
	Res                    *oci_fusion_apps.DataMaskingActivity
	DisableNotFoundRetries bool
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud) ID() string {
	return GetFusionEnvironmentDataMaskingActivityCompositeId(s.D.Get("data_masking_activity_id").(string), s.D.Get("fusion_environment_id").(string))
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fusion_apps.DataMaskingActivityLifecycleStateInProgress),
	}
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fusion_apps.DataMaskingActivityLifecycleStateSucceeded),
	}
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud) Create() error {
	request := oci_fusion_apps.CreateDataMaskingActivityRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if isResumeDataMasking, ok := s.D.GetOkExists("is_resume_data_masking"); ok {
		tmp := isResumeDataMasking.(bool)
		request.IsResumeDataMasking = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.CreateDataMaskingActivity(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFusionEnvironmentDataMaskingActivityFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"), oci_fusion_apps.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud) getFusionEnvironmentDataMaskingActivityFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fusion_apps.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fusionEnvironmentDataMaskingActivityId, err := fusionEnvironmentDataMaskingActivityWaitForWorkRequest(workId, "fusion_apps",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*fusionEnvironmentDataMaskingActivityId)

	return s.Get()
}

func fusionEnvironmentDataMaskingActivityWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fusion_apps.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func fusionEnvironmentDataMaskingActivityWaitForWorkRequest(wId *string, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fusion_apps.FusionApplicationsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fusion_apps")
	retryPolicy.ShouldRetryOperation = fusionEnvironmentDataMaskingActivityWorkRequestShouldRetryFunc(timeout)

	response := oci_fusion_apps.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fusion_apps.WorkRequestStatusInProgress),
			string(oci_fusion_apps.WorkRequestStatusAccepted),
			string(oci_fusion_apps.WorkRequestStatusCanceling),
		},
		Target: []string{
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
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_fusion_apps.WorkRequestStatusFailed || response.Status == oci_fusion_apps.WorkRequestStatusCanceled {
		return nil, getErrorFromFusionAppsFusionEnvironmentDataMaskingActivityWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFusionAppsFusionEnvironmentDataMaskingActivityWorkRequest(client *oci_fusion_apps.FusionApplicationsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum) error {
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

func (s *FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud) Get() error {
	request := oci_fusion_apps.GetDataMaskingActivityRequest{}

	if dataMaskingActivityId, ok := s.D.GetOkExists("data_masking_activity_id"); ok {
		tmp := dataMaskingActivityId.(string)
		request.DataMaskingActivityId = &tmp
	}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	dataMaskingActivityId, fusionEnvironmentId, err := parseFusionEnvironmentDataMaskingActivityCompositeId(s.D.Id())
	if err == nil {
		request.DataMaskingActivityId = &dataMaskingActivityId
		request.FusionEnvironmentId = &fusionEnvironmentId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.GetDataMaskingActivity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataMaskingActivity
	return nil
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivityResourceCrud) SetData() error {

	dataMaskingActivityId, fusionEnvironmentId, err := parseFusionEnvironmentDataMaskingActivityCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("data_masking_activity_id", &dataMaskingActivityId)
		s.D.Set("fusion_environment_id", &fusionEnvironmentId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.FusionEnvironmentId != nil {
		s.D.Set("fusion_environment_id", *s.Res.FusionEnvironmentId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeMaskingFinish != nil {
		s.D.Set("time_masking_finish", s.Res.TimeMaskingFinish.String())
	}

	if s.Res.TimeMaskingStart != nil {
		s.D.Set("time_masking_start", s.Res.TimeMaskingStart.String())
	}

	return nil
}

func GetFusionEnvironmentDataMaskingActivityCompositeId(dataMaskingActivityId string, fusionEnvironmentId string) string {
	dataMaskingActivityId = url.PathEscape(dataMaskingActivityId)
	fusionEnvironmentId = url.PathEscape(fusionEnvironmentId)
	compositeId := "fusionEnvironments/" + fusionEnvironmentId + "/dataMaskingActivities/" + dataMaskingActivityId
	return compositeId
}

func parseFusionEnvironmentDataMaskingActivityCompositeId(compositeId string) (dataMaskingActivityId string, fusionEnvironmentId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("fusionEnvironments/.*/dataMaskingActivities/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	fusionEnvironmentId, _ = url.PathUnescape(parts[1])
	dataMaskingActivityId, _ = url.PathUnescape(parts[3])

	return
}

func DataMaskingActivitySummaryToMap(obj oci_fusion_apps.DataMaskingActivitySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeMaskingFinish != nil {
		result["time_masking_finish"] = obj.TimeMaskingFinish.String()
	}

	if obj.TimeMaskingStart != nil {
		result["time_masking_start"] = obj.TimeMaskingStart.String()
	}

	return result
}
