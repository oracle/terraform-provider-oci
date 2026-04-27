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
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsDatabaseToolsSqlReportRequiredOnlyResource = DatabaseToolsDatabaseToolsSqlReportResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_sql_report", "test_database_tools_sql_report", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsSqlReportRepresentation)

	DatabaseToolsDatabaseToolsSqlReportResourceConfig = DatabaseToolsDatabaseToolsSqlReportResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_sql_report", "test_database_tools_sql_report", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsSqlReportRepresentation)

	DatabaseToolsDatabaseToolsSqlReportSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_sql_report_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_sql_report.test_database_tools_sql_report.id}`},
	}

	DatabaseToolsDatabaseToolsSqlReportDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `SQL Report1`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_DATABASE`}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsSqlReportDataSourceFilterRepresentation},
	}
	DatabaseToolsDatabaseToolsSqlReportDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_sql_report.test_database_tools_sql_report.id}`}},
	}

	DatabaseToolsDatabaseToolsSqlReportRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `SQL Report1`, Update: `displayName2`},
		"source":         acctest.Representation{RepType: acctest.Required, Create: `SELECT * FROM SYS.DBA_HIST_SYSTEM_EVENT WHERE SNAP_ID IN(SELECT SNAP_ID FROM SYS.DBA_HIST_SNAPSHOT WHERE BEGIN_INTERVAL_TIME>SYSDATE-1)`, Update: `SELECT * FROM SYS.DBA_HIST_ACTIVE_SESS_HISTORY WHERE SNAP_ID IN(SELECT SNAP_ID FROM SYS.DBA_HIST_SNAPSHOT WHERE BEGIN_INTERVAL_TIME>SYSDATE-1)`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`},
		"columns":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsSqlReportColumnsRepresentation},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Simple SQL query`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"instructions":   acctest.Representation{RepType: acctest.Optional, Create: `instructions`, Update: `instructions2`},
		"purpose":        acctest.Representation{RepType: acctest.Optional, Create: `purpose`, Update: `purpose2`},
		"variables":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsSqlReportVariablesRepresentation},
	}
	DatabaseToolsDatabaseToolsSqlReportColumnsRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Required, Create: `Simple SQL query`, Update: `description2`},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"type":        acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`, Update: `type2`},
	}
	DatabaseToolsDatabaseToolsSqlReportVariablesRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Required, Create: `Simple SQL query`, Update: `description2`},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"type":        acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`, Update: `type2`},
	}

	DatabaseToolsDatabaseToolsSqlReportResourceDependencies = ""
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsSqlReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsSqlReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_sql_report.test_database_tools_sql_report"

	var resId, resId2 string
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseToolsDatabaseToolsSqlReportResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_sql_report", "test_database_tools_sql_report", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsSqlReportRepresentation), "databasetools", "databaseToolsSqlReport", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsSqlReportDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseToolsDatabaseToolsSqlReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_sql_report", "test_database_tools_sql_report", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsSqlReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "SQL Report1"),
				resource.TestCheckResourceAttr(resourceName, "source", "SELECT * FROM SYS.DBA_HIST_SYSTEM_EVENT WHERE SNAP_ID IN(SELECT SNAP_ID FROM SYS.DBA_HIST_SNAPSHOT WHERE BEGIN_INTERVAL_TIME>SYSDATE-1)"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseToolsDatabaseToolsSqlReportResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseToolsDatabaseToolsSqlReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_sql_report", "test_database_tools_sql_report", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsSqlReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "columns.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "columns.0.description", "Simple SQL query"),
				resource.TestCheckResourceAttr(resourceName, "columns.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "columns.0.type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Simple SQL query"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "SQL Report1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instructions", "instructions"),
				resource.TestCheckResourceAttr(resourceName, "purpose", "purpose"),
				resource.TestCheckResourceAttr(resourceName, "source", "SELECT * FROM SYS.DBA_HIST_SYSTEM_EVENT WHERE SNAP_ID IN(SELECT SNAP_ID FROM SYS.DBA_HIST_SNAPSHOT WHERE BEGIN_INTERVAL_TIME>SYSDATE-1)"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.description", "Simple SQL query"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.type", "ORACLE_DATABASE"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseToolsDatabaseToolsSqlReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_sql_report", "test_database_tools_sql_report", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsSqlReportRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "columns.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "columns.0.description", "Simple SQL query"),
				resource.TestCheckResourceAttr(resourceName, "columns.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "columns.0.type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Simple SQL query"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "SQL Report1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instructions", "instructions"),
				resource.TestCheckResourceAttr(resourceName, "purpose", "purpose"),
				resource.TestCheckResourceAttr(resourceName, "source", "SELECT * FROM SYS.DBA_HIST_SYSTEM_EVENT WHERE SNAP_ID IN(SELECT SNAP_ID FROM SYS.DBA_HIST_SNAPSHOT WHERE BEGIN_INTERVAL_TIME>SYSDATE-1)"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.description", "Simple SQL query"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.type", "ORACLE_DATABASE"),

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
			Config: config + compartmentIdVariableStr + DatabaseToolsDatabaseToolsSqlReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_sql_report", "test_database_tools_sql_report", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsSqlReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "columns.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "columns.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "columns.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "columns.0.type", "type2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instructions", "instructions2"),
				resource.TestCheckResourceAttr(resourceName, "purpose", "purpose2"),
				resource.TestCheckResourceAttr(resourceName, "source", "SELECT * FROM SYS.DBA_HIST_ACTIVE_SESS_HISTORY WHERE SNAP_ID IN(SELECT SNAP_ID FROM SYS.DBA_HIST_SNAPSHOT WHERE BEGIN_INTERVAL_TIME>SYSDATE-1)"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.type", "type2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsSqlReportDatasource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsSqlReportDatasource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_tools_database_tools_sql_report.test_database_tools_sql_report"
	datasourceName := "data.oci_database_tools_database_tools_sql_reports.test_database_tools_sql_reports"
	singularDatasourceName := "data.oci_database_tools_database_tools_sql_report.test_database_tools_sql_report"

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsSqlReportDestroy, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_sql_reports", "test_database_tools_sql_reports", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsSqlReportDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseToolsDatabaseToolsSqlReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_sql_report", "test_database_tools_sql_report", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsSqlReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "database_tools_sql_report_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_sql_report_collection.0.items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_sql_report", "test_database_tools_sql_report", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsSqlReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseToolsDatabaseToolsSqlReportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_sql_report_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "columns.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "columns.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "columns.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "columns.0.type", "type2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instructions", "instructions2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "purpose", "purpose2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source", "SELECT * FROM SYS.DBA_HIST_ACTIVE_SESS_HISTORY WHERE SNAP_ID IN(SELECT SNAP_ID FROM SYS.DBA_HIST_SNAPSHOT WHERE BEGIN_INTERVAL_TIME>SYSDATE-1)"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.type", "type2"),
			),
		},

		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + DatabaseToolsDatabaseToolsSqlReportRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseToolsDatabaseToolsSqlReportDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_database_tools_sql_report" {
			noResourceFound = false
			request := oci_database_tools.GetDatabaseToolsSqlReportRequest{}

			tmp := rs.Primary.ID
			request.DatabaseToolsSqlReportId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")

			response, err := client.GetDatabaseToolsSqlReport(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_tools.DatabaseToolsSqlReportLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseToolsDatabaseToolsSqlReport") {
		resource.AddTestSweepers("DatabaseToolsDatabaseToolsSqlReport", &resource.Sweeper{
			Name:         "DatabaseToolsDatabaseToolsSqlReport",
			Dependencies: acctest.DependencyGraph["databaseToolsSqlReport"],
			F:            sweepDatabaseToolsDatabaseToolsSqlReportResource,
		})
	}
}

func sweepDatabaseToolsDatabaseToolsSqlReportResource(compartment string) error {
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()
	databaseToolsSqlReportIds, err := getDatabaseToolsDatabaseToolsSqlReportIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsSqlReportId := range databaseToolsSqlReportIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsSqlReportId]; !ok {
			deleteDatabaseToolsSqlReportRequest := oci_database_tools.DeleteDatabaseToolsSqlReportRequest{}

			deleteDatabaseToolsSqlReportRequest.DatabaseToolsSqlReportId = &databaseToolsSqlReportId

			deleteDatabaseToolsSqlReportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")
			_, error := databaseToolsClient.DeleteDatabaseToolsSqlReport(context.Background(), deleteDatabaseToolsSqlReportRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsSqlReport %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsSqlReportId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseToolsSqlReportId, DatabaseToolsDatabaseToolsSqlReportSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseToolsDatabaseToolsSqlReportSweepResponseFetchOperation, "database_tools", true)
		}
	}
	return nil
}

func getDatabaseToolsDatabaseToolsSqlReportIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsSqlReportId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()

	listDatabaseToolsSqlReportsRequest := oci_database_tools.ListDatabaseToolsSqlReportsRequest{}
	listDatabaseToolsSqlReportsRequest.CompartmentId = &compartmentId
	listDatabaseToolsSqlReportsRequest.LifecycleState = oci_database_tools.ListDatabaseToolsSqlReportsLifecycleStateActive
	listDatabaseToolsSqlReportsResponse, err := databaseToolsClient.ListDatabaseToolsSqlReports(context.Background(), listDatabaseToolsSqlReportsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseToolsSqlReport list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseToolsSqlReport := range listDatabaseToolsSqlReportsResponse.Items {
		id := *databaseToolsSqlReport.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsSqlReportId", id)
	}
	return resourceIds, nil
}

func DatabaseToolsDatabaseToolsSqlReportSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseToolsSqlReportResponse, ok := response.Response.(oci_database_tools.GetDatabaseToolsSqlReportResponse); ok {
		return databaseToolsSqlReportResponse.GetLifecycleState() != oci_database_tools.DatabaseToolsSqlReportLifecycleStateDeleted
	}
	return false
}

func DatabaseToolsDatabaseToolsSqlReportSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseToolsClient().GetDatabaseToolsSqlReport(context.Background(), oci_database_tools.GetDatabaseToolsSqlReportRequest{
		DatabaseToolsSqlReportId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
