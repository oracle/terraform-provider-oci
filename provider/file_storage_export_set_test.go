// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ExportSetResourceConfig = ExportSetResourceDependencies + `
resource "oci_file_storage_export_set" "test_export_set" {
}
`
	ExportSetPropertyVariables = `
variable "export_set_availability_domain" { default = "availabilityDomain" }
variable "export_set_display_name" { default = "displayName" }
variable "export_set_id" { default = "id" }
variable "export_set_state" { default = "state" }

`
	ExportSetResourceDependencies = ""
)

func TestFileStorageExportSetResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_file_storage_export_set.test_export_set"
	datasourceName := "data.oci_file_storage_export_sets.test_export_sets"

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
				Config:            config + ExportSetPropertyVariables + compartmentIdVariableStr + ExportSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `
variable "export_set_availability_domain" { default = "availabilityDomain2" }
variable "export_set_display_name" { default = "displayName2" }
variable "export_set_id" { default = "id2" }
variable "export_set_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr2 + ExportSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
variable "export_set_availability_domain" { default = "availabilityDomain2" }
variable "export_set_display_name" { default = "displayName2" }
variable "export_set_id" { default = "id2" }
variable "export_set_state" { default = "AVAILABLE" }

data "oci_file_storage_export_sets" "test_export_sets" {
	#Required
	availability_domain = "${var.export_set_availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.export_set_display_name}"
	id = "${var.export_set_id}"
	state = "${var.export_set_state}"

    filter {
    	name = "id"
    	values = ["${oci_file_storage_export_set.test_export_set.id}"]
    }
}
                ` + compartmentIdVariableStr2 + ExportSetResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "id", "id2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "export_sets.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.vcn_id"),
				),
			},
		},
	})
}

func TestFileStorageExportSetResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	//compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	//compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_file_storage_export_set.test_export_set"

	var resId string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + ExportSetPropertyVariables + compartmentIdVariableStr + ExportSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

		},
	})
}
