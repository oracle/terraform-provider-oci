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
	ExternalMySqlDatabaseInsightRequiredOnlyResource = ExternalMySqlDatabaseInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, externalMySqlDatabaseInsightRepresentation)

	ExternalMySqlDatabaseInsightResourceConfig = ExternalMySqlDatabaseInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update, externalMySqlDatabaseInsightRepresentation)

	ExternalMySqlDatabaseInsightSingularDataSourceRepresentation = map[string]interface{}{
		"database_insight_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_database_insight.test_database_insight.id}`},
	}

	externalMySqlDatabaseInsightDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"database_type":             acctest.Representation{RepType: acctest.Optional, Create: []string{`EXTERNAL-MYSQL`}},
		"fields":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`databaseName`, `databaseType`, `compartmentId`, `databaseDisplayName`, `freeformTags`, `definedTags`}},
		"id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_database_insight.test_database_insight.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`ENABLED`}, Update: []string{`DISABLED`}},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ExternalMySqlDatabaseInsightDataSourceFilterRepresentation},
	}

	ExternalMySqlDatabaseInsightDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_database_insight.test_database_insight.id}`}},
	}

	externalMySqlDatabaseInsightRepresentation = map[string]interface{}{
		"database_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.dbmgmt_external_mysql_database_id}`},
		"database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_connector_id}`},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"entity_source":         acctest.Representation{RepType: acctest.Required, Create: `EXTERNAL_MYSQL_DATABASE_SYSTEM`, Update: `EXTERNAL_MYSQL_DATABASE_SYSTEM`},
		"status":                acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesExternalMySqlDatabaseInsightRepresentation},
	}

	ignoreChangesExternalMySqlDatabaseInsightRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	ExternalMySqlDatabaseInsightResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiExternalMySqlDatabaseInsightResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiExternalMySqlDatabaseInsightResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	dbmgmtExternalMySqlDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_mysql_database_id")
	dbmgmtExternalMySqlDatabaseIdVariableStr := fmt.Sprintf("variable \"dbmgmt_external_mysql_database_id\" { default = \"%s\" }\n", dbmgmtExternalMySqlDatabaseId)

	databaseConnectorId := utils.GetEnvSettingWithBlankDefault("database_connector_id")
	databaseConnectorIdVariableStr := fmt.Sprintf("variable \"database_connector_id\" { default = \"%s\" }\n", databaseConnectorId)

	databaseConnectorIdUpdate := utils.GetEnvSettingWithBlankDefault("database_connector_id_for_update")
	databaseConnectorIdUpdateVariableStr := fmt.Sprintf("variable \"database_connector_id_for_update\" { default = \"%s\" }\n", databaseConnectorIdUpdate)

	resourceName := "oci_opsi_database_insight.test_database_insight"
	datasourceName := "data.oci_opsi_database_insights.test_database_insights"
	singularDatasourceName := "data.oci_opsi_database_insight.test_database_insight"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbmgmtExternalMySqlDatabaseIdVariableStr+databaseConnectorIdVariableStr+ExternalMySqlDatabaseInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Create, externalMySqlDatabaseInsightRepresentation), "opsi", "databaseInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiExternalMySqlDatabaseInsightDestroy, []resource.TestStep{
		// verify create with optionals - Step 0
		{
			Config: config + compartmentIdVariableStr + dbmgmtExternalMySqlDatabaseIdVariableStr + databaseConnectorIdVariableStr + ExternalMySqlDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Create, externalMySqlDatabaseInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "EXTERNAL_MYSQL_DATABASE_SYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
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
		// verify update to the compartment and database connector id (the compartment and database connector id will be switched back in the next step) - Step 1
		{
			Config: config + compartmentIdVariableStr + dbmgmtExternalMySqlDatabaseIdVariableStr + databaseConnectorIdUpdateVariableStr + compartmentIdUVariableStr + ExternalMySqlDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(externalMySqlDatabaseInsightRepresentation, map[string]interface{}{
						"database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_connector_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "EXTERNAL_MYSQL_DATABASE_SYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
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

		// verify updates to updatable parameters - Step 2 (Update causes status to go to disabled)
		{
			Config: config + compartmentIdVariableStr + dbmgmtExternalMySqlDatabaseIdVariableStr + databaseConnectorIdVariableStr + ExternalMySqlDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update, externalMySqlDatabaseInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "database_name"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "EXTERNAL_MYSQL_DATABASE_SYSTEM"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
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
		// verify datasource - Step 3
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_database_insights", "test_database_insights", acctest.Optional, acctest.Update, externalMySqlDatabaseInsightDataSourceRepresentation) +
				compartmentIdVariableStr + dbmgmtExternalMySqlDatabaseIdVariableStr + databaseConnectorIdVariableStr + ExternalMySqlDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update, externalMySqlDatabaseInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "database_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fields.#", "6"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "status.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_insights_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_insights_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource - Step 4
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, ExternalMySqlDatabaseInsightSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbmgmtExternalMySqlDatabaseIdVariableStr + databaseConnectorIdVariableStr + ExternalMySqlDatabaseInsightResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_source", "EXTERNAL_MYSQL_DATABASE_SYSTEM"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests - Step 5
		{
			Config: config + compartmentIdVariableStr + dbmgmtExternalMySqlDatabaseIdVariableStr + databaseConnectorIdVariableStr + ExternalMySqlDatabaseInsightResourceConfig,
		},
		// verify enable - Step 6
		{
			Config: config + compartmentIdVariableStr + dbmgmtExternalMySqlDatabaseIdVariableStr + databaseConnectorIdVariableStr + ExternalMySqlDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(externalMySqlDatabaseInsightRepresentation, map[string]interface{}{
						"status": acctest.Representation{RepType: acctest.Required, Update: `ENABLED`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//// verify resource import - Step 7
		{
			Config:            config + ExternalMySqlDatabaseInsightRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"database_connector_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOpsiExternalMySqlDatabaseInsightDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_database_insight" {
			noResourceFound = false
			request := oci_opsi.GetDatabaseInsightRequest{}

			tmp := rs.Primary.ID
			request.DatabaseInsightId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetDatabaseInsight(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.LifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("OpsiExternalMySqlDatabaseInsight") {
		resource.AddTestSweepers("OpsiExternalMySqlDatabaseInsight", &resource.Sweeper{
			Name:         "OpsiExternalMySqlDatabaseInsight",
			Dependencies: acctest.DependencyGraph["databaseInsight"],
			F:            sweepOpsiExternalMySqlDatabaseInsightResource,
		})
	}
}

func sweepOpsiExternalMySqlDatabaseInsightResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	databaseInsightIds, err := getExternalMySqlDatabaseInsightIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseInsightId := range databaseInsightIds {
		if ok := acctest.SweeperDefaultResourceId[databaseInsightId]; !ok {
			deleteDatabaseInsightRequest := oci_opsi.DeleteDatabaseInsightRequest{}

			deleteDatabaseInsightRequest.DatabaseInsightId = &databaseInsightId

			deleteDatabaseInsightRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteDatabaseInsight(context.Background(), deleteDatabaseInsightRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseInsight %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseInsightId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseInsightId, OpsiExternalMySqlDatabaseInsightSweepWaitCondition, time.Duration(3*time.Minute),
				OpsiExternalMySqlDatabaseInsightSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getExternalMySqlDatabaseInsightIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseInsightId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listDatabaseInsightsRequest := oci_opsi.ListDatabaseInsightsRequest{}
	listDatabaseInsightsRequest.CompartmentId = &compartmentId
	listDatabaseInsightsRequest.LifecycleState = []oci_opsi.LifecycleStateEnum{oci_opsi.LifecycleStateActive}
	listDatabaseInsightsResponse, err := operationsInsightsClient.ListDatabaseInsights(context.Background(), listDatabaseInsightsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseInsight list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseInsight := range listDatabaseInsightsResponse.Items {
		id := *databaseInsight.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseInsightId", id)
	}
	return resourceIds, nil
}

func OpsiExternalMySqlDatabaseInsightSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseInsightResponse, ok := response.Response.(oci_opsi.GetDatabaseInsightResponse); ok {
		return databaseInsightResponse.GetLifecycleState() != oci_opsi.LifecycleStateDeleted
	}
	return false
}

func OpsiExternalMySqlDatabaseInsightSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetDatabaseInsight(context.Background(), oci_opsi.GetDatabaseInsightRequest{
		DatabaseInsightId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
