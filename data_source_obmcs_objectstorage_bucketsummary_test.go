// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceObjectstorageBucketSummaryTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *DatasourceObjectstorageBucketSummaryTestSuite) SetupTest() {
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
	}`, nil)
	s.ResourceName = "data.oci_objectstorage_bucket_summaries.t"
}

func (s *DatasourceObjectstorageBucketSummaryTestSuite) TestAccDatasourceObjectstorageBucketSummaries_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "oci_objectstorage_bucket_summaries" "t" {
					compartment_id = "${var.compartment_id}"
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "bucket_summaries.#"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.0.name", s.Token),
				),
			},
		},
	},
	)
}

func TestDatasourceObjectstorageBucketSummaryTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceObjectstorageBucketSummaryTestSuite))
}
