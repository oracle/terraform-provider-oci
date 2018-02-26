// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	UserGroupMembershipResourceConfig = UserGroupMembershipResourceDependencies + `
resource "oci_identity_user_group_membership" "test_user_group_membership" {
	#Required
	group_id = "${oci_identity_group.test_group.id}"
	user_id = "${oci_identity_user.test_user.id}"
}
`

	UserGroupMembershipResourceConfig2 = UserGroupMembershipResourceDependencies + `
resource "oci_identity_user_group_membership" "test_user_group_membership" {
	#Required
	group_id = "${oci_identity_group.test_group2.id}"
	user_id = "${oci_identity_user.test_user.id}"
}
`

	UserGroupMembershipResourceConfig3 = UserGroupMembershipResourceDependencies + `
resource "oci_identity_user_group_membership" "test_user_group_membership" {
	#Required
	group_id = "${oci_identity_group.test_group.id}"
	user_id = "${oci_identity_user.test_user2.id}"
}
`

	UserGroupMembershipResourceConfigNew = UserGroupMembershipResourceDependencies + `
resource "oci_identity_user_group_membership" "test_user_group_membership" {
	#Required
	group_id = "${oci_identity_group.test_group2.id}"
	user_id = "${oci_identity_user.test_user2.id}"
}
`

	UserGroupMembershipPropertyVariables = `

`

	GroupResourceConfig2 = GroupResourceDependencies + `
resource "oci_identity_group" "test_group2" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.group_description2}"
	name = "${var.group_name2}"
}
`
	GroupPropertyVariables2 = `
variable "group_description2" { default = "Group 2 for network administrators" }
variable "group_name2" { default = "NetworkAdmins2" }

`

	UserResourceConfig2 = UserResourceDependencies + `
resource "oci_identity_user" "test_user2" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.user_description2}"
	name = "${var.user_name2}"
}
`
	UserPropertyVariables2 = `
variable "user_description2" { default = "John Doe" }
variable "user_name2" { default = "JohnDoe@example.com" }

`

	UserGroupMembershipResourceDependencies = GroupPropertyVariables + GroupResourceConfig + GroupPropertyVariables2 + GroupResourceConfig2 +
		UserPropertyVariables + UserResourceConfig + UserPropertyVariables2 + UserResourceConfig2
)

func TestIdentityUserGroupMembershipResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_user_group_membership.test_user_group_membership"
	datasourceName := "data.oci_identity_user_group_memberships.test_user_group_memberships"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + UserGroupMembershipPropertyVariables + compartmentIdVariableStr + UserGroupMembershipResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `

                ` + compartmentIdVariableStr + UserGroupMembershipResourceConfigNew,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `

data "oci_identity_user_group_memberships" "test_user_group_memberships" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	group_id = "${oci_identity_group.test_group.id}"
	user_id = "${oci_identity_user.test_user.id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_user_group_membership.test_user_group_membership.id}"]
    }
}
                ` + compartmentIdVariableStr + UserGroupMembershipResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "group_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "memberships.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.group_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.user_id"),
				),
			},
		},
	})
}

func TestIdentityUserGroupMembershipResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_user_group_membership.test_user_group_membership"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + UserGroupMembershipPropertyVariables + compartmentIdVariableStr + UserGroupMembershipResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
				` + compartmentIdVariableStr + UserGroupMembershipResourceConfig2,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter GroupId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
				` + compartmentIdVariableStr + UserGroupMembershipResourceConfig3,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter UserId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
