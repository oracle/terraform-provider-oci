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

type ResourceLoadBalancerCertificateTestSuite struct {
	suite.Suite
	Client      mockableClient
	Providers   map[string]terraform.ResourceProvider
	TimeCreated baremetal.Time
}

func (s *ResourceLoadBalancerCertificateTestSuite) SetupTest() {
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

func (s *ResourceLoadBalancerCertificateTestSuite) TestCreateResourceLoadBalancerCertificateMaximal() {
	resourceName := "baremetal_load_balancer_certificate.t"
	config := `
resource "baremetal_load_balancer_certificate" "t" {
  load_balancer_id   = "ocid1.loadbalancer.stub_id"
  ca_certificate     = "stub_ca_certificate"
  certificate_name   = "stub_certificate_name"
  passphrase         = "stub_passphrase"
  private_key        = "stub_private_key"
  public_certificate = "stub_public_certificate"
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
				),
			},
		},
	})
}

func TestResourceLoadBalancerCertificateTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerCertificateTestSuite))
}
