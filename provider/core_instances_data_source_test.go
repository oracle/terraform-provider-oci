// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreInstanceTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *DatasourceCoreInstanceTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + s.TokenFn(`
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-vcn"
		cidr_block = "10.0.0.0/16"
	}

	resource "oci_core_subnet" "t" {
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		display_name        = "-tf-subnet"
		cidr_block          = "10.0.1.0/24"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
	}

	variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		// Oracle-provided image "Oracle-Linux-7.4-2017.12.18-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaasc56hnpnx7swoyd2fw5gyvbn3kcdmqc2guiiuvnztl2erth62xnq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaxrqeombwty6jyqgk3fraczdd63bv66xgfsqka4ktr7c57awr3p5a"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaayxmzu6n5hsntq4wlffpb4h6qh6z3uskpbm5v3v4egqlqvwicfbyq"
	  }
	}

	resource "oci_core_instance" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "{{.token}}"
		subnet_id = "${oci_core_subnet.t.id}"
		image = "${var.InstanceImageOCID[var.region]}"
		shape = "VM.Standard1.1"
		metadata {
			ssh_authorized_keys = "${var.ssh_public_key}"
		}
		timeouts {
			create = "15m"
		}
	}`, nil)

	s.ResourceName = "data.oci_core_instances.t"
}

func (s *DatasourceCoreInstanceTestSuite) TestAccDatasourceCoreInstance_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + s.TokenFn(`
				data "oci_core_instances" "t" {
					compartment_id = "${var.compartment_id}"
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					display_name = "{{.token}}"
					filter {
						name = "id"
						values = ["${oci_core_instance.t.id}"]
					}
				}`, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.display_name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.state", "RUNNING"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.shape", "VM.Standard1.1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.region"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.image"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.metadata.%"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreInstanceTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreInstanceTestSuite))
}
