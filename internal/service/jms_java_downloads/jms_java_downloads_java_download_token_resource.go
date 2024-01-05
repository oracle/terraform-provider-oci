// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaDownloadTokenResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createJmsJavaDownloadsJavaDownloadToken,
		Read:     readJmsJavaDownloadsJavaDownloadToken,
		Update:   updateJmsJavaDownloadsJavaDownloadToken,
		Delete:   deleteJmsJavaDownloadsJavaDownloadToken,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"java_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"license_type": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_expires": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"last_updated_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"time_last_used": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createJmsJavaDownloadsJavaDownloadToken(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.CreateResource(d, sync)
}

func readJmsJavaDownloadsJavaDownloadToken(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

func updateJmsJavaDownloadsJavaDownloadToken(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteJmsJavaDownloadsJavaDownloadToken(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type JmsJavaDownloadsJavaDownloadTokenResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_jms_java_downloads.JavaDownloadClient
	Res                    *oci_jms_java_downloads.JavaDownloadToken
	DisableNotFoundRetries bool
}

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateCreating),
	}
}

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateActive),
		string(oci_jms_java_downloads.LifecycleStateNeedsAttention),
	}
}

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateDeleting),
	}
}

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_jms_java_downloads.LifecycleStateDeleted),
	}
}

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) Create() error {
	request := oci_jms_java_downloads.CreateJavaDownloadTokenRequest{}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isDefault, ok := s.D.GetOkExists("is_default"); ok {
		tmp := isDefault.(bool)
		request.IsDefault = &tmp
	}

	if javaVersion, ok := s.D.GetOkExists("java_version"); ok {
		tmp := javaVersion.(string)
		request.JavaVersion = &tmp
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		interfaces := licenseType.([]interface{})
		tmp := make([]oci_jms_java_downloads.LicenseTypeEnum, len(interfaces))
		for i, v := range interfaces {
			if value, ok := v.(string); ok {
				converted, ok := oci_jms_java_downloads.GetMappingLicenseTypeEnum(value)
				if ok {
					tmp[i] = converted
				}
			}
		}
		if len(tmp) != 0 || s.D.HasChange("license_type") {
			request.LicenseType = tmp
		}
	}

	if timeExpires, ok := s.D.GetOkExists("time_expires"); ok {
		tmp, err := time.Parse(time.RFC3339, timeExpires.(string))
		if err != nil {
			return err
		}
		request.TimeExpires = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	response, err := s.Client.CreateJavaDownloadToken(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_jms_java_downloads.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_jms_java_downloads.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "jmsjavadownloadtoken") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getJavaDownloadTokenFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads"), oci_jms_java_downloads.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) getJavaDownloadTokenFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_jms_java_downloads.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	javaDownloadTokenId, err := javaDownloadTokenWaitForWorkRequest(workId, "jmsjavadownloadtoken",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*javaDownloadTokenId)

	return s.Get()
}

func javaDownloadTokenWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "jms_java_downloads", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_jms_java_downloads.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func javaDownloadTokenWaitForWorkRequest(wId *string, entityType string, action oci_jms_java_downloads.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_jms_java_downloads.JavaDownloadClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "jms_java_downloads")
	retryPolicy.ShouldRetryOperation = javaDownloadTokenWorkRequestShouldRetryFunc(timeout)

	response := oci_jms_java_downloads.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_jms_java_downloads.ListWorkRequestsStatusInProgress),
			string(oci_jms_java_downloads.ListWorkRequestsStatusAccepted),
			string(oci_jms_java_downloads.ListWorkRequestsStatusCanceling),
		},
		Target: []string{
			string(oci_jms_java_downloads.ListWorkRequestsStatusSucceeded),
			string(oci_jms_java_downloads.ListWorkRequestsStatusFailed),
			string(oci_jms_java_downloads.ListWorkRequestsStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_jms_java_downloads.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_jms_java_downloads.OperationStatusEnum(oci_jms_java_downloads.ListWorkRequestsStatusFailed) || response.Status == oci_jms_java_downloads.OperationStatusEnum(oci_jms_java_downloads.ListWorkRequestsStatusCanceled) {
		return nil, getErrorFromJmsJavaDownloadsJavaDownloadTokenWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromJmsJavaDownloadsJavaDownloadTokenWorkRequest(client *oci_jms_java_downloads.JavaDownloadClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_jms_java_downloads.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_jms_java_downloads.ListWorkRequestErrorsRequest{
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

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) Get() error {
	request := oci_jms_java_downloads.GetJavaDownloadTokenRequest{}

	tmp := s.D.Id()
	request.JavaDownloadTokenId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	response, err := s.Client.GetJavaDownloadToken(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JavaDownloadToken
	return nil
}

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) Update() error {
	request := oci_jms_java_downloads.UpdateJavaDownloadTokenRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isDefault, ok := s.D.GetOkExists("is_default"); ok {
		tmp := isDefault.(bool)
		request.IsDefault = &tmp
	}

	tmp := s.D.Id()
	request.JavaDownloadTokenId = &tmp

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		interfaces := licenseType.([]interface{})
		tmp := make([]oci_jms_java_downloads.LicenseTypeEnum, len(interfaces))
		for i, v := range interfaces {
			if value, ok := v.(string); ok {
				converted, ok := oci_jms_java_downloads.GetMappingLicenseTypeEnum(value)
				if ok {
					tmp[i] = converted
				}
			}
		}
		if len(tmp) != 0 || s.D.HasChange("license_type") {
			request.LicenseType = tmp
		}
	}

	if timeExpires, ok := s.D.GetOkExists("time_expires"); ok {
		tmp, err := time.Parse(time.RFC3339, timeExpires.(string))
		if err != nil {
			return err
		}
		request.TimeExpires = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	response, err := s.Client.UpdateJavaDownloadToken(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getJavaDownloadTokenFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads"), oci_jms_java_downloads.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) Delete() error {
	request := oci_jms_java_downloads.DeleteJavaDownloadTokenRequest{}

	tmp := s.D.Id()
	request.JavaDownloadTokenId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_java_downloads")

	response, err := s.Client.DeleteJavaDownloadToken(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := javaDownloadTokenWaitForWorkRequest(workId, "jmsjavadownloadtoken",
		oci_jms_java_downloads.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *JmsJavaDownloadsJavaDownloadTokenResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", []interface{}{PrincipalToMap(s.Res.CreatedBy)})
	} else {
		s.D.Set("created_by", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	if s.Res.JavaVersion != nil {
		s.D.Set("java_version", *s.Res.JavaVersion)
	}

	if s.Res.LastUpdatedBy != nil {
		s.D.Set("last_updated_by", []interface{}{PrincipalToMap(s.Res.LastUpdatedBy)})
	} else {
		s.D.Set("last_updated_by", nil)
	}

	licenseType := []interface{}{}
	for _, item := range s.Res.LicenseType {
		licenseType = append(licenseType, item)
	}
	s.D.Set("license_type", licenseType)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires.Format(time.RFC3339Nano))
	}

	if s.Res.TimeLastUsed != nil {
		s.D.Set("time_last_used", s.Res.TimeLastUsed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Value != nil {
		s.D.Set("value", *s.Res.Value)
	}

	return nil
}

func JavaDownloadTokenSummaryToMap(obj oci_jms_java_downloads.JavaDownloadTokenSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = []interface{}{PrincipalToMap(obj.CreatedBy)}
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsDefault != nil {
		result["is_default"] = bool(*obj.IsDefault)
	}

	if obj.JavaVersion != nil {
		result["java_version"] = string(*obj.JavaVersion)
	}

	if obj.LastUpdatedBy != nil {
		result["last_updated_by"] = []interface{}{PrincipalToMap(obj.LastUpdatedBy)}
	}

	licenseType := []interface{}{}
	for _, item := range obj.LicenseType {
		licenseType = append(licenseType, item)
	}
	result["license_type"] = licenseType

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeExpires != nil {
		result["time_expires"] = obj.TimeExpires.String()
	}

	if obj.TimeLastUsed != nil {
		result["time_last_used"] = obj.TimeLastUsed.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func PrincipalToMap(obj *oci_jms_java_downloads.Principal) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Email != nil {
		result["email"] = string(*obj.Email)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}
