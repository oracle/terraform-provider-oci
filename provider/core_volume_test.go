// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	VolumeRequiredOnlyResource = VolumeResourceDependencies + `
resource "oci_core_volume" "test_volume" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
}
`

	VolumeResourceConfig = VolumeResourceDependencies + `
resource "oci_core_volume" "test_volume" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.volume_defined_tags_value}")}"
	display_name = "${var.volume_display_name}"
	freeform_tags = "${var.volume_freeform_tags}"
	size_in_gbs = "${var.volume_size_in_gbs}"
	source_details {
		#Required
		type = "${var.volume_source_details_type}"
		id = "${oci_core_volume.source_volume.id}"
	}
}
`
	VolumePropertyVariables = `
variable "volume_defined_tags_value" { default = "value" }
variable "volume_display_name" { default = "displayName" }
variable "volume_freeform_tags" { default = {"Department"= "Finance"} }
variable "volume_size_in_gbs" { default = 50 }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }

`
	// Uncomment once defined: VolumeBackupPropertyVariables + VolumeBackupResourceConfig
	VolumeResourceDependencies = DefinedTagsDependencies + `
data "oci_identity_availability_domains" "ADs" {
	compartment_id = "${var.compartment_id}"
}

resource "oci_core_volume" "source_volume" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
}
`
)

func TestCoreVolumeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume.test_volume"
	datasourceName := "data.oci_core_volumes.test_volumes"

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
				Config:            config + VolumePropertyVariables + compartmentIdVariableStr + VolumeRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + VolumePropertyVariables + compartmentIdVariableStr + VolumeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttr(resourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volume"),
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
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = 50 }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + VolumeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttr(resourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volume"),
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
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = 50 }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }

data "oci_core_volumes" "test_volumes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	display_name = "${var.volume_display_name}"
	state = "${var.volume_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_volume.test_volume.id}"]
    }
}
                ` + compartmentIdVariableStr + VolumeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "state"),

					resource.TestCheckResourceAttr(datasourceName, "volumes.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.size_in_gbs", "50"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.size_in_mbs", "51200"),
					resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.time_created"),
				),
			},
		},
	})
}
