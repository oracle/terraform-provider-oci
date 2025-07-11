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

var (
	DatabaseManagementCloudAsmRequiredOnlyResource = DatabaseManagementCloudAsmResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_asm", "test_cloud_asm", acctest.Required, acctest.Create, DatabaseManagementCloudAsmRepresentation)

	DatabaseManagementCloudAsmResourceConfig = DatabaseManagementCloudAsmResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_asm", "test_cloud_asm", acctest.Optional, acctest.Update, DatabaseManagementCloudAsmRepresentation)

	DatabaseManagementCloudAsmSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_asm_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_asms.test_cloud_asms.cloud_asm_collection.0.items.0.id}`},
	}

	DatabaseManagementCloudAsmDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.dbaas_dbsystem_id}`},
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}
	DatabaseManagementCloudAsmDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_cloud_asm.test_cloud_asm.id}`}},
	}

	DatabaseManagementCloudAsmRepresentation = map[string]interface{}{
		"cloud_asm_id":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_asms.test_cloud_asms.cloud_asm_collection.0.items.0.id}`},
		"cloud_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_asms.test_cloud_asms.cloud_asm_collection.0.items.0.cloud_connector_id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}

	DatabaseManagementCloudAsmResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_asms", "test_cloud_asms", acctest.Required, acctest.Create, DatabaseManagementCloudAsmDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudAsmResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudAsmResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_cloud_asm.test_cloud_asm"
	datasourceName := "data.oci_database_management_cloud_asms.test_cloud_asms"
	singularDatasourceName := "data.oci_database_management_cloud_asm.test_cloud_asm"

	dbaasDbsystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbaasDbsystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbaasDbsystemId)

	variableStr := compartmentIdVariableStr + dbaasDbsystemIdVariableStr
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+DatabaseManagementCloudAsmResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_asm", "test_cloud_asm", acctest.Optional, acctest.Create, DatabaseManagementCloudAsmRepresentation), "databasemanagement", "cloudAsm", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + variableStr + DatabaseManagementCloudAsmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_asm", "test_cloud_asm", acctest.Required, acctest.Create, DatabaseManagementCloudAsmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_asm_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Create with optionals
		{
			Config: config + variableStr + DatabaseManagementCloudAsmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_asm", "test_cloud_asm", acctest.Optional, acctest.Create, DatabaseManagementCloudAsmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_asm_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "component_name"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config + variableStr + DatabaseManagementCloudAsmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_asm", "test_cloud_asm", acctest.Optional, acctest.Update, DatabaseManagementCloudAsmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_db_system_id"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_asm_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_asm_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_asm", "test_cloud_asm", acctest.Required, acctest.Create, DatabaseManagementCloudAsmSingularDataSourceRepresentation) +
				variableStr + DatabaseManagementCloudAsmResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_asm_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grid_home"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cluster"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_flex_enabled"),
				resource.TestCheckResourceAttr(singularDatasourceName, "serviced_databases.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementCloudAsmRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cloud_asm_id",
			},
			ResourceName: resourceName,
		},
	})
}
