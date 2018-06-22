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
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.volume_group_defined_tags_value}")}"
	display_name = "${var.volume_group_display_name}"
	freeform_tags = "${var.volume_group_freeform_tags}"
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
	VolumeGroupPropertyVariables = `
variable "volume_group_defined_tags_value" { default = "value" }
variable "volume_group_display_name" { default = "displayName" }
variable "volume_group_freeform_tags" { default = {"Department"= "Finance"} }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }

`
	VolumeGroupResourceDependencies = DefinedTagsDependencies + `
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
)

func TestCoreVolumeGroupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

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
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
variable "volume_group_defined_tags_value" { default = "updatedValue" }
variable "volume_group_display_name" { default = "displayName2" }
variable "volume_group_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + VolumeGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
			// verify that the change in order of the volume ids doesn't result in a new resource
			{
				Config: config + `
variable "volume_group_display_name" { default = "displayName2" }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + VolumeGroupResourceConfigJumbledVolumeIds,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// verify that the change in list of volume ids does cause a change in the plan
			{
				Config: config + `
variable "volume_group_display_name" { default = "displayName2" }
variable "volume_group_source_details_type" { default = "volumeIds" }
variable "volume_group_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + VolumeGroupResourceConfigSingleVolumeId,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// verify datasource
			{
				Config: config + `
variable "volume_group_defined_tags_value" { default = "updatedValue" }
variable "volume_group_display_name" { default = "displayName2" }
variable "volume_group_freeform_tags" { default = {"Department"= "Accounting"} }
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
                ` + compartmentIdVariableStr + VolumeGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "volume_groups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.freeform_tags.%", "1"),
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
