// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceObjectstorageBucketSummaryTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *DatasourceObjectstorageBucketSummaryTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + s.TokenFn(`
	data "oci_objectstorage_namespace" "t" {
	}
	resource "oci_objectstorage_bucket" "t" {
		compartment_id = "${var.compartment_id}"
		namespace = "${data.oci_objectstorage_namespace.t.namespace}"
		name = "{{.token}}"
	}
	resource "oci_objectstorage_bucket" "u" {
		compartment_id = "${var.compartment_id}"
		namespace = "${data.oci_objectstorage_namespace.t.namespace}"
		name = "{{.otherToken}}"
	}`, map[string]string{"otherToken": s.Token + "-2"})
	s.ResourceName = "data.oci_objectstorage_bucket_summaries.t"
}

func (s *DatasourceObjectstorageBucketSummaryTestSuite) TestAccDatasourceObjectstorageBucketSummaries_basic() {
	compartmentID := getCompartmentIDForLegacyTests()
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			// Client-side filtering.
			{
				Config: s.Config + s.TokenFn(`
				data "oci_objectstorage_bucket_summaries" "t" {
					compartment_id = "${var.compartment_id}"
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					filter {
						name = "name"
						values = ["{{.token}}"]
					}
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					TestCheckResourceAttributesEqual(s.ResourceName, "namespace", "data.oci_objectstorage_namespace.t", "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.#", "1"),
					TestCheckResourceAttributesEqual(s.ResourceName, "bucket_summaries.0.name", "oci_objectstorage_bucket.t", "name"),
				),
			},
			{
				Config: s.Config + s.TokenFn(`
				data "oci_objectstorage_bucket_summaries" "t" {
					compartment_id = "${var.compartment_id}"
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					filter {
						name = "name"
						values = ["{{.otherToken}}"]
					}
				}`, map[string]string{"otherToken": s.Token + "-2"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					TestCheckResourceAttributesEqual(s.ResourceName, "namespace", "data.oci_objectstorage_namespace.t", "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.#", "1"),
					TestCheckResourceAttributesEqual(s.ResourceName, "bucket_summaries.0.name", "oci_objectstorage_bucket.u", "name"),
				),
			},
			{
				Config: s.Config + `
				data "oci_objectstorage_bucket_summaries" "t" {
					compartment_id = "${var.compartment_id}"
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					filter {
						name = "name"
						values = ["non-existent-bucket"]
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					TestCheckResourceAttributesEqual(s.ResourceName, "namespace", "data.oci_objectstorage_namespace.t", "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.#", "0"),
				),
			},
		},
	},
	)
}

func TestDatasourceObjectstorageBucketSummaryTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceObjectstorageBucketSummaryTestSuite))
}
