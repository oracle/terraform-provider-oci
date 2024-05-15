// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ListenerRequiredOnlyResourceCerts = ListenerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_ca",
			acctest.Required, acctest.Create, listenerRepresentationOciCerts)

	listenerRepresentationOciCerts = map[string]interface{}{
		"default_backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `myListener1`},
		"port":                     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"protocol":                 acctest.Representation{RepType: acctest.Required, Create: `HTTP`},
		"connection_configuration": acctest.RepresentationGroup{RepType: acctest.Optional, Group: listenerConnectionConfigurationRepresentation},
		"hostname_names":           acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_load_balancer_hostname.test_hostname.name}`}},
		"path_route_set_name":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_path_route_set.test_path_route_set.name}`},
		"routing_policy_name":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy.name}`},
		"rule_set_names":           acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
		"ssl_configuration":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: listenerSslConfigurationRepresentationOciCerts},
	}
	listenerConnectionConfigurationRepresentation = map[string]interface{}{
		"idle_timeout_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	listenerSslConfigurationRepresentationOciCerts = map[string]interface{}{
		// note: cannot specify certificate_name along with trusted_certificate_authority_ids
		"certificate_ids":                   acctest.Representation{RepType: acctest.Optional, Create: []string{certificateIds}, Update: []string{certificateIds2}},
		"has_session_resumption":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"cipher_suite_name":                 acctest.Representation{RepType: acctest.Optional, Create: `oci-default-ssl-cipher-suite-v1`, Update: `oci-default-ssl-cipher-suite-v1`},
		"protocols":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`TLSv1.2`}, Update: []string{`TLSv1.2`}},
		"server_order_preference":           acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"trusted_certificate_authority_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{trustedCertificateAuthorityIds}, Update: []string{trustedCertificateAuthorityIds2}},
		"verify_depth":                      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"verify_peer_certificate":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	ListenerRequiredOnlyResourceLBCert = ListenerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_lb_cert",
			acctest.Required, acctest.Create, listenerRepresentationLBCert)

	listenerRepresentationLBCert = map[string]interface{}{
		"default_backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `myListener1`},
		"port":                     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"protocol":                 acctest.Representation{RepType: acctest.Required, Create: `HTTP`},
		"connection_configuration": acctest.RepresentationGroup{RepType: acctest.Optional, Group: listenerConnectionConfigurationRepresentation},
		"hostname_names":           acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_load_balancer_hostname.test_hostname.name}`}},
		"path_route_set_name":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_path_route_set.test_path_route_set.name}`},
		"routing_policy_name":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy.name}`},
		"rule_set_names":           acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
		"ssl_configuration":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: listenerSslConfigurationRepresentationLBCert},
	}

	listenerSslConfigurationRepresentationLBCert = map[string]interface{}{
		// note: cannot specify certificate_name along with trusted_certificate_authority_ids
		"cipher_suite_name":       acctest.Representation{RepType: acctest.Optional, Create: `oci-default-ssl-cipher-suite-v1`, Update: `oci-default-ssl-cipher-suite-v1`},
		"protocols":               acctest.Representation{RepType: acctest.Optional, Create: []string{`TLSv1.2`}, Update: []string{`TLSv1.2`}},
		"server_order_preference": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"certificate_name":        acctest.Representation{RepType: acctest.Optional, Create: "example_certificate_bundle", Update: "example_certificate_bundle"},
		"verify_depth":            acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"verify_peer_certificate": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	ListenerResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer_routing_policy", "test_load_balancer_routing_policy", acctest.Required, acctest.Create, loadBalancerRoutingPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Optional, acctest.Create, certificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_path_route_set", "test_path_route_set", acctest.Required, acctest.Create, pathRouteSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_hostname", "test_hostname", acctest.Required, acctest.Create, hostnameRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Required, acctest.Create, ruleSetRepresentation) +
		caCertificateVariableStr + privateKeyVariableStr +
		`
	resource "oci_load_balancer_hostname" "test_hostname2" {
		#Required
		hostname = "app.example.com2"
		load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
		name = "example_hostname_012"
	}
`
	certificateIds  = utils.GetEnvSettingWithBlankDefault("certificate_ids")
	certificateIds2 = utils.GetEnvSettingWithBlankDefault("certificate_ids2")
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerListenerResourceOciCerts_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener_with_ca"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ListenerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
			"test_listener_with_ca", acctest.Optional, acctest.Create, listenerRepresentationOciCerts), "loadbalancer", "listener", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerListenerDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
						"test_listener_with_ca", acctest.Required, acctest.Create, listenerRepresentationOciCerts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "myListener1"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_ca", acctest.Optional, acctest.Create,
						acctest.GetUpdatedRepresentationCopy("hostname_names", acctest.Representation{RepType: acctest.Optional,
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
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
						"test_listener_with_ca", acctest.Optional, acctest.Update, listenerRepresentationOciCerts),
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
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener_with_lb_cert"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ListenerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
			"test_listener_with_lb_cert", acctest.Optional, acctest.Create, listenerRepresentationLBCert), "loadbalancer", "listener", t)

	acctest.ResourceTest(t, testAccCheckLoadBalancerListenerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
					"test_listener_with_lb_cert", acctest.Required, acctest.Create, listenerRepresentationLBCert),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "myListener1"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_lb_cert", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("hostname_names", acctest.Representation{RepType: acctest.Optional,
						Create: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"},
						Update: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"}}, listenerRepresentationLBCert)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.has_session_resumption", "false"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "example_certificate_bundle"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "oci-default-ssl-cipher-suite-v1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.server_order_preference", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "10"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_lb_cert",
					acctest.Optional, acctest.Update, listenerRepresentationLBCert),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

func TestLoadBalancerListenerResourceLBCertToOciCerts_combo(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener_with_lb_cert_to_oci_certs"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ListenerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
			"test_listener_with_lb_cert_to_oci_certs", acctest.Optional, acctest.Create, listenerRepresentationLBCert), "loadbalancer", "listener", t)

	acctest.ResourceTest(t, testAccCheckLoadBalancerListenerDestroy, []resource.TestStep{
		// verify Create with optionals with LB cert
		{
			Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_lb_cert_to_oci_certs", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("hostname_names", acctest.Representation{RepType: acctest.Optional,
						Create: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"},
						Update: []string{"${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"}}, listenerRepresentationLBCert)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
					"test_listener_with_lb_cert_to_oci_certs", acctest.Optional, acctest.Update, listenerRepresentationOciCerts),
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
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.has_session_resumption", "true"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "oci-default-ssl-cipher-suite-v1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.server_order_preference", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.trusted_certificate_authority_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "11"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "true"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

func TestLoadBalancerListenerResourceOciCertToLBCert_combo(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerListenerResourceOciCertToLBCert_combo")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener_with_oci_cert_to_lb_cert"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ListenerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener",
			"test_listener_with_oci_cert_to_lb_cert", acctest.Optional, acctest.Create, listenerRepresentationOciCerts), "loadbalancer", "listener", t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		CheckDestroy: testAccCheckLoadBalancerListenerDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_oci_cert_to_lb_cert", acctest.Optional, acctest.Create,
						acctest.GetUpdatedRepresentationCopy("hostname_names", acctest.Representation{RepType: acctest.Optional,
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
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "oci-default-ssl-cipher-suite-v1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.server_order_preference", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.trusted_certificate_authority_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "10"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener_with_oci_cert_to_lb_cert",
						acctest.Optional, acctest.Update, listenerRepresentationLBCert),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.server_order_preference", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.trusted_certificate_authority_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "11"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "true"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
	client := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()
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
