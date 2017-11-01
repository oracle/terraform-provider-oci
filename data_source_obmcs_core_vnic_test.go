// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVnicTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVnicTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + instanceDnsConfig
	s.ResourceName = "data.oci_core_vnic.t"
}

func (s *DatasourceCoreVnicTestSuite) TestAccDatasourceCoreAttachVnic_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
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
					resource.TestCheckResourceAttr(s.ResourceName, "skip_source_dest_check", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
				),
			},
		},
	})
}

func TestDatasourceCoreVnicTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVnicTestSuite))
}
