// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVnicTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVnicTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_vnic" "t" {
      vnic_id = "vnicid"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_vnic.t"
}

func (s *DatasourceCoreVnicTestSuite) TestReadVnic() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availabilitydomain"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceActive),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ip_address", "10.10.10.10"),
					resource.TestCheckResourceAttr(s.ResourceName, "public_ip_address", "52.53.54.55"),
				),
			},
		},
	},
	)

}

func TestDatasourceCoreVnicTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVnicTestSuite))
}
