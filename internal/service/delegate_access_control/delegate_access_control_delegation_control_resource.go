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

func DelegateAccessControlDelegationControlResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDelegateAccessControlDelegationControl,
		Read:     readDelegateAccessControlDelegationControl,
		Update:   updateDelegateAccessControlDelegationControl,
		Delete:   deleteDelegateAccessControlDelegationControl,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"delegation_subscription_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notification_message_format": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notification_topic_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resource_type": {
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
			"is_auto_approve_during_maintenance": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"num_approvals_required": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pre_approved_service_provider_action_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vault_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vault_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
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
			"time_deleted": {
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

func createDelegateAccessControlDelegationControl(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DelegateAccessControlWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readDelegateAccessControlDelegationControl(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

func updateDelegateAccessControlDelegationControl(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DelegateAccessControlWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDelegateAccessControlDelegationControl(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationControlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DelegateAccessControlWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type DelegateAccessControlDelegationControlResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_delegate_access_control.DelegateAccessControlClient
	Res                    *oci_delegate_access_control.DelegationControl
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_delegate_access_control.WorkRequestClient
}

func (s *DelegateAccessControlDelegationControlResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DelegateAccessControlDelegationControlResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_delegate_access_control.DelegationControlLifecycleStateCreating),
	}
}

func (s *DelegateAccessControlDelegationControlResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_delegate_access_control.DelegationControlLifecycleStateActive),
		string(oci_delegate_access_control.DelegationControlLifecycleStateNeedsAttention),
	}
}

func (s *DelegateAccessControlDelegationControlResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_delegate_access_control.DelegationControlLifecycleStateDeleting),
	}
}

func (s *DelegateAccessControlDelegationControlResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_delegate_access_control.DelegationControlLifecycleStateDeleted),
	}
}

func (s *DelegateAccessControlDelegationControlResourceCrud) Create() error {
	request := oci_delegate_access_control.CreateDelegationControlRequest{}

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

	if delegationSubscriptionIds, ok := s.D.GetOkExists("delegation_subscription_ids"); ok {
		interfaces := delegationSubscriptionIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("delegation_subscription_ids") {
			request.DelegationSubscriptionIds = tmp
		}
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

	if isAutoApproveDuringMaintenance, ok := s.D.GetOkExists("is_auto_approve_during_maintenance"); ok {
		tmp := isAutoApproveDuringMaintenance.(bool)
		request.IsAutoApproveDuringMaintenance = &tmp
	}

	if notificationMessageFormat, ok := s.D.GetOkExists("notification_message_format"); ok {
		request.NotificationMessageFormat = oci_delegate_access_control.DelegationControlNotificationMessageFormatEnum(notificationMessageFormat.(string))
	}

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	if numApprovalsRequired, ok := s.D.GetOkExists("num_approvals_required"); ok {
		tmp := numApprovalsRequired.(int)
		request.NumApprovalsRequired = &tmp
	}

	if preApprovedServiceProviderActionNames, ok := s.D.GetOkExists("pre_approved_service_provider_action_names"); ok {
		interfaces := preApprovedServiceProviderActionNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("pre_approved_service_provider_action_names") {
			request.PreApprovedServiceProviderActionNames = tmp
		}
	}

	if resourceIds, ok := s.D.GetOkExists("resource_ids"); ok {
		interfaces := resourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resource_ids") {
			request.ResourceIds = tmp
		}
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		request.ResourceType = oci_delegate_access_control.DelegationControlResourceTypeEnum(resourceType.(string))
	}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	if vaultKeyId, ok := s.D.GetOkExists("vault_key_id"); ok {
		tmp := vaultKeyId.(string)
		request.VaultKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control")

	response, err := s.Client.CreateDelegationControl(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDelegationControlFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control"), oci_delegate_access_control.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DelegateAccessControlDelegationControlResourceCrud) getDelegationControlFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_delegate_access_control.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	delegationControlId, err := delegationControlWaitForWorkRequest(workId, "delegationcontrol",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*delegationControlId)

	return s.Get()
}

func delegationControlWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func delegationControlWaitForWorkRequest(wId *string, entityType string, action oci_delegate_access_control.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_delegate_access_control.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "delegate_access_control")
	retryPolicy.ShouldRetryOperation = delegationControlWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDelegateAccessControlDelegationControlWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDelegateAccessControlDelegationControlWorkRequest(client *oci_delegate_access_control.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_delegate_access_control.ActionTypeEnum) error {
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

func (s *DelegateAccessControlDelegationControlResourceCrud) Get() error {
	request := oci_delegate_access_control.GetDelegationControlRequest{}

	tmp := s.D.Id()
	request.DelegationControlId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control")

	response, err := s.Client.GetDelegationControl(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DelegationControl
	return nil
}

func (s *DelegateAccessControlDelegationControlResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_delegate_access_control.UpdateDelegationControlRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.DelegationControlId = &tmp

	if delegationSubscriptionIds, ok := s.D.GetOkExists("delegation_subscription_ids"); ok {
		interfaces := delegationSubscriptionIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("delegation_subscription_ids") {
			request.DelegationSubscriptionIds = tmp
		}
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

	if isAutoApproveDuringMaintenance, ok := s.D.GetOkExists("is_auto_approve_during_maintenance"); ok {
		tmp := isAutoApproveDuringMaintenance.(bool)
		request.IsAutoApproveDuringMaintenance = &tmp
	}

	if notificationMessageFormat, ok := s.D.GetOkExists("notification_message_format"); ok {
		request.NotificationMessageFormat = oci_delegate_access_control.DelegationControlNotificationMessageFormatEnum(notificationMessageFormat.(string))
	}

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	if numApprovalsRequired, ok := s.D.GetOkExists("num_approvals_required"); ok {
		tmp := numApprovalsRequired.(int)
		request.NumApprovalsRequired = &tmp
	}

	if preApprovedServiceProviderActionNames, ok := s.D.GetOkExists("pre_approved_service_provider_action_names"); ok {
		interfaces := preApprovedServiceProviderActionNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("pre_approved_service_provider_action_names") {
			request.PreApprovedServiceProviderActionNames = tmp
		}
	}

	if resourceIds, ok := s.D.GetOkExists("resource_ids"); ok {
		interfaces := resourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resource_ids") {
			request.ResourceIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control")

	response, err := s.Client.UpdateDelegationControl(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDelegationControlFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control"), oci_delegate_access_control.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DelegateAccessControlDelegationControlResourceCrud) Delete() error {
	request := oci_delegate_access_control.DeleteDelegationControlRequest{}

	tmp := s.D.Id()
	request.DelegationControlId = &tmp

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control")

	response, err := s.Client.DeleteDelegationControl(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := delegationControlWaitForWorkRequest(workId, "delegationcontrol",
		oci_delegate_access_control.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DelegateAccessControlDelegationControlResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("delegation_subscription_ids", s.Res.DelegationSubscriptionIds)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoApproveDuringMaintenance != nil {
		s.D.Set("is_auto_approve_during_maintenance", *s.Res.IsAutoApproveDuringMaintenance)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	s.D.Set("notification_message_format", s.Res.NotificationMessageFormat)

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	if s.Res.NumApprovalsRequired != nil {
		s.D.Set("num_approvals_required", *s.Res.NumApprovalsRequired)
	}

	s.D.Set("pre_approved_service_provider_action_names", s.Res.PreApprovedServiceProviderActionNames)

	s.D.Set("resource_ids", s.Res.ResourceIds)

	s.D.Set("resource_type", s.Res.ResourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeleted != nil {
		s.D.Set("time_deleted", s.Res.TimeDeleted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	if s.Res.VaultKeyId != nil {
		s.D.Set("vault_key_id", *s.Res.VaultKeyId)
	}

	return nil
}

func DelegationControlSummaryToMap(obj oci_delegate_access_control.DelegationControlSummary) map[string]interface{} {
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

	result["resource_type"] = string(obj.ResourceType)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeDeleted != nil {
		result["time_deleted"] = obj.TimeDeleted.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DelegateAccessControlDelegationControlResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_delegate_access_control.ChangeDelegationControlCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DelegationControlId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control")

	response, err := s.Client.ChangeDelegationControlCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDelegationControlFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "delegate_access_control"), oci_delegate_access_control.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
