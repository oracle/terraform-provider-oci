// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
	"regexp"
)

type ResourceCoreVnicAttachmentTestSuite struct {
	suite.Suite
	Client           *baremetal.Client
	Provider         terraform.ResourceProvider
	Providers        map[string]terraform.ResourceProvider
	TimeCreated      baremetal.Time
	Config           string
	ResourceName     string
	VnicResourceName string
	Res              *baremetal.VnicAttachment
}

func (s *ResourceCoreVnicAttachmentTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + instanceDnsConfig
	s.ResourceName = "oci_core_vnic_attachment.va"
	s.VnicResourceName = "data.oci_core_vnic.v"
}

func (s *ResourceCoreVnicAttachmentTestSuite) TestAccResourceCoreVnicAttachment_basic() {

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
							display_name = "-tf-vnic"
							assign_public_ip = false
						}
					}
					data "oci_core_vnic" "v" {
						vnic_id = "${oci_core_vnic_attachment.va.vnic_id}"
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-va1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAttached),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vlan_tag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "id"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "display_name", "-tf-vnic"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "private_ip_address"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "public_ip_address", ""),
				),
			},
			{
				// Create a new VNIC and VNIC Attachment with private IP, hostname, and assign a public IP.
				ImportState:       true,
				ImportStateVerify: true,
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
							}
						}
						data "oci_core_vnic" "v" {
						  vnic_id = "${oci_core_vnic_attachment.va.vnic_id}"
						}
					`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAttached),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "id"),
					resource.TestCheckResourceAttr(s.VnicResourceName, "private_ip_address", "10.0.1.20"),
					resource.TestCheckResourceAttrSet(s.VnicResourceName, "public_ip_address"),
					resource.TestMatchResourceAttr(s.VnicResourceName, "public_ip_address", regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+\.[0-9]`)),
					resource.TestCheckResourceAttr(s.VnicResourceName, "hostname_label", "myvnichostname"),
				),
			},
		},
	})
}

func TestResourceCoreVnicAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVnicAttachmentTestSuite))
}
