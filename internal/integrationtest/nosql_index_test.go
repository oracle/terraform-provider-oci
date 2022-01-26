// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_nosql "github.com/oracle/oci-go-sdk/v56/nosql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	IndexRequiredOnlyResource = IndexResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_nosql_index", "test_index", acctest.Required, acctest.Create, indexRepresentation)

	IndexResourceConfig = IndexResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_nosql_index", "test_index", acctest.Optional, acctest.Update, indexRepresentation)

	indexSingularDataSourceRepresentation = map[string]interface{}{
		"index_name":       acctest.Representation{RepType: acctest.Required, Create: `${oci_nosql_index.test_index.id}`},
		"table_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_nosql_table.test_table.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	indexDataSourceRepresentation = map[string]interface{}{
		"table_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_nosql_table.test_table.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":             acctest.Representation{RepType: acctest.Optional, Create: `test_index`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: indexDataSourceFilterRepresentation}}
	indexDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_nosql_index.test_index.name}`}},
	}

	indexRepresentation = map[string]interface{}{
		"keys":             acctest.RepresentationGroup{RepType: acctest.Required, Group: indexKeysRepresentation},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `test_index`},
		"table_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_nosql_table.test_table.id}`},
	}
	indexKeysRepresentation = map[string]interface{}{
		"column_name": acctest.Representation{RepType: acctest.Required, Create: `name`},
	}

	indexOptionalRepresentation = map[string]interface{}{
		"keys":             acctest.RepresentationGroup{RepType: acctest.Required, Group: indexKeyWithJsonRepresentation},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `test_index`},
		"table_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_nosql_table.test_table.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"is_if_not_exists": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	indexKeyWithJsonRepresentation = map[string]interface{}{
		"column_name":     acctest.Representation{RepType: acctest.Required, Create: `info`},
		"json_field_type": acctest.Representation{RepType: acctest.Optional, Create: `STRING`},
		"json_path":       acctest.Representation{RepType: acctest.Optional, Create: `info`},
	}

	IndexResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Required, acctest.Create, tableRepresentation)
)

// issue-routing-tag: nosql/default
func TestNosqlIndexResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNosqlIndexResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_nosql_index.test_index"

	datasourceName := "data.oci_nosql_indexes.test_indexes"
	singularDatasourceName := "data.oci_nosql_index.test_index"

	var compositeId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IndexResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_nosql_index", "test_index", acctest.Optional, acctest.Create, indexRepresentation), "nosql", "index", t)

	acctest.ResourceTest(t, testAccCheckNosqlIndexDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IndexResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_nosql_index", "test_index", acctest.Required, acctest.Create, indexRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "keys.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "keys.0.column_name", "name"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_index"),
				resource.TestCheckResourceAttrSet(resourceName, "table_name_or_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IndexResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IndexResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_nosql_index", "test_index", acctest.Optional, acctest.Create, indexOptionalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_if_not_exists", "false"),
				resource.TestCheckResourceAttr(resourceName, "keys.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "keys.0.column_name", "info"),
				resource.TestCheckResourceAttr(resourceName, "keys.0.json_field_type", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "keys.0.json_path", "info"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_index"),
				resource.TestCheckResourceAttrSet(resourceName, "table_name_or_id"),

				func(s *terraform.State) (err error) {
					indexName, err := acctest.FromInstanceState(s, resourceName, "id")
					tableName, _ := acctest.FromInstanceState(s, resourceName, "table_name_or_id")
					compositeId = "tables/" + tableName + "/indexes/" + indexName
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_indexes", "test_indexes", acctest.Optional, acctest.Update, indexDataSourceRepresentation) +
				compartmentIdVariableStr + IndexResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_nosql_index", "test_index", acctest.Optional, acctest.Update, indexRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "test_index"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "table_name_or_id"),

				resource.TestCheckResourceAttr(datasourceName, "index_collection.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_index", "test_index", acctest.Required, acctest.Create, indexSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IndexResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "index_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "table_name_or_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "keys.0.column_name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "test_index"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "table_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "table_name"),
			),
		},
	})
}

func testAccCheckNosqlIndexDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NosqlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_nosql_index" {
			noResourceFound = false
			request := oci_nosql.GetIndexRequest{}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.IndexName = &value
			}

			if value, ok := rs.Primary.Attributes["table_name_or_id"]; ok {
				request.TableNameOrId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "nosql")

			response, err := client.GetIndex(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_nosql.IndexLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("NosqlIndex") {
		resource.AddTestSweepers("NosqlIndex", &resource.Sweeper{
			Name:         "NosqlIndex",
			Dependencies: acctest.DependencyGraph["index"],
			F:            sweepNosqlIndexResource,
		})
	}
}

func sweepNosqlIndexResource(compartment string) error {
	nosqlClient := acctest.GetTestClients(&schema.ResourceData{}).NosqlClient()
	indexIds, err := getIndexIds(compartment)
	if err != nil {
		return err
	}
	for _, indexId := range indexIds {
		if ok := acctest.SweeperDefaultResourceId[indexId]; !ok {
			deleteIndexRequest := oci_nosql.DeleteIndexRequest{}

			deleteIndexRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "nosql")
			_, error := nosqlClient.DeleteIndex(context.Background(), deleteIndexRequest)
			if error != nil {
				fmt.Printf("Error deleting Index %s %s, It is possible that the resource is already deleted. Please verify manually \n", indexId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &indexId, indexSweepWaitCondition, time.Duration(3*time.Minute),
				indexSweepResponseFetchOperation, "nosql", true)
		}
	}
	return nil
}

func getIndexIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IndexId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	nosqlClient := acctest.GetTestClients(&schema.ResourceData{}).NosqlClient()

	listIndexesRequest := oci_nosql.ListIndexesRequest{}
	listIndexesRequest.CompartmentId = &compartmentId

	tableNameOrIds, error := getTableIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting tableNameOrId required for Index resource requests \n")
	}
	for _, tableNameOrId := range tableNameOrIds {
		listIndexesRequest.TableNameOrId = &tableNameOrId

		listIndexesRequest.LifecycleState = oci_nosql.ListIndexesLifecycleStateActive
		listIndexesResponse, err := nosqlClient.ListIndexes(context.Background(), listIndexesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Index list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, index := range listIndexesResponse.Items {
			id := *index.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IndexId", id)
		}

	}
	return resourceIds, nil
}

func indexSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if indexResponse, ok := response.Response.(oci_nosql.GetIndexResponse); ok {
		return indexResponse.LifecycleState != oci_nosql.IndexLifecycleStateDeleted
	}
	return false
}

func indexSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.NosqlClient().GetIndex(context.Background(), oci_nosql.GetIndexRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
