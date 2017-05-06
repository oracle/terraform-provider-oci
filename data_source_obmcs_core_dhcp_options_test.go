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

type ResourceCoreDHCPOptionsDatasourceTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListDHCPOptions
}

func (s *ResourceCoreDHCPOptionsDatasourceTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_dhcp_options" "t" {
      compartment_id = "${var.compartment_id}"
      limit = 1
      page = "page"
      vcn_id = "vcn_id"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_dhcp_options.t"

	b1 := baremetal.DHCPOptions{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		ID:            "id1",
		Options: []baremetal.DHCPDNSOption{
			{
				Type:             "type",
				CustomDNSServers: []string{"server_type"},
				ServerType:       "server_type",
			},
		},
		State:       baremetal.ResourceAvailable,
		TimeCreated: baremetal.Time{Time: time.Now()},
	}

	b2 := b1
	b2.ID = "id2"

	s.List = &baremetal.ListDHCPOptions{
		DHCPOptions: []baremetal.DHCPOptions{b1, b2},
	}
}

func (s *ResourceCoreDHCPOptionsDatasourceTestSuite) TestReadDHCPOptions() {
	opts := &baremetal.ListOptions{}
	opts.Limit = 1
	opts.Page = "page"

	s.Client.On("ListDHCPOptions", "compartment_id", "vcn_id", opts).Return(s.List, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", "vcn_id"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListDHCPOptions", "compartment_id", "vcn_id", opts)
}

func (s *ResourceCoreDHCPOptionsDatasourceTestSuite) TestReadDHCPOptionsWithPagination() {
	opts := &baremetal.ListOptions{}
	opts.Limit = 1
	opts.Page = "page"

	listVal := *s.List
	list := &listVal
	list.NextPage = "nextpage"
	s.Client.On("ListDHCPOptions", "compartment_id", "vcn_id", opts).Return(list, nil)

	opts2 := &baremetal.ListOptions{}
	opts2.Limit = 1
	opts2.Page = "nextpage"

	list2Val := *s.List
	list2 := &list2Val
	b3 := s.List.DHCPOptions[0]
	b3.ID = "id3"
	b4 := s.List.DHCPOptions[1]
	b4.ID = "id4"
	list2.DHCPOptions = []baremetal.DHCPOptions{b3, b4}
	s.Client.On("ListDHCPOptions", "compartment_id", "vcn_id", opts2).Return(list2, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "options.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListDHCPOptions", "compartment_id", "vcn_id", opts2)
}

func TestResourceCoreDHCPOptionsDatasourceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDHCPOptionsDatasourceTestSuite))
}
