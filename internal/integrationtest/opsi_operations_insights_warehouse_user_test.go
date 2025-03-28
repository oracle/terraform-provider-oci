// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OpsiOperationsInsightsWarehouseUserRequiredOnlyResource = OpsiOperationsInsightsWarehouseUserResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_user", "test_operations_insights_warehouse_user", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseUserRepresentation)

	OpsiOperationsInsightsWarehouseUserResourceConfig = OpsiOperationsInsightsWarehouseUserResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_user", "test_operations_insights_warehouse_user", acctest.Optional, acctest.Update, OpsiOperationsInsightsWarehouseUserRepresentation)

	OpsiOpsiOperationsInsightsWarehouseUserSingularDataSourceRepresentation = map[string]interface{}{
		"operations_insights_warehouse_user_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_operations_insights_warehouse_user.test_operations_insights_warehouse_user.id}`},
	}

	OpsiOpsiOperationsInsightsWarehouseUserDataSourceRepresentation = map[string]interface{}{
		"operations_insights_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":                               acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_operations_insights_warehouse_user.test_operations_insights_warehouse_user.id}`},
		"state":                            acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"filter":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiOperationsInsightsWarehouseUserDataSourceFilterRepresentation}}
	OpsiOperationsInsightsWarehouseUserDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_operations_insights_warehouse_user.test_operations_insights_warehouse_user.id}`}},
	}

	OpsiOperationsInsightsWarehouseUserRepresentation = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_password":              acctest.Representation{RepType: acctest.Required, Create: `connectionPassword1`, Update: `connectionPassword2`},
		"is_awr_data_access":               acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"name":                             acctest.Representation{RepType: acctest.Required, Create: `name`},
		"operations_insights_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_em_data_access":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_opsi_data_access":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesOperationsInsightsWarehouseUserRepresentation},
	}

	ignoreChangesOperationsInsightsWarehouseUserRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	OpsiOperationsInsightsWarehouseUserResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiOperationsInsightsWarehouseUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiOperationsInsightsWarehouseUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_opsi_operations_insights_warehouse_user.test_operations_insights_warehouse_user"
	datasourceName := "data.oci_opsi_operations_insights_warehouse_users.test_operations_insights_warehouse_users"
	singularDatasourceName := "data.oci_opsi_operations_insights_warehouse_user.test_operations_insights_warehouse_user"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OpsiOperationsInsightsWarehouseUserResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_user", "test_operations_insights_warehouse_user", acctest.Optional, acctest.Create, OpsiOperationsInsightsWarehouseUserRepresentation), "operationsinsights", "operationsInsightsWarehouseUser", t)

	acctest.ResourceTest(t, testAccCheckOpsiOperationsInsightsWarehouseUserDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + OpsiOperationsInsightsWarehouseUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_user", "test_operations_insights_warehouse_user", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttr(resourceName, "connection_password", "connectionPassword1"), The Response does not return the connection password
				resource.TestCheckResourceAttr(resourceName, "is_awr_data_access", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + OpsiOperationsInsightsWarehouseUserResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + OpsiOperationsInsightsWarehouseUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_user", "test_operations_insights_warehouse_user", acctest.Optional, acctest.Create, OpsiOperationsInsightsWarehouseUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttr(resourceName, "connection_password", "connectionPassword1"), The Response does not return the connection password
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_awr_data_access", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_em_data_access", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_opsi_data_access", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + OpsiOperationsInsightsWarehouseUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_user", "test_operations_insights_warehouse_user", acctest.Optional, acctest.Update, OpsiOperationsInsightsWarehouseUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(resourceName, "connection_password", "connectionPassword2"), The Response does not return the connection password
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_awr_data_access", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_em_data_access", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_opsi_data_access", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_users", "test_operations_insights_warehouse_users", acctest.Optional, acctest.Update, OpsiOpsiOperationsInsightsWarehouseUserDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiOperationsInsightsWarehouseUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_user", "test_operations_insights_warehouse_user", acctest.Optional, acctest.Update, OpsiOperationsInsightsWarehouseUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "operations_insights_warehouse_id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "operations_insights_warehouse_user_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "operations_insights_warehouse_user_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_user", "test_operations_insights_warehouse_user", acctest.Required, acctest.Create, OpsiOpsiOperationsInsightsWarehouseUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiOperationsInsightsWarehouseUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operations_insights_warehouse_user_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttr(singularDatasourceName, "connection_password", "connectionPassword2"), The Response does not return the connection password
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_awr_data_access", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_em_data_access", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_opsi_data_access", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OpsiOperationsInsightsWarehouseUserRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"connection_password"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOpsiOperationsInsightsWarehouseUserDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_operations_insights_warehouse_user" {
			noResourceFound = false
			request := oci_opsi.GetOperationsInsightsWarehouseUserRequest{}

			tmp := rs.Primary.ID
			request.OperationsInsightsWarehouseUserId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetOperationsInsightsWarehouseUser(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.OperationsInsightsWarehouseUserLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpsiOperationsInsightsWarehouseUser") {
		resource.AddTestSweepers("OpsiOperationsInsightsWarehouseUser", &resource.Sweeper{
			Name:         "OpsiOperationsInsightsWarehouseUser",
			Dependencies: acctest.DependencyGraph["operationsInsightsWarehouseUser"],
			F:            sweepOpsiOperationsInsightsWarehouseUserResource,
		})
	}
}

func sweepOpsiOperationsInsightsWarehouseUserResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	operationsInsightsWarehouseUserIds, err := getOpsiOperationsInsightsWarehouseUserIds(compartment)
	if err != nil {
		return err
	}
	for _, operationsInsightsWarehouseUserId := range operationsInsightsWarehouseUserIds {
		if ok := acctest.SweeperDefaultResourceId[operationsInsightsWarehouseUserId]; !ok {
			deleteOperationsInsightsWarehouseUserRequest := oci_opsi.DeleteOperationsInsightsWarehouseUserRequest{}

			deleteOperationsInsightsWarehouseUserRequest.OperationsInsightsWarehouseUserId = &operationsInsightsWarehouseUserId

			deleteOperationsInsightsWarehouseUserRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteOperationsInsightsWarehouseUser(context.Background(), deleteOperationsInsightsWarehouseUserRequest)
			if error != nil {
				fmt.Printf("Error deleting OperationsInsightsWarehouseUser %s %s, It is possible that the resource is already deleted. Please verify manually \n", operationsInsightsWarehouseUserId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &operationsInsightsWarehouseUserId, OpsiOperationsInsightsWarehouseUserSweepWaitCondition, time.Duration(3*time.Minute),
				OpsiOperationsInsightsWarehouseUserSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getOpsiOperationsInsightsWarehouseUserIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OperationsInsightsWarehouseUserId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listOperationsInsightsWarehouseUsersRequest := oci_opsi.ListOperationsInsightsWarehouseUsersRequest{}
	listOperationsInsightsWarehouseUsersRequest.CompartmentId = &compartmentId

	operationsInsightsWarehouseIds, error := getOpsiOperationsInsightsWarehouseIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting operationsInsightsWarehouseId required for OperationsInsightsWarehouseUser resource requests \n")
	}
	for _, operationsInsightsWarehouseId := range operationsInsightsWarehouseIds {
		listOperationsInsightsWarehouseUsersRequest.OperationsInsightsWarehouseId = &operationsInsightsWarehouseId

		listOperationsInsightsWarehouseUsersRequest.LifecycleState = []oci_opsi.OperationsInsightsWarehouseUserLifecycleStateEnum{oci_opsi.OperationsInsightsWarehouseUserLifecycleStateActive}
		listOperationsInsightsWarehouseUsersResponse, err := operationsInsightsClient.ListOperationsInsightsWarehouseUsers(context.Background(), listOperationsInsightsWarehouseUsersRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting OperationsInsightsWarehouseUser list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, operationsInsightsWarehouseUser := range listOperationsInsightsWarehouseUsersResponse.Items {
			id := *operationsInsightsWarehouseUser.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OperationsInsightsWarehouseUserId", id)
		}

	}
	return resourceIds, nil
}

func OpsiOperationsInsightsWarehouseUserSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if operationsInsightsWarehouseUserResponse, ok := response.Response.(oci_opsi.GetOperationsInsightsWarehouseUserResponse); ok {
		return operationsInsightsWarehouseUserResponse.LifecycleState != oci_opsi.OperationsInsightsWarehouseUserLifecycleStateDeleted
	}
	return false
}

func OpsiOperationsInsightsWarehouseUserSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetOperationsInsightsWarehouseUser(context.Background(), oci_opsi.GetOperationsInsightsWarehouseUserRequest{
		OperationsInsightsWarehouseUserId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
