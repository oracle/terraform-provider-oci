package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreShapeTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreShapeTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_shape" "s" {
      compartment_id = "compartmentid"
      availability_domain = "availability_domain"
      image_id = "imageid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_shape.s"

}

func (s *ResourceCoreShapeTestSuite) TestResourceReadCoreShape() {
	opts := baremetal.ListShapesOptions{}
	opts.AvailabilityDomain = "availability_domain"
	opts.ImageID = "imageid"

	s.Client.On(
		"ListShapes",
		"compartmentid",
		opts,
	).Return(
		&baremetal.ListShapes{
			Shapes: []baremetal.Shape{
				baremetal.Shape{
					Name: "shape1",
				},
				baremetal.Shape{
					Name: "shape2",
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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "image_id", "imageid"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.0.name", "shape1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.1.name", "shape2"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListShapes", "compartmentid", opts)

}

func (s *ResourceCoreShapeTestSuite) TestResourceReadCoreShapeWithPagination() {
	opts := baremetal.ListShapesOptions{}
	opts.AvailabilityDomain = "availability_domain"
	opts.ImageID = "imageid"

	s.Client.On(
		"ListShapes",
		"compartmentid",
		opts,
	).Return(
		&baremetal.ListShapes{
			ResourceContainer: baremetal.ResourceContainer{
				NextPage: "nextpage",
			},
			Shapes: []baremetal.Shape{
				baremetal.Shape{
					Name: "shape1",
				},
				baremetal.Shape{
					Name: "shape2",
				},
			},
		},
		nil,
	)

	opts2 := &baremetal.ListShapesOptions{}
	opts2.AvailabilityDomain = "availability_domain"
	opts2.ImageID = "imageid"
	opts2.Page = "nextpage"

	s.Client.On(
		"ListShapes",
		"compartmentid",
		opts2,
	).Return(
		&baremetal.ListShapes{
			Shapes: []baremetal.Shape{
				baremetal.Shape{
					Name: "shape3",
				},
				baremetal.Shape{
					Name: "shape4",
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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "image_id", "imageid"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.0.name", "shape1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.3.name", "shape4"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListShapes", "compartmentid", opts2)

}

func TestResourceCoreShapeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreShapeTestSuite))
}
