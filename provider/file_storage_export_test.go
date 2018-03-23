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
	export_set_id = "${oci_file_storage_mount_target.test_mount_target.export_set_id}"
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	path = "${var.export_path}"
}
`
	ExportPropertyVariables = `
variable "export_id" { default = "id" }
variable "export_path" { default = "/files-5" }
variable "export_state" { default = "ACTIVE" }

`
	ExportResourceDependencies = FileSystemPropertyVariables + FileSystemResourceConfig + MountTargetPropertyVariables + MountTargetResourceConfig
)

func TestFileStorageExportResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_file_storage_export.test_export"
	datasourceName := "data.oci_file_storage_exports.test_exports"

	var resId, resId2 string

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

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `
variable "export_id" { default = "id2" }
variable "export_path" { default = "/files-6" }
variable "export_state" { default = "ACTIVE" }

                ` + compartmentIdVariableStr2 + ExportResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
					resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "path", "/files-6"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "export_path" { default = "/files-6" }
variable "export_state" { default = "ACTIVE" }

data "oci_file_storage_exports" "test_exports" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	export_set_id = "${oci_file_storage_mount_target.test_mount_target.export_set_id}"
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"

    filter {
    	name = "id"
    	values = ["${oci_file_storage_export.test_export.id}"]
    }
}
                ` + compartmentIdVariableStr2 + ExportResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "exports.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "exports.0.export_set_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "exports.0.file_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "exports.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "exports.0.path", "/files-6"),
					resource.TestCheckResourceAttr(datasourceName, "exports.0.state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "exports.0.time_created"),
				),
			},
		},
	})
}
