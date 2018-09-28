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
		// Oracle-provided image "Windows-Server-2012-R2-Standard-Edition-VM-2017.07.25-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaab2xgy6bijtudhsgsbgns6zwfqnkdb2bp4l4qap7e4mehv6bv3qca"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlfsi5npxguvhad3v5d5lu7dc3zcylr2csfdexgd6kor3f6zeqeq"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaanc7bsuauwkfonfmk52cn3mwjzgamhp4llsh754yahbv2e6no4u3q"
	  }
	}
	
	resource "oci_core_instance" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		subnet_id = "${oci_core_subnet.t.id}"
		image = "${var.InstanceImageOCID[var.region]}"
		shape = "VM.Standard1.1"
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
