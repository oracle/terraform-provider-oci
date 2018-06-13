// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ExportResourceConfig = ExportResourceDependencies + `
resource "oci_file_storage_export" "test_export" {
	#Required
	export_set_id = "${oci_file_storage_export_set.test_export_set.id}"
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	path = "${var.export_path}"
}
`
	ExportPropertyVariables = `
variable "export_path" { default = "/files-5" }
variable "export_state" { default = "ACTIVE" }

`
	ExportResourceDependencies = FileSystemPropertyVariables + FileSystemResourceConfigOnly + ExportSetPropertyVariables + ExportSetResourceConfig
)

func TestFileStorageExportResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_export.test_export"
	datasourceName := "data.oci_file_storage_exports.test_exports"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + ExportPropertyVariables + compartmentIdVariableStr + ExportResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
					resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
					resource.TestCheckResourceAttr(resourceName, "path", "/files-5"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "export_path" { default = "/files-5" }
variable "export_state" { default = "ACTIVE" }

data "oci_file_storage_exports" "test_exports" {

	#Optional
	compartment_id = "${var.compartment_id}"
	export_set_id = "${oci_file_storage_export_set.test_export_set.id}"
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"

    filter {
    	name = "id"
    	values = ["${oci_file_storage_export.test_export.id}"]
    }
}
                ` + compartmentIdVariableStr + ExportResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "exports.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "exports.0.export_set_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "exports.0.file_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "exports.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "exports.0.path", "/files-5"),
					resource.TestCheckResourceAttr(datasourceName, "exports.0.state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "exports.0.time_created"),
				),
			},
		},
	})
}
