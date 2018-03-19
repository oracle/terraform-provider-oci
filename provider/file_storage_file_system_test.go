// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	FileSystemRequiredOnlyResource = FileSystemResourceDependencies + `
resource "oci_file_storage_file_system" "test_file_system" {
	#Required
	availability_domain = "${var.file_system_availability_domain}"
	compartment_id = "${var.compartment_id}"
}
`

	FileSystemResourceConfig = FileSystemResourceDependencies + `
resource "oci_file_storage_file_system" "test_file_system" {
	#Required
	availability_domain = "${var.file_system_availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.file_system_display_name}"
}
`
	FileSystemPropertyVariables = `
variable "file_system_availability_domain" { default = "kIdk:PHX-AD-1" }
variable "file_system_display_name" { default = "media-files-1" }

`
	FileSystemResourceDependencies = ""
)

func TestFileStorageFileSystemResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_file_storage_file_system.test_file_system"
	datasourceName := "data.oci_file_storage_file_systems.test_file_systems"

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
				Config:            config + FileSystemPropertyVariables + compartmentIdVariableStr + FileSystemRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "kIdk:PHX-AD-1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + FileSystemResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + FileSystemPropertyVariables + compartmentIdVariableStr + FileSystemResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "kIdk:PHX-AD-1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "media-files-1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
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
variable "file_system_availability_domain" { default = "kIdk:PHX-AD-1" }
variable "file_system_display_name" { default = "displayName2" }

                ` + compartmentIdVariableStr + FileSystemResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "kIdk:PHX-AD-1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
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
			// verify updates to Force New parameters.
			{
				Config: config + `
variable "file_system_availability_domain" { default = "kIdk:PHX-AD-2" }
variable "file_system_display_name" { default = "displayName2" }

                ` + compartmentIdVariableStr2 + FileSystemResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "availability_domain", "kIdk:PHX-AD-2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "metered_bytes"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
variable "file_system_availability_domain" { default = "kIdk:PHX-AD-2" }
variable "file_system_display_name" { default = "displayName2" }

data "oci_file_storage_file_systems" "test_file_systems" {
	#Required
	availability_domain = "${var.file_system_availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.file_system_display_name}"
	id = "${oci_file_storage_file_system.test_file_system.id}"
	state = "${oci_file_storage_file_system.test_file_system.state}"

    filter {
    	name = "id"
    	values = ["${oci_file_storage_file_system.test_file_system.id}"]
    }
}
                ` + compartmentIdVariableStr2 + FileSystemResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "kIdk:PHX-AD-2"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					TestCheckResourceAttributesEqual(datasourceName, "state", "oci_file_storage_file_system.test_file_system", "state"),

					resource.TestCheckResourceAttr(datasourceName, "file_systems.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "file_systems.0.availability_domain", "kIdk:PHX-AD-2"),
					resource.TestCheckResourceAttr(datasourceName, "file_systems.0.compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "file_systems.0.display_name", "displayName2"),
					TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.id", "oci_file_storage_file_system.test_file_system", "id"),
					TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.metered_bytes", "oci_file_storage_file_system.test_file_system", "metered_bytes"),
					TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.state", "oci_file_storage_file_system.test_file_system", "state"),
					TestCheckResourceAttributesEqual(datasourceName, "file_systems.0.time_created", "oci_file_storage_file_system.test_file_system", "time_created"),
				),
			},
		},
	})
}
