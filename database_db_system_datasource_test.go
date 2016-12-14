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

type DBSystemDatasourceTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DBSystemDatasourceTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{"baremetal": s.Provider}
	s.Config = `
    data "baremetal_database_db_systems" "t" {
      compartment_id = "compartmentid"
      limit = 1
      page = "page"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_database_db_systems.t"
}

func (s *DBSystemDatasourceTestSuite) TestReadDBSystems() {
	opts := &baremetal.PageListOptions{}
	opts.Page = "page"
	s.Client.On(
		"ListDBSystems", "compartmentid", uint64(1), opts,
	).Return(
		&baremetal.ListDBSystems{
			ResourceContainer: baremetal.ResourceContainer{
				NextPage: "nextpage",
			},
			DBSystems: []baremetal.DBSystem{
				{Shape: "shape1"},
				{Shape: "shape2"},
			},
		},
		nil,
	)

	opts2 := &baremetal.PageListOptions{}
	opts2.Page = "nextpage"
	s.Client.On(
		"ListDBSystems", "compartmentid", uint64(1), opts2,
	).Return(
		&baremetal.ListDBSystems{
			DBSystems: []baremetal.DBSystem{
				{Shape: "shape3"},
				{Shape: "shape4"},
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
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_systems.0.shape", "shape1"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_systems.3.shape", "shape4"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_systems.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(
		s.T(), "ListDBSystems", "compartmentid", uint64(1), opts2,
	)
}

func TestDBSystemDatasourceTestSuite(t *testing.T) {
	suite.Run(t, new(DBSystemDatasourceTestSuite))
}
