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

type ResourceCoreInstancesTestSuite struct {
	suite.Suite
	Client       *client.MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreInstancesTestSuite) SetupTest() {
	s.Client = &client.MockClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_instances" "s" {
      compartment_id = "compartmentid"
      availability_domain = "availabilityid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_instances.s"

}

func (s *ResourceCoreInstancesTestSuite) TestResourceListInstances() {
	opts := []baremetal.Options{
		baremetal.Options{
			AvailabilityDomain: "availabilityid",
		},
	}

	metadata := map[string]string{
		"foo": "bar",
		"baz": "buz",
	}

	s.Client.On(
		"ListInstances",
		"compartmentid",
		opts,
	).Return(
		&baremetal.ListInstances{
			Instances: []baremetal.Instance{
				baremetal.Instance{
					ID:                 "id1",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "inst1",
					State:              baremetal.ResourceRunning,
					Metadata:           metadata,
					Region:             "here",
					Shape:              "round",
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
				},
				baremetal.Instance{
					ID:                 "id2",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "inst2",
					State:              baremetal.ResourceRunning,
					Metadata:           metadata,
					Region:             "here",
					Shape:              "round",
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
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
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListInstances", "compartmentid", opts)

}

func (s *ResourceCoreInstancesTestSuite) TestResourceListInstancesPaged() {
	opts := []baremetal.Options{
		baremetal.Options{
			AvailabilityDomain: "availabilityid",
		},
	}

	metadata := map[string]string{
		"foo": "bar",
		"baz": "buz",
	}

	s.Client.On(
		"ListInstances",
		"compartmentid",
		opts,
	).Return(
		&baremetal.ListInstances{
			ResourceContainer: baremetal.ResourceContainer{
				NextPage: "nextpage",
			},
			Instances: []baremetal.Instance{
				baremetal.Instance{
					ID:                 "id1",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "inst1",
					State:              baremetal.ResourceRunning,
					Metadata:           metadata,
					Region:             "here",
					Shape:              "round",
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
				},
				baremetal.Instance{
					ID:                 "id2",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "inst2",
					State:              baremetal.ResourceRunning,
					Metadata:           metadata,
					Region:             "here",
					Shape:              "round",
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
				},
			},
		},
		nil,
	)

	opts2 := []baremetal.Options{
		baremetal.Options{
			AvailabilityDomain: "availabilityid",
			Page:               "nextpage",
		},
	}

	s.Client.On(
		"ListInstances",
		"compartmentid",
		opts2,
	).Return(
		&baremetal.ListInstances{
			Instances: []baremetal.Instance{
				baremetal.Instance{
					ID:                 "id3",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "inst1",
					State:              baremetal.ResourceRunning,
					Metadata:           metadata,
					Region:             "here",
					Shape:              "round",
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
				},
				baremetal.Instance{
					ID:                 "id4",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "inst2",
					State:              baremetal.ResourceRunning,
					Metadata:           metadata,
					Region:             "here",
					Shape:              "round",
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
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
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListInstances", "compartmentid", opts2)

}

func TestResourceCoreInstancesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInstancesTestSuite))
}
