// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v65/aianomalydetection"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	AiAnomalyDetectionAiPrivateEndpointRequiredOnlyResource = AiAnomalyDetectionAiPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", acctest.Required, acctest.Create, aiPrivateEndpointRepresentation)

	AiAnomalyDetectionAiPrivateEndpointResourceConfig = AiAnomalyDetectionAiPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", acctest.Optional, acctest.Update, aiPrivateEndpointRepresentation)

	AiAnomalyDetectionaiPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"ai_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_anomaly_detection_ai_private_endpoint.test_ai_private_endpoint.id}`},
	}

	AiAnomalyDetectionaiPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: aiPrivateEndpointDataSourceFilterRepresentation}}
	aiPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_anomaly_detection_ai_private_endpoint.test_ai_private_endpoint.id}`}},
	}

	aiPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dns_zones":      acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.test_subnet.subnet_domain_name}`}},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	AiAnomalyDetectionAiPrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

func TestAiAnomalyDetectionAiPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiAnomalyDetectionAiPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_anomaly_detection_ai_private_endpoint.test_ai_private_endpoint"
	datasourceName := "data.oci_ai_anomaly_detection_ai_private_endpoints.test_ai_private_endpoints"
	singularDatasourceName := "data.oci_ai_anomaly_detection_ai_private_endpoint.test_ai_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiAnomalyDetectionAiPrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", acctest.Optional, acctest.Create, aiPrivateEndpointRepresentation), "aianomalydetection", "aiPrivateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckAiAnomalyDetectionAiPrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionAiPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", acctest.Required, acctest.Create, aiPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionAiPrivateEndpointResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionAiPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", acctest.Optional, acctest.Create, aiPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiAnomalyDetectionAiPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(aiPrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionAiPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", acctest.Optional, acctest.Update, aiPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoints", "test_ai_private_endpoints", acctest.Optional, acctest.Update, AiAnomalyDetectionaiPrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionAiPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", acctest.Optional, acctest.Update, aiPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_ai_private_endpoint", "test_ai_private_endpoint", acctest.Required, acctest.Create, AiAnomalyDetectionaiPrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionAiPrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ai_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "attached_data_assets.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + AiAnomalyDetectionAiPrivateEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAiAnomalyDetectionAiPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AnomalyDetectionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_anomaly_detection_ai_private_endpoint" {
			noResourceFound = false
			request := oci_ai_anomaly_detection.GetAiPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.AiPrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_anomaly_detection")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("AiAnomalyDetectionAiPrivateEndpoint") {
		resource.AddTestSweepers("AiAnomalyDetectionAiPrivateEndpoint", &resource.Sweeper{
			Name:         "AiAnomalyDetectionAiPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["aiPrivateEndpoint"],
			F:            sweepAiAnomalyDetectionAiPrivateEndpointResource,
		})
	}
}

func sweepAiAnomalyDetectionAiPrivateEndpointResource(compartment string) error {
	anomalyDetectionClient := acctest.GetTestClients(&schema.ResourceData{}).AnomalyDetectionClient()
	aiPrivateEndpointIds, err := getAiAnomalyDetectionAiPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, aiPrivateEndpointId := range aiPrivateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[aiPrivateEndpointId]; !ok {
			deleteAiPrivateEndpointRequest := oci_ai_anomaly_detection.DeleteAiPrivateEndpointRequest{}

			deleteAiPrivateEndpointRequest.AiPrivateEndpointId = &aiPrivateEndpointId

			deleteAiPrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_anomaly_detection")
			_, error := anomalyDetectionClient.DeleteAiPrivateEndpoint(context.Background(), deleteAiPrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting AiPrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", aiPrivateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &aiPrivateEndpointId, AiAnomalyDetectionaiPrivateEndpointsSweepWaitCondition, time.Duration(3*time.Minute),
				AiAnomalyDetectionaiPrivateEndpointsSweepResponseFetchOperation, "ai_anomaly_detection", true)
		}
	}
	return nil
}

func getAiAnomalyDetectionAiPrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AiPrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	anomalyDetectionClient := acctest.GetTestClients(&schema.ResourceData{}).AnomalyDetectionClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AiPrivateEndpointId", id)
	}
	return resourceIds, nil
}

func AiAnomalyDetectionaiPrivateEndpointsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if aiPrivateEndpointResponse, ok := response.Response.(oci_ai_anomaly_detection.GetAiPrivateEndpointResponse); ok {
		return aiPrivateEndpointResponse.LifecycleState != oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateDeleted
	}
	return false
}

func AiAnomalyDetectionaiPrivateEndpointsSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AnomalyDetectionClient().GetAiPrivateEndpoint(context.Background(), oci_ai_anomaly_detection.GetAiPrivateEndpointRequest{
		AiPrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
