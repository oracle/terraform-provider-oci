// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waa

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_waa "github.com/oracle/oci-go-sdk/v65/waa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WaaWebAppAccelerationPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createWaaWebAppAccelerationPolicy,
		Read:     readWaaWebAppAccelerationPolicy,
		Update:   updateWaaWebAppAccelerationPolicy,
		Delete:   deleteWaaWebAppAccelerationPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"response_caching_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_response_header_based_caching_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"response_compression_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"gzip_compression": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"is_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"lifecycle_details": {
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

func createWaaWebAppAccelerationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WaaWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readWaaWebAppAccelerationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()

	return tfresource.ReadResource(sync)
}

func updateWaaWebAppAccelerationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WaaWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteWaaWebAppAccelerationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WaaWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type WaaWebAppAccelerationPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_waa.WaaClient
	Res                    *oci_waa.WebAppAccelerationPolicy
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_waa.WorkRequestClient
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_waa.WebAppAccelerationPolicyLifecycleStateCreating),
	}
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_waa.WebAppAccelerationPolicyLifecycleStateActive),
	}
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_waa.WebAppAccelerationPolicyLifecycleStateDeleting),
	}
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_waa.WebAppAccelerationPolicyLifecycleStateDeleted),
	}
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) Create() error {
	request := oci_waa.CreateWebAppAccelerationPolicyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if responseCachingPolicy, ok := s.D.GetOkExists("response_caching_policy"); ok {
		if tmpList := responseCachingPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "response_caching_policy", 0)
			tmp, err := s.mapToResponseCachingPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResponseCachingPolicy = &tmp
		}
	}

	if responseCompressionPolicy, ok := s.D.GetOkExists("response_compression_policy"); ok {
		if tmpList := responseCompressionPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "response_compression_policy", 0)
			tmp, err := s.mapToResponseCompressionPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResponseCompressionPolicy = &tmp
		}
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa")

	response, err := s.Client.CreateWebAppAccelerationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getWebAppAccelerationPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa"), oci_waa.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) getWebAppAccelerationPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_waa.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	webAppAccelerationPolicyId, err := webAppAccelerationPolicyWaitForWorkRequest(workId, "webappaccelerationpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*webAppAccelerationPolicyId)

	return s.Get()
}

func webAppAccelerationPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "waa", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_waa.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func webAppAccelerationPolicyWaitForWorkRequest(wId *string, entityType string, action oci_waa.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_waa.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "waa")
	retryPolicy.ShouldRetryOperation = webAppAccelerationPolicyWorkRequestShouldRetryFunc(timeout)

	response := oci_waa.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_waa.WorkRequestStatusInProgress),
			string(oci_waa.WorkRequestStatusAccepted),
			string(oci_waa.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_waa.WorkRequestStatusSucceeded),
			string(oci_waa.WorkRequestStatusFailed),
			string(oci_waa.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_waa.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_waa.WorkRequestStatusFailed || response.Status == oci_waa.WorkRequestStatusCanceled {
		return nil, getErrorFromWaaWebAppAccelerationPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromWaaWebAppAccelerationPolicyWorkRequest(client *oci_waa.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_waa.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_waa.ListWorkRequestErrorsRequest{
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

func (s *WaaWebAppAccelerationPolicyResourceCrud) Get() error {
	request := oci_waa.GetWebAppAccelerationPolicyRequest{}

	tmp := s.D.Id()
	request.WebAppAccelerationPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa")

	response, err := s.Client.GetWebAppAccelerationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.WebAppAccelerationPolicy
	return nil
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_waa.UpdateWebAppAccelerationPolicyRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if responseCachingPolicy, ok := s.D.GetOkExists("response_caching_policy"); ok {
		if tmpList := responseCachingPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "response_caching_policy", 0)
			tmp, err := s.mapToResponseCachingPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResponseCachingPolicy = &tmp
		}
	}

	if responseCompressionPolicy, ok := s.D.GetOkExists("response_compression_policy"); ok {
		if tmpList := responseCompressionPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "response_compression_policy", 0)
			tmp, err := s.mapToResponseCompressionPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResponseCompressionPolicy = &tmp
		}
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	tmp := s.D.Id()
	request.WebAppAccelerationPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa")

	response, err := s.Client.UpdateWebAppAccelerationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWebAppAccelerationPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa"), oci_waa.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) Delete() error {
	request := oci_waa.DeleteWebAppAccelerationPolicyRequest{}

	tmp := s.D.Id()
	request.WebAppAccelerationPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa")

	response, err := s.Client.DeleteWebAppAccelerationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := webAppAccelerationPolicyWaitForWorkRequest(workId, "webappaccelerationpolicy",
		oci_waa.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ResponseCachingPolicy != nil {
		s.D.Set("response_caching_policy", []interface{}{ResponseCachingPolicyToMap(s.Res.ResponseCachingPolicy)})
	} else {
		s.D.Set("response_caching_policy", nil)
	}

	if s.Res.ResponseCompressionPolicy != nil {
		s.D.Set("response_compression_policy", []interface{}{ResponseCompressionPolicyToMap(s.Res.ResponseCompressionPolicy)})
	} else {
		s.D.Set("response_compression_policy", nil)
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

func (s *WaaWebAppAccelerationPolicyResourceCrud) mapToGzipCompressionPolicy(fieldKeyFormat string) (oci_waa.GzipCompressionPolicy, error) {
	result := oci_waa.GzipCompressionPolicy{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func GzipCompressionPolicyToMap(obj *oci_waa.GzipCompressionPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) mapToResponseCachingPolicy(fieldKeyFormat string) (oci_waa.ResponseCachingPolicy, error) {
	result := oci_waa.ResponseCachingPolicy{}

	if isResponseHeaderBasedCachingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_response_header_based_caching_enabled")); ok {
		tmp := isResponseHeaderBasedCachingEnabled.(bool)
		result.IsResponseHeaderBasedCachingEnabled = &tmp
	}

	return result, nil
}

func ResponseCachingPolicyToMap(obj *oci_waa.ResponseCachingPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsResponseHeaderBasedCachingEnabled != nil {
		result["is_response_header_based_caching_enabled"] = bool(*obj.IsResponseHeaderBasedCachingEnabled)
	}

	return result
}

func (s *WaaWebAppAccelerationPolicyResourceCrud) mapToResponseCompressionPolicy(fieldKeyFormat string) (oci_waa.ResponseCompressionPolicy, error) {
	result := oci_waa.ResponseCompressionPolicy{}

	if gzipCompression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "gzip_compression")); ok {
		if tmpList := gzipCompression.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "gzip_compression"), 0)
			tmp, err := s.mapToGzipCompressionPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert gzip_compression, encountered error: %v", err)
			}
			result.GzipCompression = &tmp
		}
	}

	return result, nil
}

func ResponseCompressionPolicyToMap(obj *oci_waa.ResponseCompressionPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GzipCompression != nil {
		result["gzip_compression"] = []interface{}{GzipCompressionPolicyToMap(obj.GzipCompression)}
	}

	return result
}

func WebAppAccelerationPolicySummaryToMap(obj oci_waa.WebAppAccelerationPolicySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func (s *WaaWebAppAccelerationPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_waa.ChangeWebAppAccelerationPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.WebAppAccelerationPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa")

	response, err := s.Client.ChangeWebAppAccelerationPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWebAppAccelerationPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa"), oci_waa.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
