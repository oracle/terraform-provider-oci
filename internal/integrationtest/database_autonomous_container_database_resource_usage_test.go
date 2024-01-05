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
	DatabaseAutonomousContainerDatabaseResourceUsageSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	DatabaseAutonomousContainerDatabaseResourceUsageResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentation) +
		DatabaseCloudAutonomousVmClusterRequiredOnlyResource
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseResourceUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseResourceUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_container_database_resource_usage.test_autonomous_container_database_resource_usage"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_resource_usage", "test_autonomous_container_database_resource_usage", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseResourceUsageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseResourceUsageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "autonomous_container_database_vm_usage.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "largest_provisionable_autonomous_database_in_cpus"),
				resource.TestCheckResourceAttr(singularDatasourceName, "provisionable_cpus.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "provisioned_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reclaimable_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reserved_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "used_cpus"),
			),
		},
	})
}
