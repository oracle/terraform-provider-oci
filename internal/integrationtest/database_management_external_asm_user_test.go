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
	DatabaseManagementDatabaseManagementExternalAsmUserDataSourceRepresentation = map[string]interface{}{
		"external_asm_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.external_asm_id}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementExternalAsmUserResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asms", "test_external_asms", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalAsmUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalAsmUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	externalAsmId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_asm_id")
	externalAsmIdVariableStr := fmt.Sprintf("variable \"external_asm_id\" { default = \"%s\" }\n", externalAsmId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_external_asm_users.test_external_asm_users"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asm_users", "test_external_asm_users", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmUserDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + externalAsmIdVariableStr + DatabaseManagementExternalAsmUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_asm_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_asm_user_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_asm_user_collection.0.items.#"),
			),
		},
		// verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asm_users", "test_external_asm_users", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmUserDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + externalAsmIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementExternalAsmUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_asm_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "opc_named_credential_id"),
			),
		},
	})
}
