// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	routeTableAttachmentRepresentation = map[string]interface{}{
		"subnet_id":      Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"route_table_id": Representation{repType: Required, create: `${oci_core_route_table.test_route_table.id}`},
	}

	RouteTableResourceAttachmentDependencies = SubnetResourceConfig
)

func TestCoreRouteTableAttachmentResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_route_table_attachment.test_route_table_attachment"

	var routeTableIdFromRT string

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
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
