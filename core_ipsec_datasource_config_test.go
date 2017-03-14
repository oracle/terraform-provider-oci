// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreIPSecConfigTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreIPSecConfigTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_ipsec_config" "s" {
      ipsec_id = "ipsecid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_ipsec_config.s"

}

func (s *DatasourceCoreIPSecConfigTestSuite) TestIPSecConfig() {

	s.Client.On(
		"GetIPSecConnectionDeviceConfig",
		"ipsecid",
	).Return(
		&baremetal.IPSecConnectionDeviceConfig{
			IPSecConnectionDevice: baremetal.IPSecConnectionDevice{
				CompartmentID: "compartmentid",
				ID:            "id",
				TimeCreated:   baremetal.Time{Time: time.Now()},
			},

			Tunnels: []baremetal.TunnelConfig{
				{
					IPAddress:    "10.10.10.2",
					SharedSecret: "secret1",
					TimeCreated:  baremetal.Time{Time: time.Now()},
				},
				{
					IPAddress:    "10.10.10.3",
					SharedSecret: "secret2",
					TimeCreated:  baremetal.Time{Time: time.Now()},
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "id", "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "tunnels.0.ip_address", "10.10.10.2"),
					resource.TestCheckResourceAttr(s.ResourceName, "tunnels.0.shared_secret", "secret1"),
					resource.TestCheckResourceAttr(s.ResourceName, "tunnels.1.ip_address", "10.10.10.3"),
					resource.TestCheckResourceAttr(s.ResourceName, "tunnels.1.shared_secret", "secret2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "GetIPSecConnectionDeviceConfig", "ipsecid")

}

func TestDatasourceCoreIPSecConfigTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreIPSecConfigTestSuite))
}
