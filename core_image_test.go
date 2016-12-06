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

type ResourceCoreImageTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Image
	DeletedRes   *baremetal.Image
}

func (s *ResourceCoreImageTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)
	s.Providers = map[string]terraform.ResourceProvider{"baremetal": s.Provider}

	s.ResourceName = "baremetal_core_image.t"
	s.Config = `
		resource "baremetal_core_image" "t" {
			compartment_id = "compartment_id"
			display_name = "display_name"
			instance_id = "instance_id"
		}
	`
	s.Config += testProviderConfig

	s.TimeCreated = baremetal.Time{Time: time.Now()}
	s.Res = &baremetal.Image{
		BaseImageID:            "base_image_id",
		CompartmentID:          "compartment_id",
		CreateImageAllowed:     true,
		DisplayName:            "display_name",
		ID:                     "id",
		State:                  baremetal.ResourceAvailable,
		OperatingSystem:        "operating_system",
		OperatingSystemVersion: "operating_system_version",
		TimeCreated:            s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	deletedRes := *s.Res
	s.DeletedRes = &deletedRes
	s.DeletedRes.State = baremetal.ResourceDeleted

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On("CreateImage", "compartment_id", "instance_id", opts).Return(s.Res, nil)
	s.Client.On("DeleteImage", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreImageTestSuite) TestCreateImage() {
	s.Client.On("GetImage", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetImage", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "base_image_id", s.Res.BaseImageID),
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

func (s *ResourceCoreImageTestSuite) TestCreateImageWithoutDisplayName() {
	s.Client.On("GetImage", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetImage", "id").Return(s.DeletedRes, nil)

	s.Config = `
		resource "baremetal_core_image" "t" {
			compartment_id = "compartment_id"
			instance_id = "instance_id"
		}
	`
	s.Config += testProviderConfig

	opts := &baremetal.CreateOptions{}
	s.Client.On("CreateImage", "compartment_id", "instance_id", opts).
		Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
				),
			},
		},
	})
}

func (s ResourceCoreImageTestSuite) TestUpdateImageDisplayName() {
	s.Client.On("GetImage", "id").Return(s.Res, nil).Times(3)

	config := `
		resource "baremetal_core_image" "t" {
			compartment_id = "compartment_id"
			instance_id = "instance_id"
			display_name = "new_display_name"
		}
	`
	config += testProviderConfig

	resVal := *s.Res
	res := &resVal
	res.DisplayName = "new_display_name"

	deletedResVal := *res
	deletedRes := &deletedResVal
	deletedRes.State = baremetal.ResourceDeleted

	opts := &baremetal.UpdateOptions{}
	opts.DisplayName = "new_display_name"
	s.Client.On("UpdateImage", "id", opts).Return(res, nil)
	s.Client.On("GetImage", "id").Return(res, nil).Times(2)
	s.Client.On("GetImage", "id").Return(deletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", res.DisplayName),
				),
			},
		},
	})
}

func (s *ResourceCoreImageTestSuite) TestDeleteImage() {
	s.Client.On("GetImage", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetImage", "id").Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "DeleteImage", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreImageTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreImageTestSuite))
}
