// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	routeTableAttachmentRepresentation = map[string]interface{}{
		"subnet_id":      Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"route_table_id": Representation{RepType: Required, Create: `${oci_core_route_table.test_route_table.id}`},
	}

	RouteTableResourceAttachmentDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Optional, Update, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Required, Create, dhcpOptionsRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Required, Create, securityListRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label":      Representation{RepType: Required, Create: `dnslabel`},
			"is_ipv6enabled": Representation{RepType: Optional, Create: `true`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table_2", Required, Create, routeTableRepresentation)
)

// issue-routing-tag: core/virtualNetwork
func TestCoreRouteTableAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreRouteTableAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_route_table_attachment.test_route_table_attachment"

	var routeTableIdFromRT string

	// This is manually written test for manually written resource. The resource name is not in codegen's resources.yaml file.
	// We are unable to process the resource name. Thus not saving config for this test.
	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RouteTableResourceAttachmentDependencies +
				GenerateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", Required, Create, routeTableAttachmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					routeTableIdFromRT, err := FromInstanceState(s, "oci_core_route_table.test_route_table", "id")
					if err != nil {
						return err
					}

					routeTableIdFromSubnet, err := FromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
					if routeTableIdFromRT != routeTableIdFromSubnet {
						return fmt.Errorf("requested routeTable was not attached to the subnet")
					}
					return err
				},
			),
		},
		// detach and attach another / recreate
		{
			Config: config + compartmentIdVariableStr + GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Optional, Update,
				RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
					"route_table_id": Representation{RepType: Required, Create: `${oci_core_route_table.test_route_table_2.id}`},
				})) +
				GenerateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Required, Create, dhcpOptionsRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Required, Create, securityListRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
					"dns_label":      Representation{RepType: Required, Create: `dnslabel`},
					"is_ipv6enabled": Representation{RepType: Optional, Create: `true`},
				})) +
				AvailabilityDomainConfig +
				DefinedTagsDependencies +
				GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table_2", Required, Create, routeTableRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", Required, Create,
					RepresentationCopyWithNewProperties(routeTableAttachmentRepresentation, map[string]interface{}{
						"route_table_id": Representation{RepType: Required, Create: `${oci_core_route_table.test_route_table_2.id}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					routeTableIdFromRT, err := FromInstanceState(s, "oci_core_route_table.test_route_table_2", "id")
					if err != nil {
						return err
					}

					routeTableIdFromSubnet, err := FromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
					if routeTableIdFromRT != routeTableIdFromSubnet {
						return fmt.Errorf("requested routeTable was not attached to the subnet")
					}
					return err
				},
			),
		},
		// revert to original / recreate
		{
			Config: config + compartmentIdVariableStr + RouteTableResourceAttachmentDependencies +
				GenerateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", Required, Create, routeTableAttachmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					routeTableIdFromRT, err := FromInstanceState(s, "oci_core_route_table.test_route_table", "id")
					if err != nil {
						return err
					}

					routeTableIdFromSubnet, err := FromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
					if routeTableIdFromRT != routeTableIdFromSubnet {
						return fmt.Errorf("requested routeTable was not attached to the subnet")
					}
					return err
				},
			),
		},
		// verify delete
		{
			Config:             config + compartmentIdVariableStr + RouteTableResourceAttachmentDependencies,
			ExpectNonEmptyPlan: true,
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					routeTableIdFromSubnet, err := FromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
					if routeTableIdFromRT == routeTableIdFromSubnet {
						return fmt.Errorf("requested routeTable was not detached from the subnet")
					}
					return err
				},
			),
		},
		// Create attachment to import
		{
			Config: config + compartmentIdVariableStr + RouteTableResourceAttachmentDependencies +
				GenerateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", Required, Create, routeTableAttachmentRepresentation),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
