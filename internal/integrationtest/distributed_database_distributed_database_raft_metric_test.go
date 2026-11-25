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
	DistributedDatabaseDistributedDatabaseRaftMetricSingularDataSourceRepresentation = map[string]interface{}{
		"distributed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_database.test_distributed_database.id}`},
	}

	DistributedDatabaseDistributedDatabaseRaftMetricResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, DatabaseDbHomeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Required, acctest.Create, DatabaseExascaleDbStorageVaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DatabaseVmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabaseRepresentation) +
		AvailabilityDomainConfig +
		KeyResourceDependencyConfig
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedDatabaseRaftMetricResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedDatabaseRaftMetricResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_distributed_database_distributed_database_raft_metric.test_distributed_database_raft_metric"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_database_raft_metric", "test_distributed_database_raft_metric", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabaseRaftMetricSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseRaftMetricResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "distributed_database_id"),
			),
		},
	})
}
