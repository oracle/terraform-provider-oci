// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	TagRequiredOnlyResource = TagResourceDependencies + `
resource "oci_identity_tag" "test_tag" {
	#Required
	description = "${var.tag_description}"
	name = "${var.tag_name}"
	tag_namespace_id = "${oci_identity_tag_namespace.test_tag_namespace.id}"
}
`

	TagResourceConfig = TagResourceDependencies + `
resource "oci_identity_tag" "test_tag" {
    depends_on = ["oci_identity_tag_namespace.tag-namespace1", "oci_identity_tag.tag1"]
	#Required
	description = "${var.tag_description}"
	name = "${var.tag_name}"
	tag_namespace_id = "${oci_identity_tag_namespace.test_tag_namespace.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.tag_defined_tags_value}")}"
	freeform_tags = "${var.tag_freeform_tags}"
}
`
	TagPropertyVariables = `
variable "tag_defined_tags_value" { default = "value" }
variable "tag_description" { default = "This tag will show the cost center that will be used for billing of associated resources." }
variable "tag_freeform_tags" { default = {"Department"= "Finance"} }
variable "tag_name" { default = "CostCenter" }

`
	TagResourceDependencies = TagNamespacePropertyVariables + TagNamespaceRequiredOnlyResource

	DefinedTagsDependencies = `
variable defined_tag_namespace_name { default = "" }
resource "oci_identity_tag_namespace" "tag-namespace1" {
  		#Required
		compartment_id = "${var.tenancy_ocid}"
  		description = "example tag namespace"
  		name = "${var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"}"

		is_retired = false
}

resource "oci_identity_tag" "tag1" {
  		#Required
  		description = "example tag"
  		name = "example-tag"
        tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"

		is_retired = false
}
`
)

func TestIdentityTagResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_tag.test_tag"
	datasourceName := "data.oci_identity_tags.test_tags"

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
				Config:            config + TagPropertyVariables + compartmentIdVariableStr + TagRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "This tag will show the cost center that will be used for billing of associated resources."),
					resource.TestCheckResourceAttr(resourceName, "name", "CostCenter"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + TagResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + TagPropertyVariables + compartmentIdVariableStr + TagResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "This tag will show the cost center that will be used for billing of associated resources."),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_retired"),
					resource.TestCheckResourceAttr(resourceName, "name", "CostCenter"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "tag_defined_tags_value" { default = "updatedValue" }
variable "tag_description" { default = "description2" }
variable "tag_freeform_tags" { default = {"Department"= "Accounting"} }
variable "tag_name" { default = "CostCenter" }

                ` + compartmentIdVariableStr + TagResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_retired"),
					resource.TestCheckResourceAttr(resourceName, "name", "CostCenter"),
					resource.TestCheckResourceAttrSet(resourceName, "tag_namespace_id"),
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
variable "tag_defined_tags_value" { default = "updatedValue" }
variable "tag_description" { default = "description2" }
variable "tag_freeform_tags" { default = {"Department"= "Accounting"} }
variable "tag_name" { default = "CostCenter" }

data "oci_identity_tags" "test_tags" {
	#Required
	tag_namespace_id = "${oci_identity_tag_namespace.test_tag_namespace.id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_tag.test_tag.id}"]
    }
}
                ` + compartmentIdVariableStr + TagResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "tag_namespace_id"),

					resource.TestCheckResourceAttr(datasourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.is_retired"),
					resource.TestCheckResourceAttr(datasourceName, "tags.0.name", "CostCenter"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.time_created"),
				),
			},
		},
	})
}
