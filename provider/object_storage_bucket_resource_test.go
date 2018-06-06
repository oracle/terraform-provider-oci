// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/objectstorage"
	"github.com/stretchr/testify/suite"
)

type ResourceObjectstorageBucketTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceObjectstorageBucketTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
	data "oci_objectstorage_namespace" "t" {
	}`

	s.ResourceName = "oci_objectstorage_bucket.t"
}

func (s *ResourceObjectstorageBucketTestSuite) TestAccResourceObjectstorageBucket_basic() {
	token, tokenFn := tokenize()
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + tokenFn(`
				resource "oci_objectstorage_bucket" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					compartment_id = "${var.compartment_id}"
					name = "{{.token}}"
					metadata = {
						"content-type" = "text/plain"
					}
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttr(s.ResourceName, "access_type", string(objectstorage.BucketPublicAccessTypeNopublicaccess)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "created_by"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "etag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
			// test update
			{
				Config: s.Config + tokenFn(`
				resource "oci_objectstorage_bucket" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					compartment_id = "${var.compartment_id}"
					name = "{{.token}}-changed"
					metadata = {
						"content-type" = "text/plain"
					}
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token+"-changed"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttr(s.ResourceName, "access_type", string(objectstorage.BucketPublicAccessTypeNopublicaccess)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "created_by"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "etag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func TestResourceObjectstorageBucketTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageBucketTestSuite))
}
