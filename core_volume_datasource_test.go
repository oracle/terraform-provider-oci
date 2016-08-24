package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumesTestSuite struct {
	suite.Suite
	Client       *client.MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVolumesTestSuite) SetupTest() {
	s.Client = &client.MockClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_volumes" "t" {
      availability_domain = "availability_domain"
      compartment_id = "compartment_id"
      limit = 1
      page = "page"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_volumes.t"
}

func (s *ResourceCoreVolumesTestSuite) TestReadVolumes() {
	opts := []baremetal.Options{
		baremetal.Options{
			AvailabilityDomain: "availability_domain",
			Limit:              1,
			Page:               "page",
		},
	}

	s.Client.On(
		"ListVolumes",
		"compartment_id",
		opts,
	).Return(
		&baremetal.VolumeList{
			Volumes: []baremetal.Volume{
				baremetal.Volume{
					AvailabilityDomain: "availability_domain",
					CompartmentID:      "compartment_id",
					DisplayName:        "display_name",
					ID:                 "id1",
					SizeInMBs:          "size_in_mbs",
					State:              baremetal.ResourceAvailable,
					TimeCreated:        baremetal.Time{Time: time.Now()},
				},
				baremetal.Volume{
					AvailabilityDomain: "availability_domain",
					CompartmentID:      "compartment_id",
					DisplayName:        "display_name",
					ID:                 "id2",
					SizeInMBs:          "size_in_mbs",
					State:              baremetal.ResourceAvailable,
					TimeCreated:        baremetal.Time{Time: time.Now()},
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.1.id", "id2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVolumes", "compartment_id", opts)
}

func TestResourceCoreVolumesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumesTestSuite))
}
