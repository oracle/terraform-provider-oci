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

	"github.com/oracle/terraform-provider-baremetal/client"
)

type DatasourceCoreCpeTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreCpeTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
resource "baremetal_core_cpe" "t" {
    compartment_id = "${var.compartment_id}"
    display_name = "name1"
    ip_address = "142.10.10.2"
}

data "baremetal_core_cpes" "s" {
    compartment_id = "${baremetal_core_cpe.t.compartment_id}"
}
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_cpes.s"

}

func (s *DatasourceCoreCpeTestSuite) TestCpeList() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.0.ip_address", "142.10.10.2"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.0.display_name", "name1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "cpes.#"),
				),
			},
		},
	},
	)

}

func (s *DatasourceCoreCpeTestSuite) TestCpePagedList() {
	res := &baremetal.ListCpes{}
	res.NextPage = "nextpage"
	res.Cpes = []baremetal.Cpe{
		{
			ID:            "id1",
			CompartmentID: "compartmentid",
			DisplayName:   "name",
			IPAddress:     "10.10.10.2",
			TimeCreated:   baremetal.Time{Time: time.Now()},
		},
		{
			ID:            "id2",
			CompartmentID: "compartmentid",
			DisplayName:   "name",
			IPAddress:     "10.10.10.3",
			TimeCreated:   baremetal.Time{Time: time.Now()},
		},
	}

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "cpes.0.ip_address", "10.10.10.2"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.1.ip_address", "10.10.10.3"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.2.id", "id3"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.#", "4"),
				),
			},
		},
	},
	)

	//
}

func TestDatasourceCoreCpeTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreCpeTestSuite))
}
