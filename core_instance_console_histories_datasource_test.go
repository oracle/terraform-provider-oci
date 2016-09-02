package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"crypto/rand"

	"github.com/stretchr/testify/suite"
)

type CoreInstanceConsoleHistoriesDatasourceTestSuite struct {
	suite.Suite
	Client       *client.MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *CoreInstanceConsoleHistoriesDatasourceTestSuite) SetupTest() {
	s.Client = &client.MockClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_instance_console_histories" "s" {
      instance_console_history_id = "ichid"
      limit = 100
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_instance_console_histories.s"

}

func (s *CoreInstanceConsoleHistoriesDatasourceTestSuite) TestResourceShowConsoleHistory() {

	consoleHistory := make([]byte, 100)
	rand.Read(consoleHistory)

	opts := []baremetal.Options{
		baremetal.Options{
			Length: 100,
			Offset: 0,
		},
	}

	s.Client.On(
		"ShowConsoleHistoryData",
		"ichid",
		opts,
	).Return(
		&baremetal.ShowConsoleHistoryMetadataResponse{
			BytesRemaining:     50,
			ConsoleHistoryData: string(consoleHistory[0:50]),
		},
		nil,
	)

	opts2 := []baremetal.Options{
		baremetal.Options{
			Length: 50,
			Offset: 50,
		},
	}

	s.Client.On(
		"ShowConsoleHistoryData",
		"ichid",
		opts2,
	).Return(
		&baremetal.ShowConsoleHistoryMetadataResponse{
			BytesRemaining:     0,
			ConsoleHistoryData: string(consoleHistory[50:]),
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
					resource.TestCheckResourceAttr(s.ResourceName, "console_history", string(consoleHistory)),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ShowConsoleHistoryData", "ichid", opts2)

}

func TestCoreInstanceConsoleHistoriesDatasource(t *testing.T) {
	suite.Run(t, new(CoreInstanceConsoleHistoriesDatasourceTestSuite))
}
