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
	CoreComputeHostSingularDataSourceRepresentation = map[string]interface{}{
		"compute_host_id": acctest.Representation{RepType: acctest.Required, Create: `${var.baremetalhost_id}`},
	}

	CoreComputeHostDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":          acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compute_host_health":          acctest.Representation{RepType: acctest.Optional, Create: `computeHostHealth`},
		"compute_host_lifecycle_state": acctest.Representation{RepType: acctest.Optional, Create: `computeHostLifecycleState`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"network_resource_id":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_resource.test_resource.id}`},
	}

	CloudGuardResourceDataSourceRepresentation = map[string]interface{}{}

	CoreComputeHostResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resources", "test_resources", acctest.Required, acctest.Create, CloudGuardResourceDataSourceRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeHostResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeHostResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_hosts.test_compute_hosts"
	singularDatasourceName := "data.oci_core_compute_host.test_compute_host"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_hosts", "test_compute_hosts", acctest.Required, acctest.Create, CoreComputeHostDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_host_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_host", "test_compute_host", acctest.Required, acctest.Create, CoreComputeHostSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fault_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "health"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hpc_island_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "local_block_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gpu_memory_fabric_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_block_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
