// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	GroupResourceConfig = GroupResourceDependencies + `
resource "oci_identity_group" "test_group" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.group_description}"
	name = "${var.group_name}"
}
`
	GroupPropertyVariables = `
variable "group_description" { default = "Group for network administrators" }
variable "group_name" { default = "NetworkAdmins" }

`
	GroupResourceDependencies = ""
)

func TestIdentityGroupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_group.test_group"
	datasourceName := "data.oci_identity_groups.test_groups"

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
				Config:            config + GroupPropertyVariables + compartmentIdVariableStr + GroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "Group for network administrators"),
					resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "group_description" { default = "description2" }
variable "group_name" { default = "NetworkAdmins" }

                ` + compartmentIdVariableStr + GroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "group_description" { default = "description2" }
variable "group_name" { default = "NetworkAdmins" }

data "oci_identity_groups" "test_groups" {
	#Required
	compartment_id = "${var.compartment_id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_group.test_group.id}"]
    }
}
                ` + compartmentIdVariableStr + GroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "groups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.description", "description2"),
					resource.TestCheckResourceAttrSet(datasourceName, "groups.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.name", "NetworkAdmins"),
					resource.TestCheckResourceAttrSet(datasourceName, "groups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "groups.0.time_created"),
				),
			},
		},
	})
}
