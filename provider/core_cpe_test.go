// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	CpeRequiredOnlyResource = CpeRequiredOnlyResourceDependencies + `
resource "oci_core_cpe" "test_cpe" {
	#Required
	compartment_id = "${var.compartment_id}"
	ip_address = "${var.cpe_ip_address}"
}
`

	CpeResourceConfig = CpeResourceDependencies + `
resource "oci_core_cpe" "test_cpe" {
	#Required
	compartment_id = "${var.compartment_id}"
	ip_address = "${var.cpe_ip_address}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.cpe_defined_tags_value}")}"
	display_name = "${var.cpe_display_name}"
	freeform_tags = "${var.cpe_freeform_tags}"
}
`
	CpePropertyVariables = `
variable "cpe_defined_tags_value" { default = "value" }
variable "cpe_display_name" { default = "MyCpe" }
variable "cpe_freeform_tags" { default = {"Department"= "Finance"} }
variable "cpe_ip_address" { default = "189.44.2.135" }

`
	CpeRequiredOnlyResourceDependencies = ``
	CpeResourceDependencies             = DefinedTagsDependencies
)

func TestCoreCpeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_cpe.test_cpe"
	datasourceName := "data.oci_core_cpes.test_cpes"

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
				Config:            config + CpePropertyVariables + compartmentIdVariableStr + CpeRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "189.44.2.135"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CpeResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + CpePropertyVariables + compartmentIdVariableStr + CpeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCpe"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "189.44.2.135"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "cpe_defined_tags_value" { default = "updatedValue" }
variable "cpe_display_name" { default = "displayName2" }
variable "cpe_freeform_tags" { default = {"Department"= "Accounting"} }
variable "cpe_ip_address" { default = "189.44.2.135" }

                ` + compartmentIdVariableStr + CpeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "189.44.2.135"),

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
variable "cpe_defined_tags_value" { default = "updatedValue" }
variable "cpe_display_name" { default = "displayName2" }
variable "cpe_freeform_tags" { default = {"Department"= "Accounting"} }
variable "cpe_ip_address" { default = "189.44.2.135" }

data "oci_core_cpes" "test_cpes" {
	#Required
	compartment_id = "${var.compartment_id}"

    filter {
    	name = "id"
    	values = ["${oci_core_cpe.test_cpe.id}"]
    }
}
                ` + compartmentIdVariableStr + CpeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "cpes.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cpes.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "cpes.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cpes.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "cpes.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "cpes.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "cpes.0.ip_address", "189.44.2.135"),
				),
			},
		},
	})
}
