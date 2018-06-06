// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceObjectstorageObjectHeadTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *DatasourceObjectstorageObjectHeadTestSuite) SetupTest() {
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
		content = "test content"
		metadata = {
			"content-type" = "text/plain"
		}
	}`, nil)
	s.ResourceName = "data.oci_objectstorage_object_head.t"
}

func (s *DatasourceObjectstorageObjectHeadTestSuite) TestObjectstorageObjectHead_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_objectstorage_object_head" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					object = "${oci_objectstorage_object.t.object}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "object", "-tf-object"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content-type"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content-length"),
				),
			},
		},
	})
}

func TestDatasourceObjectstorageObjectHeadTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceObjectstorageObjectHeadTestSuite))
}
