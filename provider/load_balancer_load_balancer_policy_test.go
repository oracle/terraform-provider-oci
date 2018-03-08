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
resource "oci_load_balancer_load_balancer_policy" "test_load_balancer_policy" {
}
`
	LoadBalancerPolicyPropertyVariables = `

`
	LoadBalancerPolicyResourceDependencies = ""
)

func TestLoadBalancerLoadBalancerPolicyResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_load_balancer_load_balancer_policy.test_load_balancer_policy"
	datasourceName := "data.oci_load_balancer_policies.test_load_balancer_policies"

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
				Config:            config + LoadBalancerPolicyPropertyVariables + compartmentIdVariableStr + LoadBalancerPolicyResourceConfig,
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

                ` + compartmentIdVariableStr2 + LoadBalancerPolicyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "name"),

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

data "oci_load_balancer_policies" "test_load_balancer_policies" {
	#Required
	compartment_id = "${var.compartment_id}"

    filter {
    	name = "id"
    	values = ["${oci_load_balancer_load_balancer_policy.test_load_balancer_policy.id}"]
    }
}
                ` + compartmentIdVariableStr2 + LoadBalancerPolicyResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),

					resource.TestCheckResourceAttr(datasourceName, "policies.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "policies.0.name"),
				),
			},
		},
	})
}
