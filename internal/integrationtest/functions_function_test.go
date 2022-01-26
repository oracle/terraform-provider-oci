// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_functions "github.com/oracle/oci-go-sdk/v56/functions"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	FunctionRequiredOnlyResource = FunctionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionRepresentation)

	FunctionResourceConfig = FunctionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Update, functionRepresentation)

	functionSingularDataSourceRepresentation = map[string]interface{}{
		"function_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_function.test_function.id}`},
	}

	functionDataSourceRepresentation = map[string]interface{}{
		"application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_application.test_application.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `ExampleFunction`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_functions_function.test_function.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: functionDataSourceFilterRepresentation}}
	functionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_functions_function.test_function.id}`}},
	}

	functionRepresentation = map[string]interface{}{
		"application_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_application.test_application.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `ExampleFunction`},
		"image":              acctest.Representation{RepType: acctest.Required, Create: `${var.image}`, Update: `${var.image_for_update}`},
		"memory_in_mbs":      acctest.Representation{RepType: acctest.Required, Create: `128`, Update: `256`},
		"config":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"MY_FUNCTION_CONFIG": "ConfVal"}},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"image_digest":       acctest.Representation{RepType: acctest.Optional, Create: `${var.image_digest}`, Update: `${var.image_digest_for_update}`},
		"timeout_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `31`},
		"trace_config":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: functionTraceConfigRepresentation},
	}
	functionTraceConfigRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	functionApplicationDisplayName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	FunctionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationRepresentation) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig
)

// issue-routing-tag: functions/default
func TestFunctionsFunctionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsFunctionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	image := utils.GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	imageDigest := utils.GetEnvSettingWithBlankDefault("image_digest")
	imageDigestVariableStr := fmt.Sprintf("variable \"image_digest\" { default = \"%s\" }\n", imageDigest)

	imageU := utils.GetEnvSettingWithBlankDefault("image_for_update")
	imageUVariableStr := fmt.Sprintf("variable \"image_for_update\" { default = \"%s\" }\n", imageU)

	imageDigestU := utils.GetEnvSettingWithBlankDefault("image_digest_for_update")
	imageDigestUVariableStr := fmt.Sprintf("variable \"image_digest_for_update\" { default = \"%s\" }\n", imageDigestU)

	resourceName := "oci_functions_function.test_function"
	datasourceName := "data.oci_functions_functions.test_functions"
	singularDatasourceName := "data.oci_functions_function.test_function"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FunctionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Create, functionRepresentation), "functions", "function", t)

	acctest.ResourceTest(t, testAccCheckFunctionsFunctionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ExampleFunction"),
				resource.TestCheckResourceAttr(resourceName, "image", image),
				resource.TestCheckResourceAttr(resourceName, "memory_in_mbs", "128"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FunctionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Create, functionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ExampleFunction"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image", image),
				resource.TestCheckResourceAttr(resourceName, "image_digest", imageDigest),
				resource.TestCheckResourceAttr(resourceName, "memory_in_mbs", "128"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.0.is_enabled", "false"),

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
			Config: config + compartmentIdVariableStr + imageUVariableStr + imageDigestUVariableStr + FunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Update, functionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ExampleFunction"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image", imageU),
				resource.TestCheckResourceAttr(resourceName, "image_digest", imageDigestU),
				resource.TestCheckResourceAttr(resourceName, "memory_in_mbs", "256"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "31"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.0.is_enabled", "true"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions", "test_functions", acctest.Optional, acctest.Update, functionDataSourceRepresentation) +
				compartmentIdVariableStr + imageUVariableStr + imageDigestUVariableStr + FunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Update, functionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "application_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "ExampleFunction"),
				//resource.TestCheckResourceAttr(datasourceName, "id", "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "functions.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions.0.application_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions.0.compartment_id"),
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
				resource.TestCheckResourceAttr(datasourceName, "functions.0.trace_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "functions.0.trace_config.0.is_enabled", "true"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + imageUVariableStr + imageDigestUVariableStr + FunctionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "function_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.%", "1"),
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
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_config.0.is_enabled", "true"),
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
	})
}

func testAccCheckFunctionsFunctionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FunctionsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_functions_function" {
			noResourceFound = false
			request := oci_functions.GetFunctionRequest{}

			tmp := rs.Primary.ID
			request.FunctionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "functions")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("FunctionsFunction") {
		resource.AddTestSweepers("FunctionsFunction", &resource.Sweeper{
			Name:         "FunctionsFunction",
			Dependencies: acctest.DependencyGraph["function"],
			F:            sweepFunctionsFunctionResource,
		})
	}
}

func sweepFunctionsFunctionResource(compartment string) error {
	functionsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FunctionsManagementClient()
	functionIds, err := getFunctionIds(compartment)
	if err != nil {
		return err
	}
	for _, functionId := range functionIds {
		if ok := acctest.SweeperDefaultResourceId[functionId]; !ok {
			deleteFunctionRequest := oci_functions.DeleteFunctionRequest{}

			deleteFunctionRequest.FunctionId = &functionId

			deleteFunctionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "functions")
			_, error := functionsManagementClient.DeleteFunction(context.Background(), deleteFunctionRequest)
			if error != nil {
				fmt.Printf("Error deleting Function %s %s, It is possible that the resource is already deleted. Please verify manually \n", functionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &functionId, functionSweepWaitCondition, time.Duration(3*time.Minute),
				functionSweepResponseFetchOperation, "functions", true)
		}
	}
	return nil
}

func getFunctionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FunctionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	functionsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FunctionsManagementClient()

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
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FunctionId", id)
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

func functionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FunctionsManagementClient().GetFunction(context.Background(), oci_functions.GetFunctionRequest{
		FunctionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
