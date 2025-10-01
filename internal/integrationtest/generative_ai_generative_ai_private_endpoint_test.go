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
	GenerativeAiGenerativeAiPrivateEndpointRequiredOnlyResource = GenerativeAiGenerativeAiPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoint", "test_generative_ai_private_endpoint", acctest.Required, acctest.Create, GenerativeAiGenerativeAiPrivateEndpointRepresentation)

	GenerativeAiGenerativeAiPrivateEndpointResourceConfig = GenerativeAiGenerativeAiPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoint", "test_generative_ai_private_endpoint", acctest.Optional, acctest.Update, GenerativeAiGenerativeAiPrivateEndpointRepresentation)

	GenerativeAiGenerativeAiPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"generative_ai_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint.id}`},
	}

	GenerativeAiGenerativeAiPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `pe_1234`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiGenerativeAiPrivateEndpointDataSourceFilterRepresentation}}
	GenerativeAiGenerativeAiPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint.id}`}},
	}

	GenerativeAiGenerativeAiPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dns_prefix":     acctest.Representation{RepType: acctest.Required, Create: `dnsPrefix`, Update: `dnsPrefix2`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `pe_1234`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
	}

	GenerativeAiGenerativeAiPrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiGenerativeAiPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiGenerativeAiPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint"
	datasourceName := "data.oci_generative_ai_generative_ai_private_endpoints.test_generative_ai_private_endpoints"
	singularDatasourceName := "data.oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiGenerativeAiPrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoint", "test_generative_ai_private_endpoint", acctest.Optional, acctest.Create, GenerativeAiGenerativeAiPrivateEndpointRepresentation), "generativeai", "generativeAiPrivateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiGenerativeAiPrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiGenerativeAiPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoint", "test_generative_ai_private_endpoint", acctest.Required, acctest.Create, GenerativeAiGenerativeAiPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dns_prefix", "dnsPrefix"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiGenerativeAiPrivateEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiGenerativeAiPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoint", "test_generative_ai_private_endpoint", acctest.Optional, acctest.Create, GenerativeAiGenerativeAiPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe_1234"),
				resource.TestCheckResourceAttr(resourceName, "dns_prefix", "dnsPrefix"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GenerativeAiGenerativeAiPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoint", "test_generative_ai_private_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiGenerativeAiPrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe_1234"),
				resource.TestCheckResourceAttr(resourceName, "dns_prefix", "dnsPrefix"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
			Config: config + compartmentIdVariableStr + GenerativeAiGenerativeAiPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoint", "test_generative_ai_private_endpoint", acctest.Optional, acctest.Update, GenerativeAiGenerativeAiPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_prefix", "dnsPrefix2"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoints", "test_generative_ai_private_endpoints", acctest.Optional, acctest.Update, GenerativeAiGenerativeAiPrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiGenerativeAiPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoint", "test_generative_ai_private_endpoint", acctest.Optional, acctest.Update, GenerativeAiGenerativeAiPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "generative_ai_private_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "generative_ai_private_endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_generative_ai_private_endpoint", "test_generative_ai_private_endpoint", acctest.Required, acctest.Create, GenerativeAiGenerativeAiPrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiGenerativeAiPrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "generative_ai_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fqdn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// verify resource import
		{
			Config:            config + GenerativeAiGenerativeAiPrivateEndpointRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"compartment_id",
				"defined_tags",
				"dns_prefix",
				"lifecycle_details",
				"previous_state",
				"subnet_id",
				"system_tags",
				"time_created",
				"time_updated",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckGenerativeAiGenerativeAiPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_generative_ai_private_endpoint" {
			noResourceFound = false
			request := oci_generative_ai.GetGenerativeAiPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.GenerativeAiPrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetGenerativeAiPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.GenerativeAiPrivateEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiGenerativeAiPrivateEndpoint") {
		resource.AddTestSweepers("GenerativeAiGenerativeAiPrivateEndpoint", &resource.Sweeper{
			Name:         "GenerativeAiGenerativeAiPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["generativeAiPrivateEndpoint"],
			F:            sweepGenerativeAiGenerativeAiPrivateEndpointResource,
		})
	}
}

func sweepGenerativeAiGenerativeAiPrivateEndpointResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	generativeAiPrivateEndpointIds, err := getGenerativeAiGenerativeAiPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, generativeAiPrivateEndpointId := range generativeAiPrivateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[generativeAiPrivateEndpointId]; !ok {
			deleteGenerativeAiPrivateEndpointRequest := oci_generative_ai.DeleteGenerativeAiPrivateEndpointRequest{}

			deleteGenerativeAiPrivateEndpointRequest.GenerativeAiPrivateEndpointId = &generativeAiPrivateEndpointId

			deleteGenerativeAiPrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteGenerativeAiPrivateEndpoint(context.Background(), deleteGenerativeAiPrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting GenerativeAiPrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", generativeAiPrivateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &generativeAiPrivateEndpointId, GenerativeAiGenerativeAiPrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiGenerativeAiPrivateEndpointSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiGenerativeAiPrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "GenerativeAiPrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listGenerativeAiPrivateEndpointsRequest := oci_generative_ai.ListGenerativeAiPrivateEndpointsRequest{}
	listGenerativeAiPrivateEndpointsRequest.CompartmentId = &compartmentId
	listGenerativeAiPrivateEndpointsRequest.LifecycleState = oci_generative_ai.GenerativeAiPrivateEndpointLifecycleStateActive
	listGenerativeAiPrivateEndpointsResponse, err := generativeAiClient.ListGenerativeAiPrivateEndpoints(context.Background(), listGenerativeAiPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting GenerativeAiPrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, generativeAiPrivateEndpoint := range listGenerativeAiPrivateEndpointsResponse.Items {
		id := *generativeAiPrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "GenerativeAiPrivateEndpointId", id)
	}
	return resourceIds, nil
}

func GenerativeAiGenerativeAiPrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if generativeAiPrivateEndpointResponse, ok := response.Response.(oci_generative_ai.GetGenerativeAiPrivateEndpointResponse); ok {
		return generativeAiPrivateEndpointResponse.LifecycleState != oci_generative_ai.GenerativeAiPrivateEndpointLifecycleStateDeleted
	}
	return false
}

func GenerativeAiGenerativeAiPrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetGenerativeAiPrivateEndpoint(context.Background(), oci_generative_ai.GetGenerativeAiPrivateEndpointRequest{
		GenerativeAiPrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
