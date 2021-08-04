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
	CrossConnectWithGroupResourceConfig = CrossConnectWithGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Update, crossConnectWithGroupRepresentation)

	crossConnectWithGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"cross_connect_group_id": Representation{repType: Optional, create: `${oci_core_cross_connect_group.test_cross_connect_group.id}`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"filter":                 RepresentationGroup{Required, crossConnectDataSourceFilterRepresentation}}

	crossConnectWithGroupRepresentation = map[string]interface{}{
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"location_name":           Representation{repType: Required, create: `${data.oci_core_cross_connect_locations.test_cross_connect_locations.cross_connect_locations.0.name}`},
		"port_speed_shape_name":   Representation{repType: Required, create: `10 Gbps`},
		"cross_connect_group_id":  Representation{repType: Optional, create: `${oci_core_cross_connect_group.test_cross_connect_group.id}`},
		"customer_reference_name": Representation{repType: Optional, create: `customerReferenceName`, update: `customerReferenceName2`},
		"defined_tags":            Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":           Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_active":               Representation{repType: Optional, create: `true`},
	}

	CrossConnectWithGroupResourceDependencies = CrossConnectGroupResourceConfig + generateDataSourceFromRepresentationMap("oci_core_cross_connect_locations", "test_cross_connect_locations", Required, Create, crossConnectLocationDataSourceRepresentation)
)

// issue-routing-tag: core/default
func TestResourceCoreCrossConnectResourceWithinGroup(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreCrossConnectResourceWithinGroup")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_cross_connect.test_cross_connect"
	datasourceName := "data.oci_core_cross_connects.test_cross_connects"
	singularDatasourceName := "data.oci_core_cross_connect.test_cross_connect"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreCrossConnectDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CrossConnectWithGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Create, crossConnectWithGroupRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "location_name"),
					resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CrossConnectWithGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Update, crossConnectWithGroupRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName2"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "location_name"),
					resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),

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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_cross_connects", "test_cross_connects", Optional, Update, crossConnectWithGroupDataSourceRepresentation) +
					compartmentIdVariableStr + CrossConnectWithGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Update, crossConnectWithGroupRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

					resource.TestCheckResourceAttr(datasourceName, "cross_connects.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.cross_connect_group_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.customer_reference_name", "customerReferenceName2"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.location_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.port_name"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.state", "PROVISIONED"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Required, Create, crossConnectSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CrossConnectWithGroupResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "customer_reference_name", "customerReferenceName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "location_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "port_name"),
					resource.TestCheckResourceAttr(singularDatasourceName, "port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "PROVISIONED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + CrossConnectWithGroupResourceConfig,
			},
			// verify resource import
			// import requires full configuration to handle cross connect dependency on cross connect group during destroy
			{
				Config:            config + compartmentIdVariableStr + CrossConnectWithGroupResourceConfig,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"cross_connect_id",
					"is_active",
					"far_cross_connect_or_cross_connect_group_id",
					"near_cross_connect_or_cross_connect_group_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}
