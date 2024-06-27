// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package delegate_access_control

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DelegateAccessControlDelegationSubscriptionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDelegateAccessControlDelegationSubscription,
		Read:     readDelegateAccessControlDelegationSubscription,
		Update:   updateDelegateAccessControlDelegationSubscription,
		Delete:   deleteDelegateAccessControlDelegationSubscription,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_provider_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subscribed_service_type": {
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
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_state_details": {
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

func createDelegateAccessControlDelegationSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DelegateAccessControlWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readDelegateAccessControlDelegationSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

func updateDelegateAccessControlDelegationSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DelegateAccessControlWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDelegateAccessControlDelegationSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DelegateAccessControlWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type DelegateAccessControlDelegationSubscriptionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_delegate_access_control.DelegateAccessControlClient
	Res                    *oci_delegate_access_control.DelegationSubscription
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_delegate_access_control.WorkRequestClient
}

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_delegate_access_control.DelegationSubscriptionLifecycleStateCreating),
	}
}

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_delegate_access_control.DelegationSubscriptionLifecycleStateActive),
	}
}

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_delegate_access_control.DelegationSubscriptionLifecycleStateDeleting),
	}
}

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_delegate_access_control.DelegationSubscriptionLifecycleStateDeleted),
	}
}

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) Create() error {
	request := oci_delegate_access_control.CreateDelegationSubscriptionRequest{}

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

	if serviceProviderId, ok := s.D.GetOkExists("service_provider_id"); ok {
		tmp := serviceProviderId.(string)
		request.ServiceProviderId = &tmp
	}

	if subscribedServiceType, ok := s.D.GetOkExists("subscribed_service_type"); ok {
		request.SubscribedServiceType = oci_delegate_access_control.ServiceProviderServiceTypeEnum(subscribedServiceType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control")

	response, err := s.Client.CreateDelegationSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDelegationSubscriptionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control"), oci_delegate_access_control.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) getDelegationSubscriptionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_delegate_access_control.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	delegationSubscriptionId, err := delegationSubscriptionWaitForWorkRequest(workId, "delegationsubscription",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*delegationSubscriptionId)

	return s.Get()
}

func delegationSubscriptionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "delegate_access_control", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_delegate_access_control.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func delegationSubscriptionWaitForWorkRequest(wId *string, entityType string, action oci_delegate_access_control.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_delegate_access_control.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "delegate_access_control")
	retryPolicy.ShouldRetryOperation = delegationSubscriptionWorkRequestShouldRetryFunc(timeout)

	response := oci_delegate_access_control.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_delegate_access_control.OperationStatusInProgress),
			string(oci_delegate_access_control.OperationStatusAccepted),
			string(oci_delegate_access_control.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_delegate_access_control.OperationStatusSucceeded),
			string(oci_delegate_access_control.OperationStatusFailed),
			string(oci_delegate_access_control.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_delegate_access_control.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_delegate_access_control.OperationStatusFailed || response.Status == oci_delegate_access_control.OperationStatusCanceled {
		return nil, getErrorFromDelegateAccessControlDelegationSubscriptionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDelegateAccessControlDelegationSubscriptionWorkRequest(client *oci_delegate_access_control.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_delegate_access_control.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_delegate_access_control.ListWorkRequestErrorsRequest{
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

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) Get() error {
	request := oci_delegate_access_control.GetDelegationSubscriptionRequest{}

	tmp := s.D.Id()
	request.DelegationSubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control")

	response, err := s.Client.GetDelegationSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DelegationSubscription
	return nil
}

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_delegate_access_control.UpdateDelegationSubscriptionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.DelegationSubscriptionId = &tmp

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control")

	response, err := s.Client.UpdateDelegationSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDelegationSubscriptionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control"), oci_delegate_access_control.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) Delete() error {
	request := oci_delegate_access_control.DeleteDelegationSubscriptionRequest{}

	tmp := s.D.Id()
	request.DelegationSubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control")

	response, err := s.Client.DeleteDelegationSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := delegationSubscriptionWaitForWorkRequest(workId, "delegationsubscription",
		oci_delegate_access_control.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.ServiceProviderId != nil {
		s.D.Set("service_provider_id", *s.Res.ServiceProviderId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subscribed_service_type", s.Res.SubscribedServiceType)

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

func DelegationSubscriptionSummaryToMap(obj oci_delegate_access_control.DelegationSubscriptionSummary) map[string]interface{} {
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

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.ServiceProviderId != nil {
		result["service_provider_id"] = string(*obj.ServiceProviderId)
	}

	result["state"] = string(obj.LifecycleState)

	result["subscribed_service_type"] = string(obj.SubscribedServiceType)

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

func (s *DelegateAccessControlDelegationSubscriptionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_delegate_access_control.ChangeDelegationSubscriptionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DelegationSubscriptionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control")

	response, err := s.Client.ChangeDelegationSubscriptionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDelegationSubscriptionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control"), oci_delegate_access_control.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
