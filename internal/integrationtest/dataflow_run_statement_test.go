// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataflowRunStatementRequiredOnlyResource = DataflowRunStatementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_run_statement", "test_run_statement", acctest.Required, acctest.Create, DataflowRunStatementRepresentation)

	DataflowRunStatementResourceConfig = DataflowRunStatementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_run_statement", "test_run_statement", acctest.Optional, acctest.Update, DataflowRunStatementRepresentation)

	DataflowRunStatementSingularDataSourceRepresentation = map[string]interface{}{
		"run_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_invoke_run.test_invoke_session_run.id}`},
		"statement_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_run_statement.test_run_statement.id}`},
	}

	DataflowRunStatementDataSourceRepresentation = map[string]interface{}{
		"run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_invoke_run.test_invoke_session_run.id}`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowRunStatementDataSourceFilterRepresentation}}
	DataflowRunStatementDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataflow_run_statement.test_run_statement.id}`}},
	}

	DataflowRunStatementRepresentation = map[string]interface{}{
		"code":   acctest.Representation{RepType: acctest.Required, Create: `${var.statement_code}`},
		"run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_invoke_run.test_invoke_session_run.id}`},
	}

	DataflowSessionApplicationRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `test_session_app`},
		"driver_shape":            acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"executor_shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"language":                acctest.Representation{RepType: acctest.Required, Create: `PYTHON`},
		"num_executors":           acctest.Representation{RepType: acctest.Required, Create: `1`},
		"spark_version":           acctest.Representation{RepType: acctest.Required, Create: `3.2.1`},
		"idle_timeout_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `30`},
		"max_duration_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `60`},
		"type":                    acctest.Representation{RepType: acctest.Optional, Create: `SESSION`},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForDataFlowResource},
	}

	DataflowInvokeSessionRunRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_application.test_session_application.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `test_session_run`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForDataFlowResource},
	}

	DataflowRunStatementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_session_application", acctest.Optional, acctest.Create, DataflowSessionApplicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_session_run", acctest.Optional, acctest.Create, DataflowInvokeSessionRunRepresentation)
)

// issue-routing-tag: dataflow/default
func TestDataflowRunStatementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowRunStatementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	statementCode := utils.GetEnvSettingWithBlankDefault("statement_code")
	statementCodeVariableStr := fmt.Sprintf("variable \"statement_code\" { default = \"%s\" }\n", statementCode)

	resourceName := "oci_dataflow_run_statement.test_run_statement"
	datasourceName := "data.oci_dataflow_run_statements.test_run_statements"
	singularDatasourceName := "data.oci_dataflow_run_statement.test_run_statement"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataflowRunStatementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_run_statement", "test_run_statement", acctest.Required, acctest.Create, DataflowRunStatementRepresentation), "dataflow", "runStatement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create Dependencies
		{
			Config: config + compartmentIdVariableStr + DataflowRunStatementResourceDependencies,
			Check: func(s *terraform.State) (err error) {
				// Workaround to wait for the Session Run to transition to IN_PROGRESS state.
				time.Sleep(15 * time.Minute)
				return nil
			},
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataflowRunStatementResourceDependencies + statementCodeVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_run_statement", "test_run_statement", acctest.Required, acctest.Create, DataflowRunStatementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "code", statementCode),
				resource.TestCheckResourceAttrSet(resourceName, "run_id"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_run_statements", "test_run_statements", acctest.Optional, acctest.Update, DataflowRunStatementDataSourceRepresentation) +
				compartmentIdVariableStr + statementCodeVariableStr + DataflowRunStatementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_run_statement", "test_run_statement", acctest.Optional, acctest.Update, DataflowRunStatementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "run_id"),
				resource.TestCheckResourceAttr(datasourceName, "statement_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_run_statement", "test_run_statement", acctest.Required, acctest.Create, DataflowRunStatementSingularDataSourceRepresentation) +
				compartmentIdVariableStr + statementCodeVariableStr + DataflowRunStatementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "statement_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "code"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "progress"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_completed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataflowRunStatementRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataflowRunStatementDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataFlowClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataflow_run_statement" {
			noResourceFound = false
			request := oci_dataflow.GetStatementRequest{}

			if value, ok := rs.Primary.Attributes["run_id"]; ok {
				request.RunId = &value
			}

			tmp := rs.Primary.ID
			request.StatementId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")

			_, err := client.GetStatement(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("DataflowRunStatement") {
		resource.AddTestSweepers("DataflowRunStatement", &resource.Sweeper{
			Name:         "DataflowRunStatement",
			Dependencies: acctest.DependencyGraph["runStatement"],
			F:            sweepDataflowRunStatementResource,
		})
	}
}

func sweepDataflowRunStatementResource(compartment string) error {
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()
	runStatementIds, err := getDataflowRunStatementIds(compartment)
	if err != nil {
		return err
	}
	for _, runStatementId := range runStatementIds {
		if ok := acctest.SweeperDefaultResourceId[runStatementId]; !ok {
			deleteStatementRequest := oci_dataflow.DeleteStatementRequest{}

			deleteStatementRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")
			_, error := dataFlowClient.DeleteStatement(context.Background(), deleteStatementRequest)
			if error != nil {
				fmt.Printf("Error deleting RunStatement %s %s, It is possible that the resource is already deleted. Please verify manually \n", runStatementId, error)
				continue
			}
		}
	}
	return nil
}

func getDataflowRunStatementIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RunStatementId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()

	listStatementsRequest := oci_dataflow.ListStatementsRequest{}

	runIds, error := getDataflowInvokeRunIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting runId required for RunStatement resource requests \n")
	}
	for _, runId := range runIds {
		listStatementsRequest.RunId = &runId

		listStatementsResponse, err := dataFlowClient.ListStatements(context.Background(), listStatementsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting RunStatement list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, runStatement := range listStatementsResponse.Items {
			id := strconv.FormatInt(*runStatement.Id, 10)
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RunStatementId", id)
		}

	}
	return resourceIds, nil
}
