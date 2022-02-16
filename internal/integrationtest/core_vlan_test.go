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
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	VlanRequiredOnlyResource = VlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create, vlanRepresentation)

	VlanResourceConfig = VlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Optional, acctest.Update, vlanRepresentation)

	vlanSingularDataSourceRepresentation = map[string]interface{}{
		"vlan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_vlan.id}`},
	}

	vlanDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: vlanDataSourceFilterRepresentation}}
	vlanDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_vlan.test_vlan.id}`}},
	}

	vlanRepresentation = map[string]interface{}{
		"cidr_block":          acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/24`, Update: "10.0.0.0/16"},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":             acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"route_table_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.test_route_table.id}`},
		"vlan_tag":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesNsgRepresentation},
	}

	ignoreChangesVlanRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	VlanResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, internetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, routeTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreVlanResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVlanResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_vlan.test_vlan"
	datasourceName := "data.oci_core_vlans.test_vlans"
	singularDatasourceName := "data.oci_core_vlan.test_vlan"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+VlanResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Optional, acctest.Create, vlanRepresentation), "core", "vlan", t)

	acctest.ResourceTest(t, testAccCheckCoreVlanDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create, vlanRepresentation),
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
			Config: config + compartmentIdVariableStr + VlanResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Optional, acctest.Create, vlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "vlan_tag", "10"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(vlanRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "vlan_tag", "10"),

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
			Config: config + compartmentIdVariableStr + VlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Optional, acctest.Update, vlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "vlan_tag", "10"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vlans", "test_vlans", acctest.Optional, acctest.Update, vlanDataSourceRepresentation) +
				compartmentIdVariableStr + VlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Optional, acctest.Update, vlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "vlans.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "vlans.0.cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(datasourceName, "vlans.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "vlans.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "vlans.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "vlans.0.vcn_id"),
				resource.TestCheckResourceAttr(datasourceName, "vlans.0.vlan_tag", "10"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create, vlanSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VlanResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vlan_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vlan_tag", "10"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + VlanResourceConfig,
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

func testAccCheckCoreVlanDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_vlan" {
			noResourceFound = false
			request := oci_core.GetVlanRequest{}

			tmp := rs.Primary.ID
			request.VlanId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetVlan(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VlanLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreVlan") {
		resource.AddTestSweepers("CoreVlan", &resource.Sweeper{
			Name:         "CoreVlan",
			Dependencies: acctest.DependencyGraph["vlan"],
			F:            sweepCoreVlanResource,
		})
	}
}

func sweepCoreVlanResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	vlanIds, err := getVlanIds(compartment)
	if err != nil {
		return err
	}
	for _, vlanId := range vlanIds {
		if ok := acctest.SweeperDefaultResourceId[vlanId]; !ok {
			deleteVlanRequest := oci_core.DeleteVlanRequest{}

			deleteVlanRequest.VlanId = &vlanId

			deleteVlanRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteVlan(context.Background(), deleteVlanRequest)
			if error != nil {
				fmt.Printf("Error deleting Vlan %s %s, It is possible that the resource is already deleted. Please verify manually \n", vlanId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vlanId, vlanSweepWaitCondition, time.Duration(3*time.Minute),
				vlanSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVlanIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VlanId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listVlansRequest := oci_core.ListVlansRequest{}
	listVlansRequest.CompartmentId = &compartmentId
	listVlansRequest.LifecycleState = oci_core.VlanLifecycleStateAvailable
	listVlansResponse, err := virtualNetworkClient.ListVlans(context.Background(), listVlansRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Vlan list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vlan := range listVlansResponse.Items {
		id := *vlan.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VlanId", id)
	}
	return resourceIds, nil
}

func vlanSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vlanResponse, ok := response.Response.(oci_core.GetVlanResponse); ok {
		return vlanResponse.LifecycleState != oci_core.VlanLifecycleStateTerminated
	}
	return false
}

func vlanSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetVlan(context.Background(), oci_core.GetVlanRequest{
		VlanId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
