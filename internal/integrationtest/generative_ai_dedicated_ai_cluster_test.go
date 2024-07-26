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
	GenerativeAiDedicatedAiClusterRequiredOnlyResource = GenerativeAiDedicatedAiClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Required, acctest.Create, GenerativeAiHostingDedicatedAiClusterRepresentation)

	GenerativeAiDedicatedAiClusterResourceConfig = GenerativeAiDedicatedAiClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Optional, acctest.Update, GenerativeAiHostingDedicatedAiClusterRepresentation)

	GenerativeAiDedicatedAiClusterSingularDataSourceRepresentation = map[string]interface{}{
		"dedicated_ai_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster.id}`},
	}

	GenerativeAiDedicatedAiClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiDedicatedAiClusterDataSourceFilterRepresentation}}
	GenerativeAiDedicatedAiClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster.id}`}},
	}

	// Hosting cluster, used to test integration with terraform
	GenerativeAiHostingDedicatedAiClusterRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `HOSTING`},
		"unit_count":     acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"unit_shape":     acctest.Representation{RepType: acctest.Required, Create: `SMALL_COHERE`},
		// "defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	// Fine tuning cluster, not used to test terraform, cause many parameters cannot be updated for this.
	// However AI model test will need to refer to this
	GenerativeAiFineTuningDedicatedAiClusterRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `FINE_TUNING`},
		"unit_count":     acctest.Representation{RepType: acctest.Required, Create: `2`},
		"unit_shape":     acctest.Representation{RepType: acctest.Required, Create: `SMALL_COHERE`},
		// "defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	GenerativeAiLoraFineTuningDedicatedAiClusterRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `FINE_TUNING`},
		"unit_count":     acctest.Representation{RepType: acctest.Required, Create: `2`},
		"unit_shape":     acctest.Representation{RepType: acctest.Required, Create: `LARGE_GENERIC`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `llama3testCluster`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	GenerativeAiDedicatedAiClusterResourceDependencies = `` // Cannot test from home region due to GPU, commented out - DefinedTagsDependencies
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiDedicatedAiClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiDedicatedAiClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster"
	datasourceName := "data.oci_generative_ai_dedicated_ai_clusters.test_dedicated_ai_clusters"
	singularDatasourceName := "data.oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiDedicatedAiClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Optional, acctest.Create, GenerativeAiHostingDedicatedAiClusterRepresentation), "generativeai", "dedicatedAiCluster", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiDedicatedAiClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiDedicatedAiClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Required, acctest.Create, GenerativeAiHostingDedicatedAiClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "type", "HOSTING"),
				resource.TestCheckResourceAttr(resourceName, "unit_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "unit_shape", "SMALL_COHERE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiDedicatedAiClusterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiDedicatedAiClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Optional, acctest.Create, GenerativeAiHostingDedicatedAiClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "HOSTING"),
				resource.TestCheckResourceAttr(resourceName, "unit_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "unit_shape", "SMALL_COHERE"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GenerativeAiDedicatedAiClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiHostingDedicatedAiClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "HOSTING"),
				resource.TestCheckResourceAttr(resourceName, "unit_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "unit_shape", "SMALL_COHERE"),

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
			Config: config + compartmentIdVariableStr + GenerativeAiDedicatedAiClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Optional, acctest.Update, GenerativeAiHostingDedicatedAiClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "HOSTING"),
				resource.TestCheckResourceAttr(resourceName, "unit_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "unit_shape", "SMALL_COHERE"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_dedicated_ai_clusters", "test_dedicated_ai_clusters", acctest.Optional, acctest.Update, GenerativeAiDedicatedAiClusterDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiDedicatedAiClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Optional, acctest.Update, GenerativeAiHostingDedicatedAiClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_ai_cluster_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_ai_cluster_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Required, acctest.Create, GenerativeAiDedicatedAiClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiDedicatedAiClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "HOSTING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "unit_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "unit_shape", "SMALL_COHERE"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiDedicatedAiClusterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiDedicatedAiClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_dedicated_ai_cluster" {
			noResourceFound = false
			request := oci_generative_ai.GetDedicatedAiClusterRequest{}

			tmp := rs.Primary.ID
			request.DedicatedAiClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetDedicatedAiCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.DedicatedAiClusterLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiDedicatedAiCluster") {
		resource.AddTestSweepers("GenerativeAiDedicatedAiCluster", &resource.Sweeper{
			Name:         "GenerativeAiDedicatedAiCluster",
			Dependencies: acctest.DependencyGraph["dedicatedAiCluster"],
			F:            sweepGenerativeAiDedicatedAiClusterResource,
		})
	}
}

func sweepGenerativeAiDedicatedAiClusterResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	dedicatedAiClusterIds, err := getGenerativeAiDedicatedAiClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, dedicatedAiClusterId := range dedicatedAiClusterIds {
		if ok := acctest.SweeperDefaultResourceId[dedicatedAiClusterId]; !ok {
			deleteDedicatedAiClusterRequest := oci_generative_ai.DeleteDedicatedAiClusterRequest{}

			deleteDedicatedAiClusterRequest.DedicatedAiClusterId = &dedicatedAiClusterId

			deleteDedicatedAiClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteDedicatedAiCluster(context.Background(), deleteDedicatedAiClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting DedicatedAiCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", dedicatedAiClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dedicatedAiClusterId, GenerativeAiDedicatedAiClusterSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiDedicatedAiClusterSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiDedicatedAiClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DedicatedAiClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listDedicatedAiClustersRequest := oci_generative_ai.ListDedicatedAiClustersRequest{}
	listDedicatedAiClustersRequest.CompartmentId = &compartmentId
	listDedicatedAiClustersRequest.LifecycleState = oci_generative_ai.DedicatedAiClusterLifecycleStateActive
	listDedicatedAiClustersResponse, err := generativeAiClient.ListDedicatedAiClusters(context.Background(), listDedicatedAiClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DedicatedAiCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dedicatedAiCluster := range listDedicatedAiClustersResponse.Items {
		id := *dedicatedAiCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DedicatedAiClusterId", id)
	}
	return resourceIds, nil
}

func GenerativeAiDedicatedAiClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dedicatedAiClusterResponse, ok := response.Response.(oci_generative_ai.GetDedicatedAiClusterResponse); ok {
		return dedicatedAiClusterResponse.LifecycleState != oci_generative_ai.DedicatedAiClusterLifecycleStateDeleted
	}
	return false
}

func GenerativeAiDedicatedAiClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetDedicatedAiCluster(context.Background(), oci_generative_ai.GetDedicatedAiClusterRequest{
		DedicatedAiClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
