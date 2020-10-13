// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"fmt"

	"github.com/oracle/oci-go-sdk/v27/identity"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentityGroupTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityGroupTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig()
	s.ResourceName = "oci_identity_group.t"
}

func (s *ResourceIdentityGroupTestSuite) TestAccResourceIdentityGroup_basic() {
	var resId, resId2 string
	token, tokenFn := tokenizeWithHttpReplay("identity_group_resource")
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create w/ compartment
			{
				Config: s.Config + tokenFn(`
				resource "oci_identity_group" "t0" {
					name = "{{.token}}"
					description = "tf test group"
					compartment_id = "${var.compartment_id}"
				}`, nil),
				ExpectError: regexp.MustCompile("Tenant id is not equal to compartment id"),
			},
			{
				Config: s.Config + tokenFn(`
				resource "oci_identity_group" "t0" {
					name = "{{.token}}"
					description = "tf test group"
					compartment_id = "${var.tenancy_ocid}"
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName+"0", "compartment_id", getEnvSettingWithBlankDefault("tenancy_ocid")),
					resource.TestCheckResourceAttr(s.ResourceName+"0", "name", token),
					resource.TestCheckResourceAttr(s.ResourceName+"0", "description", "tf test group"),
					resource.TestCheckResourceAttr(s.ResourceName+"0", "state", string(identity.GroupLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName+"0", "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName+"0", "inactive_state"),
				),
			},
			// verify create w/o compartment, verify that it defaults to tenancy
			{
				Config: s.Config + tokenFn(identityGroupTestStepConfigFn("tf test group"), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", getEnvSettingWithBlankDefault("tenancy_ocid")),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test group"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.GroupLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_identity_group.t", "id")
						return err
					},
				),
			},
			// verify update
			{
				Config: s.Config + tokenFn(identityGroupTestStepConfigFn("tf test group (updated)"), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test group (updated)"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.GroupLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_identity_group.t", "id")
						if resId != resId2 {
							return fmt.Errorf("resource was recreated when it should not have been")
						}
						return err
					},
				),
			},
		},
	})
}

func identityGroupTestStepConfigFn(description string) string {
	useDelegationToken := strings.EqualFold(os.Getenv("DELEGATION_TOKEN"), "true")
	if useDelegationToken {
		return fmt.Sprintf(`
				resource "oci_identity_group" "t" {
					name = "{{.token}}"
					description = "%s"
					compartment_id = "${var.tenancy_ocid}"
				}`, description)
	}

	return fmt.Sprintf(`
				resource "oci_identity_group" "t" {
					name = "{{.token}}"
					description = "%s"
				}`, description)
}

func TestResourceIdentityGroupTestSuite(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip TestResourceIdentityGroupTestSuite for httpreplay mode.")
	}
	suite.Run(t, new(ResourceIdentityGroupTestSuite))
}
