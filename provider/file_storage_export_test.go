// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
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

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_export.test_export"
	datasourceName := "data.oci_file_storage_exports.test_exports"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckFileStorageExportDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + ExportPropertyVariables + compartmentIdVariableStr + ExportResourceConfig,
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
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckFileStorageExportDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).fileStorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_export" {
			noResourceFound = false
			request := oci_file_storage.GetExportRequest{}

			tmp := rs.Primary.ID
			request.ExportId = &tmp

			response, err := client.GetExport(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_file_storage.ExportLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
