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
	DistributedDatabaseDistributedAutonomousDatabaseRuChangeLogSingularDataSourceRepresentation = map[string]interface{}{
		"distributed_autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database.id}`},
		"source_shard_name":                  acctest.Representation{RepType: acctest.Optional, Create: `sourceShardName`},
	}

	DistributedDatabaseDistributedAutonomousDatabaseRuChangeLogResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedAutonomousDatabaseRuChangeLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedAutonomousDatabaseRuChangeLogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_distributed_database_distributed_autonomous_database_ru_change_log.test_distributed_autonomous_database_ru_change_log"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database_ru_change_log", "test_distributed_autonomous_database_ru_change_log", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRuChangeLogSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseRuChangeLogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "distributed_autonomous_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_shard_name", "sourceShardName"),

				resource.TestCheckResourceAttr(singularDatasourceName, "ru_change_log.#", "1"),
			),
		},
	})
}
