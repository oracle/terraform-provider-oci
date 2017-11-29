// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type ResourcePrivateIPTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourcePrivateIPTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProvider1() + testADs() + testVCN1() + testSubnet1() + testImage1() + testInstance1() + `
	data "oci_core_vnic_attachments" "t" {
		availability_domain = "${data.oci_identity_availability_domains.t.availability_domains.0.name}"
		compartment_id = "${var.compartment_ocid}"
		instance_id = "${oci_core_instance.t.id}"
	}`

	s.ResourceName = "oci_core_private_ip.t"
}

func (s *ResourcePrivateIPTestSuite) TestAccCoreResourcePrivateIP_basic() {
	var resId string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_private_ip" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}"
					display_name = "-private-ip"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-private-ip"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_private_ip.t", "id")
						return
					},
				),
			},
			// test update
			{
				Config: s.Config + `
				resource "oci_core_private_ip" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}"
					display_name = "-private-ip2"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-private-ip2"),
				),
			},
			// test add host name label
			{
				Config: s.Config + `
				resource "oci_core_private_ip" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}"
					display_name = "-private-ip2"
					hostname_label = "ahostname"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "ahostname"),
				),
			},
			// test destructive ip address change
			{
				Config: s.Config + `
				resource "oci_core_private_ip" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}"
					display_name = "-private-ip2"	
					hostname_label = "ahostname"
					ip_address = "10.0.1.22"
				}`,
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) (err error) {
						resId2, err := fromInstanceState(s, "oci_core_private_ip.t", "id")
						if resId == resId2 {
							return fmt.Errorf("Expected new private_ip ocid, got the same")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceCorePrivateIPTestSuite(t *testing.T) {
	suite.Run(t, new(ResourcePrivateIPTestSuite))
}
