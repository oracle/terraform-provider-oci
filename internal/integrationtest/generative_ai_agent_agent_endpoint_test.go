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
	GenerativeAiAgentAgentEndpointRequiredOnlyResource = GenerativeAiAgentAgentEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoint", "test_agent_endpoint", acctest.Required, acctest.Create, GenerativeAiAgentAgentEndpointRepresentation)

	GenerativeAiAgentAgentEndpointResourceConfig = GenerativeAiAgentAgentEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoint", "test_agent_endpoint", acctest.Optional, acctest.Update, GenerativeAiAgentAgentEndpointRepresentation)

	GenerativeAiAgentAgentEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"agent_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_agent_agent_endpoint.test_agent_endpoint.id}`},
	}

	GenerativeAiAgentAgentEndpointDataSourceRepresentation = map[string]interface{}{
		"agent_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.agent_id_env}`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentAgentEndpointDataSourceFilterRepresentation}}
	GenerativeAiAgentAgentEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_agent_agent_endpoint.test_agent_endpoint.id}`}},
	}

	GenerativeAiAgentAgentEndpointRepresentation = map[string]interface{}{
		"agent_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id_env}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"content_moderation_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentAgentEndpointContentModerationConfigRepresentation},
		// "defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"session_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentAgentEndpointSessionConfigRepresentation},
		"should_enable_citation": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_enable_session":  acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"should_enable_trace":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	GenerativeAiAgentAgentEndpointContentModerationConfigRepresentation = map[string]interface{}{
		"should_enable_on_input":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_enable_on_output": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	GenerativeAiAgentAgentEndpointSessionConfigRepresentation = map[string]interface{}{
		"idle_timeout_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `4000`, Update: `5000`},
	}

	GenerativeAiAgentAgentEndpointResourceDependencies = `` //Cannot test from home region, commented out - DefinedTagsDependencies
)

// issue-routing-tag: generative_ai_agent/default
func TestGenerativeAiAgentAgentEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentAgentEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	agentId := utils.GetEnvSettingWithBlankDefault("agentId_for_create")
	agentIdUVariableStr := fmt.Sprintf("variable \"agent_id_env\" { default = \"%s\" }\n", agentId)

	resourceName := "oci_generative_ai_agent_agent_endpoint.test_agent_endpoint"
	datasourceName := "data.oci_generative_ai_agent_agent_endpoints.test_agent_endpoints"
	singularDatasourceName := "data.oci_generative_ai_agent_agent_endpoint.test_agent_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdUVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoint", "test_agent_endpoint", acctest.Optional, acctest.Create, GenerativeAiAgentAgentEndpointRepresentation), "generativeaiagent", "agentEndpoint", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentAgentEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + agentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoint", "test_agent_endpoint", acctest.Required, acctest.Create, GenerativeAiAgentAgentEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + agentIdUVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + agentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoint", "test_agent_endpoint", acctest.Optional, acctest.Create, GenerativeAiAgentAgentEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.0.should_enable_on_input", "false"),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.0.should_enable_on_output", "false"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "session_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "session_config.0.idle_timeout_in_seconds", "4000"),
				resource.TestCheckResourceAttr(resourceName, "should_enable_citation", "false"),
				resource.TestCheckResourceAttr(resourceName, "should_enable_session", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_enable_trace", "false"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + agentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoint", "test_agent_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiAgentAgentEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.0.should_enable_on_input", "false"),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.0.should_enable_on_output", "false"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "session_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "session_config.0.idle_timeout_in_seconds", "4000"),
				resource.TestCheckResourceAttr(resourceName, "should_enable_citation", "false"),
				resource.TestCheckResourceAttr(resourceName, "should_enable_session", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_enable_trace", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + agentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoint", "test_agent_endpoint", acctest.Optional, acctest.Update, GenerativeAiAgentAgentEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.0.should_enable_on_input", "true"),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.0.should_enable_on_output", "true"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "session_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "session_config.0.idle_timeout_in_seconds", "5000"),
				resource.TestCheckResourceAttr(resourceName, "should_enable_citation", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_enable_session", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_enable_trace", "true"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoints", "test_agent_endpoints", acctest.Optional, acctest.Update, GenerativeAiAgentAgentEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoint", "test_agent_endpoint", acctest.Optional, acctest.Update, GenerativeAiAgentAgentEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "agent_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "agent_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "agent_endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoint", "test_agent_endpoint", acctest.Required, acctest.Create, GenerativeAiAgentAgentEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdUVariableStr + GenerativeAiAgentAgentEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "agent_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_moderation_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_moderation_config.0.should_enable_on_input", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_moderation_config.0.should_enable_on_output", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "session_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "session_config.0.idle_timeout_in_seconds", "5000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "should_enable_citation", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "should_enable_session", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "should_enable_trace", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiAgentAgentEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiAgentAgentEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_agent_agent_endpoint" {
			noResourceFound = false
			request := oci_generative_ai_agent.GetAgentEndpointRequest{}

			tmp := rs.Primary.ID
			request.AgentEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")

			response, err := client.GetAgentEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai_agent.AgentEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiAgentAgentEndpoint") {
		resource.AddTestSweepers("GenerativeAiAgentAgentEndpoint", &resource.Sweeper{
			Name:         "GenerativeAiAgentAgentEndpoint",
			Dependencies: acctest.DependencyGraph["agentEndpoint"],
			F:            sweepGenerativeAiAgentAgentEndpointResource,
		})
	}
}

func sweepGenerativeAiAgentAgentEndpointResource(compartment string) error {
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()
	agentEndpointIds, err := getGenerativeAiAgentAgentEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, agentEndpointId := range agentEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[agentEndpointId]; !ok {
			deleteAgentEndpointRequest := oci_generative_ai_agent.DeleteAgentEndpointRequest{}

			deleteAgentEndpointRequest.AgentEndpointId = &agentEndpointId

			deleteAgentEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")
			_, error := generativeAiAgentClient.DeleteAgentEndpoint(context.Background(), deleteAgentEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting AgentEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", agentEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &agentEndpointId, GenerativeAiAgentAgentEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiAgentAgentEndpointSweepResponseFetchOperation, "generative_ai_agent", true)
		}
	}
	return nil
}

func getGenerativeAiAgentAgentEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AgentEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()

	listAgentEndpointsRequest := oci_generative_ai_agent.ListAgentEndpointsRequest{}
	listAgentEndpointsRequest.CompartmentId = &compartmentId
	listAgentEndpointsRequest.LifecycleState = oci_generative_ai_agent.AgentEndpointLifecycleStateActive
	listAgentEndpointsResponse, err := generativeAiAgentClient.ListAgentEndpoints(context.Background(), listAgentEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AgentEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, agentEndpoint := range listAgentEndpointsResponse.Items {
		id := *agentEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AgentEndpointId", id)
	}
	return resourceIds, nil
}

func GenerativeAiAgentAgentEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if agentEndpointResponse, ok := response.Response.(oci_generative_ai_agent.GetAgentEndpointResponse); ok {
		return agentEndpointResponse.LifecycleState != oci_generative_ai_agent.AgentEndpointLifecycleStateDeleted
	}
	return false
}

func GenerativeAiAgentAgentEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiAgentClient().GetAgentEndpoint(context.Background(), oci_generative_ai_agent.GetAgentEndpointRequest{
		AgentEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
