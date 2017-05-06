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

type ResourceCoreInstanceCredentialTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreInstanceCredentialTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_instance_credentials" "s" {
      instance_id = "instanceid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_instance_credentials.s"

}

func (s *ResourceCoreInstanceCredentialTestSuite) TestResourceReadCoreInstanceCredential() {
	s.Client.On(
		"GetWindowsInstanceInitialCredentials",
		"instanceid",
	).Return(
		&baremetal.InstanceCredentials{"username", "password"},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "instance_id", "instanceid"),
					resource.TestCheckResourceAttr(s.ResourceName, "username", "username"),
					resource.TestCheckResourceAttr(s.ResourceName, "password", "password"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "GetWindowsInstanceInitialCredentials", "instanceid")
}

func TestResourceCoreInstanceCredentialTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInstanceCredentialTestSuite))
}
