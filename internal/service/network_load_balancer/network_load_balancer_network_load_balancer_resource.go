// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_load_balancer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"
)

func NetworkLoadBalancerNetworkLoadBalancerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkLoadBalancerNetworkLoadBalancer,
		Read:     readNetworkLoadBalancerNetworkLoadBalancer,
		Update:   updateNetworkLoadBalancerNetworkLoadBalancer,
		Delete:   deleteNetworkLoadBalancerNetworkLoadBalancer,
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

			// Optional
			"assigned_ipv6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"assigned_private_ipv4": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
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
			"is_preserve_source_destination": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_private": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_symmetric_hash_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"network_security_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"nlb_ip_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserved_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"subnet_ipv6cidr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},

			// Computed
			"ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_public": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"reserved_ip": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
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

func createNetworkLoadBalancerNetworkLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkLoadBalancerNetworkLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkLoadBalancerNetworkLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkLoadBalancerNetworkLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkLoadBalancerNetworkLoadBalancerResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_load_balancer.NetworkLoadBalancerClient
	Res                    *oci_network_load_balancer.NetworkLoadBalancer
	DisableNotFoundRetries bool
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_network_load_balancer.LifecycleStateCreating),
	}
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_network_load_balancer.LifecycleStateActive),
		string(oci_network_load_balancer.LifecycleStateFailed),
	}
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_network_load_balancer.LifecycleStateDeleting),
	}
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_network_load_balancer.LifecycleStateDeleted),
	}
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) Create() error {
	request := oci_network_load_balancer.CreateNetworkLoadBalancerRequest{}

	if assignedIpv6, ok := s.D.GetOkExists("assigned_ipv6"); ok {
		tmp := assignedIpv6.(string)
		request.AssignedIpv6 = &tmp
	}

	if assignedPrivateIpv4, ok := s.D.GetOkExists("assigned_private_ipv4"); ok {
		tmp := assignedPrivateIpv4.(string)
		request.AssignedPrivateIpv4 = &tmp
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

	if isPreserveSourceDestination, ok := s.D.GetOkExists("is_preserve_source_destination"); ok {
		tmp := isPreserveSourceDestination.(bool)
		request.IsPreserveSourceDestination = &tmp
	}

	if isPrivate, ok := s.D.GetOkExists("is_private"); ok {
		tmp := isPrivate.(bool)
		request.IsPrivate = &tmp
	}

	if isSymmetricHashEnabled, ok := s.D.GetOkExists("is_symmetric_hash_enabled"); ok {
		tmp := isSymmetricHashEnabled.(bool)
		request.IsSymmetricHashEnabled = &tmp
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
	if nlbIpVersion, ok := s.D.GetOkExists("nlb_ip_version"); ok {
		request.NlbIpVersion = oci_network_load_balancer.NlbIpVersionEnum(nlbIpVersion.(string))
	}
	if reservedIps, ok := s.D.GetOkExists("reserved_ips"); ok {
		interfaces := reservedIps.([]interface{})
		tmp := make([]oci_network_load_balancer.ReservedIp, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "reserved_ips", stateDataIndex)
			converted, err := s.mapToNetworkLoadBalancerReservedIp(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("reserved_ips") {
			request.ReservedIps = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if subnetIpv6Cidr, ok := s.D.GetOkExists("subnet_ipv6cidr"); ok {
		tmp := subnetIpv6Cidr.(string)
		request.SubnetIpv6Cidr = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.CreateNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getNetworkLoadBalancerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) getNetworkLoadBalancerFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_network_load_balancer.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	networkLoadBalancerId, err := networkLoadBalancerWaitForWorkRequest(workId,
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*networkLoadBalancerId)

	return s.Get()
}

func networkLoadBalancerWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "network_load_balancer", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_network_load_balancer.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func networkLoadBalancerWaitForWorkRequest(wId *string, action oci_network_load_balancer.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_network_load_balancer.NetworkLoadBalancerClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "network_load_balancer")
	retryPolicy.ShouldRetryOperation = networkLoadBalancerWorkRequestShouldRetryFunc(timeout)

	response := oci_network_load_balancer.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_network_load_balancer.OperationStatusInProgress),
			string(oci_network_load_balancer.OperationStatusAccepted),
			string(oci_network_load_balancer.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_network_load_balancer.OperationStatusSucceeded),
			string(oci_network_load_balancer.OperationStatusFailed),
			string(oci_network_load_balancer.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_network_load_balancer.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), "networkloadbalancer") {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	var workRequestErr error
	if response.WorkRequest.Status == oci_network_load_balancer.OperationStatusFailed {
		errorMessage := getErrorFromNetworkLoadBalancerWorkRequest(response.WorkRequest, client, retryPolicy)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, action: %s. Message: %s", *wId, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromNetworkLoadBalancerWorkRequest(wr oci_network_load_balancer.WorkRequest,
	client *oci_network_load_balancer.NetworkLoadBalancerClient, retryPolicy *oci_common.RetryPolicy) string {
	// Fetch the list of work request errors
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_network_load_balancer.ListWorkRequestErrorsRequest{
			WorkRequestId: wr.Id,
			CompartmentId: wr.CompartmentId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})

	if err != nil {
		return "Unknown failure reason"
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.WorkRequestErrorCollection.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")
	return errorMessage
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) Get() error {
	request := oci_network_load_balancer.GetNetworkLoadBalancerRequest{}

	tmp := s.D.Id()
	request.NetworkLoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.GetNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkLoadBalancer
	return nil
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	if s.D.HasChange("network_security_group_ids") {
		err := s.updateNetworkSecurityGroups()
		if err != nil {
			return fmt.Errorf("unable to update 'network_security_group_ids', error: %v", err)
		}
	}

	request := oci_network_load_balancer.UpdateNetworkLoadBalancerRequest{}

	if assignedIpv6, ok := s.D.GetOkExists("assigned_ipv6"); ok &&
		s.D.HasChange("assigned_ipv6") {
		tmp := assignedIpv6.(string)
		request.AssignedIpv6 = &tmp
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

	if isPreserveSourceDestination, ok := s.D.GetOkExists("is_preserve_source_destination"); ok {
		tmp := isPreserveSourceDestination.(bool)
		request.IsPreserveSourceDestination = &tmp
	}

	if isSymmetricHashEnabled, ok := s.D.GetOkExists("is_symmetric_hash_enabled"); ok {
		tmp := isSymmetricHashEnabled.(bool)
		request.IsSymmetricHashEnabled = &tmp
	}

	tmp := s.D.Id()
	request.NetworkLoadBalancerId = &tmp
	if nlbIpVersion, ok := s.D.GetOkExists("nlb_ip_version"); ok {
		request.NlbIpVersion = oci_network_load_balancer.NlbIpVersionEnum(nlbIpVersion.(string))
	}

	if subnetIpv6Cidr, ok := s.D.GetOkExists("subnet_ipv6cidr"); ok &&
		s.D.HasChange("subnet_ipv6cidr") {
		tmp := subnetIpv6Cidr.(string)
		request.SubnetIpv6Cidr = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.UpdateNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkLoadBalancerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) Delete() error {
	request := oci_network_load_balancer.DeleteNetworkLoadBalancerRequest{}

	tmp := s.D.Id()
	request.NetworkLoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.DeleteNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := networkLoadBalancerWaitForWorkRequest(workId,
		oci_network_load_balancer.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) SetData() error {
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

	ipAddresses := []interface{}{}
	for _, item := range s.Res.IpAddresses {
		ipAddresses = append(ipAddresses, NetworkLoadBalancerIpAddressToMap(item))
	}
	s.D.Set("ip_addresses", ipAddresses)

	if s.Res.IsPreserveSourceDestination != nil {
		s.D.Set("is_preserve_source_destination", *s.Res.IsPreserveSourceDestination)
	}

	if s.Res.IsPrivate != nil {
		s.D.Set("is_private", *s.Res.IsPrivate)
	}

	if s.Res.IsSymmetricHashEnabled != nil {
		s.D.Set("is_symmetric_hash_enabled", *s.Res.IsSymmetricHashEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	networkSecurityGroupIds := []interface{}{}
	for _, item := range s.Res.NetworkSecurityGroupIds {
		networkSecurityGroupIds = append(networkSecurityGroupIds, item)
	}
	s.D.Set("network_security_group_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, networkSecurityGroupIds))
	s.D.Set("nlb_ip_version", s.Res.NlbIpVersion)
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

func NetworkLoadBalancerIpAddressToMap(obj oci_network_load_balancer.IpAddress) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}
	result["ip_version"] = string(obj.IpVersion)
	if obj.IsPublic != nil {
		result["is_public"] = bool(*obj.IsPublic)
	}

	if obj.ReservedIp != nil {
		result["reserved_ip"] = []interface{}{NetworkLoadBalancerReservedIpToMap(*obj.ReservedIp)}
	}

	return result
}

func NetworkLoadBalancerSummaryToMap(obj oci_network_load_balancer.NetworkLoadBalancerSummary, datasource bool) map[string]interface{} {
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

	ipAddresses := []interface{}{}
	for _, item := range obj.IpAddresses {
		ipAddresses = append(ipAddresses, NetworkLoadBalancerIpAddressToMap(item))
	}
	result["ip_addresses"] = ipAddresses

	if obj.IsPreserveSourceDestination != nil {
		result["is_preserve_source_destination"] = bool(*obj.IsPreserveSourceDestination)
	}

	if obj.IsPrivate != nil {
		result["is_private"] = bool(*obj.IsPrivate)
	}

	if obj.IsSymmetricHashEnabled != nil {
		result["is_symmetric_hash_enabled"] = bool(*obj.IsSymmetricHashEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	networkSecurityGroupIds := []interface{}{}
	for _, item := range obj.NetworkSecurityGroupIds {
		networkSecurityGroupIds = append(networkSecurityGroupIds, item)
	}
	if datasource {
		result["network_security_group_ids"] = networkSecurityGroupIds
	} else {
		result["network_security_group_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, networkSecurityGroupIds)
	}
	result["nlb_ip_version"] = string(obj.NlbIpVersion)
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

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) mapToNetworkLoadBalancerReservedIp(fieldKeyFormat string) (oci_network_load_balancer.ReservedIp, error) {
	result := oci_network_load_balancer.ReservedIp{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func NetworkLoadBalancerReservedIpToMap(obj oci_network_load_balancer.ReservedIp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) updateNetworkSecurityGroups() error {
	updateNsgIdsRequest := oci_network_load_balancer.UpdateNetworkSecurityGroupsRequest{}

	//@Codegen: Unless explicitly specified by the user, network_security_group_ids will not be supplied as the feature may or may not be supported
	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		updateNsgIdsRequest.NetworkSecurityGroupIds = tmp
	}

	tmp := s.D.Id()
	updateNsgIdsRequest.NetworkLoadBalancerId = &tmp

	updateNsgIdsRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.UpdateNetworkSecurityGroups(context.Background(), updateNsgIdsRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkLoadBalancerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_network_load_balancer.ChangeNetworkLoadBalancerCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.NetworkLoadBalancerId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.ChangeNetworkLoadBalancerCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkLoadBalancerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
