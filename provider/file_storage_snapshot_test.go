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
variable "snapshot_state" { default = "state" }

`
	SnapshotResourceDependencies = FileSystemPropertyVariables + FileSystemResourceConfig
)

func TestFileStorageSnapshotResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_file_storage_snapshot.test_snapshot"
	datasourceName := "data.oci_file_storage_snapshots.test_snapshots"

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
				Config:            config + SnapshotPropertyVariables + compartmentIdVariableStr + SnapshotResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "snapshot-1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `
variable "snapshot_name" { default = "name2" }
variable "snapshot_state" { default = "ACTIVE" }

                ` + compartmentIdVariableStr2 + SnapshotResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
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
variable "snapshot_name" { default = "name2" }
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
                ` + compartmentIdVariableStr2 + SnapshotResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "file_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "snapshots.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.file_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "snapshots.0.name", "name2"),
					resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.time_created"),
				),
			},
		},
	})
}
