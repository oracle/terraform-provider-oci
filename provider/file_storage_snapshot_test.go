// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	SnapshotResourceConfig = SnapshotResourceDependencies + `
resource "oci_file_storage_snapshot" "test_snapshot" {
	#Required
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	name = "${var.snapshot_name}"
}
`
	SnapshotPropertyVariables = `
variable "snapshot_name" { default = "snapshot-1" }
variable "snapshot_state" { default = "ACTIVE" }

`
	SnapshotResourceDependencies = FileSystemPropertyVariables + FileSystemResourceConfig
)

func TestFileStorageSnapshotResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_snapshot.test_snapshot"
	datasourceName := "data.oci_file_storage_snapshots.test_snapshots"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + SnapshotPropertyVariables + compartmentIdVariableStr + SnapshotResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "snapshot-1"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "snapshot_name" { default = "snapshot-1" }
variable "snapshot_state" { default = "ACTIVE" }

data "oci_file_storage_snapshots" "test_snapshots" {
	#Required
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"

	#Optional
	id = "${oci_file_storage_snapshot.test_snapshot.id}"
	state = "${var.snapshot_state}"

    filter {
    	name = "id"
    	values = ["${oci_file_storage_snapshot.test_snapshot.id}"]
    }
}
                ` + compartmentIdVariableStr + SnapshotResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "file_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "snapshots.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.file_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "snapshots.0.name", "snapshot-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.time_created"),
				),
			},
		},
	})
}
