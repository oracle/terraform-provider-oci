// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	DatabaseCloudAutonomousVmClusterAcdResourceUsageDataSourceRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compartment_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DatabaseCloudAutonomousVmClusterAcdResourceUsageResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseCloudAutonomousVmClusterAcdResourceUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudAutonomousVmClusterAcdResourceUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_cloud_autonomous_vm_cluster_acd_resource_usages.test_cloud_autonomous_vm_cluster_acd_resource_usages"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster_acd_resource_usages", "test_cloud_autonomous_vm_cluster_acd_resource_usages", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterAcdResourceUsageDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseCloudAutonomousVmClusterAcdResourceUsageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.#"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_resource_usages.0.autonomous_container_database_vm_usage.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.available_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.largest_provisionable_autonomous_database_in_cpus"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_resource_usages.0.provisionable_cpus.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.provisioned_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.reclaimable_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.reserved_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.used_cpus"),
			),
		},
	})
}
