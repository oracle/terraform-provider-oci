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
	GenerativeAiHostedApplicationIamRequiredOnlyResource = GenerativeAiHostedApplicationIamResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_iam", "test_hosted_application_iam", acctest.Required, acctest.Create, GenerativeAiHostedApplicationIamRepresentation)

	GenerativeAiHostedApplicationIamResourceConfig = GenerativeAiHostedApplicationIamResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_iam", "test_hosted_application_iam", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationIamRepresentation)

	GenerativeAiHostedApplicationIamSingularDataSourceRepresentation = map[string]interface{}{
		"hosted_application_iam_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_hosted_application_iam.test_hosted_application_iam.id}`},
	}

	GenerativeAiHostedApplicationIamDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `qa_plus_agent_test_hosted_application_iam`, Update: `qa_plus_agent_test_hosted_application_iam_updated`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_hosted_application_iam.test_hosted_application_iam.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiHostedApplicationIamDataSourceFilterRepresentation}}
	GenerativeAiHostedApplicationIamDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_hosted_application_iam.test_hosted_application_iam.id}`}},
	}

	GenerativeAiHostedApplicationIamRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `qa_plus_agent_test_hosted_application_iam`, Update: `qa_plus_agent_test_hosted_application_iam_updated`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"environment_variables": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiHostedApplicationIamEnvironmentVariablesRepresentation},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Finance"}},
		"networking_config":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiHostedApplicationIamNetworkingConfigRepresentation},
		"scaling_config":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiHostedApplicationIamScalingConfigRepresentation},
	}
	GenerativeAiHostedApplicationIamEnvironmentVariablesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `PLAINTEXT`, Update: `PLAINTEXT`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `{\"dummyKey\":\"dummyValue\"}`},
	}
	GenerativeAiHostedApplicationIamNetworkingConfigRepresentation = map[string]interface{}{
		"inbound_networking_config":  acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiHostedApplicationIamNetworkingConfigInboundNetworkingConfigRepresentation},
		"outbound_networking_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiHostedApplicationIamNetworkingConfigOutboundNetworkingConfigRepresentation},
	}
	GenerativeAiHostedApplicationIamScalingConfigRepresentation = map[string]interface{}{
		"scaling_type":                 acctest.Representation{RepType: acctest.Required, Create: `CPU`, Update: `CPU`},
		"max_replica":                  acctest.Representation{RepType: acctest.Optional, Create: `3`, Update: `3`},
		"min_replica":                  acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"target_concurrency_threshold": acctest.Representation{RepType: acctest.Optional},
		"target_cpu_threshold":         acctest.Representation{RepType: acctest.Optional, Create: `70`, Update: `70`},
		"target_memory_threshold":      acctest.Representation{RepType: acctest.Optional},
		"target_rps_threshold":         acctest.Representation{RepType: acctest.Optional},
	}
	GenerativeAiHostedApplicationIamNetworkingConfigInboundNetworkingConfigRepresentation = map[string]interface{}{
		"endpoint_mode": acctest.Representation{RepType: acctest.Required, Create: `PUBLIC`},
	}
	GenerativeAiHostedApplicationIamNetworkingConfigOutboundNetworkingConfigRepresentation = map[string]interface{}{
		"network_mode": acctest.Representation{RepType: acctest.Required, Create: `MANAGED`},
	}

	GenerativeAiHostedApplicationIamResourceDependencies = ""
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiHostedApplicationIamResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiHostedApplicationIamResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_hosted_application_iam.test_hosted_application_iam"
	datasourceName := "data.oci_generative_ai_hosted_application_iams.test_hosted_application_iams"
	singularDatasourceName := "data.oci_generative_ai_hosted_application_iam.test_hosted_application_iam"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiHostedApplicationIamResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_iam", "test_hosted_application_iam", acctest.Optional, acctest.Create, GenerativeAiHostedApplicationIamRepresentation), "generativeai", "hostedApplicationIam", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiHostedApplicationIamDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationIamResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_iam", "test_hosted_application_iam", acctest.Required, acctest.Create, GenerativeAiHostedApplicationIamRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "qa_plus_agent_test_hosted_application_iam"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationIamResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationIamResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_iam", "test_hosted_application_iam", acctest.Optional, acctest.Create, GenerativeAiHostedApplicationIamRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "qa_plus_agent_test_hosted_application_iam"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.type", "PLAINTEXT"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.value", "{\"dummyKey\":\"dummyValue\"}"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.0.endpoint_mode", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.0.network_mode", "MANAGED"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.max_replica", "3"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.min_replica", "1"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.scaling_type", "CPU"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_concurrency_threshold", "0"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_cpu_threshold", "70"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_memory_threshold", "0"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_rps_threshold", "0"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GenerativeAiHostedApplicationIamResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_iam", "test_hosted_application_iam", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiHostedApplicationIamRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "qa_plus_agent_test_hosted_application_iam"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.type", "PLAINTEXT"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.value", "{\"dummyKey\":\"dummyValue\"}"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.0.endpoint_mode", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.0.network_mode", "MANAGED"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.max_replica", "3"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.min_replica", "1"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.scaling_type", "CPU"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_concurrency_threshold", "0"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_cpu_threshold", "70"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_memory_threshold", "0"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_rps_threshold", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationIamResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_iam", "test_hosted_application_iam", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationIamRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "qa_plus_agent_test_hosted_application_iam_updated"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.type", "PLAINTEXT"),
				resource.TestCheckResourceAttr(resourceName, "environment_variables.0.value", "{\"dummyKey\":\"dummyValue\"}"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.inbound_networking_config.0.endpoint_mode", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "networking_config.0.outbound_networking_config.0.network_mode", "MANAGED"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.max_replica", "3"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.min_replica", "1"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.scaling_type", "CPU"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_concurrency_threshold", "0"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_cpu_threshold", "70"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_memory_threshold", "0"),
				resource.TestCheckResourceAttr(resourceName, "scaling_config.0.target_rps_threshold", "0"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_hosted_application_iams", "test_hosted_application_iams", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationIamDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiHostedApplicationIamResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_iam", "test_hosted_application_iam", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationIamRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "qa_plus_agent_test_hosted_application_iam_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "hosted_application_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "hosted_application_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_hosted_application_iam", "test_hosted_application_iam", acctest.Required, acctest.Create, GenerativeAiHostedApplicationIamSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiHostedApplicationIamResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hosted_application_iam_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "qa_plus_agent_test_hosted_application_iam_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "environment_variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "environment_variables.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "environment_variables.0.type", "PLAINTEXT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "environment_variables.0.value", "{\"dummyKey\":\"dummyValue\"}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networking_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networking_config.0.inbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networking_config.0.inbound_networking_config.0.endpoint_mode", "PUBLIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networking_config.0.outbound_networking_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networking_config.0.outbound_networking_config.0.network_mode", "MANAGED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.max_replica", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.min_replica", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.scaling_type", "CPU"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.target_concurrency_threshold", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.target_cpu_threshold", "70"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.target_memory_threshold", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scaling_config.0.target_rps_threshold", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiHostedApplicationIamRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiHostedApplicationIamDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_hosted_application_iam" {
			noResourceFound = false
			request := oci_generative_ai.GetHostedApplicationIamRequest{}

			tmp := rs.Primary.ID
			request.HostedApplicationIamId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetHostedApplicationIam(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.HostedApplicationIamLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiHostedApplicationIam") {
		resource.AddTestSweepers("GenerativeAiHostedApplicationIam", &resource.Sweeper{
			Name:         "GenerativeAiHostedApplicationIam",
			Dependencies: acctest.DependencyGraph["hostedApplicationIam"],
			F:            sweepGenerativeAiHostedApplicationIamResource,
		})
	}
}

func sweepGenerativeAiHostedApplicationIamResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	hostedApplicationIamIds, err := getGenerativeAiHostedApplicationIamIds(compartment)
	if err != nil {
		return err
	}
	for _, hostedApplicationIamId := range hostedApplicationIamIds {
		if ok := acctest.SweeperDefaultResourceId[hostedApplicationIamId]; !ok {
			deleteHostedApplicationIamRequest := oci_generative_ai.DeleteHostedApplicationIamRequest{}

			deleteHostedApplicationIamRequest.HostedApplicationIamId = &hostedApplicationIamId

			deleteHostedApplicationIamRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteHostedApplicationIam(context.Background(), deleteHostedApplicationIamRequest)
			if error != nil {
				fmt.Printf("Error deleting HostedApplicationIam %s %s, It is possible that the resource is already deleted. Please verify manually \n", hostedApplicationIamId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &hostedApplicationIamId, GenerativeAiHostedApplicationIamSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiHostedApplicationIamSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiHostedApplicationIamIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "HostedApplicationIamId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listHostedApplicationsIamRequest := oci_generative_ai.ListHostedApplicationsIamRequest{}
	listHostedApplicationsIamRequest.CompartmentId = &compartmentId
	listHostedApplicationsIamRequest.LifecycleState = oci_generative_ai.HostedApplicationBaseLifecycleStateActive
	listHostedApplicationsIamResponse, err := generativeAiClient.ListHostedApplicationsIam(context.Background(), listHostedApplicationsIamRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting HostedApplicationIam list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, hostedApplicationIam := range listHostedApplicationsIamResponse.Items {
		id := *hostedApplicationIam.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "HostedApplicationIamId", id)
	}
	return resourceIds, nil
}

func GenerativeAiHostedApplicationIamSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if hostedApplicationIamResponse, ok := response.Response.(oci_generative_ai.GetHostedApplicationIamResponse); ok {
		return hostedApplicationIamResponse.LifecycleState != oci_generative_ai.HostedApplicationIamLifecycleStateDeleted
	}
	return false
}

func GenerativeAiHostedApplicationIamSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetHostedApplicationIam(context.Background(), oci_generative_ai.GetHostedApplicationIamRequest{
		HostedApplicationIamId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
