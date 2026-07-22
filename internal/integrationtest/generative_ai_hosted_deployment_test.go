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
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiHostedDeploymentRequiredOnlyResource = GenerativeAiHostedDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_deployment", "test_hosted_deployment", acctest.Required, acctest.Create, GenerativeAiHostedDeploymentRepresentation)

	GenerativeAiHostedDeploymentResourceConfig = GenerativeAiHostedDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_deployment", "test_hosted_deployment", acctest.Optional, acctest.Update, GenerativeAiHostedDeploymentRepresentation)

	GenerativeAiHostedDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"hosted_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_hosted_deployment.test_hosted_deployment.id}`},
	}

	GenerativeAiHostedDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"application_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_hosted_application.test_hosted_application.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_hosted_deployment.test_hosted_deployment.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiHostedDeploymentDataSourceFilterRepresentation}}
	GenerativeAiHostedDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_hosted_deployment.test_hosted_deployment.id}`}},
	}

	GenerativeAiHostedDeploymentRepresentation = map[string]interface{}{
		"active_artifact":       acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiHostedDeploymentActiveArtifactRepresentation},
		"hosted_application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_hosted_application.test_hosted_application.id}`},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	GenerativeAiHostedDeploymentActiveArtifactRepresentation = map[string]interface{}{
		"artifact_type": acctest.Representation{RepType: acctest.Optional, Create: `SIMPLE_DOCKER_ARTIFACT`},
		"container_uri": acctest.Representation{RepType: acctest.Optional, Create: `${var.region}.ocir.io/axk4z7krhqfx/cost-service`, Update: `${var.region}.ocir.io/axk4z7krhqfx/cost-service`},
		"tag":           acctest.Representation{RepType: acctest.Optional, Create: `latest`, Update: `latest`},
	}

	GenerativeAiHostedDeploymentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application", "test_hosted_application", acctest.Required, acctest.Create, GenerativeAiHostedApplicationRepresentation)
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiHostedDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiHostedDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_generative_ai_hosted_deployment.test_hosted_deployment"
	datasourceName := "data.oci_generative_ai_hosted_deployments.test_hosted_deployments"
	singularDatasourceName := "data.oci_generative_ai_hosted_deployment.test_hosted_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiHostedDeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_deployment", "test_hosted_deployment", acctest.Optional, acctest.Create, GenerativeAiHostedDeploymentRepresentation), "generativeai", "hostedDeployment", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiHostedDeploymentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_deployment", "test_hosted_deployment", acctest.Required, acctest.Create, GenerativeAiHostedDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active_artifact.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hosted_application_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedDeploymentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_deployment", "test_hosted_deployment", acctest.Optional, acctest.Create, GenerativeAiHostedDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active_artifact.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "active_artifact.0.artifact_type", "SIMPLE_DOCKER_ARTIFACT"),
				resource.TestCheckResourceAttr(resourceName, "active_artifact.0.container_uri", fmt.Sprintf("%s.ocir.io/axk4z7krhqfx/cost-service", utils.GetEnvSettingWithBlankDefault("region"))),
				resource.TestCheckResourceAttrSet(resourceName, "active_artifact.0.hosted_deployment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "active_artifact.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "active_artifact.0.status"),
				resource.TestCheckResourceAttr(resourceName, "active_artifact.0.tag", "latest"),
				resource.TestCheckResourceAttrSet(resourceName, "active_artifact.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "artifacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hosted_application_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + GenerativeAiHostedDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_deployment", "test_hosted_deployment", acctest.Optional, acctest.Update, GenerativeAiHostedDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active_artifact.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "active_artifact.0.artifact_type", "SIMPLE_DOCKER_ARTIFACT"),
				resource.TestCheckResourceAttr(resourceName, "active_artifact.0.container_uri", fmt.Sprintf("%s.ocir.io/axk4z7krhqfx/cost-service", utils.GetEnvSettingWithBlankDefault("region"))),
				resource.TestCheckResourceAttrSet(resourceName, "active_artifact.0.hosted_deployment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "active_artifact.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "active_artifact.0.status"),
				resource.TestCheckResourceAttr(resourceName, "active_artifact.0.tag", "latest"),
				resource.TestCheckResourceAttrSet(resourceName, "active_artifact.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "artifacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hosted_application_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_hosted_deployments", "test_hosted_deployments", acctest.Optional, acctest.Update, GenerativeAiHostedDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiHostedDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_deployment", "test_hosted_deployment", acctest.Optional, acctest.Update, GenerativeAiHostedDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "application_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "hosted_deployment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "hosted_deployment_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_hosted_deployment", "test_hosted_deployment", acctest.Required, acctest.Create, GenerativeAiHostedDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiHostedDeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hosted_deployment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "active_artifact.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "active_artifact.0.artifact_type", "SIMPLE_DOCKER_ARTIFACT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "active_artifact.0.container_uri", fmt.Sprintf("%s.ocir.io/axk4z7krhqfx/cost-service", utils.GetEnvSettingWithBlankDefault("region"))),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "active_artifact.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "active_artifact.0.status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "active_artifact.0.tag", "latest"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "active_artifact.0.time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "artifacts.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiHostedDeploymentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiHostedDeploymentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_hosted_deployment" {
			noResourceFound = false
			request := oci_generative_ai.GetHostedDeploymentRequest{}

			tmp := rs.Primary.ID
			request.HostedDeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetHostedDeployment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.HostedDeploymentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiHostedDeployment") {
		resource.AddTestSweepers("GenerativeAiHostedDeployment", &resource.Sweeper{
			Name:         "GenerativeAiHostedDeployment",
			Dependencies: acctest.DependencyGraph["hostedDeployment"],
			F:            sweepGenerativeAiHostedDeploymentResource,
		})
	}
}

func sweepGenerativeAiHostedDeploymentResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	hostedDeploymentIds, err := getGenerativeAiHostedDeploymentIds(compartment)
	if err != nil {
		return err
	}
	for _, hostedDeploymentId := range hostedDeploymentIds {
		if ok := acctest.SweeperDefaultResourceId[hostedDeploymentId]; !ok {
			deleteHostedDeploymentRequest := oci_generative_ai.DeleteHostedDeploymentRequest{}

			deleteHostedDeploymentRequest.HostedDeploymentId = &hostedDeploymentId

			deleteHostedDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteHostedDeployment(context.Background(), deleteHostedDeploymentRequest)
			if error != nil {
				fmt.Printf("Error deleting HostedDeployment %s %s, It is possible that the resource is already deleted. Please verify manually \n", hostedDeploymentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &hostedDeploymentId, GenerativeAiHostedDeploymentSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiHostedDeploymentSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiHostedDeploymentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "HostedDeploymentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listHostedDeploymentsRequest := oci_generative_ai.ListHostedDeploymentsRequest{}
	listHostedDeploymentsRequest.CompartmentId = &compartmentId
	listHostedDeploymentsRequest.LifecycleState = oci_generative_ai.HostedDeploymentLifecycleStateActive
	listHostedDeploymentsResponse, err := generativeAiClient.ListHostedDeployments(context.Background(), listHostedDeploymentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting HostedDeployment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, hostedDeployment := range listHostedDeploymentsResponse.Items {
		id := *hostedDeployment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "HostedDeploymentId", id)
	}
	return resourceIds, nil
}

func GenerativeAiHostedDeploymentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if hostedDeploymentResponse, ok := response.Response.(oci_generative_ai.GetHostedDeploymentResponse); ok {
		return hostedDeploymentResponse.LifecycleState != oci_generative_ai.HostedDeploymentLifecycleStateDeleted
	}
	return false
}

func GenerativeAiHostedDeploymentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetHostedDeployment(context.Background(), oci_generative_ai.GetHostedDeploymentRequest{
		HostedDeploymentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
