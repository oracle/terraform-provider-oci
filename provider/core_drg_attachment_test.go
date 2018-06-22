// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	DrgAttachmentRequiredOnlyResource = DrgAttachmentResourceDependencies + `
resource "oci_core_drg_attachment" "test_drg_attachment" {
	#Required
	drg_id = "${oci_core_drg.test_drg.id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`

	DrgAttachmentResourceConfig = DrgAttachmentResourceDependencies + `
resource "oci_core_drg_attachment" "test_drg_attachment" {
	#Required
	drg_id = "${oci_core_drg.test_drg.id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.drg_attachment_display_name}"
}
`
	DrgAttachmentPropertyVariables = `
variable "drg_attachment_display_name" { default = "displayName" }

`
	DrgAttachmentResourceDependencies = DefinedTagsDependencies + DrgPropertyVariables + DrgRequiredOnlyResource + VcnPropertyVariables + VcnRequiredOnlyResource
)

func TestCoreDrgAttachmentResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachment.test_drg_attachment"
	datasourceName := "data.oci_core_drg_attachments.test_drg_attachments"

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
				Config:            config + DrgAttachmentPropertyVariables + compartmentIdVariableStr + DrgAttachmentRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + DrgAttachmentPropertyVariables + compartmentIdVariableStr + DrgAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "drg_attachment_display_name" { default = "displayName2" }

                ` + compartmentIdVariableStr + DrgAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
variable "drg_attachment_display_name" { default = "displayName2" }

data "oci_core_drg_attachments" "test_drg_attachments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	drg_id = "${oci_core_drg.test_drg.id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_drg_attachment.test_drg_attachment.id}"]
    }
}
                ` + compartmentIdVariableStr + DrgAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "drg_attachments.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.drg_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.vcn_id"),
				),
			},
		},
	})
}
