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
	DatabaseManagementDatabaseManagementExternalAsmConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"external_asm_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.external_asm_id}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementExternalAsmConfigurationResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asms", "test_external_asms", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalAsmConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalAsmConfigurationResource_basic")
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

	singularDatasourceName := "data.oci_database_management_external_asm_configuration.test_external_asm_configuration"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asm_configuration", "test_external_asm_configuration", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + externalAsmIdVariableStr + DatabaseManagementExternalAsmConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_asm_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "init_parameters.#"),
			),
		},
		// verify singular datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asm_configuration", "test_external_asm_configuration", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + externalAsmIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementExternalAsmConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_asm_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opc_named_credential_id"),
			),
		},
	})
}
