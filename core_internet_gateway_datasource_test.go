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

type CoreInternetGatewayDatasourceTestSuite struct {
	suite.Suite
	Client       *client.MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *CoreInternetGatewayDatasourceTestSuite) SetupTest() {
	s.Client = &client.MockClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_internet_gateways" "s" {
      compartment_id = "compartmentid"
      vcn_id = "vcnid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_internet_gateways.s"

}

func (s *CoreInternetGatewayDatasourceTestSuite) TestResourceListIPConnections() {

	s.Client.On(
		"ListInternetGateways",
		"compartmentid",
		"vcnid",
		[]baremetal.Options{},
	).Return(
		&baremetal.ListInternetGateways{
			Gateways: []baremetal.InternetGateway{
				baremetal.InternetGateway{
					CompartmentID: "compartmentid",
					DisplayName:   "display_name",
					ID:            "id1",
					State:         baremetal.ResourceAvailable,
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
					ModifiedTime: baremetal.Time{
						Time: time.Now(),
					},
				},
				baremetal.InternetGateway{
					CompartmentID: "compartmentid",
					DisplayName:   "display_name",
					ID:            "id2",
					State:         baremetal.ResourceAvailable,
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
					ModifiedTime: baremetal.Time{
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
					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", "vcnid"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.1.id", "id2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListInternetGateways", "compartmentid", "vcnid", []baremetal.Options{})

}

func TestCoreInternetGatewayDatasource(t *testing.T) {
	suite.Run(t, new(CoreInternetGatewayDatasourceTestSuite))
}
