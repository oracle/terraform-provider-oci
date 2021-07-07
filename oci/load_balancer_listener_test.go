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
	ListenerRequiredOnlyResourceCerts = ListenerResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_ca",
			Required, Create, listenerRepresentationOciCerts)

	listenerRepresentationOciCerts = map[string]interface{}{
		"default_backend_set_name": Representation{RepType: Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id":         Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                     Representation{RepType: Required, Create: `myListener1`},
		"port":                     Representation{RepType: Required, Create: `10`, Update: `11`},
		"protocol":                 Representation{RepType: Required, Create: `HTTP`},
		"connection_configuration": RepresentationGroup{Optional, listenerConnectionConfigurationRepresentation},
		"hostname_names":           Representation{RepType: Optional, Create: []string{`${oci_load_balancer_hostname.test_hostname.name}`}},
		"path_route_set_name":      Representation{RepType: Optional, Create: `${oci_load_balancer_path_route_set.test_path_route_set.name}`},
		"routing_policy_name":      Representation{RepType: Optional, Create: `${oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy.name}`},
		"rule_set_names":           Representation{RepType: Optional, Create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
		"ssl_configuration":        RepresentationGroup{Optional, listenerSslConfigurationRepresentationOciCerts},
	}
	listenerConnectionConfigurationRepresentation = map[string]interface{}{
		"idle_timeout_in_seconds": Representation{RepType: Required, Create: `10`, Update: `11`},
	}
	listenerSslConfigurationRepresentationOciCerts = map[string]interface{}{
		// note: cannot specify certificate_name along with trusted_certificate_authority_ids
		"certificate_ids":                   Representation{RepType: Optional, Create: []string{certificateIds}, Update: []string{certificateIds2}},
		"cipher_suite_name":                 Representation{RepType: Optional, Create: `oci-default-ssl-cipher-suite-v1`, Update: `oci-default-ssl-cipher-suite-v1`},
		"protocols":                         Representation{RepType: Optional, Create: []string{`TLSv1.2`}, Update: []string{`TLSv1.2`}},
		"server_order_preference":           Representation{RepType: Optional, Create: `ENABLED`, Update: `DISABLED`},
		"trusted_certificate_authority_ids": Representation{RepType: Optional, Create: []string{trustedCertificateAuthorityIds}, Update: []string{trustedCertificateAuthorityIds2}},
		"verify_depth":                      Representation{RepType: Optional, Create: `10`, Update: `11`},
		"verify_peer_certificate":           Representation{RepType: Optional, Create: `false`, Update: `true`},
	}

	ListenerRequiredOnlyResourceLBCert = ListenerResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_lb_cert",
			Required, Create, listenerRepresentationLBCert)

	listenerRepresentationLBCert = map[string]interface{}{
		"default_backend_set_name": Representation{RepType: Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id":         Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                     Representation{RepType: Required, Create: `myListener1`},
		"port":                     Representation{RepType: Required, Create: `10`, Update: `11`},
		"protocol":                 Representation{RepType: Required, Create: `HTTP`},
		"connection_configuration": RepresentationGroup{Optional, listenerConnectionConfigurationRepresentation},
		"hostname_names":           Representation{RepType: Optional, Create: []string{`${oci_load_balancer_hostname.test_hostname.name}`}},
		"path_route_set_name":      Representation{RepType: Optional, Create: `${oci_load_balancer_path_route_set.test_path_route_set.name}`},
		"routing_policy_name":      Representation{RepType: Optional, Create: `${oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy.name}`},
		"rule_set_names":           Representation{RepType: Optional, Create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
		"ssl_configuration":        RepresentationGroup{Optional, listenerSslConfigurationRepresentationLBCert},
	}

	listenerSslConfigurationRepresentationLBCert = map[string]interface{}{
		// note: cannot specify certificate_name along with trusted_certificate_authority_ids
		"cipher_suite_name":       Representation{RepType: Optional, Create: `oci-default-ssl-cipher-suite-v1`, Update: `oci-default-ssl-cipher-suite-v1`},
		"protocols":               Representation{RepType: Optional, Create: []string{`TLSv1.2`}, Update: []string{`TLSv1.2`}},
		"server_order_preference": Representation{RepType: Optional, Create: `ENABLED`, Update: `DISABLED`},
		"certificate_name":        Representation{RepType: Optional, Create: "example_certificate_bundle", Update: "example_certificate_bundle"},
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
		name = "example_hostname_012"
	}
`
	certificateIds  = getEnvSettingWithBlankDefault("certificate_ids")
	certificateIds2 = getEnvSettingWithBlankDefault("certificate_ids2")
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerListenerResourceOciCerts_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener_with_ca"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ListenerResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
			"test_listener_with_ca", Optional, Create, listenerRepresentationOciCerts), "loadbalancer", "listener", t)

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
					GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
						"test_listener_with_ca", Required, Create, listenerRepresentationOciCerts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "myListener1"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
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
					GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_ca", Optional, Create,
						GetUpdatedRepresentationCopy("hostname_names", Representation{RepType: Optional,
							Create: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"},
							Update: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"}}, listenerRepresentationOciCerts)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.idle_timeout_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_names.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "hostname_names.0", "example_hostname_001"),
					resource.TestCheckResourceAttr(resourceName, "hostname_names.1", "example_hostname_012"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "myListener1"),
					resource.TestCheckResourceAttrSet(resourceName, "path_route_set_name"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttrSet(resourceName, "routing_policy_name"),
					resource.TestCheckResourceAttr(resourceName, "rule_set_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "oci-default-ssl-cipher-suite-v1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.server_order_preference", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.trusted_certificate_authority_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "10"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_ids.#", "1"),
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
					GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
						"test_listener_with_ca", Optional, Update, listenerRepresentationOciCerts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.idle_timeout_in_seconds", "11"),
					resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_names.0", "example_hostname_001"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "myListener1"),
					resource.TestCheckResourceAttrSet(resourceName, "path_route_set_name"),
					resource.TestCheckResourceAttr(resourceName, "port", "11"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttrSet(resourceName, "routing_policy_name"),
					resource.TestCheckResourceAttr(resourceName, "rule_set_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "oci-default-ssl-cipher-suite-v1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.server_order_preference", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.trusted_certificate_authority_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "11"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_ids.#", "1"),
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
		},
	})
}

func TestLoadBalancerListenerResourceLBCert_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener_with_lb_cert"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ListenerResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
			"test_listener_with_lb_cert", Optional, Create, listenerRepresentationLBCert), "loadbalancer", "listener", t)

	ResourceTest(t, testAccCheckLoadBalancerListenerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
					"test_listener_with_lb_cert", Required, Create, listenerRepresentationLBCert),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "myListener1"),
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
				GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_lb_cert", Optional, Create,
					GetUpdatedRepresentationCopy("hostname_names", Representation{RepType: Optional,
						Create: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"},
						Update: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"}}, listenerRepresentationLBCert)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "connection_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.idle_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", "backendSet1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_names.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "hostname_names.0", "example_hostname_001"),
				resource.TestCheckResourceAttr(resourceName, "hostname_names.1", "example_hostname_012"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "myListener1"),
				resource.TestCheckResourceAttrSet(resourceName, "path_route_set_name"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttrSet(resourceName, "routing_policy_name"),
				resource.TestCheckResourceAttr(resourceName, "rule_set_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "example_certificate_bundle"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "oci-default-ssl-cipher-suite-v1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.server_order_preference", "ENABLED"),
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
				GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_lb_cert",
					Optional, Update, listenerRepresentationLBCert),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "connection_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.idle_timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", "backendSet1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_names.0", "example_hostname_001"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "myListener1"),
				resource.TestCheckResourceAttrSet(resourceName, "path_route_set_name"),
				resource.TestCheckResourceAttr(resourceName, "port", "11"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttrSet(resourceName, "routing_policy_name"),
				resource.TestCheckResourceAttr(resourceName, "rule_set_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "oci-default-ssl-cipher-suite-v1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "example_certificate_bundle"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.server_order_preference", "DISABLED"),
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
