// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"fmt"

	"github.com/stretchr/testify/suite"
)

type ResourceObjectstorageObjectTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *ResourceObjectstorageObjectTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + s.TokenFn(`
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
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "object", "-tf-object"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_type", "application/octet-stream"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_language", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "content_encoding", ""),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content_length"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content_md5"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.version", "1"),
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
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "content"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_type", "text/json"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_language", "*"),
					resource.TestCheckResourceAttr(s.ResourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content_length"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.version", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.modified", "10-18-2017"),
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
