// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	LoadBalancerProtocolResourceConfig = LoadBalancerProtocolResourceDependencies + `

`
	LoadBalancerProtocolPropertyVariables = `

`
	LoadBalancerProtocolResourceDependencies = ""
)

func TestLoadBalancerLoadBalancerProtocolResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	datasourceName := "data.oci_load_balancer_protocols.test_load_balancer_protocols"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_load_balancer_protocols" "test_load_balancer_protocols" {
	#Required
	compartment_id = "${var.compartment_id}"
}
                ` + compartmentIdVariableStr2 + LoadBalancerProtocolResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),

					resource.TestCheckResourceAttrSet(datasourceName, "protocols.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "protocols.0.name"),
				),
			},
		},
	})
}
