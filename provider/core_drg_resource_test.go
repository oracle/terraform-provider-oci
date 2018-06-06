// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"fmt"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreDrgTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreDrgTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig()
	s.ResourceName = "oci_core_drg.t"
}

func (s *ResourceCoreDrgTestSuite) TestAccResourceCoreDrg_basic() {
	var resId, resId2 string

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify a drg can be created
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: legacyTestProviderConfig() + `
				resource "oci_core_drg" "t" {
					compartment_id = "${var.compartment_id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("oci_core_drg.t", "id"),
					resource.TestCheckResourceAttrSet("oci_core_drg.t", "time_created"),
					resource.TestCheckResourceAttrSet("oci_core_drg.t", "display_name"),
					resource.TestCheckResourceAttr("oci_core_drg.t", "state", string(core.DrgLifecycleStateAvailable)),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_drg.t", "id")
						return err
					},
				),
			},
			// verify drg update
			{
				Config: legacyTestProviderConfig() + `
				resource "oci_core_drg" "t" {
					compartment_id = "${var.compartment_id}"
					display_name = "-tf-drg"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("oci_core_drg.t", "display_name", "-tf-drg"),
					resource.TestCheckResourceAttrSet("oci_core_drg.t", "time_created"),
					resource.TestCheckResourceAttr("oci_core_drg.t", "state", string(core.DrgLifecycleStateAvailable)),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_core_drg.t", "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceCoreDrgTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDrgTestSuite))
}
