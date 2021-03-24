// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v37/common"
	oci_core "github.com/oracle/oci-go-sdk/v37/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VirtualCircuitRequiredOnlyResource = VirtualCircuitResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitRequiredOnlyRepresentation)

	virtualCircuitSingularDataSourceRepresentation = map[string]interface{}{
		"virtual_circuit_id": Representation{repType: Required, create: `${oci_core_virtual_circuit.test_virtual_circuit.id}`},
	}

	virtualCircuitDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `PROVISIONED`},
		"filter":         RepresentationGroup{Required, virtualCircuitDataSourceFilterRepresentation}}
	virtualCircuitDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_virtual_circuit.test_virtual_circuit.id}`}},
	}

	virtualCircuitPublicRequiredOnlyRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"type":                   Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"cross_connect_mappings": RepresentationGroup{Required, crossConnectMappingsPublicRequiredOnlyRepresentation},
		"customer_bgp_asn":       Representation{repType: Required, create: `10`, update: `11`},
		"public_prefixes":        RepresentationGroup{Required, virtualCircuitPublicPrefixesRepresentation},
	}
	virtualCircuitPublicRequiredOnlyWithoutDeprecatedRepresentation = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(virtualCircuitPublicRequiredOnlyRepresentation, []string{"customer_bgp_asn"}), map[string]interface{}{
			"customer_asn": Representation{repType: Required, create: `10`, update: `11`},
		})

	virtualCircuitRequiredOnlyRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"type":                   Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"cross_connect_mappings": RepresentationGroup{Required, crossConnectMappingsRequiredOnlyRepresentation},
		"customer_asn":           Representation{repType: Required, create: `10`, update: `11`},
		"gateway_id":             Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
	}
	virtualCircuitRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"type":                   Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   Representation{repType: Optional, create: `10 Gbps`, update: `20 Gbps`},
		"cross_connect_mappings": RepresentationGroup{Required, crossConnectMappingsRequiredOnlyRepresentation},
		"customer_asn":           Representation{repType: Required, create: `10`, update: `11`},
		"defined_tags":           Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":          Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"gateway_id":             Representation{repType: Optional, create: `${oci_core_drg.test_drg.id}`},
		"region":                 Representation{repType: Optional, create: `us-phoenix-1`},
	}

	virtualCircuitWithProviderRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"type":                   Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   Representation{repType: Optional, create: "${data.oci_core_virtual_circuit_bandwidth_shapes.test_virtual_circuit_bandwidth_shapes.virtual_circuit_bandwidth_shapes.0.name}"},
		"cross_connect_mappings": RepresentationGroup{Required, virtualCircuitCrossConnectMappingsRepresentation},
		"customer_asn":           Representation{repType: Required, create: `10`, update: `11`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"gateway_id":             Representation{repType: Optional, create: `${oci_core_drg.test_drg.id}`},
		"provider_service_id":    Representation{repType: Optional, create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
		// provider_service_key_name can only be updated by a Fast Connect Service Provider
		// "provider_service_key_name": Representation{repType: Optional, create: `d8f7a443-28c2-4dcf-996c-286351908c58`},
		"region": Representation{repType: Optional, create: `us-phoenix-1`},
	}

	crossConnectMappingsPublicRequiredOnlyRepresentation = map[string]interface{}{
		"cross_connect_or_cross_connect_group_id": Representation{repType: Required, create: `${oci_core_cross_connect.test_cross_connect.cross_connect_group_id}`},
		"vlan": Representation{repType: Required, create: `200`, update: `300`},
	}
	crossConnectMappingsRequiredOnlyRepresentation = map[string]interface{}{
		"cross_connect_or_cross_connect_group_id": Representation{repType: Required, create: `${oci_core_cross_connect.test_cross_connect.cross_connect_group_id}`},
		"customer_bgp_peering_ip":                 Representation{repType: Required, create: `10.0.0.18/31`, update: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":                   Representation{repType: Required, create: `10.0.0.19/31`, update: `10.0.0.21/31`},
		"vlan":                                    Representation{repType: Required, create: `200`, update: `300`},
	}
	virtualCircuitCrossConnectMappingsRepresentation = map[string]interface{}{
		"customer_bgp_peering_ip": Representation{repType: Required, create: `10.0.0.18/31`, update: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":   Representation{repType: Required, create: `10.0.0.19/31`, update: `10.0.0.21/31`},
	}
	virtualCircuitPublicPrefixesRepresentation = map[string]interface{}{
		"cidr_block": Representation{repType: Required, create: `11.0.0.0/24`, update: `11.0.1.0/24`},
	}

	VirtualCircuitWithProviderResourceConfigFilter = `
data "oci_core_fast_connect_provider_services" "test_fast_connect_provider_services" {
	#Required
	compartment_id = "${var.compartment_id}"

	filter {
		name = "type"
		values = [ "LAYER2" ]
	}

	filter {
		name = "private_peering_bgp_management"
		values = [ "CUSTOMER_MANAGED" ]
	}

	filter {
		name = "supported_virtual_circuit_types"
		values = [ "${var.virtual_circuit_type}" ]
	}

	filter {
		name = "public_peering_bgp_management"
		values = [ "ORACLE_MANAGED" ]
	}

	filter {
		name = "provider_service_key_management"
		values = ["PROVIDER_MANAGED"]
	}
}

data "oci_core_virtual_circuit_bandwidth_shapes" "test_virtual_circuit_bandwidth_shapes" {
  #Required
  provider_service_id = "${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}"
}
`

	VirtualCircuitPrivatePropertyVariables = `
variable "virtual_circuit_type" { default = "PRIVATE" }

`

	VirtualCircuitPublicPropertyVariables = `
variable "virtual_circuit_type" { default = "PUBLIC" }

`
	VirtualCircuitResourceDependencies = DrgRequiredOnlyResource + CrossConnectWithGroupResourceConfig
)

func TestCoreVirtualCircuitResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVirtualCircuitResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_virtual_circuit.test_virtual_circuit"
	datasourceName := "data.oci_core_virtual_circuits.test_virtual_circuits"
	singularDatasourceName := "data.oci_core_virtual_circuit.test_virtual_circuit"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+VirtualCircuitResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, virtualCircuitRepresentation), "core", "virtualCircuit", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVirtualCircuitDestroy,
		Steps: []resource.TestStep{
			// verify create - PUBLIC Virtual Circuit
			{
				Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPublicPropertyVariables +
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitPublicRequiredOnlyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
			// verify update from customer_bgp_asn to customer_asn
			{
				Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPublicPropertyVariables +
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitPublicRequiredOnlyWithoutDeprecatedRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
					resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
					resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
						"cidr_block": "11.0.0.0/24",
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
			// verify update - PUBLIC Virtual Circuit
			{
				Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPublicPropertyVariables +
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Update, virtualCircuitPublicRequiredOnlyWithoutDeprecatedRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "300"),
					resource.TestCheckResourceAttr(resourceName, "customer_asn", "11"),
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
				Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables + VirtualCircuitWithProviderResourceConfigFilter +
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, virtualCircuitWithProviderRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
					resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
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
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitWithProviderRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
					resource.TestCheckResourceAttr(resourceName, "customer_asn", "11"),
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
				Config: config + VirtualCircuitPrivatePropertyVariables + compartmentIdVariableStr + VirtualCircuitRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
					resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
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
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, virtualCircuitRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
					resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-phoenix-1"),
					resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create,
						representationCopyWithNewProperties(virtualCircuitRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
					resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-phoenix-1"),
					resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "20 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "300"),
					resource.TestCheckResourceAttr(resourceName, "customer_asn", "11"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.vlan", "300"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.customer_asn", "11"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.freeform_tags.%", "1"),
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
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "gateway_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_circuit_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "bandwidth_shape_name", "20 Gbps"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_management"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_session_state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.vlan", "300"),
					resource.TestCheckResourceAttr(singularDatasourceName, "customer_asn", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_bgp_asn"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "service_type"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "oracle_bgp_asn", "31898"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_type", "COLOCATED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "PROVISIONED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "type", "PRIVATE"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentation),
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
		},
	})
}

func testAccCheckCoreVirtualCircuitDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_virtual_circuit" {
			noResourceFound = false
			request := oci_core.GetVirtualCircuitRequest{}

			tmp := rs.Primary.ID
			request.VirtualCircuitId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetVirtualCircuit(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VirtualCircuitLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreVirtualCircuit") {
		resource.AddTestSweepers("CoreVirtualCircuit", &resource.Sweeper{
			Name:         "CoreVirtualCircuit",
			Dependencies: DependencyGraph["virtualCircuit"],
			F:            sweepCoreVirtualCircuitResource,
		})
	}
}

func sweepCoreVirtualCircuitResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	virtualCircuitIds, err := getVirtualCircuitIds(compartment)
	if err != nil {
		return err
	}
	for _, virtualCircuitId := range virtualCircuitIds {
		if ok := SweeperDefaultResourceId[virtualCircuitId]; !ok {
			deleteVirtualCircuitRequest := oci_core.DeleteVirtualCircuitRequest{}

			deleteVirtualCircuitRequest.VirtualCircuitId = &virtualCircuitId

			deleteVirtualCircuitRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteVirtualCircuit(context.Background(), deleteVirtualCircuitRequest)
			if error != nil {
				fmt.Printf("Error deleting VirtualCircuit %s %s, It is possible that the resource is already deleted. Please verify manually \n", virtualCircuitId, error)
				continue
			}
			waitTillCondition(testAccProvider, &virtualCircuitId, virtualCircuitSweepWaitCondition, time.Duration(3*time.Minute),
				virtualCircuitSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVirtualCircuitIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VirtualCircuitId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listVirtualCircuitsRequest := oci_core.ListVirtualCircuitsRequest{}
	listVirtualCircuitsRequest.CompartmentId = &compartmentId
	listVirtualCircuitsRequest.LifecycleState = oci_core.VirtualCircuitLifecycleStateProvisioned
	listVirtualCircuitsResponse, err := virtualNetworkClient.ListVirtualCircuits(context.Background(), listVirtualCircuitsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VirtualCircuit list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, virtualCircuit := range listVirtualCircuitsResponse.Items {
		id := *virtualCircuit.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "VirtualCircuitId", id)
	}
	return resourceIds, nil
}

func virtualCircuitSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if virtualCircuitResponse, ok := response.Response.(oci_core.GetVirtualCircuitResponse); ok {
		return virtualCircuitResponse.LifecycleState != oci_core.VirtualCircuitLifecycleStateTerminated
	}
	return false
}

func virtualCircuitSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetVirtualCircuit(context.Background(), oci_core.GetVirtualCircuitRequest{
		VirtualCircuitId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
