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
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v52/aianomalydetection"
	"github.com/oracle/oci-go-sdk/v52/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AiPrivateEndpointRequiredOnlyResource = AiPrivateEndpointResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", Required, Create, aiPrivateEndpointRepresentation)

	AiPrivateEndpointResourceConfig = AiPrivateEndpointResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", Optional, Update, aiPrivateEndpointRepresentation)

	aiPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"ai_private_endpoint_id": Representation{RepType: Required, Create: `${oci_ai_anomaly_detection_ai_private_endpoint.test_ai_private_endpoint.id}`},
	}

	aiPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, aiPrivateEndpointDataSourceFilterRepresentation}}
	aiPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_ai_anomaly_detection_ai_private_endpoint.test_ai_private_endpoint.id}`}},
	}

	aiPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"dns_zones":      Representation{RepType: Required, Create: []string{`${oci_core_subnet.test_subnet.subnet_domain_name}`}},
		"subnet_id":      Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      RepresentationGroup{Required, ignoreDefinedTagsChangesRep},
	}

	AiPrivateEndpointResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Optional, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

func TestAiAnomalyDetectionAiPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiAnomalyDetectionAiPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_anomaly_detection_ai_private_endpoint.test_ai_private_endpoint"
	datasourceName := "data.oci_ai_anomaly_detection_ai_private_endpoints.test_ai_private_endpoints"
	singularDatasourceName := "data.oci_ai_anomaly_detection_ai_private_endpoint.test_ai_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+AiPrivateEndpointResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", Optional, Create, aiPrivateEndpointRepresentation), "aianomalydetection", "aiPrivateEndpoint", t)

	ResourceTest(t, testAccCheckAiAnomalyDetectionAiPrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiPrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", Required, Create, aiPrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiPrivateEndpointResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiPrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", Optional, Create, aiPrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				// 					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiPrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", Optional, Create,
					RepresentationCopyWithNewProperties(aiPrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				// 					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AiPrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", Optional, Update, aiPrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				// 					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoints", "test_ai_private_endpoints", Optional, Update, aiPrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + AiPrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", Optional, Update, aiPrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "ai_private_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ai_private_endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", Required, Create, aiPrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiPrivateEndpointResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ai_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "attached_data_assets.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				// 					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AiPrivateEndpointResourceConfig,
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

func testAccCheckAiAnomalyDetectionAiPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).anomalyDetectionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_anomaly_detection_ai_private_endpoint" {
			noResourceFound = false
			request := oci_ai_anomaly_detection.GetAiPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.AiPrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "ai_anomaly_detection")

			response, err := client.GetAiPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("AiAnomalyDetectionAiPrivateEndpoint") {
		resource.AddTestSweepers("AiAnomalyDetectionAiPrivateEndpoint", &resource.Sweeper{
			Name:         "AiAnomalyDetectionAiPrivateEndpoint",
			Dependencies: DependencyGraph["aiPrivateEndpoint"],
			F:            sweepAiAnomalyDetectionAiPrivateEndpointResource,
		})
	}
}

func sweepAiAnomalyDetectionAiPrivateEndpointResource(compartment string) error {
	anomalyDetectionClient := GetTestClients(&schema.ResourceData{}).anomalyDetectionClient()
	aiPrivateEndpointIds, err := getAiPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, aiPrivateEndpointId := range aiPrivateEndpointIds {
		if ok := SweeperDefaultResourceId[aiPrivateEndpointId]; !ok {
			deleteAiPrivateEndpointRequest := oci_ai_anomaly_detection.DeleteAiPrivateEndpointRequest{}

			deleteAiPrivateEndpointRequest.AiPrivateEndpointId = &aiPrivateEndpointId

			deleteAiPrivateEndpointRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "ai_anomaly_detection")
			_, error := anomalyDetectionClient.DeleteAiPrivateEndpoint(context.Background(), deleteAiPrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting AiPrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", aiPrivateEndpointId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &aiPrivateEndpointId, aiPrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				aiPrivateEndpointSweepResponseFetchOperation, "ai_anomaly_detection", true)
		}
	}
	return nil
}

func getAiPrivateEndpointIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "AiPrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	anomalyDetectionClient := GetTestClients(&schema.ResourceData{}).anomalyDetectionClient()

	listAiPrivateEndpointsRequest := oci_ai_anomaly_detection.ListAiPrivateEndpointsRequest{}
	listAiPrivateEndpointsRequest.CompartmentId = &compartmentId
	listAiPrivateEndpointsRequest.LifecycleState = oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateActive
	listAiPrivateEndpointsResponse, err := anomalyDetectionClient.ListAiPrivateEndpoints(context.Background(), listAiPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AiPrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, aiPrivateEndpoint := range listAiPrivateEndpointsResponse.Items {
		id := *aiPrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "AiPrivateEndpointId", id)
	}
	return resourceIds, nil
}

func aiPrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if aiPrivateEndpointResponse, ok := response.Response.(oci_ai_anomaly_detection.GetAiPrivateEndpointResponse); ok {
		return aiPrivateEndpointResponse.LifecycleState != oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateDeleted
	}
	return false
}

func aiPrivateEndpointSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.anomalyDetectionClient().GetAiPrivateEndpoint(context.Background(), oci_ai_anomaly_detection.GetAiPrivateEndpointRequest{
		AiPrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
