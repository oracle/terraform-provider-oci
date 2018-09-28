// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

var (
	VirtualCircuitPublicRequiredOnlyResource = VirtualCircuitResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitPublicRequiredOnlyRepresentation)

	VirtualCircuitRequiredOnlyResource = VirtualCircuitResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitRequiredOnlyRepresentation)

	VirtualCircuitResourceConfig = VirtualCircuitResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, virtualCircuitRepresentation)

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
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"type":           Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"cross_connect_mappings": RepresentationGroup{Required, crossConnectMappingsPublicRequiredOnlyRepresentation},
		"customer_bgp_asn":       Representation{repType: Required, create: `10`, update: `11`},
		"public_prefixes":        RepresentationGroup{Required, virtualCircuitPublicPrefixesRepresentation},
	}
	virtualCircuitRequiredOnlyRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"type":           Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"cross_connect_mappings": RepresentationGroup{Required, crossConnectMappingsRequiredOnlyRepresentation},
		"customer_bgp_asn":       Representation{repType: Required, create: `10`, update: `11`},
		"gateway_id":             Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
	}
	virtualCircuitRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"type":                   Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   Representation{repType: Optional, create: `10 Gbps`, update: `20 Gbps`},
		"cross_connect_mappings": RepresentationGroup{Required, virtualCircuitCrossConnectMappingsRepresentation},
		"customer_bgp_asn":       Representation{repType: Required, create: `10`, update: `11`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"gateway_id":             Representation{repType: Optional, create: `${oci_core_drg.test_drg.id}`},
		"provider_service_id":    Representation{repType: Optional, create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
		"region":                 Representation{repType: Optional, create: `us-phoenix-1`},
	}
	virtualCircuitRepresentationWithoutProvider = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"type":                   Representation{repType: Required, create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   Representation{repType: Optional, create: `10 Gbps`, update: `20 Gbps`},
		"cross_connect_mappings": RepresentationGroup{Required, virtualCircuitCrossConnectMappingsRepresentation},
		"customer_bgp_asn":       Representation{repType: Required, create: `10`, update: `11`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"gateway_id":             Representation{repType: Optional, create: `${oci_core_drg.test_drg.id}`},
		"region":                 Representation{repType: Optional, create: `us-phoenix-1`},
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
		"cidr_block": Representation{repType: Required, create: `0.0.0.0/5`},
	}

	VirtualCircuitResourceConfigFilter = `
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
}`

	VirtualCircuitPrivatePropertyVariables = `
variable "virtual_circuit_type" { default = "PRIVATE" }

`

	VirtualCircuitPublicPropertyVariables = `
variable "virtual_circuit_type" { default = "PUBLIC" }

`
	VirtualCircuitResourceDependencies = DrgRequiredOnlyResource + CrossConnectResourceConfig
)

func TestCoreVirtualCircuitResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_virtual_circuit.test_virtual_circuit"
	datasourceName := "data.oci_core_virtual_circuits.test_virtual_circuits"
	singularDatasourceName := "data.oci_core_virtual_circuit.test_virtual_circuit"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVirtualCircuitDestroy,
		Steps: []resource.TestStep{
			// verify create - PUBLIC Virtual Circuit
			{
				Config: config + VirtualCircuitPublicPropertyVariables + compartmentIdVariableStr + VirtualCircuitResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitPublicRequiredOnlyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
					resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "10"),
					resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
						"cidr_block": "0.0.0.0/5",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

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
			// verify create - PRIVATE Virtual Circuit with Provider
			{
				Config: config + VirtualCircuitPrivatePropertyVariables + compartmentIdVariableStr + VirtualCircuitResourceConfigFilter + VirtualCircuitResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
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
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, virtualCircuitRepresentationWithoutProvider),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
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
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentationWithoutProvider),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "20 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
					resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
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
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentationWithoutProvider),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.bandwidth_shape_name", "20 Gbps"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.vlan", "300"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.customer_bgp_asn", "11"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.gateway_id"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.state", "PROVISIONED"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.type", "PRIVATE"),
				),
			},
			// verify singular datasource
			{
				Config: config + generateDataSourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentationWithoutProvider),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "gateway_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_circuit_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "bandwidth_shape_name", "20 Gbps"),
					resource.TestCheckResourceAttr(singularDatasourceName, "bgp_management", "CUSTOMER_MANAGED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "bgp_session_state", "DOWN"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.vlan", "300"),
					resource.TestCheckResourceAttr(singularDatasourceName, "customer_bgp_asn", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
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
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentationWithoutProvider),
			},
			// verify resource import
			{
				Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables +
					generateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentationWithoutProvider),
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"public_prefixes",
					"region",
					"virtual_circuit_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckCoreVirtualCircuitDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_virtual_circuit" {
			noResourceFound = false
			request := oci_core.GetVirtualCircuitRequest{}

			tmp := rs.Primary.ID
			request.VirtualCircuitId = &tmp

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
