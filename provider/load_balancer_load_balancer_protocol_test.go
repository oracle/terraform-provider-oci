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

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

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
                ` + compartmentIdVariableStr + LoadBalancerProtocolResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "protocols.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "protocols.0.name"),
				),
			},
		},
	})
}
