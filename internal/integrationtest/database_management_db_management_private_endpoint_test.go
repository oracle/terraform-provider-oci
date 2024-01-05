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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseManagementDbManagementPrivateEndpointRequiredOnlyResource = DatabaseManagementDbManagementPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, DatabaseManagementDbManagementPrivateEndpointRepresentation)

	DatabaseManagementDbManagementPrivateEndpointResourceConfig = DatabaseManagementDbManagementPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Optional, acctest.Update, DatabaseManagementDbManagementPrivateEndpointRepresentation)

	DatabaseManagementDatabaseManagementDbManagementPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"db_management_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id}`},
	}

	DatabaseManagementDatabaseManagementDbManagementPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_cluster":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementDbManagementPrivateEndpointDataSourceFilterRepresentation}}
	DatabaseManagementDbManagementPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id}`}},
	}

	DatabaseManagementDbManagementPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"is_cluster":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
	}

	DatabaseManagementDbManagementPrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementDbManagementPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementDbManagementPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint"
	datasourceName := "data.oci_database_management_db_management_private_endpoints.test_db_management_private_endpoints"
	singularDatasourceName := "data.oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseManagementDbManagementPrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Optional, acctest.Create, DatabaseManagementDbManagementPrivateEndpointRepresentation), "databasemanagement", "dbManagementPrivateEndpoint", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseManagementDbManagementPrivateEndpointDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + DatabaseManagementDbManagementPrivateEndpointResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, DatabaseManagementDbManagementPrivateEndpointRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "is_cluster", "false"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + DatabaseManagementDbManagementPrivateEndpointResourceDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + DatabaseManagementDbManagementPrivateEndpointResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Optional, acctest.Create, DatabaseManagementDbManagementPrivateEndpointRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_cluster", "false"),

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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseManagementDbManagementPrivateEndpointResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(DatabaseManagementDbManagementPrivateEndpointRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttr(resourceName, "is_cluster", "false"),

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
				Config: config + compartmentIdVariableStr + DatabaseManagementDbManagementPrivateEndpointResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Optional, acctest.Update, DatabaseManagementDbManagementPrivateEndpointRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "is_cluster", "false"),

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
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_db_management_private_endpoints", "test_db_management_private_endpoints", acctest.Optional, acctest.Update, DatabaseManagementDatabaseManagementDbManagementPrivateEndpointDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseManagementDbManagementPrivateEndpointResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Optional, acctest.Update, DatabaseManagementDbManagementPrivateEndpointRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),
					resource.TestCheckResourceAttr(datasourceName, "db_management_private_endpoint_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "db_management_private_endpoint_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "is_cluster", "false"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementDbManagementPrivateEndpointSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseManagementDbManagementPrivateEndpointResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_management_private_endpoint_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_cluster", "false"),
				),
			},
			// verify resource import
			{
				Config:                  config + DatabaseManagementDbManagementPrivateEndpointRequiredOnlyResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckDatabaseManagementDbManagementPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_db_management_private_endpoint" {
			noResourceFound = false
			request := oci_database_management.GetDbManagementPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.DbManagementPrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetDbManagementPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementDbManagementPrivateEndpoint") {
		resource.AddTestSweepers("DatabaseManagementDbManagementPrivateEndpoint", &resource.Sweeper{
			Name:         "DatabaseManagementDbManagementPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["dbManagementPrivateEndpoint"],
			F:            sweepDatabaseManagementDbManagementPrivateEndpointResource,
		})
	}
}

func sweepDatabaseManagementDbManagementPrivateEndpointResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	dbManagementPrivateEndpointIds, err := getDatabaseManagementDbManagementPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, dbManagementPrivateEndpointId := range dbManagementPrivateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[dbManagementPrivateEndpointId]; !ok {
			deleteDbManagementPrivateEndpointRequest := oci_database_management.DeleteDbManagementPrivateEndpointRequest{}

			deleteDbManagementPrivateEndpointRequest.DbManagementPrivateEndpointId = &dbManagementPrivateEndpointId

			deleteDbManagementPrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteDbManagementPrivateEndpoint(context.Background(), deleteDbManagementPrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting DbManagementPrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", dbManagementPrivateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dbManagementPrivateEndpointId, DatabaseManagementDbManagementPrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementDbManagementPrivateEndpointSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementDbManagementPrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DbManagementPrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listDbManagementPrivateEndpointsRequest := oci_database_management.ListDbManagementPrivateEndpointsRequest{}
	listDbManagementPrivateEndpointsRequest.CompartmentId = &compartmentId
	listDbManagementPrivateEndpointsRequest.LifecycleState = oci_database_management.ListDbManagementPrivateEndpointsLifecycleStateActive
	listDbManagementPrivateEndpointsResponse, err := dbManagementClient.ListDbManagementPrivateEndpoints(context.Background(), listDbManagementPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DbManagementPrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dbManagementPrivateEndpoint := range listDbManagementPrivateEndpointsResponse.Items {
		id := *dbManagementPrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DbManagementPrivateEndpointId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementDbManagementPrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dbManagementPrivateEndpointResponse, ok := response.Response.(oci_database_management.GetDbManagementPrivateEndpointResponse); ok {
		return dbManagementPrivateEndpointResponse.LifecycleState != oci_database_management.LifecycleStatesDeleted
	}
	return false
}

func DatabaseManagementDbManagementPrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetDbManagementPrivateEndpoint(context.Background(), oci_database_management.GetDbManagementPrivateEndpointRequest{
		DbManagementPrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
