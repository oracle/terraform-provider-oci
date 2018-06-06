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

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

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
                ` + compartmentIdVariableStr + LoadBalancerPolicyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "policies.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "policies.0.name"),
				),
			},
		},
	})
}
