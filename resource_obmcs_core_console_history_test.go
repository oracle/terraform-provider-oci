// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreConsoleHistoryTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Res          *baremetal.ConsoleHistoryMetadata
	ResourceName string
}

func (s *ResourceCoreConsoleHistoryTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	p := s.Provider.(*schema.Provider)
	res := p.ResourcesMap["baremetal_core_console_history"]
	res.Delete = func(d *schema.ResourceData, m interface{}) (e error) {
		return nil
	}

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    resource "baremetal_core_console_history" "t" {
			instance_id = "instance_id"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "baremetal_core_console_history.t"
	s.Res = &baremetal.ConsoleHistoryMetadata{
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartmentid",
		DisplayName:        "display_name",
		InstanceID:         "instance_id",
		ID:                 "id",
		State:              baremetal.ResourceSucceeded,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.Client.On("CaptureConsoleHistory", s.Res.InstanceID, (*baremetal.RetryTokenOptions)(nil)).Return(s.Res, nil)
}

func (s *ResourceCoreConsoleHistoryTestSuite) TestCreateResourceCoreInstanceConsoleHistory() {
	s.Client.On("GetConsoleHistory", "id").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
				),
			},
		},
	})
}

func TestResourceCoreConsoleHistoryTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreConsoleHistoryTestSuite))
}
