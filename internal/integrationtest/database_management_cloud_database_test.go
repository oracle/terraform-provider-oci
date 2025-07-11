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
	DatabaseManagementCloudDatabaseDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.dbaas_dbsystem_id}`},
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	DatabaseManagementCloudDatabaseResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbaasDbsystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbaasDbsystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbaasDbsystemId)

	datasourceName := "data.oci_database_management_cloud_databases.test_cloud_databases"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_databases", "test_cloud_databases", acctest.Required, acctest.Create, DatabaseManagementCloudDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + dbaasDbsystemIdVariableStr + DatabaseManagementCloudDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_database_collection.0.items.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_database_collection.#"),
			),
		},
	})
}
