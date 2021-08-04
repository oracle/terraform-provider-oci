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
		"subnet_id":      Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"route_table_id": Representation{repType: Required, create: `${oci_core_route_table.test_route_table.id}`},
	}

	RouteTableResourceAttachmentDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Optional, Update, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Required, Create, dhcpOptionsRepresentation) +
		generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		generateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Required, Create, securityListRepresentation) +
		generateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label":      Representation{repType: Required, create: `dnslabel`},
			"is_ipv6enabled": Representation{repType: Optional, create: `true`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table_2", Required, Create, routeTableRepresentation)
)

// issue-routing-tag: core/virtualNetwork
func TestCoreRouteTableAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreRouteTableAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_route_table_attachment.test_route_table_attachment"

	var routeTableIdFromRT string

	// This is manually written test for manually written resource. The resource name is not in codegen's resources.yaml file.
	// We are unable to process the resource name. Thus not saving config for this test.
	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + RouteTableResourceAttachmentDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", Required, Create, routeTableAttachmentRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						routeTableIdFromRT, err := fromInstanceState(s, "oci_core_route_table.test_route_table", "id")
						if err != nil {
							return err
						}

						routeTableIdFromSubnet, err := fromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
						if routeTableIdFromRT != routeTableIdFromSubnet {
							return fmt.Errorf("requested routeTable was not attached to the subnet")
						}
						return err
					},
				),
			},
			// detach and attach another / recreate
			{
				Config: config + compartmentIdVariableStr + generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Optional, Update,
					representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
						"route_table_id": Representation{repType: Required, create: `${oci_core_route_table.test_route_table_2.id}`},
					})) +
					generateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", Required, Create, dhcpOptionsRepresentation) +
					generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
					generateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Required, Create, securityListRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
						"dns_label":      Representation{repType: Required, create: `dnslabel`},
						"is_ipv6enabled": Representation{repType: Optional, create: `true`},
					})) +
					AvailabilityDomainConfig +
					DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table_2", Required, Create, routeTableRepresentation) +
					generateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", Required, Create,
						representationCopyWithNewProperties(routeTableAttachmentRepresentation, map[string]interface{}{
							"route_table_id": Representation{repType: Required, create: `${oci_core_route_table.test_route_table_2.id}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						routeTableIdFromRT, err := fromInstanceState(s, "oci_core_route_table.test_route_table_2", "id")
						if err != nil {
							return err
						}

						routeTableIdFromSubnet, err := fromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
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
					generateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", Required, Create, routeTableAttachmentRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						routeTableIdFromRT, err := fromInstanceState(s, "oci_core_route_table.test_route_table", "id")
						if err != nil {
							return err
						}

						routeTableIdFromSubnet, err := fromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
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
						routeTableIdFromSubnet, err := fromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
						if routeTableIdFromRT == routeTableIdFromSubnet {
							return fmt.Errorf("requested routeTable was not detached from the subnet")
						}
						return err
					},
				),
			},
			// create attachment to import
			{
				Config: config + compartmentIdVariableStr + RouteTableResourceAttachmentDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", Required, Create, routeTableAttachmentRepresentation),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
