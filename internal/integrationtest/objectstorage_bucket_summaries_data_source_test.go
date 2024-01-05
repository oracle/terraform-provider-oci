// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/stretchr/testify/suite"
)

type DatasourceObjectstorageBucketSummaryTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *DatasourceObjectstorageBucketSummaryTestSuite) SetupTest() {
	s.Token, s.TokenFn = acctest.TokenizeWithHttpReplay("objectstorage")
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + s.TokenFn(`
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
	compartmentID := acctest.GetCompartmentIDForLegacyTests()
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "namespace", "data.oci_objectstorage_namespace.t", "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "bucket_summaries.0.name", "oci_objectstorage_bucket.t", "name"),
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "namespace", "data.oci_objectstorage_namespace.t", "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "bucket_summaries.0.name", "oci_objectstorage_bucket.u", "name"),
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "namespace", "data.oci_objectstorage_namespace.t", "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.#", "0"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: object_storage/default
func TestDatasourceObjectstorageBucketSummaryTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceObjectstorageBucketSummaryTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceObjectstorageBucketSummaryTestSuite))
}
