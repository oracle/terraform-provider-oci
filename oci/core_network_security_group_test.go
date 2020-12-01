// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v30/common"
	oci_core "github.com/oracle/oci-go-sdk/v30/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NetworkSecurityGroupRequiredOnlyResource = NetworkSecurityGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation)

	NetworkSecurityGroupResourceConfig = NetworkSecurityGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Optional, Update, networkSecurityGroupRepresentation)

	networkSecurityGroupSingularDataSourceRepresentation = map[string]interface{}{
		"network_security_group_id": Representation{repType: Required, create: `${oci_core_network_security_group.test_network_security_group.id}`},
	}

	networkSecurityGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"vcn_id":         Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, networkSecurityGroupDataSourceFilterRepresentation}}
	networkSecurityGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
	}

	networkSecurityGroupRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"vcn_id":         Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	NetworkSecurityGroupResourceDependencies = generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestCoreNetworkSecurityGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreNetworkSecurityGroupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_network_security_group.test_network_security_group"
	datasourceName := "data.oci_core_network_security_groups.test_network_security_groups"
	singularDatasourceName := "data.oci_core_network_security_group.test_network_security_group"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreNetworkSecurityGroupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + NetworkSecurityGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + NetworkSecurityGroupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + NetworkSecurityGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Optional, Create, networkSecurityGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NetworkSecurityGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Optional, Create,
						representationCopyWithNewProperties(networkSecurityGroupRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + NetworkSecurityGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Optional, Update, networkSecurityGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
					generateDataSourceFromRepresentationMap("oci_core_network_security_groups", "test_network_security_groups", Optional, Update, networkSecurityGroupDataSourceRepresentation) +
					compartmentIdVariableStr + NetworkSecurityGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Optional, Update, networkSecurityGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "network_security_groups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "network_security_groups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "network_security_groups.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + NetworkSecurityGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "network_security_group_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + NetworkSecurityGroupResourceConfig,
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

func testAccCheckCoreNetworkSecurityGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_network_security_group" {
			noResourceFound = false
			request := oci_core.GetNetworkSecurityGroupRequest{}

			tmp := rs.Primary.ID
			request.NetworkSecurityGroupId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreNetworkSecurityGroup") {
		resource.AddTestSweepers("CoreNetworkSecurityGroup", &resource.Sweeper{
			Name:         "CoreNetworkSecurityGroup",
			Dependencies: DependencyGraph["networkSecurityGroup"],
			F:            sweepCoreNetworkSecurityGroupResource,
		})
	}
}

func sweepCoreNetworkSecurityGroupResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	networkSecurityGroupIds, err := getNetworkSecurityGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, networkSecurityGroupId := range networkSecurityGroupIds {
		if ok := SweeperDefaultResourceId[networkSecurityGroupId]; !ok {
			deleteNetworkSecurityGroupRequest := oci_core.DeleteNetworkSecurityGroupRequest{}

			deleteNetworkSecurityGroupRequest.NetworkSecurityGroupId = &networkSecurityGroupId

			deleteNetworkSecurityGroupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteNetworkSecurityGroup(context.Background(), deleteNetworkSecurityGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkSecurityGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkSecurityGroupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &networkSecurityGroupId, networkSecurityGroupSweepWaitCondition, time.Duration(3*time.Minute),
				networkSecurityGroupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getNetworkSecurityGroupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "NetworkSecurityGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "NetworkSecurityGroupId", id)
	}
	return resourceIds, nil
}

func networkSecurityGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if networkSecurityGroupResponse, ok := response.Response.(oci_core.GetNetworkSecurityGroupResponse); ok {
		return networkSecurityGroupResponse.LifecycleState != oci_core.NetworkSecurityGroupLifecycleStateTerminated
	}
	return false
}

func networkSecurityGroupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetNetworkSecurityGroup(context.Background(), oci_core.GetNetworkSecurityGroupRequest{
		NetworkSecurityGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
