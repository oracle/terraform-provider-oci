// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	VnicResourceConfig = VnicResourceDependencies + `
resource "oci_core_vnic" "test_vnic" {
}
`
	VnicPropertyVariables = `
variable "vnic_vnic_id" { default = "vnicId" }

`
	VnicResourceDependencies = ""
)

func TestCoreVnicResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_vnic.test_vnic"
	datasourceName := "data.oci_core_vnics.test_vnics"

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
				Config:            config + VnicPropertyVariables + compartmentIdVariableStr + VnicResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "vnic_vnic_id" { default = "vnicId" }

                ` + compartmentIdVariableStr + VnicResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
variable "vnic_vnic_id" { default = "vnicId" }

data "oci_core_vnics" "test_vnics" {
	#Required
	vnic_id = "${var.vnic_vnic_id}"

    filter {
    	name = "id"
    	values = ["${oci_core_vnic.test_vnic.id}"]
    }
}
                ` + compartmentIdVariableStr + VnicResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "vnic_id", "vnicId"),

					resource.TestCheckResourceAttr(datasourceName, "vnic.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic.0.private_ip_address"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic.0.time_created"),
				),
			},
		},
	})
}
