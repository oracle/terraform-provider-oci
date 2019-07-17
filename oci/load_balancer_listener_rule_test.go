// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	listenerRuleDataSourceRepresentation = map[string]interface{}{
		"listener_name":    Representation{repType: Required, create: `${oci_load_balancer_listener.test_listener.name}`},
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
	}

	ListenerRuleResourceConfig = ListenerResourceDependencies +
		generateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Optional, Create, listenerRepresentation)
)

func TestLoadBalancerListenerRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerRuleResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_load_balancer_listener_rules.test_listener_rules"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_load_balancer_listener_rules", "test_listener_rules", Required, Create, listenerRuleDataSourceRepresentation) +
					compartmentIdVariableStr + ListenerRuleResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "listener_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "listener_rules.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "listener_rules.0.name"),
					resource.TestCheckResourceAttr(datasourceName, "listener_rules.0.rule.#", "1"),
				),
			},
		},
	})
}
