// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreConsoleHistoryTestSuite struct {
	suite.Suite
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreConsoleHistoryTestSuite) SetupTest() {
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + instanceConfig

	p := s.Provider.(*schema.Provider)
	res := p.ResourcesMap["oci_core_console_history"]
	res.Delete = func(d *schema.ResourceData, m interface{}) (e error) {
		return nil
	}

	s.ResourceName = "oci_core_console_history.t"
}

func (s *ResourceCoreConsoleHistoryTestSuite) TestAccResourceCoreInstanceConsoleHistory_basic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_console_history" "t" {
					instance_id = "${oci_core_instance.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.ConsoleHistoryLifecycleStateSucceeded)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func TestResourceCoreConsoleHistoryTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreConsoleHistoryTestSuite))
}
