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

type ResourceIdentityCompartmentTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityCompartmentTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig()
	s.ResourceName = "oci_identity_compartment.t"
}

func (s *ResourceIdentityCompartmentTestSuite) TestAccResourceIdentityCompartment_basic() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// compartments are permanent, sync to existing test compartment
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_identity_compartment" "t" {
					name = "terraform-update-test-compartment"
					description = "for name and description update tests"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "terraform-update-test-compartment"),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "for name and description update tests"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceActive),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_identity_compartment.t", "id")
						return err
					},
				),
			},
			// verify update
			{
				Config: s.Config + `
				resource "oci_identity_compartment" "t" {
					name = "terraform-update-test-compartment2"
					description = "for name and description update tests2"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "terraform-update-test-compartment2"),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "for name and description update tests2"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_identity_compartment.t", "id")
						if resId != resId2 {
							return fmt.Errorf("Expected same ocid, got the different.")
						}
						return err
					},
				),
			},
			// restore compartment to original state (for future tests)
			{
				Config: s.Config + `
				resource "oci_identity_compartment" "t" {
					name = "terraform-update-test-compartment"
					description = "for name and description update tests"
				}`,
			},
		},
	})
}

func TestResourceIdentityCompartmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityCompartmentTestSuite))
}
