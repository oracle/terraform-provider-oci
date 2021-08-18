// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	govSpecificVirtualCircuitRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"type":                   Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   Representation{repType: Optional, create: `10 Gbps`, update: `20 Gbps`},
		"cross_connect_mappings": RepresentationGroup{Required, govSpecificCrossConnectMappingsRequiredOnlyRepresentation},
		"customer_bgp_asn":       Representation{repType: Required, create: `10`, update: `11`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"gateway_id":             Representation{repType: Optional, create: `${oci_core_drg.test_drg.id}`},
		"region":                 Representation{repType: Optional, create: `us-phoenix-1`},
	}

	govSpecificVirtualCircuitWithProviderRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"type":                   Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   Representation{repType: Optional, create: "${data.oci_core_virtual_circuit_bandwidth_shapes.test_virtual_circuit_bandwidth_shapes.virtual_circuit_bandwidth_shapes.0.name}"},
		"cross_connect_mappings": RepresentationGroup{Required, govSpecificVirtualCircuitCrossConnectMappingsRepresentation},
		"customer_bgp_asn":       Representation{repType: Required, create: `10`, update: `11`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"gateway_id":             Representation{repType: Optional, create: `${oci_core_drg.test_drg.id}`},
		"provider_service_id":    Representation{repType: Optional, create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
		// provider_service_key_name can only be updated by a Fast Connect Service Provider
		// "provider_service_key_name": Representation{repType: Optional, create: `d8f7a443-28c2-4dcf-996c-286351908c58`},
		"region": Representation{repType: Optional, create: `us-phoenix-1`},
	}

	govSpecificVirtualCircuitRequiredOnlyRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"type":                   Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"cross_connect_mappings": RepresentationGroup{Required, govSpecificCrossConnectMappingsRequiredOnlyRepresentation},
		"customer_bgp_asn":       Representation{repType: Required, create: `10`, update: `11`},
		"gateway_id":             Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
	}

	govSpecificCrossConnectMappingsRequiredOnlyRepresentation = map[string]interface{}{
		"cross_connect_or_cross_connect_group_id": Representation{repType: Required, create: `${oci_core_cross_connect.test_cross_connect.cross_connect_group_id}`},
		"customer_bgp_peering_ip":                 Representation{repType: Required, create: `10.0.0.18/31`, update: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":                   Representation{repType: Required, create: `10.0.0.19/31`, update: `10.0.0.21/31`},
		"customer_bgp_peering_ipv6":               Representation{repType: Required, create: `fd00:aaaa:0123::4/64`, update: `fd00:aaaa:0124::4/64`},
		"oracle_bgp_peering_ipv6":                 Representation{repType: Required, create: `fd00:aaaa:0123::5/64`, update: `fd00:aaaa:0124::5/64`},
		"vlan":                                    Representation{repType: Required, create: `200`, update: `300`},
	}

	govSpecificVirtualCircuitCrossConnectMappingsRepresentation = map[string]interface{}{
		"customer_bgp_peering_ip":   Representation{repType: Required, create: `10.0.0.18/31`, update: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":     Representation{repType: Required, create: `10.0.0.19/31`, update: `10.0.0.21/31`},
		"customer_bgp_peering_ipv6": Representation{repType: Required, create: `fd00:aaaa:0123::4/64`, update: `fd00:aaaa:0124::4/64`},
		"oracle_bgp_peering_ipv6":   Representation{repType: Required, create: `fd00:aaaa:0123::5/64`, update: `fd00:aaaa:0124::5/64`},
	}
)

// issue-routing-tag: core/default
func TestGovSpecificCoreVirtualCircuitResource_basic(t *testing.T) {
	if !strings.Contains(getEnvSettingWithBlankDefault("enabled_tests"), "IPv6") {
		t.Skip("DoDIPv6 test not supported in this realm")
	}
	httpreplay.SetScenario("TestCoreVirtualCircuitResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_virtual_circuit.test_virtual_circuit"
	datasourceName := "data.oci_core_virtual_circuits.test_virtual_circuits"
	singularDatasourceName := "data.oci_core_virtual_circuit.test_virtual_circuit"

	var resId, resId2 string

	ResourceTest(t, testAccCheckCoreVirtualCircuitDestroy, []resource.TestStep{
		// verify create - PUBLIC Virtual Circuit
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPublicPropertyVariables +
				generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitPublicRequiredOnlyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
					"cidr_block": "11.0.0.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify update - PUBLIC Virtual Circuit
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPublicPropertyVariables +
				generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Update, virtualCircuitPublicRequiredOnlyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
					"cidr_block": "11.0.1.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies,
		},
		// verify create - PRIVATE Virtual Circuit with Provider
		{
			PreConfig: func() {
				log.Printf("%s", config+compartmentIdVariableStr+VirtualCircuitResourceDependencies+VirtualCircuitPrivatePropertyVariables+VirtualCircuitWithProviderResourceConfigFilter+
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, getUpdatedRepresentationCopy("cross_connect_mappings", govSpecificVirtualCircuitCrossConnectMappingsRepresentation, virtualCircuitWithProviderRepresentation)))
				fmt.Printf("%s", config+compartmentIdVariableStr+VirtualCircuitResourceDependencies+VirtualCircuitPrivatePropertyVariables+VirtualCircuitWithProviderResourceConfigFilter+
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, getUpdatedRepresentationCopy("cross_connect_mappings", govSpecificVirtualCircuitCrossConnectMappingsRepresentation, virtualCircuitWithProviderRepresentation)))
			},
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables + VirtualCircuitWithProviderResourceConfigFilter +
				generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, govSpecificVirtualCircuitWithProviderRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ipv6", "fd00:aaaa:123::4/64"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ipv6", "fd00:aaaa:123::5/64"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_service_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify update - PRIVATE Virtual Circuit with Provider
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables + VirtualCircuitWithProviderResourceConfigFilter +
				generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, govSpecificVirtualCircuitWithProviderRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ipv6", "fd00:aaaa:124::4/64"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ipv6", "fd00:aaaa:124::5/64"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_service_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies,
		},

		// verify create - PRIVATE Virtual Circuit
		{
			Config: config + VirtualCircuitPrivatePropertyVariables + compartmentIdVariableStr + VirtualCircuitResourceDependencies +
				generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, govSpecificVirtualCircuitRequiredOnlyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ipv6", "fd00:aaaa:123::4/64"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ipv6", "fd00:aaaa:123::5/64"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies,
		},

		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
				generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, govSpecificVirtualCircuitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "10 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ipv6", "fd00:aaaa:123::4/64"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ipv6", "fd00:aaaa:123::5/64"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
				generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, govSpecificVirtualCircuitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "20 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ipv6", "fd00:aaaa:124::4/64"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ipv6", "fd00:aaaa:124::5/64"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + generateDataSourceFromRepresentationMap("oci_core_virtual_circuits", "test_virtual_circuits", Optional, Update, virtualCircuitDataSourceRepresentation) +
				compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
				generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, govSpecificVirtualCircuitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.bandwidth_shape_name", "20 Gbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.bgp_management"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.bgp_session_state"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ipv6", "fd00:aaaa:124::4/64"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ipv6", "fd00:aaaa:124::5/64"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.customer_bgp_asn", "11"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.gateway_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.oracle_bgp_asn"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.service_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.state", "PROVISIONED"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.type", "PRIVATE"),
			),
		},
		// verify singular datasource
		{
			Config: config + generateDataSourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
				generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, govSpecificVirtualCircuitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_circuit_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "bandwidth_shape_name", "20 Gbps"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_management"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_session_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ipv6", "fd00:aaaa:124::4/64"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ipv6", "fd00:aaaa:124::5/64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_bgp_asn", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_bgp_asn"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "oracle_bgp_asn", "46558"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_type", "COLOCATED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "PROVISIONED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "PRIVATE"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
				generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, govSpecificVirtualCircuitRepresentation),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"region",
			},
			ResourceName: resourceName,
		},
	})
}
