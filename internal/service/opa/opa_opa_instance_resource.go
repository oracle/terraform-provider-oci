// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opa

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opa "github.com/oracle/oci-go-sdk/v65/opa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpaOpaInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpaOpaInstance,
		Read:     readOpaOpaInstance,
		Update:   updateOpaOpaInstance,
		Delete:   deleteOpaOpaInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"consumption_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"idcs_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_breakglass_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metering_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_opa.OpaInstanceLifecycleStateInactive),
					string(oci_opa.OpaInstanceLifecycleStateActive),
				}, true),
			},

			// Computed
			"attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_implicit": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_instance_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_service_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"identity_app_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_app_guid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_app_opc_service_instance_guid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_domain_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_url": {
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

func createOpaOpaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OpaOpaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpaInstanceClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_opa.OpaInstanceLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_opa.OpaInstanceLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopOpaInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_opa.OpaInstanceLifecycleStateInactive)
	}
	return nil

}

func readOpaOpaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OpaOpaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpaInstanceClient()

	return tfresource.ReadResource(sync)
}

func updateOpaOpaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OpaOpaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpaInstanceClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_opa.OpaInstanceLifecycleStateActive == oci_opa.OpaInstanceLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_opa.OpaInstanceLifecycleStateInactive == oci_opa.OpaInstanceLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartOpaInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_opa.OpaInstanceLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopOpaInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_opa.OpaInstanceLifecycleStateInactive)
	}

	return nil
}

func deleteOpaOpaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OpaOpaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpaInstanceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpaOpaInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opa.OpaInstanceClient
	Res                    *oci_opa.OpaInstance
	DisableNotFoundRetries bool
}

func (s *OpaOpaInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OpaOpaInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opa.OpaInstanceLifecycleStateCreating),
	}
}

func (s *OpaOpaInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opa.OpaInstanceLifecycleStateActive),
	}
}

func (s *OpaOpaInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opa.OpaInstanceLifecycleStateDeleting),
	}
}

func (s *OpaOpaInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opa.OpaInstanceLifecycleStateDeleted),
	}
}

func (s *OpaOpaInstanceResourceCrud) Create() error {
	request := oci_opa.CreateOpaInstanceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if consumptionModel, ok := s.D.GetOkExists("consumption_model"); ok {
		request.ConsumptionModel = oci_opa.OpaInstanceConsumptionModelEnum(consumptionModel.(string))
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

	if idcsAt, ok := s.D.GetOkExists("idcs_at"); ok {
		tmp := idcsAt.(string)
		request.IdcsAt = &tmp
	}

	if isBreakglassEnabled, ok := s.D.GetOkExists("is_breakglass_enabled"); ok {
		tmp := isBreakglassEnabled.(bool)
		request.IsBreakglassEnabled = &tmp
	}

	if meteringType, ok := s.D.GetOkExists("metering_type"); ok {
		request.MeteringType = oci_opa.OpaInstanceMeteringTypeEnum(meteringType.(string))
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		request.ShapeName = oci_opa.OpaInstanceShapeNameEnum(shapeName.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa")

	response, err := s.Client.CreateOpaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_opa.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_opa.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "opa") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getOpaInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa"), oci_opa.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OpaOpaInstanceResourceCrud) getOpaInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opa.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	opaInstanceId, err := opaInstanceWaitForWorkRequest(workId, "opa",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, opaInstanceId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_opa.CancelWorkRequestRequest{
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
	s.D.SetId(*opaInstanceId)

	return s.Get()
}

func opaInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opa", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opa.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func opaInstanceWaitForWorkRequest(wId *string, entityType string, action oci_opa.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opa.OpaInstanceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opa")
	retryPolicy.ShouldRetryOperation = opaInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_opa.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_opa.OperationStatusInProgress),
			string(oci_opa.OperationStatusAccepted),
			string(oci_opa.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opa.OperationStatusSucceeded),
			string(oci_opa.OperationStatusFailed),
			string(oci_opa.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opa.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opa.OperationStatusFailed || response.Status == oci_opa.OperationStatusCanceled {
		return nil, getErrorFromOpaOpaInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpaOpaInstanceWorkRequest(client *oci_opa.OpaInstanceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opa.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opa.ListWorkRequestErrorsRequest{
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

func (s *OpaOpaInstanceResourceCrud) Get() error {
	request := oci_opa.GetOpaInstanceRequest{}

	tmp := s.D.Id()
	request.OpaInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa")

	response, err := s.Client.GetOpaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OpaInstance
	return nil
}

func (s *OpaOpaInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_opa.UpdateOpaInstanceRequest{}

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

	tmp := s.D.Id()
	request.OpaInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa")

	response, err := s.Client.UpdateOpaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOpaInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa"), oci_opa.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpaOpaInstanceResourceCrud) Delete() error {
	request := oci_opa.DeleteOpaInstanceRequest{}

	tmp := s.D.Id()
	request.OpaInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa")

	response, err := s.Client.DeleteOpaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	if _, err := opaInstanceWaitForWorkRequest(workId, "opainstance",
		oci_opa.ActionTypeRelated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client); err != nil {
		return err
	}
	return nil
}

func (s *OpaOpaInstanceResourceCrud) SetData() error {
	attachments := []interface{}{}
	if s.Res.Attachments != nil {
		for _, item := range s.Res.Attachments {
			attachments = append(attachments, AttachmentDetailsToMap(item))
		}
	}
	s.D.Set("attachments", attachments)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("consumption_model", s.Res.ConsumptionModel)

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

	if s.Res.IdentityAppDisplayName != nil {
		s.D.Set("identity_app_display_name", *s.Res.IdentityAppDisplayName)
	}

	if s.Res.IdentityAppGuid != nil {
		s.D.Set("identity_app_guid", *s.Res.IdentityAppGuid)
	}

	if s.Res.IdentityAppOpcServiceInstanceGuid != nil {
		s.D.Set("identity_app_opc_service_instance_guid", *s.Res.IdentityAppOpcServiceInstanceGuid)
	}

	if s.Res.IdentityDomainUrl != nil {
		s.D.Set("identity_domain_url", *s.Res.IdentityDomainUrl)
	}

	if s.Res.InstanceUrl != nil {
		s.D.Set("instance_url", *s.Res.InstanceUrl)
	}

	if s.Res.IsBreakglassEnabled != nil {
		s.D.Set("is_breakglass_enabled", *s.Res.IsBreakglassEnabled)
	}

	s.D.Set("metering_type", s.Res.MeteringType)

	s.D.Set("shape_name", s.Res.ShapeName)

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

func (s *OpaOpaInstanceResourceCrud) StartOpaInstance() error {
	request := oci_opa.StartOpaInstanceRequest{}

	idTmp := s.D.Id()
	request.OpaInstanceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa")

	_, err := s.Client.StartOpaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_opa.OpaInstanceLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpaOpaInstanceResourceCrud) StopOpaInstance() error {
	request := oci_opa.StopOpaInstanceRequest{}

	idTmp := s.D.Id()
	request.OpaInstanceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa")

	_, err := s.Client.StopOpaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_opa.OpaInstanceLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func AttachmentDetailsToMap(obj oci_opa.AttachmentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsImplicit != nil {
		result["is_implicit"] = bool(*obj.IsImplicit)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TargetInstanceUrl != nil {
		result["target_instance_url"] = string(*obj.TargetInstanceUrl)
	}

	result["target_role"] = string(obj.TargetRole)

	if obj.TargetServiceType != nil {
		result["target_service_type"] = string(*obj.TargetServiceType)
	}

	return result
}

func OpaInstanceSummaryToMap(obj oci_opa.OpaInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["consumption_model"] = string(obj.ConsumptionModel)

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

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InstanceUrl != nil {
		result["instance_url"] = string(*obj.InstanceUrl)
	}

	if obj.IsBreakglassEnabled != nil {
		result["is_breakglass_enabled"] = bool(*obj.IsBreakglassEnabled)
	}

	result["metering_type"] = string(obj.MeteringType)

	result["shape_name"] = string(obj.ShapeName)

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

func (s *OpaOpaInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_opa.ChangeOpaInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OpaInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa")

	response, err := s.Client.ChangeOpaInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOpaInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opa"), oci_opa.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
