// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	VnicAttachmentRequiredOnlyResource = VnicAttachmentResourceDependencies + `
resource "oci_core_vnic_attachment" "test_vnic_attachment" {
	#Required
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.test_subnet.id}"
	}
	instance_id = "${oci_core_instance.test_instance.id}"
}
`

	VnicAttachmentResourceConfig = VnicAttachmentResourceDependencies + `
resource "oci_core_vnic_attachment" "test_vnic_attachment" {
	#Required
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.test_subnet.id}"

		#Optional
		assign_public_ip = "${var.vnic_attachment_create_vnic_details_assign_public_ip}"
		display_name = "${var.vnic_attachment_create_vnic_details_display_name}"
		hostname_label = "${var.vnic_attachment_create_vnic_details_hostname_label}"
		private_ip = "${var.vnic_attachment_create_vnic_details_private_ip}"
		skip_source_dest_check = "${var.vnic_attachment_create_vnic_details_skip_source_dest_check}"
	}
	instance_id = "${oci_core_instance.test_instance.id}"

	#Optional
	display_name = "${var.vnic_attachment_display_name}"
	nic_index = "${var.vnic_attachment_nic_index}"
}
`
	VnicAttachmentPropertyVariables = `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = false }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = 10 }
variable "vnic_attachment_vnic_id" { default = "vnicId" }

`

	// TODO Replace these with actual definitions of test Instance properties and configs
	InstancePropertyVariables          = ""
	InstanceResourceConfig             = ""
	VnicAttachmentResourceDependencies = InstancePropertyVariables + InstanceResourceConfig
)

func TestCoreVnicAttachmentResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_vnic_attachment.test_vnic_attachment"
	datasourceName := "data.oci_core_vnic_attachments.test_vnic_attachments"

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
				Config:            config + VnicAttachmentPropertyVariables + compartmentIdVariableStr + VnicAttachmentRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VnicAttachmentResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + VnicAttachmentPropertyVariables + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain2" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = true }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName2" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel2" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp2" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = true }
variable "vnic_attachment_display_name" { default = "displayName2" }
variable "vnic_attachment_nic_index" { default = 11 }
variable "vnic_attachment_vnic_id" { default = "vnicId2" }

                ` + compartmentIdVariableStr2 + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "11"),
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
variable "vnic_attachment_availability_domain" { default = "availabilityDomain2" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = true }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName2" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel2" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp2" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = true }
variable "vnic_attachment_display_name" { default = "displayName2" }
variable "vnic_attachment_nic_index" { default = 11 }
variable "vnic_attachment_vnic_id" { default = "vnicId2" }

data "oci_core_vnic_attachments" "test_vnic_attachments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.vnic_attachment_availability_domain}"
	instance_id = "${oci_core_instance.test_instance.id}"
	vnic_id = "${var.vnic_attachment_vnic_id}"

    filter {
    	name = "id"
    	values = ["${oci_core_vnic_attachment.test_vnic_attachment.id}"]
    }
}
                ` + compartmentIdVariableStr2 + VnicAttachmentResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "vnic_id", "vnicId2"),

					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.0.nic_index", "11"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.time_created"),
				),
			},
		},
	})
}

func TestCoreVnicAttachmentResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_vnic_attachment.test_vnic_attachment"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + VnicAttachmentPropertyVariables + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
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
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = true }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = 10 }
variable "vnic_attachment_vnic_id" { default = "vnicId" }
				` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter AssignPublicIp but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = true }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName2" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = 10 }
variable "vnic_attachment_vnic_id" { default = "vnicId" }
				` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter DisplayName but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = true }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName2" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel2" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = 10 }
variable "vnic_attachment_vnic_id" { default = "vnicId" }
				` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
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
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = true }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName2" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel2" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp2" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = 10 }
variable "vnic_attachment_vnic_id" { default = "vnicId" }
				` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter PrivateIp but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = true }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName2" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel2" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp2" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = true }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = 10 }
variable "vnic_attachment_vnic_id" { default = "vnicId" }
				` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter SkipSourceDestCheck but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = true }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName2" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel2" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp2" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = true }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = 10 }
variable "vnic_attachment_vnic_id" { default = "vnicId" }
				` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
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

			{
				Config: config + `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = false }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName2" }
variable "vnic_attachment_nic_index" { default = 10 }
variable "vnic_attachment_vnic_id" { default = "vnicId" }
				` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter DisplayName but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = false }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName2" }
variable "vnic_attachment_nic_index" { default = 10 }
variable "vnic_attachment_vnic_id" { default = "vnicId" }
				` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter InstanceId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = false }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName2" }
variable "vnic_attachment_nic_index" { default = 11 }
variable "vnic_attachment_vnic_id" { default = "vnicId" }
				` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "11"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter NicIndex but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
