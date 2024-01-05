// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CrossConnectWithGroupOptionalResource = acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Optional, acctest.Create, crossConnectWithGroupRepresentation)

	CrossConnectWithGroupResourceConfig = CrossConnectWithGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Optional, acctest.Update, crossConnectWithGroupRepresentation)

	CrossConnectWithGroupResourceConfigCopyForVC = CrossConnectWithGroupResourceDependenciesCopyForVC +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Optional, acctest.Update, crossConnectWithGroupRepresentation)

	crossConnectWithGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cross_connect_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_cross_connect_group.test_cross_connect_group.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreCrossConnectDataSourceFilterRepresentation}}

	crossConnectWithGroupRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"location_name":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_cross_connect_locations.test_cross_connect_locations.cross_connect_locations.1.name}`}, // not all locations support macsec
		"port_speed_shape_name":   acctest.Representation{RepType: acctest.Required, Create: `10 Gbps`},
		"cross_connect_group_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_cross_connect_group.test_cross_connect_group.id}`},
		"customer_reference_name": acctest.Representation{RepType: acctest.Optional, Create: `customerReferenceName`, Update: `customerReferenceName2`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_active":               acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	CrossConnectWithGroupResourceDependencies          = CoreCrossConnectGroupResourceConfig + acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_locations", "test_cross_connect_locations", acctest.Required, acctest.Create, CoreCoreCrossConnectLocationDataSourceRepresentation)
	CrossConnectWithGroupResourceDependenciesCopyForVC = CrossConnectGroupResourceConfigCopyForVC + acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_locations", "test_cross_connect_locations", acctest.Required, acctest.Create, CoreCoreCrossConnectLocationDataSourceRepresentation)
)

// issue-routing-tag: core/default
func TestResourceCoreCrossConnectResourceWithinGroup(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreCrossConnectResourceWithinGroup")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	secretIdCKN := utils.GetEnvSettingWithBlankDefault("secret_ocid_ckn")
	secretIdVariableStrCKN := fmt.Sprintf("variable \"secret_ocid_ckn\" { default = \"%s\" }\n", secretIdCKN)

	secretIdCAK := utils.GetEnvSettingWithBlankDefault("secret_ocid_cak")
	secretIdVariableStrCAK := fmt.Sprintf("variable \"secret_ocid_cak\" { default = \"%s\" }\n", secretIdCAK)

	// cross connect group dependency is with the updated values
	secretVersionCAKU := utils.GetEnvSettingWithBlankDefault("secret_version_cak_for_update")
	secretVersionUStrCAK := fmt.Sprintf("variable \"secret_version_cak_for_update\" { default = \"%s\" }\n", secretVersionCAKU)

	secretVersionCKN := utils.GetEnvSettingWithBlankDefault("secret_version_ckn")
	secretVersionStrCKN := fmt.Sprintf("variable \"secret_version_ckn\" { default = \"%s\" }\n", secretVersionCKN)

	resourceName := "oci_core_cross_connect.test_cross_connect"
	datasourceName := "data.oci_core_cross_connects.test_cross_connects"
	singularDatasourceName := "data.oci_core_cross_connect.test_cross_connect"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreCrossConnectDestroy,
		Steps: []resource.TestStep{
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + CrossConnectWithGroupResourceDependenciesCopyForVC + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionUStrCAK + secretVersionStrCKN +
					acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Optional, acctest.Create, crossConnectWithGroupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "location_name"),
					resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CrossConnectWithGroupResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionUStrCAK + secretVersionStrCKN +
					acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Optional, acctest.Update, crossConnectWithGroupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "location_name"),
					resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),

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
				Config: config + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionUStrCAK + secretVersionStrCKN +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connects", "test_cross_connects", acctest.Optional, acctest.Update, crossConnectWithGroupDataSourceRepresentation) +
					compartmentIdVariableStr + CrossConnectWithGroupResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Optional, acctest.Update, crossConnectWithGroupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

					resource.TestCheckResourceAttr(datasourceName, "cross_connects.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.cross_connect_group_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.customer_reference_name", "customerReferenceName2"),
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
				Config: config + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionUStrCAK + secretVersionStrCKN +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Required, acctest.Create, CoreCoreCrossConnectSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CrossConnectWithGroupResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "customer_reference_name", "customerReferenceName2"),
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
			// verify resource import
			{
				Config:            config + CrossConnectWithGroupOptionalResource,
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
