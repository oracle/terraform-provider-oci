// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	listenerRuleDataSourceRepresentation = map[string]interface{}{
		"listener_name":    acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_listener.test_listener.name}`},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
	}

	ListenerRuleResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(listenerRepresentationOciCerts, map[string]interface{}{
			"rule_set_names": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Required, acctest.Create, ruleSetRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerListenerRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_load_balancer_listener_rules.test_listener_rules"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_listener_rules", "test_listener_rules", acctest.Required, acctest.Create, listenerRuleDataSourceRepresentation) +
				compartmentIdVariableStr + ListenerRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "listener_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "listener_rules.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "listener_rules.0.name"),
				resource.TestCheckResourceAttr(datasourceName, "listener_rules.0.rule.#", "1"),
			),
		},
	})
}
