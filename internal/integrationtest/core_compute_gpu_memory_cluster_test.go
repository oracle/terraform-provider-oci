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
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// uncomment for testing with tenancies that automatically added tags like Oracle-Tags
	//IgnoreDefinedTagsRepresentation = map[string]interface{}{
	//	"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	//}
	CoreComputeGpuMemoryClusterRequiredOnlyResource = CoreComputeGpuMemoryClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster", "test_compute_gpu_memory_cluster", acctest.Required, acctest.Create, CoreComputeGpuMemoryClusterRepresentation)

	CoreComputeGpuMemoryClusterResourceConfig = CoreComputeGpuMemoryClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster", "test_compute_gpu_memory_cluster", acctest.Optional, acctest.Update, CoreComputeGpuMemoryClusterRepresentation)

	CoreComputeGpuMemoryClusterSingularDataSourceRepresentation = map[string]interface{}{
		"compute_gpu_memory_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster.id}`},
	}

	CoreComputeGpuMemoryClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":           acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compute_cluster_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_cluster.test_compute_cluster.id}`},
		"compute_gpu_memory_cluster_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster.id}`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreComputeGpuMemoryClusterDataSourceFilterRepresentation}}
	CoreComputeGpuMemoryClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster.id}`}},
	}

	CoreComputeGpuMemoryClusterRepresentation = map[string]interface{}{
		"availability_domain":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_cluster_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_cluster.test_compute_cluster.id}`},
		"instance_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"gpu_memory_fabric_id":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_compute_gpu_memory_fabric.test_compute_gpu_memory_fabric.id}`},
		"size":                      acctest.Representation{RepType: acctest.Required, Create: `${var.gmc_size}`},
		// uncomment for testing with tenancies that automatically added tags like Oracle-Tags
		//"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreDefinedTagsRepresentation},
	}

	GpuMemoryClusterInstanceConfigurationInstanceDetailsRepresentation = map[string]interface{}{
		"instance_type":  acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"launch_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GpuMemoryClusterInstanceConfigurationInstanceDetailsLaunchDetailsRepresentation},
	}

	GpuMemoryClusterInstanceConfigurationInstanceDetailsLaunchDetailsRepresentation = acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetMultipleUpdatedRepresenationCopy(
			[]string{"shape", "source_details"},
			[]interface{}{
				acctest.Representation{RepType: acctest.Optional, Create: `BM.GPU.GB200.4`}, // modified shape to GB200
				acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("image_id", acctest.Representation{RepType: acctest.Optional, Create: `${var.image_id}`}, CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation)},
			},
			CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentation),
		[]string{"shape_config", "dedicated_vm_host_id", "is_pv_encryption_in_transit_enabled", "preferred_maintenance_action", "defined_tags", "freeform_tags"})

	GpuMemoryClusterInstanceConfigurationRepresentation = acctest.RepresentationCopyWithRemovedProperties(
		CoreInstanceConfigurationRepresentation,
		[]string{"defined_tags", "freeform_tags"})

	CoreComputeGpuMemoryClusterResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_compute_cluster", "test_compute_cluster", acctest.Required, acctest.Create, CoreComputeClusterRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
			acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: GpuMemoryClusterInstanceConfigurationInstanceDetailsRepresentation}, GpuMemoryClusterInstanceConfigurationRepresentation)) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_gpu_memory_fabric", "test_compute_gpu_memory_fabric", acctest.Required, acctest.Create, CoreComputeGpuMemoryFabricSingularDataSourceRepresentation)
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeGpuMemoryClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeGpuMemoryClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	gpuMemoryClusterSize := utils.GetEnvSettingWithDefault("gmc_size", "2")
	gpuMemoryClusterSizeVariableStr := fmt.Sprintf("variable \"gmc_size\" { default = \"%s\" }\n", gpuMemoryClusterSize)

	imageId := utils.GetEnvSettingWithBlankDefault("gb200_image_id")
	imageIdVariableStr := fmt.Sprintf("variable \"image_id\" { default = \"%s\" }\n", imageId)

	computeGpuMemoryFabricId := utils.GetEnvSettingWithBlankDefault("compute_gpu_memory_fabric_id")
	computeGpuMemoryFabricIdVariableStr := fmt.Sprintf("variable \"compute_gpu_memory_fabric_id\" { default = \"%s\" }\n", computeGpuMemoryFabricId)

	resourceName := "oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster"
	datasourceName := "data.oci_core_compute_gpu_memory_clusters.test_compute_gpu_memory_clusters"
	singularDatasourceName := "data.oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+gpuMemoryClusterSizeVariableStr+imageIdVariableStr+computeGpuMemoryFabricIdVariableStr+CoreComputeGpuMemoryClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster", "test_compute_gpu_memory_cluster", acctest.Optional, acctest.Create, CoreComputeGpuMemoryClusterRepresentation), "core", "computeGpuMemoryCluster", t)

	acctest.ResourceTest(t, testAccCheckCoreComputeGpuMemoryClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + gpuMemoryClusterSizeVariableStr + imageIdVariableStr + computeGpuMemoryFabricIdVariableStr + CoreComputeGpuMemoryClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster", "test_compute_gpu_memory_cluster", acctest.Required, acctest.Create, CoreComputeGpuMemoryClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compute_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gpu_memory_fabric_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + gpuMemoryClusterSizeVariableStr + imageIdVariableStr + computeGpuMemoryFabricIdVariableStr + CoreComputeGpuMemoryClusterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + gpuMemoryClusterSizeVariableStr + imageIdVariableStr + computeGpuMemoryFabricIdVariableStr + CoreComputeGpuMemoryClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster", "test_compute_gpu_memory_cluster", acctest.Optional, acctest.Create, CoreComputeGpuMemoryClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compute_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gpu_memory_fabric_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "size", gpuMemoryClusterSize),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + gpuMemoryClusterSizeVariableStr + imageIdVariableStr + computeGpuMemoryFabricIdVariableStr + CoreComputeGpuMemoryClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster", "test_compute_gpu_memory_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreComputeGpuMemoryClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "compute_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gpu_memory_fabric_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "size", gpuMemoryClusterSize),
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
			Config: config + compartmentIdVariableStr + gpuMemoryClusterSizeVariableStr + imageIdVariableStr + computeGpuMemoryFabricIdVariableStr + CoreComputeGpuMemoryClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster", "test_compute_gpu_memory_cluster", acctest.Optional, acctest.Update, CoreComputeGpuMemoryClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compute_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gpu_memory_fabric_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "size", gpuMemoryClusterSize),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_gpu_memory_clusters", "test_compute_gpu_memory_clusters", acctest.Optional, acctest.Update, CoreComputeGpuMemoryClusterDataSourceRepresentation) +
				compartmentIdVariableStr + gpuMemoryClusterSizeVariableStr + imageIdVariableStr + computeGpuMemoryFabricIdVariableStr + CoreComputeGpuMemoryClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster", "test_compute_gpu_memory_cluster", acctest.Optional, acctest.Update, CoreComputeGpuMemoryClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_gpu_memory_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "compute_gpu_memory_cluster_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compute_gpu_memory_cluster_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster", "test_compute_gpu_memory_cluster", acctest.Required, acctest.Create, CoreComputeGpuMemoryClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + gpuMemoryClusterSizeVariableStr + imageIdVariableStr + computeGpuMemoryFabricIdVariableStr + CoreComputeGpuMemoryClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_gpu_memory_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "size", gpuMemoryClusterSize),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreComputeGpuMemoryClusterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreComputeGpuMemoryClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_compute_gpu_memory_cluster" {
			noResourceFound = false
			request := oci_core.GetComputeGpuMemoryClusterRequest{}

			tmp := rs.Primary.ID
			request.ComputeGpuMemoryClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetComputeGpuMemoryCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.ComputeGpuMemoryClusterLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CoreComputeGpuMemoryCluster") {
		resource.AddTestSweepers("CoreComputeGpuMemoryCluster", &resource.Sweeper{
			Name:         "CoreComputeGpuMemoryCluster",
			Dependencies: acctest.DependencyGraph["computeGpuMemoryCluster"],
			F:            sweepCoreComputeGpuMemoryClusterResource,
		})
	}
}

func sweepCoreComputeGpuMemoryClusterResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	computeGpuMemoryClusterIds, err := getCoreComputeGpuMemoryClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, computeGpuMemoryClusterId := range computeGpuMemoryClusterIds {
		if ok := acctest.SweeperDefaultResourceId[computeGpuMemoryClusterId]; !ok {
			deleteComputeGpuMemoryClusterRequest := oci_core.DeleteComputeGpuMemoryClusterRequest{}

			deleteComputeGpuMemoryClusterRequest.ComputeGpuMemoryClusterId = &computeGpuMemoryClusterId

			deleteComputeGpuMemoryClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteComputeGpuMemoryCluster(context.Background(), deleteComputeGpuMemoryClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting ComputeGpuMemoryCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", computeGpuMemoryClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &computeGpuMemoryClusterId, CoreComputeGpuMemoryClusterSweepWaitCondition, time.Duration(3*time.Minute),
				CoreComputeGpuMemoryClusterSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreComputeGpuMemoryClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ComputeGpuMemoryClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

	listComputeGpuMemoryClustersRequest := oci_core.ListComputeGpuMemoryClustersRequest{}
	listComputeGpuMemoryClustersRequest.CompartmentId = &compartmentId
	listComputeGpuMemoryClustersResponse, err := computeClient.ListComputeGpuMemoryClusters(context.Background(), listComputeGpuMemoryClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ComputeGpuMemoryCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, computeGpuMemoryCluster := range listComputeGpuMemoryClustersResponse.Items {
		id := *computeGpuMemoryCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ComputeGpuMemoryClusterId", id)
	}
	return resourceIds, nil
}

func CoreComputeGpuMemoryClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if computeGpuMemoryClusterResponse, ok := response.Response.(oci_core.GetComputeGpuMemoryClusterResponse); ok {
		return computeGpuMemoryClusterResponse.LifecycleState != oci_core.ComputeGpuMemoryClusterLifecycleStateDeleted
	}
	return false
}

func CoreComputeGpuMemoryClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeClient().GetComputeGpuMemoryCluster(context.Background(), oci_core.GetComputeGpuMemoryClusterRequest{
		ComputeGpuMemoryClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
