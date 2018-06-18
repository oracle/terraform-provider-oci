// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	VolumeBackupPolicyAssignmentResourceConfig = VolumeBackupPolicyAssignmentResourceDependencies + `
data "oci_core_volume_backup_policies" "test_volume_backup_policies" {
	filter {
        name = "display_name"
        values = [ "silver" ]
    }
}

resource "oci_core_volume_backup_policy_assignment" "test_volume_backup_policy_assignment" {
	asset_id = "${oci_core_volume.test_volume.id}"
	policy_id = "${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}"
}
`
	VolumeBackupPolicyAssignmentPropertyVariables = `

`
	VolumeBackupPolicyAssignmentResourceDependencies = VolumePropertyVariables + VolumeResourceConfig
)

func TestCoreVolumeBackupPolicyAssignmentResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_backup_policy_assignment.test_volume_backup_policy_assignment"
	datasourceName := "data.oci_core_volume_backup_policy_assignments.test_volume_backup_policy_assignments"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + VolumeBackupPolicyAssignmentPropertyVariables + compartmentIdVariableStr + VolumeBackupPolicyAssignmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "asset_id"),
					resource.TestCheckResourceAttrSet(resourceName, "policy_id"),
				),
			},

			// verify datasource
			{
				Config: config + `

data "oci_core_volume_backup_policy_assignments" "test_volume_backup_policy_assignments" {
	#Required
	asset_id = "${oci_core_volume.test_volume.id}"

    filter {
        name = "id"
        values = ["${oci_core_volume_backup_policy_assignment.test_volume_backup_policy_assignment.id}"]
    }
}
                ` + compartmentIdVariableStr + VolumeBackupPolicyAssignmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "asset_id"),

					resource.TestCheckResourceAttr(datasourceName, "volume_backup_policy_assignments.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policy_assignments.0.asset_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policy_assignments.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policy_assignments.0.policy_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policy_assignments.0.time_created"),
				),
			},
		},
	})
}
