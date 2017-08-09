// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DatasourceObjectstorageNamespaceTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Namespace
}

func (s *DatasourceObjectstorageNamespaceTestSuite) SetupTest() {
	s.Client = GetTestProvider()

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
		data "baremetal_objectstorage_namespace" "t" {}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_objectstorage_namespace.t"
	namespace := baremetal.Namespace("namespaceID")
	s.Res = &namespace
}

func (s *DatasourceObjectstorageNamespaceTestSuite) TestObjectstorageNamespace() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "namespace", "namespaceID"),
				),
			},
		},
	})

}

func TestDatasourceobjectstorageNamespaceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageObjectTestSuite))
}
