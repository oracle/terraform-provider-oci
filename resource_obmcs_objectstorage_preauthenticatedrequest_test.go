// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type ResourceObjectstoragePARTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *ResourceObjectstoragePARTestSuite) SetupTest() {
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

	s.ResourceName = "oci_objectstorage_preauthrequest.t"
}

func (s *ResourceObjectstoragePARTestSuite) TestAccResourceObjectstoragePAR_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_objectstorage_preauthrequest" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					name = "-tf-par"
					access_type = "AnyObjectWrite"
					time_expires = "2019-11-10T23:00:00Z"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "access_type", "AnyObjectWrite"),
					resource.TestCheckResourceAttr(s.ResourceName, "time_expires", "2019-11-10T23:00:00Z"),
				),
			},
		},
	})
}

func TestResourceObjectstoragePARTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstoragePARTestSuite))
}
