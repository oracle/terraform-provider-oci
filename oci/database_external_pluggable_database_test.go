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
	"github.com/oracle/oci-go-sdk/v36/common"
	oci_database "github.com/oracle/oci-go-sdk/v36/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExternalPluggableDatabaseRequiredOnlyResource = ExternalPluggableDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", Required, Create, externalPluggableDatabaseRepresentation)

	ExternalPluggableDatabaseResourceConfig = ExternalPluggableDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", Optional, Update, externalPluggableDatabaseRepresentation)

	externalPluggableDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"external_pluggable_database_id": Representation{repType: Required, create: `${oci_database_external_pluggable_database.test_external_pluggable_database.id}`},
	}

	externalPluggableDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                 Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                   Representation{repType: Optional, create: `myTestExternalPdb`},
		"external_container_database_id": Representation{repType: Optional, create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"state":                          Representation{repType: Optional, create: `NOT_CONNECTED`},
		"filter":                         RepresentationGroup{Required, externalPluggableDatabaseDataSourceFilterRepresentation}}
	externalPluggableDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_external_pluggable_database.test_external_pluggable_database.id}`}},
	}

	externalPluggableDatabaseRepresentation = map[string]interface{}{
		"compartment_id":                 Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                   Representation{repType: Required, create: `myTestExternalPdb`},
		"external_container_database_id": Representation{repType: Required, create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"defined_tags":                   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	ExternalPluggableDatabaseResourceDependencies = generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Required, Create, externalContainerDatabaseRepresentation) +
		DefinedTagsDependencies
)

func TestDatabaseExternalPluggableDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalPluggableDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_external_pluggable_database.test_external_pluggable_database"
	datasourceName := "data.oci_database_external_pluggable_databases.test_external_pluggable_databases"
	singularDatasourceName := "data.oci_database_external_pluggable_database.test_external_pluggable_database"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseExternalPluggableDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", Required, Create, externalPluggableDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalPdb"),
					resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", Optional, Create, externalPluggableDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalPdb"),
					resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),

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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ExternalPluggableDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", Optional, Create,
						representationCopyWithNewProperties(externalPluggableDatabaseRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalPdb"),
					resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),

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
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", Optional, Update, externalPluggableDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalPdb"),
					resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),

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
					generateDataSourceFromRepresentationMap("oci_database_external_pluggable_databases", "test_external_pluggable_databases", Optional, Update, externalPluggableDatabaseDataSourceRepresentation) +
					compartmentIdVariableStr + ExternalPluggableDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", Optional, Update, externalPluggableDatabaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "myTestExternalPdb"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_container_database_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "NOT_CONNECTED"),

					resource.TestCheckResourceAttr(datasourceName, "external_pluggable_databases.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "external_pluggable_databases.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "external_pluggable_databases.0.database_management_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "external_pluggable_databases.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "external_pluggable_databases.0.display_name", "myTestExternalPdb"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_pluggable_databases.0.external_container_database_id"),

					resource.TestCheckResourceAttr(datasourceName, "external_pluggable_databases.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_pluggable_databases.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_pluggable_databases.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "external_pluggable_databases.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", Required, Create, externalPluggableDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ExternalPluggableDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "external_pluggable_database_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "database_management_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "myTestExternalPdb"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseResourceConfig,
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

func testAccCheckDatabaseExternalPluggableDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_external_pluggable_database" {
			noResourceFound = false
			request := oci_database.GetExternalPluggableDatabaseRequest{}

			tmp := rs.Primary.ID
			request.ExternalPluggableDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetExternalPluggableDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExternalPluggableDatabaseLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("DatabaseExternalPluggableDatabase") {
		resource.AddTestSweepers("DatabaseExternalPluggableDatabase", &resource.Sweeper{
			Name:         "DatabaseExternalPluggableDatabase",
			Dependencies: DependencyGraph["externalPluggableDatabase"],
			F:            sweepDatabaseExternalPluggableDatabaseResource,
		})
	}
}

func sweepDatabaseExternalPluggableDatabaseResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	externalPluggableDatabaseIds, err := getExternalPluggableDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, externalPluggableDatabaseId := range externalPluggableDatabaseIds {
		if ok := SweeperDefaultResourceId[externalPluggableDatabaseId]; !ok {
			deleteExternalPluggableDatabaseRequest := oci_database.DeleteExternalPluggableDatabaseRequest{}

			deleteExternalPluggableDatabaseRequest.ExternalPluggableDatabaseId = &externalPluggableDatabaseId

			deleteExternalPluggableDatabaseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExternalPluggableDatabase(context.Background(), deleteExternalPluggableDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalPluggableDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalPluggableDatabaseId, error)
				continue
			}
			waitTillCondition(testAccProvider, &externalPluggableDatabaseId, externalPluggableDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				externalPluggableDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getExternalPluggableDatabaseIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ExternalPluggableDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listExternalPluggableDatabasesRequest := oci_database.ListExternalPluggableDatabasesRequest{}
	listExternalPluggableDatabasesRequest.CompartmentId = &compartmentId
	listExternalPluggableDatabasesRequest.LifecycleState = oci_database.ExternalDatabaseBaseLifecycleStateAvailable
	listExternalPluggableDatabasesResponse, err := databaseClient.ListExternalPluggableDatabases(context.Background(), listExternalPluggableDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalPluggableDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalPluggableDatabase := range listExternalPluggableDatabasesResponse.Items {
		id := *externalPluggableDatabase.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ExternalPluggableDatabaseId", id)
	}
	return resourceIds, nil
}

func externalPluggableDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalPluggableDatabaseResponse, ok := response.Response.(oci_database.GetExternalPluggableDatabaseResponse); ok {
		return externalPluggableDatabaseResponse.LifecycleState != oci_database.ExternalPluggableDatabaseLifecycleStateTerminated
	}
	return false
}

func externalPluggableDatabaseSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetExternalPluggableDatabase(context.Background(), oci_database.GetExternalPluggableDatabaseRequest{
		ExternalPluggableDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
