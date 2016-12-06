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

type DatasourceCoreIPSecTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreIPSecTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_ipsec_connections" "s" {
      compartment_id = "compartmentid"
      cpe_id = "cpeid"
      drg_id = "drgid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_ipsec_connections.s"

}

func (s *DatasourceCoreIPSecTestSuite) TestResourceListIPConnections() {
	opts := &baremetal.ListIPSecConnsOptions{}
	opts.DrgID = "drgid"
	opts.CpeID = "cpeid"

	s.Client.On(
		"ListIPSecConnections",
		"compartmentid",
		opts,
	).Return(
		&baremetal.ListIPSecConnections{
			Connections: []baremetal.IPSecConnection{
				baremetal.IPSecConnection{
					CompartmentID: "compartmentid",
					CpeID:         "cpeid",
					DisplayName:   "display_name",
					DrgID:         "drgid",
					ID:            "id1",
					State:         baremetal.ResourceUp,
					StaticRoutes: []string{
						"route1",
						"route2",
					},
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
				},
				baremetal.IPSecConnection{
					CompartmentID: "compartmentid",
					CpeID:         "cpeid",
					DisplayName:   "display_name",
					DrgID:         "drgid",
					ID:            "id2",
					State:         baremetal.ResourceUp,
					StaticRoutes: []string{
						"route1",
						"route2",
					},
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
					resource.TestCheckResourceAttr(s.ResourceName, "drg_id", "drgid"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpe_id", "cpeid"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.0.compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListIPSecConnections", "compartmentid", opts)

}

func (s *DatasourceCoreIPSecTestSuite) TestResourceListPagedIPConnections() {
	opts := baremetal.ListIPSecConnsOptions{}
	opts.DrgID = "drgid"
	opts.CpeID = "cpeid"

	s.Client.On(
		"ListIPSecConnections",
		"compartmentid",
		opts,
	).Return(
		&baremetal.ListIPSecConnections{
			ResourceContainer: baremetal.ResourceContainer{
				NextPage: "nextpage",
			},
			Connections: []baremetal.IPSecConnection{
				baremetal.IPSecConnection{
					CompartmentID: "compartmentid",
					CpeID:         "cpeid",
					DisplayName:   "display_name",
					DrgID:         "drgid",
					ID:            "id1",
					State:         baremetal.ResourceUp,
					StaticRoutes: []string{
						"route1",
						"route2",
					},
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
				},
				baremetal.IPSecConnection{
					CompartmentID: "compartmentid",
					CpeID:         "cpeid",
					DisplayName:   "display_name",
					DrgID:         "drgid",
					ID:            "id2",
					State:         baremetal.ResourceUp,
					StaticRoutes: []string{
						"route1",
						"route2",
					},
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
				},
			},
		},
		nil,
	)

	opts2 := baremetal.ListIPSecConnsOptions{}
	opts2.DrgID = "drgid"
	opts2.CpeID = "cpeid"
	opts2.Page = "nextpage"

	s.Client.On(
		"ListIPSecConnections",
		"compartmentid",
		opts2,
	).Return(
		&baremetal.ListIPSecConnections{
			Connections: []baremetal.IPSecConnection{
				baremetal.IPSecConnection{
					CompartmentID: "compartmentid",
					CpeID:         "cpeid",
					DisplayName:   "display_name",
					DrgID:         "drgid",
					ID:            "id3",
					State:         baremetal.ResourceUp,
					StaticRoutes: []string{
						"route1",
						"route2",
					},
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
				},
				baremetal.IPSecConnection{
					CompartmentID: "compartmentid",
					CpeID:         "cpeid",
					DisplayName:   "display_name",
					DrgID:         "drgid",
					ID:            "id4",
					State:         baremetal.ResourceUp,
					StaticRoutes: []string{
						"route1",
						"route2",
					},
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
					resource.TestCheckResourceAttr(s.ResourceName, "drg_id", "drgid"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpe_id", "cpeid"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.0.compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListIPSecConnections", "compartmentid", opts2)

}

func TestDatasourceCoreIPSecTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreIPSecTestSuite))
}
