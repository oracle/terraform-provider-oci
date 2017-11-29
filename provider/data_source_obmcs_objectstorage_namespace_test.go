// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceObjectstorageNamespaceTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *DatasourceObjectstorageNamespaceTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	data "oci_objectstorage_namespace" "t" {
	}`
	s.ResourceName = "data.oci_objectstorage_namespace.t"
}

func (s *DatasourceObjectstorageNamespaceTestSuite) TestAccDatasourceObjectstorageNamespace_basic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
				),
			},
		},
	})
}

func TestDatasourceObjectstorageNamespaceTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceObjectstorageNamespaceTestSuite))
}
