// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	DrgRequiredOnlyResource = DrgResourceDependencies + `
resource "oci_core_drg" "test_drg" {
	#Required
	compartment_id = "${var.compartment_id}"
}
`

	DrgResourceConfig = DrgResourceDependencies + `
resource "oci_core_drg" "test_drg" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.drg_display_name}"
}
`
	DrgPropertyVariables = `
variable "drg_display_name" { default = "MyDrg" }

`
	DrgResourceDependencies = ""
)

func TestCoreDrgResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg.test_drg"
	datasourceName := "data.oci_core_drgs.test_drgs"

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
				Config:            config + DrgPropertyVariables + compartmentIdVariableStr + DrgRequiredOnlyResource,
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
				Config: config + compartmentIdVariableStr + DrgResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + DrgPropertyVariables + compartmentIdVariableStr + DrgResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyDrg"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "drg_display_name" { default = "displayName2" }

                ` + compartmentIdVariableStr + DrgResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
variable "drg_display_name" { default = "displayName2" }

data "oci_core_drgs" "test_drgs" {
	#Required
	compartment_id = "${var.compartment_id}"

    filter {
    	name = "id"
    	values = ["${oci_core_drg.test_drg.id}"]
    }
}
                ` + compartmentIdVariableStr + DrgResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "drgs.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "drgs.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "drgs.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "drgs.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drgs.0.state"),
				),
			},
		},
	})
}
