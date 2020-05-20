// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	CoreInstanceResourceConfig = OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies + `
resource "oci_core_instance" "test_instance" {
	#Required
	availability_domain = "${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}"
	compartment_id = "${var.compartment_id}"
	shape = "${var.instance_shape}"

	#Optional
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.test_subnet.id}"
        skip_source_dest_check = false

		#Optional
		defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.instance_create_vnic_details_defined_tags_value}")}"
		display_name = "${var.instance_create_vnic_details_display_name}"
		freeform_tags = "${var.instance_create_vnic_details_freeform_tags}"
	}
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.instance_defined_tags_value}")}"
	display_name = "${var.instance_display_name}"
	extended_metadata = "${var.instance_extended_metadata}"
	freeform_tags = "${var.instance_freeform_tags}"
	metadata = "${var.instance_metadata}"
	source_details {
		#Required
		source_type = "image"
        source_id = "${var.InstanceImageOCID[var.region]}"
	}
	subnet_id = "${oci_core_subnet.test_subnet.id}"
}
`
)

func TestResourceCoreApplyFiltersIntegration_basic(t *testing.T) {
	httpreplay.SetScenario("TestApplyFiltersIntegration_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource id
			{
				Config: config + `
variable "instance_create_vnic_details_defined_tags_value" { 
	default = "definedTags" 
}
variable "instance_create_vnic_details_display_name" { 
	default = "displayName" 
}
variable "instance_create_vnic_details_freeform_tags" { 
	default = {
		"Department"= "Accounting"
	} 
}
variable "instance_create_vnic_details_hostname_label" { 
	default = "hostnameLabel" 
}
variable "instance_defined_tags_value" { 
	default = "updatedValue" 
}
variable "instance_display_name" { 
	default = "displayName2" 
}
variable "instance_extended_metadata" { 
	default = {
				keyA = "valA"
				keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
	} 
}
variable "instance_freeform_tags" { 
	default = {
		"Department"= "Accounting"
	} 
}
variable "instance_hostname_label" { 
	default = "hostnameLabel" 
}
variable "instance_image" { 
	default = "image" 
}
variable "instance_ipxe_script" { 
	default = "ipxeScript" 
}
variable "instance_metadata" { 
	default = {
				ssh_authorized_keys = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
				user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
	}
}
variable "instance_shape" { 
	default = "VM.Standard2.1" 
}
variable "instance_source_details_source_type" { 
	default = "sourceType" 
}
variable "instance_state" { 
	default = "AVAILABLE" 
}

data "oci_core_instances" "test_instances" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	display_name = "${var.instance_display_name}"

    filter {
    	name = "id"
    	values = ["${oci_core_instance.test_instance.id}"]
    }

	filter {
    	name = "freeform_tags.Department"
    	values = ["Accounting"]
    }

	filter {
    	name = "defined_tags.${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}"
    	values = ["${var.instance_defined_tags_value}"]
    }

	filter {
    	name = "source_details.source_type"
    	values = ["image"]
    }
}

data "oci_core_instances" "test_instances_filter_out_with_nested_structure" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	display_name = "${var.instance_display_name}"

   filter {
    	name = "id"
    	values = ["${oci_core_instance.test_instance.id}"]
    }
	
	filter {
    	name = "source_details.source_type"
    	values = ["image.blah"]
    }
}

data "oci_core_instances" "test_instances_filter_out_with_map" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	display_name = "${var.instance_display_name}"

   filter {
    	name = "id"
    	values = ["${oci_core_instance.test_instance.id}"]
    }
	
	filter {
    	name = "defined_tags.${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}"
    	values = ["${var.instance_defined_tags_value}.blah"]
    }
}
                ` + compartmentIdVariableStr + CoreInstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.oci_core_instances.test_instances", "instances.#", "1"),
					resource.TestCheckResourceAttr("data.oci_core_instances.test_instances_filter_out_with_nested_structure", "instances.#", "0"),
					resource.TestCheckResourceAttr("data.oci_core_instances.test_instances_filter_out_with_map", "instances.#", "0"),
				),
			},
		},
	})
}
