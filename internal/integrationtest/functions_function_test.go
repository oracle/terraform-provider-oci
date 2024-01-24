// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

const pbfFunctionDisplayName = "ExampleFunctionFromPbf"

var (
	// Uncomment this when running test in a Tenancy created after 2019 as they will have Oracle internal tags
	//for Creator and Created time added automatically
	/*fnDefinedTagsIgnoreRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}*/
	FunctionsFunctionRequiredOnlyResource = FunctionsFunctionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionImageSourceRepresentation)

	FunctionsFunctionResourceConfig = FunctionsFunctionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Update, FunctionsFunctionImageSourceRepresentation)

	FunctionsFunctionsFunctionSingularDataSourceRepresentation = map[string]interface{}{
		"function_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_function.test_function.id}`},
	}

	FunctionsFunctionsFunctionDataSourceRepresentation = map[string]interface{}{
		"application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_application.test_application.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `ExampleFunction`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_functions_function.test_function.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FunctionsFunctionDataSourceFilterRepresentation}}
	FunctionsFunctionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_functions_function.test_function.id}`}},
	}

	FunctionsFunctionRepresentation = map[string]interface{}{
		"application_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_application.test_application.id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `ExampleFunction`},
		"memory_in_mbs":                  acctest.Representation{RepType: acctest.Required, Create: `128`, Update: `256`},
		"config":                         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"MY_FUNCTION_CONFIG": "ConfVal"}},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"image":                          acctest.Representation{RepType: acctest.Required, Create: `${var.image}`, Update: `${var.image_for_update}`},
		"image_digest":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.image_digest}`, Update: `${var.image_digest_for_update}`},
		"provisioned_concurrency_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FunctionsFunctionProvisionedConcurrencyConfigRepresentation},
		"source_details":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: FunctionsFunctionSourceDetailsRepresentation},
		"timeout_in_seconds":             acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `31`},
		"trace_config":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: FunctionsFunctionTraceConfigRepresentation},
		//"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: fnDefinedTagsIgnoreRepresentation},
	}

	FunctionsFunctionImageSourceRepresentation            = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"source_details"}, FunctionsFunctionRepresentation)
	FunctionsFunctionPbfSourceRepresentationWithoutConfig = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"image", "image_digest", "config", "display_name"},
		FunctionsFunctionRepresentation)
	FunctionsFunctionPbfSourceRepresentation = acctest.RepresentationCopyWithNewProperties(FunctionsFunctionPbfSourceRepresentationWithoutConfig, map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: pbfFunctionDisplayName},
		"config":       acctest.Representation{RepType: acctest.Required, Create: map[string]string{`${var.function_config_key}`: `${var.function_config_value}`}},
	})

	FunctionsFunctionProvisionedConcurrencyConfigRepresentation = map[string]interface{}{
		"strategy": acctest.Representation{RepType: acctest.Required, Create: `CONSTANT`, Update: `NONE`},
		"count":    acctest.Representation{RepType: acctest.Optional, Create: `40`, Update: `0`},
	}
	FunctionsFunctionSourceDetailsRepresentation = map[string]interface{}{
		"pbf_listing_id": acctest.Representation{RepType: acctest.Required, Create: `${var.pbf_listing_id}`},
		"source_type":    acctest.Representation{RepType: acctest.Required, Create: `PRE_BUILT_FUNCTIONS`},
	}
	FunctionsFunctionTraceConfigRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	functionApplicationDisplayName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	FunctionsFunctionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
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

	pbfListingId := utils.GetEnvSettingWithBlankDefault("pbf_listing_id")
	pbfListingIdVariableStr := fmt.Sprintf("variable \"pbf_listing_id\" { default = \"%s\" }\n", pbfListingId)

	functionConfigKey := utils.GetEnvSettingWithDefault("function_config_key", "MY_FUNCTION_CONFIG")
	functionConfigKeyVariableStr := fmt.Sprintf("variable \"function_config_key\" { default = \"%s\" }\n", functionConfigKey)

	functionConfigValue := utils.GetEnvSettingWithDefault("function_config_value", "ConfVal")
	functionConfigValueVariableStr := fmt.Sprintf("variable \"function_config_value\" { default = \"%s\" }\n", functionConfigValue)

	resourceName := "oci_functions_function.test_function"
	datasourceName := "data.oci_functions_functions.test_functions"
	singularDatasourceName := "data.oci_functions_function.test_function"

	log.Printf("imageDigest value is: %s", imageDigest)
	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FunctionsFunctionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Create, FunctionsFunctionRepresentation), "functions", "function", t)

	acctest.ResourceTest(t, testAccCheckFunctionsFunctionDestroy, []resource.TestStep{
		// verify Create with PBF
		{
			Config: config + compartmentIdVariableStr + pbfListingIdVariableStr + functionConfigKeyVariableStr + functionConfigValueVariableStr + FunctionsFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionPbfSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", pbfFunctionDisplayName),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.pbf_listing_id", pbfListingId),
				resource.TestCheckResourceAttr(resourceName, "memory_in_mbs", "128"),
				resource.TestCheckResourceAttr(resourceName, "config.%", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FunctionsFunctionResourceDependencies,
		},
		// verify Create with PBF with optionals
		{
			Config: config + compartmentIdVariableStr + pbfListingIdVariableStr + functionConfigKeyVariableStr + functionConfigValueVariableStr + FunctionsFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Create, FunctionsFunctionPbfSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", pbfFunctionDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.pbf_listing_id", pbfListingId),
				resource.TestCheckResourceAttr(resourceName, "memory_in_mbs", "128"),
				resource.TestCheckResourceAttr(resourceName, "provisioned_concurrency_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "provisioned_concurrency_config.0.count", "40"),
				resource.TestCheckResourceAttr(resourceName, "provisioned_concurrency_config.0.strategy", "CONSTANT"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.pbf_listing_id"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "PRE_BUILT_FUNCTIONS"),
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
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FunctionsFunctionResourceDependencies,
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionImageSourceRepresentation),
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
			Config: config + compartmentIdVariableStr + FunctionsFunctionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Create, FunctionsFunctionImageSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ExampleFunction"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image", image),
				resource.TestCheckResourceAttr(resourceName, "image_digest", imageDigest),
				resource.TestCheckResourceAttr(resourceName, "memory_in_mbs", "128"),
				resource.TestCheckResourceAttr(resourceName, "provisioned_concurrency_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "provisioned_concurrency_config.0.count", "40"),
				resource.TestCheckResourceAttr(resourceName, "provisioned_concurrency_config.0.strategy", "CONSTANT"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "trace_config.0.is_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					log.Printf("enable_export_compartment : %s", utils.GetEnvSettingWithDefault("enable_export_compartment", "true"))
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
			Config: config + compartmentIdVariableStr + imageUVariableStr + imageDigestUVariableStr + FunctionsFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Update, FunctionsFunctionImageSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "config.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ExampleFunction"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image", imageU),
				resource.TestCheckResourceAttr(resourceName, "image_digest", imageDigestU),
				resource.TestCheckResourceAttr(resourceName, "memory_in_mbs", "256"),
				resource.TestCheckResourceAttr(resourceName, "provisioned_concurrency_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "provisioned_concurrency_config.0.count", "0"),
				resource.TestCheckResourceAttr(resourceName, "provisioned_concurrency_config.0.strategy", "NONE"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions", "test_functions", acctest.Optional, acctest.Update, FunctionsFunctionsFunctionDataSourceRepresentation) +
				compartmentIdVariableStr + imageUVariableStr + imageDigestUVariableStr + FunctionsFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Optional, acctest.Update, FunctionsFunctionImageSourceRepresentation),
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
				resource.TestCheckResourceAttr(datasourceName, "functions.0.provisioned_concurrency_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "functions.0.provisioned_concurrency_config.0.count", "0"),
				resource.TestCheckResourceAttr(datasourceName, "functions.0.provisioned_concurrency_config.0.strategy", "NONE"),
				resource.TestCheckResourceAttr(datasourceName, "functions.0.shape", "GENERIC_X86"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionsFunctionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + imageUVariableStr + imageDigestUVariableStr + FunctionsFunctionResourceConfig,
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
				resource.TestCheckResourceAttr(singularDatasourceName, "provisioned_concurrency_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "provisioned_concurrency_config.0.count", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "provisioned_concurrency_config.0.strategy", "NONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape", "GENERIC_X86"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "31"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_config.0.is_enabled", "true"),
			),
		},
		// verify resource import
		{
			Config:                  config + FunctionsFunctionRequiredOnlyResource,
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
	functionIds, err := getFunctionsFunctionIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &functionId, FunctionsFunctionSweepWaitCondition, time.Duration(3*time.Minute),
				FunctionsFunctionSweepResponseFetchOperation, "functions", true)
		}
	}
	return nil
}

func getFunctionsFunctionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FunctionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	functionsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FunctionsManagementClient()

	listFunctionsRequest := oci_functions.ListFunctionsRequest{}

	applicationIds, error := getFunctionsApplicationIds(compartment)
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

func FunctionsFunctionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if functionResponse, ok := response.Response.(oci_functions.GetFunctionResponse); ok {
		return functionResponse.LifecycleState != oci_functions.FunctionLifecycleStateDeleted
	}
	return false
}

func FunctionsFunctionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FunctionsManagementClient().GetFunction(context.Background(), oci_functions.GetFunctionRequest{
		FunctionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
