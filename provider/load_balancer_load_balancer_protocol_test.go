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
resource "oci_load_balancer_load_balancer_protocol" "test_load_balancer_protocol" {
}
`
	LoadBalancerProtocolPropertyVariables = `

`
	LoadBalancerProtocolResourceDependencies = ""
)

func TestLoadBalancerLoadBalancerProtocolResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_load_balancer_load_balancer_protocol.test_load_balancer_protocol"
	datasourceName := "data.oci_load_balancer_protocols.test_load_balancer_protocols"

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
				Config:            config + LoadBalancerProtocolPropertyVariables + compartmentIdVariableStr + LoadBalancerProtocolResourceConfig,
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

                ` + compartmentIdVariableStr2 + LoadBalancerProtocolResourceConfig,
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

data "oci_load_balancer_protocols" "test_load_balancer_protocols" {
	#Required
	compartment_id = "${var.compartment_id}"

    filter {
    	name = "id"
    	values = ["${oci_load_balancer_load_balancer_protocol.test_load_balancer_protocol.id}"]
    }
}
                ` + compartmentIdVariableStr2 + LoadBalancerProtocolResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),

					resource.TestCheckResourceAttr(datasourceName, "protocols.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "protocols.0.name"),
				),
			},
		},
	})
}
