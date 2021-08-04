// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	dbVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                       Representation{repType: Required, create: `${var.compartment_id}`},
		"db_system_id":                         Representation{repType: Optional, create: `${oci_database_db_system.test_db_system.id}`},
		"db_system_shape":                      Representation{repType: Optional, create: `BM.DenseIO2.52`},
		"is_database_software_image_supported": Representation{repType: Optional, create: `false`},
		"is_upgrade_supported":                 Representation{repType: Optional, create: `false`},
		"storage_management":                   Representation{repType: Optional, create: `ASM`},
	}
	dbVersionDataSourceRepresentationRequiredOnly = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}
	dbVersionDataSourceRepresentationWithDbSystemIdOptional = representationCopyWithNewProperties(dbVersionDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"db_system_id": Representation{repType: Optional, create: `${oci_database_db_system.test_db_system.id}`},
	})
	dbVersionDataSourceRepresentationWithUpgradeSupportedOptional = representationCopyWithNewProperties(dbVersionDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"db_system_shape":      Representation{repType: Optional, create: `BM.DenseIO2.52`},
		"is_upgrade_supported": Representation{repType: Optional, create: `false`},
	})
	dbVersionDataSourceRepresentationWithDbSystemShapeOptional = representationCopyWithNewProperties(dbVersionDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"db_system_shape": Representation{repType: Optional, create: `BM.DenseIO2.52`},
	})
	dbVersionDataSourceRepresentationWithStorageManagementOptional = representationCopyWithNewProperties(dbVersionDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"storage_management": Representation{repType: Optional, create: `ASM`},
	})
	DbVersionResourceConfig = DbSystemResourceConfig
)

// issue-routing-tag: database/default
func TestDatabaseDbVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbVersionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_versions.test_db_versions"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_db_versions", "test_db_versions", Required, Create, dbVersionDataSourceRepresentationRequiredOnly) +
					generateDataSourceFromRepresentationMap("oci_database_db_versions", "test_db_versions_by_db_system_id", Optional, Create, dbVersionDataSourceRepresentationWithDbSystemIdOptional) +
					generateDataSourceFromRepresentationMap("oci_database_db_versions", "test_db_versions_by_db_system_shape", Optional, Create, dbVersionDataSourceRepresentationWithDbSystemShapeOptional) +
					generateDataSourceFromRepresentationMap("oci_database_db_versions", "test_db_versions_by_is_upgrade_supported", Optional, Create, dbVersionDataSourceRepresentationWithUpgradeSupportedOptional) +
					generateDataSourceFromRepresentationMap("oci_database_db_versions", "test_db_versions_by_storage_management", Optional, Create, dbVersionDataSourceRepresentationWithStorageManagementOptional) +
					compartmentIdVariableStr + DbVersionResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
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
		},
	})
}
