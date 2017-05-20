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

type ResourceObjectstorageObjectTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Object
}

func (s *ResourceObjectstorageObjectTestSuite) SetupTest() {
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
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_objectstorage_object.t"

}

func (s *ResourceObjectstorageObjectTestSuite) TestCreateResourceObjectstorageObject() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "content", "bodyContent"),
				),
			},
		},
	})
}

func (s *ResourceObjectstorageObjectTestSuite) TestUpdateResourceObjectstorageObject() {

	config := `
		resource "baremetal_objectstorage_bucket" "t" {
			compartment_id = "${var.compartment_id}"
			name = "bucketID"
			namespace = "${var.namespace}"
			metadata = {
				"foo" = "bar"
			}
		}

		resource "baremetal_objectstorage_object" "t" {
			object = "objectID"
			bucket = "${baremetal_objectstorage_bucket.t.name}"
			namespace = "${var.namespace}"
			content = "bodyContent2"
			metadata = {
				"foo" = "bar"
			}
		}
	`
	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "content", "bodyContent2"),
				),
			},
		},
	})
}

func (s *ResourceObjectstorageObjectTestSuite) TestDeleteResourceObjectstorageObject() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

}

func TestResourceObjectstorageObjectTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageObjectTestSuite))
}
