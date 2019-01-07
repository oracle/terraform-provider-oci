// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

var (
	AutonomousDataWarehouseRequiredOnlyResource = AutonomousDataWarehouseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_data_warehouse", "test_autonomous_data_warehouse", Required, Create, autonomousDataWarehouseRepresentation)

	AutonomousDataWarehouseResourceConfig = AutonomousDataWarehouseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_data_warehouse", "test_autonomous_data_warehouse", Optional, Update, autonomousDataWarehouseRepresentation)

	autonomousDataWarehouseSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_data_warehouse_id": Representation{repType: Required, create: `${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}`},
	}

	autonomousDataWarehouseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `example_autonomous_data_warehouse`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, autonomousDataWarehouseDataSourceFilterRepresentation}}
	autonomousDataWarehouseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}`}},
	}

	adwName = GenerateTestResourceName("adw", 14)

	autonomousDataWarehouseRepresentation = map[string]interface{}{
		"admin_password":           Representation{repType: Required, create: `BEstrO0ng_#11`, update: `BEstrO0ng_#12`},
		"compartment_id":           Representation{repType: Required, create: `${var.compartment_id}`},
		"cpu_core_count":           Representation{repType: Required, create: `1`},
		"data_storage_size_in_tbs": Representation{repType: Required, create: `1`},
		"db_name":                  Representation{repType: Required, create: adwName},
		"defined_tags":             Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":             Representation{repType: Optional, create: `example_autonomous_data_warehouse`, update: `displayName2`},
		"freeform_tags":            Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"license_model":            Representation{repType: Optional, create: `LICENSE_INCLUDED`},
	}

	AutonomousDataWarehouseResourceDependencies = DefinedTagsDependencies
)

func TestDatabaseAutonomousDataWarehouseResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse"
	datasourceName := "data.oci_database_autonomous_data_warehouses.test_autonomous_data_warehouses"
	singularDatasourceName := "data.oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDataWarehouseDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousDataWarehouseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_data_warehouse", "test_autonomous_data_warehouse", Required, Create, autonomousDataWarehouseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adwName),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDataWarehouseResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDataWarehouseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_data_warehouse", "test_autonomous_data_warehouse", Optional, Create, autonomousDataWarehouseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adwName),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_data_warehouse"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutonomousDataWarehouseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_data_warehouse", "test_autonomous_data_warehouse", Optional, Update, autonomousDataWarehouseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adwName),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_data_warehouses", "test_autonomous_data_warehouses", Optional, Update, autonomousDataWarehouseDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDataWarehouseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_data_warehouse", "test_autonomous_data_warehouse", Optional, Update, autonomousDataWarehouseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouses.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouses.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouses.0.cpu_core_count", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouses.0.data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouses.0.db_name", adwName),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouses.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouses.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouses.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouses.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouses.0.license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouses.0.state"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_data_warehouse", "test_autonomous_data_warehouse", Required, Create, autonomousDataWarehouseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDataWarehouseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_data_warehouse_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.0.all_connection_strings.%", "4"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_strings.0.high"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_strings.0.low"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_strings.0.medium"),
					resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "db_name", adwName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + AutonomousDataWarehouseResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"admin_password",
					"lifecycle_details",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckDatabaseAutonomousDataWarehouseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_data_warehouse" {
			noResourceFound = false
			request := oci_database.GetAutonomousDataWarehouseRequest{}

			tmp := rs.Primary.ID
			request.AutonomousDataWarehouseId = &tmp

			response, err := client.GetAutonomousDataWarehouse(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AutonomousDataWarehouseLifecycleStateTerminated): true,
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

func initDatabaseAutonomousDataWarehouseSweeper() {
	resource.AddTestSweepers("DatabaseAutonomousDataWarehouse", &resource.Sweeper{
		Name:         "DatabaseAutonomousDataWarehouse",
		Dependencies: DependencyGraph["autonomousDataWarehouse"],
		F:            sweepDatabaseAutonomousDataWarehouseResource,
	})
}

func sweepDatabaseAutonomousDataWarehouseResource(compartment string) error {
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient

	listAutonomousDataWarehousesRequest := oci_database.ListAutonomousDataWarehousesRequest{}
	listAutonomousDataWarehousesRequest.CompartmentId = &compartmentId
	listAutonomousDataWarehousesRequest.LifecycleState = oci_database.AutonomousDataWarehouseSummaryLifecycleStateAvailable
	listAutonomousDataWarehousesResponse, err := databaseClient.ListAutonomousDataWarehouses(context.Background(), listAutonomousDataWarehousesRequest)

	if err != nil {
		return fmt.Errorf("Error getting AutonomousDataWarehouse list for compartment id : %s , %s \n", compartmentId, err)
	}

	for _, autonomousDataWarehouse := range listAutonomousDataWarehousesResponse.Items {
		if autonomousDataWarehouse.LifecycleState != oci_database.AutonomousDataWarehouseSummaryLifecycleStateTerminated {
			log.Printf("deleting autonomousDataWarehouse %s ", *autonomousDataWarehouse.Id)

			deleteAutonomousDataWarehouseRequest := oci_database.DeleteAutonomousDataWarehouseRequest{}

			deleteAutonomousDataWarehouseRequest.AutonomousDataWarehouseId = autonomousDataWarehouse.Id

			deleteAutonomousDataWarehouseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteAutonomousDataWarehouse(context.Background(), deleteAutonomousDataWarehouseRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousDataWarehouse %s %s, It is possible that the resource is already deleted. Please verify manually \n", *autonomousDataWarehouse.Id, error)
				continue
			}

			getAutonomousDataWarehouseRequest := oci_database.GetAutonomousDataWarehouseRequest{}

			getAutonomousDataWarehouseRequest.AutonomousDataWarehouseId = autonomousDataWarehouse.Id

			_, error = databaseClient.GetAutonomousDataWarehouse(context.Background(), getAutonomousDataWarehouseRequest)
			if error != nil {
				fmt.Printf("Error retrieving AutonomousDataWarehouse state %s \n", error)
				continue
			}

			waitTillCondition(testAccProvider, autonomousDataWarehouse.Id, autonomousDataWarehouseSweepWaitCondition, time.Duration(3*time.Minute),
				autonomousDataWarehouseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func autonomousDataWarehouseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousDataWarehouseResponse, ok := response.Response.(oci_database.GetAutonomousDataWarehouseResponse); ok {
		return autonomousDataWarehouseResponse.LifecycleState == oci_database.AutonomousDataWarehouseLifecycleStateTerminated
	}
	return false
}

func autonomousDataWarehouseSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient.GetAutonomousDataWarehouse(context.Background(), oci_database.GetAutonomousDataWarehouseRequest{
		AutonomousDataWarehouseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
