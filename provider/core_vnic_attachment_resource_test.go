// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVnicAttachmentTestSuite struct {
	suite.Suite
	Providers        map[string]terraform.ResourceProvider
	Config           string
	ResourceName     string
	VnicResourceName string
}

func (s *ResourceCoreVnicAttachmentTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + instanceDnsConfig
	s.ResourceName = "oci_core_vnic_attachment.va"
	s.VnicResourceName = "data.oci_core_vnic.v"
}

func (s *ResourceCoreVnicAttachmentTestSuite) TestAccResourceCoreVnicAttachment_basic() {

	var vaId string

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_vnic_attachment" "va" {
						instance_id = "${oci_core_instance.t.id}"
						display_name = "-tf-va1"
						create_vnic_details {
							subnet_id = "${oci_core_subnet.t.id}"
							assign_public_ip = false
							defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
							freeform_tags = { "Department" = "Accounting" }
						}
					}
					data "oci_core_vnic" "v" {
						vnic_id = "${oci_core_vnic_attachment.va.vnic_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-va1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VnicAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vlan_tag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "private_ip_address"),
					// @SDK 1/2018: Since we don't assign a public IP to this vnic, we will get a response from server
					// without a public_ip_address. Old SDK would have set it to empty, but new SDK will set it to nil.
					// Commenting out until we have a better way of handling this.
					//resource.TestCheckResourceAttr(s.VnicResourceName, "public_ip_address", ""),
					resource.TestCheckNoResourceAttr(s.VnicResourceName, "public_ip_address"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "skip_source_dest_check", "false"),
					func(ts *terraform.State) (err error) {
						vaId, err = fromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			{
				// Update the VNIC
				Config: s.Config + `
					resource "oci_core_vnic_attachment" "va" {
						instance_id = "${oci_core_instance.t.id}"
						display_name = "-tf-va1"
						create_vnic_details {
							subnet_id = "${oci_core_subnet.t.id}"
							display_name = "-tf-vnic-2"
							assign_public_ip = false
							hostname_label = "myvnichostname"
							skip_source_dest_check = true
							defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}"
							freeform_tags = { "Department" = "Finance" }
						}
					}
					data "oci_core_vnic" "v" {
						vnic_id = "${oci_core_vnic_attachment.va.vnic_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-va1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VnicAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vlan_tag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.display_name", "-tf-vnic-2"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "private_ip_address"),
					// @SDK 1/2018: Since we don't assign a public IP to this vnic, we will get a response from server
					// without a public_ip_address. Old SDK would have set it to empty, but new SDK will set it to nil.
					// Commenting out until we have a better way of handling this.
					//resource.TestCheckResourceAttr(s.VnicResourceName, "public_ip_address", ""),
					resource.TestCheckNoResourceAttr(s.VnicResourceName, "public_ip_address"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "true"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId != vaId {
							return fmt.Errorf("Expected same ocid, got different.")
						}
						return err
					},
				),
			},
			{
				// Create a new VNIC and VNIC Attachment with different options.
				Config: s.Config + `
						resource "oci_core_vnic_attachment" "va" {
							instance_id = "${oci_core_instance.t.id}"
							display_name = "-tf-va1"
							create_vnic_details {
								subnet_id = "${oci_core_subnet.t.id}"
								display_name = "-tf-vnic"
								assign_public_ip = true
								private_ip = "10.0.1.20"
								hostname_label = "myvnichostname"
								skip_source_dest_check = true
							}
						}
						data "oci_core_vnic" "v" {
						  vnic_id = "${oci_core_vnic_attachment.va.vnic_id}"
						}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VnicAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "id"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "private_ip_address", "10.0.1.20"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "public_ip_address"),
					resource.TestMatchResourceAttr(s.VnicResourceName, "public_ip_address", regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+\.[0-9]`)),
					resource.TestCheckResourceAttr(s.VnicResourceName, "hostname_label", "myvnichostname"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "skip_source_dest_check", "true"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId == vaId {
							return fmt.Errorf("Expected new ocid, got the same.")
						}
						vaId = newId
						return err
					},
				),
			},
			{
				// Switching skip_source_dest_check and assign_public_ip from true to "true" will destroy and recreate, but should result in a
				// VNIC with the same value.
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
						resource "oci_core_vnic_attachment" "va" {
							instance_id = "${oci_core_instance.t.id}"
							display_name = "-tf-va1"
							create_vnic_details {
								subnet_id = "${oci_core_subnet.t.id}"
								display_name = "-tf-vnic"
								assign_public_ip = "true"
								private_ip = "10.0.1.20"
								hostname_label = "myvnichostname"
								skip_source_dest_check = "true"
							}
						}
						data "oci_core_vnic" "v" {
						  vnic_id = "${oci_core_vnic_attachment.va.vnic_id}"
						}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VnicAttachmentLifecycleStateAttached)),
					resource.TestCheckResourceAttr(s.VnicResourceName, "private_ip_address", "10.0.1.20"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "skip_source_dest_check", "true"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId != vaId {
							return fmt.Errorf("Expected same ocid, got different.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceCoreVnicAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVnicAttachmentTestSuite))
}
