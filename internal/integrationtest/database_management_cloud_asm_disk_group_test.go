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
	DatabaseManagementCloudAsmDiskGroupDataSourceRepresentation = map[string]interface{}{
		"cloud_asm_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_asm_id}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementCloudAsmDiskGroupResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_asms", "test_cloud_asms", acctest.Required, acctest.Create, DatabaseManagementCloudAsmDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudAsmDiskGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudAsmDiskGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbaasDbsystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbaasDbsystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbaasDbsystemId)

	cloudAsmId := utils.GetEnvSettingWithBlankDefault("dbmgmt_cloud_asm_id")
	cloudAsmIdVariableStr := fmt.Sprintf("variable \"cloud_asm_id\" { default = \"%s\" }\n", cloudAsmId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_cloud_asm_disk_groups.test_cloud_asm_disk_groups"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_asm_disk_groups", "test_cloud_asm_disk_groups", acctest.Required, acctest.Create, DatabaseManagementCloudAsmDiskGroupDataSourceRepresentation) +
				compartmentIdVariableStr + dbaasDbsystemIdVariableStr + cloudAsmIdVariableStr + DatabaseManagementCloudAsmDiskGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_asm_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "cloud_asm_disk_group_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_asm_disk_group_collection.0.items.#"),
			),
		},
		// verify singular datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_asm_disk_groups", "test_cloud_asm_disk_groups", acctest.Optional, acctest.Create, DatabaseManagementCloudAsmDiskGroupDataSourceRepresentation) +
				compartmentIdVariableStr + dbaasDbsystemIdVariableStr + cloudAsmIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementCloudAsmDiskGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_asm_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "opc_named_credential_id"),
			),
		},
	})
}
