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
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiAgentProvisionedCapacityRequiredOnlyResource = GenerativeAiAgentProvisionedCapacityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Required, acctest.Create, GenerativeAiAgentProvisionedCapacityRepresentation)

	GenerativeAiAgentProvisionedCapacityResourceConfig = GenerativeAiAgentProvisionedCapacityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Optional, acctest.Update, GenerativeAiAgentProvisionedCapacityRepresentation)

	GenerativeAiAgentProvisionedCapacitySingularDataSourceRepresentation = map[string]interface{}{
		"provisioned_capacity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_agent_provisioned_capacity.test_provisioned_capacity.id}`},
	}

	GenerativeAiAgentProvisionedCapacityDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `provisionedCapacity display name`, Update: `displayName2`},
		"provisioned_capacity_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_agent_provisioned_capacity.test_provisioned_capacity.id}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentProvisionedCapacityDataSourceFilterRepresentation}}
	GenerativeAiAgentProvisionedCapacityDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_agent_provisioned_capacity.test_provisioned_capacity.id}`}},
	}

	GenerativeAiAgentProvisionedCapacityRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `provisionedCapacity display name`, Update: `displayName2`},
		"number_of_units": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	GenerativeAiAgentProvisionedCapacityResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Required, acctest.Create, GenerativeAiAgentProvisionedCapacityRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: generative_ai_agent/default
func TestGenerativeAiAgentProvisionedCapacityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentProvisionedCapacityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_agent_provisioned_capacity.test_provisioned_capacity"
	datasourceName := "data.oci_generative_ai_agent_provisioned_capacities.test_provisioned_capacities"
	singularDatasourceName := "data.oci_generative_ai_agent_provisioned_capacity.test_provisioned_capacity"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiAgentProvisionedCapacityResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Optional, acctest.Create, GenerativeAiAgentProvisionedCapacityRepresentation), "generativeaiagent", "provisionedCapacity", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentProvisionedCapacityDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiAgentProvisionedCapacityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Required, acctest.Create, GenerativeAiAgentProvisionedCapacityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "provisionedCapacity display name"),
				resource.TestCheckResourceAttr(resourceName, "number_of_units", "10"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiAgentProvisionedCapacityResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiAgentProvisionedCapacityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Optional, acctest.Create, GenerativeAiAgentProvisionedCapacityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "provisionedCapacity display name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "number_of_units", "10"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GenerativeAiAgentProvisionedCapacityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiAgentProvisionedCapacityRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "provisionedCapacity display name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "number_of_units", "10"),
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
			Config: config + compartmentIdVariableStr + GenerativeAiAgentProvisionedCapacityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Optional, acctest.Update, GenerativeAiAgentProvisionedCapacityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "number_of_units", "11"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacities", "test_provisioned_capacities", acctest.Optional, acctest.Update, GenerativeAiAgentProvisionedCapacityDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiAgentProvisionedCapacityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Optional, acctest.Update, GenerativeAiAgentProvisionedCapacityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "provisioned_capacity_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "provisioned_capacity_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "provisioned_capacity_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Required, acctest.Create, GenerativeAiAgentProvisionedCapacitySingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiAgentProvisionedCapacityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "provisioned_capacity_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "number_of_units", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiAgentProvisionedCapacityRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiAgentProvisionedCapacityDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_agent_provisioned_capacity" {
			noResourceFound = false
			request := oci_generative_ai_agent.GetProvisionedCapacityRequest{}

			tmp := rs.Primary.ID
			request.ProvisionedCapacityId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")

			response, err := client.GetProvisionedCapacity(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai_agent.ProvisionedCapacityLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiAgentProvisionedCapacity") {
		resource.AddTestSweepers("GenerativeAiAgentProvisionedCapacity", &resource.Sweeper{
			Name:         "GenerativeAiAgentProvisionedCapacity",
			Dependencies: acctest.DependencyGraph["provisionedCapacity"],
			F:            sweepGenerativeAiAgentProvisionedCapacityResource,
		})
	}
}

func sweepGenerativeAiAgentProvisionedCapacityResource(compartment string) error {
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()
	provisionedCapacityIds, err := getGenerativeAiAgentProvisionedCapacityIds(compartment)
	if err != nil {
		return err
	}
	for _, provisionedCapacityId := range provisionedCapacityIds {
		if ok := acctest.SweeperDefaultResourceId[provisionedCapacityId]; !ok {
			deleteProvisionedCapacityRequest := oci_generative_ai_agent.DeleteProvisionedCapacityRequest{}

			deleteProvisionedCapacityRequest.ProvisionedCapacityId = &provisionedCapacityId

			deleteProvisionedCapacityRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")
			_, error := generativeAiAgentClient.DeleteProvisionedCapacity(context.Background(), deleteProvisionedCapacityRequest)
			if error != nil {
				fmt.Printf("Error deleting ProvisionedCapacity %s %s, It is possible that the resource is already deleted. Please verify manually \n", provisionedCapacityId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &provisionedCapacityId, GenerativeAiAgentProvisionedCapacitySweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiAgentProvisionedCapacitySweepResponseFetchOperation, "generative_ai_agent", true)
		}
	}
	return nil
}

func getGenerativeAiAgentProvisionedCapacityIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProvisionedCapacityId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()

	listProvisionedCapacitiesRequest := oci_generative_ai_agent.ListProvisionedCapacitiesRequest{}
	listProvisionedCapacitiesRequest.CompartmentId = &compartmentId
	listProvisionedCapacitiesRequest.LifecycleState = oci_generative_ai_agent.ProvisionedCapacityLifecycleStateActive
	listProvisionedCapacitiesResponse, err := generativeAiAgentClient.ListProvisionedCapacities(context.Background(), listProvisionedCapacitiesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ProvisionedCapacity list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, provisionedCapacity := range listProvisionedCapacitiesResponse.Items {
		id := *provisionedCapacity.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProvisionedCapacityId", id)
	}
	return resourceIds, nil
}

func GenerativeAiAgentProvisionedCapacitySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if provisionedCapacityResponse, ok := response.Response.(oci_generative_ai_agent.GetProvisionedCapacityResponse); ok {
		return provisionedCapacityResponse.LifecycleState != oci_generative_ai_agent.ProvisionedCapacityLifecycleStateDeleted
	}
	return false
}

func GenerativeAiAgentProvisionedCapacitySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiAgentClient().GetProvisionedCapacity(context.Background(), oci_generative_ai_agent.GetProvisionedCapacityRequest{
		ProvisionedCapacityId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
