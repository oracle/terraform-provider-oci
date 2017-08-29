// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourcePrivateIPTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.PrivateIP
	DeletedRes   *baremetal.PrivateIP
}

func (s *ResourcePrivateIPTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = vnicConfig + `
resource "oci_core_private_ip" "testPrivateIP" {
	vnic_id = "${lookup(data.oci_core_vnic_attachments.vnics.vnic_attachments[0],"vnic_id")}"
	display_name = "display_name"
}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "oci_core_private_ip.testPrivateIP"

}

func (s *ResourcePrivateIPTestSuite) TestCreateResourcePrivateIP() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
				),
			},
		},
	})
}

func (s ResourcePrivateIPTestSuite) TestUpdatePrivateIPDisplayName() {
	config := vnicConfig + `
resource "oci_core_private_ip" "testPrivateIP" {
	vnic_id = "${lookup(data.oci_core_vnic_attachments.vnics.vnic_attachments[0],"vnic_id")}"
	display_name = "newDisplayName"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "newDisplayName"),
				),
			},
		},
	})
}
func (s ResourcePrivateIPTestSuite) TestUpdatePrivateIPHostnameLabel() {
	config := vnicConfig + `
resource "oci_core_private_ip" "testPrivateIP" {
	vnic_id = "${lookup(data.oci_core_vnic_attachments.vnics.vnic_attachments[0],"vnic_id")}"
	hostname_label = "newhostnamelabel"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "newhostnamelabel"),
				),
			},
		},
	})
}

func (s ResourcePrivateIPTestSuite) TestUpdateIPAddressForcesNewPrivateIP() {

	config := vnicConfig + `
resource "oci_core_private_ip" "testPrivateIP" {
	vnic_id = "${lookup(data.oci_core_vnic_attachments.vnics.vnic_attachments[0],"vnic_id")}"
	ip_address = "10.0.1.22"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "ip_address", "10.0.1.22"),
				),
			},
		},
	})
}

func TestResourcePrivateIPTestSuite(t *testing.T) {
	suite.Run(t, new(ResourcePrivateIPTestSuite))
}
