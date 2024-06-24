// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package integrationtest

//
//import (
//	"context"
//	"fmt"
//	"strconv"
//	"testing"
//	"time"
//
//	"github.com/oracle/terraform-provider-oci/internal/acctest"
//	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
//	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
//	"github.com/oracle/terraform-provider-oci/internal/tfresource"
//	"github.com/oracle/terraform-provider-oci/internal/utils"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
//	"github.com/oracle/oci-go-sdk/v65/common"
//	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"
//
//	"github.com/oracle/terraform-provider-oci/httpreplay"
//)
//
//var (
//	DatabaseMigrationAgentRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", acctest.Required, acctest.Create, agentRepresentation2)
//
//	DatabaseMigrationAgentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", acctest.Optional, acctest.Update, agentRepresentation2)
//
//	DatabaseMigrationagentSingularDataSourceRepresentation = map[string]interface{}{
//		"agent_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_agent.test_agent.id}`},
//	}
//
//	DatabaseMigrationagentDataSourceRepresentation = map[string]interface{}{
//		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
//		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TF_displayName`, Update: `TF_displayName2`},
//		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
//		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: agentDataSourceFilterRepresentation}}
//	agentDataSourceRepresentation2 = map[string]interface{}{
//		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
//		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: agentDataSourceFilterRepresentation}}
//	agentDataSourceFilterRepresentation = map[string]interface{}{
//		"name":   acctest.Representation{RepType: acctest.Required, Create: `agent_id`},
//		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_agent.test_agent.id}`}},
//	}
//
//	agentRepresentation = map[string]interface{}{
//		"agent_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_agent.test_agent.id}`},
//		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
//		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TF_displayName`, Update: `TF_displayName2`},
//		"stream_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_streaming_stream.test_stream.id}`},
//		"version":        acctest.Representation{RepType: acctest.Optional, Create: `version`, Update: `version2`},
//	}
//
//	agentRepresentation2 = map[string]interface{}{
//		"agent_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_agent.test_agent.id}`},
//		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
//		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TF_displayName`, Update: `TF_displayName2`},
//		"stream_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_streaming_stream.test_stream.id}`},
//		"version":        acctest.Representation{RepType: acctest.Optional, Create: `version`, Update: `version2`},
//	}
//
//	DatabaseMigrationAgentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", acctest.Required, acctest.Create, agentRepresentation) +
//		DefinedTagsDependencies + acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Required, acctest.Create, StreamingStreamRepresentation)
//)
//
//// issue-routing-tag: database_migration/default
//func TestDatabaseMigrationAgentResource_basic(t *testing.T) {
//	t.Skip("Skip this test agent creation is an independent operation.")
//	httpreplay.SetScenario("TestDatabaseMigrationAgentResource_basic")
//	defer httpreplay.SaveScenario()
//
//	config := acctest.ProviderTestConfig()
//
//	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
//	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
//
//	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
//	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
//
//	resourceName := "oci_database_migration_agent.test_agent"
//	datasourceName := "data.oci_database_migration_agents.test_agents"
//	singularDatasourceName := "data.oci_database_migration_agent.test_agent"
//
//	var resId, resId2 string
//	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
//	acctest.SaveConfigContent(config+compartmentIdVariableStr+ //DatabaseMigrationAgentResourceDependencies+
//		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", acctest.Optional, acctest.Create, agentRepresentation), "databasemigration", "agent", t)
//
//	acctest.ResourceTest(t, testAccCheckDatabaseMigrationAgentDestroy, []resource.TestStep{
//		// verify Create
//		{
//			Config: config + compartmentIdVariableStr +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", acctest.Required, acctest.Create, agentRepresentation),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
//
//				func(s *terraform.State) (err error) {
//					resId, err = acctest.FromInstanceState(s, resourceName, "agent_id")
//					return err
//				},
//			),
//		},
//
//		// delete before next Create
//		{
//			Config: config + compartmentIdVariableStr,
//		},
//		// verify Create with optionals
//		{
//			Config: config + compartmentIdVariableStr +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", acctest.Optional, acctest.Create, agentRepresentation2),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
//				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
//				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_displayName"),
//				resource.TestCheckResourceAttrSet(resourceName, "id"),
//				resource.TestCheckResourceAttrSet(resourceName, "state"),
//				resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
//				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
//				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
//				resource.TestCheckResourceAttr(resourceName, "version", "version"),
//
//				func(s *terraform.State) (err error) {
//					resId, err = acctest.FromInstanceState(s, resourceName, "agent_id")
//					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
//						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
//							return errExport
//						}
//					}
//					return err
//				},
//			),
//		},
//
//		// verify Update to the compartment (the compartment will be switched back in the next step)
//		{
//			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", acctest.Optional, acctest.Create,
//					acctest.RepresentationCopyWithNewProperties(agentRepresentation2, map[string]interface{}{
//						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
//					})),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
//				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
//				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_displayName"),
//				resource.TestCheckResourceAttrSet(resourceName, "id"),
//				resource.TestCheckResourceAttrSet(resourceName, "state"),
//				resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
//				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
//				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
//				resource.TestCheckResourceAttr(resourceName, "version", "version"),
//
//				func(s *terraform.State) (err error) {
//					resId2, err = acctest.FromInstanceState(s, resourceName, "agent_id")
//					if resId != resId2 {
//						return fmt.Errorf("resource recreated when it was supposed to be updated")
//					}
//					return err
//				},
//			),
//		},
//
//		// verify updates to updatable parameters
//		{
//			Config: config + compartmentIdVariableStr +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", acctest.Optional, acctest.Update, agentRepresentation2),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
//				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
//				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_displayName2"),
//				resource.TestCheckResourceAttrSet(resourceName, "id"),
//				resource.TestCheckResourceAttrSet(resourceName, "state"),
//				resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
//				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
//				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
//				resource.TestCheckResourceAttr(resourceName, "version", "version2"),
//
//				func(s *terraform.State) (err error) {
//					resId2, err = acctest.FromInstanceState(s, resourceName, "agent_id")
//					if resId != resId2 {
//						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
//					}
//					return err
//				},
//			),
//		},
//		// verify datasource
//		{
//			Config: config +
//				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_agents", "test_agents", acctest.Optional, acctest.Update, agentDataSourceRepresentation2) +
//				compartmentIdVariableStr +
//				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_agent", "test_agent", acctest.Optional, acctest.Update, agentRepresentation2),
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
//				resource.TestCheckResourceAttr(datasourceName, "agent_collection.#", "1"),
//				resource.TestCheckResourceAttr(datasourceName, "agent_collection.0.items.#", "0"),
//			),
//		},
//		// verify singular datasource
//		{
//			Config: config +
//				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_agent", "test_agent", acctest.Required, acctest.Create, DatabaseMigrationagentSingularDataSourceRepresentation) +
//				compartmentIdVariableStr + DatabaseMigrationAgentResourceConfig,
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(singularDatasourceName, "agent_id"),
//				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
//				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TF_displayName2"),
//				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
//				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
//				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
//				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
//				resource.TestCheckResourceAttr(singularDatasourceName, "version", "version2"),
//			),
//		},
//		// verify resource import
//		{
//			Config:                  config + DatabaseMigrationAgentRequiredOnlyResource,
//			ImportState:             true,
//			ImportStateVerify:       false,
//			ImportStateVerifyIgnore: []string{},
//			ResourceName:            resourceName,
//		},
//	})
//}
//
//func testAccCheckDatabaseMigrationAgentDestroy(s *terraform.State) error {
//	noResourceFound := true
//	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseMigrationClient()
//	for _, rs := range s.RootModule().Resources {
//		if rs.Type == "oci_database_migration_agent" {
//			noResourceFound = false
//			request := oci_database_migration.GetAgentRequest{}
//
//			tmp := rs.Primary.ID
//			request.AgentId = &tmp
//
//			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")
//
//			response, err := client.GetAgent(context.Background(), request)
//
//			if err == nil {
//				deletedLifecycleStates := map[string]bool{
//					string(oci_database_migration.LifecycleStatesDeleted): true,
//				}
//				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
//					//resource lifecycle state is not in expected deleted lifecycle states.
//					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
//				}
//				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
//				continue
//			}
//
//			//Verify that exception is for '404 not found'.
//			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
//				return err
//			}
//		}
//	}
//	if noResourceFound {
//		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
//	}
//
//	return nil
//}
//
//func init() {
//	if acctest.DependencyGraph == nil {
//		acctest.InitDependencyGraph()
//	}
//	if !acctest.InSweeperExcludeList("DatabaseMigrationAgent") {
//		resource.AddTestSweepers("DatabaseMigrationAgent", &resource.Sweeper{
//			Name:         "DatabaseMigrationAgent",
//			Dependencies: acctest.DependencyGraph["agent"],
//			F:            sweepDatabaseMigrationAgentResource,
//		})
//	}
//}
//
//func sweepDatabaseMigrationAgentResource(compartment string) error {
//	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()
//	agentIds, err := getDatabaseMigrationAgentIds(compartment)
//	if err != nil {
//		return err
//	}
//	for _, agentId := range agentIds {
//		if ok := acctest.SweeperDefaultResourceId[agentId]; !ok {
//			deleteAgentRequest := oci_database_migration.DeleteAgentRequest{}
//
//			deleteAgentRequest.AgentId = &agentId
//
//			deleteAgentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")
//			_, error := databaseMigrationClient.DeleteAgent(context.Background(), deleteAgentRequest)
//			if error != nil {
//				fmt.Printf("Error deleting Agent %s %s, It is possible that the resource is already deleted. Please verify manually \n", agentId, error)
//				continue
//			}
//			acctest.WaitTillCondition(acctest.TestAccProvider, &agentId, DatabaseMigrationagentsSweepWaitCondition, time.Duration(3*time.Minute),
//				DatabaseMigrationagentsSweepResponseFetchOperation, "database_migration", true)
//		}
//	}
//	return nil
//}
//
//func getDatabaseMigrationAgentIds(compartment string) ([]string, error) {
//	ids := acctest.GetResourceIdsToSweep(compartment, "AgentId")
//	if ids != nil {
//		return ids, nil
//	}
//	var resourceIds []string
//	compartmentId := compartment
//	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()
//
//	listAgentsRequest := oci_database_migration.ListAgentsRequest{}
//	listAgentsRequest.CompartmentId = &compartmentId
//	listAgentsRequest.LifecycleState = oci_database_migration.ListAgentsLifecycleStateActive
//	listAgentsResponse, err := databaseMigrationClient.ListAgents(context.Background(), listAgentsRequest)
//
//	if err != nil {
//		return resourceIds, fmt.Errorf("Error getting Agent list for compartment id : %s , %s \n", compartmentId, err)
//	}
//	for _, agent := range listAgentsResponse.Items {
//		id := *agent.Id
//		resourceIds = append(resourceIds, id)
//		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AgentId", id)
//	}
//	return resourceIds, nil
//}
//
//func DatabaseMigrationagentsSweepWaitCondition(response common.OCIOperationResponse) bool {
//	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
//	if agentResponse, ok := response.Response.(oci_database_migration.GetAgentResponse); ok {
//		return agentResponse.LifecycleState != oci_database_migration.LifecycleStatesDeleted
//	}
//	return false
//}
//
//func DatabaseMigrationagentsSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
//	_, err := client.DatabaseMigrationClient().GetAgent(context.Background(), oci_database_migration.GetAgentRequest{
//		AgentId: resourceId,
//		RequestMetadata: common.RequestMetadata{
//			RetryPolicy: retryPolicy,
//		},
//	})
//	return err
//}
