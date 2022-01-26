// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_apigateway "github.com/oracle/oci-go-sdk/v56/apigateway"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func ApigatewayApiResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApigatewayApi,
		Read:     readApigatewayApi,
		Update:   updateApigatewayApi,
		Delete:   deleteApigatewayApi,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"content": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"specification_type": {
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
			"validation_results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"result": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createApigatewayApi(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayApiResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiGatewayClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.CreateResource(d, sync)
}

func readApigatewayApi(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayApiResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiGatewayClient()

	return tfresource.ReadResource(sync)
}

func updateApigatewayApi(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayApiResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiGatewayClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApigatewayApi(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayApiResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiGatewayClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.DeleteResource(d, sync)
}

type ApigatewayApiResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apigateway.ApiGatewayClient
	Res                    *oci_apigateway.Api
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_apigateway.WorkRequestsClient
}

func (s *ApigatewayApiResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ApigatewayApiResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_apigateway.ApiLifecycleStateCreating),
	}
}

func (s *ApigatewayApiResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_apigateway.ApiLifecycleStateActive),
	}
}

func (s *ApigatewayApiResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_apigateway.ApiLifecycleStateDeleting),
	}
}

func (s *ApigatewayApiResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_apigateway.ApiLifecycleStateDeleted),
	}
}

func (s *ApigatewayApiResourceCrud) Create() error {
	request := oci_apigateway.CreateApiRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if content, ok := s.D.GetOkExists("content"); ok {
		tmp := content.(string)
		request.Content = &tmp
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.CreateApi(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getApiFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	return apiWaitForValidation(response.Id, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.Client)
}

func (s *ApigatewayApiResourceCrud) getApiFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_apigateway.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	apiId, err := apiWaitForWorkRequest(workId, "api",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, apiId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(context.Background(),
			oci_apigateway.CancelWorkRequestRequest{
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
	s.D.SetId(*apiId)

	return s.Get()
}

func apiWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "apigateway", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_apigateway.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func apiWaitForWorkRequest(wId *string, entityType string, action oci_apigateway.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_apigateway.WorkRequestsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "apigateway")
	retryPolicy.ShouldRetryOperation = apiWorkRequestShouldRetryFunc(timeout)

	response := oci_apigateway.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_apigateway.WorkRequestStatusInProgress),
			string(oci_apigateway.WorkRequestStatusAccepted),
			string(oci_apigateway.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_apigateway.WorkRequestStatusSucceeded),
			string(oci_apigateway.WorkRequestStatusFailed),
			string(oci_apigateway.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_apigateway.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_apigateway.WorkRequestStatusFailed || response.Status == oci_apigateway.WorkRequestStatusCanceled {
		return nil, getErrorFromApigatewayApiWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromApigatewayApiWorkRequest(client *oci_apigateway.WorkRequestsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_apigateway.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_apigateway.ListWorkRequestErrorsRequest{
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

func (s *ApigatewayApiResourceCrud) Get() error {
	request := oci_apigateway.GetApiRequest{}

	tmp := s.D.Id()
	request.ApiId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.GetApi(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Api
	return nil
}

func (s *ApigatewayApiResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_apigateway.UpdateApiRequest{}

	apiId := s.D.Id()
	request.ApiId = &apiId

	if content, ok := s.D.GetOkExists("content"); ok {
		tmp := content.(string)
		request.Content = &tmp
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.UpdateApi(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getApiFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	return apiWaitForValidation(&apiId, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
}

func (s *ApigatewayApiResourceCrud) Delete() error {
	request := oci_apigateway.DeleteApiRequest{}

	tmp := s.D.Id()
	request.ApiId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.DeleteApi(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := apiWaitForWorkRequest(workId, "api",
		oci_apigateway.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *ApigatewayApiResourceCrud) SetData() error {
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

	if s.Res.SpecificationType != nil {
		s.D.Set("specification_type", *s.Res.SpecificationType)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	validationResults := []interface{}{}
	for _, item := range s.Res.ValidationResults {
		validationResults = append(validationResults, ApiValidationResultToMap(item))
	}
	s.D.Set("validation_results", validationResults)

	return nil
}

func ApiSummaryToMap(obj oci_apigateway.ApiSummary) map[string]interface{} {
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

	if obj.SpecificationType != nil {
		result["specification_type"] = string(*obj.SpecificationType)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	validationResults := []interface{}{}
	for _, item := range obj.ValidationResults {
		validationResults = append(validationResults, ApiValidationResultToMap(item))
	}
	result["validation_results"] = validationResults

	return result
}

func ApiValidationResultToMap(obj oci_apigateway.ApiValidationResult) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["result"] = string(obj.Result)

	return result
}

func (s *ApigatewayApiResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_apigateway.ChangeApiCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ApiId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.ChangeApiCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getApiFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func apiResourceShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		if time.Now().After(stopTime) {
			return false
		}

		if tfresource.ShouldRetry(response, false, "apigateway", startTime) {
			return true
		}

		if apiResponse, ok := response.Response.(oci_apigateway.GetApiResponse); ok {
			return *apiResponse.LifecycleDetails == "Validating"
		}
		return false
	}
}

func apiWaitForValidation(apiId *string, timeout time.Duration, disableFoundRetries bool, client *oci_apigateway.ApiGatewayClient) error {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "apigateway")
	retryPolicy.ShouldRetryOperation = apiResourceShouldRetryFunc(timeout)

	response := oci_apigateway.GetApiResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			"Validating",
		},
		Target: []string{
			"New",
			"Valid",
			"Warning",
			"Error",
			"Failed",
			"Canceled",
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetApi(context.Background(),
				oci_apigateway.GetApiRequest{
					ApiId: apiId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			api := &response.Api
			return api, *api.LifecycleDetails, err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return e
	}
	return nil
}
