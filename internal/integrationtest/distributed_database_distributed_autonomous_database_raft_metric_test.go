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

	DistributedDatabaseDistributedAutonomousDatabaseRaftMetricResourceConfig = DistributedDatabaseDistributedAutonomousDatabaseResourceConfig
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedAutonomousDatabaseRaftMetricResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedAutonomousDatabaseRaftMetricResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	privateEndpointId := utils.GetEnvSettingWithBlankDefault("private_endpoint_id")
	privateEndpointIdVariableStr := fmt.Sprintf("variable \"private_endpoint_id\" { default = \"%s\" }\n", privateEndpointId)
	cloudAutonomousVmClusterId := utils.GetEnvSettingWithBlankDefault("cloud_autonomous_vm_cluster_id")
	cloudAutonomousVmClusterIdVariableStr := fmt.Sprintf("variable \"cloud_autonomous_vm_cluster_id\" { default = \"%s\" }\n", cloudAutonomousVmClusterId)
	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)
	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	resourceName := "oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database"
	singularDatasourceName := "data.oci_distributed_database_distributed_autonomous_database_raft_metric.test_distributed_autonomous_database_raft_metric"
	resourceConfig := config + compartmentIdVariableStr + privateEndpointIdVariableStr + cloudAutonomousVmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseRaftMetricResourceConfig
	dataSourceConfig := acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database_raft_metric", "test_distributed_autonomous_database_raft_metric", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRaftMetricSingularDataSourceRepresentation) +
		resourceConfig

	acctest.SaveConfigContent(resourceConfig, "distributeddatabase", "distributedAutonomousDatabaseRaftMetric", t)

	acctest.ResourceTest(t, testAccCheckDistributedDatabaseDistributedAutonomousDatabaseDestroy, []resource.TestStep{
		// verify singular datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "distributed_autonomous_database_id"),
			),
		},
	})
}
