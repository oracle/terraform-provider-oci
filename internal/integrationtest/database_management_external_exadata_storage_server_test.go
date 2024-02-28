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
	DatabaseManagementExternalExadataStorageServerRequiredOnlyResource = DatabaseManagementExternalExadataStorageServerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_server", "test_external_exadata_storage_server", acctest.Required, acctest.Create, DatabaseManagementExternalExadataStorageServerRepresentation)

	DatabaseManagementExternalExadataStorageServerResourceConfig = DatabaseManagementExternalExadataStorageServerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_server", "test_external_exadata_storage_server", acctest.Optional, acctest.Update, DatabaseManagementExternalExadataStorageServerRepresentation)

	DatabaseManagementDatabaseManagementExternalExadataStorageServerSingularDataSourceRepresentation = map[string]interface{}{
		"external_exadata_storage_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.storage_server_id}`},
	}

	DatabaseManagementExternalExadataStorageServerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"external_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${var.connector_exadata_infra_id}`},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"filter":                             acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalExadataStorageServerDataSourceFilterRepresentation}}
	DatabaseManagementExternalExadataStorageServerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_external_exadata_storage_server.test_external_exadata_storage_server.id}`}},
	}

	DatabaseManagementExternalExadataStorageServerRepresentation = map[string]interface{}{
		"external_exadata_storage_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.storage_server_id}`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}

	DatabaseManagementExternalExadataStorageServerResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalExadataStorageServerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalExadataStorageServerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_external_exadata_storage_server.test_external_exadata_storage_server"
	storageServerId := utils.GetEnvSettingWithBlankDefault("storage_server_id")
	storageServerIdVariableStr := fmt.Sprintf("variable \"storage_server_id\" { default = \"%s\" }\n", storageServerId)

	connectorExadataInfraId := utils.GetEnvSettingWithBlankDefault("connector_exadata_infra_id")
	connectorExadataInfraIdVariableStr := fmt.Sprintf("variable \"connector_exadata_infra_id\" { default = \"%s\" }\n", connectorExadataInfraId)

	datasourceName := "data.oci_database_management_external_exadata_storage_servers.test_external_exadata_storage_servers"
	singularDatasourceName := "data.oci_database_management_external_exadata_storage_server.test_external_exadata_storage_server"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+storageServerIdVariableStr+connectorExadataInfraIdVariableStr+DatabaseManagementExternalExadataStorageServerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_server", "test_external_exadata_storage_server", acctest.Optional, acctest.Create, DatabaseManagementExternalExadataStorageServerRepresentation), "databasemanagement", "externalExadataStorageServer", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + storageServerIdVariableStr + connectorExadataInfraIdVariableStr + DatabaseManagementExternalExadataStorageServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_server", "test_external_exadata_storage_server", acctest.Optional, acctest.Create, DatabaseManagementExternalExadataStorageServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "external_exadata_storage_server_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttrSet(resourceName, "resource_type"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + storageServerIdVariableStr + connectorExadataInfraIdVariableStr + DatabaseManagementExternalExadataStorageServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_server", "test_external_exadata_storage_server", acctest.Optional, acctest.Update, DatabaseManagementExternalExadataStorageServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "external_exadata_storage_server_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttrSet(resourceName, "resource_type"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + storageServerIdVariableStr + connectorExadataInfraIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_storage_servers", "test_external_exadata_storage_servers", acctest.Required, acctest.Update, DatabaseManagementExternalExadataStorageServerDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementExternalExadataStorageServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_server", "test_external_exadata_storage_server", acctest.Optional, acctest.Update, DatabaseManagementExternalExadataStorageServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "external_exadata_infrastructure_id"),

				resource.TestCheckResourceAttr(datasourceName, "external_exadata_storage_server_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_exadata_storage_server_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + storageServerIdVariableStr + connectorExadataInfraIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_storage_server", "test_external_exadata_storage_server", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalExadataStorageServerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementExternalExadataStorageServerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_exadata_storage_server_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "connector.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "internal_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "make_model"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_flash_disk_iops"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_flash_disk_throughput"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_hard_disk_iops"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_hard_disk_throughput"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_gb"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalExadataStorageServerRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"external_exadata_storage_server_id",
			},
			ResourceName: resourceName,
		},
	})
}
