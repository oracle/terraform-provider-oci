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
	VcnRealmOptionalsResource = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
		"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`}}))
)

// issue-routing-tag: core/virtualNetwork
func TestGovSpecificCoreVcnResource_basic(t *testing.T) {
	//if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "IPv6") {
	//	t.Skip("DoDIPv6 test not supported in this realm")
	//}
	httpreplay.SetScenario("TestGovSpecificCoreVcnResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_vcn.test_vcn"
	datasourceName := "data.oci_core_vcns.test_vcns"
	singularDatasourceName := "data.oci_core_vcn.test_vcn"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckCoreVcnDestroy, []resource.TestStep{
		// verify Create with optionals which includes IPV6 parameters
		{
			Config: config + compartmentIdVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
					"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_ipv6enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
					"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_ipv6enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcns", "test_vcns", acctest.Optional, acctest.Update, CoreCoreVcnDataSourceRepresentation) +
				compartmentIdVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
					"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.default_dhcp_options_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.default_route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.default_security_list_id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.vcn_domain_name"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreCoreVcnSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VcnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
					"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_dhcp_options_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_route_table_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_security_list_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_domain_name"),
			),
		},
		// verify resource import
		{
			Config:            config + VcnRealmOptionalsResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_ipv6enabled",
			},
			ResourceName: resourceName,
		},
	})
}
