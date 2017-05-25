// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
	"testing"
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
	s.Config = loadbalancerConfig + certificateConfig + `

resource "baremetal_load_balancer_backendset" "no_cert" {
  load_balancer_id = "${baremetal_load_balancer.t.id}"
  name             = "stub_backendset_name_no_cert"
  policy           = "ROUND_ROBIN"

  health_checker {
    interval_ms         = 30000
    port                = 1234
    protocol            = "HTTP"
    response_body_regex = ".*"
    url_path = "/"
  }
}

resource "baremetal_load_balancer_backendset" "t" {
  load_balancer_id = "${baremetal_load_balancer.t.id}"
  name             = "stub_backendset_name"
  policy           = "ROUND_ROBIN"

  health_checker {
    interval_ms         = 30000
    port                = 1234
    protocol            = "HTTP"
    response_body_regex = ".*"
    url_path = "/"
  }

  ssl_configuration {
    certificate_name        = "${baremetal_load_balancer_certificate.t.certificate_name}"
    verify_depth            = 6
    verify_peer_certificate = false
  }
}

resource "baremetal_load_balancer_listener" "t" {
  load_balancer_id         = "${baremetal_load_balancer.t.id}"
  name                     = "stub_listener_name"
  default_backend_set_name = "${baremetal_load_balancer_backendset.t.name}"
  port                     = 443
  protocol                 = "HTTP"

  ssl_configuration {
      certificate_name        = "${baremetal_load_balancer_certificate.t.certificate_name}"
      verify_depth            = 6
      verify_peer_certificate = false
  }
}

resource "baremetal_load_balancer_backend" "t" {
  load_balancer_id = "${baremetal_load_balancer.t.id}"
  backendset_name  = "${baremetal_load_balancer_backendset.t.name}"
  ip_address       = "1.2.3.4"
  port             = 1234
  backup           = true
  drain            = true
  offline          = true
  weight           = 1
}
`
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
					resource.TestCheckResourceAttr("baremetal_load_balancer.t", "display_name", "lb_display_name"),
					resource.TestCheckResourceAttrSet("baremetal_load_balancer.t", "id"),
					// Computed
					resource.TestCheckResourceAttrSet("baremetal_load_balancer.t", "ip_addresses.#"),

					resource.TestCheckResourceAttr("baremetal_load_balancer_listener.t", "ssl_configuration.#", "1"),

					// Certificate
					resource.TestCheckResourceAttr("baremetal_load_balancer_certificate.t", "certificate_name", "stub_certificate_name"),

					// BackendSet
					resource.TestCheckResourceAttr("baremetal_load_balancer_backendset.t", "name", "stub_backendset_name"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_backendset.t", "health_checker.0.port", "1234"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_backendset.t", "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_backendset.t", "ssl_configuration.0.certificate_name", "stub_certificate_name"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_backendset.t", "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_backendset.t", "ssl_configuration.0.verify_peer_certificate", "false"),

					resource.TestCheckResourceAttr("baremetal_load_balancer_backendset.no_cert", "name", "stub_backendset_name_no_cert"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_backendset.no_cert", "health_checker.0.port", "1234"),

					// Listener
					resource.TestCheckResourceAttr("baremetal_load_balancer_listener.t", "name", "stub_listener_name"),

					resource.TestCheckResourceAttr("baremetal_load_balancer_listener.t", "ssl_configuration.0.certificate_name", "stub_certificate_name"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_listener.t", "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_listener.t", "ssl_configuration.0.verify_peer_certificate", "false"),

					// Backend
					resource.TestCheckResourceAttr("baremetal_load_balancer_backend.t", "ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_backend.t", "backup", "true"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_backend.t", "drain", "true"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_backend.t", "offline", "true"),
					resource.TestCheckResourceAttr("baremetal_load_balancer_backend.t", "weight", "1"),
				),
			},
		},
	})
}

func TestResourceLoadBalancerTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerTestSuite))
}
