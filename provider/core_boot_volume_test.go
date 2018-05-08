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
variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
	  }
	}
`
	InstanceResourceConfigs = `
resource "oci_core_instance" "test_instance" {
	availability_domain = "${var.subnet_availability_domain}"
	compartment_id = "${var.compartment_id}"
	display_name = "-tf-instance"
	image = "${var.InstanceImageOCID[var.region]}"
	shape = "VM.Standard1.1"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
	metadata {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}

	timeouts {
		create = "15m"
	}
}`
	BootVolumeResourceDependencies = BootVolumePropertyVariables + SubnetPropertyVariables + SubnetRequiredOnlyResource + InstanceResourceConfigs
)

func TestCoreBootVolumeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)
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
						availability_domain = "${var.subnet_availability_domain}"
						compartment_id = "${var.compartment_id}"

					}
                ` + compartmentIdVariableStr2 ,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),

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
