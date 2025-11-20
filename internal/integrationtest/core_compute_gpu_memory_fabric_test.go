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
	CoreComputeGpuMemoryFabricSingularDataSourceRepresentation = map[string]interface{}{
		"compute_gpu_memory_fabric_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compute_gpu_memory_fabric_id}`},
	}

	CoreComputeGpuMemoryFabricDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_gpu_memory_fabric_health":          acctest.Representation{RepType: acctest.Optional, Create: `HEALTHY`},
		"compute_gpu_memory_fabric_lifecycle_state": acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"display_name":                              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}

	CoreComputeGpuMemoryFabricRepresentation = map[string]interface{}{
		"compute_gpu_memory_fabric_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compute_gpu_memory_fabric_id}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"memory_fabric_preferences":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreComputeGpuMemoryFabricMemoryFabricPreferencesRepresentation},
	}
	CoreComputeGpuMemoryFabricMemoryFabricPreferencesRepresentation = map[string]interface{}{
		"customer_desired_firmware_bundle_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.firmware_bundle_id}`},
		"fabric_recycle_level":                acctest.Representation{RepType: acctest.Optional, Create: `FULL_RECYCLE`, Update: `SKIP_RECYCLE`},
	}
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeGpuMemoryFabricResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeGpuMemoryFabricResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	computeGpuMemoryFabricId := utils.GetEnvSettingWithBlankDefault("compute_gpu_memory_fabric_id")
	computeGpuMemoryFabricIdVariableStr := fmt.Sprintf("variable \"compute_gpu_memory_fabric_id\" { default = \"%s\" }\n", computeGpuMemoryFabricId)

	// resource create/destroy not supported for compute_gpu_memory_fabric
	datasourceName := "data.oci_core_compute_gpu_memory_fabrics.test_compute_gpu_memory_fabrics"
	singularDatasourceName := "data.oci_core_compute_gpu_memory_fabric.test_compute_gpu_memory_fabric"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + computeGpuMemoryFabricIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap(
				"oci_core_compute_gpu_memory_fabrics", "test_compute_gpu_memory_fabrics", acctest.Optional, acctest.Create, CoreComputeGpuMemoryFabricDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compute_gpu_memory_fabric_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + computeGpuMemoryFabricIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap(
				"oci_core_compute_gpu_memory_fabric", "test_compute_gpu_memory_fabric", acctest.Required, acctest.Create, CoreComputeGpuMemoryFabricSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_gpu_memory_fabric_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_host_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_local_block_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_firmware_bundle_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fabric_health"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "firmware_update_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "healthy_host_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_platform_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "memory_fabric_preferences.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "switch_platform_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_firmware_bundle_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_host_count"),
			),
		},
	})
}
