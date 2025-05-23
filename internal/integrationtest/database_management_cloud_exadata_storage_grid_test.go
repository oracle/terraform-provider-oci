// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementCloudExadataStorageGridRequiredOnlyResource = DatabaseManagementCloudExadataStorageGridResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_grid", "test_cloud_exadata_storage_grid", acctest.Required, acctest.Create, DatabaseManagementCloudExadataStorageGridRepresentation)

	DatabaseManagementCloudExadataStorageGridResourceConfig = DatabaseManagementCloudExadataStorageGridResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_grid", "test_cloud_exadata_storage_grid", acctest.Optional, acctest.Update, DatabaseManagementCloudExadataStorageGridRepresentation)

	DatabaseManagementCloudExadataStorageGridSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_exadata_storage_grid_id": acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_exadata_storage_grid_id}`},
	}

	DatabaseManagementCloudExadataStorageGridRepresentation = map[string]interface{}{
		"cloud_exadata_storage_grid_id": acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_exadata_storage_grid_id}`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseManagementCloudExadataStorageGridResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudExadataStorageGridResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudExadataStorageGridResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	cloudExadataStorageGridId := utils.GetEnvSettingWithBlankDefault("cloud_exadata_storage_grid_id")
	cloudExadataStorageGridIdVariableStr := fmt.Sprintf("variable \"cloud_exadata_storage_grid_id\" { default = \"%s\" }\n", cloudExadataStorageGridId)

	resourceName := "oci_database_management_cloud_exadata_storage_grid.test_cloud_exadata_storage_grid"

	singularDatasourceName := "data.oci_database_management_cloud_exadata_storage_grid.test_cloud_exadata_storage_grid"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseManagementCloudExadataStorageGridResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_grid", "test_cloud_exadata_storage_grid", acctest.Optional, acctest.Create, DatabaseManagementCloudExadataStorageGridRepresentation), "databasemanagement", "cloudExadataStorageGrid", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + cloudExadataStorageGridIdVariableStr + DatabaseManagementCloudExadataStorageGridResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_grid", "test_cloud_exadata_storage_grid", acctest.Required, acctest.Create, DatabaseManagementCloudExadataStorageGridRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_storage_grid_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + cloudExadataStorageGridIdVariableStr + DatabaseManagementCloudExadataStorageGridResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + cloudExadataStorageGridIdVariableStr + DatabaseManagementCloudExadataStorageGridResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_grid", "test_cloud_exadata_storage_grid", acctest.Optional, acctest.Create, DatabaseManagementCloudExadataStorageGridRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_storage_grid_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + cloudExadataStorageGridIdVariableStr + DatabaseManagementCloudExadataStorageGridResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_grid", "test_cloud_exadata_storage_grid", acctest.Optional, acctest.Update, DatabaseManagementCloudExadataStorageGridRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_storage_grid_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_grid", "test_cloud_exadata_storage_grid", acctest.Required, acctest.Create, DatabaseManagementCloudExadataStorageGridSingularDataSourceRepresentation) +
				compartmentIdVariableStr + cloudExadataStorageGridIdVariableStr + DatabaseManagementCloudExadataStorageGridResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_exadata_storage_grid_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "server_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_servers.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementCloudExadataStorageGridRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cloud_exadata_storage_grid_id",
			},
			ResourceName: resourceName,
		},
	})
}
