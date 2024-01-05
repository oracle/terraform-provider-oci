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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CloudBridgeAgentRequiredOnlyResource = CloudBridgeAgentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Required, acctest.Create, CloudBridgeAgentRepresentation)

	CloudBridgeAgentResourceConfig = CloudBridgeAgentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Optional, acctest.Update, CloudBridgeAgentRepresentation)

	CloudBridgeCloudBridgeAgentSingularDataSourceRepresentation = map[string]interface{}{
		"agent_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_bridge_agent.test_agent.id}`},
	}

	CloudBridgeCloudBridgeAgentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"agent_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_bridge_agent.test_agent.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"environment_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_bridge_environment.test_environment.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAgentDataSourceFilterRepresentation}}
	CloudBridgeAgentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_bridge_agent.test_agent.id}`}},
	}

	CloudBridgeAgentRepresentation = map[string]interface{}{
		"agent_type":     acctest.Representation{RepType: acctest.Required, Create: `APPLIANCE`},
		"agent_version":  acctest.Representation{RepType: acctest.Required, Create: `agentVersion`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_bridge_environment.test_environment.id}`},
		"os_version":     acctest.Representation{RepType: acctest.Required, Create: `osVersion`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}

	CloudBridgeAgentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Required, acctest.Create, CloudBridgeEnvironmentRepresentation)
)

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeAgentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeAgentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_bridge_agent.test_agent"
	datasourceName := "data.oci_cloud_bridge_agents.test_agents"
	singularDatasourceName := "data.oci_cloud_bridge_agent.test_agent"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudBridgeAgentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Optional, acctest.Create, CloudBridgeAgentRepresentation), "cloudbridge", "agent", t)

	acctest.ResourceTest(t, testAccCheckCloudBridgeAgentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudBridgeAgentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Required, acctest.Create, CloudBridgeAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_type", "APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "agent_version", "agentVersion"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "environment_id"),
				resource.TestCheckResourceAttr(resourceName, "os_version", "osVersion"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudBridgeAgentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudBridgeAgentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Optional, acctest.Create, CloudBridgeAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_type", "APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "agent_version", "agentVersion"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_version", "osVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CloudBridgeAgentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Optional, acctest.Update, CloudBridgeAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_type", "APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "agent_version", "agentVersion"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_version", "osVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_agents", "test_agents", acctest.Optional, acctest.Update, CloudBridgeCloudBridgeAgentDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeAgentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Optional, acctest.Update, CloudBridgeAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "agent_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "environment_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "agent_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "agent_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Required, acctest.Create, CloudBridgeCloudBridgeAgentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeAgentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "agent_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "agent_type", "APPLIANCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "agent_version", "agentVersion"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_version", "osVersion"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + CloudBridgeAgentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"plugin_list",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCloudBridgeAgentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OcbAgentSvcClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_bridge_agent" {
			noResourceFound = false
			request := oci_cloud_bridge.GetAgentRequest{}

			tmp := rs.Primary.ID
			request.AgentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")

			response, err := client.GetAgent(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_bridge.AgentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudBridgeAgent") {
		resource.AddTestSweepers("CloudBridgeAgent", &resource.Sweeper{
			Name:         "CloudBridgeAgent",
			Dependencies: acctest.DependencyGraph["agent"],
			F:            sweepCloudBridgeAgentResource,
		})
	}
}

func sweepCloudBridgeAgentResource(compartment string) error {
	ocbAgentSvcClient := acctest.GetTestClients(&schema.ResourceData{}).OcbAgentSvcClient()
	agentIds, err := getCloudBridgeAgentIds(compartment)
	if err != nil {
		return err
	}
	for _, agentId := range agentIds {
		if ok := acctest.SweeperDefaultResourceId[agentId]; !ok {
			deleteAgentRequest := oci_cloud_bridge.DeleteAgentRequest{}

			deleteAgentRequest.AgentId = &agentId

			deleteAgentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")
			_, error := ocbAgentSvcClient.DeleteAgent(context.Background(), deleteAgentRequest)
			if error != nil {
				fmt.Printf("Error deleting Agent %s %s, It is possible that the resource is already deleted. Please verify manually \n", agentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &agentId, CloudBridgeAgentSweepWaitCondition, time.Duration(3*time.Minute),
				CloudBridgeAgentSweepResponseFetchOperation, "cloud_bridge", true)
		}
	}
	return nil
}

func getCloudBridgeAgentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AgentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	ocbAgentSvcClient := acctest.GetTestClients(&schema.ResourceData{}).OcbAgentSvcClient()

	listAgentsRequest := oci_cloud_bridge.ListAgentsRequest{}
	listAgentsRequest.CompartmentId = &compartmentId
	listAgentsRequest.LifecycleState = oci_cloud_bridge.AgentLifecycleStateActive
	listAgentsResponse, err := ocbAgentSvcClient.ListAgents(context.Background(), listAgentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Agent list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, agent := range listAgentsResponse.Items {
		id := *agent.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AgentId", id)
	}
	return resourceIds, nil
}

func CloudBridgeAgentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if agentResponse, ok := response.Response.(oci_cloud_bridge.GetAgentResponse); ok {
		return agentResponse.LifecycleState != oci_cloud_bridge.AgentLifecycleStateDeleted
	}
	return false
}

func CloudBridgeAgentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OcbAgentSvcClient().GetAgent(context.Background(), oci_cloud_bridge.GetAgentRequest{
		AgentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
