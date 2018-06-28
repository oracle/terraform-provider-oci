// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreInstanceCredentialTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreInstanceCredentialTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
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
	s.ResourceName = "data.oci_core_instance_credentials.s"
}

func (s *DatasourceCoreInstanceCredentialTestSuite) TestAccDatasourceCoreInstanceCredentials_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + `
				data "oci_core_instance_credentials" "s" {
					instance_id = "${oci_core_instance.t.id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "username"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreInstanceCredentialTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreInstanceCredentialTestSuite))
}
