// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

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
	s.Token, s.TokenFn = acctest.TokenizeWithHttpReplay("object_storage_data_source")
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + s.TokenFn(`
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
		storage_tier = "InfrequentAccess"
		metadata = {
			"content-type" = "text/plain"
		}
	}`, nil)
	s.ResourceName = "data.oci_objectstorage_object_head.t"
}

func (s *DatasourceObjectstorageObjectHeadTestSuite) TestDatasourceObjectHead_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + `
				data "oci_objectstorage_object_head" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					object = "${oci_objectstorage_object.t.object}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "object", "-tf-object"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content_type"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "content_length"),
					resource.TestCheckResourceAttr(s.ResourceName, "storage_tier", "InfrequentAccess"),
				),
			},
		},
	})
}

// issue-routing-tag: object_storage/default
func TestDatasourceObjectstorageObjectHeadTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceObjectstorageObjectHeadTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceObjectstorageObjectHeadTestSuite))
}
