// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_apigateway "github.com/oracle/oci-go-sdk/v37/apigateway"
	"github.com/oracle/oci-go-sdk/v37/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ApiRequiredOnlyResource = ApiResourceDependencies +
		generateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Required, Create, apiRepresentation)

	ApiResourceConfig = ApiResourceDependencies +
		generateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Optional, Update, apiRepresentation)

	apiSingularDataSourceRepresentation = map[string]interface{}{
		"api_id": Representation{repType: Required, create: `${oci_apigateway_api.test_api.id}`},
	}

	apiDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, apiDataSourceFilterRepresentation}}
	apiDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_apigateway_api.test_api.id}`}},
	}

	apiRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"content": Representation{repType: Optional,
			create: `{\"openapi\":\"3.0.0\",\"info\":{\"version\":\"1.0.0\",\"title\":\"test\",\"license\":{\"name\":\"MIT\"}},\"paths\":{\"/ping\":{\"get\":{\"responses\":{\"200\":{\"description\":\"OK\"}}}}}}`,
			update: `{\"openapi\":\"3.0.0\",\"info\":{\"version\":\"1.0.0\",\"title\":\"test\"}}`},
	}

	ApiResourceDependencies = DefinedTagsDependencies
)

func TestApigatewayApiResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayApiResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_apigateway_api.test_api"
	datasourceName := "data.oci_apigateway_apis.test_apis"
	singularDatasourceName := "data.oci_apigateway_api.test_api"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ApiResourceDependencies+
		generateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Optional, Create, apiRepresentation), "apigateway", "api", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApigatewayApiDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ApiResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Required, Create, apiRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ApiResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ApiResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Optional, Create, apiRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApiResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Optional, Create,
						representationCopyWithNewProperties(apiRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ApiResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Optional, Update, apiRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_apigateway_apis", "test_apis", Optional, Update, apiDataSourceRepresentation) +
					compartmentIdVariableStr + ApiResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Optional, Update, apiRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_apigateway_api", "test_api", Required, Create, apiSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ApiResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "api_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "specification_type"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ApiResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"content",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckApigatewayApiDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).apiGatewayClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apigateway_api" {
			noResourceFound = false
			request := oci_apigateway.GetApiRequest{}

			tmp := rs.Primary.ID
			request.ApiId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apigateway")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("ApigatewayApi") {
		resource.AddTestSweepers("ApigatewayApi", &resource.Sweeper{
			Name:         "ApigatewayApi",
			Dependencies: DependencyGraph["api"],
			F:            sweepApigatewayApiResource,
		})
	}
}

func sweepApigatewayApiResource(compartment string) error {
	apiGatewayClient := GetTestClients(&schema.ResourceData{}).apiGatewayClient()
	apiIds, err := getApiIds(compartment)
	if err != nil {
		return err
	}
	for _, apiId := range apiIds {
		if ok := SweeperDefaultResourceId[apiId]; !ok {
			deleteApiRequest := oci_apigateway.DeleteApiRequest{}

			deleteApiRequest.ApiId = &apiId

			deleteApiRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apigateway")
			_, error := apiGatewayClient.DeleteApi(context.Background(), deleteApiRequest)
			if error != nil {
				fmt.Printf("Error deleting Api %s %s, It is possible that the resource is already deleted. Please verify manually \n", apiId, error)
				continue
			}
			waitTillCondition(testAccProvider, &apiId, apiSweepWaitCondition, time.Duration(3*time.Minute),
				apiSweepResponseFetchOperation, "apigateway", true)
		}
	}
	return nil
}

func getApiIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ApiId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apiGatewayClient := GetTestClients(&schema.ResourceData{}).apiGatewayClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "ApiId", id)
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

func apiSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.apiGatewayClient().GetApi(context.Background(), oci_apigateway.GetApiRequest{
		ApiId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
