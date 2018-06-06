// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourcePrivateIPTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourcePrivateIPTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + testADs() + testVCN1() + testSubnet1() + testImage1() + testInstance1() + `
	data "oci_core_vnic_attachments" "t" {
		compartment_id = "${var.compartment_id}"
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.0.ip_address", "10.0.1.23"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.0.hostname_label", ""),
				),
			},
			// list by vnic id
			{
				Config: s.Config + `
				data "oci_core_private_ips" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.#", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.time_created"),
					TestCheckResourceAttributesEqual(s.ResourceName, "private_ips.0.hostname_label", "oci_core_instance.t", "create_vnic_details.0.hostname_label"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.1.hostname_label", ""),
				),
			},
			// list by subnet id
			{
				Config: s.Config + `
				data "oci_core_private_ips" "t" {
					subnet_id = "${oci_core_subnet.t.id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.#", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.time_created"),
					TestCheckResourceAttributesEqual(s.ResourceName, "private_ips.0.hostname_label", "oci_core_instance.t", "create_vnic_details.0.hostname_label"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.1.hostname_label", ""),
				),
			},
		},
	},
	)
}

func TestDatasourceCorePrivateIPTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourcePrivateIPTestSuite))
}
