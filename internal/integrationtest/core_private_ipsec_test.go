// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (

	// Private CPE
	CorePrivateCpeRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ip_address":          acctest.Representation{RepType: acctest.Required, Create: `10.1.2.3`},
		"cpe_device_shape_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_cpe_device_shapes.test_cpe_device_shapes_ipsec_over_fc.cpe_device_shapes.0.cpe_device_shape_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `cpe-for-ipsec-over-fc-tf-provider-test`},
		"is_private":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	// Drg
	CorePrivateDrgRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `drg-for-ipsec-over-fc-tf-provider-test`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}

	// Drg route table
	CorePrivateDrgRouteTableRepresentation = map[string]interface{}{
		"drg_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg_private.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `drg-route-table-for-ipsec-over-fc-tf-test`},
	}

	// Drg route table
	CorePrivateDrgRouteTableTwoRepresentation = map[string]interface{}{
		"drg_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg_private.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `drg-route-table-for-ipsec-over-fc-tf-test2`},
	}

	// CC
	CorePrivateCrossConnectRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"location_name":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_cross_connect_locations.test_cross_connect_locations.cross_connect_locations.0.name}`},
		"port_speed_shape_name": acctest.Representation{RepType: acctest.Required, Create: `10 Gbps`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `cc-for-vc-for-ipsec-over-fc-tf-provider-test`},
		"is_active":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	// VC
	CoreVCforIPSecRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"bandwidth_shape_name":   acctest.Representation{RepType: acctest.Optional, Create: `10 Gbps`},
		"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: CorePrivateVirtualCircuitCrossConnectMappingsRepresentation},
		"customer_asn":           acctest.Representation{RepType: acctest.Required, Create: `12345`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `vc-for-ipsec-over-fc-tf-provider-test`},
		"gateway_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg.test_drg_private.id}`},
	}

	CorePrivateVirtualCircuitCrossConnectMappingsRepresentation = map[string]interface{}{
		"cross_connect_or_cross_connect_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cross_connect.test_cross_connect_private.id}`},
		"customer_bgp_peering_ip":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.18/31`},
		"oracle_bgp_peering_ip":                   acctest.Representation{RepType: acctest.Required, Create: `10.0.0.19/31`},
		"vlan":                                    acctest.Representation{RepType: acctest.Required, Create: `200`},
	}
	CorePrivateVirtualCircuit2CrossConnectMappingsRepresentation = map[string]interface{}{
		"cross_connect_or_cross_connect_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cross_connect.test_cross_connect_private.id}`},
		"customer_bgp_peering_ip":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":                   acctest.Representation{RepType: acctest.Required, Create: `10.0.0.21/31`},
		"vlan":                                    acctest.Representation{RepType: acctest.Required, Create: `300`},
	}
	CoreVC2forIPSecRepresentation = acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit_private_2", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(
			acctest.RepresentationCopyWithRemovedProperties(CoreVCforIPSecRepresentation, []string{"cross_connect_mappings"}), map[string]interface{}{
				"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: CorePrivateVirtualCircuit2CrossConnectMappingsRepresentation},
			}))

	// dependencies
	CorePrivateIpSecConnectionResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shapes", "test_cpe_device_shapes_ipsec_over_fc", acctest.Required, acctest.Create, CoreCoreCpeDeviceShapeDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_locations", "test_cross_connect_locations", acctest.Required, acctest.Create, CoreCoreCrossConnectLocationDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe_private", acctest.Optional, acctest.Create, CorePrivateCpeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg_private", acctest.Optional, acctest.Create, CorePrivateDrgRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table_private", acctest.Optional, acctest.Create, CorePrivateDrgRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect_private", acctest.Optional, acctest.Create, CorePrivateCrossConnectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit_private", acctest.Optional, acctest.Create, CoreVCforIPSecRepresentation) +
		DefinedTagsDependencies

	// Private IPSec
	CorePrivateIpSecConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpe_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_cpe.test_cpe_private.id}`},
		"drg_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg.test_drg_private.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CorePrivateConnectionDataSourceFilterRepresentation}}
	CorePrivateConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_ipsec.test_ip_sec_connection_private.id}`}},
	}

	CorePrivateIpSecConnectionRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpe_id":                    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cpe.test_cpe_private.id}`},
		"cpe_local_identifier":      acctest.Representation{RepType: acctest.Optional, Create: `10.1.2.3`, Update: `fakehostname`},
		"cpe_local_identifier_type": acctest.Representation{RepType: acctest.Optional, Create: `IP_ADDRESS`, Update: `HOSTNAME`},
		"drg_id":                    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg_private.id}`},
		"static_routes":             acctest.Representation{RepType: acctest.Required, Create: []string{`10.1.16.0/28`}, Update: []string{`10.1.17.0/28`}},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `MyIPSecConnectionOverFC`, Update: `displayName2`},
		"tunnel_configuration":      CorePrivateIpSecTunnelConfig,
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTags},
	}

	CorePrivateIpSecTunnelConfig = []acctest.RepresentationGroup{
		{
			RepType: acctest.Optional,
			Group: map[string]interface{}{
				"oracle_tunnel_ip":            acctest.Representation{RepType: acctest.Required, Create: `10.1.5.5`},
				"associated_virtual_circuits": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_virtual_circuit.test_virtual_circuit_private.id}`}},
				"drg_route_table_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table_private.id}`},
			},
		},
		{
			RepType: acctest.Optional,
			Group: map[string]interface{}{
				"oracle_tunnel_ip":            acctest.Representation{RepType: acctest.Required, Create: `10.1.7.7`},
				"associated_virtual_circuits": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_virtual_circuit.test_virtual_circuit_private.id}`}},
				"drg_route_table_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table_private.id}`},
			},
		},
	}
)

// issue-routing-tag: core/default
func TestCorePrivateIpSecConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePrivateIpSecConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	ipsecResourceName := "oci_core_ipsec.test_ip_sec_connection_private"
	datasourceName := "data.oci_core_ipsec_connections.test_private_ip_sec_connections"

	var resId string
	var resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CorePrivateIpSecConnectionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection_private", acctest.Optional, acctest.Create, CorePrivateIpSecConnectionRepresentation), "core", "privateIpSecConnection", t)

	acctest.ResourceTest(t, testAccCheckCoreIpSecConnectionDestroy, []resource.TestStep{
		// verify Create private IPSec (IPSec-over-FastConnect)
		{
			Config: config + compartmentIdVariableStr + CorePrivateIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection_private", acctest.Optional, acctest.Create, CorePrivateIpSecConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(ipsecResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "cpe_id"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "drg_id"),
				resource.TestCheckResourceAttr(ipsecResourceName, "static_routes.0", "10.1.16.0/28"),
				resource.TestCheckResourceAttr(ipsecResourceName, "tunnel_configuration.#", "2"),
				resource.TestCheckResourceAttr(ipsecResourceName, "tunnel_configuration.0.oracle_tunnel_ip", "10.1.5.5"),
				resource.TestCheckResourceAttr(ipsecResourceName, "tunnel_configuration.1.oracle_tunnel_ip", "10.1.7.7"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "tunnel_configuration.0.associated_virtual_circuits.#"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "tunnel_configuration.1.associated_virtual_circuits.#"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "tunnel_configuration.0.drg_route_table_id"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "tunnel_configuration.1.drg_route_table_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, ipsecResourceName, "id")
					log.Printf("created ipsec %s", resId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, ipsecResourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CorePrivateIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection_private", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CorePrivateIpSecConnectionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(ipsecResourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "cpe_id"),
				resource.TestCheckResourceAttr(ipsecResourceName, "cpe_local_identifier", "10.1.2.3"),
				resource.TestCheckResourceAttr(ipsecResourceName, "cpe_local_identifier_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(ipsecResourceName, "display_name", "MyIPSecConnectionOverFC"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "drg_id"),
				resource.TestCheckResourceAttr(ipsecResourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "id"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "state"),
				resource.TestCheckResourceAttr(ipsecResourceName, "static_routes.#", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, ipsecResourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify Update to the oracle tunnel ip returns error
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CorePrivateIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection_private", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CorePrivateIpSecConnectionRepresentation, map[string]interface{}{
						"tunnel_configuration": []acctest.RepresentationGroup{
							{
								RepType: acctest.Optional,
								Group: map[string]interface{}{
									"oracle_tunnel_ip":            acctest.Representation{RepType: acctest.Required, Create: `10.1.5.6`},
									"associated_virtual_circuits": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_virtual_circuit.test_virtual_circuit_private.id}`}},
									"drg_route_table_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table_private.id}`},
								},
							},
							{
								RepType: acctest.Optional,
								Group: map[string]interface{}{
									"oracle_tunnel_ip":            acctest.Representation{RepType: acctest.Required, Create: `10.1.7.8`},
									"associated_virtual_circuits": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_virtual_circuit.test_virtual_circuit_private.id}`}},
									"drg_route_table_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table_private.id}`},
								},
							},
						},
					})),
			ExpectError: regexp.MustCompile("oracle_tunnel_ip field cannot be updated after create ipsec connection"),
		},

		// verify Update to the associated VC returns error
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CorePrivateIpSecConnectionResourceDependencies + CoreVC2forIPSecRepresentation +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection_private", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CorePrivateIpSecConnectionRepresentation, map[string]interface{}{
						"tunnel_configuration": []acctest.RepresentationGroup{
							{
								RepType: acctest.Optional,
								Group: map[string]interface{}{
									"oracle_tunnel_ip":            acctest.Representation{RepType: acctest.Required, Create: `10.1.5.5`},
									"associated_virtual_circuits": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_virtual_circuit.test_virtual_circuit_private_2.id}`}},
									"drg_route_table_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table_private.id}`},
								},
							},
							{
								RepType: acctest.Optional,
								Group: map[string]interface{}{
									"oracle_tunnel_ip":            acctest.Representation{RepType: acctest.Required, Create: `10.1.7.7`},
									"associated_virtual_circuits": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_virtual_circuit.test_virtual_circuit_private_2.id}`}},
									"drg_route_table_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table_private.id}`},
								},
							},
						},
					})),
			ExpectError: regexp.MustCompile("associated_virtual_circuits field cannot be updated after create ipsec connection"),
		},
		// verify Update to the DRG Route Table Id returns error
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CorePrivateIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table_private_two", acctest.Optional, acctest.Create, CorePrivateDrgRouteTableTwoRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection_private", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CorePrivateIpSecConnectionRepresentation, map[string]interface{}{
						"tunnel_configuration": []acctest.RepresentationGroup{
							{
								RepType: acctest.Optional,
								Group: map[string]interface{}{
									"oracle_tunnel_ip":            acctest.Representation{RepType: acctest.Required, Create: `10.1.5.5`},
									"associated_virtual_circuits": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_virtual_circuit.test_virtual_circuit_private.id}`}},
									"drg_route_table_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table_private_two.id}`},
								},
							},
							{
								RepType: acctest.Optional,
								Group: map[string]interface{}{
									"oracle_tunnel_ip":            acctest.Representation{RepType: acctest.Required, Create: `10.1.7.7`},
									"associated_virtual_circuits": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_virtual_circuit.test_virtual_circuit_private.id}`}},
									"drg_route_table_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg_route_table.test_drg_route_table_private_two.id}`},
								},
							},
						},
					})),
			ExpectError: regexp.MustCompile("drg_route_table_id cannot be updated through oci_core_ipsec, use oci_core_drg_attachment_management resource instead to update"),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CorePrivateIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection_private", acctest.Optional, acctest.Update, CorePrivateIpSecConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(ipsecResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "cpe_id"),
				resource.TestCheckResourceAttr(ipsecResourceName, "cpe_local_identifier", "fakehostname"),
				resource.TestCheckResourceAttr(ipsecResourceName, "cpe_local_identifier_type", "HOSTNAME"),
				resource.TestCheckResourceAttr(ipsecResourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "drg_id"),
				resource.TestCheckResourceAttr(ipsecResourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "id"),
				resource.TestCheckResourceAttrSet(ipsecResourceName, "state"),
				resource.TestCheckResourceAttr(ipsecResourceName, "static_routes.0", "10.1.17.0/28"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, ipsecResourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connections", "test_private_ip_sec_connections", acctest.Optional, acctest.Update, CorePrivateIpSecConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + CorePrivateIpSecConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection_private", acctest.Optional, acctest.Update, CorePrivateIpSecConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "cpe_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),

				resource.TestCheckResourceAttr(datasourceName, "connections.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.cpe_id"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.cpe_local_identifier", "fakehostname"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.cpe_local_identifier_type", "HOSTNAME"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.drg_id"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.static_routes.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.transport_type"),
				resource.TestCheckResourceAttr(datasourceName, "connections.0.tunnel_configuration.#", "2"),
				resource.TestCheckResourceAttrPair(datasourceName, "connections.0.tunnel_configuration.0.associated_virtual_circuits.0", "oci_core_virtual_circuit.test_virtual_circuit_private", "id"),
				resource.TestCheckResourceAttrPair(datasourceName, "connections.0.tunnel_configuration.1.associated_virtual_circuits.0", "oci_core_virtual_circuit.test_virtual_circuit_private", "id"),
				resource.TestCheckResourceAttrPair(datasourceName, "connections.0.tunnel_configuration.0.drg_route_table_id", "oci_core_drg_route_table.test_drg_route_table_private", "id"),
				resource.TestCheckResourceAttrPair(datasourceName, "connections.0.tunnel_configuration.1.drg_route_table_id", "oci_core_drg_route_table.test_drg_route_table_private", "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.tunnel_configuration.0.oracle_tunnel_ip"),
				resource.TestCheckResourceAttrSet(datasourceName, "connections.0.tunnel_configuration.1.oracle_tunnel_ip")),
		},
		// verify resource import
		{
			Config:                  config + acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection_private", acctest.Optional, acctest.Create, CorePrivateIpSecConnectionRepresentation),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"tunnel_configuration"},
			ResourceName:            ipsecResourceName,
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CorePrivateIpSecConnection") {
		resource.AddTestSweepers("CorePrivateIpSecConnection", &resource.Sweeper{
			Name:         "CorePrivateIpSecConnection",
			Dependencies: acctest.DependencyGraph["privateIpSecConnection"],
			F:            sweepCorePrivateIpSecConnectionResource,
		})
	}
}

func sweepCorePrivateIpSecConnectionResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	ipSecConnectionIds, err := getCoreIpSecConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, ipSecConnectionId := range ipSecConnectionIds {
		if ok := acctest.SweeperDefaultResourceId[ipSecConnectionId]; !ok {
			deleteIPSecConnectionRequest := oci_core.DeleteIPSecConnectionRequest{}

			deleteIPSecConnectionRequest.IpscId = &ipSecConnectionId

			deleteIPSecConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteIPSecConnection(context.Background(), deleteIPSecConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting IpSecConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", ipSecConnectionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &ipSecConnectionId, CoreIpSecConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				CoreIpSecConnectionSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}
