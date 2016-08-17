package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeTestSuite struct {
	suite.Suite
	Client       *MockClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Volume
}

func (s *ResourceCoreVolumeTestSuite) SetupTest() {
	s.Client = &MockClient{}

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
		resource "baremetal_core_volume" "t" {
			availability_domain = "availability_domain"
			compartment_id = "compartment_id"
			display_name = "didsplay_name"
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_core_volume.t"
	s.Res = &baremetal.Volume{
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "display_name",
		ID:                 "id",
		SizeInMBs:          "size_in_mbs",
		State:              "state",
		TimeCreated:        s.TimeCreated,
		ETag:               "etag",
		OPCRequestID:       "opc_request_id",
	}

	opts := baremetal.Options{DisplayName: "display_name"}
	s.Client.On(
		"CreateVolume",
		"availability_domain",
		"compartment_id", []baremetal.Options{opts}).Return(s.Res, nil)
}

func (s *ResourceCoreVolumeTestSuite) TestCreateResourceCoreVolume() {
	s.Client.On("GetVolume", "id", []baremetal.Options(nil)).Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", s.Res.AvailabilityDomain),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", s.Res.SizeInMBs),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
				),
			},
		},
	})
}
