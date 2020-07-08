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
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CrossConnectGroupRequiredOnlyResource = CrossConnectGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", Required, Create, crossConnectGroupRepresentation)

	CrossConnectGroupResourceConfig = CrossConnectGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", Optional, Update, crossConnectGroupRepresentation)

	crossConnectGroupSingularDataSourceRepresentation = map[string]interface{}{
		"cross_connect_group_id": Representation{repType: Required, create: `${oci_core_cross_connect_group.test_cross_connect_group.id}`},
	}

	crossConnectGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"filter":         RepresentationGroup{Required, crossConnectGroupDataSourceFilterRepresentation}}
	crossConnectGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_cross_connect_group.test_cross_connect_group.id}`}},
	}

	crossConnectGroupRepresentation = map[string]interface{}{
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"customer_reference_name": Representation{repType: Optional, create: `customerReferenceName`, update: `customerReferenceName2`},
		"defined_tags":            Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":           Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	CrossConnectGroupResourceDependencies = DefinedTagsDependencies
)

func TestCoreCrossConnectGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCrossConnectGroupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_cross_connect_group.test_cross_connect_group"
	datasourceName := "data.oci_core_cross_connect_groups.test_cross_connect_groups"
	singularDatasourceName := "data.oci_core_cross_connect_group.test_cross_connect_group"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreCrossConnectGroupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + CrossConnectGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", Required, Create, crossConnectGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CrossConnectGroupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CrossConnectGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", Optional, Create, crossConnectGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CrossConnectGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", Optional, Create,
						representationCopyWithNewProperties(crossConnectGroupRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

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
				Config: config + compartmentIdVariableStr + CrossConnectGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", Optional, Update, crossConnectGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName2"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

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
					generateDataSourceFromRepresentationMap("oci_core_cross_connect_groups", "test_cross_connect_groups", Optional, Update, crossConnectGroupDataSourceRepresentation) +
					compartmentIdVariableStr + CrossConnectGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", Optional, Update, crossConnectGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

					resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.customer_reference_name", "customerReferenceName2"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_groups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_groups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_groups.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", Required, Create, crossConnectGroupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CrossConnectGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_group_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "customer_reference_name", "customerReferenceName2"),
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
				Config: config + compartmentIdVariableStr + CrossConnectGroupResourceConfig,
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

func testAccCheckCoreCrossConnectGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_cross_connect_group" {
			noResourceFound = false
			request := oci_core.GetCrossConnectGroupRequest{}

			tmp := rs.Primary.ID
			request.CrossConnectGroupId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetCrossConnectGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.CrossConnectGroupLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("CoreCrossConnectGroup") {
		resource.AddTestSweepers("CoreCrossConnectGroup", &resource.Sweeper{
			Name:         "CoreCrossConnectGroup",
			Dependencies: DependencyGraph["crossConnectGroup"],
			F:            sweepCoreCrossConnectGroupResource,
		})
	}
}

func sweepCoreCrossConnectGroupResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	crossConnectGroupIds, err := getCrossConnectGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, crossConnectGroupId := range crossConnectGroupIds {
		if ok := SweeperDefaultResourceId[crossConnectGroupId]; !ok {
			deleteCrossConnectGroupRequest := oci_core.DeleteCrossConnectGroupRequest{}

			deleteCrossConnectGroupRequest.CrossConnectGroupId = &crossConnectGroupId

			deleteCrossConnectGroupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteCrossConnectGroup(context.Background(), deleteCrossConnectGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting CrossConnectGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", crossConnectGroupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &crossConnectGroupId, crossConnectGroupSweepWaitCondition, time.Duration(3*time.Minute),
				crossConnectGroupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCrossConnectGroupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "CrossConnectGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listCrossConnectGroupsRequest := oci_core.ListCrossConnectGroupsRequest{}
	listCrossConnectGroupsRequest.CompartmentId = &compartmentId
	listCrossConnectGroupsRequest.LifecycleState = oci_core.CrossConnectGroupLifecycleStateProvisioned
	listCrossConnectGroupsResponse, err := virtualNetworkClient.ListCrossConnectGroups(context.Background(), listCrossConnectGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CrossConnectGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, crossConnectGroup := range listCrossConnectGroupsResponse.Items {
		id := *crossConnectGroup.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "CrossConnectGroupId", id)
	}
	return resourceIds, nil
}

func crossConnectGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if crossConnectGroupResponse, ok := response.Response.(oci_core.GetCrossConnectGroupResponse); ok {
		return crossConnectGroupResponse.LifecycleState != oci_core.CrossConnectGroupLifecycleStateTerminated
	}
	return false
}

func crossConnectGroupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetCrossConnectGroup(context.Background(), oci_core.GetCrossConnectGroupRequest{
		CrossConnectGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
