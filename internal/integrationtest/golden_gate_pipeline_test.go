// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GoldenGatePipelineRequiredOnlyResource = GoldenGatePipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Required, acctest.Create, GoldenGatePipelineRepresentation)

	GoldenGatePipelineResourceConfig = GoldenGatePipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Optional, acctest.Update, GoldenGatePipelineRepresentation)

	GoldenGatePipelineSingularDataSourceRepresentation = map[string]interface{}{
		"pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_pipeline.test_pipeline.id}`},
	}

	GoldenGatePipelineDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"lifecycle_sub_state": acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: GoldenGatePipelineDataSourceFilterRepresentation}}
	GoldenGatePipelineDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_golden_gate_pipeline.test_pipeline.id}`}},
	}

	GoldenGatePipelineRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"license_model":             acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"recipe_type":               acctest.Representation{RepType: acctest.Required, Create: `ZERO_ETL`},
		"source_connection_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: GoldenGatePipelineSourceConnectionDetailsRepresentation},
		"target_connection_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: GoldenGatePipelineTargetConnectionDetailsRepresentation},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		//"locks":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: GoldenGatePipelineLocksRepresentation},
		"process_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GoldenGatePipelineProcessOptionsRepresentation},
	}
	GoldenGatePipelineSourceConnectionDetailsRepresentation = map[string]interface{}{
		//"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_connection.test_connection.id}`},
		"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_id}`},
	}
	GoldenGatePipelineTargetConnectionDetailsRepresentation = map[string]interface{}{
		//"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_connection.test_connection.id}`},
		"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.target_connection_id}`},
	}
	//GoldenGatePipelineLocksRepresentation = map[string]interface{}{
	//	"type":    acctest.Representation{RepType: acctest.Required, Create: `FULL`},
	//	"message": acctest.Representation{RepType: acctest.Optional, Create: `message`},
	//}
	GoldenGatePipelineProcessOptionsRepresentation = map[string]interface{}{
		"initial_data_load":           acctest.RepresentationGroup{RepType: acctest.Required, Group: GoldenGatePipelineProcessOptionsInitialDataLoadRepresentation},
		"replicate_schema_change":     acctest.RepresentationGroup{RepType: acctest.Required, Group: GoldenGatePipelineProcessOptionsReplicateSchemaChangeRepresentation},
		"should_restart_on_failure":   acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `DISABLED`},
		"start_using_default_mapping": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
	}
	GoldenGatePipelineProcessOptionsInitialDataLoadRepresentation = map[string]interface{}{
		"is_initial_load":          acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `DISABLED`},
		"action_on_existing_table": acctest.Representation{RepType: acctest.Optional, Create: `TRUNCATE`, Update: `REPLACE`},
	}
	GoldenGatePipelineProcessOptionsReplicateSchemaChangeRepresentation = map[string]interface{}{
		"can_replicate_schema_change": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `DISABLED`},
		"action_on_ddl_error":         acctest.Representation{RepType: acctest.Optional, Create: `TERMINATE`, Update: `DISCARD`},
		"action_on_dml_error":         acctest.Representation{RepType: acctest.Optional, Create: `TERMINATE`, Update: `DISCARD`},
	}

	GoldenGatePipelineResourceDependencies = ""
)

// issue-routing-tag: golden_gate/default
func TestGoldenGatePipelineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGatePipelineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() +
		makeVariableStr("source_connection_id", t) +
		makeVariableStr("target_connection_id", t)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_golden_gate_pipeline.test_pipeline"
	datasourceName := "data.oci_golden_gate_pipelines.test_pipelines"
	singularDatasourceName := "data.oci_golden_gate_pipeline.test_pipeline"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GoldenGatePipelineResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Optional, acctest.Create, GoldenGatePipelineRepresentation), "goldengate", "pipeline", t)

	acctest.ResourceTest(t, testAccCheckGoldenGatePipelineDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GoldenGatePipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Required, acctest.Create, GoldenGatePipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "recipe_type", "ZERO_ETL"),
				resource.TestCheckResourceAttr(resourceName, "source_connection_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_connection_details.0.connection_id"),
				resource.TestCheckResourceAttr(resourceName, "target_connection_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_connection_details.0.connection_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GoldenGatePipelineResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GoldenGatePipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Optional, acctest.Create, GoldenGatePipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpu_core_count"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_auto_scaling_enabled"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.initial_data_load.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.initial_data_load.0.action_on_existing_table", "TRUNCATE"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.initial_data_load.0.is_initial_load", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.0.action_on_ddl_error", "TERMINATE"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.0.action_on_dml_error", "TERMINATE"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.0.can_replicate_schema_change", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.should_restart_on_failure", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.start_using_default_mapping", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "recipe_type", "ZERO_ETL"),
				resource.TestCheckResourceAttr(resourceName, "source_connection_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_connection_details.0.connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_connection_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_connection_details.0.connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GoldenGatePipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GoldenGatePipelineRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "cpu_core_count"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_auto_scaling_enabled"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.initial_data_load.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.initial_data_load.0.action_on_existing_table", "TRUNCATE"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.initial_data_load.0.is_initial_load", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.0.action_on_ddl_error", "TERMINATE"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.0.action_on_dml_error", "TERMINATE"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.0.can_replicate_schema_change", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.should_restart_on_failure", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.start_using_default_mapping", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "recipe_type", "ZERO_ETL"),
				resource.TestCheckResourceAttr(resourceName, "source_connection_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_connection_details.0.connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_connection_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_connection_details.0.connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + GoldenGatePipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Optional, acctest.Update, GoldenGatePipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpu_core_count"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_auto_scaling_enabled"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.initial_data_load.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.initial_data_load.0.action_on_existing_table", "REPLACE"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.initial_data_load.0.is_initial_load", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.0.action_on_ddl_error", "DISCARD"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.0.action_on_dml_error", "DISCARD"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.replicate_schema_change.0.can_replicate_schema_change", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.should_restart_on_failure", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "process_options.0.start_using_default_mapping", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "recipe_type", "ZERO_ETL"),
				resource.TestCheckResourceAttr(resourceName, "source_connection_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_connection_details.0.connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_connection_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_connection_details.0.connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_pipelines", "test_pipelines", acctest.Optional, acctest.Update, GoldenGatePipelineDataSourceRepresentation) +
				compartmentIdVariableStr + GoldenGatePipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Optional, acctest.Update, GoldenGatePipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "lifecycle_sub_state", "STOPPED"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "pipeline_collection.#", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "pipeline_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Required, acctest.Create, GoldenGatePipelineSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GoldenGatePipelineResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pipeline_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_core_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_auto_scaling_enabled"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_sub_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mapping_rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pipeline_diagnostic_data.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "process_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "process_options.0.initial_data_load.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "process_options.0.initial_data_load.0.action_on_existing_table", "REPLACE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "process_options.0.initial_data_load.0.is_initial_load", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "process_options.0.replicate_schema_change.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "process_options.0.replicate_schema_change.0.action_on_ddl_error", "DISCARD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "process_options.0.replicate_schema_change.0.action_on_dml_error", "DISCARD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "process_options.0.replicate_schema_change.0.can_replicate_schema_change", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "process_options.0.should_restart_on_failure", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "process_options.0.start_using_default_mapping", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recipe_type", "ZERO_ETL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_connection_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_connection_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_recorded"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GoldenGatePipelineRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGoldenGatePipelineDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GoldenGateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_golden_gate_pipeline" {
			noResourceFound = false
			request := oci_golden_gate.GetPipelineRequest{}

			tmp := rs.Primary.ID
			request.PipelineId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")

			response, err := client.GetPipeline(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_golden_gate.PipelineLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GoldenGatePipeline") {
		resource.AddTestSweepers("GoldenGatePipeline", &resource.Sweeper{
			Name:         "GoldenGatePipeline",
			Dependencies: acctest.DependencyGraph["pipeline"],
			F:            sweepGoldenGatePipelineResource,
		})
	}
}

func sweepGoldenGatePipelineResource(compartment string) error {
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()
	pipelineIds, err := getGoldenGatePipelineIds(compartment)
	if err != nil {
		return err
	}
	for _, pipelineId := range pipelineIds {
		if ok := acctest.SweeperDefaultResourceId[pipelineId]; !ok {
			deletePipelineRequest := oci_golden_gate.DeletePipelineRequest{}

			deletePipelineRequest.PipelineId = &pipelineId

			deletePipelineRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")
			_, error := goldenGateClient.DeletePipeline(context.Background(), deletePipelineRequest)
			if error != nil {
				fmt.Printf("Error deleting Pipeline %s %s, It is possible that the resource is already deleted. Please verify manually \n", pipelineId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &pipelineId, GoldenGatePipelineSweepWaitCondition, time.Duration(3*time.Minute),
				GoldenGatePipelineSweepResponseFetchOperation, "golden_gate", true)
		}
	}
	return nil
}

func getGoldenGatePipelineIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PipelineId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()

	listPipelinesRequest := oci_golden_gate.ListPipelinesRequest{}
	listPipelinesRequest.CompartmentId = &compartmentId
	listPipelinesRequest.LifecycleState = oci_golden_gate.PipelineLifecycleStateActive
	listPipelinesResponse, err := goldenGateClient.ListPipelines(context.Background(), listPipelinesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Pipeline list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, pipeline := range listPipelinesResponse.Items {
		id := *pipeline.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PipelineId", id)
	}
	return resourceIds, nil
}

func GoldenGatePipelineSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if pipelineResponse, ok := response.Response.(oci_golden_gate.GetPipelineResponse); ok {
		return pipelineResponse.GetLifecycleState() != oci_golden_gate.PipelineLifecycleStateDeleted
	}
	return false
}

func GoldenGatePipelineSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GoldenGateClient().GetPipeline(context.Background(), oci_golden_gate.GetPipelineRequest{
		PipelineId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
