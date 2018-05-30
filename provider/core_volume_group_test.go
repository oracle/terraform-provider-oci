// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	VolumeGroupRequiredOnlyResource = VolumeGroupResourceDependencies + `
resource "oci_core_volume_group" "test_volume_group" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	source_details {
		#Required
		type = "${var.volume_group_source_details_type}"
		volume_ids = ["${oci_core_volume.source_volume_list.*.id}"]
	}
}
`
	VolumeGroupNewADResourceConfig = VolumeGroupNewADResourceDependencies + `
resource "oci_core_volume_group" "test_volume_group" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.1.name}"
	compartment_id = "${var.compartment_id}"
	source_details {
		#Required
		type = "${var.volume_group_source_details_type}"
		volume_ids = ["${oci_core_volume.source_volume.id}"]
	}
	#Optional
	display_name = "${var.volume_group_display_name}"
}
`
	VolumeGroupSourceDetailsConfig = `
	source_details {
		#Required
		type = "${var.volume_group_source_details_type}"
		volume_ids = ["${oci_core_volume.source_volume_list.*.id}"]
	}
`
	VolumeGroupSourceDetailsConfigJumbledVolumeIds = `
	source_details {
		#Required
		type = "${var.volume_group_source_details_type}"
		volume_ids = ["${oci_core_volume.source_volume_list.*.id[1]}", "${oci_core_volume.source_volume_list.*.id[0]}"]
	}
`
	VolumeGroupSourceDetailsConfigSingleVolumeIdSourceDetails = `
	source_details {
		#Required
		type = "${var.volume_group_source_details_type}"
		volume_ids = ["${oci_core_volume.source_volume_list.*.id[1]}"]
	}
`
	VolumeGroupResourceConfig = VolumeGroupResourceDependencies + `
resource "oci_core_volume_group" "test_volume_group" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
` + VolumeGroupSourceDetailsConfig + `
	#Optional
	display_name = "${var.volume_group_display_name}"
}
`
	VolumeGroupResourceConfigJumbledVolumeIds = VolumeGroupResourceDependencies + `
resource "oci_core_volume_group" "test_volume_group" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
` + VolumeGroupSourceDetailsConfigJumbledVolumeIds + `
	#Optional
	display_name = "${var.volume_group_display_name}"
}
`
	VolumeGroupResourceConfigSingleVolumeId = VolumeGroupResourceDependencies + `
resource "oci_core_volume_group" "test_volume_group" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
` + VolumeGroupSourceDetailsConfigSingleVolumeIdSourceDetails + `
	#Optional
	display_name = "${var.volume_group_display_name}"
}
`
	VolumeGroupWithVolumeGroupSourceResourceConfig = VolumeGroupWithVolumeGroupSourceResourceDependencies + `
resource "oci_core_volume_group" "test_volume_group" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	source_details {
		#Required
		type = "${var.volume_group_source_details_type}"
		volume_group_id = "${oci_core_volume_group.source_volume_group.id}"
	}

	#Optional
	display_name = "${var.volume_group_display_name}"
}
`
	VolumeGroupWithVolumeGroupBackupSourceResourceConfig = VolumeGroupWithVolumeGroupBackupSourceResourceDependencies + `
resource "oci_core_volume_group" "test_volume_group" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	source_details {
		#Required
		type = "${var.volume_group_source_details_type}"
		volume_group_backup_id = "${oci_core_volume_group_backup.source_volume_group_backup.id}"
	}

	#Optional
	display_name = "${var.volume_group_display_name}"
}
`
	VolumeGroupPropertyVariables = `
variable "volume_group_display_name" { default = "displayName" }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }

`
	VolumeGroupResourceDependencies = `
data "oci_identity_availability_domains" "ADs" {
	compartment_id = "${var.compartment_id}"
}

resource "oci_core_volume" "source_volume_list" {
	count = 2
	display_name = "${format("source-volume-%d", count.index + 1)}"

	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
}
`
	VolumeGroupWithVolumeGroupSourceResourceDependencies = `
data "oci_identity_availability_domains" "ADs" {
	compartment_id = "${var.compartment_id}"
}

resource "oci_core_volume_group" "source_volume_group" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	source_details {
		#Required
		type = "volumeIds" #Hardcoded for the source dependency
		volume_ids = ["${oci_core_volume.source_volume.id}"]
	}

	#Optional
	display_name = "source-volume-group"
}
resource "oci_core_volume" "source_volume" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "source-volume-for-source-volume-group"
}
`
	VolumeGroupWithVolumeGroupBackupSourceResourceDependencies = `
data "oci_identity_availability_domains" "ADs" {
	compartment_id = "${var.compartment_id}"
}

resource "oci_core_volume_group" "source_volume_group" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	source_details {
		#Required
		type = "volumeIds" #Hardcoded for the source dependency
		volume_ids = ["${oci_core_volume.source_volume.id}"]
	}

	#Optional
	display_name = "source-volume-group"
}
resource "oci_core_volume" "source_volume" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "source-volume-for-source-volume-group"
}
resource "oci_core_volume_group_backup" "source_volume_group_backup" {
  #Required
  volume_group_id = "${oci_core_volume_group.source_volume_group.id}"

  #Optional
  display_name = "source-volume-group-backup"
  type = "INCREMENTAL"
}
`
	VolumeGroupNewADResourceDependencies = `
data "oci_identity_availability_domains" "ADs" {
	compartment_id = "${var.compartment_id}"
}

resource "oci_core_volume" "source_volume" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.1.name}"
	compartment_id = "${var.compartment_id}"
}
`
)

func TestCoreVolumeGroupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_volume_group.test_volume_group"
	datasourceName := "data.oci_core_volume_groups.test_volume_groups"

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
				Config:            config + VolumeGroupPropertyVariables + compartmentIdVariableStr + VolumeGroupRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// We need to assert that the volume group created above did cause the source volume to have the volume
			// group id property populated correctly. Since the TF framework doesn't have a RefreshOnly directive, we are
			// using PlanOnly to trigger a refresh, and then assert on the value
			{
				Config:   config + VolumeGroupPropertyVariables + compartmentIdVariableStr + VolumeGroupRequiredOnlyResource,
				PlanOnly: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("oci_core_volume.source_volume_list.0", "volume_group_id"),
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeGroupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + VolumeGroupPropertyVariables + compartmentIdVariableStr + VolumeGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "volume_group_availability_domain" { default = "availabilityDomain" }
variable "volume_group_display_name" { default = "displayName2" }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + VolumeGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates to Force New parameters.
			{
				Config: config + `
variable "volume_group_display_name" { default = "displayName2" }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr2 + VolumeGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify that the change in order of the volume ids doesn't result in a new resource
			{
				Config: config + `
variable "volume_group_display_name" { default = "displayName2" }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr2 + VolumeGroupResourceConfigJumbledVolumeIds,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// verify that the change in list of volume ids does cause a change in the plan
			{
				Config: config + `
variable "volume_group_display_name" { default = "displayName2" }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr2 + VolumeGroupResourceConfigSingleVolumeId,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// verify datasource
			{
				Config: config + `
variable "volume_group_display_name" { default = "displayName2" }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }

data "oci_core_volume_groups" "test_volume_groups" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	display_name = "${var.volume_group_display_name}"
	state = "${var.volume_group_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_volume_group.test_volume_group.id}"]
    }
}
                ` + compartmentIdVariableStr2 + VolumeGroupResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "volume_groups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.size_in_mbs"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.source_details.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),
				),
			},
		},
	})
}

func TestCoreVolumeGroupResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_volume_group.test_volume_group"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + VolumeGroupPropertyVariables + compartmentIdVariableStr + VolumeGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "volume_group_display_name" { default = "displayName" }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }
				` + compartmentIdVariableStr + VolumeGroupNewADResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter AvailabilityDomain but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "volume_group_display_name" { default = "displayName" }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }
				` + compartmentIdVariableStr2 + VolumeGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter CompartmentId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "volume_group_display_name" { default = "displayName" }
variable "volume_group_source_details_type" { default = "volumeGroup" }
variable "volume_group_state" { default = "AVAILABLE" }
				` + compartmentIdVariableStr2 + VolumeGroupWithVolumeGroupSourceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter Type but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
			// Check for volume group sourced from volume group backup
			{
				Config: config + `
variable "volume_group_display_name" { default = "displayName" }
variable "volume_group_source_details_type" { default = "volumeGroupBackup" }
variable "volume_group_state" { default = "AVAILABLE" }
				` + compartmentIdVariableStr2 + VolumeGroupWithVolumeGroupBackupSourceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeGroupBackup"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter Type but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
