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
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiEndpointRequiredOnlyResource = GenerativeAiEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_endpoint", "test_endpoint", acctest.Required, acctest.Create, GenerativeAiEndpointRepresentation)

	GenerativeAiEndpointResourceConfig = GenerativeAiEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_endpoint", "test_endpoint", acctest.Optional, acctest.Update, GenerativeAiEndpointRepresentation)

	GenerativeAiEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_endpoint.test_endpoint.id}`},
	}

	GenerativeAiEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_endpoint.test_endpoint.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiEndpointDataSourceFilterRepresentation}}
	GenerativeAiEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_endpoint.test_endpoint.id}`}},
	}

	GenerativeAiEndpointRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dedicated_ai_cluster_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster.id}`},
		"model_id":                  acctest.Representation{RepType: acctest.Required, Create: `${local.servering_model_id}`},
		"content_moderation_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiEndpointContentModerationConfigRepresentation},
		// "defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	GenerativeAiEndpointContentModerationConfigRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	GenerativeAiEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Required, acctest.Create, GenerativeAiHostingDedicatedAiClusterRepresentation) +
		servingModelDependencies
	// DefinedTagsDependencies + - no test in home region

	servingModelDependencies = `
	locals {
	
	  filtered_models = [
		for item in data.oci_generative_ai_models.serving_models.model_collection[0].items : item
		  if (
			(item.version == "14.2")
			&& length(item.capabilities) == 1
			&& (item.display_name == "cohere.command-light")
		  )
		]
	
	  servering_model_id = local.filtered_models[0].id
	}
	
	data "oci_generative_ai_models" "serving_models" {
	  compartment_id = var.compartment_id
	  display_name = "cohere.command-light"
	}
	`
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_endpoint.test_endpoint"
	datasourceName := "data.oci_generative_ai_endpoints.test_endpoints"
	singularDatasourceName := "data.oci_generative_ai_endpoint.test_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_endpoint", "test_endpoint", acctest.Optional, acctest.Create, GenerativeAiEndpointRepresentation), "generativeai", "endpoint", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_endpoint", "test_endpoint", acctest.Required, acctest.Create, GenerativeAiEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_endpoint", "test_endpoint", acctest.Optional, acctest.Create, GenerativeAiEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.0.is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GenerativeAiEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_endpoint", "test_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.0.is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
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
			Config: config + compartmentIdVariableStr + GenerativeAiEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_endpoint", "test_endpoint", acctest.Optional, acctest.Update, GenerativeAiEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_moderation_config.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_endpoints", "test_endpoints", acctest.Optional, acctest.Update, GenerativeAiEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_endpoint", "test_endpoint", acctest.Optional, acctest.Update, GenerativeAiEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_endpoint", "test_endpoint", acctest.Required, acctest.Create, GenerativeAiEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "endpoint_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_moderation_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_moderation_config.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_endpoint" {
			noResourceFound = false
			request := oci_generative_ai.GetEndpointRequest{}

			tmp := rs.Primary.ID
			request.EndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.EndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiEndpoint") {
		resource.AddTestSweepers("GenerativeAiEndpoint", &resource.Sweeper{
			Name:         "GenerativeAiEndpoint",
			Dependencies: acctest.DependencyGraph["endpoint"],
			F:            sweepGenerativeAiEndpointResource,
		})
	}
}

func sweepGenerativeAiEndpointResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	endpointIds, err := getGenerativeAiEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, endpointId := range endpointIds {
		if ok := acctest.SweeperDefaultResourceId[endpointId]; !ok {
			deleteEndpointRequest := oci_generative_ai.DeleteEndpointRequest{}

			deleteEndpointRequest.EndpointId = &endpointId

			deleteEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteEndpoint(context.Background(), deleteEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting Endpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", endpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &endpointId, GenerativeAiEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiEndpointSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listEndpointsRequest := oci_generative_ai.ListEndpointsRequest{}
	listEndpointsRequest.CompartmentId = &compartmentId
	listEndpointsRequest.LifecycleState = oci_generative_ai.EndpointLifecycleStateActive
	listEndpointsResponse, err := generativeAiClient.ListEndpoints(context.Background(), listEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Endpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, endpoint := range listEndpointsResponse.Items {
		id := *endpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "EndpointId", id)
	}
	return resourceIds, nil
}

func GenerativeAiEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is ACTIVE beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if endpointResponse, ok := response.Response.(oci_generative_ai.GetEndpointResponse); ok {
		return endpointResponse.LifecycleState != oci_generative_ai.EndpointLifecycleStateDeleted
	}
	return false
}

func GenerativeAiEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetEndpoint(context.Background(), oci_generative_ai.GetEndpointRequest{
		EndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
