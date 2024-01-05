// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSqlCollectionRequiredOnlyResource = DataSafeSqlCollectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Required, acctest.Create, DataSafeSqlCollectionRepresentation)

	DataSafeSqlCollectionResourceConfig = DataSafeSqlCollectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Update, DataSafeSqlCollectionRepresentation)

	DataSafeSqlCollectionSingularDataSourceRepresentation = map[string]interface{}{
		"sql_collection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sql_collection.test_sql_collection.id}`},
	}

	DataSafeSqlCollectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                          acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"db_user_name":                          acctest.Representation{RepType: acctest.Optional, Create: `${var.db_user_name}`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `SampleSqlCollection`, Update: `displayName2`},
		"sql_collection_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sql_collection.test_sql_collection.id}`},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `COMPLETED`},
		"target_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSqlCollectionDataSourceFilterRepresentation}}
	DataSafeSqlCollectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_sql_collection.test_sql_collection.id}`}},
	}

	DataSafeSqlCollectionRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_user_name":   acctest.Representation{RepType: acctest.Required, Create: `${var.db_user_name}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Sample SqlCollection description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `SampleSqlCollection`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"sql_level":      acctest.Representation{RepType: acctest.Optional, Create: `ALL_SQL`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSqlCollectionSystemTagsChangesRep},
	}

	DataSafeSqlCollectionResourceDependencies = DefinedTagsDependencies

	ignoreSqlCollectionSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSqlCollectionStartRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_user_name":   acctest.Representation{RepType: acctest.Required, Create: `${var.db_user_name}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"start_trigger":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSqlCollectionSystemTagsChangesRep},
	}

	DataSafeSqlCollectionStopRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_user_name":   acctest.Representation{RepType: acctest.Required, Create: `${var.db_user_name}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"stop_trigger":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSqlCollectionSystemTagsChangesRep},
	}

	DataSafeSqlCollectionRefreshInsightsRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_user_name":                 acctest.Representation{RepType: acctest.Required, Create: `${var.db_user_name}`},
		"target_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"refresh_log_insights_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSqlCollectionSystemTagsChangesRep},
	}

	DataSafeSqlCollectionPurgeLogsRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_user_name":       acctest.Representation{RepType: acctest.Required, Create: `${var.db_user_name}`},
		"target_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"purge_logs_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSqlCollectionSystemTagsChangesRep},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeSqlCollectionResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the target ocid and dbUserName are hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSqlCollectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	dbUserName := utils.GetEnvSettingWithBlankDefault("data_safe_db_user_name")
	dbUserNameVariableStr := fmt.Sprintf("variable \"db_user_name\" { default = \"%s\" }\n", dbUserName)

	resourceName := "oci_data_safe_sql_collection.test_sql_collection"
	datasourceName := "data.oci_data_safe_sql_collections.test_sql_collections"
	singularDatasourceName := "data.oci_data_safe_sql_collection.test_sql_collection"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+targetIdVariableStr+dbUserNameVariableStr+DataSafeSqlCollectionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Create, DataSafeSqlCollectionRepresentation), "datasafe", "sqlCollection", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSqlCollectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Required, acctest.Create, DataSafeSqlCollectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_user_name"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// stop SQL Collection
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Update, DataSafeSqlCollectionStopRepresentation),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Create, DataSafeSqlCollectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_user_name"),
				resource.TestCheckResourceAttr(resourceName, "description", "Sample SqlCollection description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "SampleSqlCollection"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "sql_level", "ALL_SQL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
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
		// start SQL Collection
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Update, DataSafeSqlCollectionStartRepresentation),
		},
		// stop SQL Collection
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Update, DataSafeSqlCollectionStopRepresentation),
		},
		// refresh SQL Collection
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Update, DataSafeSqlCollectionRefreshInsightsRepresentation),
		},
		// purge SQL Collection logs
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Update, DataSafeSqlCollectionPurgeLogsRepresentation),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeSqlCollectionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "db_user_name"),
				resource.TestCheckResourceAttr(resourceName, "description", "Sample SqlCollection description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "SampleSqlCollection"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "sql_level", "ALL_SQL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Update, DataSafeSqlCollectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_user_name"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "sql_level", "ALL_SQL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_collections", "test_sql_collections", acctest.Optional, acctest.Update, DataSafeSqlCollectionDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Optional, acctest.Update, DataSafeSqlCollectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_user_name"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_collection_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "COMPLETED"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),

				resource.TestCheckResourceAttr(datasourceName, "sql_collection_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sql_collection_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_collection", "test_sql_collection", acctest.Required, acctest.Create, DataSafeSqlCollectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + dbUserNameVariableStr + DataSafeSqlCollectionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_collection_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sql_level", "ALL_SQL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_stopped"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeSqlCollectionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`purge_logs_trigger`, `refresh_log_insights_trigger`, `start_trigger`, `stop_trigger`, `generate_sql_firewall_policy_trigger`},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSqlCollectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_sql_collection" {
			noResourceFound = false
			request := oci_data_safe.GetSqlCollectionRequest{}

			tmp := rs.Primary.ID
			request.SqlCollectionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSqlCollection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SqlCollectionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeSqlCollection") {
		resource.AddTestSweepers("DataSafeSqlCollection", &resource.Sweeper{
			Name:         "DataSafeSqlCollection",
			Dependencies: acctest.DependencyGraph["sqlCollection"],
			F:            sweepDataSafeSqlCollectionResource,
		})
	}
}

func sweepDataSafeSqlCollectionResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	sqlCollectionIds, err := getDataSafeSqlCollectionIds(compartment)
	if err != nil {
		return err
	}
	for _, sqlCollectionId := range sqlCollectionIds {
		if ok := acctest.SweeperDefaultResourceId[sqlCollectionId]; !ok {
			deleteSqlCollectionRequest := oci_data_safe.DeleteSqlCollectionRequest{}

			deleteSqlCollectionRequest.SqlCollectionId = &sqlCollectionId

			deleteSqlCollectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSqlCollection(context.Background(), deleteSqlCollectionRequest)
			if error != nil {
				fmt.Printf("Error deleting SqlCollection %s %s, It is possible that the resource is already deleted. Please verify manually \n", sqlCollectionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sqlCollectionId, DataSafeSqlCollectionSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSqlCollectionSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSqlCollectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SqlCollectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSqlCollectionsRequest := oci_data_safe.ListSqlCollectionsRequest{}
	listSqlCollectionsRequest.CompartmentId = &compartmentId
	listSqlCollectionsRequest.LifecycleState = oci_data_safe.ListSqlCollectionsLifecycleStateCollecting
	listSqlCollectionsResponse, err := dataSafeClient.ListSqlCollections(context.Background(), listSqlCollectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SqlCollection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sqlCollection := range listSqlCollectionsResponse.Items {
		id := *sqlCollection.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SqlCollectionId", id)
	}
	return resourceIds, nil
}

func DataSafeSqlCollectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sqlCollectionResponse, ok := response.Response.(oci_data_safe.GetSqlCollectionResponse); ok {
		return sqlCollectionResponse.LifecycleState != oci_data_safe.SqlCollectionLifecycleStateDeleted
	}
	return false
}

func DataSafeSqlCollectionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSqlCollection(context.Background(), oci_data_safe.GetSqlCollectionRequest{
		SqlCollectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
