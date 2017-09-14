// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DataSourceCoreVnicTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DataSourceCoreVnicTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + instanceDnsConfig
	s.ResourceName = "data.oci_core_vnic.v"
}

func (s *DataSourceCoreVnicTestSuite) TestAccDatasrouceCoreAttachVnic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_core_vnic_attachments" "va" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "v" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.va.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "mac_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance-vnic"),
					resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "testinstance"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_primary", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
				),
			},
		},
	})
}

func TestDataSourceCoreVnicTestSuite(t *testing.T) {
	suite.Run(t, new(DataSourceCoreVnicTestSuite))
}
