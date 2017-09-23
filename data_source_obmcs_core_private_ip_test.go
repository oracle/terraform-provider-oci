// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourcePrivateIPTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListPrivateIPs
}

func (s *DatasourcePrivateIPTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = vnicConfig + `
resource "oci_core_private_ip" "testPrivateIP" {
	vnic_id = "${lookup(data.oci_core_vnic_attachments.vnics.vnic_attachments[0],"vnic_id")}"
	ip_address = "10.0.1.23"
}
	`
	s.Config += testProviderConfig()
	s.ResourceName = "data.oci_core_private_ips.testPrivateIPs"
}

func (s *DatasourcePrivateIPTestSuite) TestListPrivateIPsByVnicID() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
data "oci_core_private_ips" "testPrivateIPs" {
	vnic_id = "${lookup(data.oci_core_vnic_attachments.vnics.vnic_attachments[0],"vnic_id")}"
}
				`,
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

func (s *DatasourcePrivateIPTestSuite) TestListPrivateIPsBySubnetID() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
data "oci_core_private_ips" "testPrivateIPs" {
	subnet_id = "${oci_core_subnet.WebSubnetAD1.id}"
}
				`,
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

func (s *DatasourcePrivateIPTestSuite) TestListPrivateIPsByIPAddress() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
data "oci_core_private_ips" "testPrivateIPs" {
	ip_address = "10.0.1.23"
	subnet_id = "${oci_core_subnet.WebSubnetAD1.id}"
}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.0.ip_address", "10.0.1.23"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.#", "1"),
				),
			},
		},
	},
	)
}

func TestDatasourcePrivateIPTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourcePrivateIPTestSuite))
}
