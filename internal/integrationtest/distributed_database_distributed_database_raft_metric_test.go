// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func TestDistributedDatabaseDistributedDatabaseRaftMetricResource_existingDdb(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedDatabaseRaftMetricResource_existingDdb")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	ddbId := utils.GetEnvSettingWithBlankDefault("distributed_database_id")
	if ddbId == "" {
		t.Fatal("TF_VAR_distributed_database_id must be set")
	}
	ddbIdVariableStr := fmt.Sprintf("variable \"distributed_database_id\" { default = \"%s\" }\n", ddbId)

	dsName := "data.oci_distributed_database_distributed_database_raft_metric.test"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + ddbIdVariableStr + `
data "oci_distributed_database_distributed_database_raft_metric" "test" {
  distributed_database_id = var.distributed_database_id
}
`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				testCheckResourcePrimaryIDSet(dsName),
			),
		},
	})
}

func testCheckResourcePrimaryIDSet(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		resourceState, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource %q not found in state", name)
		}

		if resourceState.Primary.ID == "" {
			return fmt.Errorf("resource %q primary id is empty", name)
		}

		return nil
	}
}

// issue-routing-tag: distributed_database/default
//func TestDistributedDatabaseDistributedDatabaseRaftMetricResource_basic(t *testing.T) {
//	httpreplay.SetScenario("TestDistributedDatabaseDistributedDatabaseRaftMetricResource_basic")
//	defer httpreplay.SaveScenario()
//
//	config := acctest.ProviderTestConfig()
//
//	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
//	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
//
//	singularDatasourceName := "data.oci_distributed_database_distributed_database_raft_metric.test_distributed_database_raft_metric"
//
//	acctest.SaveConfigContent("", "", "", t)
//
//	acctest.ResourceTest(t, nil, []resource.TestStep{
//		// verify singular datasource
//		{
//			Config: config +
//				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_database_raft_metric", "test_distributed_database_raft_metric", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabaseRaftMetricSingularDataSourceRepresentation) +
//				compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseRaftMetricResourceConfig,
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(singularDatasourceName, "distributed_database_id"),
//			),
//		},
//	})
//}
