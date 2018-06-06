// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"fmt"

	"regexp"

	"github.com/stretchr/testify/suite"
)

type ResourceObjectstorageObjectTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *ResourceObjectstorageObjectTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + s.TokenFn(`
	data "oci_objectstorage_namespace" "t" {
	}
	
	resource "oci_objectstorage_bucket" "t" {
		compartment_id = "${var.compartment_id}"
		namespace = "${data.oci_objectstorage_namespace.t.namespace}"
		name = "{{.token}}"
		access_type="ObjectRead"
	}`, nil)
	s.ResourceName = "oci_objectstorage_object.t"
}

func (s *ResourceObjectstorageObjectTestSuite) TestAccResourceObjectstorageObject_basic() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify error from validation function if metadata key is not lowercase
			{
				Config: s.Config + `
				resource "oci_objectstorage_object" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					object = "-tf-object"
					content = "123"
					metadata = {
						"version" = "1"
						"MY-KEY" = "aBc"
					}
				}`,
				ExpectError: regexp.MustCompile("All 'metadata' keys must be lowercase"),
			},
			// verify create with required only
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_objectstorage_object" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					object = "-tf-object-required-only"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "object", "-tf-object-required-only"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_type", "application/octet-stream"),
					// New SDK doesn't set omitted values from response, check they are missing from state.
					resource.TestCheckNoResourceAttr(s.ResourceName, "content"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "content_language"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "content_encoding"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_length", "0"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content_md5"),
				),
			},
			// verify create with expected defaults
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_objectstorage_object" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					object = "-tf-object"
					content = "123"
					metadata = {
						"version" = "1"
						"my-key" = "aBc"
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "object", "-tf-object"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_type", "application/octet-stream"),
					// New SDK doesn't set omitted values from response, check they are missing from state.
					resource.TestCheckNoResourceAttr(s.ResourceName, "content_language"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "content_encoding"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content_length"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content_md5"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.version", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.my-key", "aBc"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_objectstorage_object.t", "content")
						return err
					},
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
					content_type = "text/json"
					content_language = "*"
					content_encoding = "identity"
					metadata = {
						"version" = "2"
						"modified" = "10-18-2017"
						"my-key" = "ABC"
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "content"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_type", "text/json"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_language", "*"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content_length"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.version", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.modified", "10-18-2017"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.my-key", "ABC"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_objectstorage_object.t", "content")
						if resId == resId2 {
							return fmt.Errorf("Expected different content hash, got same.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceObjectstorageObjectTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageObjectTestSuite))
}
