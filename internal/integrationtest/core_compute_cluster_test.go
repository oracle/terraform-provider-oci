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
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreComputeClusterRequiredOnlyResource = CoreComputeClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_cluster", "test_compute_cluster", acctest.Required, acctest.Create, CoreComputeClusterRepresentation)

	CoreComputeClusterResourceConfig = CoreComputeClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_cluster", "test_compute_cluster", acctest.Optional, acctest.Update, CoreComputeClusterRepresentation)

	CoreComputeClusterSingularDataSourceRepresentation = map[string]interface{}{
		"compute_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_cluster.test_compute_cluster.id}`},
	}

	CoreComputeClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreComputeClusterDataSourceFilterRepresentation}}
	CoreComputeClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_compute_cluster.test_compute_cluster.id}`}},
	}

	CoreComputeClusterRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CoreComputeClusterResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/computeBm
func TestCoreComputeClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_compute_cluster.test_compute_cluster"
	datasourceName := "data.oci_core_compute_clusters.test_compute_clusters"
	singularDatasourceName := "data.oci_core_compute_cluster.test_compute_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreComputeClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_cluster", "test_compute_cluster", acctest.Optional, acctest.Create, CoreComputeClusterRepresentation), "core", "computeCluster", t)

	acctest.ResourceTest(t, testAccCheckCoreComputeClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreComputeClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_cluster", "test_compute_cluster", acctest.Required, acctest.Create, CoreComputeClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreComputeClusterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreComputeClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_cluster", "test_compute_cluster", acctest.Optional, acctest.Create, CoreComputeClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreComputeClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_cluster", "test_compute_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreComputeClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + CoreComputeClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_cluster", "test_compute_cluster", acctest.Optional, acctest.Update, CoreComputeClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_clusters", "test_compute_clusters", acctest.Optional, acctest.Update, CoreComputeClusterDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_cluster", "test_compute_cluster", acctest.Optional, acctest.Update, CoreComputeClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "compute_cluster_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compute_cluster_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_cluster", "test_compute_cluster", acctest.Required, acctest.Create, CoreComputeClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreComputeClusterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreComputeClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_compute_cluster" {
			noResourceFound = false
			request := oci_core.GetComputeClusterRequest{}

			tmp := rs.Primary.ID
			request.ComputeClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetComputeCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.ComputeClusterLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CoreComputeCluster") {
		resource.AddTestSweepers("CoreComputeCluster", &resource.Sweeper{
			Name:         "CoreComputeCluster",
			Dependencies: acctest.DependencyGraph["computeCluster"],
			F:            sweepCoreComputeClusterResource,
		})
	}
}

func sweepCoreComputeClusterResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	computeClusterIds, err := getCoreComputeClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, computeClusterId := range computeClusterIds {
		if ok := acctest.SweeperDefaultResourceId[computeClusterId]; !ok {
			deleteComputeClusterRequest := oci_core.DeleteComputeClusterRequest{}

			deleteComputeClusterRequest.ComputeClusterId = &computeClusterId

			deleteComputeClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteComputeCluster(context.Background(), deleteComputeClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting ComputeCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", computeClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &computeClusterId, CoreComputeClusterSweepWaitCondition, time.Duration(3*time.Minute),
				CoreComputeClusterSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreComputeClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ComputeClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

	listComputeClustersRequest := oci_core.ListComputeClustersRequest{}
	listComputeClustersRequest.CompartmentId = &compartmentId
	listComputeClustersResponse, err := computeClient.ListComputeClusters(context.Background(), listComputeClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ComputeCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, computeCluster := range listComputeClustersResponse.Items {
		id := *computeCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ComputeClusterId", id)
	}
	return resourceIds, nil
}

func CoreComputeClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if computeClusterResponse, ok := response.Response.(oci_core.GetComputeClusterResponse); ok {
		return computeClusterResponse.LifecycleState != oci_core.ComputeClusterLifecycleStateDeleted
	}
	return false
}

func CoreComputeClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeClient().GetComputeCluster(context.Background(), oci_core.GetComputeClusterRequest{
		ComputeClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
