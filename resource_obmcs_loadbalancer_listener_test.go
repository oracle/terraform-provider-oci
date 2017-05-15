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

type ResourceLoadBalancerListenerTestSuite struct {
	suite.Suite
	Client      mockableClient
	Providers   map[string]terraform.ResourceProvider
	TimeCreated baremetal.Time
}

func (s *ResourceLoadBalancerListenerTestSuite) SetupTest() {
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

func (s *ResourceLoadBalancerListenerTestSuite) TestCreateResourceLoadBalancerListenerrMaximal() {
	resourceName := "baremetal_load_balancer_listener.t"
	config := `
resource "baremetal_load_balancer_listener" "t" {
  load_balancer_id         = "stub_load_balancer_id"
  name                     = "stub_name"
  default_backend_set_name = "stub_backend_set_name"
  port                     = 1234
  protocol                 = "stub_protocol"

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
					resource.TestCheckResourceAttr(resourceName, "name", "stub_name"),
					resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", "stub_backend_set_name"),
					resource.TestCheckResourceAttr(resourceName, "port", "1234"),

					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "stub_certificate_name"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
				),
			},
		},
	})
}

func TestResourceLoadBalancerListenerTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerListenerTestSuite))
}
