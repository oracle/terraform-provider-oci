// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementExternalAsmInstanceRequiredOnlyResource = DatabaseManagementExternalAsmInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_asm_instance", "test_external_asm_instance", acctest.Required, acctest.Create, DatabaseManagementExternalAsmInstanceRepresentation)

	DatabaseManagementExternalAsmInstanceResourceConfig = DatabaseManagementExternalAsmInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_asm_instance", "test_external_asm_instance", acctest.Optional, acctest.Update, DatabaseManagementExternalAsmInstanceRepresentation)

	DatabaseManagementDatabaseManagementExternalAsmInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"external_asm_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_asm_instances.test_external_asm_instances.external_asm_instance_collection.0.items.0.id}`},
	}

	DatabaseManagementDatabaseManagementExternalAsmInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"external_asm_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_asms.test_external_asms.external_asm_collection.0.items.0.id}`},
	}

	DatabaseManagementExternalAsmInstanceRepresentation = map[string]interface{}{
		"external_asm_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_asm_instances.test_external_asm_instances.external_asm_instance_collection.0.items.0.id}`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}

	DatabaseManagementExternalAsmInstanceResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asms", "test_external_asms", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_asm_instances", "test_external_asm_instances", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalAsmInstanceDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalAsmInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalAsmInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_external_asm_instance.test_external_asm_instance"
	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	datasourceName := "data.oci_database_management_external_asm_instances.test_external_asm_instances"
	singularDatasourceName := "data.oci_database_management_external_asm_instance.test_external_asm_instance"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbSystemIdVariableStr+DatabaseManagementExternalAsmInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_asm_instance", "test_external_asm_instance", acctest.Optional, acctest.Create, DatabaseManagementExternalAsmInstanceRepresentation), "databasemanagement", "externalAsmInstance", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalAsmInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_asm_instance", "test_external_asm_instance", acctest.Optional, acctest.Create, DatabaseManagementExternalAsmInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "component_name"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "external_asm_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_asm_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalAsmInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_asm_instance", "test_external_asm_instance", acctest.Optional, acctest.Update, DatabaseManagementExternalAsmInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_asm_id"),

				resource.TestCheckResourceAttr(datasourceName, "external_asm_instance_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_asm_instance_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
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
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalAsmInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"external_asm_instance_id",
			},
			ResourceName: resourceName,
		},
	})
}
