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
		generateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_tcp", Required, Create, listenerTcpRepresentation)
	listenerTcpRepresentation = map[string]interface{}{
		"default_backend_set_name": Representation{repType: Required, create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id":         Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                     Representation{repType: Required, create: `mylistener`},
		"port":                     Representation{repType: Required, create: `10`, update: `11`},
		"protocol":                 Representation{repType: Required, create: `TCP`},
		"connection_configuration": RepresentationGroup{Optional, listenerTcpConnectionConfigurationRepresentation},
		"rule_set_names":           Representation{repType: Optional, create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
	}
	listenerTcpConnectionConfigurationRepresentation = map[string]interface{}{
		"idle_timeout_in_seconds":            Representation{repType: Required, create: `10`, update: `11`},
		"backend_tcp_proxy_protocol_version": Representation{repType: Optional, create: `1`, update: `2`},
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
			// verify create with TCP optional
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_tcp", Optional, Create, listenerTcpRepresentation),
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

			// verify update with TCP optional
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_tcp", Optional, Update, listenerTcpRepresentation),
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
