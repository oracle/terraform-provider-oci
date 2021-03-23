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
	"github.com/oracle/oci-go-sdk/v37/common"
	oci_database "github.com/oracle/oci-go-sdk/v37/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExternalNonContainerDatabaseRequiredOnlyResource = ExternalNonContainerDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Required, Create, externalNonContainerDatabaseRepresentation)

	ExternalNonContainerDatabaseResourceConfig = ExternalNonContainerDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Optional, Update, externalNonContainerDatabaseRepresentation)

	externalNonContainerDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"external_non_container_database_id": Representation{repType: Required, create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
	}

	externalNonContainerDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `myTestExternalNonCdb`},
		"state":          Representation{repType: Optional, create: `NOT_CONNECTED`},
		"filter":         RepresentationGroup{Required, externalNonContainerDatabaseDataSourceFilterRepresentation}}
	externalNonContainerDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_external_non_container_database.test_external_non_container_database.id}`}},
	}

	externalNonContainerDatabaseRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `myTestExternalNonCdb`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	ExternalNonContainerDatabaseResourceDependencies = DefinedTagsDependencies
)

func TestDatabaseExternalNonContainerDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalNonContainerDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_external_non_container_database.test_external_non_container_database"
	datasourceName := "data.oci_database_external_non_container_databases.test_external_non_container_databases"
	singularDatasourceName := "data.oci_database_external_non_container_database.test_external_non_container_database"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ExternalNonContainerDatabaseResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Optional, Create, externalNonContainerDatabaseRepresentation), "database", "externalNonContainerDatabase", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseExternalNonContainerDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Required, Create, externalNonContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalNonCdb"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Optional, Create, externalNonContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalNonCdb"),
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ExternalNonContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Optional, Create,
						representationCopyWithNewProperties(externalNonContainerDatabaseRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalNonCdb"),
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
				Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Optional, Update, externalNonContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalNonCdb"),
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
					generateDataSourceFromRepresentationMap("oci_database_external_non_container_databases", "test_external_non_container_databases", Optional, Update, externalNonContainerDatabaseDataSourceRepresentation) +
					compartmentIdVariableStr + ExternalNonContainerDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Optional, Update, externalNonContainerDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "myTestExternalNonCdb"),
					resource.TestCheckResourceAttr(datasourceName, "state", "NOT_CONNECTED"),

					resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.0.database_management_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.0.display_name", "myTestExternalNonCdb"),
					resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_non_container_databases.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_non_container_databases.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_non_container_databases.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Required, Create, externalNonContainerDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ExternalNonContainerDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "external_non_container_database_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "database_management_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "myTestExternalNonCdb"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseResourceConfig,
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

func testAccCheckDatabaseExternalNonContainerDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_external_non_container_database" {
			noResourceFound = false
			request := oci_database.GetExternalNonContainerDatabaseRequest{}

			tmp := rs.Primary.ID
			request.ExternalNonContainerDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetExternalNonContainerDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExternalNonContainerDatabaseLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("DatabaseExternalNonContainerDatabase") {
		resource.AddTestSweepers("DatabaseExternalNonContainerDatabase", &resource.Sweeper{
			Name:         "DatabaseExternalNonContainerDatabase",
			Dependencies: DependencyGraph["externalNonContainerDatabase"],
			F:            sweepDatabaseExternalNonContainerDatabaseResource,
		})
	}
}

func sweepDatabaseExternalNonContainerDatabaseResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	externalNonContainerDatabaseIds, err := getExternalNonContainerDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, externalNonContainerDatabaseId := range externalNonContainerDatabaseIds {
		if ok := SweeperDefaultResourceId[externalNonContainerDatabaseId]; !ok {
			deleteExternalNonContainerDatabaseRequest := oci_database.DeleteExternalNonContainerDatabaseRequest{}

			deleteExternalNonContainerDatabaseRequest.ExternalNonContainerDatabaseId = &externalNonContainerDatabaseId

			deleteExternalNonContainerDatabaseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExternalNonContainerDatabase(context.Background(), deleteExternalNonContainerDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalNonContainerDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalNonContainerDatabaseId, error)
				continue
			}
			waitTillCondition(testAccProvider, &externalNonContainerDatabaseId, externalNonContainerDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				externalNonContainerDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getExternalNonContainerDatabaseIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ExternalNonContainerDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listExternalNonContainerDatabasesRequest := oci_database.ListExternalNonContainerDatabasesRequest{}
	listExternalNonContainerDatabasesRequest.CompartmentId = &compartmentId
	listExternalNonContainerDatabasesRequest.LifecycleState = oci_database.ExternalDatabaseBaseLifecycleStateAvailable
	listExternalNonContainerDatabasesResponse, err := databaseClient.ListExternalNonContainerDatabases(context.Background(), listExternalNonContainerDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalNonContainerDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalNonContainerDatabase := range listExternalNonContainerDatabasesResponse.Items {
		id := *externalNonContainerDatabase.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ExternalNonContainerDatabaseId", id)
	}
	return resourceIds, nil
}

func externalNonContainerDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalNonContainerDatabaseResponse, ok := response.Response.(oci_database.GetExternalNonContainerDatabaseResponse); ok {
		return externalNonContainerDatabaseResponse.LifecycleState != oci_database.ExternalNonContainerDatabaseLifecycleStateTerminated
	}
	return false
}

func externalNonContainerDatabaseSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetExternalNonContainerDatabase(context.Background(), oci_database.GetExternalNonContainerDatabaseRequest{
		ExternalNonContainerDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
