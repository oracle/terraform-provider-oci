package main

import (
	"testing"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreShapeTestSuite struct {
	suite.Suite
	Client       *MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreShapeTestSuite) SetupTest() {
	s.Client = &MockClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_shape" "s" {
      compartment_id = "compartmentid"
      availability_domain = "availabilityid"
      image_id = "imageid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_shape.s"

}

func (s *ResourceCoreShapeTestSuite) TestResourceReadCoreShape() {
	opts := []baremtlsdk.CoreOptions{
		baremtlsdk.CoreOptions{
			AvailabilityDomain: "availabilityid",
			ImageID:            "imageid",
		},
	}

	s.Client.On(
		"ListShapes",
		"compartmentid",
		opts,
	).Return(
		&baremtlsdk.ShapeList{
			Shapes: []string{
				"shape1",
				"shape2",
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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "image_id", "imageid"),
					testCheckAttributeTypeList(s.ResourceName, "shapes", []string{"shape1", "shape2"}),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListShapes", "compartmentid", opts)

}

func TestResourceCoreShapeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreShapeTestSuite))
}
