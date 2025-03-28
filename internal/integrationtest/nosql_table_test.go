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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_nosql "github.com/oracle/oci-go-sdk/v65/nosql"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	NosqlTableRequiredOnlyResource = NosqlTableResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Required, acctest.Create, NosqlTableRepresentation)

	NosqlTableResourceConfig = NosqlTableResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Optional, acctest.Update, NosqlTableRepresentation)

	NosqlNosqlTableSingularDataSourceRepresentation = map[string]interface{}{
		"table_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_nosql_table.test_table.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}
	ddlStatement = "CREATE TABLE IF NOT EXISTS test_table(id INTEGER GENERATED ALWAYS AS IDENTITY, name STRING, age STRING, info JSON, guid STRING AS UUID, PRIMARY KEY(SHARD(id)))"

	NosqlNosqlTableDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `test_table`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: NosqlTableDataSourceFilterRepresentation}}
	NosqlTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_nosql_table.test_table.id}`}},
	}

	NosqlTableRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ddl_statement":       acctest.Representation{RepType: acctest.Required, Create: ddlStatement},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `test_table`},
		"table_limits":        acctest.RepresentationGroup{RepType: acctest.Required, Group: NosqlTableTableLimitsRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_auto_reclaimable": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreTableDefinedTags},
	}
	NosqlTableTableLimitsRepresentation = map[string]interface{}{
		"max_read_units":     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"max_storage_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"max_write_units":    acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"capacity_mode":      acctest.Representation{RepType: acctest.Optional, Create: `PROVISIONED`},
	}
	ignoreTableDefinedTags = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	NosqlTableResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: nosql/default
func TestNosqlTableResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNosqlTableResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_nosql_table.test_table"

	datasourceName := "data.oci_nosql_tables.test_tables"
	singularDatasourceName := "data.oci_nosql_table.test_table"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NosqlTableResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Optional, acctest.Create, NosqlTableRepresentation), "nosql", "table", t)

	acctest.ResourceTest(t, testAccCheckNosqlTableDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NosqlTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Required, acctest.Create, NosqlTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "ddl_statement", ddlStatement),
				resource.TestCheckResourceAttr(resourceName, "name", "test_table"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NosqlTableResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NosqlTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Optional, acctest.Create, NosqlTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "ddl_statement", ddlStatement),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_reclaimable", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.capacity_mode", "PROVISIONED"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_read_units", "10"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_storage_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_write_units", "10"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NosqlTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(NosqlTableRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "ddl_statement", ddlStatement),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_reclaimable", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.capacity_mode", "PROVISIONED"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_read_units", "10"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_storage_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_write_units", "10"),

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
			Config: config + compartmentIdVariableStr + NosqlTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Optional, acctest.Update, NosqlTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "ddl_statement", ddlStatement),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_reclaimable", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.capacity_mode", "PROVISIONED"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_read_units", "11"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_storage_in_gbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_write_units", "11"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_tables", "test_tables", acctest.Optional, acctest.Update, NosqlNosqlTableDataSourceRepresentation) +
				compartmentIdVariableStr + NosqlTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Optional, acctest.Update, NosqlTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "table_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "table_collection.0.id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Required, acctest.Create, NosqlNosqlTableSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NosqlTableResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "table_name_or_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "ddl_statement", ddlStatement),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_reclaimable", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_multi_region"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replicas.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schema_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.identity.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.identity.0.column_name", "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.identity.0.is_always", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.identity.0.is_null", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.columns.#", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.columns.0.name", "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.columns.0.is_as_uuid", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.columns.0.is_generated", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.columns.4.name", "guid"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.columns.4.is_as_uuid", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.0.columns.4.is_generated", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.capacity_mode", "PROVISIONED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.max_read_units", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.max_storage_in_gbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.max_write_units", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + NosqlTableRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNosqlTableDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NosqlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_nosql_table" {
			noResourceFound = false
			request := oci_nosql.GetTableRequest{}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			if value, ok := rs.Primary.Attributes["table_name_or_id"]; ok {
				request.TableNameOrId = &value
			} else if rs.Primary.ID != "" {
				tmp := rs.Primary.ID
				request.TableNameOrId = &tmp
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "nosql")

			response, err := client.GetTable(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_nosql.TableLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("NosqlTable") {
		resource.AddTestSweepers("NosqlTable", &resource.Sweeper{
			Name:         "NosqlTable",
			Dependencies: acctest.DependencyGraph["table"],
			F:            sweepNosqlTableResource,
		})
	}
}

func sweepNosqlTableResource(compartment string) error {
	nosqlClient := acctest.GetTestClients(&schema.ResourceData{}).NosqlClient()
	tableIds, err := getNosqlTableIds(compartment)
	if err != nil {
		return err
	}
	for _, tableId := range tableIds {
		if ok := acctest.SweeperDefaultResourceId[tableId]; !ok {
			deleteTableRequest := oci_nosql.DeleteTableRequest{}

			deleteTableRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "nosql")
			_, error := nosqlClient.DeleteTable(context.Background(), deleteTableRequest)
			if error != nil {
				fmt.Printf("Error deleting Table %s %s, It is possible that the resource is already deleted. Please verify manually \n", tableId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &tableId, NosqlTableSweepWaitCondition, time.Duration(3*time.Minute),
				NosqlTableSweepResponseFetchOperation, "nosql", true)
		}
	}
	return nil
}

func getNosqlTableIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TableId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	nosqlClient := acctest.GetTestClients(&schema.ResourceData{}).NosqlClient()

	listTablesRequest := oci_nosql.ListTablesRequest{}
	listTablesRequest.CompartmentId = &compartmentId
	listTablesRequest.LifecycleState = oci_nosql.ListTablesLifecycleStateActive
	listTablesResponse, err := nosqlClient.ListTables(context.Background(), listTablesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Table list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, table := range listTablesResponse.Items {
		id := *table.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TableId", id)
	}
	return resourceIds, nil
}

func NosqlTableSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if tableResponse, ok := response.Response.(oci_nosql.GetTableResponse); ok {
		return tableResponse.LifecycleState != oci_nosql.TableLifecycleStateDeleted
	}
	return false
}

func NosqlTableSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.NosqlClient().GetTable(context.Background(), oci_nosql.GetTableRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
