// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	BootVolumeResourceConfig = BootVolumeResourceDependencies + `

`
	BootVolumePropertyVariables = `
`

	BootVolumeResourceDependencies = BootVolumePropertyVariables + InstancePropertyVariables + InstanceResourceAsDependencyConfig
)

func TestCoreBootVolumeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_boot_volumes.test_boot_volumes"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + BootVolumeResourceConfig + `
					data "oci_core_boot_volumes" "test_boot_volumes" {
						#Required
						availability_domain = "${oci_core_instance.test_instance.availability_domain}"
						compartment_id = "${var.compartment_id}"

					}
                ` + compartmentIdVariableStr,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.size_in_mbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.time_created"),
				),
			},
		},
	})
}
