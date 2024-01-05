// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	govSpecificVirtualCircuitRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   acctest.Representation{RepType: acctest.Optional, Create: `10 Gbps`, Update: `20 Gbps`},
		"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: govSpecificCrossConnectMappingsRequiredOnlyRepresentation},
		"customer_bgp_asn":       acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"gateway_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg.test_drg.id}`},
		"region":                 acctest.Representation{RepType: acctest.Optional, Create: `us-phoenix-1`},
	}

	govSpecificVirtualCircuitWithProviderRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   acctest.Representation{RepType: acctest.Optional, Create: "${data.oci_core_virtual_circuit_bandwidth_shapes.test_virtual_circuit_bandwidth_shapes.virtual_circuit_bandwidth_shapes.0.name}"},
		"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: govSpecificVirtualCircuitCrossConnectMappingsRepresentation},
		"customer_bgp_asn":       acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"gateway_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg.test_drg.id}`},
		"provider_service_id":    acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
		// provider_service_key_name can only be updated by a Fast Connect Service Provider
		// "provider_service_key_name": acctest.Representation{RepType: acctest.Optional, Create: `d8f7a443-28c2-4dcf-996c-286351908c58`},
		"region": acctest.Representation{RepType: acctest.Optional, Create: `us-phoenix-1`},
	}

	govSpecificVirtualCircuitRequiredOnlyRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `${var.virtual_circuit_type}`},
		"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: govSpecificCrossConnectMappingsRequiredOnlyRepresentation},
		"customer_bgp_asn":       acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"gateway_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
	}

	govSpecificCrossConnectMappingsRequiredOnlyRepresentation = map[string]interface{}{
		"cross_connect_or_cross_connect_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cross_connect.test_cross_connect.cross_connect_group_id}`},
		"customer_bgp_peering_ip":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.18/31`, Update: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":                   acctest.Representation{RepType: acctest.Required, Create: `10.0.0.19/31`, Update: `10.0.0.21/31`},
		"customer_bgp_peering_ipv6":               acctest.Representation{RepType: acctest.Required, Create: `fd00:aaaa:0123::4/64`, Update: `fd00:aaaa:0124::4/64`},
		"oracle_bgp_peering_ipv6":                 acctest.Representation{RepType: acctest.Required, Create: `fd00:aaaa:0123::5/64`, Update: `fd00:aaaa:0124::5/64`},
		"vlan":                                    acctest.Representation{RepType: acctest.Required, Create: `200`, Update: `300`},
	}

	govSpecificVirtualCircuitCrossConnectMappingsRepresentation = map[string]interface{}{
		"customer_bgp_peering_ip":   acctest.Representation{RepType: acctest.Required, Create: `10.0.0.18/31`, Update: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":     acctest.Representation{RepType: acctest.Required, Create: `10.0.0.19/31`, Update: `10.0.0.21/31`},
		"customer_bgp_peering_ipv6": acctest.Representation{RepType: acctest.Required, Create: `fd00:aaaa:0123::4/64`, Update: `fd00:aaaa:0124::4/64`},
		"oracle_bgp_peering_ipv6":   acctest.Representation{RepType: acctest.Required, Create: `fd00:aaaa:0123::5/64`, Update: `fd00:aaaa:0124::5/64`},
	}
)

// issue-routing-tag: core/default
func TestGovSpecificCoreVirtualCircuitResource_basic(t *testing.T) {
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "IPv6") {
		t.Skip("DoDIPv6 test not supported in this realm")
	}
	httpreplay.SetScenario("TestCoreVirtualCircuitResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_virtual_circuit.test_virtual_circuit"
	datasourceName := "data.oci_core_virtual_circuits.test_virtual_circuits"
	singularDatasourceName := "data.oci_core_virtual_circuit.test_virtual_circuit"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckCoreVirtualCircuitDestroy, []resource.TestStep{
		// verify Create - PUBLIC Virtual Circuit
		{
			Config: config + compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies + VirtualCircuitPublicPropertyVariables +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Create, CoreVirtualCircuitPublicRequiredOnlyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
					"cidr_block": "11.0.0.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update - PUBLIC Virtual Circuit
		{
			Config: config + compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies + VirtualCircuitPublicPropertyVariables +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Update, CoreVirtualCircuitPublicRequiredOnlyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
					"cidr_block": "11.0.1.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies,
		},
		// verify Create - PRIVATE Virtual Circuit with Provider
		{
			PreConfig: func() {
				log.Printf("%s", config+compartmentIdVariableStr+CoreVirtualCircuitResourceDependencies+VirtualCircuitPrivatePropertyVariables+VirtualCircuitWithProviderResourceConfigFilter+
					acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Create, acctest.GetUpdatedRepresentationCopy("cross_connect_mappings", govSpecificVirtualCircuitCrossConnectMappingsRepresentation, CoreVirtualCircuitWithProviderRepresentation)))
				fmt.Printf("%s", config+compartmentIdVariableStr+CoreVirtualCircuitResourceDependencies+VirtualCircuitPrivatePropertyVariables+VirtualCircuitWithProviderResourceConfigFilter+
					acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Create, acctest.GetUpdatedRepresentationCopy("cross_connect_mappings", govSpecificVirtualCircuitCrossConnectMappingsRepresentation, CoreVirtualCircuitWithProviderRepresentation)))
			},
			Config: config + compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables + VirtualCircuitWithProviderResourceConfigFilter +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Create, govSpecificVirtualCircuitWithProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update - PRIVATE Virtual Circuit with Provider
		{
			Config: config + compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables + VirtualCircuitWithProviderResourceConfigFilter +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Update, govSpecificVirtualCircuitWithProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies,
		},

		// verify Create - PRIVATE Virtual Circuit
		{
			Config: config + VirtualCircuitPrivatePropertyVariables + compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Create, govSpecificVirtualCircuitRequiredOnlyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Create, govSpecificVirtualCircuitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Update, govSpecificVirtualCircuitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_core_virtual_circuits", "test_virtual_circuits", acctest.Optional, acctest.Update, CoreCoreVirtualCircuitDataSourceRepresentation) +
				compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Update, govSpecificVirtualCircuitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Create, CoreCoreVirtualCircuitSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreVirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Update, govSpecificVirtualCircuitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
		// verify resource import
		{
			Config:            config + CoreVirtualCircuitRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"region",
			},
			ResourceName: resourceName,
		},
	})
}
