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

type ResourceCoreInternetGatewayTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.InternetGateway
	DeletedRes   *baremetal.InternetGateway
}

func (s *ResourceCoreInternetGatewayTestSuite) SetupTest() {
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
		resource "baremetal_core_internet_gateway" "t" {
			compartment_id = "compartment_id"
			display_name = "display_name"
      vcn_id = "vcnid"
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_core_internet_gateway.t"
	s.Res = &baremetal.InternetGateway{
		ETaggedResource: baremetal.ETaggedResource{
			RequestableResource: baremetal.RequestableResource{
				RequestID: "requestid",
			},
			ETag: "etag",
		},
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		ID:            "id",
		IsEnabled:     true,
		State:         baremetal.ResourceAvailable,
		ModifiedTime:  s.TimeCreated,
		TimeCreated:   s.TimeCreated,
	}

	s.DeletedRes = s.Res
	s.DeletedRes.State = baremetal.ResourceTerminated

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On(
		"CreateInternetGateway",
		s.Res.CompartmentID,
		"vcnid",
		s.Res.IsEnabled,
		opts).Return(s.Res, nil)
	s.Client.On("DeleteInternetGateway", s.Res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreInternetGatewayTestSuite) TestCreateResourceCoreInternetGateway() {
	s.Client.On("GetInternetGateway", "id").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
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

func (s ResourceCoreInternetGatewayTestSuite) TestUpdateCompartmentIDForcesNewInternetGateway() {
	s.Client.On("GetInternetGateway", "id").Return(s.Res, nil)

	config := `
  resource "baremetal_core_internet_gateway" "t" {
    compartment_id = "new_compartment_id"
    display_name = "display_name"
    vcn_id = "vcnid"
  }
	`

	config += testProviderConfig

	res := &baremetal.InternetGateway{
		ETaggedResource: baremetal.ETaggedResource{
			RequestableResource: baremetal.RequestableResource{
				RequestID: "requestid",
			},
			ETag: "etag",
		},
		CompartmentID: "new_compartment_id",
		DisplayName:   "display_name",
		ID:            "id2",
		IsEnabled:     true,
		State:         baremetal.ResourceAvailable,
		ModifiedTime:  s.TimeCreated,
		TimeCreated:   s.TimeCreated,
	}

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On(
		"CreateInternetGateway",
		res.CompartmentID,
		"vcnid",
		res.IsEnabled,
		opts).Return(res, nil)

	s.Client.On("GetInternetGateway", res.ID).Return(res, nil)
	s.Client.On("DeleteInternetGateway", res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "id", res.ID),
				),
			},
		},
	})
}

func (s *ResourceCoreInternetGatewayTestSuite) TestDeleteInternetGateway() {
	s.Client.On("GetInternetGateway", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetInternetGateway", "id").Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "DeleteInternetGateway", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreInternetGatewayTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInternetGatewayTestSuite))
}
