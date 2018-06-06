// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceObjectstorageObjectTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *DatasourceObjectstorageObjectTestSuite) SetupTest() {
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
	}
	
	resource "oci_objectstorage_object" "t" {
		namespace = "${data.oci_objectstorage_namespace.t.namespace}"
		bucket = "${oci_objectstorage_bucket.t.name}"
		object = "-tf-object"
		content = "123"
	}
	
	resource "oci_objectstorage_object" "u" {
		namespace = "${data.oci_objectstorage_namespace.t.namespace}"
		bucket = "${oci_objectstorage_bucket.t.name}"
		object = "-tf-object2"
		content = "456"
	}`, nil)

	s.ResourceName = "data.oci_objectstorage_objects.t"
}

func (s *DatasourceObjectstorageObjectTestSuite) TestAccDatasourceObjectstorageObjects_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_objectstorage_objects" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
				}`,
			},
			{
				Config: s.Config + `
				data "oci_objectstorage_objects" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					TestCheckResourceAttributesEqual(s.ResourceName, "namespace", "data.oci_objectstorage_namespace.t", "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "objects.#", "2"),
				),
			},
			// Client-side filtering.
			{
				Config: s.Config + `
				data "oci_objectstorage_objects" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					filter {
						name = "name"
						values = ["-tf-object"]
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					TestCheckResourceAttributesEqual(s.ResourceName, "namespace", "data.oci_objectstorage_namespace.t", "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "objects.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "objects.0.name", "-tf-object"),
					resource.TestCheckResourceAttr(s.ResourceName, "objects.0.size", "3"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "objects.0.md5"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "objects.0.time_created"),
				),
			},
		},
	})
}

func TestDatasourceObjectstorageObjectTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceObjectstorageObjectTestSuite))
}
