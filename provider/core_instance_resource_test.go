// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/terraform-provider-oci/crud"
)

type ResourceCoreInstanceTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreInstanceTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
		dns_label = "examplevcn"
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
		dns_label = "examplesubnet"
	}

	variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		// Oracle-provided image "Oracle-Linux-7.4-2017.12.18-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaasc56hnpnx7swoyd2fw5gyvbn3kcdmqc2guiiuvnztl2erth62xnq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaxrqeombwty6jyqgk3fraczdd63bv66xgfsqka4ktr7c57awr3p5a"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaayxmzu6n5hsntq4wlffpb4h6qh6z3uskpbm5v3v4egqlqvwicfbyq"
	  }
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
					hostname_label = "hostname1"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard1.1"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					extended_metadata {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
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
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "image"),
					// only set if specified
					resource.TestCheckNoResourceAttr(s.ResourceName, "ipxe_script"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "hostname1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape", "VM.Standard1.1"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.user_data", "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="),
					resource.TestCheckResourceAttrSet(s.ResourceName, "metadata.ssh_authorized_keys"),
					resource.TestCheckResourceAttr(s.ResourceName, "extended_metadata.%", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "extended_metadata.keyA", "valA"),
					resource.TestCheckResourceAttr(s.ResourceName, "extended_metadata.keyB", "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "region"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.hostname_label"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.private_ip"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.InstanceLifecycleStateRunning)),
					func(ts *terraform.State) (err error) {
						instanceId, err = fromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// Switching to create_vnic_details for subnet_id and hostname_label should not lead to a change.
			// Changing the letter case in the hostname_label of the instance should also not result in a change.
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						hostname_label = "hostNAME1"
					}
					image = "${var.InstanceImageOCID[var.region]}"
					hostname_label = "HOSTName1"
					shape = "VM.Standard1.1"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					extended_metadata {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
					}
					timeouts {
						create = "15m"
					}
				}`,
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			// verify update - adds display name
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard1.1"
					display_name = "-tf-instance"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					extended_metadata {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId != instanceId {
							return fmt.Errorf("expected same instance ocid, got different")
						}
						return err
					},
				),
			},
			// Adding create_vnic_details with the same subnet_id and an updatable fields should cause an update only.
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard1.1"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					extended_metadata {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
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
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.hostname_label"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.private_ip"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "true"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId != instanceId {
							return fmt.Errorf("expected same instance ocid, got different")
						}
						return err
					},
				),
			},
			// Adding create_vnic_details flags with default values should not lead to a change.
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard1.1"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					extended_metadata {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						skip_source_dest_check = false
						assign_public_ip = true
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			// Update create_vnic_details
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard1.1"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					extended_metadata {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						display_name = "-tf-vnic-2"
						skip_source_dest_check = true
						hostname_label = "mytftesthostname"
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
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.display_name", "-tf-vnic-2"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.hostname_label", "mytftesthostname"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.private_ip"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId != instanceId {
							return fmt.Errorf("Expected same instance ocid, got different.")
						}
						return err
					},
				),
			},
			// verify force new by setting non-updateable VNIC details.
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard1.1"
					display_name = "-tf-instance"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					extended_metadata {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						display_name = "-tf-vnic-2"
						assign_public_ip = false
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
					resource.TestCheckResourceAttr(s.ResourceName, "private_ip", "10.0.1.20"),
					resource.TestCheckResourceAttr(vnicResourceName, "display_name", "-tf-vnic-2"),
					resource.TestCheckResourceAttr(vnicResourceName, "skip_source_dest_check", "true"),
					resource.TestCheckNoResourceAttr(vnicResourceName, "public_ip_address"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId == instanceId {
							return fmt.Errorf("expected new instance ocid, got the same")
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
