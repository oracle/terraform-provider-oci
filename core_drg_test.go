package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreDrgTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Drg
	DeletedRes   *baremetal.Drg
}

func (s *ResourceCoreDrgTestSuite) SetupTest() {
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
		resource "baremetal_core_drg" "t" {
			compartment_id = "compartment_id"
			display_name = "display_name"
		}
	`
	s.Config += testProviderConfig

	s.ResourceName = "baremetal_core_drg.t"
	s.Res = &baremetal.Drg{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		ID:            "id",
		State:         baremetal.ResourceAvailable,
		TimeCreated:   s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.DeletedRes = &baremetal.Drg{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		ID:            "id",
		State:         baremetal.ResourceTerminated,
		TimeCreated:   s.TimeCreated,
	}
	s.DeletedRes.ETag = "etag"
	s.DeletedRes.RequestID = "opcrequestid"

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On(
		"CreateDrg",
		"compartment_id",
		opts).Return(s.Res, nil)
	s.Client.On("DeleteDrg", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreDrgTestSuite) TestCreateResourceCoreDrg() {
	s.Client.On("GetDrg", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetDrg", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
				),
			},
		},
	})
}

func (s *ResourceCoreDrgTestSuite) TestCreateResourceCoreDrgWithoutDisplayName() {
	s.Client.On("GetDrg", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetDrg", "id").Return(s.DeletedRes, nil)

	s.Config = `
		resource "baremetal_core_drg" "t" {
			compartment_id = "compartment_id"
		}
	`
	s.Config += testProviderConfig

	opts := &baremetal.CreateOptions{}
	s.Client.On(
		"CreateDrg",
		"compartment_id", opts).Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
				),
			},
		},
	})
}

func (s *ResourceCoreDrgTestSuite) TestDeleteDrg() {
	s.Client.On("GetDrg", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetDrg", "id").Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "DeleteDrg", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreDrgTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDrgTestSuite))
}
