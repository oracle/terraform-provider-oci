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
	DistributedDatabaseDistributedAutonomousDatabaseRaftMetricSingularDataSourceRepresentation = map[string]interface{}{
		"distributed_autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database.id}`},
	}

	DistributedDatabaseDistributedAutonomousDatabaseRaftMetricResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedAutonomousDatabaseRaftMetricResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedAutonomousDatabaseRaftMetricResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_distributed_database_distributed_autonomous_database_raft_metric.test_distributed_autonomous_database_raft_metric"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database_raft_metric", "test_distributed_autonomous_database_raft_metric", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRaftMetricSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseRaftMetricResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "distributed_autonomous_database_id"),
			),
		},
	})
}
