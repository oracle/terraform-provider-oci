// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApigatewayUsagePlanResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApigatewayUsagePlan,
		Read:     readApigatewayUsagePlan,
		Update:   updateApigatewayUsagePlan,
		Delete:   deleteApigatewayUsagePlan,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entitlements": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quota": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"operation_on_breach": {
										Type:     schema.TypeString,
										Required: true,
									},
									"reset_policy": {
										Type:     schema.TypeString,
										Required: true,
									},
									"unit": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeInt,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"rate_limit": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"unit": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeInt,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"targets": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"deployment_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
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

func createApigatewayUsagePlan(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayUsagePlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsagePlansClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.CreateResource(d, sync)
}

func readApigatewayUsagePlan(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayUsagePlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsagePlansClient()

	return tfresource.ReadResource(sync)
}

func updateApigatewayUsagePlan(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayUsagePlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsagePlansClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApigatewayUsagePlan(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayUsagePlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsagePlansClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.DeleteResource(d, sync)
}

type ApigatewayUsagePlanResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apigateway.UsagePlansClient
	Res                    *oci_apigateway.UsagePlan
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_apigateway.WorkRequestsClient
}

func (s *ApigatewayUsagePlanResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ApigatewayUsagePlanResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_apigateway.UsagePlanLifecycleStateCreating),
	}
}

func (s *ApigatewayUsagePlanResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_apigateway.UsagePlanLifecycleStateActive),
	}
}

func (s *ApigatewayUsagePlanResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_apigateway.UsagePlanLifecycleStateDeleting),
	}
}

func (s *ApigatewayUsagePlanResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_apigateway.UsagePlanLifecycleStateDeleted),
	}
}

func (s *ApigatewayUsagePlanResourceCrud) Create() error {
	request := oci_apigateway.CreateUsagePlanRequest{}

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

	if entitlements, ok := s.D.GetOkExists("entitlements"); ok {
		interfaces := entitlements.([]interface{})
		tmp := make([]oci_apigateway.Entitlement, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "entitlements", stateDataIndex)
			converted, err := s.mapToEntitlement(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("entitlements") {
			request.Entitlements = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.CreateUsagePlan(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getUsagePlanFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ApigatewayUsagePlanResourceCrud) getUsagePlanFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_apigateway.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	usagePlanId, err := usagePlanWaitForWorkRequest(workId, "usageplan",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, usagePlanId)
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
	s.D.SetId(*usagePlanId)

	return s.Get()
}

func usagePlanWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func usagePlanWaitForWorkRequest(wId *string, entityType string, action oci_apigateway.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_apigateway.WorkRequestsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "apigateway")
	retryPolicy.ShouldRetryOperation = usagePlanWorkRequestShouldRetryFunc(timeout)

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
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_apigateway.WorkRequestStatusFailed || response.Status == oci_apigateway.WorkRequestStatusCanceled {
		return nil, getErrorFromApigatewayUsagePlanWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromApigatewayUsagePlanWorkRequest(client *oci_apigateway.WorkRequestsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_apigateway.WorkRequestResourceActionTypeEnum) error {
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

func (s *ApigatewayUsagePlanResourceCrud) Get() error {
	request := oci_apigateway.GetUsagePlanRequest{}

	tmp := s.D.Id()
	request.UsagePlanId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.GetUsagePlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UsagePlan
	return nil
}

func (s *ApigatewayUsagePlanResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_apigateway.UpdateUsagePlanRequest{}

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

	if entitlements, ok := s.D.GetOkExists("entitlements"); ok {
		interfaces := entitlements.([]interface{})
		tmp := make([]oci_apigateway.Entitlement, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "entitlements", stateDataIndex)
			converted, err := s.mapToEntitlement(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("entitlements") {
			request.Entitlements = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.UsagePlanId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.UpdateUsagePlan(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUsagePlanFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ApigatewayUsagePlanResourceCrud) Delete() error {
	request := oci_apigateway.DeleteUsagePlanRequest{}

	tmp := s.D.Id()
	request.UsagePlanId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.DeleteUsagePlan(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := usagePlanWaitForWorkRequest(workId, "usageplan",
		oci_apigateway.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *ApigatewayUsagePlanResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	entitlements := []interface{}{}
	for _, item := range s.Res.Entitlements {
		entitlements = append(entitlements, EntitlementToMap(item))
	}
	s.D.Set("entitlements", entitlements)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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

func (s *ApigatewayUsagePlanResourceCrud) mapToEntitlement(fieldKeyFormat string) (oci_apigateway.Entitlement, error) {
	result := oci_apigateway.Entitlement{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if quota, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "quota")); ok {
		if tmpList := quota.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "quota"), 0)
			tmp, err := s.mapToQuota(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert quota, encountered error: %v", err)
			}
			result.Quota = &tmp
		}
	}

	if rateLimit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rate_limit")); ok {
		if tmpList := rateLimit.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rate_limit"), 0)
			tmp, err := s.mapToRateLimit(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert rate_limit, encountered error: %v", err)
			}
			result.RateLimit = &tmp
		}
	}

	if targets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "targets")); ok {
		interfaces := targets.([]interface{})
		tmp := make([]oci_apigateway.EntitlementTarget, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "targets"), stateDataIndex)
			converted, err := s.mapToEntitlementTarget(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "targets")) {
			result.Targets = tmp
		}
	}

	return result, nil
}

func (s *ApigatewayUsagePlanResourceCrud) mapToQuota(fieldKeyFormat string) (oci_apigateway.Quota, error) {
	result := oci_apigateway.Quota{}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(int)
		result.Value = &tmp
	}
	if unit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit")); ok {
		result.Unit = oci_apigateway.QuotaUnitEnum(unit.(string))
	}
	if resetPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reset_policy")); ok {
		log.Printf("[DEBUG] resetPolicy of Usage Plan: %s\n", resetPolicy)
		result.ResetPolicy = oci_apigateway.QuotaResetPolicyEnum(resetPolicy.(string))
	}
	if operationOnBreach, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_on_breach")); ok {
		log.Printf("[DEBUG] operationOnBreach of Usage Plan: %s\n", operationOnBreach)
		result.OperationOnBreach = oci_apigateway.QuotaOperationOnBreachEnum(operationOnBreach.(string))
	}

	return result, nil
}

func (s *ApigatewayUsagePlanResourceCrud) mapToRateLimit(fieldKeyFormat string) (oci_apigateway.RateLimit, error) {
	result := oci_apigateway.RateLimit{}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(int)
		result.Value = &tmp
	}
	if unit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit")); ok {
		result.Unit = oci_apigateway.RateLimitUnitEnum(unit.(string))
	}

	return result, nil
}

func EntitlementToMap(obj oci_apigateway.Entitlement) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Quota != nil {
		result["quota"] = []interface{}{QuotaToMap(obj.Quota)}
	}

	if obj.RateLimit != nil {
		result["rate_limit"] = []interface{}{RateLimitToMap(obj.RateLimit)}
	}

	targets := []interface{}{}
	for _, item := range obj.Targets {
		targets = append(targets, EntitlementTargetToMap(item))
	}
	result["targets"] = targets

	return result
}

func EntitlementSummaryToMap(obj oci_apigateway.EntitlementSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Quota != nil {
		result["quota"] = []interface{}{QuotaToMap(obj.Quota)}
	}

	if obj.RateLimit != nil {
		result["rate_limit"] = []interface{}{RateLimitToMap(obj.RateLimit)}
	}

	return result
}

func (s *ApigatewayUsagePlanResourceCrud) mapToEntitlementTarget(fieldKeyFormat string) (oci_apigateway.EntitlementTarget, error) {
	result := oci_apigateway.EntitlementTarget{}

	if deploymentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deployment_id")); ok {
		tmp := deploymentId.(string)
		result.DeploymentId = &tmp
	}

	return result, nil
}

func EntitlementTargetToMap(obj oci_apigateway.EntitlementTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DeploymentId != nil {
		result["deployment_id"] = string(*obj.DeploymentId)
	}

	return result
}

func QuotaToMap(obj *oci_apigateway.Quota) map[string]interface{} {
	result := map[string]interface{}{}

	result["operation_on_breach"] = string(obj.OperationOnBreach)

	result["reset_policy"] = string(obj.ResetPolicy)

	result["unit"] = string(obj.Unit)

	if obj.Value != nil {
		result["value"] = int(*obj.Value)
	}

	return result
}

func RateLimitToMap(obj *oci_apigateway.RateLimit) map[string]interface{} {
	result := map[string]interface{}{}

	result["unit"] = string(obj.Unit)

	if obj.Value != nil {
		result["value"] = int(*obj.Value)
	}

	return result
}

func UsagePlanSummaryToMap(obj oci_apigateway.UsagePlanSummary) map[string]interface{} {
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

	entitlements := []interface{}{}
	for _, item := range obj.Entitlements {
		entitlements = append(entitlements, EntitlementSummaryToMap(item))
	}
	result["entitlements"] = entitlements

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func (s *ApigatewayUsagePlanResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_apigateway.ChangeUsagePlanCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.UsagePlanId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.ChangeUsagePlanCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUsagePlanFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
