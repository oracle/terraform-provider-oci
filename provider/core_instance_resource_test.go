// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"regexp"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"

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
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
	  }
	}` + DefinedTagsDependencies

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
                    defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value"
									)}"
                    freeform_tags = { "Department" = "Accounting"}
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "image"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_mode", "NATIVE"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.network_type", "VFIO"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
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
					resource.TestCheckResourceAttr(s.ResourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "region"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.hostname_label"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.private_ip"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.InstanceLifecycleStateRunning)),
					resource.TestCheckResourceAttr(s.ResourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "source_details.0.source_id"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "preserve_boot_volume"),
					func(ts *terraform.State) (err error) {
						instanceId, err = fromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// Switching to create_vnic_details for subnet_id and hostname_label should not lead to a change.
			// Changing the letter case in the hostname_label of the instance should also not result in a change.
			// Changing the defined and freeform tags should
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
			// Switching to source_details for the same image ID should not lead to a change.
			// Also, check that source_type is case insensitive.
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					source_details {
						source_type = "ImAgE"
						source_id = "${var.InstanceImageOCID[var.region]}"
					}
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
			// verify update - adds display name, update tags
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
					defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value2"
									)}"
                    freeform_tags = { "CostCenter" = "42"}
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					extended_metadata {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
			// verify force new by setting non-updateable VNIC details and also add tags to the VNIC details
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
						defined_tags = "${map(
							"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value"
							)}"
                    	freeform_tags = { "Department" = "Accounting" }
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ip", "10.0.1.20"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(vnicResourceName, "display_name", "-tf-vnic-2"),
					resource.TestCheckResourceAttr(vnicResourceName, "skip_source_dest_check", "true"),
					resource.TestCheckNoResourceAttr(vnicResourceName, "public_ip_address"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId == instanceId {
							return fmt.Errorf("expected new instance ocid, got the same")
						}
						instanceId = newId
						return err
					},
				),
			},
			// verify updating vnic tags result in an update only
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
						defined_tags = "${map(
							"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue"
							)}"
                    	freeform_tags = { "Department" = "Finance" }
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId != instanceId {
							return fmt.Errorf("Expected same instance ocid, got different.")
						}
						return err
					},
				),
			},
		},
	})
}

// Tests preserve boot volume and attach behavior using source details
func (s *ResourceCoreInstanceTestSuite) TestAccResourceCoreInstance_preserveBootVolume() {

	var instanceId string
	var preservedBootVolumeId string

	// This is a reference to the TestSteps. We will use this reference to change the TF configs while test steps are
	// being run. This is necessary because some configs require a computed boot volume ID from a previous test step.
	// We cannot set the boot volume id here (it will be nil), so we have to do it within a function closure that gets
	// executed at test step execution time.
	var testStepsReference []resource.TestStep

	testSteps := []resource.TestStep{
		// verify create of an instance with source_details and that we can get a boot volume id
		{
			ImportState:       true,
			ImportStateVerify: true,
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					source_details {
						source_type = "image"
						source_id = "${var.InstanceImageOCID[var.region]}"
					}
					shape = "VM.Standard1.1"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					timeouts {
						create = "15m"
					}
				}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "image"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_mode", "NATIVE"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_options.#", "1"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.boot_volume_type", "ISCSI"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.firmware", "UEFI_64"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.network_type", "VFIO"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
				// only set if specified
				resource.TestCheckNoResourceAttr(s.ResourceName, "ipxe_script"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
				resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "hostname1"),
				resource.TestCheckResourceAttr(s.ResourceName, "shape", "VM.Standard1.1"),
				resource.TestCheckResourceAttr(s.ResourceName, "metadata.%", "2"),
				resource.TestCheckResourceAttr(s.ResourceName, "metadata.user_data", "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="),
				resource.TestCheckResourceAttrSet(s.ResourceName, "metadata.ssh_authorized_keys"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "region"),
				resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.display_name"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.hostname_label"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.private_ip"),
				resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "true"),
				resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.InstanceLifecycleStateRunning)),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.0.source_type", "image"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "source_details.0.source_id"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "source_details.0.boot_volume_size_in_gbs"),
				resource.TestCheckNoResourceAttr(s.ResourceName, "preserve_boot_volume"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "boot_volume_id"),
				// Store the instance ID for future verification
				func(ts *terraform.State) (err error) {
					instanceId, err = fromInstanceState(ts, s.ResourceName, "id")
					return err
				},
			),
		},
		// Switching from source_details back to image ID should not lead to a change.
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
					timeouts {
						create = "15m"
					}
				}`,
			ExpectNonEmptyPlan: false,
			PlanOnly:           true,
		},
		// verify the preserve_boot_volume setting can be applied and doesn't cause a ForceNew instance
		{
			ImportState:       true,
			ImportStateVerify: true,
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					source_details {
						source_type = "image"
						source_id = "${var.InstanceImageOCID[var.region]}"
					}
					preserve_boot_volume = "true"
					shape = "VM.Standard1.1"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					timeouts {
						create = "15m"
					}
				}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(s.ResourceName, "preserve_boot_volume", "true"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "boot_volume_id"),
				// Verify that we didn't get a new Instance
				func(ts *terraform.State) (err error) {
					newId, err := fromInstanceState(ts, s.ResourceName, "id")
					if newId != instanceId {
						return fmt.Errorf("expected same instance ocid, got different")
					}
					return err
				},
				// Store the boot volume id, so we can use it for attaching to an Instance later
				// Also update all the config test steps to use the computed boot volume ID
				func(ts *terraform.State) (err error) {
					preservedBootVolumeId, err = fromInstanceState(ts, s.ResourceName, "boot_volume_id")

					_, tokenFn := tokenize()
					for idx := range testStepsReference {
						testStepsReference[idx].Config = tokenFn(testStepsReference[idx].Config, map[string]string{"preservedBootVolumeId": preservedBootVolumeId})
					}

					return err
				},
			),
		},
		// ForceNew an instance by changing its hostname_label and boot volume size
		// Verify that the boot volume was preserved and can be attached to the new instance as a data volume.
		// Also verify that the new boot volume size is being used.
		{
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname2"
					source_details {
						source_type = "image"
						source_id = "${var.InstanceImageOCID[var.region]}"
						boot_volume_size_in_gbs = "60"
					}
					preserve_boot_volume = "false"
					shape = "VM.Standard1.1"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					timeouts {
						create = "15m"
					}
				}

				resource "oci_core_volume_attachment" "volume_attach" {
					attachment_type = "iscsi"
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "{{.preservedBootVolumeId}}"
				}
				`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(s.ResourceName, "preserve_boot_volume", "false"),
				TestCheckResourceAttributesEqual("oci_core_volume_attachment.volume_attach", "instance_id", s.ResourceName, "id"),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.0.source_type", "image"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "source_details.0.source_id"),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
				// Verify that we got a new Instance
				func(ts *terraform.State) (err error) {
					newId, err := fromInstanceState(ts, s.ResourceName, "id")
					if newId == instanceId {
						return fmt.Errorf("expected different instance ocid, got same")
					}

					instanceId = newId
					return err
				},
				// Verify that the volume attachment's ID is the same as the preserved boot volume
				func(ts *terraform.State) (err error) {
					volumeId, err := fromInstanceState(ts, "oci_core_volume_attachment.volume_attach", "volume_id")
					if preservedBootVolumeId != volumeId {
						return fmt.Errorf("expected attached volume id to be same as preserved boot volume, got different one")
					}

					return err
				},
			),
		},
		// Detach the boot volume and force a new instance by attaching preserved boot volume in the source details.
		{
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname2"
					source_details {
						source_type = "bootVolume"
						source_id = "{{.preservedBootVolumeId}}"
					}
					preserve_boot_volume = "false"
					shape = "VM.Standard1.1"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					timeouts {
						create = "15m"
					}
				}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(s.ResourceName, "preserve_boot_volume", "false"),
				// Verify that we got a new Instance
				func(ts *terraform.State) (err error) {
					newId, err := fromInstanceState(ts, s.ResourceName, "id")
					if newId == instanceId {
						return fmt.Errorf("expected different instance ocid, got same")
					}

					instanceId = newId
					return err
				},
				// Verify that the boot volume attachment is the same as the preserved boot volume
				func(ts *terraform.State) (err error) {
					volumeId, err := fromInstanceState(ts, s.ResourceName, "boot_volume_id")
					if preservedBootVolumeId != volumeId {
						return fmt.Errorf("expected attached boot volume ID to be same as preserved boot volume, got different one")
					}

					return err
				},
			),
		},
		// ForceNew an instance and attach to the old boot volume, which should have been deleted because we didn't set
		// the preserve flag in the previous step. This should result in an error from service.
		{
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname2"
					source_details {
						source_type = "bootVolume"
						source_id = "{{.preservedBootVolumeId}}"
					}
					preserve_boot_volume = "false"
					shape = "VM.Standard1.1"
					metadata {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					timeouts {
						create = "15m"
					}
				}`,
			ExpectError: regexp.MustCompile("One or more of the specified volumes are not found"),
		},
	}

	testStepsReference = testSteps
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps:     testSteps,
	})
}

func TestIsStatefulResource(t *testing.T) {
	var _ crud.StatefulResource = (*InstanceResourceCrud)(nil)
}

func TestResourceCoreInstanceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInstanceTestSuite))
}
