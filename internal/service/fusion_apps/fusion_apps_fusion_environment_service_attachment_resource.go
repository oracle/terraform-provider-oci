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

func FusionAppsFusionEnvironmentServiceAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFusionAppsFusionEnvironmentServiceAttachment,
		Read:     readFusionAppsFusionEnvironmentServiceAttachment,
		Delete:   deleteFusionAppsFusionEnvironmentServiceAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_instance_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Optional: true,
				Elem:     schema.TypeString,
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
			"is_sku_based": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"service_url": {
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

func createFusionAppsFusionEnvironmentServiceAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentServiceAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.CreateResource(d, sync)
}

func readFusionAppsFusionEnvironmentServiceAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentServiceAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

func deleteFusionAppsFusionEnvironmentServiceAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentServiceAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FusionAppsFusionEnvironmentServiceAttachmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fusion_apps.FusionApplicationsClient
	Res                    *oci_fusion_apps.ServiceAttachment
	DisableNotFoundRetries bool
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) ID() string {
	return GetFusionEnvironmentServiceAttachmentCompositeId(s.D.Get("fusion_environment_id").(string), *(s.Res).Id)
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fusion_apps.ServiceAttachmentLifecycleStateCreating),
	}
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fusion_apps.ServiceAttachmentLifecycleStateActive),
	}
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fusion_apps.ServiceAttachmentLifecycleStateDeleting),
	}
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fusion_apps.ServiceAttachmentLifecycleStateDeleted),
	}
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) Create() error {
	request := oci_fusion_apps.CreateServiceAttachmentRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if serviceInstanceId, ok := s.D.GetOkExists("service_instance_id"); ok {
		tmp := serviceInstanceId.(string)
		request.ServiceInstanceId = &tmp
	}

	if serviceInstanceType, ok := s.D.GetOkExists("service_instance_type"); ok {
		request.ServiceInstanceType = oci_fusion_apps.ServiceAttachmentServiceInstanceTypeEnum(serviceInstanceType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.CreateServiceAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.setIdFromWorkRequest(workId)
	return s.getFusionEnvironmentServiceAttachmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"), oci_fusion_apps.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) setIdFromWorkRequest(workId *string) {
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
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "serviceattachments") {
				identifier = res.Identifier
				break
			}
		}
	}
	if identifier != nil {
		compositeId := GetFusionEnvironmentServiceAttachmentCompositeId(s.D.Get("fusion_environment_id").(string), *identifier)
		s.D.SetId(compositeId)
	}
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) getFusionEnvironmentServiceAttachmentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fusion_apps.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fusionEnvironmentServiceAttachmentId, err := fusionEnvironmentServiceAttachmentWaitForWorkRequest(workId,
		"serviceattachments",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	compositeId := *fusionEnvironmentServiceAttachmentId
	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		compositeId = GetFusionEnvironmentServiceAttachmentCompositeId(tmp, *fusionEnvironmentServiceAttachmentId)
	} else {
		log.Printf("[WARN] Unable to set composite id")
	}

	s.D.SetId(compositeId)

	return s.Get()
}

func fusionEnvironmentServiceAttachmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func fusionEnvironmentServiceAttachmentWaitForWorkRequest(wId *string, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fusion_apps.FusionApplicationsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fusion_apps")
	retryPolicy.ShouldRetryOperation = fusionEnvironmentServiceAttachmentWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromFusionAppsFusionEnvironmentServiceAttachmentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFusionAppsFusionEnvironmentServiceAttachmentWorkRequest(client *oci_fusion_apps.FusionApplicationsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum) error {
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

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) Get() error {
	request := oci_fusion_apps.GetServiceAttachmentRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if serviceAttachmentId, ok := s.D.GetOkExists("service_attachment_id"); ok {
		tmp := serviceAttachmentId.(string)
		request.ServiceAttachmentId = &tmp
	}

	tmp := s.D.Id()
	request.ServiceAttachmentId = &tmp
	fusionEnvironmentId, serviceAttachmentId, err := parseFusionEnvironmentServiceAttachmentCompositeId(s.D.Id())
	if err == nil {
		request.FusionEnvironmentId = &fusionEnvironmentId
		request.ServiceAttachmentId = &serviceAttachmentId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.GetServiceAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceAttachment
	return nil
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) Delete() error {
	request := oci_fusion_apps.DeleteServiceAttachmentRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if serviceAttachmentId, ok := s.D.GetOkExists("service_attachment_id"); ok {
		tmp := serviceAttachmentId.(string)
		request.ServiceAttachmentId = &tmp
	}

	tmp := s.D.Id()
	request.ServiceAttachmentId = &tmp
	fusionEnvironmentId, serviceAttachmentId, err := parseFusionEnvironmentServiceAttachmentCompositeId(s.D.Id())
	if err == nil {
		request.FusionEnvironmentId = &fusionEnvironmentId
		request.ServiceAttachmentId = &serviceAttachmentId
	} else {
		log.Printf("[WARN] DELETE() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.DeleteServiceAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := fusionEnvironmentServiceAttachmentWaitForWorkRequest(workId, "serviceattachments",
		oci_fusion_apps.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries,
		s.Client)

	return delWorkRequestErr
}

func (s *FusionAppsFusionEnvironmentServiceAttachmentResourceCrud) SetData() error {
	fusionEnvironmentId, _, err := parseFusionEnvironmentServiceAttachmentCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("fusion_environment_id", &fusionEnvironmentId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSkuBased != nil {
		s.D.Set("is_sku_based", *s.Res.IsSkuBased)
	}

	if s.Res.ServiceInstanceId != nil {
		s.D.Set("service_instance_id", *s.Res.ServiceInstanceId)
	}

	s.D.Set("service_instance_type", s.Res.ServiceInstanceType)

	if s.Res.ServiceUrl != nil {
		s.D.Set("service_url", *s.Res.ServiceUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetFusionEnvironmentServiceAttachmentCompositeId(fusionEnvironmentId string, serviceAttachmentId string) string {
	fusionEnvironmentId = url.PathEscape(fusionEnvironmentId)
	serviceAttachmentId = url.PathEscape(serviceAttachmentId)
	compositeId := "fusionEnvironments/" + fusionEnvironmentId + "/serviceAttachments/" + serviceAttachmentId
	return compositeId
}

func parseFusionEnvironmentServiceAttachmentCompositeId(compositeId string) (fusionEnvironmentId string, serviceAttachmentId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("fusionEnvironments/.*/serviceAttachments/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	fusionEnvironmentId, _ = url.PathUnescape(parts[1])
	serviceAttachmentId, _ = url.PathUnescape(parts[3])

	return
}

func ServiceAttachmentSummaryToMap(obj oci_fusion_apps.ServiceAttachmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsSkuBased != nil {
		result["is_sku_based"] = bool(*obj.IsSkuBased)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ServiceInstanceId != nil {
		result["service_instance_id"] = string(*obj.ServiceInstanceId)
	}

	result["service_instance_type"] = string(obj.ServiceInstanceType)

	if obj.ServiceUrl != nil {
		result["service_url"] = string(*obj.ServiceUrl)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
