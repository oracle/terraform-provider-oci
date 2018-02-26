// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityCompartmentsTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *identity.ListCompartmentsResponse
}

func (s *DatasourceIdentityCompartmentsTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
	resource "oci_identity_compartment" "t" {
		name = "-tf-compartment"
		description = "tf test compartment"
	}`
	s.ResourceName = "data.oci_identity_compartments.t"
}

func (s *DatasourceIdentityCompartmentsTestSuite) TestAccIdentityCompartments_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_identity_compartments" "t" {
					compartment_id = "${var.tenancy_ocid}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartments.#"),
				),
			},
			{
				Config: s.Config + `
				data "oci_identity_compartments" "t" {
					compartment_id = "${var.tenancy_ocid}"
					filter {
						name   = "name"
						values = ["-tf-compartment"]
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartments.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartments.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartments.0.compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartments.0.name", "-tf-compartment"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartments.0.description", "tf test compartment"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartments.0.state", string(identity.CompartmentLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartments.0.time_created"),
					// TODO: This field is not being returned by the service call but is still showing up in the datasource
					// resource.TestCheckNoResourceAttr(s.ResourceName, "compartments.0.inactive_state"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityCompartmentsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityCompartmentsTestSuite))
}
