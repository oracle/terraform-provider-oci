// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v49/common"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v49/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ListenerRequiredOnlyResource = ListenerResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Required, Create, listenerRepresentation)

	listenerRepresentation = map[string]interface{}{
		"default_backend_set_name": Representation{RepType: Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id":         Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                     Representation{RepType: Required, Create: `mylistener`},
		"port":                     Representation{RepType: Required, Create: `10`, Update: `11`},
		"protocol":                 Representation{RepType: Required, Create: `HTTP`},
		"connection_configuration": RepresentationGroup{Optional, listenerConnectionConfigurationRepresentation},
		"hostname_names":           Representation{RepType: Optional, Create: []string{`${oci_load_balancer_hostname.test_hostname.name}`}},
		"path_route_set_name":      Representation{RepType: Optional, Create: `${oci_load_balancer_path_route_set.test_path_route_set.name}`},
		"routing_policy_name":      Representation{RepType: Optional, Create: `${oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy.name}`},
		"rule_set_names":           Representation{RepType: Optional, Create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
		"ssl_configuration":        RepresentationGroup{Optional, listenerSslConfigurationRepresentation},
	}
	listenerConnectionConfigurationRepresentation = map[string]interface{}{
		"idle_timeout_in_seconds": Representation{RepType: Required, Create: `10`, Update: `11`},
	}
	listenerSslConfigurationRepresentation = map[string]interface{}{
		"certificate_name":        Representation{RepType: Required, Create: `${oci_load_balancer_certificate.test_certificate.certificate_name}`},
		"verify_depth":            Representation{RepType: Optional, Create: `10`, Update: `11`},
		"verify_peer_certificate": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}

	ListenerResourceDependencies = GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer_routing_policy", "test_load_balancer_routing_policy", Required, Create, loadBalancerRoutingPolicyRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Optional, Create, certificateRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies +
		GenerateResourceFromRepresentationMap("oci_load_balancer_path_route_set", "test_path_route_set", Required, Create, pathRouteSetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_hostname", "test_hostname", Required, Create, hostnameRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Required, Create, ruleSetRepresentation) +
		caCertificateVariableStr + privateKeyVariableStr +
		`
	resource "oci_load_balancer_hostname" "test_hostname2" {
		#Required
		hostname = "app.example.com2"
		load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
		name = "example_hostname_0012"
	}
`
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerListenerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ListenerResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Optional, Create, listenerRepresentation), "loadbalancer", "listener", t)

	ResourceTest(t, testAccCheckLoadBalancerListenerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Required, Create, listenerRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "mylistener"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ListenerResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Optional, Create,
					GetUpdatedRepresentationCopy("hostname_names", Representation{RepType: Optional, Create: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"}, Update: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"}}, listenerRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "connection_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.idle_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", "backendSet1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_names.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "hostname_names.0", "example_hostname_001"),
				resource.TestCheckResourceAttr(resourceName, "hostname_names.1", "example_hostname_0012"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "mylistener"),
				resource.TestCheckResourceAttrSet(resourceName, "path_route_set_name"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttrSet(resourceName, "routing_policy_name"),
				resource.TestCheckResourceAttr(resourceName, "rule_set_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "ssl_configuration.0.certificate_name"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "10"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Optional, Update, listenerRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "connection_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.idle_timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", "backendSet1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_names.0", "example_hostname_001"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "mylistener"),
				resource.TestCheckResourceAttrSet(resourceName, "path_route_set_name"),
				resource.TestCheckResourceAttr(resourceName, "port", "11"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttrSet(resourceName, "routing_policy_name"),
				resource.TestCheckResourceAttr(resourceName, "rule_set_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "ssl_configuration.0.certificate_name"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "11"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "true"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

func testAccCheckLoadBalancerListenerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_listener" {
			noResourceFound = false
			request := oci_load_balancer.GetLoadBalancerRequest{}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			response, err := client.GetLoadBalancer(context.Background(), request)
			if err == nil {
				lb := &response.LoadBalancer
				listenerName := rs.Primary.Attributes["name"]
				if lb != nil && lb.Listeners != nil {
					if l, ok := lb.Listeners[listenerName]; ok {
						if l.Name != nil && *l.Name == listenerName {
							return fmt.Errorf("listener still exists")
						}
					}
				}
				// no error and item not found, item is deleted
				return nil
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}

		}
	}

	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
