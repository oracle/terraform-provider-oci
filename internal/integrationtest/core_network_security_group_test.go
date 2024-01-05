// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreNetworkSecurityGroupRequiredOnlyResource = CoreNetworkSecurityGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation)

	CoreNetworkSecurityGroupResourceConfig = CoreNetworkSecurityGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Optional, acctest.Update, CoreNetworkSecurityGroupRepresentation)

	CoreCoreNetworkSecurityGroupSingularDataSourceRepresentation = map[string]interface{}{
		"network_security_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
	}

	CoreCoreNetworkSecurityGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreNetworkSecurityGroupDataSourceFilterRepresentation},
	}

	CoreNetworkSecurityNetworkSecurityGroupVlanDataSourceRepresentation = map[string]interface{}{
		"vlan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_vlan.id}`},
	}

	CoreNetworkSecurityGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
	}

	CoreNetworkSecurityGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreNetworkSecurityIgnoreChangesNsgRepresentation},
	}

	CoreNetworkSecurityIgnoreChangesNsgRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	CoreNetworkSecurityVlanNsgRepresentation = acctest.RepresentationCopyWithRemovedProperties(CoreVlanRepresentation, []string{"route_table_id"})

	CoreNetworkSecurityGroupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Optional, acctest.Create, CoreNetworkSecurityVlanNsgRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreNetworkSecurityGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreNetworkSecurityGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_network_security_group.test_network_security_group"
	datasourceName := "data.oci_core_network_security_groups.test_network_security_groups"
	singularDatasourceName := "data.oci_core_network_security_group.test_network_security_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreNetworkSecurityGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Optional, acctest.Create, CoreNetworkSecurityGroupRepresentation), "core", "networkSecurityGroup", t)

	acctest.ResourceTest(t, testAccCheckCoreNetworkSecurityGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreNetworkSecurityGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreNetworkSecurityGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Optional, acctest.Create, CoreNetworkSecurityGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreNetworkSecurityGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreNetworkSecurityGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + CoreNetworkSecurityGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Optional, acctest.Update, CoreNetworkSecurityGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_network_security_groups", "test_network_security_groups", acctest.Optional, acctest.Update, CoreCoreNetworkSecurityGroupDataSourceRepresentation) +
				compartmentIdVariableStr + CoreNetworkSecurityGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Optional, acctest.Update, CoreNetworkSecurityGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "network_security_groups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "network_security_groups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "network_security_groups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "network_security_groups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_groups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_groups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_groups.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_groups.0.vcn_id"),
			),
		},

		// verify with vlan query parameter only
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_network_security_groups", "test_network_security_groups", acctest.Optional, acctest.Update, CoreNetworkSecurityNetworkSecurityGroupVlanDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Optional, acctest.Update, CoreNetworkSecurityGroupRepresentation) +
				compartmentIdVariableStr + CoreNetworkSecurityGroupResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "vlan_id"),

				resource.TestCheckResourceAttr(datasourceName, "network_security_groups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "network_security_groups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "network_security_groups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "network_security_groups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_groups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_groups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_groups.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_groups.0.vcn_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreCoreNetworkSecurityGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreNetworkSecurityGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_security_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreNetworkSecurityGroupRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreNetworkSecurityGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_network_security_group" {
			noResourceFound = false
			request := oci_core.GetNetworkSecurityGroupRequest{}

			tmp := rs.Primary.ID
			request.NetworkSecurityGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetNetworkSecurityGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.NetworkSecurityGroupLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreNetworkSecurityGroup") {
		resource.AddTestSweepers("CoreNetworkSecurityGroup", &resource.Sweeper{
			Name:         "CoreNetworkSecurityGroup",
			Dependencies: acctest.DependencyGraph["networkSecurityGroup"],
			F:            sweepCoreNetworkSecurityGroupResource,
		})
	}
}

func sweepCoreNetworkSecurityGroupResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	networkSecurityGroupIds, err := getCoreNetworkSecurityGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, networkSecurityGroupId := range networkSecurityGroupIds {
		if ok := acctest.SweeperDefaultResourceId[networkSecurityGroupId]; !ok {
			deleteNetworkSecurityGroupRequest := oci_core.DeleteNetworkSecurityGroupRequest{}

			deleteNetworkSecurityGroupRequest.NetworkSecurityGroupId = &networkSecurityGroupId

			deleteNetworkSecurityGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteNetworkSecurityGroup(context.Background(), deleteNetworkSecurityGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkSecurityGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkSecurityGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &networkSecurityGroupId, CoreNetworkSecurityGroupSweepWaitCondition, time.Duration(3*time.Minute),
				CoreNetworkSecurityGroupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreNetworkSecurityGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkSecurityGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listNetworkSecurityGroupsRequest := oci_core.ListNetworkSecurityGroupsRequest{}
	listNetworkSecurityGroupsRequest.CompartmentId = &compartmentId
	listNetworkSecurityGroupsRequest.LifecycleState = oci_core.NetworkSecurityGroupLifecycleStateAvailable
	listNetworkSecurityGroupsResponse, err := virtualNetworkClient.ListNetworkSecurityGroups(context.Background(), listNetworkSecurityGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NetworkSecurityGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, networkSecurityGroup := range listNetworkSecurityGroupsResponse.Items {
		id := *networkSecurityGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkSecurityGroupId", id)
	}
	return resourceIds, nil
}

func CoreNetworkSecurityGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if networkSecurityGroupResponse, ok := response.Response.(oci_core.GetNetworkSecurityGroupResponse); ok {
		return networkSecurityGroupResponse.LifecycleState != oci_core.NetworkSecurityGroupLifecycleStateTerminated
	}
	return false
}

func CoreNetworkSecurityGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetNetworkSecurityGroup(context.Background(), oci_core.GetNetworkSecurityGroupRequest{
		NetworkSecurityGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
