// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	CompartmentResourceConfig = CompartmentResourceDependencies + `
resource "oci_identity_compartment" "test_compartment" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.compartment_description}"
	name = "${var.compartment_name}"
}
`
	CompartmentPropertyVariables = `
variable "compartment_description" { default = "For network components" }
variable "compartment_name" { default = "Network" }

`
	CompartmentResourceDependencies = ""
)

func TestIdentityCompartmentResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_compartment.test_compartment"
	datasourceName := "data.oci_identity_compartments.test_compartments"

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
				Config:            config + CompartmentPropertyVariables + compartmentIdVariableStr + CompartmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "For network components"),
					resource.TestCheckResourceAttr(resourceName, "name", "Network"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "compartment_description" { default = "description2" }
variable "compartment_name" { default = "name2" }

                ` + compartmentIdVariableStr + CompartmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
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
variable "compartment_description" { default = "description2" }
variable "compartment_name" { default = "name2" }

data "oci_identity_compartments" "test_compartments" {
	#Required
	compartment_id = "${var.compartment_id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_compartment.test_compartment.id}"]
    }
}
                ` + compartmentIdVariableStr + CompartmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "compartments.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "compartments.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "compartments.0.description", "description2"),
					resource.TestCheckResourceAttrSet(datasourceName, "compartments.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "compartments.0.name", "name2"),
					resource.TestCheckResourceAttrSet(datasourceName, "compartments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "compartments.0.time_created"),
				),
			},
			// verify reset to name and description
			{
				Config: config + CompartmentPropertyVariables + compartmentIdVariableStr + CompartmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "For network components"),
					resource.TestCheckResourceAttr(resourceName, "name", "Network"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}
