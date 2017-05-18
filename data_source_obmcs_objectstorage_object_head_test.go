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

type DatasourceObjectstorageObjectHeadTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.HeadObject
}

func (s *DatasourceObjectstorageObjectHeadTestSuite) SetupTest() {
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
		resource "baremetal_objectstorage_bucket" "t" {
			compartment_id = "${var.compartment_id}"
			name = "bucketID"
			namespace = "${var.namespace}"
			metadata = {
				"foo" = "bar"
			}
		}

		resource "baremetal_objectstorage_object" "t" {
			namespace = "${var.namespace}"
			bucket = "${baremetal_objectstorage_bucket.t.name}"
			object = "objectID"
			content = "bodyContent"
			metadata = {
				"foo" = "bar"
			}
		}
		data "baremetal_objectstorage_object_head" "t" {
			namespace = "${var.namespace}"
			bucket = "${baremetal_objectstorage_bucket.t.name}"
			object = "${baremetal_objectstorage_object.t.object}"
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "data.baremetal_objectstorage_object_head.t"
}

func (s *DatasourceObjectstorageObjectHeadTestSuite) TestObjectstorageHeadObject() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", "bucketID"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.foo", "bar"),
				),
			},
		},
	})

}

func TestDatasourceobjectstorageObjectHeadTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceObjectstorageObjectHeadTestSuite))
}
