// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ListenerTcpRequiredOnlyResource = ListenerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_tcp", acctest.Required, acctest.Create, listenerTcpRepresentation)
	listenerTcpRepresentation = map[string]interface{}{
		"default_backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `mylistener`},
		"port":                     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"protocol":                 acctest.Representation{RepType: acctest.Required, Create: `TCP`},
		"connection_configuration": acctest.RepresentationGroup{RepType: acctest.Optional, Group: listenerTcpConnectionConfigurationRepresentation},
		"rule_set_names":           acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
	}
	listenerTcpConnectionConfigurationRepresentation = map[string]interface{}{
		"idle_timeout_in_seconds":            acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"backend_tcp_proxy_protocol_version": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerListenerTcpResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerTcpResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener_tcp"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerListenerDestroy,
		Steps: []resource.TestStep{
			// verify Create with TCP optional
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_tcp", acctest.Optional, acctest.Create, listenerTcpRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.backend_tcp_proxy_protocol_version", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.idle_timeout_in_seconds", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "mylistener"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "rule_set_names.#", "1"),
				),
			},

			// verify Update with TCP optional
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_tcp", acctest.Optional, acctest.Update, listenerTcpRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.backend_tcp_proxy_protocol_version", "2"),
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.idle_timeout_in_seconds", "11"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "mylistener"),
					resource.TestCheckResourceAttr(resourceName, "port", "11"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "rule_set_names.#", "1"),
				),
			},
		},
	})
}
