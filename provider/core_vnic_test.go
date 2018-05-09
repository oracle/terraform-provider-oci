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

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

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

			// verify updates to Force New parameters.
			{
				Config: config + `
variable "vnic_vnic_id" { default = "vnicId2" }

                ` + compartmentIdVariableStr2 + VnicResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
variable "vnic_vnic_id" { default = "vnicId2" }

data "oci_core_vnics" "test_vnics" {
	#Required
	vnic_id = "${var.vnic_vnic_id}"

    filter {
    	name = "id"
    	values = ["${oci_core_vnic.test_vnic.id}"]
    }
}
                ` + compartmentIdVariableStr2 + VnicResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "vnic_id", "vnicId2"),

					resource.TestCheckResourceAttr(datasourceName, "vnics.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnics.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnics.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnics.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnics.0.private_ip"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnics.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnics.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnics.0.time_created"),
				),
			},
		},
	})
}

func TestCoreVnicResource_forcenew(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_vnic.test_vnic"

	var resId string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + VnicPropertyVariables + compartmentIdVariableStr + VnicResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
