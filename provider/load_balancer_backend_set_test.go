// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	BackendSetRequiredOnlyResource = BackendSetResourceDependencies + `
resource "oci_load_balancer_backendset" "test_backend_set" {
	#Required
	health_checker {
		#Required
		protocol = "${var.backend_set_health_checker_protocol}"
		url_path = "${var.backend_set_health_checker_url_path}"
	}
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.backend_set_name}"
	policy = "${var.backend_set_policy}"
}
`

	BackendSetResourceConfig = BackendSetResourceDependencies + `
resource "oci_load_balancer_backendset" "test_backend_set" {
	#Required
	health_checker {
		#Required
		protocol = "${var.backend_set_health_checker_protocol}"

		#Optional
		interval_ms = "${var.backend_set_health_checker_interval_ms}"
		port = "${var.backend_set_health_checker_port}"
		response_body_regex = "${var.backend_set_health_checker_response_body_regex}"
		retries = "${var.backend_set_health_checker_retries}"
		return_code = "${var.backend_set_health_checker_return_code}"
		timeout_in_millis = "${var.backend_set_health_checker_timeout_in_millis}"
		url_path = "${var.backend_set_health_checker_url_path}"
	}
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.backend_set_name}"
	policy = "${var.backend_set_policy}"

	#Optional
	session_persistence_configuration {
		#Required
		cookie_name = "${var.backend_set_session_persistence_configuration_cookie_name}"

		#Optional
		disable_fallback = "${var.backend_set_session_persistence_configuration_disable_fallback}"
	}
	ssl_configuration {
		#Required
		certificate_name = "${var.backend_set_ssl_configuration_certificate_name}"

		#Optional
		verify_depth = "${var.backend_set_ssl_configuration_verify_depth}"
		verify_peer_certificate = "${var.backend_set_ssl_configuration_verify_peer_certificate}"
	}
}
`
	BackendSetPropertyVariables = `
variable "backend_set_health_checker_interval_ms" { default = "1000" }
variable "backend_set_health_checker_port" { default = "80" }
variable "backend_set_health_checker_protocol" { default = "HTTP" }
variable "backend_set_health_checker_response_body_regex" { default = "responseBodyRegex" }
variable "backend_set_health_checker_retries" { default = "3" }
variable "backend_set_health_checker_return_code" { default = "200" }
variable "backend_set_health_checker_timeout_in_millis" { default = "10000" }
variable "backend_set_health_checker_url_path" { default = "/healthcheck" }
variable "backend_set_name" { default = "backendSet1" }
variable "backend_set_policy" { default = "ROUND_ROBIN" }
variable "backend_set_session_persistence_configuration_cookie_name" { default = "cookieName" }
variable "backend_set_session_persistence_configuration_disable_fallback" { default = false }
variable "backend_set_ssl_configuration_certificate_name" { default = "My_certificate_bundle" }
variable "backend_set_ssl_configuration_verify_depth" { default = "3" }
variable "backend_set_ssl_configuration_verify_peer_certificate" { default = false }

`
	BackendSetResourceDependencies = LoadBalancerPropertyVariables + LoadBalancerResourceConfig
)

func TestLoadBalancerBackendSetResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_load_balancer_backendset.test_backend_set"
	datasourceName := "data.oci_load_balancer_backendsets.test_backend_sets"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + BackendSetPropertyVariables + compartmentIdVariableStr + BackendSetRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "policy", "ROUND_ROBIN"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
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
				Config: config + BackendSetPropertyVariables + compartmentIdVariableStr + BackendSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "1000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "80"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "responseBodyRegex"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/healthcheck"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "policy", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.cookie_name", "cookieName"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.disable_fallback", "false"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "My_certificate_bundle"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "3"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "backend_set_health_checker_interval_ms" { default = "2000" }
variable "backend_set_health_checker_port" { default = 11 }
variable "backend_set_health_checker_protocol" { default = "HTTP" }
variable "backend_set_health_checker_response_body_regex" { default = "responseBodyRegex2" }
variable "backend_set_health_checker_retries" { default = 11 }
variable "backend_set_health_checker_return_code" { default = 11 }
variable "backend_set_health_checker_timeout_in_millis" { default = 11 }
variable "backend_set_health_checker_url_path" { default = "urlPath2" }
variable "backend_set_name" { default = "backendSet1" }
variable "backend_set_policy" { default = "ROUND_ROBIN" }
variable "backend_set_session_persistence_configuration_cookie_name" { default = "cookieName" }
variable "backend_set_session_persistence_configuration_disable_fallback" { default = true }
variable "backend_set_ssl_configuration_certificate_name" { default = "My_certificate_bundle" }
variable "backend_set_ssl_configuration_verify_depth" { default = 11 }
variable "backend_set_ssl_configuration_verify_peer_certificate" { default = true }

                ` + compartmentIdVariableStr + BackendSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "policy", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.cookie_name", "cookieName"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "My_certificate_bundle"),
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
			// verify updates to Force New parameters.
			{
				Config: config + `
variable "backend_set_health_checker_interval_ms" { default = "2000" }
variable "backend_set_health_checker_port" { default = 11 }
variable "backend_set_health_checker_protocol" { default = "TCP" }
variable "backend_set_health_checker_response_body_regex" { default = "responseBodyRegex2" }
variable "backend_set_health_checker_retries" { default = 11 }
variable "backend_set_health_checker_return_code" { default = 11 }
variable "backend_set_health_checker_timeout_in_millis" { default = 11 }
variable "backend_set_health_checker_url_path" { default = "urlPath2" }
variable "backend_set_name" { default = "backendSet2" }
variable "backend_set_policy" { default = "policy2" }
variable "backend_set_session_persistence_configuration_cookie_name" { default = "cookieName2" }
variable "backend_set_session_persistence_configuration_disable_fallback" { default = true }
variable "backend_set_ssl_configuration_certificate_name" { default = "certificateName2" }
variable "backend_set_ssl_configuration_verify_depth" { default = 11 }
variable "backend_set_ssl_configuration_verify_peer_certificate" { default = true }

                ` + compartmentIdVariableStr2 + BackendSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "2000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "responseBodyRegex2"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "11"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "urlPath2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet2"),
					resource.TestCheckResourceAttr(resourceName, "policy", "policy2"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.cookie_name", "cookieName2"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "certificateName2"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "11"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "true"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "backend_set_health_checker_interval_ms" { default = "2000" }
variable "backend_set_health_checker_port" { default = 11 }
variable "backend_set_health_checker_protocol" { default = "TCP" }
variable "backend_set_health_checker_response_body_regex" { default = "responseBodyRegex2" }
variable "backend_set_health_checker_retries" { default = 11 }
variable "backend_set_health_checker_return_code" { default = 11 }
variable "backend_set_health_checker_timeout_in_millis" { default = 11 }
variable "backend_set_health_checker_url_path" { default = "urlPath2" }
variable "backend_set_name" { default = "backendSet2" }
variable "backend_set_policy" { default = "policy2" }
variable "backend_set_session_persistence_configuration_cookie_name" { default = "cookieName2" }
variable "backend_set_session_persistence_configuration_disable_fallback" { default = true }
variable "backend_set_ssl_configuration_certificate_name" { default = "certificateName2" }
variable "backend_set_ssl_configuration_verify_depth" { default = 11 }
variable "backend_set_ssl_configuration_verify_peer_certificate" { default = true }

data "oci_load_balancer_backendsets" "test_backend_sets" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"

    filter {
    	name = "id"
    	values = ["${oci_load_balancer_backendset.test_backend_set.id}"]
    }
}
                ` + compartmentIdVariableStr2 + BackendSetResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.health_checker.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.health_checker.0.interval_ms", "2000"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.health_checker.0.port", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.health_checker.0.response_body_regex", "responseBodyRegex2"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.health_checker.0.retries", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.health_checker.0.return_code", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.health_checker.0.timeout_in_millis", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.health_checker.0.url_path", "urlPath2"),
					resource.TestCheckResourceAttrSet(datasourceName, "backend_sets.0.load_balancer_id"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.name", "backendSet2"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.policy", "policy2"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.session_persistence_configuration.0.cookie_name", "cookieName2"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.ssl_configuration.0.certificate_name", "certificateName2"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.ssl_configuration.0.verify_depth", "11"),
					resource.TestCheckResourceAttr(datasourceName, "backend_sets.0.ssl_configuration.0.verify_peer_certificate", "true"),
				),
			},
		},
	})
}

func TestLoadBalancerBackendSetResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_backendset.test_backend_set"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + BackendSetPropertyVariables + compartmentIdVariableStr + BackendSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "backend.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.backup", "false"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.drain", "false"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "backend.0.name"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.offline", "false"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.port", "80"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.weight", "10"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "1000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "80"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "responseBodyRegex"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/healthcheck"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "policy", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.cookie_name", "cookieName"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.disable_fallback", "false"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "My_certificate_bundle"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "3"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "backend_set_backend_backup" { default = false }
variable "backend_set_backend_drain" { default = false }
variable "backend_set_backend_ip_address" { default = "ipAddress" }
variable "backend_set_backend_offline" { default = false }
variable "backend_set_backend_port" { default = "80" }
variable "backend_set_backend_weight" { default = 10 }
variable "backend_set_health_checker_interval_ms" { default = "1000" }
variable "backend_set_health_checker_port" { default = "80" }
variable "backend_set_health_checker_protocol" { default = "HTTP" }
variable "backend_set_health_checker_response_body_regex" { default = "responseBodyRegex" }
variable "backend_set_health_checker_retries" { default = "3" }
variable "backend_set_health_checker_return_code" { default = "200" }
variable "backend_set_health_checker_timeout_in_millis" { default = "10000" }
variable "backend_set_health_checker_url_path" { default = "/healthcheck" }
variable "backend_set_name" { default = "backendSet2" }
variable "backend_set_policy" { default = "ROUND_ROBIN" }
variable "backend_set_session_persistence_configuration_cookie_name" { default = "cookieName" }
variable "backend_set_session_persistence_configuration_disable_fallback" { default = false }
variable "backend_set_ssl_configuration_certificate_name" { default = "My_certificate_bundle" }
variable "backend_set_ssl_configuration_verify_depth" { default = "3" }
variable "backend_set_ssl_configuration_verify_peer_certificate" { default = false }
				` + compartmentIdVariableStr + BackendSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "backend.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.backup", "false"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.drain", "false"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "backend.0.name"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.offline", "false"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.port", "80"),
					resource.TestCheckResourceAttr(resourceName, "backend.0.weight", "10"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_ms", "1000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "80"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "responseBodyRegex"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
					resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/healthcheck"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "backendSet2"),
					resource.TestCheckResourceAttr(resourceName, "policy", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.cookie_name", "cookieName"),
					resource.TestCheckResourceAttr(resourceName, "session_persistence_configuration.0.disable_fallback", "false"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "My_certificate_bundle"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "3"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter Name but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
