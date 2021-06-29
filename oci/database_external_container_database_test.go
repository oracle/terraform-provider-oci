// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v43/common"
	oci_database "github.com/oracle/oci-go-sdk/v43/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExternalContainerDatabaseRequiredOnlyResource = ExternalContainerDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Required, Create, externalContainerDatabaseRepresentation)

	ExternalContainerDatabaseResourceConfig = ExternalContainerDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Optional, Update, externalContainerDatabaseRepresentation)

	externalContainerDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"external_container_database_id": Representation{repType: Required, create: `${oci_database_external_container_database.test_external_container_database.id}`},
	}

	externalContainerDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `myTestExternalCdb`},
		"state":          Representation{repType: Optional, create: `NOT_CONNECTED`},
		"filter":         RepresentationGroup{Required, externalContainerDatabaseDataSourceFilterRepresentation}}
	externalContainerDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_external_container_database.test_external_container_database.id}`}},
	}

	externalContainerDatabaseRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `myTestExternalCdb`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	ExternalContainerDatabaseResourceDependencies = DefinedTagsDependencies
)

func TestDatabaseExternalContainerDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalContainerDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_external_container_database.test_external_container_database"
	datasourceName := "data.oci_database_external_container_databases.test_external_container_databases"
	singularDatasourceName := "data.oci_database_external_container_database.test_external_container_database"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ExternalContainerDatabaseResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Optional, Create, externalContainerDatabaseRepresentation), "database", "externalContainerDatabase", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseExternalContainerDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ExternalContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Required, Create, externalContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalCdb"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ExternalContainerDatabaseResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ExternalContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Optional, Create, externalContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalCdb"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ExternalContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Optional, Create,
						representationCopyWithNewProperties(externalContainerDatabaseRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalCdb"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				Config: config + compartmentIdVariableStr + ExternalContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Optional, Update, externalContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalCdb"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
					generateDataSourceFromRepresentationMap("oci_database_external_container_databases", "test_external_container_databases", Optional, Update, externalContainerDatabaseDataSourceRepresentation) +
					compartmentIdVariableStr + ExternalContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Optional, Update, externalContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "myTestExternalCdb"),
					resource.TestCheckResourceAttr(datasourceName, "state", "NOT_CONNECTED"),

					resource.TestCheckResourceAttr(datasourceName, "external_container_databases.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "external_container_databases.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "external_container_databases.0.database_management_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "external_container_databases.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "external_container_databases.0.display_name", "myTestExternalCdb"),
					resource.TestCheckResourceAttr(datasourceName, "external_container_databases.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_container_databases.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_container_databases.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_container_databases.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Required, Create, externalContainerDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ExternalContainerDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "external_container_database_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "database_management_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "myTestExternalCdb"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ExternalContainerDatabaseResourceConfig,
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

func testAccCheckDatabaseExternalContainerDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_external_container_database" {
			noResourceFound = false
			request := oci_database.GetExternalContainerDatabaseRequest{}

			tmp := rs.Primary.ID
			request.ExternalContainerDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetExternalContainerDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExternalContainerDatabaseLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("DatabaseExternalContainerDatabase") {
		resource.AddTestSweepers("DatabaseExternalContainerDatabase", &resource.Sweeper{
			Name:         "DatabaseExternalContainerDatabase",
			Dependencies: DependencyGraph["externalContainerDatabase"],
			F:            sweepDatabaseExternalContainerDatabaseResource,
		})
	}
}

func sweepDatabaseExternalContainerDatabaseResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	externalContainerDatabaseIds, err := getExternalContainerDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, externalContainerDatabaseId := range externalContainerDatabaseIds {
		if ok := SweeperDefaultResourceId[externalContainerDatabaseId]; !ok {
			deleteExternalContainerDatabaseRequest := oci_database.DeleteExternalContainerDatabaseRequest{}

			deleteExternalContainerDatabaseRequest.ExternalContainerDatabaseId = &externalContainerDatabaseId

			deleteExternalContainerDatabaseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExternalContainerDatabase(context.Background(), deleteExternalContainerDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalContainerDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalContainerDatabaseId, error)
				continue
			}
			waitTillCondition(testAccProvider, &externalContainerDatabaseId, externalContainerDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				externalContainerDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getExternalContainerDatabaseIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ExternalContainerDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listExternalContainerDatabasesRequest := oci_database.ListExternalContainerDatabasesRequest{}
	listExternalContainerDatabasesRequest.CompartmentId = &compartmentId
	listExternalContainerDatabasesRequest.LifecycleState = oci_database.ExternalDatabaseBaseLifecycleStateAvailable
	listExternalContainerDatabasesResponse, err := databaseClient.ListExternalContainerDatabases(context.Background(), listExternalContainerDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalContainerDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalContainerDatabase := range listExternalContainerDatabasesResponse.Items {
		id := *externalContainerDatabase.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ExternalContainerDatabaseId", id)
	}
	return resourceIds, nil
}

func externalContainerDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalContainerDatabaseResponse, ok := response.Response.(oci_database.GetExternalContainerDatabaseResponse); ok {
		return externalContainerDatabaseResponse.LifecycleState != oci_database.ExternalContainerDatabaseLifecycleStateTerminated
	}
	return false
}

func externalContainerDatabaseSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetExternalContainerDatabase(context.Background(), oci_database.GetExternalContainerDatabaseRequest{
		ExternalContainerDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
