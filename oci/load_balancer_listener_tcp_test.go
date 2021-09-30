// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ListenerTcpRequiredOnlyResource = ListenerResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_tcp", Required, Create, listenerTcpRepresentation)
	listenerTcpRepresentation = map[string]interface{}{
		"default_backend_set_name": Representation{RepType: Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id":         Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                     Representation{RepType: Required, Create: `mylistener`},
		"port":                     Representation{RepType: Required, Create: `10`, Update: `11`},
		"protocol":                 Representation{RepType: Required, Create: `TCP`},
		"connection_configuration": RepresentationGroup{Optional, listenerTcpConnectionConfigurationRepresentation},
		"rule_set_names":           Representation{RepType: Optional, Create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
	}
	listenerTcpConnectionConfigurationRepresentation = map[string]interface{}{
		"idle_timeout_in_seconds":            Representation{RepType: Required, Create: `10`, Update: `11`},
		"backend_tcp_proxy_protocol_version": Representation{RepType: Optional, Create: `1`, Update: `2`},
	}
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerListenerTcpResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerTcpResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener_tcp"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerListenerDestroy,
		Steps: []resource.TestStep{
			// verify Create with TCP optional
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_tcp", Optional, Create, listenerTcpRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_tcp", Optional, Update, listenerTcpRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
