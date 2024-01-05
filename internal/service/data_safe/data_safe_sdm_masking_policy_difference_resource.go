// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSdmMaskingPolicyDifferenceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeSdmMaskingPolicyDifference,
		Read:     readDataSafeSdmMaskingPolicyDifference,
		Update:   updateDataSafeSdmMaskingPolicyDifference,
		Delete:   deleteDataSafeSdmMaskingPolicyDifference,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"masking_policy_id": {
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
			"difference_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"sensitive_data_model_id": {
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
			"time_creation_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeSdmMaskingPolicyDifference(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSdmMaskingPolicyDifferenceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeSdmMaskingPolicyDifference(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSdmMaskingPolicyDifferenceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeSdmMaskingPolicyDifference(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSdmMaskingPolicyDifferenceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeSdmMaskingPolicyDifference(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSdmMaskingPolicyDifferenceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeSdmMaskingPolicyDifferenceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SdmMaskingPolicyDifference
	DisableNotFoundRetries bool
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.SdmMaskingPolicyDifferenceLifecycleStateCreating),
	}
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.SdmMaskingPolicyDifferenceLifecycleStateActive),
	}
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.SdmMaskingPolicyDifferenceLifecycleStateDeleting),
	}
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.SdmMaskingPolicyDifferenceLifecycleStateDeleted),
	}
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) Create() error {
	request := oci_data_safe.CreateSdmMaskingPolicyDifferenceRequest{}

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

	if differenceType, ok := s.D.GetOkExists("difference_type"); ok {
		request.DifferenceType = oci_data_safe.SdmMaskingPolicyDifferenceDifferenceTypeEnum(differenceType.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateSdmMaskingPolicyDifference(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getSdmMaskingPolicyDifferenceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) getSdmMaskingPolicyDifferenceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	sdmMaskingPolicyDifferenceId, err := sdmMaskingPolicyDifferenceWaitForWorkRequest(workId, "sdmmaskingpolicydifference",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, sdmMaskingPolicyDifferenceId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_data_safe.CancelWorkRequestRequest{
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
	s.D.SetId(*sdmMaskingPolicyDifferenceId)

	return s.Get()
}

func sdmMaskingPolicyDifferenceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func sdmMaskingPolicyDifferenceWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = sdmMaskingPolicyDifferenceWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
			string(oci_data_safe.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_data_safe.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeSdmMaskingPolicyDifferenceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeSdmMaskingPolicyDifferenceWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_data_safe.ListWorkRequestErrorsRequest{
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

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) Get() error {
	request := oci_data_safe.GetSdmMaskingPolicyDifferenceRequest{}

	tmp := s.D.Id()
	request.SdmMaskingPolicyDifferenceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSdmMaskingPolicyDifference(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SdmMaskingPolicyDifference
	return nil
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateSdmMaskingPolicyDifferenceRequest{}

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

	tmp := s.D.Id()
	request.SdmMaskingPolicyDifferenceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSdmMaskingPolicyDifference(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSdmMaskingPolicyDifferenceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) Delete() error {
	request := oci_data_safe.DeleteSdmMaskingPolicyDifferenceRequest{}

	tmp := s.D.Id()
	request.SdmMaskingPolicyDifferenceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteSdmMaskingPolicyDifference(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := sdmMaskingPolicyDifferenceWaitForWorkRequest(workId, "sdmmaskingpolicydifference",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("difference_type", s.Res.DifferenceType)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MaskingPolicyId != nil {
		s.D.Set("masking_policy_id", *s.Res.MaskingPolicyId)
	}

	if s.Res.SensitiveDataModelId != nil {
		s.D.Set("sensitive_data_model_id", *s.Res.SensitiveDataModelId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeCreationStarted != nil {
		s.D.Set("time_creation_started", s.Res.TimeCreationStarted.String())
	}

	return nil
}

func SdmMaskingPolicyDifferenceSummaryToMap(obj oci_data_safe.SdmMaskingPolicyDifferenceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["difference_type"] = string(obj.DifferenceType)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MaskingPolicyId != nil {
		result["masking_policy_id"] = string(*obj.MaskingPolicyId)
	}

	if obj.SensitiveDataModelId != nil {
		result["sensitive_data_model_id"] = string(*obj.SensitiveDataModelId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeCreationStarted != nil {
		result["time_creation_started"] = obj.TimeCreationStarted.String()
	}

	return result
}

func (s *DataSafeSdmMaskingPolicyDifferenceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeSdmMaskingPolicyDifferenceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SdmMaskingPolicyDifferenceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeSdmMaskingPolicyDifferenceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
