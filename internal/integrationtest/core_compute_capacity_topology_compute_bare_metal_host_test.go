// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreComputeCapacityTopologyComputeBareMetalHostDataSourceRepresentation = map[string]interface{}{
		"compute_capacity_topology_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_capacity_topology.test_compute_capacity_topology.id}`},
		"availability_domain":          acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	CoreComputeCapacityTopologyComputeBareMetalHostResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_topology", "test_compute_capacity_topology", acctest.Required, acctest.Create, CoreComputeCapacityTopologyRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: core/computeBm
func TestCoreComputeCapacityTopologyComputeBareMetalHostResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityTopologyComputeBareMetalHostResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_capacity_topology_compute_bare_metal_hosts.test_compute_capacity_topology_compute_bare_metal_hosts"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_topology_compute_bare_metal_hosts", "test_compute_capacity_topology_compute_bare_metal_hosts", acctest.Required, acctest.Create, CoreComputeCapacityTopologyComputeBareMetalHostDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeCapacityTopologyComputeBareMetalHostResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_topology_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "compute_bare_metal_host_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "compute_bare_metal_host_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_bare_metal_host_collection.0.items.0.compute_hpc_island_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_bare_metal_host_collection.0.items.0.compute_local_block_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_bare_metal_host_collection.0.items.0.compute_network_block_id"),
			),
		},
	})
}
