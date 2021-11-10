// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v51/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v51/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BackendSetRequiredOnlyResource = BackendSetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentationOciCerts)

	backendSetDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           RepresentationGroup{Required, backendSetDataSourceFilterRepresentation}}
	backendSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `name`},
		"values": Representation{RepType: Required, Create: []string{`${oci_load_balancer_backend_set.test_backend_set.name}`}},
	}

	backendSetRepresentationOciCerts = map[string]interface{}{
		"health_checker":                    RepresentationGroup{Required, backendSetHealthCheckerRepresentation},
		"load_balancer_id":                  Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                              Representation{RepType: Required, Create: `backendSet1`},
		"policy":                            Representation{RepType: Required, Create: `LEAST_CONNECTIONS`},
		"session_persistence_configuration": RepresentationGroup{Optional, backendSetSessionPersistenceConfigurationRepresentation},
		"ssl_configuration":                 RepresentationGroup{Optional, backendSetSslConfigurationRepresentationOciCerts},
	}

	backendSetRepresentation = map[string]interface{}{
		"health_checker":                    RepresentationGroup{Required, backendSetHealthCheckerRepresentation},
		"load_balancer_id":                  Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                              Representation{RepType: Required, Create: `backendSet1`},
		"policy":                            Representation{RepType: Required, Create: `LEAST_CONNECTIONS`},
		"session_persistence_configuration": RepresentationGroup{Optional, backendSetSessionPersistenceConfigurationRepresentation},
		"ssl_configuration":                 RepresentationGroup{Optional, backendSetSslConfigurationRepresentationLB},
	}

	backendSetSslConfigurationRepresentationOciCerts = map[string]interface{}{
		"cipher_suite_name":                 Representation{RepType: Optional, Create: `example_cipher_suite`, Update: `cipherSuiteName2`},
		"protocols":                         Representation{RepType: Optional, Create: []string{`protocols`}, Update: []string{`protocols2`}},
		"trusted_certificate_authority_ids": Representation{RepType: Optional, Create: []string{trustedCertificateAuthorityIds}, Update: []string{trustedCertificateAuthorityIds2}},
		"verify_depth":                      Representation{RepType: Optional, Create: `6`},
		"verify_peer_certificate":           Representation{RepType: Optional, Create: `false`},
	}

	backendSetSslConfigurationRepresentationLB = map[string]interface{}{
		"cipher_suite_name":       Representation{RepType: Optional, Create: `oci-default-ssl-cipher-suite-v1`, Update: `oci-default-ssl-cipher-suite-v1`},
		"protocols":               Representation{RepType: Optional, Create: []string{`TLSv1.2`}, Update: []string{`TLSv1.2`}},
		"certificate_name":        Representation{RepType: Optional, Create: "example_certificate_bundle", Update: "example_certificate_bundle"},
		"verify_depth":            Representation{RepType: Optional, Create: `6`},
		"verify_peer_certificate": Representation{RepType: Optional, Create: `false`},
	}

	backendSet2Representation = map[string]interface{}{
		"health_checker":   RepresentationGroup{Required, backendSetHealthCheckerRepresentation},
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer2.id}`},
		"name":             Representation{RepType: Required, Create: `backendSet2`},
		"policy":           Representation{RepType: Required, Create: `LEAST_CONNECTIONS`},
	}

	backendSetLBRepresentation = RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(backendSetRepresentation, []string{`session_persistence_configuration`}), map[string]interface{}{
		"lb_cookie_session_persistence_configuration": RepresentationGroup{Optional, backendSetLbCookieSessionPersistenceConfigurationRepresentation},
	})

	backendSetLBRepresentationOciCerts = RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(backendSetRepresentationOciCerts, []string{`session_persistence_configuration`}), map[string]interface{}{
		"lb_cookie_session_persistence_configuration": RepresentationGroup{Optional, backendSetLbCookieSessionPersistenceConfigurationRepresentation},
	})

	backendSetHealthCheckerRepresentation = map[string]interface{}{
		"protocol":            Representation{RepType: Required, Create: `HTTP`},
		"interval_ms":         Representation{RepType: Optional, Create: `1000`, Update: `2000`},
		"port":                Representation{RepType: Optional, Create: `10`, Update: `11`},
		"response_body_regex": Representation{RepType: Optional, Create: `.*`, Update: `responseBodyRegex2`},
		"retries":             Representation{RepType: Optional, Create: `10`, Update: `11`},
		"return_code":         Representation{RepType: Optional, Create: `200`, Update: `11`},
		"timeout_in_millis":   Representation{RepType: Optional, Create: `10000`, Update: `11`},
		"url_path":            Representation{RepType: Required, Create: `/healthcheck`, Update: `urlPath2`},
	}
	backendSetLbCookieSessionPersistenceConfigurationRepresentation = map[string]interface{}{
		"cookie_name":        Representation{RepType: Optional, Create: `example_cookie`, Update: `cookieName2`},
		"disable_fallback":   Representation{RepType: Optional, Create: `false`, Update: `true`},
		"domain":             Representation{RepType: Optional, Create: `example.oracle.com`, Update: `domain2`},
		"is_http_only":       Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_secure":          Representation{RepType: Optional, Create: `false`, Update: `true`},
		"max_age_in_seconds": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"path":               Representation{RepType: Optional, Create: `/tmp/example`, Update: `/tmp/example2`},
	}
	backendSetSessionPersistenceConfigurationRepresentation = map[string]interface{}{
		"cookie_name":      Representation{RepType: Required, Create: `example_cookie`},
		"disable_fallback": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	backendSetSslConfigurationRepresentation = map[string]interface{}{
		"certificate_name":        Representation{RepType: Required, Create: `${oci_load_balancer_certificate.test_certificate.certificate_name}`},
		"verify_depth":            Representation{RepType: Optional, Create: `6`},
		"verify_peer_certificate": Representation{RepType: Optional, Create: `false`},
	}

	BackendSetResourceDependencies = GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Optional, Create, certificateRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerResourceDependencies + caCertificateVariableStr + privateKeyVariableStr

	trustedCertificateAuthorityIds  = getEnvSettingWithBlankDefault("trusted_certificate_authority_ids")
	trustedCertificateAuthorityIds2 = getEnvSettingWithBlankDefault("trusted_certificate_authority_ids2")
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerBackendSetResourceOciCerts_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerBackendSetResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_backend_set.test_backend_set"
	datasourceName := "data.oci_load_balancer_backend_sets.test_backend_sets"

	var resId, resId2 string

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+BackendSetResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Create, backendSetRepresentationOciCerts), "loadbalancer", "backendSet", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerBackendSetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentationOciCerts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BackendSetResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Create, backendSetRepresentationOciCerts) +
					// @CODEGEN Add a backend to load balancer to validate TypeSet schema on backends during a GET in the following test steps: updates and data sources
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "1000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "10"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "10"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/healthcheck"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.disable_fallback", "false"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.trusted_certificate_authority_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Update, backendSetRepresentationOciCerts) +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					CheckResourceSetContainsElementWithProperties(resourceName, "backend", map[string]string{
						"backup":     "true",
						"drain":      "true",
						"ip_address": "10.0.0.3",
						"offline":    "true",
						"port":       "10",
						"weight":     "11",
					},
						[]string{
							"name",
						}),
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "2000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "responseBodyRegex2"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "urlPath2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.trusted_certificate_authority_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_load_balancer_backend_sets", "test_backend_sets", Optional, Update, backendSetDataSourceRepresentation) +
					compartmentIdVariableStr + BackendSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Update, backendSetRepresentationOciCerts) +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "backendsets.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.backup", "true"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.drain", "true"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.ip_address", "10.0.0.3"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.offline", "true"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.port", "10"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.weight", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.interval_ms", "2000"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.port", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.response_body_regex", "responseBodyRegex2"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.retries", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.return_code", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.timeout_in_millis", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.url_path", "urlPath2"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.name", "backendSet1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.verify_peer_certificate", "false"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.trusted_certificate_authority_ids.#", "1"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"state",
				},
				ResourceName: resourceName,
			},
			// verify update with LB session persistence
			{
				Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Create, backendSetLBRepresentationOciCerts) +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "1000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "10"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "10"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/healthcheck"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.disable_fallback", "false"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_http_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.path", "/tmp/example"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BackendSetResourceDependencies,
			},
			// verify create with LB session persistence
			{
				Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Create, backendSetLBRepresentationOciCerts) +
					// @CODEGEN Add a backend to load balancer to validate TypeSet schema on backends during a GET in the following test steps: updates and data sources
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "1000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "10"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "10"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/healthcheck"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.disable_fallback", "false"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_http_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.path", "/tmp/example"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "example_cipher_suite"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.trusted_certificate_authority_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
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
				Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Update, backendSetLBRepresentationOciCerts) +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					CheckResourceSetContainsElementWithProperties(resourceName, "backend", map[string]string{
						"backup":     "true",
						"drain":      "true",
						"ip_address": "10.0.0.3",
						"offline":    "true",
						"port":       "10",
						"weight":     "11",
					},
						[]string{
							"name",
						}),
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "2000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "responseBodyRegex2"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "urlPath2"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "cookieName2"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.domain", "domain2"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_http_only", "true"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_secure", "true"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "11"),
					resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.path", "/tmp/example2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "cipherSuiteName2"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.trusted_certificate_authority_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_load_balancer_backend_sets", "test_backend_sets", Optional, Update, backendSetDataSourceRepresentation) +
					compartmentIdVariableStr + BackendSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Update, backendSetLBRepresentationOciCerts) +
					GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "backendsets.#", "1"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "backendsets.0.backend", map[string]string{
						"backup":     "true",
						"drain":      "true",
						"ip_address": "10.0.0.3",
						"offline":    "true",
						"port":       "10",
						"weight":     "11",
					},
						[]string{
							"backup",
							"drain",
							"ip_address",
							"name",
							"offline",
							"port",
							"weight",
						}),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.interval_ms", "2000"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.port", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.response_body_regex", "responseBodyRegex2"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.retries", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.return_code", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.timeout_in_millis", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.url_path", "urlPath2"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.cookie_name", "cookieName2"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.domain", "domain2"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.is_http_only", "true"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.is_secure", "true"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.path", "/tmp/example2"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.name", "backendSet1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.cipher_suite_name", "cipherSuiteName2"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.protocols.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.trusted_certificate_authority_ids.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.verify_peer_certificate", "false"),
				),
			},
		},
	})
}

func TestLoadBalancerBackendSetResourceLB_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerBackendSetResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_backend_set.test_backend_set"
	datasourceName := "data.oci_load_balancer_backend_sets.test_backend_sets"

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+BackendSetResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Create, backendSetRepresentation), "loadbalancer", "backendSet", t)

	ResourceTest(t, testAccCheckLoadBalancerBackendSetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
				resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BackendSetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Create, backendSetRepresentation) +
				// @CODEGEN Add a backend to load balancer to validate TypeSet schema on backends during a GET in the following test steps: updates and data sources
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "1000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", ".*"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "10"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "200"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/healthcheck"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
				resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),
				resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.cookie_name", "example_cookie"),
				resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.disable_fallback", "false"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "example_certificate_bundle"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Update, backendSetRepresentation) +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				CheckResourceSetContainsElementWithProperties(resourceName, "backend", map[string]string{
					"backup":     "true",
					"drain":      "true",
					"ip_address": "10.0.0.3",
					"offline":    "true",
					"port":       "10",
					"weight":     "11",
				},
					[]string{
						"name",
					}),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "2000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "11"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "responseBodyRegex2"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "11"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "11"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "11"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "urlPath2"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
				resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),
				resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.cookie_name", "example_cookie"),
				resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.disable_fallback", "true"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "example_certificate_bundle"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_load_balancer_backend_sets", "test_backend_sets", Optional, Update, backendSetDataSourceRepresentation) +
				compartmentIdVariableStr + BackendSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Update, backendSetRepresentation) +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "backendsets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.backup", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.drain", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.offline", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.port", "10"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.backend.0.weight", "11"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.interval_ms", "2000"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.port", "11"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.response_body_regex", "responseBodyRegex2"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.retries", "11"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.return_code", "11"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.timeout_in_millis", "11"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.url_path", "urlPath2"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.name", "backendSet1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.policy", "LEAST_CONNECTIONS"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.session_persistence_configuration.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.session_persistence_configuration.0.cookie_name", "example_cookie"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.session_persistence_configuration.0.disable_fallback", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.verify_depth", "6"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.verify_peer_certificate", "false"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.certificate_name", "example_certificate_bundle"),
			),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"state",
			},
			ResourceName: resourceName,
		},
		// verify Update with LB session persistence
		{
			Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Create, backendSetLBRepresentation) +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "1000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", ".*"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "10"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "200"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/healthcheck"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.disable_fallback", "false"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_http_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_secure", "false"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.path", "/tmp/example"),
				resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
				resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BackendSetResourceDependencies,
		},
		// verify Create with LB session persistence
		{
			Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Create, backendSetLBRepresentation) +
				// @CODEGEN Add a backend to load balancer to validate TypeSet schema on backends during a GET in the following test steps: updates and data sources
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "1000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", ".*"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "10"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "200"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/healthcheck"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.disable_fallback", "false"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_http_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_secure", "false"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.path", "/tmp/example"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
				resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "oci-default-ssl-cipher-suite-v1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "example_certificate_bundle"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
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
			Config: config + compartmentIdVariableStr + BackendSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Update, backendSetLBRepresentation) +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				CheckResourceSetContainsElementWithProperties(resourceName, "backend", map[string]string{
					"backup":     "true",
					"drain":      "true",
					"ip_address": "10.0.0.3",
					"offline":    "true",
					"port":       "10",
					"weight":     "11",
				},
					[]string{
						"name",
					}),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "2000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "11"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "responseBodyRegex2"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "11"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "11"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "11"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "urlPath2"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "cookieName2"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.disable_fallback", "true"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.domain", "domain2"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_http_only", "true"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.is_secure", "true"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "lb_cookie_session_persistence_configuration.0.path", "/tmp/example2"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
				resource.TestCheckResourceAttr(resourceName, "policy", "LEAST_CONNECTIONS"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.cipher_suite_name", "oci-default-ssl-cipher-suite-v1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.protocols.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "example_certificate_bundle"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "6"),
				resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_load_balancer_backend_sets", "test_backend_sets", Optional, Update, backendSetDataSourceRepresentation) +
				compartmentIdVariableStr + BackendSetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Optional, Update, backendSetLBRepresentation) +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "backendsets.#", "1"),
				CheckResourceSetContainsElementWithProperties(datasourceName, "backendsets.0.backend", map[string]string{
					"backup":     "true",
					"drain":      "true",
					"ip_address": "10.0.0.3",
					"offline":    "true",
					"port":       "10",
					"weight":     "11",
				},
					[]string{
						"backup",
						"drain",
						"ip_address",
						"name",
						"offline",
						"port",
						"weight",
					}),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.interval_ms", "2000"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.port", "11"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.response_body_regex", "responseBodyRegex2"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.retries", "11"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.return_code", "11"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.timeout_in_millis", "11"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.health_checker.0.url_path", "urlPath2"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.cookie_name", "cookieName2"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.disable_fallback", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.domain", "domain2"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.is_http_only", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.is_secure", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "11"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.lb_cookie_session_persistence_configuration.0.path", "/tmp/example2"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.name", "backendSet1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.policy", "LEAST_CONNECTIONS"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.cipher_suite_name", "oci-default-ssl-cipher-suite-v1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.protocols.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.certificate_name", "example_certificate_bundle"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.verify_depth", "6"),
				resource.TestCheckResourceAttr(datasourceName, "backendsets.0.ssl_configuration.0.verify_peer_certificate", "false"),
			),
		},
	})
}

func testAccCheckLoadBalancerBackendSetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_backend_set" {
			noResourceFound = false
			request := oci_load_balancer.GetBackendSetRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.BackendSetName = &value
			}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "load_balancer")

			_, err := client.GetBackendSet(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("LoadBalancerBackendSet") {
		resource.AddTestSweepers("LoadBalancerBackendSet", &resource.Sweeper{
			Name:         "LoadBalancerBackendSet",
			Dependencies: DependencyGraph["backendSet"],
			F:            sweepLoadBalancerBackendSetResource,
		})
	}
}

func sweepLoadBalancerBackendSetResource(compartment string) error {
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()
	backendSetIds, err := getBackendSetIds(compartment)
	if err != nil {
		return err
	}
	for _, backendSetId := range backendSetIds {
		if ok := SweeperDefaultResourceId[backendSetId]; !ok {
			deleteBackendSetRequest := oci_load_balancer.DeleteBackendSetRequest{}

			deleteBackendSetRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeleteBackendSet(context.Background(), deleteBackendSetRequest)
			if error != nil {
				fmt.Printf("Error deleting BackendSet %s %s, It is possible that the resource is already deleted. Please verify manually \n", backendSetId, error)
				continue
			}
		}
	}
	return nil
}

func getBackendSetIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "BackendSetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()

	listBackendSetsRequest := oci_load_balancer.ListBackendSetsRequest{}

	loadBalancerIds, error := getLoadBalancerIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting loadBalancerId required for BackendSet resource requests \n")
	}
	for _, loadBalancerId := range loadBalancerIds {
		listBackendSetsRequest.LoadBalancerId = &loadBalancerId

		listBackendSetsResponse, err := loadBalancerClient.ListBackendSets(context.Background(), listBackendSetsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting BackendSet list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, backendSet := range listBackendSetsResponse.Items {
			id := *backendSet.Name
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "BackendSetId", id)
		}

	}
	return resourceIds, nil
}
