// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	LoadBalancerPolicyResourceConfig = LoadBalancerPolicyResourceDependencies + `

`
	LoadBalancerPolicyPropertyVariables = `

`
	LoadBalancerPolicyResourceDependencies = ""
)

func TestLoadBalancerLoadBalancerPolicyResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	datasourceName := "data.oci_load_balancer_policies.test_load_balancer_policies"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_load_balancer_policies" "test_load_balancer_policies" {
	#Required
	compartment_id = "${var.compartment_id}"
}
                ` + compartmentIdVariableStr2 + LoadBalancerPolicyResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),

					resource.TestCheckResourceAttrSet(datasourceName, "policies.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "policies.0.name"),
				),
			},
		},
	})
}
