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
	DatabaseCloudAutonomousVmClusterResourceUsageSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
	}

	DatabaseCloudAutonomousVmClusterResourceUsageResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseCloudAutonomousVmClusterResourceUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudAutonomousVmClusterResourceUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_cloud_autonomous_vm_cluster_resource_usage.test_cloud_autonomous_vm_cluster_resource_usage"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster_resource_usage", "test_cloud_autonomous_vm_cluster_resource_usage", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterResourceUsageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseCloudAutonomousVmClusterResourceUsageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_autonomous_vm_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "autonomous_vm_resource_usage.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_autonomous_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_storage_size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_storage_in_tbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_per_oracle_compute_unit_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "non_provisionable_autonomous_container_databases"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "provisionable_autonomous_container_databases"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "provisioned_autonomous_container_databases"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "provisioned_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reclaimable_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reserved_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_container_databases"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "used_autonomous_data_storage_size_in_tbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "used_cpus"),
			),
		},
	})
}
