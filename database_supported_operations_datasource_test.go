// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/client/mocks"

	"github.com/stretchr/testify/suite"
)

type DatasourceDatabaseSupportedOperationTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.ListSupportedOperations
}

func (s *DatasourceDatabaseSupportedOperationTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
		data "baremetal_database_supported_operations" "t" {}
	`

	s.Config += testProviderConfig

	s.ResourceName = "data.baremetal_database_supported_operations.t"
	s.Res = &baremetal.ListSupportedOperations{

		SupportedOperations: []baremetal.SupportedOperation{
			{
				ID: "test-1",
			},
			{
				ID: "test-2",
			},
		},
	}
}

func (s *DatasourceDatabaseSupportedOperationTestSuite) TestDatabaseListSupportedOperations() {
	s.Client.On("ListSupportedOperations").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "supported_operations.0.id", "test-1"),
				),
			},
		},
	})
	s.Client.AssertCalled(s.T(), "ListSupportedOperations")
}

func TestDatasourceDatabaseSupportedOperationTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceDatabaseSupportedOperationTestSuite))
}
