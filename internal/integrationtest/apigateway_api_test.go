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
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApigatewayApiRequiredOnlyResource = ApigatewayApiResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Required, acctest.Create, ApigatewayApiRepresentation)

	ApigatewayApiResourceConfig = ApigatewayApiResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Optional, acctest.Update, ApigatewayApiRepresentation)

	ApigatewayApiSingularDataSourceRepresentation = map[string]interface{}{
		"api_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_api.test_api.id}`},
	}

	ApigatewayApiDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewayApiDataSourceFilterRepresentation}}
	ApigatewayApiDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apigateway_api.test_api.id}`}},
	}

	ApigatewayApiRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"content": acctest.Representation{RepType: acctest.Optional,
			Create: `{\"openapi\":\"3.0.0\",\"info\":{\"version\":\"1.0.0\",\"title\":\"test\",\"license\":{\"name\":\"MIT\"}},\"paths\":{\"/ping\":{\"get\":{\"responses\":{\"200\":{\"description\":\"OK\"}}}}}}`,
			Update: `{\"openapi\":\"3.0.0\",\"info\":{\"version\":\"1.0.0\",\"title\":\"test\"}}`},
	}

	ApigatewayApiResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: apigateway/default
func TestApigatewayApiResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayApiResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_apigateway_api.test_api"
	datasourceName := "data.oci_apigateway_apis.test_apis"
	singularDatasourceName := "data.oci_apigateway_api.test_api"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApigatewayApiResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Optional, acctest.Create, ApigatewayApiRepresentation), "apigateway", "api", t)

	acctest.ResourceTest(t, testAccCheckApigatewayApiDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ApigatewayApiResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Required, acctest.Create, ApigatewayApiRepresentation),
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
			Config: config + compartmentIdVariableStr + ApigatewayApiResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApigatewayApiResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Optional, acctest.Create, ApigatewayApiRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApigatewayApiResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ApigatewayApiRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config + compartmentIdVariableStr + ApigatewayApiResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Optional, acctest.Update, ApigatewayApiRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_apis", "test_apis", acctest.Optional, acctest.Update, ApigatewayApiDataSourceRepresentation) +
				compartmentIdVariableStr + ApigatewayApiResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Optional, acctest.Update, ApigatewayApiRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "api_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "api_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Required, acctest.Create, ApigatewayApiSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApigatewayApiResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "api_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "specification_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + ApigatewayApiRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"content",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckApigatewayApiDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ApiGatewayClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apigateway_api" {
			noResourceFound = false
			request := oci_apigateway.GetApiRequest{}

			tmp := rs.Primary.ID
			request.ApiId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apigateway")

			response, err := client.GetApi(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_apigateway.ApiLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ApigatewayApi") {
		resource.AddTestSweepers("ApigatewayApi", &resource.Sweeper{
			Name:         "ApigatewayApi",
			Dependencies: acctest.DependencyGraph["api"],
			F:            sweepApigatewayApiResource,
		})
	}
}

func sweepApigatewayApiResource(compartment string) error {
	apiGatewayClient := acctest.GetTestClients(&schema.ResourceData{}).ApiGatewayClient()
	apiIds, err := getApiIds(compartment)
	if err != nil {
		return err
	}
	for _, apiId := range apiIds {
		if ok := acctest.SweeperDefaultResourceId[apiId]; !ok {
			deleteApiRequest := oci_apigateway.DeleteApiRequest{}

			deleteApiRequest.ApiId = &apiId

			deleteApiRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apigateway")
			_, error := apiGatewayClient.DeleteApi(context.Background(), deleteApiRequest)
			if error != nil {
				fmt.Printf("Error deleting Api %s %s, It is possible that the resource is already deleted. Please verify manually \n", apiId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &apiId, apiSweepWaitCondition, time.Duration(3*time.Minute),
				apiSweepResponseFetchOperation, "apigateway", true)
		}
	}
	return nil
}

func getApiIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ApiId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apiGatewayClient := acctest.GetTestClients(&schema.ResourceData{}).ApiGatewayClient()

	listApisRequest := oci_apigateway.ListApisRequest{}
	listApisRequest.CompartmentId = &compartmentId
	listApisRequest.LifecycleState = oci_apigateway.ApiSummaryLifecycleStateActive
	listApisResponse, err := apiGatewayClient.ListApis(context.Background(), listApisRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Api list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, api := range listApisResponse.Items {
		id := *api.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ApiId", id)
	}
	return resourceIds, nil
}

func apiSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if apiResponse, ok := response.Response.(oci_apigateway.GetApiResponse); ok {
		return apiResponse.LifecycleState != oci_apigateway.ApiLifecycleStateDeleted
	}
	return false
}

func apiSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ApiGatewayClient().GetApi(context.Background(), oci_apigateway.GetApiRequest{
		ApiId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
