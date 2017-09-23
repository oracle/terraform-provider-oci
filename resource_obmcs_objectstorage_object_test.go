// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type ResourceObjectstorageObjectTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Object
}

func (s *ResourceObjectstorageObjectTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	data "oci_objectstorage_namespace" "t" {
	}
	
	resource "oci_objectstorage_bucket" "t" {
		compartment_id = "${var.compartment_id}"
		namespace = "${data.oci_objectstorage_namespace.t.namespace}"
		name = "-tf-bucket"
		access_type="ObjectRead"
	}`
	s.ResourceName = "oci_objectstorage_object.t"
}

func (s *ResourceObjectstorageObjectTestSuite) TestAccResourceObjectstorageObject_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_objectstorage_object" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					object = "-tf-object"
					content = "test content"
					metadata = {
						"content-type" = "text/plain"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", "-tf-bucket"),
					resource.TestCheckResourceAttr(s.ResourceName, "object", "-tf-object"),
					resource.TestCheckResourceAttr(s.ResourceName, "content", "test content"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.content-type", "text/plain"),
				),
			},
			// verify update
			{
				Config: s.Config + `
				resource "oci_objectstorage_object" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					object = "-tf-object"
					content = "{}"
					metadata = {
						"content-type" = "text/json"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "content", "{}"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.content-type", "text/json"),
				),
			},
		},
	})
}

func TestResourceObjectstorageObjectTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageObjectTestSuite))
}
