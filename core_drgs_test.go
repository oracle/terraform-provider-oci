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

type ResourceCoreDrgsTestSuite struct {
	suite.Suite
	Client       *client.MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreDrgsTestSuite) SetupTest() {
	s.Client = &client.MockClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_drgs" "t" {
      compartment_id = "compartment_id"
      limit = 1
      page = "page"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_drgs.t"
}

func (s *ResourceCoreDrgsTestSuite) TestReadDrgs() {
	opts := []baremetal.Options{
		baremetal.Options{
			Limit: 1,
			Page:  "page",
		},
	}

	s.Client.On(
		"ListDrgs",
		"compartment_id",
		opts,
	).Return(
		&baremetal.DrgList{
			Drgs: []baremetal.Drg{
				baremetal.Drg{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id1",
					State:         baremetal.ResourceAttached,
					TimeCreated:   baremetal.Time{Time: time.Now()},
				},
				baremetal.Drg{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id2",
					State:         baremetal.ResourceAttached,
					TimeCreated:   baremetal.Time{Time: time.Now()},
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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.0.compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.1.id", "id2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListDrgs", "compartment_id", opts)
}

func TestResourceCoreDrgsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDrgsTestSuite))
}
