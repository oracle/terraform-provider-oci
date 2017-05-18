// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"crypto/rand"

	"github.com/stretchr/testify/suite"
)

type CoreConsoleHistoryDataDatasourceTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *CoreConsoleHistoryDataDatasourceTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `

    resource "baremetal_core_console_history" "t" {
			instance_id = "instance_id"
    }
    data "baremetal_core_console_history_data" "s" {
      console_history_id = "${baremetal_core_console_history.t.id}"
      length = 1
      offset = 1
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_console_history_data.s"
}

func (s *CoreConsoleHistoryDataDatasourceTestSuite) TestResourceShowConsoleHistory() {
	data := make([]byte, 100)
	rand.Read(data)

	opts := &baremetal.ConsoleHistoryDataOptions{}
	opts.Length = 1
	opts.Offset = 1

	s.Client.On("ShowConsoleHistoryData", "ichid", opts).
		Return(
			&baremetal.ConsoleHistoryData{
				BytesRemaining: 50,
				Data:           string(data),
			},
			nil,
		)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "data", string(data)),
				),
			},
		},
	},
	)
}

func TestCoreInstanceConsoleHistoriesDatasource(t *testing.T) {
	suite.Run(t, new(CoreConsoleHistoryDataDatasourceTestSuite))
}
