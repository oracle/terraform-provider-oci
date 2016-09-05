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

type DatasourceCoreVnicTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVnicTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_vnic" "t" {
      vnic_id = "vnicid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_vnic.t"
}

func (s *DatasourceCoreVnicTestSuite) TestReadVnic() {

	s.Client.On(
		"GetVnic",
		"vnicid",
	).Return(
		&baremetal.Vnic{
			AvailabilityDomain: "availabilitydomain",
			CompartmentID:      "compartmentid",
			DisplayName:        "displayname",
			ID:                 "vncid",
			State:              baremetal.ResourceActive,
			PrivateIPAddress:   "10.10.10.10",
			PublicIPAddress:    "52.53.54.55",
			SubnetID:           "subnetid",
			TimeCreated:        baremetal.Time{Time: time.Now()},
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
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availabilitydomain"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceActive),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ip_address", "10.10.10.10"),
					resource.TestCheckResourceAttr(s.ResourceName, "public_ip_address", "52.53.54.55"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "GetVnic", "vnicid")
}

func TestDatasourceCoreVnicTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVnicTestSuite))
}
