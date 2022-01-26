// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	SubnetRequiredOnlyResource = SubnetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation)

	SubnetResourceConfig = SubnetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, subnetRepresentation)

	subnetSingularDataSourceRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	subnetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MySubnet`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: subnetDataSourceFilterRepresentation}}
	subnetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.test_subnet.id}`}},
	}

	subnetRepresentation = map[string]interface{}{
		"cidr_block":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/24`, Update: "10.0.0.0/16"},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain":        acctest.Representation{RepType: acctest.Optional, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dhcp_options_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`, Update: `${oci_core_dhcp_options.test_dhcp_options.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `MySubnet`, Update: `displayName2`},
		"dns_label":                  acctest.Representation{RepType: acctest.Optional, Create: `dnslabel`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"prohibit_internet_ingress":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"route_table_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`, Update: `${oci_core_route_table.test_route_table.id}`},
		"security_list_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}, Update: []string{`${oci_core_security_list.test_security_list.id}`}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	SubnetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", acctest.Required, acctest.Create, dhcpOptionsRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, internetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, routeTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, securityListRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, serviceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
			"is_ipv6enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
	AnotherSecurityListRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, securityListRepresentation)
	SubnetRequiredOnlyResourceDependencies  = AvailabilityDomainConfig + VcnResourceConfig
)

// issue-routing-tag: core/virtualNetwork
func TestCoreSubnetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreSubnetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_subnet.test_subnet"
	datasourceName := "data.oci_core_subnets.test_subnets"
	singularDatasourceName := "data.oci_core_subnet.test_subnet"

	// Get subnet CIDR block based on its VCN CIDR Block
	// For example: VCN CIDR Block: 2607:9b80:9a0f:0100::/56, Subnet CIDR Block: 2607:9b80:9a0f:0100::/64
	subnetCidrBlock := `${substr(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn.ipv6cidr_blocks[0]) - 2)}${64}`
	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SubnetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, subnetRepresentation), "core", "subnet", t)

	acctest.ResourceTest(t, testAccCheckCoreSubnetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SubnetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
					"ipv6cidr_block": acctest.Representation{RepType: acctest.Optional, Create: subnetCidrBlock},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MySubnet"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MySubnet"),
				resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + SubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
					"ipv6cidr_block": acctest.Representation{RepType: acctest.Optional, Update: subnetCidrBlock},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_subnets", "test_subnets", acctest.Optional, acctest.Update, subnetDataSourceRepresentation) +
				compartmentIdVariableStr + SubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, subnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "subnets.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.dhcp_options_id"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.ipv6cidr_block"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.subnet_domain_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.vcn_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.virtual_router_ip"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.virtual_router_mac"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SubnetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(singularDatasourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_domain_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_router_mac"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + SubnetResourceConfig,
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

func testAccCheckCoreSubnetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_subnet" {
			noResourceFound = false
			request := oci_core.GetSubnetRequest{}

			tmp := rs.Primary.ID
			request.SubnetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetSubnet(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.SubnetLifecycleStateTerminated): true,
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreSubnet") {
		resource.AddTestSweepers("CoreSubnet", &resource.Sweeper{
			Name:         "CoreSubnet",
			Dependencies: acctest.DependencyGraph["subnet"],
			F:            sweepCoreSubnetResource,
		})
	}
}

func sweepCoreSubnetResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	subnetIds, err := getSubnetIds(compartment)
	if err != nil {
		return err
	}
	for _, subnetId := range subnetIds {
		if ok := acctest.SweeperDefaultResourceId[subnetId]; !ok {
			deleteSubnetRequest := oci_core.DeleteSubnetRequest{}

			deleteSubnetRequest.SubnetId = &subnetId

			deleteSubnetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteSubnet(context.Background(), deleteSubnetRequest)
			if error != nil {
				fmt.Printf("Error deleting Subnet %s %s, It is possible that the resource is already deleted. Please verify manually \n", subnetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &subnetId, subnetSweepWaitCondition, time.Duration(3*time.Minute),
				subnetSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getSubnetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubnetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listSubnetsRequest := oci_core.ListSubnetsRequest{}
	listSubnetsRequest.CompartmentId = &compartmentId
	listSubnetsRequest.LifecycleState = oci_core.SubnetLifecycleStateAvailable
	listSubnetsResponse, err := virtualNetworkClient.ListSubnets(context.Background(), listSubnetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Subnet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subnet := range listSubnetsResponse.Items {
		id := *subnet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubnetId", id)
	}
	return resourceIds, nil
}

func subnetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if subnetResponse, ok := response.Response.(oci_core.GetSubnetResponse); ok {
		return subnetResponse.LifecycleState != oci_core.SubnetLifecycleStateTerminated
	}
	return false
}

func subnetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetSubnet(context.Background(), oci_core.GetSubnetRequest{
		SubnetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
