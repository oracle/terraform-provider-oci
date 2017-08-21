// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
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
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.Config = instanceDnsConfig
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_vnic.v"
}

func (s *DataSourceCoreVnicTestSuite) TestAttachVnic() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
						data "baremetal_core_vnic_attachments" "va" {
						  compartment_id = "${var.compartment_id}"
						  instance_id = "${baremetal_core_instance.t.id}"
						}
						data "baremetal_core_vnic" "v" {
						  vnic_id = "${lookup(data.baremetal_core_vnic_attachments.va.vnic_attachments[0],"vnic_id")}"
						}
					`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance-vnic"),
					resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "testinstance"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_primary", "true"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "mac_address"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func TestDataSourceCoreVnicTestSuite(t *testing.T) {
	suite.Run(t, new(DataSourceCoreVnicTestSuite))
}
