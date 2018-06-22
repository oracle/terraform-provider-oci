// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"fmt"

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
	s.Config = legacyTestProviderConfig() + instanceConfig + DefinedTagsDependencies

	p := s.Provider.(*schema.Provider)
	res := p.ResourcesMap["oci_core_console_history"]
	res.Delete = func(d *schema.ResourceData, m interface{}) (e error) {
		return nil
	}

	s.ResourceName = "oci_core_console_history.t"
}

func (s *ResourceCoreConsoleHistoryTestSuite) TestAccResourceCoreInstanceConsoleHistory_basic() {
	var consoleHistoryId string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_console_history" "t" {
					instance_id = "${oci_core_instance.t.id}"

					#Optional
					defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value"
									)}"
                    freeform_tags = { "Department" = "Accounting"}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.ConsoleHistoryLifecycleStateSucceeded)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					func(ts *terraform.State) (err error) {
						consoleHistoryId, err = fromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_console_history" "t" {
					instance_id = "${oci_core_instance.t.id}"

					#Optional
					display_name = "updatedDisplayName"
					defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue"
									)}"
                    freeform_tags = { "Department" = "Finance"}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "freeform_tags.%", "1"),
					func(ts *terraform.State) (err error) {
						newId, err := fromInstanceState(ts, s.ResourceName, "id")
						if newId != consoleHistoryId {
							return fmt.Errorf("expected same console history ocid, got different")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceCoreConsoleHistoryTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreConsoleHistoryTestSuite))
}
