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
	"github.com/oracle/oci-go-sdk/v31/common"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v31/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ListenerRequiredOnlyResource = ListenerResourceDependencies +
		generateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Required, Create, listenerRepresentation)

	listenerRepresentation = map[string]interface{}{
		"default_backend_set_name": Representation{repType: Required, create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id":         Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                     Representation{repType: Required, create: `mylistener`},
		"port":                     Representation{repType: Required, create: `10`, update: `11`},
		"protocol":                 Representation{repType: Required, create: `HTTP`},
		"connection_configuration": RepresentationGroup{Optional, listenerConnectionConfigurationRepresentation},
		"hostname_names":           Representation{repType: Optional, create: []string{`${oci_load_balancer_hostname.test_hostname.name}`}},
		"path_route_set_name":      Representation{repType: Optional, create: `${oci_load_balancer_path_route_set.test_path_route_set.name}`},
		"rule_set_names":           Representation{repType: Optional, create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
		"ssl_configuration":        RepresentationGroup{Optional, listenerSslConfigurationRepresentation},
	}
	listenerConnectionConfigurationRepresentation = map[string]interface{}{
		"idle_timeout_in_seconds": Representation{repType: Required, create: `10`, update: `11`},
	}
	listenerSslConfigurationRepresentation = map[string]interface{}{
		"certificate_name":        Representation{repType: Required, create: `${oci_load_balancer_certificate.test_certificate.certificate_name}`},
		"verify_depth":            Representation{repType: Optional, create: `10`, update: `11`},
		"verify_peer_certificate": Representation{repType: Optional, create: `false`, update: `true`},
	}

	ListenerResourceDependencies = generateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Optional, Create, certificateRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies +
		generateResourceFromRepresentationMap("oci_load_balancer_path_route_set", "test_path_route_set", Required, Create, pathRouteSetRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_hostname", "test_hostname", Required, Create, hostnameRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Required, Create, ruleSetRepresentation) +
		`
	resource "oci_load_balancer_hostname" "test_hostname2" {
		#Required
		hostname = "app.example.com2"
		load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
		name = "example_hostname_0012"
	}
`
)

func TestLoadBalancerListenerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerListenerDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Required, Create, listenerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "mylistener"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Optional, Create,
						getUpdatedRepresentationCopy("hostname_names", Representation{repType: Optional, create: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"}, update: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"}}, listenerRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "rule_set_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "ssl_configuration.0.certificate_name"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "10"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", Optional, Update, listenerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "rule_set_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "ssl_configuration.0.certificate_name"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "11"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "true"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
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
