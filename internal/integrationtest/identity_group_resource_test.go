// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"fmt"

	"github.com/oracle/oci-go-sdk/v56/identity"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentityGroupTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityGroupTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig()
	s.ResourceName = "oci_identity_group.t"
}

func (s *ResourceIdentityGroupTestSuite) TestAccResourceIdentityGroup_basic() {
	var resId, resId2 string
	token, tokenFn := acctest.TokenizeWithHttpReplay("identity_group_resource")
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify Create w/ compartment
			{
				Config: s.Config + tokenFn(`
				resource "oci_identity_group" "t0" {
					name = "{{.token}}"
					description = "tf test Group"
					compartment_id = "${var.compartment_id}"
				}`, nil),
				ExpectError: regexp.MustCompile("Tenant id is not equal to compartment id"),
			},
			{
				Config: s.Config + tokenFn(`
				resource "oci_identity_group" "t0" {
					name = "{{.token}}"
					description = "tf test Group"
					compartment_id = "${var.tenancy_ocid}"
				}`, nil),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName+"0", "compartment_id", utils.GetEnvSettingWithBlankDefault("tenancy_ocid")),
					resource.TestCheckResourceAttr(s.ResourceName+"0", "name", token),
					resource.TestCheckResourceAttr(s.ResourceName+"0", "description", "tf test Group"),
					resource.TestCheckResourceAttr(s.ResourceName+"0", "state", string(identity.GroupLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName+"0", "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName+"0", "inactive_state"),
				),
			},
			// verify Create w/o compartment, verify that it defaults to tenancy
			{
				Config: s.Config + tokenFn(identityGroupTestStepConfigFn("tf test Group"), nil),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", utils.GetEnvSettingWithBlankDefault("tenancy_ocid")),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test Group"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.GroupLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_identity_group.t", "id")
						return err
					},
				),
			},
			// verify Update
			{
				Config: s.Config + tokenFn(identityGroupTestStepConfigFn("tf test Group (updated)"), nil),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test Group (updated)"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.GroupLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, "oci_identity_group.t", "id")
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

// issue-routing-tag: identity/default
func TestResourceIdentityGroupTestSuite(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip TestResourceIdentityGroupTestSuite for httpreplay mode.")
	}
	suite.Run(t, new(ResourceIdentityGroupTestSuite))
}
