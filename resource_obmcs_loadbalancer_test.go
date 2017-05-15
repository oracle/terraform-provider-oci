// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourceLoadBalancerTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceLoadBalancerTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.ResourceName = "baremetal_load_balancer.t"
	s.Config = loadbalancerConfig
	s.Config += testProviderConfig()
}

func (s *ResourceLoadBalancerTestSuite) TestCreateResourceLoadBalancerMaximal() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					// Assigned
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "lb_display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					// Computed
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_addresses.#"),
				),
			},
		},
	})
}

func TestResourceLoadBalancerTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerTestSuite))
}
