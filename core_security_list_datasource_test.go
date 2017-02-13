// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type CoreSecurityListDatasourceTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *CoreSecurityListDatasourceTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_security_lists" "t" {
      compartment_id = "compartment_id"
      limit = 1
      page = "page"
      vcn_id = "vcn_id"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_security_lists.t"
}

func (s *CoreSecurityListDatasourceTestSuite) TestReadSecurityLists() {
	opts := &baremetal.ListOptions{}
	opts.Limit = 1
	opts.Page = "page"

	res := &baremetal.ListSecurityLists{}
	res.NextPage = "nextpage"
	res.SecurityLists = []baremetal.SecurityList{
		{
			CompartmentID: "compartment_id",
			ID:            "id1",
			EgressSecurityRules: []baremetal.EgressSecurityRule{
				baremetal.EgressSecurityRule{
					Destination: "destination",
					ICMPOptions: &baremetal.ICMPOptions{Code: 1, Type: 2},
					Protocol:    "protocol",
				},
			},
			IngressSecurityRules: []baremetal.IngressSecurityRule{
				baremetal.IngressSecurityRule{
					TCPOptions: &baremetal.TCPOptions{
						baremetal.PortRange{Max: 2, Min: 1},
					},
					Protocol: "protocol",
					Source:   "source",
				},
			},
		},
		{
			CompartmentID: "compartment_id",
			ID:            "id2",
		},
	}

	s.Client.On(
		"ListSecurityLists",
		"compartment_id", "vcn_id",
		opts,
	).Return(res, nil)

	opts2 := &baremetal.ListOptions{}
	opts2.Limit = 1
	opts2.Page = "nextpage"

	s.Client.On(
		"ListSecurityLists",
		"compartment_id", "vcn_id",
		opts2,
	).Return(
		&baremetal.ListSecurityLists{
			SecurityLists: []baremetal.SecurityList{{ID: "id3"}, {ID: "id4"}},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.egress_security_rules.0.icmp_options.0.code", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.ingress_security_rules.0.tcp_options.0.max", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.#", "4"),
				),
			},
		},
	},
	)
}

func TestCoreSecurityListDatasourceTestSuite(t *testing.T) {
	suite.Run(t, new(CoreSecurityListDatasourceTestSuite))
}
