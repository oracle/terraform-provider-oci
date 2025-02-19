// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreComputeGpuMemoryClusterInstanceDataSourceRepresentation = map[string]interface{}{
		"compute_gpu_memory_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster.id}`},
	}

	CoreComputeGpuMemoryClusterInstanceResourceConfig = CoreComputeGpuMemoryClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster", "test_compute_gpu_memory_cluster", acctest.Required, acctest.Create, CoreComputeGpuMemoryClusterRepresentation)
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeGpuMemoryClusterInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeGpuMemoryClusterInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	gpuMemoryClusterSize := utils.GetEnvSettingWithDefault("gmc_size", "2")
	gpuMemoryClusterSizeVariableStr := fmt.Sprintf("variable \"gmc_size\" { default = \"%s\" }\n", gpuMemoryClusterSize)

	imageId := utils.GetEnvSettingWithBlankDefault("gb200_image_id")
	imageIdVariableStr := fmt.Sprintf("variable \"image_id\" { default = \"%s\" }\n", imageId)

	computeGpuMemoryFabricId := utils.GetEnvSettingWithBlankDefault("compute_gpu_memory_fabric_id")
	computeGpuMemoryFabricIdVariableStr := fmt.Sprintf("variable \"compute_gpu_memory_fabric_id\" { default = \"%s\" }\n", computeGpuMemoryFabricId)

	datasourceName := "data.oci_core_compute_gpu_memory_cluster_instances.test_compute_gpu_memory_cluster_instances"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_gpu_memory_cluster_instances", "test_compute_gpu_memory_cluster_instances", acctest.Required, acctest.Create, CoreComputeGpuMemoryClusterInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + gpuMemoryClusterSizeVariableStr + imageIdVariableStr + computeGpuMemoryFabricIdVariableStr + CoreComputeGpuMemoryClusterInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compute_gpu_memory_cluster_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "compute_gpu_memory_cluster_instance_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "compute_gpu_memory_cluster_instance_collection.0.items.#", gpuMemoryClusterSize),
			),
		},
	})
}
