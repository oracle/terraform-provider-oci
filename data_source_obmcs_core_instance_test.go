// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
	"fmt"
	"time"
)

type DatasourceCoreInstanceTestSuite struct {
	suite.Suite
	Client 			*baremetal.Client
	Config			string
	Provider		terraform.ResourceProvider
	Providers 		map[string]terraform.ResourceProvider
	ResourceName	string
}


var instance_display_name string
var metadata_ssh_key_entry = "ssh_authorized_keys"

func timestampNameSuffix() string {
	var t = time.Now()
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func (s *DatasourceCoreInstanceTestSuite) SetupTest() {
	instance_display_name = "instance" + timestampNameSuffix()

	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
		data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_virtual_network" "vn" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
	}

	resource "oci_core_subnet" "sb" {
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.vn.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		route_table_id      = "${oci_core_virtual_network.vn.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.vn.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.vn.default_dhcp_options_id}"
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet"
	}

	data "oci_core_images" "img" {
		compartment_id = "${var.compartment_id}"
		operating_system = "Oracle Linux"
		operating_system_version = "7.3"
		limit = 1
	}

	resource "oci_core_instance" "inst" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "` + instance_display_name + `"
		subnet_id = "${oci_core_subnet.sb.id}"
		image = "${data.oci_core_images.img.images.0.id}"
		shape = "VM.Standard1.1"
		metadata {` +
			metadata_ssh_key_entry + `= "${var.ssh_public_key}"
		}
		timeouts {
			create = "15m"
		}
	}`

	s.ResourceName = "data.oci_core_instances.inst"
}

func (s *DatasourceCoreInstanceTestSuite) TestAccDatasourceCoreInstance_basic() {

	resource.Test(s.T(), resource.TestCase {
		PreventPostDestroyRefresh: 	true,
		Providers:					s.Providers,
		Steps:	[]resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "oci_core_instances" "inst" {
					compartment_id = "${var.compartment_id}"
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					display_name = "` + instance_display_name + `"
					limit = 1
				}`,
				Check: resource.ComposeTestCheckFunc(
					// check to make sure that the display_name matches what we set it to first.
					// Otherwise, the rest of the check could produce a false positive
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName,"instances.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName,"instances.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.region"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.state"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.shape"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.image"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.metadata.%"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instances.0.metadata." + metadata_ssh_key_entry),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreInstanceTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreInstanceTestSuite))
}