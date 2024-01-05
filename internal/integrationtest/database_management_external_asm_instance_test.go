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
	DatabaseManagementDatabaseManagementExternalAsmInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"external_asm_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_asm_instances.test_external_asm_instances.external_asm_instance_collection.0.items.0.id}`},
	}

	DatabaseManagementDatabaseManagementExternalAsmInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"external_asm_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_asms.test_external_asms.external_asm_collection.0.items.0.id}`},
	}

	DatabaseManagementExternalAsmInstanceResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asms", "test_external_asms", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalAsmInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalAsmInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	datasourceName := "data.oci_database_management_external_asm_instances.test_external_asm_instances"
	singularDatasourceName := "data.oci_database_management_external_asm_instance.test_external_asm_instance"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asm_instances", "test_external_asm_instances", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalAsmInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_asm_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "external_asm_instance_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asm_instances", "test_external_asm_instances", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmInstanceDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asm_instance", "test_external_asm_instance", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalAsmInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_asm_instance_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "adr_home_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_db_node_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_db_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
