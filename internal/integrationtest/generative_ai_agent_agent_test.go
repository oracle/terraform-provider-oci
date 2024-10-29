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
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiAgentAgentRequiredOnlyResource = GenerativeAiAgentAgentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent", "test_agent", acctest.Required, acctest.Create, GenerativeAiAgentAgentRepresentation)

	GenerativeAiAgentAgentResourceConfig = GenerativeAiAgentAgentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent", "test_agent", acctest.Optional, acctest.Update, GenerativeAiAgentAgentRepresentation)

	GenerativeAiAgentAgentSingularDataSourceRepresentation = map[string]interface{}{
		"agent_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_agent_agent.test_agent.id}`},
	}

	GenerativeAiAgentAgentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentAgentDataSourceFilterRepresentation}}
	GenerativeAiAgentAgentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_agent_agent.test_agent.id}`}},
	}

	GenerativeAiAgentAgentRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		// "defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"knowledge_base_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.knowledgeBaseId_env}`}},
		"welcome_message":    acctest.Representation{RepType: acctest.Optional, Create: `welcomeMessage`, Update: `welcomeMessage2`},
	}

	GenerativeAiAgentAgentResourceDependencies = `` //Cannot test from home region, commented out - DefinedTagsDependencies
)

// issue-routing-tag: generative_ai_agent/default
func TestGenerativeAiAgentAgentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentAgentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	knowledgeBaseId := utils.GetEnvSettingWithBlankDefault("knowledgeBaseId_for_update")
	knowledgeBaseIdUVariableStr := fmt.Sprintf("variable \"knowledgeBaseId_env\" { default = \"%s\" }\n", knowledgeBaseId)

	resourceName := "oci_generative_ai_agent_agent.test_agent"
	datasourceName := "data.oci_generative_ai_agent_agents.test_agents"
	singularDatasourceName := "data.oci_generative_ai_agent_agent.test_agent"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+knowledgeBaseIdUVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent", "test_agent", acctest.Optional, acctest.Create, GenerativeAiAgentAgentRepresentation), "generativeaiagent", "agent", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentAgentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent", "test_agent", acctest.Required, acctest.Create, GenerativeAiAgentAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + knowledgeBaseIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent", "test_agent", acctest.Optional, acctest.Create, GenerativeAiAgentAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "knowledge_base_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "welcome_message", "welcomeMessage"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + knowledgeBaseIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent", "test_agent", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiAgentAgentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "knowledge_base_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "welcome_message", "welcomeMessage"),

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
			Config: config + compartmentIdVariableStr + knowledgeBaseIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent", "test_agent", acctest.Optional, acctest.Update, GenerativeAiAgentAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "knowledge_base_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "welcome_message", "welcomeMessage2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_agents", "test_agents", acctest.Optional, acctest.Update, GenerativeAiAgentAgentDataSourceRepresentation) +
				compartmentIdVariableStr + knowledgeBaseIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent", "test_agent", acctest.Optional, acctest.Update, GenerativeAiAgentAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "agent_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "agent_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_agent", "test_agent", acctest.Required, acctest.Create, GenerativeAiAgentAgentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + knowledgeBaseIdUVariableStr + GenerativeAiAgentAgentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "agent_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "knowledge_base_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "welcome_message", "welcomeMessage2"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiAgentAgentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiAgentAgentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_agent_agent" {
			noResourceFound = false
			request := oci_generative_ai_agent.GetAgentRequest{}

			tmp := rs.Primary.ID
			request.AgentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")

			response, err := client.GetAgent(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai_agent.AgentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiAgentAgent") {
		resource.AddTestSweepers("GenerativeAiAgentAgent", &resource.Sweeper{
			Name:         "GenerativeAiAgentAgent",
			Dependencies: acctest.DependencyGraph["agent"],
			F:            sweepGenerativeAiAgentAgentResource,
		})
	}
}

func sweepGenerativeAiAgentAgentResource(compartment string) error {
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()
	agentIds, err := getGenerativeAiAgentAgentIds(compartment)
	if err != nil {
		return err
	}
	for _, agentId := range agentIds {
		if ok := acctest.SweeperDefaultResourceId[agentId]; !ok {
			deleteAgentRequest := oci_generative_ai_agent.DeleteAgentRequest{}

			deleteAgentRequest.AgentId = &agentId

			deleteAgentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")
			_, error := generativeAiAgentClient.DeleteAgent(context.Background(), deleteAgentRequest)
			if error != nil {
				fmt.Printf("Error deleting Agent %s %s, It is possible that the resource is already deleted. Please verify manually \n", agentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &agentId, GenerativeAiAgentAgentSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiAgentAgentSweepResponseFetchOperation, "generative_ai_agent", true)
		}
	}
	return nil
}

func getGenerativeAiAgentAgentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AgentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()

	listAgentsRequest := oci_generative_ai_agent.ListAgentsRequest{}
	listAgentsRequest.CompartmentId = &compartmentId
	listAgentsRequest.LifecycleState = oci_generative_ai_agent.AgentLifecycleStateActive
	listAgentsResponse, err := generativeAiAgentClient.ListAgents(context.Background(), listAgentsRequest)

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

func GenerativeAiAgentAgentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if agentResponse, ok := response.Response.(oci_generative_ai_agent.GetAgentResponse); ok {
		return agentResponse.LifecycleState != oci_generative_ai_agent.AgentLifecycleStateDeleted
	}
	return false
}

func GenerativeAiAgentAgentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiAgentClient().GetAgent(context.Background(), oci_generative_ai_agent.GetAgentRequest{
		AgentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
