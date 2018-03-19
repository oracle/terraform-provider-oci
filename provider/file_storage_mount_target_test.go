// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	MountTargetRequiredOnlyResource = MountTargetResourceDependencies + `
resource "oci_file_storage_mount_target" "test_mount_target" {
	#Required
	availability_domain = "${var.mount_target_availability_domain}"
	compartment_id = "${var.compartment_id}"
	subnet_id = "${oci_file_storage_subnet.test_subnet.id}"
}
`

	MountTargetResourceConfig = MountTargetResourceDependencies + `
resource "oci_file_storage_mount_target" "test_mount_target" {
	#Required
	availability_domain = "${var.mount_target_availability_domain}"
	compartment_id = "${var.compartment_id}"
	subnet_id = "${oci_file_storage_subnet.test_subnet.id}"

	#Optional
	display_name = "${var.mount_target_display_name}"
	hostname_label = "${var.mount_target_hostname_label}"
	ip_address = "${var.mount_target_ip_address}"
}
`
	MountTargetPropertyVariables = `
variable "mount_target_availability_domain" { default = "pWEh:PHX-AD-2" }
variable "mount_target_display_name" { default = "mount-target-5" }
variable "mount_target_export_set_id" { default = "exportSetId" }
variable "mount_target_hostname_label" { default = "hostnameLabel" }
variable "mount_target_id" { default = "id" }
variable "mount_target_ip_address" { default = "ipAddress" }
variable "mount_target_state" { default = "state" }

`
	MountTargetResourceDependencies = SubnetPropertyVariables + SubnetResourceConfig
)

func TestFileStorageMountTargetResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_file_storage_mount_target.test_mount_target"
	datasourceName := "data.oci_file_storage_mount_targets.test_mount_targets"

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
				Config:            config + MountTargetPropertyVariables + compartmentIdVariableStr + MountTargetRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "pWEh:PHX-AD-2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + MountTargetResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + MountTargetPropertyVariables + compartmentIdVariableStr + MountTargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "pWEh:PHX-AD-2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "lifecycle_details"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
variable "mount_target_availability_domain" { default = "pWEh:PHX-AD-2" }
variable "mount_target_display_name" { default = "displayName2" }
variable "mount_target_export_set_id" { default = "exportSetId" }
variable "mount_target_hostname_label" { default = "hostnameLabel" }
variable "mount_target_id" { default = "id" }
variable "mount_target_ip_address" { default = "ipAddress" }
variable "mount_target_state" { default = "state" }

                ` + compartmentIdVariableStr + MountTargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "pWEh:PHX-AD-2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "lifecycle_details"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
			// verify updates to Force New parameters.
			{
				Config: config + `
variable "mount_target_availability_domain" { default = "availabilityDomain2" }
variable "mount_target_display_name" { default = "displayName2" }
variable "mount_target_export_set_id" { default = "exportSetId2" }
variable "mount_target_hostname_label" { default = "hostnameLabel2" }
variable "mount_target_id" { default = "id2" }
variable "mount_target_ip_address" { default = "ipAddress2" }
variable "mount_target_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr2 + MountTargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress2"),
					resource.TestCheckResourceAttrSet(resourceName, "lifecycle_details"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
variable "mount_target_availability_domain" { default = "availabilityDomain2" }
variable "mount_target_display_name" { default = "displayName2" }
variable "mount_target_export_set_id" { default = "exportSetId2" }
variable "mount_target_hostname_label" { default = "hostnameLabel2" }
variable "mount_target_id" { default = "id2" }
variable "mount_target_ip_address" { default = "ipAddress2" }
variable "mount_target_state" { default = "AVAILABLE" }

data "oci_file_storage_mount_targets" "test_mount_targets" {
	#Required
	availability_domain = "${var.mount_target_availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.mount_target_display_name}"
	export_set_id = "${var.mount_target_export_set_id}"
	id = "${var.mount_target_id}"
	state = "${var.mount_target_state}"

    filter {
    	name = "id"
    	values = ["${oci_file_storage_mount_target.test_mount_target.id}"]
    }
}
                ` + compartmentIdVariableStr2 + MountTargetResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "export_set_id", "exportSetId2"),
					resource.TestCheckResourceAttr(datasourceName, "id", "id2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(datasourceName, "mount_targets.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.lifecycle_details"),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.private_ip_ids"),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.time_created"),
				),
			},
		},
	})
}

func TestFileStorageMountTargetResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_file_storage_mount_target.test_mount_target"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + MountTargetPropertyVariables + compartmentIdVariableStr + MountTargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "pWEh:PHX-AD-2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "lifecycle_details"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "mount_target_availability_domain" { default = "availabilityDomain2" }
variable "mount_target_display_name" { default = "mount-target-5" }
variable "mount_target_export_set_id" { default = "exportSetId" }
variable "mount_target_hostname_label" { default = "hostnameLabel" }
variable "mount_target_id" { default = "id" }
variable "mount_target_ip_address" { default = "ipAddress" }
variable "mount_target_state" { default = "state" }
				` + compartmentIdVariableStr + MountTargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "lifecycle_details"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
variable "mount_target_availability_domain" { default = "availabilityDomain2" }
variable "mount_target_display_name" { default = "mount-target-5" }
variable "mount_target_export_set_id" { default = "exportSetId" }
variable "mount_target_hostname_label" { default = "hostnameLabel" }
variable "mount_target_id" { default = "id" }
variable "mount_target_ip_address" { default = "ipAddress" }
variable "mount_target_state" { default = "state" }
				` + compartmentIdVariableStr2 + MountTargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "lifecycle_details"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
variable "mount_target_availability_domain" { default = "availabilityDomain2" }
variable "mount_target_display_name" { default = "mount-target-5" }
variable "mount_target_export_set_id" { default = "exportSetId" }
variable "mount_target_hostname_label" { default = "hostnameLabel2" }
variable "mount_target_id" { default = "id" }
variable "mount_target_ip_address" { default = "ipAddress" }
variable "mount_target_state" { default = "state" }
				` + compartmentIdVariableStr2 + MountTargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "lifecycle_details"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter HostnameLabel but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "mount_target_availability_domain" { default = "availabilityDomain2" }
variable "mount_target_display_name" { default = "mount-target-5" }
variable "mount_target_export_set_id" { default = "exportSetId" }
variable "mount_target_hostname_label" { default = "hostnameLabel2" }
variable "mount_target_id" { default = "id" }
variable "mount_target_ip_address" { default = "ipAddress2" }
variable "mount_target_state" { default = "state" }
				` + compartmentIdVariableStr2 + MountTargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress2"),
					resource.TestCheckResourceAttrSet(resourceName, "lifecycle_details"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter IpAddress but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "mount_target_availability_domain" { default = "availabilityDomain2" }
variable "mount_target_display_name" { default = "mount-target-5" }
variable "mount_target_export_set_id" { default = "exportSetId" }
variable "mount_target_hostname_label" { default = "hostnameLabel2" }
variable "mount_target_id" { default = "id" }
variable "mount_target_ip_address" { default = "ipAddress2" }
variable "mount_target_state" { default = "state" }
				` + compartmentIdVariableStr2 + MountTargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress2"),
					resource.TestCheckResourceAttrSet(resourceName, "lifecycle_details"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter SubnetId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
