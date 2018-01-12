// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreInstanceCredentialTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreInstanceCredentialTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
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
		// Oracle-provided image "Windows-Server-2012-R2-Standard-Edition-VM-Gen2-2017.10.31-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaae5cbnhdfyd75lzxypq2vk4n7w7rj6uj7rqvquoghn6n67omy5lqq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaamjrkigt6tdtfq2ovkaosmuyksbqd4y562chfd7qg4ujxpv7fe6pa"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaajfu4alt5mdx6hvdh2xpnqjuistonr4vndj37aipuec7c23ifarsq"
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
				Check: resource.ComposeTestCheckFunc(
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
