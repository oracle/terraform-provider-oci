// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var (
	CapacityManagementOccCustomerGroupRequiredOnlyResource = CapacityManagementOccCustomerGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group", "test_occ_customer_group", acctest.Required, acctest.Create, CapacityManagementOccCustomerGroupRepresentation)
	CapacityManagementOccCustomerGroupResourceConfig = CapacityManagementOccCustomerGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group", "test_occ_customer_group", acctest.Optional, acctest.Update, CapacityManagementOccCustomerGroupRepresentation)

	occCustomerGroupDisplayName        = "TersiCg"
	occCustomerGroupUpdatedDisplayName = "Updated TERSI Test"
	occCustomerGroupDescription        = "This is a Customer group created via Terraform"

	CapacityManagementOccCustomerGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: occCustomerGroupDisplayName, Update: occCustomerGroupUpdatedDisplayName},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: occCustomerGroupDescription},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
	}

	CapacityManagementOccCustomerGroupSingularDataSourceRepresentation = map[string]interface{}{
		"occ_customer_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_occ_customer_group.test_occ_customer_group.id}`},
	}

	CapacityManagementOccCustomerGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: occCustomerGroupUpdatedDisplayName},
	}

	CapacityManagementOccCustomerGroupResourceDependencies = ""
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccCustomerGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccCustomerGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	occCustomerGroupId := utils.GetEnvSettingWithBlankDefault("occ_customer_group_id")
	occCustomerGroupIdVariableStr := fmt.Sprintf("variable \"occ_customer_group_id\" { default = \"%s\" }\n", occCustomerGroupId)

	resourceName := "oci_capacity_management_occ_customer_group.test_occ_customer_group"
	datasourceName := "data.oci_capacity_management_occ_customer_groups.test_occ_customer_groups"
	singularDatasourceName := "data.oci_capacity_management_occ_customer_group.test_occ_customer_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CapacityManagementOccCustomerGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group", "test_occ_customer_group", acctest.Optional, acctest.Create, CapacityManagementOccCustomerGroupRepresentation), "capacitymanagement", "occCustomerGroup", t)

	acctest.ResourceTest(t, testAccCheckCapacityManagementOccCustomerGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CapacityManagementOccCustomerGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group", "test_occ_customer_group", acctest.Required, acctest.Create, CapacityManagementOccCustomerGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", occCustomerGroupDisplayName),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CapacityManagementOccCustomerGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CapacityManagementOccCustomerGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group", "test_occ_customer_group", acctest.Optional, acctest.Create, CapacityManagementOccCustomerGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "customers_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "description", occCustomerGroupDescription),
				resource.TestCheckResourceAttr(resourceName, "display_name", occCustomerGroupDisplayName),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CapacityManagementOccCustomerGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_customer_group", "test_occ_customer_group", acctest.Optional, acctest.Update, CapacityManagementOccCustomerGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "customers_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "description", occCustomerGroupDescription),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Updated TERSI Test"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_customer_groups", "test_occ_customer_groups", acctest.Required, acctest.Create, CapacityManagementOccCustomerGroupDataSourceRepresentation) +
				compartmentIdVariableStr + CapacityManagementOccCustomerGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "occ_customer_group_collection.#", "1"),
				// TODO: Update when the test fails accordingly
				resource.TestCheckResourceAttr(datasourceName, "occ_customer_group_collection.0.items.#", "28"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_customer_group", "test_occ_customer_group", acctest.Optional, acctest.Create, CapacityManagementOccCustomerGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccCustomerGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", occCustomerGroupUpdatedDisplayName),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", occCustomerGroupDescription),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "customers_list.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + CapacityManagementOccCustomerGroupResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCapacityManagementOccCustomerGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CapacityManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_capacity_management_occ_customer_group" {
			noResourceFound = false
			request := oci_capacity_management.GetOccCustomerGroupRequest{}

			tmp := rs.Primary.ID
			request.OccCustomerGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")

			response, err := client.GetOccCustomerGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_capacity_management.OccCustomerGroupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CapacityManagementOccCustomerGroup") {
		resource.AddTestSweepers("CapacityManagementOccCustomerGroup", &resource.Sweeper{
			Name:         "CapacityManagementOccCustomerGroup",
			Dependencies: acctest.DependencyGraph["occCustomerGroup"],
			F:            sweepCapacityManagementOccCustomerGroupResource,
		})
	}
}

func sweepCapacityManagementOccCustomerGroupResource(compartment string) error {
	capacityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).CapacityManagementClient()
	occCustomerGroupIds, err := getCapacityManagementOccCustomerGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, occCustomerGroupId := range occCustomerGroupIds {
		if ok := acctest.SweeperDefaultResourceId[occCustomerGroupId]; !ok {
			deleteOccCustomerGroupRequest := oci_capacity_management.DeleteOccCustomerGroupRequest{}

			deleteOccCustomerGroupRequest.OccCustomerGroupId = &occCustomerGroupId

			deleteOccCustomerGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")
			_, error := capacityManagementClient.DeleteOccCustomerGroup(context.Background(), deleteOccCustomerGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting OccCustomerGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", occCustomerGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &occCustomerGroupId, CapacityManagementOccCustomerGroupSweepWaitCondition, time.Duration(3*time.Minute),
				CapacityManagementOccCustomerGroupSweepResponseFetchOperation, "capacity_management", true)
		}
	}
	return nil
}

func getCapacityManagementOccCustomerGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OccCustomerGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	capacityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).CapacityManagementClient()

	listOccCustomerGroupsRequest := oci_capacity_management.ListOccCustomerGroupsRequest{}
	listOccCustomerGroupsRequest.CompartmentId = &compartmentId
	listOccCustomerGroupsResponse, err := capacityManagementClient.ListOccCustomerGroups(context.Background(), listOccCustomerGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OccCustomerGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, occCustomerGroup := range listOccCustomerGroupsResponse.Items {
		id := *occCustomerGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OccCustomerGroupId", id)
	}
	return resourceIds, nil
}

func CapacityManagementOccCustomerGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if occCustomerGroupResponse, ok := response.Response.(oci_capacity_management.GetOccCustomerGroupResponse); ok {
		return occCustomerGroupResponse.LifecycleState != oci_capacity_management.OccCustomerGroupLifecycleStateDeleted
	}
	return false
}

func CapacityManagementOccCustomerGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CapacityManagementClient().GetOccCustomerGroup(context.Background(), oci_capacity_management.GetOccCustomerGroupRequest{
		OccCustomerGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
