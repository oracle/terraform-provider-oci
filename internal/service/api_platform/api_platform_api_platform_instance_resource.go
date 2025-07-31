// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package api_platform

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_api_platform "github.com/oracle/oci-go-sdk/v65/apiplatform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApiPlatformApiPlatformInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
		},
		Create: createApiPlatformApiPlatformInstance,
		Read:   readApiPlatformApiPlatformInstance,
		Update: updateApiPlatformApiPlatformInstance,
		Delete: deleteApiPlatformApiPlatformInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
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

			// Computed
			"idcs_app": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"url": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uris": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"developers_portal_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_portal_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createApiPlatformApiPlatformInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ApiPlatformApiPlatformInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiPlatformClient()

	return tfresource.CreateResource(d, sync)
}

func readApiPlatformApiPlatformInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ApiPlatformApiPlatformInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiPlatformClient()

	return tfresource.ReadResource(sync)
}

func updateApiPlatformApiPlatformInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ApiPlatformApiPlatformInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiPlatformClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApiPlatformApiPlatformInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ApiPlatformApiPlatformInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiPlatformClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ApiPlatformApiPlatformInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_api_platform.ApiPlatformClient
	Res                    *oci_api_platform.ApiPlatformInstance
	DisableNotFoundRetries bool
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_api_platform.ApiPlatformInstanceLifecycleStateCreating),
	}
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_api_platform.ApiPlatformInstanceLifecycleStateActive),
	}
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_api_platform.ApiPlatformInstanceLifecycleStateDeleting),
	}
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_api_platform.ApiPlatformInstanceLifecycleStateDeleted),
	}
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) Create() error {
	request := oci_api_platform.CreateApiPlatformInstanceRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "api_platform")

	response, err := s.Client.CreateApiPlatformInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getApiPlatformInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "api_platform"), oci_api_platform.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) getApiPlatformInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_api_platform.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	apiPlatformInstanceId, err := apiPlatformInstanceWaitForWorkRequest(workId, "apiplatforminstance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*apiPlatformInstanceId)

	return s.Get()
}

func apiPlatformInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "api_platform", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_api_platform.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func apiPlatformInstanceWaitForWorkRequest(wId *string, entityType string, action oci_api_platform.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_api_platform.ApiPlatformClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "api_platform")
	retryPolicy.ShouldRetryOperation = apiPlatformInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_api_platform.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_api_platform.OperationStatusInProgress),
			string(oci_api_platform.OperationStatusAccepted),
			string(oci_api_platform.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_api_platform.OperationStatusSucceeded),
			string(oci_api_platform.OperationStatusFailed),
			string(oci_api_platform.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_api_platform.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_api_platform.OperationStatusFailed || response.Status == oci_api_platform.OperationStatusCanceled {
		return nil, getErrorFromApiPlatformApiPlatformInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromApiPlatformApiPlatformInstanceWorkRequest(client *oci_api_platform.ApiPlatformClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_api_platform.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_api_platform.ListWorkRequestErrorsRequest{
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

func (s *ApiPlatformApiPlatformInstanceResourceCrud) Get() error {
	request := oci_api_platform.GetApiPlatformInstanceRequest{}

	tmp := s.D.Id()
	request.ApiPlatformInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "api_platform")

	response, err := s.Client.GetApiPlatformInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApiPlatformInstance
	return nil
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_api_platform.UpdateApiPlatformInstanceRequest{}

	tmp := s.D.Id()
	request.ApiPlatformInstanceId = &tmp

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "api_platform")

	response, err := s.Client.UpdateApiPlatformInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApiPlatformInstance
	return nil
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) Delete() error {
	request := oci_api_platform.DeleteApiPlatformInstanceRequest{}

	tmp := s.D.Id()
	request.ApiPlatformInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "api_platform")

	response, err := s.Client.DeleteApiPlatformInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := apiPlatformInstanceWaitForWorkRequest(workId, "apiplatforminstance",
		oci_api_platform.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdcsApp != nil {
		s.D.Set("idcs_app", []interface{}{IdcsAppToMap(s.Res.IdcsApp)})
	} else {
		s.D.Set("idcs_app", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
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

	if s.Res.Uris != nil {
		s.D.Set("uris", []interface{}{UrisToMap(s.Res.Uris)})
	} else {
		s.D.Set("uris", nil)
	}

	return nil
}

func ApiPlatformInstanceSummaryToMap(obj oci_api_platform.ApiPlatformInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
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

func IdcsAppToMap(obj *oci_api_platform.IdcsApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func UrisToMap(obj *oci_api_platform.Uris) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DevelopersPortalUri != nil {
		result["developers_portal_uri"] = string(*obj.DevelopersPortalUri)
	}

	if obj.ManagementPortalUri != nil {
		result["management_portal_uri"] = string(*obj.ManagementPortalUri)
	}

	return result
}

func (s *ApiPlatformApiPlatformInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_api_platform.ChangeApiPlatformInstanceCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ApiPlatformInstanceId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "api_platform")

	_, err := s.Client.ChangeApiPlatformInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
