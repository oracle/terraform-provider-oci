// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RouteTableResourceAttachmentRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", acctest.Required, acctest.Create, routeTableAttachmentRepresentation)

	routeTableAttachmentRepresentation = map[string]interface{}{
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`},
	}

	RouteTableResourceAttachmentDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", acctest.Required, acctest.Create, CoreDhcpOptionsRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, CoreInternetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, CoreSecurityListRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, CoreCoreServiceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
			"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table_2", acctest.Required, acctest.Create, CoreRouteTableRepresentation)
)

// issue-routing-tag: core/virtualNetwork
func TestCoreRouteTableAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreRouteTableAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_route_table_attachment.test_route_table_attachment"

	var routeTableIdFromRT string

	// This is manually written test for manually written resource. The resource name is not in codegen's resources.yaml file.
	// We are unable to process the resource name. Thus not saving config for this test.
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RouteTableResourceAttachmentDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", acctest.Required, acctest.Create, routeTableAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					routeTableIdFromRT, err := acctest.FromInstanceState(s, "oci_core_route_table.test_route_table", "id")
					if err != nil {
						return err
					}

					routeTableIdFromSubnet, err := acctest.FromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
					if routeTableIdFromRT != routeTableIdFromSubnet {
						return fmt.Errorf("requested routeTable was not attached to the subnet")
					}
					return err
				},
			),
		},
		// detach and attach another / recreate
		{
			Config: config + compartmentIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update,
				acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
					"route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table_2.id}`},
				})) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", acctest.Required, acctest.Create, CoreDhcpOptionsRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, CoreInternetGatewayRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, CoreSecurityListRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, CoreCoreServiceDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
					"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
					"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
				})) +
				AvailabilityDomainConfig +
				DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table_2", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(routeTableAttachmentRepresentation, map[string]interface{}{
						"route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table_2.id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					routeTableIdFromRT, err := acctest.FromInstanceState(s, "oci_core_route_table.test_route_table_2", "id")
					if err != nil {
						return err
					}

					routeTableIdFromSubnet, err := acctest.FromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", acctest.Required, acctest.Create, routeTableAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					routeTableIdFromRT, err := acctest.FromInstanceState(s, "oci_core_route_table.test_route_table", "id")
					if err != nil {
						return err
					}

					routeTableIdFromSubnet, err := acctest.FromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
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
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					routeTableIdFromSubnet, err := acctest.FromInstanceState(s, "oci_core_subnet.test_subnet", "route_table_id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table_attachment", "test_route_table_attachment", acctest.Required, acctest.Create, routeTableAttachmentRepresentation),
		},
		// verify resource import
		{
			Config:                  config + RouteTableResourceAttachmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
