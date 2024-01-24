// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func NetworkFirewallNetworkFirewallResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("20m"),
		},
		Create: createNetworkFirewallNetworkFirewall,
		Read:   readNetworkFirewallNetworkFirewall,
		Update: updateNetworkFirewallNetworkFirewall,
		Delete: deleteNetworkFirewallNetworkFirewall,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_firewall_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"availability_domain": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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
			"ipv4address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ipv6address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"network_security_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createNetworkFirewallNetworkFirewall(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewall(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewall(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewall(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.NetworkFirewall
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateCreating),
		string(oci_network_firewall.LifecycleStateAttaching),
	}
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateActive),
		string(oci_network_firewall.LifecycleStateNeedsAttention),
	}
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateDeleting),
		string(oci_network_firewall.LifecycleStateDetaching),
	}
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateDeleted),
	}
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) Create() error {
	request := oci_network_firewall.CreateNetworkFirewallRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

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

	if ipv4Address, ok := s.D.GetOkExists("ipv4address"); ok {
		tmp := ipv4Address.(string)
		request.Ipv4Address = &tmp
	}

	if ipv6Address, ok := s.D.GetOkExists("ipv6address"); ok {
		tmp := ipv6Address.(string)
		request.Ipv6Address = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("network_security_group_ids") {
			request.NetworkSecurityGroupIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateNetworkFirewall(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getNetworkFirewallFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall"), oci_network_firewall.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) getNetworkFirewallFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_network_firewall.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	networkFirewallId, err := networkFirewallWaitForWorkRequest(workId, "networkfirewall",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)
	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, networkFirewallId)
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
	s.D.SetId(*networkFirewallId)

	return s.Get()
}

func networkFirewallWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

		return false
	}
}

func networkFirewallWaitForWorkRequest(wId *string, entityType string, action oci_network_firewall.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_network_firewall.NetworkFirewallClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "network_firewall")
	retryPolicy.ShouldRetryOperation = networkFirewallWorkRequestShouldRetryFunc(timeout)
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
			/*			if response.WorkRequest.OperationType == "CREATE_NETWORK_FIREWALL" && response.WorkRequest.Status == "SUCCEEDED" {
						for key, rsc := range response.WorkRequest.Resources {
							rsc.ActionType = "CREATED"
							response.WorkRequest.Resources[key].ActionType = "CREATED"
						}
					}*/
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Delay:        2 * time.Minute,
		PollInterval: 150 * time.Second,
		Timeout:      1 * time.Hour,
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
		return nil, getErrorFromNetworkFirewallNetworkFirewallWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromNetworkFirewallNetworkFirewallWorkRequest(client *oci_network_firewall.NetworkFirewallClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_network_firewall.ActionTypeEnum) error {
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

func (s *NetworkFirewallNetworkFirewallResourceCrud) Get() error {
	request := oci_network_firewall.GetNetworkFirewallRequest{}

	tmp := s.D.Id()
	request.NetworkFirewallId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetNetworkFirewall(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkFirewall
	return nil
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_network_firewall.UpdateNetworkFirewallRequest{}

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
	request.NetworkFirewallId = &tmp

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("network_security_group_ids") {
			request.NetworkSecurityGroupIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateNetworkFirewall(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkFirewallFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall"), oci_network_firewall.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteNetworkFirewallRequest{}

	tmp := s.D.Id()
	request.NetworkFirewallId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.DeleteNetworkFirewall(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := networkFirewallWaitForWorkRequest(workId, "networkfirewall",
		oci_network_firewall.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
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

	if s.Res.Ipv4Address != nil {
		s.D.Set("ipv4address", *s.Res.Ipv4Address)
	}

	if s.Res.Ipv6Address != nil {
		s.D.Set("ipv6address", *s.Res.Ipv6Address)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NetworkFirewallPolicyId != nil {
		s.D.Set("network_firewall_policy_id", *s.Res.NetworkFirewallPolicyId)
	}

	networkSecurityGroupIds := []interface{}{}
	for _, item := range s.Res.NetworkSecurityGroupIds {
		networkSecurityGroupIds = append(networkSecurityGroupIds, item)
	}
	s.D.Set("network_security_group_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, networkSecurityGroupIds))

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

	return nil
}

func NetworkFirewallSummaryToMap(obj oci_network_firewall.NetworkFirewallSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

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

	if obj.Ipv4Address != nil {
		result["ipv4address"] = string(*obj.Ipv4Address)
	}

	if obj.Ipv6Address != nil {
		result["ipv6address"] = string(*obj.Ipv6Address)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.NetworkFirewallPolicyId != nil {
		result["network_firewall_policy_id"] = string(*obj.NetworkFirewallPolicyId)
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

	return result
}

func (s *NetworkFirewallNetworkFirewallResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_network_firewall.ChangeNetworkFirewallCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.NetworkFirewallId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.ChangeNetworkFirewallCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkFirewallFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall"), oci_network_firewall.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
