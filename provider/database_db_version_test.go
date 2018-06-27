// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	DbVersionResourceConfig = DbVersionResourceDependencies + `

`
	DbVersionPropertyVariables = `
variable "db_version_db_system_shape" { default = "BM.DenseIO1.36" }

`
	DbVersionResourceDependencies = DbHomePatchResourceDependencies
)

func TestDatabaseDbVersionResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_versions.test_db_versions"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `
variable "db_version_db_system_shape" { default = "BM.DenseIO1.36" }

data "oci_database_db_versions" "test_db_versions" {
	#Required
	compartment_id = "${var.compartment_id}"
}

data "oci_database_db_versions" "test_db_versions_by_db_system_id" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}

data "oci_database_db_versions" "test_db_versions_by_db_system_shape" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	db_system_shape = "${var.db_version_db_system_shape}"
}
                ` + compartmentIdVariableStr + DbVersionResourceConfig,
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
