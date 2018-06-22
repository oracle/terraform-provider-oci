// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	TagNamespaceRequiredOnlyResource = TagNamespaceResourceDependencies + `
resource "oci_identity_tag_namespace" "test_tag_namespace" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.tag_namespace_description}"
	name = "${var.tag_namespace_name}"
}
`

	TagNamespaceResourceConfig = TagNamespaceResourceDependencies + `
resource "oci_identity_tag_namespace" "test_tag_namespace" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.tag_namespace_description}"
	name = "${var.tag_namespace_name}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.tag_namespace_defined_tags_value}")}"
	freeform_tags = "${var.tag_namespace_freeform_tags}"
}
`
	TagNamespacePropertyVariables = `
variable "tag_namespace_defined_tags_value" { default = "value" }
variable "tag_namespace_description" { default = "This namespace contains tags that will be used in billing." }
variable "tag_namespace_freeform_tags" { default = {"Department"= "Finance"} }
variable "tag_namespace_include_subcompartments" { default = false }
variable "tag_namespace_name" { default = "BillingTags" }

`
	TagNamespaceResourceDependencies = DefinedTagsDependencies
)

func TestIdentityTagNamespaceResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_tag_namespace.test_tag_namespace"
	datasourceName := "data.oci_identity_tag_namespaces.test_tag_namespaces"

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
				Config:            config + TagNamespacePropertyVariables + compartmentIdVariableStr + TagNamespaceRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "This namespace contains tags that will be used in billing."),
					resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + TagNamespaceResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + TagNamespacePropertyVariables + compartmentIdVariableStr + TagNamespaceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "This namespace contains tags that will be used in billing."),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_retired"),
					resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),
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
variable "tag_namespace_defined_tags_value" { default = "updatedValue" }
variable "tag_namespace_description" { default = "description2" }
variable "tag_namespace_freeform_tags" { default = {"Department"= "Accounting"} }
variable "tag_namespace_include_subcompartments" { default = false }
variable "tag_namespace_name" { default = "BillingTags" }

                ` + compartmentIdVariableStr + TagNamespaceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_retired"),
					resource.TestCheckResourceAttr(resourceName, "name", "BillingTags"),
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
variable "tag_namespace_defined_tags_value" { default = "updatedValue" }
variable "tag_namespace_description" { default = "description2" }
variable "tag_namespace_freeform_tags" { default = {"Department"= "Accounting"} }
variable "tag_namespace_include_subcompartments" { default = false }
variable "tag_namespace_name" { default = "BillingTags" }

data "oci_identity_tag_namespaces" "test_tag_namespaces" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	include_subcompartments = "${var.tag_namespace_include_subcompartments}"

    filter {
    	name = "id"
    	values = ["${oci_identity_tag_namespace.test_tag_namespace.id}"]
    }
}
                ` + compartmentIdVariableStr + TagNamespaceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "include_subcompartments", "false"),

					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "tag_namespaces.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "tag_namespaces.0.is_retired"),
					resource.TestCheckResourceAttr(datasourceName, "tag_namespaces.0.name", "BillingTags"),
					resource.TestCheckResourceAttrSet(datasourceName, "tag_namespaces.0.time_created"),
				),
			},
		},
	})
}
