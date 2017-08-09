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

type DatasourceObjectstorageObjectTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.ListObjects
	Res2         *baremetal.ListObjects
}

func (s *DatasourceObjectstorageObjectTestSuite) SetupTest() {
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
		data "baremetal_objectstorage_objects" "t" {
			namespace = "namespaceID"
			bucket = "bucketID"
			prefix = "testprefix"
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_objectstorage_objects.t"
	s.Res = &baremetal.ListObjects{
		NextStartWith: "testprefix-2",
		Objects: []baremetal.ObjectSummary{
			{
				Name: "testprefix-1",
			},
			{
				Name: "testprefix-2",
			},
		},
	}
	s.Res2 = &baremetal.ListObjects{

		Objects: []baremetal.ObjectSummary{
			{
				Name: "testprefix-3",
			},
			{
				Name: "testprefix-4",
			},
		},
	}
}

func (s *DatasourceObjectstorageObjectTestSuite) TestObjectstorageListObjects() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "prefix", "testprefix"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", "bucketID"),
					resource.TestCheckResourceAttr(s.ResourceName, "namespace", "namespaceID"),
					resource.TestCheckResourceAttr(s.ResourceName, "namespace", "namespaceID"),
					resource.TestCheckResourceAttr(s.ResourceName, "objects.2.name", "testprefix-3"),
				),
			},
		},
	})

}

func TestDatasourceobjectstorageObjectTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageObjectTestSuite))
}
