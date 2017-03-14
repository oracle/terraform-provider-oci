// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreCpeTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Cpe
}

func (s *ResourceCoreCpeTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated = baremetal.Time{Time: time.Now()}
	s.Config = `

		resource "baremetal_core_cpe" "t" {
			compartment_id = "compartmentid"
			display_name = "displayname"
      ip_address = "123.123.123.123"
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_core_cpe.t"
	s.Res = &baremetal.Cpe{
		ID:            "cpeid",
		CompartmentID: "compartmentid",
		DisplayName:   "displayname",
		IPAddress:     "123.123.123.123",
		TimeCreated:   s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "displayname"
	s.Client.On("CreateCpe", "compartmentid", "123.123.123.123", opts).Return(s.Res, nil)
	s.Client.On("DeleteCpe", "cpeid", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreCpeTestSuite) TestCreateResourceCoreCpe() {
	s.Client.On("GetCpe", "cpeid").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
					resource.TestCheckResourceAttr(s.ResourceName, "ip_address", s.Res.IPAddress),
				),
			},
		},
	})
}

func (s ResourceCoreCpeTestSuite) TestUpdateForcesNewCoreCpe() {
	s.Client.On("GetCpe", "cpeid").Return(s.Res, nil)

	updateForcingChangeConfig := `

  resource "baremetal_core_cpe" "t" {
    compartment_id = "compartmentid"
    display_name = "displayname"
    ip_address = "111.222.111.222"
  }

  `
	updateForcingChangeConfig += testProviderConfig

	result := &baremetal.Cpe{
		ID:            "cpeid2",
		CompartmentID: "compartmentid",
		DisplayName:   "displayname",
		IPAddress:     "111.222.111.222",
		TimeCreated:   s.TimeCreated,
	}
	result.ETag = "etag"
	result.RequestID = "opcrequestid"

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "displayname"
	s.Client.On("CreateCpe", "compartmentid", "111.222.111.222", opts).Return(result, nil)

	s.Client.On("GetCpe", "cpeid2").Return(result, nil)
	s.Client.On("DeleteCpe", "cpeid2", (*baremetal.IfMatchOptions)(nil)).Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config: updateForcingChangeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "ip_address", result.IPAddress),
				),
			},
		},
	})

}

func (s *ResourceCoreCpeTestSuite) TestDeleteResourceCoreCpe() {
	s.Client.On("GetCpe", "cpeid").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeleteCpe", "cpeid", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreCpeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreCpeTestSuite))
}
