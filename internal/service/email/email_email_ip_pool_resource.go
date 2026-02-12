// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func EmailEmailIpPoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createEmailEmailIpPoolWithContext,
		ReadContext:   readEmailEmailIpPoolWithContext,
		UpdateContext: updateEmailEmailIpPoolWithContext,
		DeleteContext: deleteEmailEmailIpPoolWithContext,
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
			"outbound_ips": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"outbound_ips_response": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Computed
						"assignment_state": {
							Type:     schema.TypeString,
							Computed: true},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"outbound_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
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
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createEmailEmailIpPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &EmailEmailIpPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readEmailEmailIpPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &EmailEmailIpPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateEmailEmailIpPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &EmailEmailIpPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteEmailEmailIpPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &EmailEmailIpPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type EmailEmailIpPoolResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_email.EmailClient
	Res                    *oci_email.EmailIpPool
	DisableNotFoundRetries bool
}

func (s *EmailEmailIpPoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *EmailEmailIpPoolResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_email.EmailIpPoolLifecycleStateCreating),
	}
}

func (s *EmailEmailIpPoolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_email.EmailIpPoolLifecycleStateActive),
	}
}

func (s *EmailEmailIpPoolResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_email.EmailIpPoolLifecycleStateDeleting),
	}
}

func (s *EmailEmailIpPoolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_email.EmailIpPoolLifecycleStateDeleted),
	}
}

func (s *EmailEmailIpPoolResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_email.CreateEmailIpPoolRequest{}

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
	if outboundIps, ok := s.D.GetOkExists("outbound_ips"); ok {
		interfaces := outboundIps.([]interface{})
		tmp := make([]string, 0, len(interfaces))
		for _, v := range interfaces {
			if v == nil {
				continue
			}
			tmp = append(tmp, v.(string))
		}
		request.OutboundIps = tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.CreateEmailIpPool(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getEmailIpPoolFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email"), oci_email.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *EmailEmailIpPoolResourceCrud) getEmailIpPoolFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_email.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	emailIpPoolId, err := emailIpPoolWaitForWorkRequest(ctx, workId, "emailippool",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*emailIpPoolId)

	return s.GetWithContext(ctx)
}

func emailIpPoolWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "email", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_email.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func emailIpPoolWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_email.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_email.EmailClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "email")
	retryPolicy.ShouldRetryOperation = emailIpPoolWorkRequestShouldRetryFunc(timeout)

	response := oci_email.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_email.OperationStatusInProgress),
			string(oci_email.OperationStatusAccepted),
			string(oci_email.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_email.OperationStatusSucceeded),
			string(oci_email.OperationStatusFailed),
			string(oci_email.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_email.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_email.OperationStatusFailed || response.Status == oci_email.OperationStatusCanceled {
		return nil, getErrorFromEmailEmailIpPoolWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromEmailEmailIpPoolWorkRequest(ctx context.Context, client *oci_email.EmailClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_email.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_email.ListWorkRequestErrorsRequest{
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

func (s *EmailEmailIpPoolResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_email.GetEmailIpPoolRequest{}

	tmp := s.D.Id()
	request.EmailIpPoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.GetEmailIpPool(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.EmailIpPool
	return nil
}

func (s *EmailEmailIpPoolResourceCrud) UpdateWithContext(ctx context.Context) error {

	if s.D.HasChange("outbound_ips") {
		oldRaw, newRaw := s.D.GetChange("outbound_ips")

		oldList := expandStringList(oldRaw.([]interface{}))
		newList := expandStringList(newRaw.([]interface{}))

		toAdd, toRemove := diffStringSlices(oldList, newList)

		if len(toAdd) > 0 {
			if err := s.AddEmailOutboundIp(ctx, toAdd); err != nil {
				return err
			}
		}
		if len(toRemove) > 0 {
			if err := s.RemoveEmailOutboundIp(ctx, toRemove); err != nil {
				return err
			}
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_email.UpdateEmailIpPoolRequest{}

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

	tmp := s.D.Id()
	request.EmailIpPoolId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.UpdateEmailIpPool(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEmailIpPoolFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email"), oci_email.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func expandStringList(raw []interface{}) []string {
	out := make([]string, 0, len(raw))
	for _, v := range raw {
		if s, ok := v.(string); ok && s != "" {
			out = append(out, s)
		}
	}
	return out
}

func diffStringSlices(oldList, newList []string) (toAdd, toRemove []string) {
	oldSet := make(map[string]struct{}, len(oldList))
	newSet := make(map[string]struct{}, len(newList))

	for _, v := range oldList {
		oldSet[v] = struct{}{}
	}
	for _, v := range newList {
		newSet[v] = struct{}{}
	}

	// in new but not in old , then we will add
	for v := range newSet {
		if _, exists := oldSet[v]; !exists {
			toAdd = append(toAdd, v)
		}
	}

	// in old but not in new , then we will remove
	for v := range oldSet {
		if _, exists := newSet[v]; !exists {
			toRemove = append(toRemove, v)
		}
	}

	return
}

func (s *EmailEmailIpPoolResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_email.DeleteEmailIpPoolRequest{}

	tmp := s.D.Id()
	request.EmailIpPoolId = &tmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.DeleteEmailIpPool(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := emailIpPoolWaitForWorkRequest(ctx, workId, "emailippool",
		oci_email.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *EmailEmailIpPoolResourceCrud) SetData() error {
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, IpPoolResourceLockToMap(item))
	}
	s.D.Set("locks", locks)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	outboundIpsResponse := []interface{}{}
	for _, item := range s.Res.OutboundIps {
		outboundIpsResponse = append(outboundIpsResponse, IpPoolOutboundIpToMap(item))
	}
	s.D.Set("outbound_ips_response", outboundIpsResponse)

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

func IpPoolOutboundIpToMap(obj oci_email.EmailOutboundIpSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssignmentState != "" {
		result["assignment_state"] = string(obj.AssignmentState)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = *obj.LifecycleDetails
	}

	if obj.OutboundIp != nil {
		result["outbound_ip"] = *obj.OutboundIp
	}

	if obj.LifecycleState != "" {
		result["state"] = string(obj.LifecycleState)
	}

	return result
}

func (s *EmailEmailIpPoolResourceCrud) AddEmailOutboundIp(ctx context.Context, outboundIps []string) error {

	if len(outboundIps) == 0 {
		return nil
	}

	request := oci_email.AddEmailOutboundIpRequest{}

	idTmp := s.D.Id()
	request.EmailIpPoolId = &idTmp

	request.OutboundIps = outboundIps

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	_, err := s.Client.AddEmailOutboundIp(ctx, request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *EmailEmailIpPoolResourceCrud) RemoveEmailOutboundIp(ctx context.Context, outboundIps []string) error {

	if len(outboundIps) == 0 {
		return nil
	}

	request := oci_email.RemoveEmailOutboundIpRequest{}

	idTmp := s.D.Id()
	request.EmailIpPoolId = &idTmp

	// IPs to remove
	request.OutboundIps = outboundIps

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	_, err := s.Client.RemoveEmailOutboundIp(ctx, request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func EmailIpPoolSummaryToMap(obj oci_email.EmailIpPoolSummary) map[string]interface{} {
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

	locks := []interface{}{}
	for _, item := range obj.Locks {
		locks = append(locks, IpPoolResourceLockToMap(item))
	}
	result["locks"] = locks

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

func IpPoolResourceLockToMap(obj oci_email.ResourceLock) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	if obj.RelatedResourceId != nil {
		result["related_resource_id"] = string(*obj.RelatedResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *EmailEmailIpPoolResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_email.ChangeEmailIpPoolCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.EmailIpPoolId = &idTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.ChangeEmailIpPoolCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEmailIpPoolFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email"), oci_email.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
