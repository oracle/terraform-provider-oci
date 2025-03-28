// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"
)

func LoadBalancerLoadBalancerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLoadBalancerLoadBalancer,
		Read:     readLoadBalancerLoadBalancer,
		Update:   updateLoadBalancerLoadBalancer,
		Delete:   deleteLoadBalancerLoadBalancer,
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
			"shape": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"ip_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_delete_protection_enabled": {
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
			"is_request_id_enabled": {
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
			"request_id_header": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserved_ips": {
				Type:     schema.TypeList,
				Optional: true,
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
			"security_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"shape_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"maximum_bandwidth_in_mbps": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"minimum_bandwidth_in_mbps": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"ip_address_details": {
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
			"ip_addresses": {
				Type:       schema.TypeList,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedForAnother("ip_addresses", "ip_address_details"),
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
		},
	}
}

func createLoadBalancerLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readLoadBalancerLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateLoadBalancerLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoadBalancerLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoadBalancerLoadBalancerResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.LoadBalancer
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *LoadBalancerLoadBalancerResourceCrud) ID() string {
	id, workSuccess := loadBalancerResourceID(s.Res, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return *s.WorkRequest.LoadBalancerId
	}
	return ""
}

func (s *LoadBalancerLoadBalancerResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateCreating),
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerLoadBalancerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateActive),
		string(oci_load_balancer.LoadBalancerLifecycleStateFailed),
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerLoadBalancerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateDeleting),
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerLoadBalancerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateDeleted),
	}
}

func (s *LoadBalancerLoadBalancerResourceCrud) Create() error {
	request := oci_load_balancer.CreateLoadBalancerRequest{}

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

	if ipMode, ok := s.D.GetOkExists("ip_mode"); ok {
		request.IpMode = oci_load_balancer.CreateLoadBalancerDetailsIpModeEnum(ipMode.(string))
	}

	if isDeleteProtectionEnabled, ok := s.D.GetOkExists("is_delete_protection_enabled"); ok {
		tmp := isDeleteProtectionEnabled.(bool)
		request.IsDeleteProtectionEnabled = &tmp
	}

	if isPrivate, ok := s.D.GetOkExists("is_private"); ok {
		tmp := isPrivate.(bool)
		request.IsPrivate = &tmp
	}

	if isRequestIdEnabled, ok := s.D.GetOkExists("is_request_id_enabled"); ok {
		tmp := isRequestIdEnabled.(bool)
		request.IsRequestIdEnabled = &tmp
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

	if requestIdHeader, ok := s.D.GetOkExists("request_id_header"); ok {
		if requestIdHeader != nil {
			tmp := requestIdHeader.(string)
			request.RequestIdHeader = &tmp
		}
	}

	if reservedIps, ok := s.D.GetOkExists("reserved_ips"); ok {
		interfaces := reservedIps.([]interface{})
		tmp := make([]oci_load_balancer.ReservedIp, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "reserved_ips", stateDataIndex)
			converted, err := s.mapToReservedIP(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("reserved_ips") {
			request.ReservedIps = tmp
		}
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		convertedAttributes := tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		request.SecurityAttributes = convertedAttributes
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.ShapeName = &tmp
	}

	if shapeDetails, ok := s.D.GetOkExists("shape_details"); ok {
		if tmpList := shapeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shape_details", 0)
			tmp, err := s.mapToShapeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ShapeDetails = &tmp
		}
	}

	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok {
		interfaces := subnetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subnet_ids") {
			request.SubnetIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	var identifier *string
	identifier = workRequestResponse.LoadBalancerId
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerLoadBalancerResourceCrud) Get() error {
	id, stillWorking, err := loadBalancerResourceGet(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	if id == "" && s.WorkRequest != nil {
		id = *s.WorkRequest.LoadBalancerId
		s.D.SetId(id)
	}

	request := oci_load_balancer.GetLoadBalancerRequest{}

	tmp := s.D.Id()
	request.LoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LoadBalancer
	return nil
}

func (s *LoadBalancerLoadBalancerResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok && s.D.HasChange("shape") {
		oldRaw, newRaw := s.D.GetChange("shape")
		if newRaw != "" && oldRaw != "" {
			err := s.updateShape(shape)
			if err != nil {
				return err
			}
		}
	} else if _, ok := s.D.GetOkExists("shape_details"); ok && s.D.HasChange("shape_details") {
		if shape, ok := s.D.GetOkExists("shape"); ok {
			if shape == "flexible" {
				err := s.updateShape(shape)
				if err != nil {
					return err
				}
			}
		}
	}

	if s.D.HasChange("network_security_group_ids") {
		err := s.updateNetworkSecurityGroups()
		if err != nil {
			return fmt.Errorf("unable to update 'network_security_group_ids', error: %v", err)
		}
	}
	request := oci_load_balancer.UpdateLoadBalancerRequest{}

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

	if isDeleteProtectionEnabled, ok := s.D.GetOkExists("is_delete_protection_enabled"); ok {
		tmp := isDeleteProtectionEnabled.(bool)
		request.IsDeleteProtectionEnabled = &tmp
	}

	if isRequestIdEnabled, ok := s.D.GetOkExists("is_request_id_enabled"); ok {
		tmp := isRequestIdEnabled.(bool)
		request.IsRequestIdEnabled = &tmp
	}

	tmp := s.D.Id()
	request.LoadBalancerId = &tmp

	if requestIdHeader, ok := s.D.GetOkExists("request_id_header"); ok {
		tmp := requestIdHeader.(string)
		request.RequestIdHeader = &tmp
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		convertedAttributes := tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		request.SecurityAttributes = convertedAttributes
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *LoadBalancerLoadBalancerResourceCrud) Delete() error {
	request := oci_load_balancer.DeleteLoadBalancerRequest{}

	tmp := s.D.Id()
	request.LoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerLoadBalancerResourceCrud) SetData() error {
	if s.Res == nil || s.Res.Id == nil {
		return nil
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

	ipAddresses := []string{}
	ipMode := "IPV4"
	for _, ad := range s.Res.IpAddresses {
		if ad.IpAddress != nil {
			ipAddresses = append(ipAddresses, *ad.IpAddress)
		}
		tmp := *ad.IpAddress
		if !isIPV4(tmp) {
			ipMode = "IPV6"
		}
	}
	s.D.Set("ip_mode", ipMode)
	s.D.Set("ip_addresses", ipAddresses)

	ipAddressDetails := []interface{}{}

	for _, item := range s.Res.IpAddresses {
		ipAddressDetails = append(ipAddressDetails, IpAddressToMap(item))
	}

	s.D.Set("ip_address_details", ipAddressDetails)

	if s.Res.IsDeleteProtectionEnabled != nil {
		s.D.Set("is_delete_protection_enabled", *s.Res.IsDeleteProtectionEnabled)
	}

	if s.Res.IsPrivate != nil {
		s.D.Set("is_private", *s.Res.IsPrivate)
	}

	if s.Res.IsRequestIdEnabled != nil {
		s.D.Set("is_request_id_enabled", *s.Res.IsRequestIdEnabled)
	}

	networkSecurityGroupIds := []interface{}{}
	for _, item := range s.Res.NetworkSecurityGroupIds {
		networkSecurityGroupIds = append(networkSecurityGroupIds, item)
	}
	s.D.Set("network_security_group_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, networkSecurityGroupIds))

	if s.Res.RequestIdHeader != nil {
		s.D.Set("request_id_header", *s.Res.RequestIdHeader)
	}

	s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))

	if s.Res.ShapeName != nil {
		s.D.Set("shape", *s.Res.ShapeName)
	}

	if s.Res.ShapeDetails != nil {
		s.D.Set("shape_details", []interface{}{ShapeDetailsToMap(s.Res.ShapeDetails)})
	} else {
		s.D.Set("shape_details", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subnet_ids", s.Res.SubnetIds)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func IpAddressToMap(obj oci_load_balancer.IpAddress) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.IsPublic != nil {
		result["is_public"] = bool(*obj.IsPublic)
	}

	if obj.ReservedIp != nil {
		result["reserved_ip"] = []interface{}{ReservedIPToMap(*obj.ReservedIp)}
	}

	return result
}

func (s *LoadBalancerLoadBalancerResourceCrud) mapToReservedIP(fieldKeyFormat string) (oci_load_balancer.ReservedIp, error) {
	result := oci_load_balancer.ReservedIp{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func ReservedIPToMap(obj oci_load_balancer.ReservedIp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *LoadBalancerLoadBalancerResourceCrud) mapToShapeDetails(fieldKeyFormat string) (oci_load_balancer.ShapeDetails, error) {
	result := oci_load_balancer.ShapeDetails{}

	if maximumBandwidthInMbps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_bandwidth_in_mbps")); ok {
		tmp := maximumBandwidthInMbps.(int)
		result.MaximumBandwidthInMbps = &tmp
	}

	if minimumBandwidthInMbps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "minimum_bandwidth_in_mbps")); ok {
		tmp := minimumBandwidthInMbps.(int)
		result.MinimumBandwidthInMbps = &tmp
	}

	return result, nil
}

func ShapeDetailsToMap(obj *oci_load_balancer.ShapeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaximumBandwidthInMbps != nil {
		result["maximum_bandwidth_in_mbps"] = int(*obj.MaximumBandwidthInMbps)
	}

	if obj.MinimumBandwidthInMbps != nil {
		result["minimum_bandwidth_in_mbps"] = int(*obj.MinimumBandwidthInMbps)
	}

	return result
}

func isIPV4(ipAddress string) bool {
	return strings.Contains(ipAddress, ".")
}

func (s *LoadBalancerLoadBalancerResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_load_balancer.ChangeLoadBalancerCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.LoadBalancerId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	_, err := s.Client.ChangeLoadBalancerCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *LoadBalancerLoadBalancerResourceCrud) updateNetworkSecurityGroups() error {
	updateNsgIdsRequest := oci_load_balancer.UpdateNetworkSecurityGroupsRequest{}

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
	updateNsgIdsRequest.LoadBalancerId = &tmp

	updateNsgIdsRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateNetworkSecurityGroups(context.Background(), updateNsgIdsRequest)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerLoadBalancerResourceCrud) updateShape(shape interface{}) error {
	changeShapeRequest := oci_load_balancer.UpdateLoadBalancerShapeRequest{}

	shapeTmp := shape.(string)
	changeShapeRequest.ShapeName = &shapeTmp

	idTmp := s.D.Id()
	changeShapeRequest.LoadBalancerId = &idTmp

	changeShapeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	if shapeDetails, ok := s.D.GetOkExists("shape_details"); ok {
		if tmpList := shapeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shape_details", 0)
			tmp, err := s.mapToShapeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			changeShapeRequest.ShapeDetails = &tmp
		}
	}

	response, err := s.Client.UpdateLoadBalancerShape(context.Background(), changeShapeRequest)
	if err != nil {
		return err
	}
	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}
