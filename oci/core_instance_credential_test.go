// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	instanceCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": Representation{repType: Required, create: `${oci_core_instance.t.id}`},
	}

	InstanceCredentialResourceConfig = InstanceCredentialResourceDependencies + `

`
	InstanceCredentialPropertyVariables = `

`
	InstanceCredentialResourceDependencies = ""

	WindowsInstanceDnsConfig = `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
	}
	
	resource "oci_core_subnet" "t" {
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet"
	}
	
	variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		# The below Image OCIDs are for Windows-Server-2012-R2-Standard-Edition-VM-Gen2-2018.12.12-0
    	# See https://docs.cloud.oracle.com/iaas/images/image/5e34cde5-6cef-4cc3-b8f1-c8fc3a088302/
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaarlo3ace3wq34aompwj3u2z2xteonboapg663woz6d2iovarowhja"
    	us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaabzwak2haqxh3r7h6dajgu4enp7q7hcrreql45awryd5frjsd5l6a"
    	eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaourcjktoe3gprvwfksxc36r4rxgbcjs5qvtrja6w6euivci635vq"
    	uk-london-1  = "ocid1.image.oc1.uk-london-1.aaaaaaaadb4mg7ii73wkrntmiunr7x7qrh7ompczvy3xbggm27pkhotpgj2q"
	  }
	}
	
	resource "oci_core_instance" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		subnet_id = "${oci_core_subnet.t.id}"
		image = "${var.InstanceImageOCID[var.region]}"
		shape = "VM.Standard2.1"
		metadata {
			ssh_authorized_keys = "${var.ssh_public_key}"
		}
		timeouts {
			create = "15m"
		}
	}`
)

func TestCoreInstanceCredentialResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_instance_credentials.test_instance_credentials"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + WindowsInstanceDnsConfig +
					generateDataSourceFromRepresentationMap("oci_core_instance_credentials", "test_instance_credentials", Required, Create, instanceCredentialSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceCredentialResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "password"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "username"),
				),
			},
		},
	})
}
