// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"fmt"
	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-oci/crud"
)

type ResourceCoreInstanceTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Instance
	DeletedRes   *baremetal.Instance
}

func (s *ResourceCoreInstanceTestSuite) SetupTest() {
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
	
	data "oci_core_images" "t" {
		compartment_id = "${var.compartment_id}"
		operating_system = "Oracle Linux"
		operating_system_version = "7.3"
		limit = 1
	}`

	s.ResourceName = "oci_core_instance.t"
}

func (s *ResourceCoreInstanceTestSuite) TestAccResourceCoreInstance_basic() {

	var instanceId string
	vnicResourceName := "data.oci_core_vnic.t"

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					image = "${data.oci_core_images.t.images.0.id}"
					shape = "VM.Standard1.1"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
					}
					timeouts {
						create = "15m"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceRunning),
					func(ts *terraform.State) (err error) {
						instanceId, err = fromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// verify update
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					image = "${data.oci_core_images.t.images.0.id}"
					shape = "VM.Standard1.1"
					display_name = "-tf-instance"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId != instanceId {
							return fmt.Errorf("Expected same instance ocid, got the different.")
						}
						return err
					},
				),
			},
			// verify force new by creating an instance with create_vnic_details.
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${data.oci_core_images.t.images.0.id}"
					shape = "VM.Standard1.1"
					display_name = "-tf-instance"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						display_name = "-tf-vnic"
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					resource.TestCheckResourceAttr(vnicResourceName, "display_name", "-tf-vnic"),
					resource.TestCheckResourceAttr(vnicResourceName, "skip_source_dest_check", "false"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId == instanceId {
							return fmt.Errorf("Expected new instance ocid, got the same.")
						}
						instanceId = newId
						return err
					},
				),
			},
			// verify force new by creating an instance with different VNIC details.
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${data.oci_core_images.t.images.0.id}"
					shape = "VM.Standard1.1"
					display_name = "-tf-instance"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						display_name = "-tf-vnic"
						assign_public_ip = true
						private_ip = "10.0.1.20"
						skip_source_dest_check = true
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					resource.TestCheckResourceAttr(vnicResourceName, "display_name", "-tf-vnic"),
					resource.TestCheckResourceAttr(vnicResourceName, "skip_source_dest_check", "true"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId == instanceId {
							return fmt.Errorf("Expected new instance ocid, got the same.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestIsStatefulResource(t *testing.T) {
	var _ crud.StatefulResource = (*InstanceResourceCrud)(nil)
}

func TestResourceCoreInstanceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInstanceTestSuite))
}
