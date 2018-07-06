// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	CrossConnectGroupRequiredOnlyResource = CrossConnectGroupResourceDependencies + `
resource "oci_core_cross_connect_group" "test_cross_connect_group" {
	#Required
	compartment_id = "${var.compartment_id}"
}
`

	CrossConnectGroupResourceConfig = CrossConnectGroupResourceDependencies + `
resource "oci_core_cross_connect_group" "test_cross_connect_group" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.cross_connect_group_display_name}"
}
`
	CrossConnectGroupPropertyVariables = `
variable "cross_connect_group_display_name" { default = "displayName" }
variable "cross_connect_group_state" { default = "AVAILABLE" }

`
	CrossConnectGroupResourceDependencies = ""
)

func TestCoreCrossConnectGroupResource_basic(t *testing.T) {
	region := getRequiredEnvSetting("region")
	if !strings.EqualFold("r1", region) {
		t.Skip("FastConnect tests are not yet supported in production regions")
	}

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_cross_connect_group.test_cross_connect_group"
	datasourceName := "data.oci_core_cross_connect_groups.test_cross_connect_groups"
	singularDatasourceName := "data.oci_core_cross_connect_group.test_cross_connect_group"

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
				Config:            config + CrossConnectGroupPropertyVariables + compartmentIdVariableStr + CrossConnectGroupRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CrossConnectGroupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + CrossConnectGroupPropertyVariables + compartmentIdVariableStr + CrossConnectGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "cross_connect_group_display_name" { default = "displayName2" }
variable "cross_connect_group_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + CrossConnectGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),

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
variable "cross_connect_group_display_name" { default = "displayName2" }
variable "cross_connect_group_state" { default = "AVAILABLE" }

data "oci_core_cross_connect_groups" "test_cross_connect_groups" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.cross_connect_group_display_name}"
	#state = "${var.cross_connect_group_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_cross_connect_group.test_cross_connect_group.id}"]
    }
}
                ` + compartmentIdVariableStr + CrossConnectGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					//resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.display_name", "displayName2"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
variable "cross_connect_group_display_name" { default = "displayName2" }
variable "cross_connect_group_state" { default = "AVAILABLE" }

data "oci_core_cross_connect_group" "test_cross_connect_group" {
	#Required
	cross_connect_group_id = "${oci_core_cross_connect_group.test_cross_connect_group.id}"
}
                ` + compartmentIdVariableStr + CrossConnectGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_group_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					//resource.TestCheckResourceAttr(singularDatasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}
