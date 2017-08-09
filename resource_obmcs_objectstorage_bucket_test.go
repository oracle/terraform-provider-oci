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

	"github.com/oracle/terraform-provider-baremetal/client"
)

type ResourceObjectstorageBucketTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Bucket
	Namespace    baremetal.Namespace
	AccessType   baremetal.BucketAccessType
}

func (s *ResourceObjectstorageBucketTestSuite) SetupTest() {
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
			name = "name"
			namespace = "${var.namespace}"
			metadata = {
				"foo" = "bar"
			}
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_objectstorage_bucket.t"
	metadata := map[string]string{
		"foo": "bar",
	}
	s.Namespace = baremetal.Namespace("namespace")
	s.Res = &baremetal.Bucket{
		CompartmentID: "compartment_id",
		Name:          "name",
		Namespace:     s.Namespace,
		Metadata:      metadata,
		CreatedBy:     "created_by",
		TimeCreated:   s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

}

func (s *ResourceObjectstorageBucketTestSuite) TestCreateResourceObjectstorageBucket() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "name", s.Res.Name),
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
				),
			},
		},
	})
}

func (s *ResourceObjectstorageBucketTestSuite) TestDeleteResourceObjectstorageBucket() {

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

func TestResourceObjectstorageBucketTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageBucketTestSuite))
}
