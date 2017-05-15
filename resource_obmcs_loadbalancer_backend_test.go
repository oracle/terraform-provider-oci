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

type ResourceLoadBalancerBackendTestSuite struct {
	suite.Suite
	Client      mockableClient
	Providers   map[string]terraform.ResourceProvider
	TimeCreated baremetal.Time
}

func (s *ResourceLoadBalancerBackendTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": Provider(
			func(d *schema.ResourceData) (interface{}, error) {
				return s.Client, nil
			},
		),
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}
}

func (s *ResourceLoadBalancerBackendTestSuite) TestCreateResourceLoadBalancerBackendMaximal() {
	resourceName := "baremetal_load_balancer_backend.t"
	config := `
resource "baremetal_load_balancer_backend" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  backendset_name  = "stub_backendset_name"
  name             = "stub_backend_name"
  ip_address       = "1.2.3.4"
  port             = 1234
  backup           = true
  drain            = true
  offline          = true
  weight           = 1
}
`
	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "backendset_name", "stub_backendset_name"),
					resource.TestCheckResourceAttr(resourceName, "name", "stub_backend_name"),

					resource.TestCheckResourceAttr(resourceName, "ip_address", "1.2.3.4"),

					resource.TestCheckResourceAttr(resourceName, "backup", "true"),
					resource.TestCheckResourceAttr(resourceName, "drain", "true"),
					resource.TestCheckResourceAttr(resourceName, "offline", "true"),
					resource.TestCheckResourceAttr(resourceName, "weight", "1"),
				),
			},
		},
	})
}

func TestResourceLoadBalancerBackendTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerBackendTestSuite))
}
