// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	dbVersionDataSourceRepresentationRequiredOnly = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}
	dbVersionDataSourceRepresentationWithDbSystemIdOptional = representationCopyWithNewProperties(dbVersionDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"db_system_id": Representation{repType: Optional, create: `${oci_database_db_system.test_db_system.id}`},
	})
	dbVersionDataSourceRepresentationWithDbSystemShapeOptional = representationCopyWithNewProperties(dbVersionDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"db_system_shape": Representation{repType: Optional, create: `BM.DenseIO1.36`},
	})

	DbVersionResourceConfig = DbSystemResourceConfig
)

func TestDatabaseDbVersionResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_versions.test_db_versions"

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
					compartmentIdVariableStr + DbVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "db_versions.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_versions.0.supports_pdb"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_versions.0.version"),

					resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_id", "db_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_id", "db_versions.#"),
					resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_id", "db_versions.0.supports_pdb"),
					resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_id", "db_versions.0.version"),

					resource.TestCheckResourceAttr(datasourceName+"_by_db_system_shape", "db_system_shape", "BM.DenseIO1.36"),
					resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_shape", "db_versions.#"),
					resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_shape", "db_versions.0.supports_pdb"),
					resource.TestCheckResourceAttrSet(datasourceName+"_by_db_system_shape", "db_versions.0.version"),
				),
			},
		},
	})
}
