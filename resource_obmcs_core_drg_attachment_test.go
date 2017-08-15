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

type ResourceCoreDrgAttachmentTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.DrgAttachment
	DetachedRes  *baremetal.DrgAttachment
}

func (s *ResourceCoreDrgAttachmentTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
		resource "baremetal_core_virtual_network" "t" {
			cidr_block = "10.0.0.0/16"
			compartment_id = "${var.compartment_id}"
			display_name = "network_name"
		}
		resource "baremetal_core_drg" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
		}
		resource "baremetal_core_drg_attachment" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
			drg_id = "${baremetal_core_drg.t.id}"
			vcn_id = "${baremetal_core_virtual_network.t.id}"
		}
	`
	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_drg_attachment.t"

}

func (s *ResourceCoreDrgAttachmentTestSuite) TestCreateResourceCoreDrgAttachment() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAttached),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
				),
			},
		},
	})
}

func (s *ResourceCoreDrgAttachmentTestSuite) TestDetachVolume() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

}

func TestResourceCoreDrgAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDrgAttachmentTestSuite))
}
