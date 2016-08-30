package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreInstanceConsoleHistoriesTestSuite struct {
	suite.Suite
	Client       *client.MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Res          *baremetal.ConsoleHistoryMetadata
	ResourceName string
	DeletedRes   *baremetal.ConsoleHistoryMetadata
	Opts         []baremetal.Options
}

func (s *ResourceCoreInstanceConsoleHistoriesTestSuite) SetupTest() {
	s.Client = &client.MockClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    resource "baremetal_core_instance_console_histories" "t" {
      compartment_id = "compartmentid"
			instance_id = "instance_id"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "baremetal_core_instance_console_histories.t"
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

	opts := baremetal.Options{}
	s.Opts = []baremetal.Options{opts}
	s.Client.On(
		"CaptureConsoleHistory",
		s.Res.InstanceID,
	).Return(s.Res, nil)
	s.Client.On("DeleteConsoleHistory", s.Res.ID).Return(nil)
	resCopy := *s.Res
	s.DeletedRes = &resCopy
	s.DeletedRes.State = baremetal.ResourceTerminated
}

func (s *ResourceCoreInstanceConsoleHistoriesTestSuite) TestCreateResourceCoreInstanceConsoleHistory() {
	s.Client.On("GetConsoleHistory", "id", s.Opts).Return(s.Res, nil).Times(2)
	s.Client.On("GetConsoleHistory", "id", s.Opts).Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
				),
			},
		},
	})
}

func (s *ResourceCoreInstanceConsoleHistoriesTestSuite) TestTerminateInstanceConsoleHistories() {
	s.Client.On("GetConsoleHistory", "id", s.Opts).Return(s.Res, nil).Times(2)
	s.Client.On("GetConsoleHistory", "id", s.Opts).Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeleteConsoleHistory", "id")
}

func TestResourceCoreInstanceConsoleHistoriesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInstanceConsoleHistoriesTestSuite))
}
