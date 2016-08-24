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

type DatasourceCoreIPSecTestSuite struct {
	suite.Suite
	Client       *client.MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreIPSecTestSuite) SetupTest() {
	s.Client = &client.MockClient{}
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
	opts := []baremetal.Options{
		baremetal.Options{
			DrgID: "drgid",
			CpeID: "cpeid",
		},
	}

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
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListIPSecConnections", "compartmentid", opts)

}

func TestDatasourceCoreIPSecTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreIPSecTestSuite))
}
