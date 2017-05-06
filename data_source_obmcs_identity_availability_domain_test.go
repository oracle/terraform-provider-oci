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

type ResourceIdentityAvailabilityDomainsTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListAvailabilityDomains
}

func (s *ResourceIdentityAvailabilityDomainsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_identity_availability_domains" "t" {
      compartment_id = "${var.compartment_id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_identity_availability_domains.t"

	a1 := baremetal.AvailabilityDomain{
		Name:          "AD1",
		CompartmentID: "compartmentID",
	}

	a2 := a1
	a2.Name = "AD2"

	s.List = &baremetal.ListAvailabilityDomains{
		AvailabilityDomains: []baremetal.AvailabilityDomain{a1, a2},
	}
}

func (s *ResourceIdentityAvailabilityDomainsTestSuite) TestReadAPIKeys() {
	s.Client.On("ListAvailabilityDomains", "compartmentID").Return(s.List, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "CompartmentID"),
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domains.0.name", "AD1"),
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domains.1.name", "AD2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListAPIKeys", "user_id")
}

func TestResourceIdentityAvailabilityDomainsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityAPIKeysTestSuite))
}
