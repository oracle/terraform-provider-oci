// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	InstanceRequiredOnlyResource = InstanceResourceDependenciesRequiredOnly + `
resource "oci_core_instance" "test_instance" {
	#Required
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	compartment_id = "${var.compartment_id}"
	shape = "${var.instance_shape}"
}
`
	InstanceResourceAsDependencyConfig = InstanceResourceDependenciesRequiredOnly + `
resource "oci_core_instance" "test_instance" {
	#Required
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	compartment_id = "${var.compartment_id}"
	shape = "${var.instance_shape}"
	image = "${var.InstanceImageOCID[var.region]}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
	metadata {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}

	timeouts {
		create = "15m"
	}
}
`

	InstanceResourceConfig = InstanceResourceDependencies + `
resource "oci_core_instance" "test_instance" {
	#Required
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	compartment_id = "${var.compartment_id}"
	shape = "${var.instance_shape}"

	#Optional
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.test_subnet.id}"

		#Optional
		assign_public_ip = "${var.instance_create_vnic_details_assign_public_ip}"
		defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.instance_create_vnic_details_defined_tags_value}")}"
		display_name = "${var.instance_create_vnic_details_display_name}"
		freeform_tags = "${var.instance_create_vnic_details_freeform_tags}"
		hostname_label = "${var.instance_create_vnic_details_hostname_label}"
		private_ip = "${var.instance_create_vnic_details_private_ip}"
		skip_source_dest_check = "${var.instance_create_vnic_details_skip_source_dest_check}"
	}
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.instance_defined_tags_value}")}"
	display_name = "${var.instance_display_name}"
	extended_metadata = "${var.instance_extended_metadata}"
	freeform_tags = "${var.instance_freeform_tags}"
	hostname_label = "${var.instance_hostname_label}"
	ipxe_script = "${var.instance_ipxe_script}"
	metadata = "${var.instance_metadata}"
	source_details {
		#Required
		source_type = "image"
        source_id = "${var.InstanceImageOCID[var.region]}"
	}
	subnet_id = "${oci_core_subnet.test_subnet.id}"
}
`
	InstancePropertyVariables = `
variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
	  }
}

variable "instance_availability_domain" { default = "availabilityDomain" }
variable "instance_create_vnic_details_assign_public_ip" { default = false }
variable "instance_create_vnic_details_defined_tags_value" { default = "definedTags" }
variable "instance_create_vnic_details_display_name" { default = "displayName" }
variable "instance_create_vnic_details_freeform_tags" { default = "freeformTags" }
variable "instance_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "instance_create_vnic_details_private_ip" { default = "privateIp" }
variable "instance_create_vnic_details_skip_source_dest_check" { default = false }
variable "instance_defined_tags_value" { default = "value" }
variable "instance_display_name" { default = "displayName" }
variable "instance_extended_metadata" { default = "extendedMetadata" }
variable "instance_freeform_tags" { default = {"Department"= "Finance"} }
variable "instance_hostname_label" { default = "hostnameLabel" }
variable "instance_image" { default = "image" }
variable "instance_ipxe_script" { default = "ipxeScript" }
variable "instance_metadata" { default = "metadata" }
variable "instance_shape" { default = "VM.Standard1.8" }
variable "instance_source_details_source_type" { default = "sourceType" }
variable "instance_state" { default = "AVAILABLE" }

`
	InstanceResourceDependenciesRequiredOnly = SubnetPropertyVariables + SubnetRequiredOnlyResource
	InstanceResourceDependencies             = DefinedTagsDependencies + InstanceResourceDependenciesRequiredOnly
)

func TestCoreInstanceResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"

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
				Config:            config + InstancePropertyVariables + compartmentIdVariableStr + InstanceRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "availabilityDomain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard1.8"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + InstanceResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + InstancePropertyVariables + compartmentIdVariableStr + InstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "availabilityDomain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image", "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard1.8"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "sourceType"),
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
variable "instance_availability_domain" { default = "availabilityDomain" }
variable "instance_create_vnic_details_assign_public_ip" { default = false }
variable "instance_create_vnic_details_defined_tags_value" { default = "definedTags" }
variable "instance_create_vnic_details_display_name" { default = "displayName" }
variable "instance_create_vnic_details_freeform_tags" { default = "freeformTags" }
variable "instance_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "instance_create_vnic_details_private_ip" { default = "privateIp" }
variable "instance_create_vnic_details_skip_source_dest_check" { default = false }
variable "instance_defined_tags_value" { default = "updatedValue" }
variable "instance_display_name" { default = "displayName2" }
variable "instance_extended_metadata" { default = "extendedMetadata" }
variable "instance_freeform_tags" { default = {"Department"= "Accounting"} }
variable "instance_hostname_label" { default = "hostnameLabel" }
variable "instance_image" { default = "image" }
variable "instance_ipxe_script" { default = "ipxeScript" }
variable "instance_metadata" { default = "metadata" }
variable "instance_shape" { default = "VM.Standard1.8" }
variable "instance_source_details_source_type" { default = "sourceType" }
variable "instance_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + InstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "availabilityDomain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image", "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard1.8"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "sourceType"),
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
			// verify datasource
			{
				Config: config + `
variable "instance_availability_domain" { default = "availabilityDomain" }
variable "instance_create_vnic_details_assign_public_ip" { default = false }
variable "instance_create_vnic_details_defined_tags_value" { default = "definedTags" }
variable "instance_create_vnic_details_display_name" { default = "displayName" }
variable "instance_create_vnic_details_freeform_tags" { default = "freeformTags" }
variable "instance_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "instance_create_vnic_details_private_ip" { default = "privateIp" }
variable "instance_create_vnic_details_skip_source_dest_check" { default = false }
variable "instance_defined_tags_value" { default = "updatedValue" }
variable "instance_display_name" { default = "displayName2" }
variable "instance_extended_metadata" { default = "extendedMetadata" }
variable "instance_freeform_tags" { default = {"Department"= "Accounting"} }
variable "instance_hostname_label" { default = "hostnameLabel" }
variable "instance_image" { default = "image" }
variable "instance_ipxe_script" { default = "ipxeScript" }
variable "instance_metadata" { default = "metadata" }
variable "instance_shape" { default = "VM.Standard1.8" }
variable "instance_source_details_source_type" { default = "sourceType" }
variable "instance_state" { default = "AVAILABLE" }

data "oci_core_instances" "test_instances" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.instance_availability_domain}"
	display_name = "${var.instance_display_name}"
	state = "${var.instance_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_instance.test_instance.id}"]
    }
}
                ` + compartmentIdVariableStr + InstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.availability_domain", "availabilityDomain"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.extended_metadata.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.image", "image"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.metadata.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard1.8"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.0.source_type", "sourceType"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.time_created"),
				),
			},
		},
	})
}
