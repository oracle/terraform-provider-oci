// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreVirtualCircuitResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreVirtualCircuit,
		Read:     readCoreVirtualCircuit,
		Update:   updateCoreVirtualCircuit,
		Delete:   deleteCoreVirtualCircuit,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"bandwidth_shape_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bgp_admin_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cross_connect_mappings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"bgp_md5auth_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cross_connect_or_cross_connect_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							Default:  nil,
						},
						"customer_bgp_peering_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"customer_bgp_peering_ipv6": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: ipv6CompressionDiffSuppressFunction,
						},
						"oracle_bgp_peering_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oracle_bgp_peering_ipv6": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: ipv6CompressionDiffSuppressFunction,
						},
						"vlan": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							Default:  nil,
						},

						// Computed
					},
				},
			},
			"customer_asn": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
				ConflictsWith:    []string{"customer_bgp_asn"},
			},
			"customer_bgp_asn": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"customer_asn"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("customer_bgp_asn", "customer_asn"),
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
			"gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_mtu": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_bfd_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_transport_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"provider_service_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"provider_service_key_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"public_prefixes": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      publicPrefixesHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"cidr_block": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"routing_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"bgp_ipv6session_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bgp_management": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedButSupportedThroughAnotherDataSource("bgp_management", "oci_core_fast_connect_provider_service"),
			},
			"bgp_session_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_bgp_asn": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"provider_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reference_comment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_circuit_redundancy_metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"configured_redundancy_level": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv4bgp_session_redundancy_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6bgp_session_redundancy_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createCoreVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreVirtualCircuitResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.VirtualCircuit
	DisableNotFoundRetries bool
}

func (s *CoreVirtualCircuitResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVirtualCircuitResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStateVerifying),
		string(oci_core.VirtualCircuitLifecycleStateProvisioning),
	}
}

func (s *CoreVirtualCircuitResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStatePendingProvider),
		string(oci_core.VirtualCircuitLifecycleStateProvisioned),
	}
}

func (s *CoreVirtualCircuitResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStateTerminating),
	}
}

func (s *CoreVirtualCircuitResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStateTerminated),
	}
}

func (s *CoreVirtualCircuitResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStateProvisioning),
	}
}

func (s *CoreVirtualCircuitResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStatePendingProvider),
		string(oci_core.VirtualCircuitLifecycleStateProvisioned),
	}
}

func (s *CoreVirtualCircuitResourceCrud) Create() error {
	request := oci_core.CreateVirtualCircuitRequest{}

	if bandwidthShapeName, ok := s.D.GetOkExists("bandwidth_shape_name"); ok {
		tmp := bandwidthShapeName.(string)
		request.BandwidthShapeName = &tmp
	}

	if bgpAdminState, ok := s.D.GetOkExists("bgp_admin_state"); ok {
		request.BgpAdminState = oci_core.CreateVirtualCircuitDetailsBgpAdminStateEnum(bgpAdminState.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if crossConnectMappings, ok := s.D.GetOkExists("cross_connect_mappings"); ok {
		interfaces := crossConnectMappings.([]interface{})
		tmp := make([]oci_core.CrossConnectMapping, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cross_connect_mappings", stateDataIndex)
			converted, err := s.mapToCrossConnectMapping(fieldKeyFormat, false)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("cross_connect_mappings") {
			request.CrossConnectMappings = tmp
		}
	}

	if customerAsn, ok := s.D.GetOkExists("customer_asn"); ok {
		tmp := customerAsn.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert customerAsn string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.CustomerAsn = &tmpInt64
	}

	if customerBgpAsn, ok := s.D.GetOkExists("customer_bgp_asn"); ok {
		tmp := int64(customerBgpAsn.(int))
		request.CustomerAsn = &tmp
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

	if gatewayId, ok := s.D.GetOkExists("gateway_id"); ok {
		tmp := gatewayId.(string)
		request.GatewayId = &tmp
	}

	if ipMtu, ok := s.D.GetOkExists("ip_mtu"); ok {
		request.IpMtu = oci_core.VirtualCircuitIpMtuEnum(ipMtu.(string))
	}

	if isBfdEnabled, ok := s.D.GetOkExists("is_bfd_enabled"); ok {
		tmp := isBfdEnabled.(bool)
		request.IsBfdEnabled = &tmp
	}

	if vcType, ok := s.D.GetOkExists("type"); ok && !strings.EqualFold(vcType.(string), string(oci_core.VirtualCircuitTypePublic)) {
		if isTransportMode, ok := s.D.GetOkExists("is_transport_mode"); ok {
			tmp := isTransportMode.(bool)
			request.IsTransportMode = &tmp
		}
	}

	if providerServiceId, ok := s.D.GetOkExists("provider_service_id"); ok {
		tmp := providerServiceId.(string)
		request.ProviderServiceId = &tmp
	}

	if providerServiceKeyName, ok := s.D.GetOkExists("provider_service_key_name"); ok {
		tmp := providerServiceKeyName.(string)
		request.ProviderServiceKeyName = &tmp
	}

	if publicPrefixes, ok := s.D.GetOkExists("public_prefixes"); ok {
		set := publicPrefixes.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.CreateVirtualCircuitPublicPrefixDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := publicPrefixesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "public_prefixes", stateDataIndex)
			converted, err := s.mapToCreateVirtualCircuitPublicPrefixDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("public_prefixes") {
			request.PublicPrefixes = tmp
		}
	}

	// Virtual Circuit of type 'PRIVATE' does not support publicPrefixes in payload
	if len(request.PublicPrefixes) == 0 {
		request.PublicPrefixes = nil
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if routingPolicy, ok := s.D.GetOkExists("routing_policy"); ok {
		interfaces := routingPolicy.([]interface{})
		tmp := make([]oci_core.CreateVirtualCircuitDetailsRoutingPolicyEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_core.CreateVirtualCircuitDetailsRoutingPolicyEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("routing_policy") {
			request.RoutingPolicy = tmp
		}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_core.CreateVirtualCircuitDetailsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVirtualCircuit(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VirtualCircuit
	return nil
}

func (s *CoreVirtualCircuitResourceCrud) Get() error {
	request := oci_core.GetVirtualCircuitRequest{}

	tmp := s.D.Id()
	request.VirtualCircuitId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVirtualCircuit(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VirtualCircuit

	ppRequest := oci_core.ListVirtualCircuitPublicPrefixesRequest{}
	ppRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
	ppRequest.VirtualCircuitId = request.VirtualCircuitId

	ppResponse, ppErr := s.Client.ListVirtualCircuitPublicPrefixes(context.Background(), ppRequest)
	if ppErr != nil {
		return ppErr
	}

	publicPrefixes := []string{}
	for _, item := range ppResponse.Items {
		publicPrefixes = append(publicPrefixes, *item.CidrBlock)
	}

	s.Res.PublicPrefixes = publicPrefixes

	return nil
}

func (s *CoreVirtualCircuitResourceCrud) Update() error {
	// Update public prefixes, if changed
	// Cannot Update PublicPrefix when the VirtualCircuit is in state PROVISIONING so public prefixes should be updated first
	if s.D.HasChange("public_prefixes") {
		err := s.updatePublicPrefixes()
		if err != nil {
			return fmt.Errorf("unable to Update 'public_prefixes', error: %v", err)
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}

			if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
				return waitErr
			}
		}
	}

	request := oci_core.UpdateVirtualCircuitRequest{}

	if bandwidthShapeName, ok := s.D.GetOkExists("bandwidth_shape_name"); ok && s.D.HasChange("bandwidth_shape_name") {
		tmp := bandwidthShapeName.(string)
		request.BandwidthShapeName = &tmp
	}

	if bgpAdminState, ok := s.D.GetOkExists("bgp_admin_state"); ok {
		request.BgpAdminState = oci_core.UpdateVirtualCircuitDetailsBgpAdminStateEnum(bgpAdminState.(string))
	}

	if crossConnectMappings, ok := s.D.GetOkExists("cross_connect_mappings"); ok && s.D.HasChange("cross_connect_mappings") {
		interfaces := crossConnectMappings.([]interface{})
		tmp := make([]oci_core.CrossConnectMapping, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cross_connect_mappings", stateDataIndex)
			converted, err := s.mapToCrossConnectMapping(fieldKeyFormat, true)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("cross_connect_mappings") {
			request.CrossConnectMappings = tmp
		}
	}

	if customerAsn, ok := s.D.GetOkExists("customer_asn"); ok && s.D.HasChange("customer_asn") {
		tmp := customerAsn.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert customerAsn string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.CustomerAsn = &tmpInt64
	}

	if customerBgpAsn, ok := s.D.GetOkExists("customer_bgp_asn"); ok && s.D.HasChange("customer_bgp_asn") {
		tmp := int64(customerBgpAsn.(int))
		request.CustomerAsn = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok && s.D.HasChange("display_name") {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gatewayId, ok := s.D.GetOkExists("gateway_id"); ok && s.D.HasChange("gateway_id") {
		tmp := gatewayId.(string)
		request.GatewayId = &tmp
	}

	if ipMtu, ok := s.D.GetOkExists("ip_mtu"); ok && s.D.HasChange("ip_mtu") {
		request.IpMtu = oci_core.VirtualCircuitIpMtuEnum(ipMtu.(string))
	}

	if isBfdEnabled, ok := s.D.GetOkExists("is_bfd_enabled"); ok {
		tmp := isBfdEnabled.(bool)
		request.IsBfdEnabled = &tmp
	}

	if vcType, ok := s.D.GetOkExists("type"); ok && !strings.EqualFold(vcType.(string), string(oci_core.VirtualCircuitTypePublic)) {
		if isTransportMode, ok := s.D.GetOkExists("is_transport_mode"); ok {
			tmp := isTransportMode.(bool)
			request.IsTransportMode = &tmp
		}
	}

	if providerServiceKeyName, ok := s.D.GetOkExists("provider_service_key_name"); ok && s.D.HasChange("provider_service_key_name") {
		tmp := providerServiceKeyName.(string)
		request.ProviderServiceKeyName = &tmp
	}

	// @CODEGEN - 20190315 - provider_state can only be updated by Fast Connect Providers

	if referenceComment, ok := s.D.GetOkExists("reference_comment"); ok && s.D.HasChange("reference_comment") {
		tmp := referenceComment.(string)
		request.ReferenceComment = &tmp
	}

	if routingPolicy, ok := s.D.GetOkExists("routing_policy"); ok && s.D.HasChange("routing_policy") {
		interfaces := routingPolicy.([]interface{})
		tmp := make([]oci_core.UpdateVirtualCircuitDetailsRoutingPolicyEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_core.UpdateVirtualCircuitDetailsRoutingPolicyEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("routing_policy") {
			request.RoutingPolicy = tmp
		}
	}

	tmp := s.D.Id()
	request.VirtualCircuitId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVirtualCircuit(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VirtualCircuit
	return nil
}

// Update public prefixes using BulkAdd and BulkDelete APIs by computing the diff
func (s *CoreVirtualCircuitResourceCrud) updatePublicPrefixes() error {
	virtualCircuitId := s.D.Id()

	o, n := s.D.GetChange("public_prefixes")
	if o == nil {
		o = new(schema.Set)
	}
	if n == nil {
		n = new(schema.Set)
	}

	os := o.(*schema.Set)
	ns := n.(*schema.Set)

	newPublicPrefixesToAdd := ns.Difference(os).List()
	oldPublicPrefixesToDelete := os.Difference(ns).List()

	var publicPrefixesToAdd []oci_core.CreateVirtualCircuitPublicPrefixDetails
	var publicPrefixesToDelete []oci_core.DeleteVirtualCircuitPublicPrefixDetails

	for _, nppMap := range newPublicPrefixesToAdd {
		npp := mapToCreateVirtualCircuitPublicPrefixDetails(nppMap.(map[string]interface{}))
		publicPrefixesToAdd = append(publicPrefixesToAdd, npp)
	}

	for _, oppMap := range oldPublicPrefixesToDelete {
		opp := mapToDeleteVirtualCircuitPublicPrefixDetails(oppMap.(map[string]interface{}))
		publicPrefixesToDelete = append(publicPrefixesToDelete, opp)
	}

	// Add the public prefixes first, if any
	// And, wait for the Update to complete before proceeding for subsequent updates if state is PROVISIONING
	if len(publicPrefixesToAdd) > 0 {
		addRequest := oci_core.BulkAddVirtualCircuitPublicPrefixesRequest{}
		addRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
		addRequest.PublicPrefixes = publicPrefixesToAdd
		addRequest.VirtualCircuitId = &virtualCircuitId
		_, addErr := s.Client.BulkAddVirtualCircuitPublicPrefixes(context.Background(), addRequest)
		if addErr != nil {
			return fmt.Errorf("failed to add public prefixes, error: %v", addErr)
		}

		if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
			return waitErr
		}
	}

	// Delete the old public prefixes, if any, after adding new ones
	// And, wait for the Update to complete before proceeding for subsequent updates if state is PROVISIONING
	if len(publicPrefixesToDelete) > 0 {
		deleteRequest := oci_core.BulkDeleteVirtualCircuitPublicPrefixesRequest{}
		deleteRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
		deleteRequest.PublicPrefixes = publicPrefixesToDelete
		deleteRequest.VirtualCircuitId = &virtualCircuitId
		_, deleteErr := s.Client.BulkDeleteVirtualCircuitPublicPrefixes(context.Background(), deleteRequest)
		if deleteErr != nil {
			return fmt.Errorf("failed to delete public prefixes, error: %v", deleteErr)
		}

		if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
			return waitErr
		}
	}

	return nil
}

func (s *CoreVirtualCircuitResourceCrud) Delete() error {
	request := oci_core.DeleteVirtualCircuitRequest{}

	tmp := s.D.Id()
	request.VirtualCircuitId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVirtualCircuit(context.Background(), request)
	return err
}

func (s *CoreVirtualCircuitResourceCrud) SetData() error {
	if s.Res.BandwidthShapeName != nil {
		s.D.Set("bandwidth_shape_name", *s.Res.BandwidthShapeName)
	}

	s.D.Set("bgp_admin_state", s.Res.BgpAdminState)

	s.D.Set("bgp_ipv6session_state", s.Res.BgpIpv6SessionState)

	s.D.Set("bgp_management", s.Res.BgpManagement)

	s.D.Set("bgp_session_state", s.Res.BgpSessionState)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	crossConnectMappings := []interface{}{}
	for _, item := range s.Res.CrossConnectMappings {
		crossConnectMappings = append(crossConnectMappings, CrossConnectMappingToMap(item))
	}
	s.D.Set("cross_connect_mappings", crossConnectMappings)

	if s.Res.CustomerAsn != nil {
		s.D.Set("customer_asn", strconv.FormatInt(*s.Res.CustomerAsn, 10))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GatewayId != nil {
		s.D.Set("gateway_id", *s.Res.GatewayId)
	}

	s.D.Set("ip_mtu", s.Res.IpMtu)

	if s.Res.IsBfdEnabled != nil {
		s.D.Set("is_bfd_enabled", *s.Res.IsBfdEnabled)
	}

	if s.Res.IsTransportMode != nil {
		s.D.Set("is_transport_mode", *s.Res.IsTransportMode)
	}

	if s.Res.OracleBgpAsn != nil {
		s.D.Set("oracle_bgp_asn", *s.Res.OracleBgpAsn)
	}

	if s.Res.ProviderServiceId != nil {
		s.D.Set("provider_service_id", *s.Res.ProviderServiceId)
	}

	if s.Res.ProviderServiceKeyName != nil {
		s.D.Set("provider_service_key_name", *s.Res.ProviderServiceKeyName)
	}

	s.D.Set("provider_state", s.Res.ProviderState)

	publicPrefixes := []interface{}{}
	for _, item := range s.Res.PublicPrefixes {
		publicPrefixes = append(publicPrefixes, CreateVirtualCircuitPublicPrefixDetailsToMap(item))
	}
	s.D.Set("public_prefixes", schema.NewSet(publicPrefixesHashCodeForSets, publicPrefixes))

	if s.Res.ReferenceComment != nil {
		s.D.Set("reference_comment", *s.Res.ReferenceComment)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	s.D.Set("routing_policy", s.Res.RoutingPolicy)

	s.D.Set("service_type", s.Res.ServiceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.VirtualCircuitRedundancyMetadata != nil {
		s.D.Set("virtual_circuit_redundancy_metadata", []interface{}{VirtualCircuitRedundancyMetadataToMap(s.Res.VirtualCircuitRedundancyMetadata)})
	} else {
		s.D.Set("virtual_circuit_redundancy_metadata", nil)
	}

	return nil
}

// Converting raw set data from state diff to DeleteVirtualCircuitPublicPrefixDetails
func mapToDeleteVirtualCircuitPublicPrefixDetails(publicPrefix map[string]interface{}) oci_core.DeleteVirtualCircuitPublicPrefixDetails {
	result := oci_core.DeleteVirtualCircuitPublicPrefixDetails{}

	if cidrBlock, ok := publicPrefix["cidr_block"]; ok {
		tmp := cidrBlock.(string)
		result.CidrBlock = &tmp
	}

	return result
}

// Converting raw set data from state diff to CreateVirtualCircuitPublicPrefixDetails
func mapToCreateVirtualCircuitPublicPrefixDetails(publicPrefix map[string]interface{}) oci_core.CreateVirtualCircuitPublicPrefixDetails {
	result := oci_core.CreateVirtualCircuitPublicPrefixDetails{}

	if cidrBlock, ok := publicPrefix["cidr_block"]; ok {
		tmp := cidrBlock.(string)
		result.CidrBlock = &tmp
	}

	return result
}

func (s *CoreVirtualCircuitResourceCrud) mapToCreateVirtualCircuitPublicPrefixDetails(fieldKeyFormat string) (oci_core.CreateVirtualCircuitPublicPrefixDetails, error) {
	result := oci_core.CreateVirtualCircuitPublicPrefixDetails{}

	if cidrBlock, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cidr_block")); ok {
		tmp := cidrBlock.(string)
		result.CidrBlock = &tmp
	}

	return result, nil
}

func CreateVirtualCircuitPublicPrefixDetailsToMap(obj string) map[string]interface{} {
	result := map[string]interface{}{}

	result["cidr_block"] = obj

	return result
}

func (s *CoreVirtualCircuitResourceCrud) mapToCrossConnectMapping(fieldKeyFormat string, isUpdate bool) (oci_core.CrossConnectMapping, error) {
	result := oci_core.CrossConnectMapping{}
	_, hasProviderId := s.D.GetOkExists("provider_service_id")
	_, hasProviderName := s.D.GetOkExists("provider_service_name")
	isProvider := hasProviderId || hasProviderName

	// Do not include default empty bgp_md5auth_key in request payload unless it has changed
	if bgpMd5AuthKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bgp_md5auth_key")); ok &&
		(bgpMd5AuthKey != "" || s.D.HasChange("bgp_md5auth_key")) {
		tmp := bgpMd5AuthKey.(string)
		result.BgpMd5AuthKey = &tmp
	}

	// Do not include default empty cross_connect_or_cross_connect_group_id in request payload unless it has changed
	if crossConnectOrCrossConnectGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cross_connect_or_cross_connect_group_id")); ok &&
		((!isUpdate && crossConnectOrCrossConnectGroupId != "") || !isProvider) ||
		(isUpdate && s.D.HasChange("cross_connect_or_cross_connect_group_id")) {
		tmp := crossConnectOrCrossConnectGroupId.(string)
		result.CrossConnectOrCrossConnectGroupId = &tmp
	}

	// customer_bgp_peering_ip, oracleBgpPeeringIp are not allowed in requests for PUBLIC virtual circuit
	if vcType, ok := s.D.GetOkExists("type"); ok && !strings.EqualFold(vcType.(string), string(oci_core.VirtualCircuitTypePublic)) {
		if customerBgpPeeringIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_bgp_peering_ip")); ok {
			tmp := customerBgpPeeringIp.(string)
			result.CustomerBgpPeeringIp = &tmp
		}

		if customerBgpPeeringIpv6, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_bgp_peering_ipv6")); ok {
			tmp := customerBgpPeeringIpv6.(string)
			if tmp != "" {
				result.CustomerBgpPeeringIpv6 = &tmp
			}
		}

		if oracleBgpPeeringIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oracle_bgp_peering_ip")); ok {
			tmp := oracleBgpPeeringIp.(string)
			result.OracleBgpPeeringIp = &tmp
		}

		if oracleBgpPeeringIpv6, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oracle_bgp_peering_ipv6")); ok {
			tmp := oracleBgpPeeringIpv6.(string)
			if tmp != "" {
				result.OracleBgpPeeringIpv6 = &tmp
			}
		}
	}

	if vlan, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vlan")); ok {
		tmp := vlan.(int)
		// Do not include default 0 vlan in request payload unless it has changed
		if ((!isUpdate && tmp > 0) || !isProvider) || (isUpdate && s.D.HasChange("vlan")) {
			result.Vlan = &tmp
		}
	}

	return result, nil
}

func CrossConnectMappingToMap(obj oci_core.CrossConnectMapping) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BgpMd5AuthKey != nil {
		result["bgp_md5auth_key"] = string(*obj.BgpMd5AuthKey)
	}

	if obj.CrossConnectOrCrossConnectGroupId != nil {
		result["cross_connect_or_cross_connect_group_id"] = string(*obj.CrossConnectOrCrossConnectGroupId)
	}

	if obj.CustomerBgpPeeringIp != nil {
		result["customer_bgp_peering_ip"] = string(*obj.CustomerBgpPeeringIp)
	}

	if obj.CustomerBgpPeeringIpv6 != nil {
		result["customer_bgp_peering_ipv6"] = string(*obj.CustomerBgpPeeringIpv6)
	}

	if obj.OracleBgpPeeringIp != nil {
		result["oracle_bgp_peering_ip"] = string(*obj.OracleBgpPeeringIp)
	}

	if obj.OracleBgpPeeringIpv6 != nil {
		result["oracle_bgp_peering_ipv6"] = string(*obj.OracleBgpPeeringIpv6)
	}

	if obj.Vlan != nil {
		result["vlan"] = int(*obj.Vlan)
	}

	return result
}

func VirtualCircuitRedundancyMetadataToMap(obj *oci_core.VirtualCircuitRedundancyMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	result["configured_redundancy_level"] = string(obj.ConfiguredRedundancyLevel)

	result["ipv4bgp_session_redundancy_status"] = string(obj.Ipv4bgpSessionRedundancyStatus)

	result["ipv6bgp_session_redundancy_status"] = string(obj.Ipv6bgpSessionRedundancyStatus)

	return result
}

func publicPrefixesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if cidrBlock, ok := m["cidr_block"]; ok && cidrBlock != "" {
		buf.WriteString(fmt.Sprintf("%v-", cidrBlock))
	}
	return utils.GetStringHashcode(buf.String())
}
func (s *CoreVirtualCircuitResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeVirtualCircuitCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VirtualCircuitId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeVirtualCircuitCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
