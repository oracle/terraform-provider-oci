// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type ResourceObjectstorageBucketTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceObjectstorageBucketTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
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
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.content-type", "text/plain"),
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
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token+"-changed"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.content-type", "text/plain"),
				),
			},
		},
	})
}

func TestResourceObjectstorageBucketTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageBucketTestSuite))
}
