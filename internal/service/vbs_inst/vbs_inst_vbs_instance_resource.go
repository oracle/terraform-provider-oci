// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vbs_inst

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_vbs_inst "github.com/oracle/oci-go-sdk/v65/vbsinst"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func VbsInstVbsInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createVbsInstVbsInstance,
		Read:     readVbsInstVbsInstance,
		Update:   updateVbsInstVbsInstance,
		Delete:   deleteVbsInstVbsInstance,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"idcs_access_token": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_resource_usage_agreement_granted": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"resource_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecyle_details": {
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
			"vbs_access_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createVbsInstVbsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &VbsInstVbsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbsInstanceClient()

	return tfresource.CreateResource(d, sync)
}

func readVbsInstVbsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &VbsInstVbsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbsInstanceClient()

	return tfresource.ReadResource(sync)
}

func updateVbsInstVbsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &VbsInstVbsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbsInstanceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteVbsInstVbsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &VbsInstVbsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbsInstanceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type VbsInstVbsInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_vbs_inst.VbsInstanceClient
	Res                    *oci_vbs_inst.VbsInstance
	DisableNotFoundRetries bool
}

func (s *VbsInstVbsInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VbsInstVbsInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_vbs_inst.LifecycleStateCreating),
	}
}

func (s *VbsInstVbsInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_vbs_inst.LifecycleStateActive),
	}
}

func (s *VbsInstVbsInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_vbs_inst.LifecycleStateDeleting),
	}
}

func (s *VbsInstVbsInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_vbs_inst.LifecycleStateDeleted),
	}
}

func (s *VbsInstVbsInstanceResourceCrud) Create() error {
	request := oci_vbs_inst.CreateVbsInstanceRequest{}

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

	if idcsAccessToken, ok := s.D.GetOkExists("idcs_access_token"); ok {
		tmp := idcsAccessToken.(string)
		request.IdcsAccessToken = &tmp
	}

	if isResourceUsageAgreementGranted, ok := s.D.GetOkExists("is_resource_usage_agreement_granted"); ok {
		tmp := isResourceUsageAgreementGranted.(bool)
		request.IsResourceUsageAgreementGranted = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if resourceCompartmentId, ok := s.D.GetOkExists("resource_compartment_id"); ok {
		tmp := resourceCompartmentId.(string)
		request.ResourceCompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vbs_inst")

	response, err := s.Client.CreateVbsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_vbs_inst.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_vbs_inst.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vbs_inst"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "vbsinstance") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getVbsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vbs_inst"), oci_vbs_inst.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *VbsInstVbsInstanceResourceCrud) getVbsInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_vbs_inst.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	vbsInstanceId, err := vbsInstanceWaitForWorkRequest(workId, "vbsinstance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*vbsInstanceId)

	return s.Get()
}

func vbsInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "vbs_inst", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_vbs_inst.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func vbsInstanceWaitForWorkRequest(wId *string, entityType string, action oci_vbs_inst.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_vbs_inst.VbsInstanceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "vbs_inst")
	retryPolicy.ShouldRetryOperation = vbsInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_vbs_inst.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_vbs_inst.OperationStatusInProgress),
			string(oci_vbs_inst.OperationStatusAccepted),
			string(oci_vbs_inst.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_vbs_inst.OperationStatusSucceeded),
			string(oci_vbs_inst.OperationStatusFailed),
			string(oci_vbs_inst.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_vbs_inst.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_vbs_inst.OperationStatusFailed || response.Status == oci_vbs_inst.OperationStatusCanceled {
		return nil, getErrorFromVbsInstVbsInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromVbsInstVbsInstanceWorkRequest(client *oci_vbs_inst.VbsInstanceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_vbs_inst.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_vbs_inst.ListWorkRequestErrorsRequest{
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

func (s *VbsInstVbsInstanceResourceCrud) Get() error {
	request := oci_vbs_inst.GetVbsInstanceRequest{}

	tmp := s.D.Id()
	request.VbsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vbs_inst")

	response, err := s.Client.GetVbsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VbsInstance
	return nil
}

func (s *VbsInstVbsInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_vbs_inst.UpdateVbsInstanceRequest{}

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

	if isResourceUsageAgreementGranted, ok := s.D.GetOkExists("is_resource_usage_agreement_granted"); ok {
		tmp := isResourceUsageAgreementGranted.(bool)
		request.IsResourceUsageAgreementGranted = &tmp
	}

	if resourceCompartmentId, ok := s.D.GetOkExists("resource_compartment_id"); ok {
		tmp := resourceCompartmentId.(string)
		request.ResourceCompartmentId = &tmp
	}

	tmp := s.D.Id()
	request.VbsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vbs_inst")

	response, err := s.Client.UpdateVbsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVbsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vbs_inst"), oci_vbs_inst.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *VbsInstVbsInstanceResourceCrud) Delete() error {
	request := oci_vbs_inst.DeleteVbsInstanceRequest{}

	tmp := s.D.Id()
	request.VbsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vbs_inst")

	response, err := s.Client.DeleteVbsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := vbsInstanceWaitForWorkRequest(workId, "vbsinstance",
		oci_vbs_inst.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *VbsInstVbsInstanceResourceCrud) SetData() error {
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

	if s.Res.IsResourceUsageAgreementGranted != nil {
		s.D.Set("is_resource_usage_agreement_granted", *s.Res.IsResourceUsageAgreementGranted)
	}

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ResourceCompartmentId != nil {
		s.D.Set("resource_compartment_id", *s.Res.ResourceCompartmentId)
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

	if s.Res.VbsAccessUrl != nil {
		s.D.Set("vbs_access_url", *s.Res.VbsAccessUrl)
	}

	return nil
}

func VbsInstanceSummaryToMap(obj oci_vbs_inst.VbsInstanceSummary) map[string]interface{} {
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

	if obj.IsResourceUsageAgreementGranted != nil {
		result["is_resource_usage_agreement_granted"] = bool(*obj.IsResourceUsageAgreementGranted)
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

func (s *VbsInstVbsInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_vbs_inst.ChangeVbsInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VbsInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vbs_inst")

	response, err := s.Client.ChangeVbsInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVbsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vbs_inst"), oci_vbs_inst.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
