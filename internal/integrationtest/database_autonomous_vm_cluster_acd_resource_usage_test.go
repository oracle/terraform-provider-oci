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
	DatabaseAutonomousVmClusterAcdResourceUsageDataSourceRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"compartment_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}
)

// issue-routing-tag: database/ExaCC
func TestDatabaseAutonomousVmClusterAcdResourceUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousVmClusterAcdResourceUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_vm_cluster_acd_resource_usages.test_autonomous_vm_cluster_acd_resource_usages"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceDependencies + ExaccAcdResourceConfig,
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_vm_cluster_acd_resource_usages", "test_autonomous_vm_cluster_acd_resource_usages", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterAcdResourceUsageDataSourceRepresentation) +
				compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceDependencies + ExaccAcdResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_cluster_id"),
				//resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.#"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_resource_usages.0.autonomous_container_database_vm_usage.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.available_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.largest_provisionable_autonomous_database_in_cpus"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_resource_usages.0.provisionable_cpus.#", "29"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.provisioned_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.reclaimable_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.reserved_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_resource_usages.0.used_cpus"),
			),
		},
	})
}
