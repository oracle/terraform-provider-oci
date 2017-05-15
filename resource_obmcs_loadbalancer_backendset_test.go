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

type ResourceLoadBalancerBackendsetTestSuite struct {
	suite.Suite
	Client      mockableClient
	Providers   map[string]terraform.ResourceProvider
	TimeCreated baremetal.Time
}

func (s *ResourceLoadBalancerBackendsetTestSuite) SetupTest() {
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

func (s *ResourceLoadBalancerBackendsetTestSuite) TestCreateResourceLoadBalancerBackendsetrMaximal() {
	resourceName := "baremetal_load_balancer_backendset.t"
	config := `
resource "baremetal_load_balancer_backendset" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  name             = "stub_backendset_name"
  policy           = "stub_policy"

  health_checker {
    interval_ms         = 30001
    port                = 1234
    protocol            = "stub_protocol"
    response_body_regex = "stub_regex"
  }

  ssl_configuration {
    certificate_name        = "stub_certificate_name"
    verify_depth            = 6
    verify_peer_certificate = false
  }
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
					resource.TestCheckResourceAttr(resourceName, "health_checker.port", "1234"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "stub_certificate_name"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
				),
			},
		},
	})
}

func TestResourceLoadBalancerBackendsetTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerBackendsetTestSuite))
}
