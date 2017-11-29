// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourcePrivateIPTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourcePrivateIPTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProvider1() + testADs() + testVCN1() + testSubnet1() + testImage1() + testInstance1() + `	
	data "oci_core_vnic_attachments" "t" {
		compartment_id = "${var.compartment_ocid}"
		availability_domain = "${data.oci_identity_availability_domains.t.availability_domains.0.name}"
		instance_id = "${oci_core_instance.t.id}"
	}
	
	resource "oci_core_private_ip" "t" {
		vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}"
		ip_address = "10.0.1.23"
	}`

	s.ResourceName = "data.oci_core_private_ips.t"
}

func (s *DatasourcePrivateIPTestSuite) TestAccCorePrivateIPs_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			// list by ip address
			{
				Config: s.Config + `
				data "oci_core_private_ips" "t" {
					ip_address = "10.0.1.23"
					subnet_id = "${oci_core_subnet.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.0.ip_address", "10.0.1.23"),
				),
			},
			// list by vnic id
			{
				Config: s.Config + `
				data "oci_core_private_ips" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.#", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.id"),
				),
			},
			// list by subnet id
			{
				Config: s.Config + `
				data "oci_core_private_ips" "t" {
					subnet_id = "${oci_core_subnet.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.#", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.id"),
				),
			},
		},
	},
	)
}

func TestDatasourceCorePrivateIPTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourcePrivateIPTestSuite))
}
