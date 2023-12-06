// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RecoveryRecoveryServiceSubnetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRecoveryRecoveryServiceSubnet,
		Read:     readRecoveryRecoveryServiceSubnet,
		Update:   updateRecoveryRecoveryServiceSubnet,
		Delete:   deleteRecoveryRecoveryServiceSubnet,
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
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": {
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

			// Computed
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
		},
	}
}

func createRecoveryRecoveryServiceSubnet(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryRecoveryServiceSubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.CreateResource(d, sync)
}

func readRecoveryRecoveryServiceSubnet(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryRecoveryServiceSubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.ReadResource(sync)
}

func updateRecoveryRecoveryServiceSubnet(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryRecoveryServiceSubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteRecoveryRecoveryServiceSubnet(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryRecoveryServiceSubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type RecoveryRecoveryServiceSubnetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_recovery.DatabaseRecoveryClient
	Res                    *oci_recovery.RecoveryServiceSubnet
	DisableNotFoundRetries bool
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_recovery.LifecycleStateCreating),
	}
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_recovery.LifecycleStateActive),
	}
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_recovery.LifecycleStateDeleting),
	}
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_recovery.LifecycleStateDeleted),
	}
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) Create() error {
	request := oci_recovery.CreateRecoveryServiceSubnetRequest{}

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

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.CreateRecoveryServiceSubnet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getRecoveryServiceSubnetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) getRecoveryServiceSubnetFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_recovery.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	recoveryServiceSubnetId, err := recoveryServiceSubnetWaitForWorkRequest(workId, "recoveryservicesubnet",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*recoveryServiceSubnetId)

	return s.Get()
}

func recoveryServiceSubnetWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "recovery", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_recovery.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func recoveryServiceSubnetWaitForWorkRequest(wId *string, entityType string, action oci_recovery.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_recovery.DatabaseRecoveryClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "recovery")
	retryPolicy.ShouldRetryOperation = recoveryServiceSubnetWorkRequestShouldRetryFunc(timeout)

	response := oci_recovery.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_recovery.OperationStatusInProgress),
			string(oci_recovery.OperationStatusAccepted),
			string(oci_recovery.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_recovery.OperationStatusSucceeded),
			string(oci_recovery.OperationStatusFailed),
			string(oci_recovery.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_recovery.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_recovery.OperationStatusFailed || response.Status == oci_recovery.OperationStatusCanceled {
		return nil, getErrorFromRecoveryRecoveryServiceSubnetWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromRecoveryRecoveryServiceSubnetWorkRequest(client *oci_recovery.DatabaseRecoveryClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_recovery.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_recovery.ListWorkRequestErrorsRequest{
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

func (s *RecoveryRecoveryServiceSubnetResourceCrud) Get() error {
	request := oci_recovery.GetRecoveryServiceSubnetRequest{}

	tmp := s.D.Id()
	request.RecoveryServiceSubnetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.GetRecoveryServiceSubnet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RecoveryServiceSubnet
	return nil
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_recovery.UpdateRecoveryServiceSubnetRequest{}

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
	request.RecoveryServiceSubnetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.UpdateRecoveryServiceSubnet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRecoveryServiceSubnetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) Delete() error {
	request := oci_recovery.DeleteRecoveryServiceSubnetRequest{}

	tmp := s.D.Id()
	request.RecoveryServiceSubnetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.DeleteRecoveryServiceSubnet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := recoveryServiceSubnetWaitForWorkRequest(workId, "recoveryservicesubnet",
		oci_recovery.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) SetData() error {
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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func RecoveryServiceSubnetSummaryToMap(obj oci_recovery.RecoveryServiceSubnetSummary) map[string]interface{} {
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

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}

func (s *RecoveryRecoveryServiceSubnetResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_recovery.ChangeRecoveryServiceSubnetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.RecoveryServiceSubnetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.ChangeRecoveryServiceSubnetCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRecoveryServiceSubnetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
