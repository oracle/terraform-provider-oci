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
	DatabaseManagementExternalExadataStorageGridRequiredOnlyResource = DatabaseManagementExternalExadataStorageGridResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_grid", "test_external_exadata_storage_grid", acctest.Required, acctest.Create, DatabaseManagementExternalExadataStorageGridRepresentation)

	DatabaseManagementExternalExadataStorageGridResourceConfig = DatabaseManagementExternalExadataStorageGridResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_grid", "test_external_exadata_storage_grid", acctest.Optional, acctest.Update, DatabaseManagementExternalExadataStorageGridRepresentation)

	DatabaseManagementDatabaseManagementExternalExadataStorageGridSingularDataSourceRepresentation = map[string]interface{}{
		"external_exadata_storage_grid_id": acctest.Representation{RepType: acctest.Required, Create: `${var.storage_server_grid}`},
	}

	DatabaseManagementExternalExadataStorageGridRepresentation = map[string]interface{}{
		"external_exadata_storage_grid_id": acctest.Representation{RepType: acctest.Required, Create: `${var.storage_server_grid}`},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}

	DatabaseManagementExternalExadataStorageGridResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalExadataStorageGridResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalExadataStorageGridResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_external_exadata_storage_grid.test_external_exadata_storage_grid"
	storageServerGridIdUri := utils.GetEnvSettingWithBlankDefault("storage_server_grid")
	storageServerGridIdUriVariableStr := fmt.Sprintf("variable \"storage_server_grid\" { default = \"%s\" }\n", storageServerGridIdUri)

	singularDatasourceName := "data.oci_database_management_external_exadata_storage_grid.test_external_exadata_storage_grid"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+storageServerGridIdUriVariableStr+DatabaseManagementExternalExadataStorageGridResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_grid", "test_external_exadata_storage_grid", acctest.Optional, acctest.Create, DatabaseManagementExternalExadataStorageGridRepresentation), "databasemanagement", "externalExadataStorageGrid", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + storageServerGridIdUriVariableStr + DatabaseManagementExternalExadataStorageGridResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_grid", "test_external_exadata_storage_grid", acctest.Optional, acctest.Create, DatabaseManagementExternalExadataStorageGridRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "external_exadata_storage_grid_id"),
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
			Config: config + compartmentIdVariableStr + storageServerGridIdUriVariableStr + DatabaseManagementExternalExadataStorageGridResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_grid", "test_external_exadata_storage_grid", acctest.Optional, acctest.Update, DatabaseManagementExternalExadataStorageGridRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "external_exadata_storage_grid_id"),
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
		// verify singular datasource
		{
			Config: config + storageServerGridIdUriVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_storage_grid", "test_external_exadata_storage_grid", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalExadataStorageGridSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementExternalExadataStorageGridResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_exadata_storage_grid_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "server_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_servers.#", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalExadataStorageGridRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"external_exadata_storage_grid_id",
			},
			ResourceName: resourceName,
		},
	})
}
