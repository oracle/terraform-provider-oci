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
	ExportRequiredOnlyResource = ExportResourceDependencies + `
resource "oci_file_storage_export" "test_export" {
	#Required
	export_set_id = "${oci_file_storage_export_set.test_export_set.id}"
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	path = "${var.export_path}"
}
`

	ExportResourceConfig = ExportResourceDependencies + `
resource "oci_file_storage_export" "test_export" {
	#Required
	export_set_id = "${oci_file_storage_export_set.test_export_set.id}"
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	path = "${var.export_path}"

	#Optional
	export_options {
		#Required
		source = "${var.export_export_options_source}"

		#Optional
		access = "${var.export_export_options_access}"
		anonymous_gid = "${var.export_export_options_anonymous_gid}"
		anonymous_uid = "${var.export_export_options_anonymous_uid}"
		identity_squash = "${var.export_export_options_identity_squash}"
		require_privileged_source_port = "${var.export_export_options_require_privileged_source_port}"
	}
}
`
	ExportPropertyVariables = `
variable "export_export_options_access" { default = "READ_WRITE" }
variable "export_export_options_anonymous_gid" { default = 10 }
variable "export_export_options_anonymous_uid" { default = 10 }
variable "export_export_options_identity_squash" { default = "NONE" }
variable "export_export_options_require_privileged_source_port" { default = false }
variable "export_export_options_source" { default = "0.0.0.0/0" }
variable "export_id" { default = "id" }
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

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckFileStorageExportDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + ExportPropertyVariables + compartmentIdVariableStr + ExportRequiredOnlyResource,
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

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ExportResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + ExportPropertyVariables + compartmentIdVariableStr + ExportResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "export_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.access", "READ_WRITE"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.anonymous_gid", "10"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.anonymous_uid", "10"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.identity_squash", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.require_privileged_source_port", "false"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.source", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
					resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "path", "/files-5"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "export_export_options_access" { default = "READ_ONLY" }
variable "export_export_options_anonymous_gid" { default = 11 }
variable "export_export_options_anonymous_uid" { default = 11 }
variable "export_export_options_identity_squash" { default = "ALL" }
variable "export_export_options_require_privileged_source_port" { default = true }
variable "export_export_options_source" { default = "0.0.0.0/0" }
variable "export_id" { default = "id" }
variable "export_path" { default = "/files-5" }
variable "export_state" { default = "ACTIVE" }

                ` + compartmentIdVariableStr + ExportResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "export_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.access", "READ_ONLY"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.anonymous_gid", "11"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.anonymous_uid", "11"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.identity_squash", "ALL"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.require_privileged_source_port", "true"),
					resource.TestCheckResourceAttr(resourceName, "export_options.0.source", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
					resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "path", "/files-5"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "export_export_options_access" { default = "READ_ONLY" }
variable "export_export_options_anonymous_gid" { default = 11 }
variable "export_export_options_anonymous_uid" { default = 11 }
variable "export_export_options_identity_squash" { default = "ALL" }
variable "export_export_options_require_privileged_source_port" { default = true }
variable "export_export_options_source" { default = "0.0.0.0/0" }
variable "export_id" { default = "id" }
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
