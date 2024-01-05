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
	DatabaseDatabaseDbVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_system_id":                         acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_db_system.test_db_system.id}`},
		"db_system_shape":                      acctest.Representation{RepType: acctest.Optional, Create: `BM.DenseIO2.52`},
		"is_database_software_image_supported": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_upgrade_supported":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"storage_management":                   acctest.Representation{RepType: acctest.Optional, Create: `ASM`},
	}
	dbVersionDataSourceRepresentationRequiredOnly = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}
	dbVersionDataSourceRepresentationWithDbSystemIdOptional = acctest.RepresentationCopyWithNewProperties(dbVersionDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_db_system.test_db_system.id}`},
	})
	dbVersionDataSourceRepresentationWithUpgradeSupportedOptional = acctest.RepresentationCopyWithNewProperties(dbVersionDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"db_system_shape":      acctest.Representation{RepType: acctest.Optional, Create: `BM.DenseIO2.52`},
		"is_upgrade_supported": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	})
	dbVersionDataSourceRepresentationWithDbSystemShapeOptional = acctest.RepresentationCopyWithNewProperties(dbVersionDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"db_system_shape": acctest.Representation{RepType: acctest.Optional, Create: `BM.DenseIO2.52`},
	})
	dbVersionDataSourceRepresentationWithStorageManagementOptional = acctest.RepresentationCopyWithNewProperties(dbVersionDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"storage_management": acctest.Representation{RepType: acctest.Optional, Create: `ASM`},
	})
	DatabaseDbVersionResourceConfig = DbSystemResourceConfig
)

// issue-routing-tag: database/default
func TestDatabaseDbVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_versions.test_db_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_versions", "test_db_versions", acctest.Required, acctest.Create, dbVersionDataSourceRepresentationRequiredOnly) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_versions", "test_db_versions_by_db_system_id", acctest.Optional, acctest.Create, dbVersionDataSourceRepresentationWithDbSystemIdOptional) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_versions", "test_db_versions_by_db_system_shape", acctest.Optional, acctest.Create, dbVersionDataSourceRepresentationWithDbSystemShapeOptional) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_versions", "test_db_versions_by_is_upgrade_supported", acctest.Optional, acctest.Create, dbVersionDataSourceRepresentationWithUpgradeSupportedOptional) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_versions", "test_db_versions_by_storage_management", acctest.Optional, acctest.Create, dbVersionDataSourceRepresentationWithStorageManagementOptional) +
				compartmentIdVariableStr + DatabaseDbVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "db_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_versions.0.is_latest_for_major_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_versions.0.is_preview_db_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_versions.0.is_upgrade_supported"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_versions.0.supports_pdb"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_versions.0.version"),

				resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_id", "db_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_id", "db_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_id", "db_versions.0.supports_pdb"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_id", "db_versions.0.version"),

				resource.TestCheckResourceAttr(datasourceName+"_by_db_system_shape", "db_system_shape", "BM.DenseIO2.52"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_shape", "db_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_shape", "db_versions.0.is_latest_for_major_version"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_shape", "db_versions.0.supports_pdb"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_shape", "db_versions.0.version"),

				resource.TestCheckResourceAttr(datasourceName+"_by_storage_management", "storage_management", "ASM"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_storage_management", "db_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_storage_management", "db_versions.0.is_latest_for_major_version"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_storage_management", "db_versions.0.supports_pdb"),
				resource.TestCheckResourceAttrSet(datasourceName+"_by_storage_management", "db_versions.0.version"),

				resource.TestCheckResourceAttr(datasourceName+"_by_is_upgrade_supported", "db_versions.0.is_upgrade_supported", "false"),
			),
		},
	})
}
