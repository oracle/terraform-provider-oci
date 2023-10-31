// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicy,
		Read:     readNetworkFirewallNetworkFirewallPolicy,
		Update:   updateNetworkFirewallNetworkFirewallPolicy,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"attached_network_firewall_count": {
				Type:     schema.TypeInt,
				Computed: true,
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
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("apply_policy_trigger"); ok {
		err := sync.ApplyNetworkFirewallPolicy()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("clone_policy_trigger"); ok {
		err := sync.CloneNetworkFirewallPolicy()
		if err != nil {
			return err
		}
	}
	return nil
}

func readNetworkFirewallNetworkFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteNetworkFirewallNetworkFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.NetworkFirewallPolicy
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateCreating),
		string(oci_network_firewall.LifecycleStateAttaching),
	}
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateActive),
		string(oci_network_firewall.LifecycleStateNeedsAttention),
	}
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateDeleting),
		string(oci_network_firewall.LifecycleStateDetaching),
	}
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateDeleted),
	}
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) Create() error {
	request := oci_network_firewall.CreateNetworkFirewallPolicyRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getNetworkFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall"), oci_network_firewall.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) getNetworkFirewallPolicyFromWorkRequest(
	workId *string,
	retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_network_firewall.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	networkFirewallPolicyId, err := networkFirewallPolicyWaitForWorkRequest(workId, "networkfirewallpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, networkFirewallPolicyId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_network_firewall.CancelWorkRequestRequest{
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
	s.D.SetId(*networkFirewallPolicyId)

	return s.Get()
}

func networkFirewallPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "network_firewall", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_network_firewall.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func networkFirewallPolicyWaitForWorkRequest(wId *string, entityType string, action oci_network_firewall.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_network_firewall.NetworkFirewallClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "network_firewall")
	retryPolicy.ShouldRetryOperation = networkFirewallPolicyWorkRequestShouldRetryFunc(timeout)

	response := oci_network_firewall.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_network_firewall.OperationStatusInProgress),
			string(oci_network_firewall.OperationStatusAccepted),
			string(oci_network_firewall.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_network_firewall.OperationStatusSucceeded),
			string(oci_network_firewall.OperationStatusFailed),
			string(oci_network_firewall.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_network_firewall.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_network_firewall.OperationStatusFailed || response.Status == oci_network_firewall.OperationStatusCanceled {
		return nil, getErrorFromNetworkFirewallNetworkFirewallPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromNetworkFirewallNetworkFirewallPolicyWorkRequest(client *oci_network_firewall.NetworkFirewallClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_network_firewall.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_network_firewall.ListWorkRequestErrorsRequest{
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

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) Get() error {
	request := oci_network_firewall.GetNetworkFirewallPolicyRequest{}

	tmp := s.D.Id()
	request.NetworkFirewallPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkFirewallPolicy
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_network_firewall.UpdateNetworkFirewallPolicyRequest{}

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
	request.NetworkFirewallPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall"), oci_network_firewall.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteNetworkFirewallPolicyRequest{}

	tmp := s.D.Id()
	request.NetworkFirewallPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.DeleteNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := networkFirewallPolicyWaitForWorkRequest(workId, "networkfirewallpolicy",
		oci_network_firewall.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) SetData() error {
	if s.Res.AttachedNetworkFirewallCount != nil {
		s.D.Set("attached_network_firewall_count", *s.Res.AttachedNetworkFirewallCount)
	}

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

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) ApplyNetworkFirewallPolicy() error {
	request := oci_network_firewall.ApplyNetworkFirewallPolicyRequest{}

	if firewalls, ok := s.D.GetOkExists("firewalls"); ok {
		interfaces := firewalls.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("firewalls") {
			request.Firewalls = tmp
		}
	}

	idTmp := s.D.Id()
	request.NetworkFirewallPolicyId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.ApplyNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("apply_policy_trigger")
	s.D.Set("apply_policy_trigger", val)

	s.Res = &response.NetworkFirewallPolicy
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) CloneNetworkFirewallPolicy() error {
	request := oci_network_firewall.CloneNetworkFirewallPolicyRequest{}

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

	idTmp := s.D.Id()
	request.NetworkFirewallPolicyId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CloneNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("clone_policy_trigger")
	s.D.Set("clone_policy_trigger", val)

	s.Res = &response.NetworkFirewallPolicy
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) MigrateNetworkFirewallPolicy() error {
	request := oci_network_firewall.MigrateNetworkFirewallPolicyRequest{}

	idTmp := s.D.Id()
	request.NetworkFirewallPolicyId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.MigrateNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("migrate_trigger")
	s.D.Set("migrate_trigger", val)

	return nil
}

func NetworkFirewallPolicySummaryToMap(obj oci_network_firewall.NetworkFirewallPolicySummary) map[string]interface{} {
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
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_network_firewall.ChangeNetworkFirewallPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.NetworkFirewallPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.ChangeNetworkFirewallPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
