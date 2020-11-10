// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v28/common"
	oci_functions "github.com/oracle/oci-go-sdk/v28/functions"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	FunctionRequiredOnlyResource = FunctionResourceDependencies +
		generateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation)

	FunctionResourceConfig = FunctionResourceDependencies +
		generateResourceFromRepresentationMap("oci_functions_function", "test_function", Optional, Update, functionRepresentation)

	functionSingularDataSourceRepresentation = map[string]interface{}{
		"function_id": Representation{repType: Required, create: `${oci_functions_function.test_function.id}`},
	}

	functionDataSourceRepresentation = map[string]interface{}{
		"application_id": Representation{repType: Required, create: `${oci_functions_application.test_application.id}`},
		"display_name":   Representation{repType: Optional, create: `ExampleFunction`},
		"id":             Representation{repType: Optional, create: `${oci_functions_function.test_function.id}`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, functionDataSourceFilterRepresentation}}
	functionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_functions_function.test_function.id}`}},
	}

	functionRepresentation = map[string]interface{}{
		"application_id":     Representation{repType: Required, create: `${oci_functions_application.test_application.id}`},
		"display_name":       Representation{repType: Required, create: `ExampleFunction`},
		"image":              Representation{repType: Required, create: `${var.image}`, update: `${var.image_for_update}`},
		"memory_in_mbs":      Representation{repType: Required, create: `128`, update: `256`},
		"config":             Representation{repType: Optional, create: map[string]string{"MY_FUNCTION_CONFIG": "ConfVal"}},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"image_digest":       Representation{repType: Optional, create: `${var.image_digest}`, update: `${var.image_digest_for_update}`},
		"timeout_in_seconds": Representation{repType: Optional, create: `30`, update: `31`},
	}

	FunctionResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
		DefinedTagsDependencies
)

func TestFunctionsFunctionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsFunctionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	image := getEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	imageDigest := getEnvSettingWithBlankDefault("image_digest")
	imageDigestVariableStr := fmt.Sprintf("variable \"image_digest\" { default = \"%s\" }\n", imageDigest)

	imageU := getEnvSettingWithBlankDefault("image_for_update")
	imageUVariableStr := fmt.Sprintf("variable \"image_for_update\" { default = \"%s\" }\n", imageU)

	imageDigestU := getEnvSettingWithBlankDefault("image_digest_for_update")
	imageDigestUVariableStr := fmt.Sprintf("variable \"image_digest_for_update\" { default = \"%s\" }\n", imageDigestU)

	resourceName := "oci_functions_function.test_function"
	datasourceName := "data.oci_functions_functions.test_functions"
	singularDatasourceName := "data.oci_functions_function.test_function"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckFunctionsFunctionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "ExampleFunction"),
					resource.TestCheckResourceAttr(resourceName, "image", image),
					resource.TestCheckResourceAttr(resourceName, "memory_in_mbs", "128"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + FunctionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_function", "test_function", Optional, Create, functionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "ExampleFunction"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image", image),
					resource.TestCheckResourceAttr(resourceName, "image_digest", imageDigest),
					resource.TestCheckResourceAttr(resourceName, "memory_in_mbs", "128"),
					resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "30"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + imageUVariableStr + imageDigestUVariableStr + FunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_function", "test_function", Optional, Update, functionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "ExampleFunction"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image", imageU),
					resource.TestCheckResourceAttr(resourceName, "image_digest", imageDigestU),
					resource.TestCheckResourceAttr(resourceName, "memory_in_mbs", "256"),
					resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "31"),

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
					generateDataSourceFromRepresentationMap("oci_functions_functions", "test_functions", Optional, Update, functionDataSourceRepresentation) +
					compartmentIdVariableStr + imageUVariableStr + imageDigestUVariableStr + FunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_function", "test_function", Optional, Update, functionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "application_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "ExampleFunction"),
					//resource.TestCheckResourceAttr(datasourceName, "id", "id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "functions.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "functions.0.application_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "functions.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "functions.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "functions.0.display_name", "ExampleFunction"),
					resource.TestCheckResourceAttr(datasourceName, "functions.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "functions.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "functions.0.image", imageU),
					resource.TestCheckResourceAttr(datasourceName, "functions.0.image_digest", imageDigestU),
					resource.TestCheckResourceAttrSet(datasourceName, "functions.0.invoke_endpoint"),
					resource.TestCheckResourceAttr(datasourceName, "functions.0.memory_in_mbs", "256"),
					resource.TestCheckResourceAttrSet(datasourceName, "functions.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "functions.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "functions.0.time_updated"),
					resource.TestCheckResourceAttr(datasourceName, "functions.0.timeout_in_seconds", "31"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + imageUVariableStr + imageDigestUVariableStr + FunctionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "function_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "config.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "ExampleFunction"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "image", imageU),
					resource.TestCheckResourceAttr(singularDatasourceName, "image_digest", imageDigestU),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "invoke_endpoint"),
					resource.TestCheckResourceAttr(singularDatasourceName, "memory_in_mbs", "256"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "31"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + imageUVariableStr + imageDigestUVariableStr + FunctionResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckFunctionsFunctionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).functionsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_functions_function" {
			noResourceFound = false
			request := oci_functions.GetFunctionRequest{}

			tmp := rs.Primary.ID
			request.FunctionId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "functions")

			response, err := client.GetFunction(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_functions.FunctionLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("FunctionsFunction") {
		resource.AddTestSweepers("FunctionsFunction", &resource.Sweeper{
			Name:         "FunctionsFunction",
			Dependencies: DependencyGraph["function"],
			F:            sweepFunctionsFunctionResource,
		})
	}
}

func sweepFunctionsFunctionResource(compartment string) error {
	functionsManagementClient := GetTestClients(&schema.ResourceData{}).functionsManagementClient()
	functionIds, err := getFunctionIds(compartment)
	if err != nil {
		return err
	}
	for _, functionId := range functionIds {
		if ok := SweeperDefaultResourceId[functionId]; !ok {
			deleteFunctionRequest := oci_functions.DeleteFunctionRequest{}

			deleteFunctionRequest.FunctionId = &functionId

			deleteFunctionRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "functions")
			_, error := functionsManagementClient.DeleteFunction(context.Background(), deleteFunctionRequest)
			if error != nil {
				fmt.Printf("Error deleting Function %s %s, It is possible that the resource is already deleted. Please verify manually \n", functionId, error)
				continue
			}
			waitTillCondition(testAccProvider, &functionId, functionSweepWaitCondition, time.Duration(3*time.Minute),
				functionSweepResponseFetchOperation, "functions", true)
		}
	}
	return nil
}

func getFunctionIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "FunctionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	functionsManagementClient := GetTestClients(&schema.ResourceData{}).functionsManagementClient()

	listFunctionsRequest := oci_functions.ListFunctionsRequest{}

	applicationIds, error := getApplicationIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting applicationId required for Function resource requests \n")
	}
	for _, applicationId := range applicationIds {
		listFunctionsRequest.ApplicationId = &applicationId

		listFunctionsRequest.LifecycleState = oci_functions.FunctionLifecycleStateActive
		listFunctionsResponse, err := functionsManagementClient.ListFunctions(context.Background(), listFunctionsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Function list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, function := range listFunctionsResponse.Items {
			id := *function.Id
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "FunctionId", id)
		}

	}
	return resourceIds, nil
}

func functionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if functionResponse, ok := response.Response.(oci_functions.GetFunctionResponse); ok {
		return functionResponse.LifecycleState != oci_functions.FunctionLifecycleStateDeleted
	}
	return false
}

func functionSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.functionsManagementClient().GetFunction(context.Background(), oci_functions.GetFunctionRequest{
		FunctionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
