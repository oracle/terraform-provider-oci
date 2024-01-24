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
	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AiLanguageEndpointRequiredOnlyResource = AiLanguageEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Required, acctest.Create, AiLanguageEndpointRepresentation)

	AiLanguageEndpointResourceConfig = AiLanguageEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Optional, acctest.Update, AiLanguageEndpointRepresentation)

	AiLanguageAiLanguageEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_language_endpoint.test_endpoint.id}`},
	}

	AiLanguageAiLanguageEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		// "endpoint_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_language_endpoint.test_endpoint.id}`},
		"model_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_language_model.test_model.id}`},
		"project_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_language_project.test_project.id}`},
		"state":      acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":     acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageEndpointDataSourceFilterRepresentation}}
	AiLanguageEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_language_endpoint.test_endpoint.id}`}},
	}

	AiLanguageEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_language_model.test_model.id}`},
		// "defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"inference_units": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}

	AiLanguageEndpointResourceDependencies =
	// acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Required, acctest.Create, AiLanguageEndpointRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Required, acctest.Create, AiLanguageModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_project", "test_project", acctest.Required, acctest.Create, AiLanguageProjectRepresentation) +
		// acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Required, acctest.Create, datasetRepresentation) +
		DefinedTagsDependencies
	// +
	// acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
	// acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: ai_language/default
func TestAiLanguageEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiLanguageEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_language_endpoint.test_endpoint"
	datasourceName := "data.oci_ai_language_endpoints.test_endpoints"
	singularDatasourceName := "data.oci_ai_language_endpoint.test_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiLanguageEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Optional, acctest.Create, AiLanguageEndpointRepresentation), "ailanguage", "endpoint", t)

	acctest.ResourceTest(t, testAccCheckAiLanguageEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiLanguageEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Required, acctest.Create, AiLanguageEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiLanguageEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiLanguageEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Optional, acctest.Create, AiLanguageEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inference_units", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiLanguageEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AiLanguageEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inference_units", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + AiLanguageEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Optional, acctest.Update, AiLanguageEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inference_units", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_language_endpoints", "test_endpoints", acctest.Optional, acctest.Update, AiLanguageAiLanguageEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + AiLanguageEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Optional, acctest.Update, AiLanguageEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				// resource.TestCheckResourceAttrSet(datasourceName, "endpoint_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Required, acctest.Create, AiLanguageAiLanguageEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiLanguageEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inference_units", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + AiLanguageEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAiLanguageEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AiServiceLanguageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_language_endpoint" {
			noResourceFound = false
			request := oci_ai_language.GetEndpointRequest{}

			tmp := rs.Primary.ID
			request.EndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_language")

			response, err := client.GetEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_language.EndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiLanguageEndpoint") {
		resource.AddTestSweepers("AiLanguageEndpoint", &resource.Sweeper{
			Name:         "AiLanguageEndpoint",
			Dependencies: acctest.DependencyGraph["endpoint"],
			F:            sweepAiLanguageEndpointResource,
		})
	}
}

func sweepAiLanguageEndpointResource(compartment string) error {
	aiServiceLanguageClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceLanguageClient()
	endpointIds, err := getAiLanguageEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, endpointId := range endpointIds {
		if ok := acctest.SweeperDefaultResourceId[endpointId]; !ok {
			deleteEndpointRequest := oci_ai_language.DeleteEndpointRequest{}

			deleteEndpointRequest.EndpointId = &endpointId

			deleteEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_language")
			_, error := aiServiceLanguageClient.DeleteEndpoint(context.Background(), deleteEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting Endpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", endpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &endpointId, AiLanguageEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				AiLanguageEndpointSweepResponseFetchOperation, "ai_language", true)
		}
	}
	return nil
}

func getAiLanguageEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	aiServiceLanguageClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceLanguageClient()

	listEndpointsRequest := oci_ai_language.ListEndpointsRequest{}
	listEndpointsRequest.CompartmentId = &compartmentId
	listEndpointsRequest.LifecycleState = oci_ai_language.EndpointLifecycleStateActive
	listEndpointsResponse, err := aiServiceLanguageClient.ListEndpoints(context.Background(), listEndpointsRequest)

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

func AiLanguageEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if endpointResponse, ok := response.Response.(oci_ai_language.GetEndpointResponse); ok {
		return endpointResponse.LifecycleState != oci_ai_language.EndpointLifecycleStateDeleted
	}
	return false
}

func AiLanguageEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AiServiceLanguageClient().GetEndpoint(context.Background(), oci_ai_language.GetEndpointRequest{
		EndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
