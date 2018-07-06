// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	CrossConnectLocationResourceConfig = CrossConnectLocationResourceDependencies + `

`
	CrossConnectLocationPropertyVariables = `

`
	CrossConnectLocationResourceDependencies = ""
)

func TestCoreCrossConnectLocationResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_cross_connect_locations.test_cross_connect_locations"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_core_cross_connect_locations" "test_cross_connect_locations" {
	#Required
	compartment_id = "${var.compartment_id}"
}
                ` + compartmentIdVariableStr + CrossConnectLocationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_locations.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_locations.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_locations.0.name"),
				),
			},
		},
	})
}
