// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	NetworkLoadBalancerBackendSetRequiredOnlyResource = NetworkLoadBalancerBackendSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, NetworkLoadBalancerBackendSetRepresentation)

	NetworkLoadBalancerBackendSetResourceConfig = NetworkLoadBalancerBackendSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Update, NetworkLoadBalancerBackendSetRepresentation)

	NetworkLoadBalancerNetworkLoadBalancerBackendSetSingularDataSourceRepresentation = map[string]interface{}{
		"backend_set_name":         acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
	}

	NetworkLoadBalancerNetworkLoadBalancerBackendSetDataSourceRepresentation = map[string]interface{}{
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"filter":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: NetworkLoadBalancerBackendSetDataSourceFilterRepresentation}}
	NetworkLoadBalancerBackendSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_load_balancer_backend_set.test_backend_set.name}`}},
	}

	NetworkLoadBalancerBackendSetRepresentation = map[string]interface{}{
		"health_checker":              acctest.RepresentationGroup{RepType: acctest.Required, Group: NetworkLoadBalancerBackendSetHealthCheckerRepresentation},
		"name":                        acctest.Representation{RepType: acctest.Required, Create: `example_backend_set`},
		"network_load_balancer_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"policy":                      acctest.Representation{RepType: acctest.Required, Create: `FIVE_TUPLE`, Update: `THREE_TUPLE`},
		"ip_version":                  acctest.Representation{RepType: acctest.Optional, Create: `IPV4`},
		"is_fail_open":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_preserve_source":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_instant_failover_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	NetworkLoadBalancerBackendSetHealthCheckerRepresentation = map[string]interface{}{
		"protocol":           acctest.Representation{RepType: acctest.Required, Create: `TCP`, Update: `TCP`},
		"interval_in_millis": acctest.Representation{RepType: acctest.Optional, Create: `10000`, Update: `30000`},
		"port":               acctest.Representation{RepType: acctest.Optional, Create: `80`, Update: `8080`},
		"request_data":       acctest.Representation{RepType: acctest.Optional, Create: `SGVsbG9Xb3JsZA==`, Update: `QnllV29ybGQ=`},
		"response_data":      acctest.Representation{RepType: acctest.Optional, Create: `SGVsbG9Xb3JsZA==`, Update: `QnllV29ybGQ=`},
		"retries":            acctest.Representation{RepType: acctest.Optional, Create: `3`, Update: `5`},
		"timeout_in_millis":  acctest.Representation{RepType: acctest.Optional, Create: `10000`, Update: `30000`},
	}

	NetworkLoadBalancerHttpBackendSetRepresentation = map[string]interface{}{
		"health_checker":           acctest.RepresentationGroup{RepType: acctest.Required, Group: nlbHttpBackendSetHealthCheckerRepresentation},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `example_backend_set`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"policy":                   acctest.Representation{RepType: acctest.Required, Create: `FIVE_TUPLE`, Update: `TWO_TUPLE`},
		"is_preserve_source":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_fail_open":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	nlbHttpBackendSetHealthCheckerRepresentation = map[string]interface{}{
		"protocol":            acctest.Representation{RepType: acctest.Required, Create: `HTTP`, Update: `HTTPS`},
		"dns":                 acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"interval_in_millis":  acctest.Representation{RepType: acctest.Optional, Create: `10000`, Update: `30000`},
		"port":                acctest.Representation{RepType: acctest.Optional, Create: `80`, Update: `8080`},
		"response_body_regex": acctest.Representation{RepType: acctest.Optional, Create: `^(?i)(true)$`, Update: `^(?i)(false)$`},
		"retries":             acctest.Representation{RepType: acctest.Optional, Create: `3`, Update: `5`},
		"return_code":         acctest.Representation{RepType: acctest.Optional, Create: `202`, Update: `204`},
		"timeout_in_millis":   acctest.Representation{RepType: acctest.Optional, Create: `10000`, Update: `30000`},
		"url_path":            acctest.Representation{RepType: acctest.Optional, Create: `/urlPath`, Update: `/urlPath2`},
	}
	NetworkLoadBalancerDnsBackendSetRepresentation = map[string]interface{}{
		"health_checker":           acctest.RepresentationGroup{RepType: acctest.Required, Group: NlbDnsBackendSetHealthCheckerRepresentation},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `example_backend_set`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"policy":                   acctest.Representation{RepType: acctest.Required, Create: `THREE_TUPLE`, Update: `FIVE_TUPLE`},
		"is_preserve_source":       acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_fail_open":             acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
	}
	NlbDnsBackendSetHealthCheckerRepresentation = map[string]interface{}{
		"protocol":           acctest.Representation{RepType: acctest.Required, Create: `DNS`, Update: `DNS`},
		"dns":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: NetworkLoadBalancerBackendSetHealthCheckerDnsRepresentation},
		"interval_in_millis": acctest.Representation{RepType: acctest.Optional, Create: `10000`, Update: `30000`},
		"port":               acctest.Representation{RepType: acctest.Optional, Create: `53`, Update: `8053`},
		"retries":            acctest.Representation{RepType: acctest.Optional, Create: `3`, Update: `5`},
		"timeout_in_millis":  acctest.Representation{RepType: acctest.Optional, Create: `10000`, Update: `30000`},
	}
	NetworkLoadBalancerBackendSetHealthCheckerDnsRepresentation = map[string]interface{}{
		"domain_name":        acctest.Representation{RepType: acctest.Required, Create: `oracle.com`, Update: `example.com`},
		"query_class":        acctest.Representation{RepType: acctest.Optional, Create: `IN`, Update: `CH`},
		"query_type":         acctest.Representation{RepType: acctest.Optional, Create: `A`, Update: `TXT`},
		"rcodes":             acctest.Representation{RepType: acctest.Optional, Create: []string{`NOERROR`}, Update: []string{`NOERROR`, `SERVFAIL`}},
		"transport_protocol": acctest.Representation{RepType: acctest.Optional, Create: `UDP`, Update: `TCP`},
	}

	NetworkLoadBalancerNlbBackendRepresentation = map[string]interface{}{
		"backend_set_name":         acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"port":                     acctest.Representation{RepType: acctest.Required, Create: `10`},
		"ip_address":               acctest.Representation{RepType: acctest.Required, Create: `10.0.0.3`},
		"is_backup":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_drain":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_offline":               acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"name":                     acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"weight":                   acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	NetworkLoadBalancerBackendSetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerRepresentation)
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerBackendSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerBackendSetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_load_balancer_backend_set.test_backend_set"
	datasourceName := "data.oci_network_load_balancer_backend_sets.test_backend_sets"
	singularDatasourceName := "data.oci_network_load_balancer_backend_set.test_backend_set"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckNetworkLoadBalancerBackendSetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, NetworkLoadBalancerBackendSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backends.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "0"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "3"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "3000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "FIVE_TUPLE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Create, NetworkLoadBalancerBackendSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backends.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "80"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.request_data", "SGVsbG9Xb3JsZA=="),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", ""),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_data", "SGVsbG9Xb3JsZA=="),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "3"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "0"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", ""),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "is_instant_failover_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_fail_open", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "FIVE_TUPLE"),

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
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Update, NetworkLoadBalancerBackendSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(resourceName, "backends.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.request_data", "QnllV29ybGQ="),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", ""),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_data", "QnllV29ybGQ="),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "5"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", ""),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "THREE_TUPLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// Update with HTTP health checker with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Create, NetworkLoadBalancerHttpBackendSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backends.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "80"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "^(?i)(true)$"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "3"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "202"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/urlPath"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_fail_open", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "FIVE_TUPLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// Update with HTTPS health checker
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Update, NetworkLoadBalancerHttpBackendSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backends.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTPS"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "^(?i)(false)$"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "5"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "is_fail_open", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_instant_failover_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "204"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/urlPath2"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "TWO_TUPLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// Create with DNS health checker with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Create, NetworkLoadBalancerDnsBackendSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backends.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.domain_name", "oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.query_class", "IN"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.query_type", "A"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.rcodes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.rcodes.0", "NOERROR"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.transport_protocol", "UDP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "53"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "DNS"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "3"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_fail_open", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "THREE_TUPLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Update with DNS health checker
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Update, NetworkLoadBalancerDnsBackendSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backends.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.domain_name", "example.com"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.query_class", "CH"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.query_type", "TXT"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.rcodes.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.rcodes.0", "NOERROR"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.rcodes.1", "SERVFAIL"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.dns.0.transport_protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "8053"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "DNS"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "5"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "is_fail_open", "false"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "FIVE_TUPLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Update with backends
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Update, NetworkLoadBalancerHttpBackendSetRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", acctest.Required, acctest.Create, NetworkLoadBalancerNlbBackendRepresentation) +
				`data "oci_network_load_balancer_backend_sets" "test_backend_sets" {
						depends_on = ["oci_network_load_balancer_backend_set.test_backend_set", "oci_network_load_balancer_backend.test_backend"]
						network_load_balancer_id = "${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}"
					}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// The state file could show either 0 or 1 backends in backend_set; depending on the order of operations.
				// If UpdateBackendSet happens first, then you will see 0. If CreateBackend happens first, then you will see 1.
				//resource.TestCheckResourceAttr(resourceName, "backends.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.0.items.0.backends.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.0.items.0.backends.0.ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.0.items.0.backends.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTPS"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "^(?i)(false)$"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "5"),
				resource.TestCheckResourceAttr(resourceName, "is_fail_open", "true"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "204"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/urlPath2"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "TWO_TUPLE"),
				// getting a validation error. so this update of ipversion is not allowed and the expected error is thrown
				//resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV6"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
			ExpectNonEmptyPlan: true,
		},

		// Force Create new by changing backend port
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Update, NetworkLoadBalancerBackendSetRepresentation) +
				`resource "oci_network_load_balancer_backend" "test_backend" {
						network_load_balancer_id = "${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}"
						backend_set_name = "${oci_network_load_balancer_backend_set.test_backend_set.name}"
						ip_address = "10.0.0.3"
						port = 80
					}

					data "oci_network_load_balancer_backend_sets" "test_backend_sets" {
						depends_on = ["oci_network_load_balancer_backend_set.test_backend_set", "oci_network_load_balancer_backend.test_backend"]
						network_load_balancer_id = "${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}"
					}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// The state file could show either 0 or 1 backends in backend_set; depending on the order of operations.
				// If UpdateBackendSet happens first, then you will see 0. If CreateBackend happens first, then you will see 1.
				//resource.TestCheckResourceAttr(resourceName, "backends.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.0.items.0.backends.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.0.items.0.backends.0.ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.0.items.0.backends.0.port", "80"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.request_data", "QnllV29ybGQ="),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", ""),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_data", "QnllV29ybGQ="),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "5"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", ""),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "THREE_TUPLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
			ExpectNonEmptyPlan: true,
		},

		// Remove backends while updating backendset
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Update, NetworkLoadBalancerHttpBackendSetRepresentation) +
				`data "oci_network_load_balancer_backend_sets" "test_backend_sets" {
						depends_on = ["oci_network_load_balancer_backend_set.test_backend_set"]
						network_load_balancer_id = "${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}"
					}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.0.items.#", "1"),
				resource.TestCheckNoResourceAttr(datasourceName, "backend_set_collection.0.items.0.backends"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "HTTPS"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", "^(?i)(false)$"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "5"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.return_code", "204"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.url_path", "/urlPath2"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_fail_open", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "TWO_TUPLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
			ExpectNonEmptyPlan: true,
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_load_balancer_backend_sets", "test_backend_sets", acctest.Optional, acctest.Update, NetworkLoadBalancerNetworkLoadBalancerBackendSetDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Update, NetworkLoadBalancerBackendSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backend_set_collection.0.items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config + NetworkLoadBalancerBackendSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Optional, acctest.Update, NetworkLoadBalancerHttpBackendSetRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerBackendSetSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_set_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backends.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.0.interval_in_millis", "30000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.0.port", "8080"),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.0.protocol", "HTTPS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.0.request_data", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.0.response_body_regex", "^(?i)(false)$"),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.0.response_data", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.0.retries", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_instant_failover_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.0.return_code", "204"),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.0.timeout_in_millis", "30000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "health_checker.0.url_path", "/urlPath2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_fail_open", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_preserve_source", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy", "TWO_TUPLE"),
			),
		},
		// verify resource import
		{
			Config:                  config + NetworkLoadBalancerBackendSetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkLoadBalancerBackendSetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkLoadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_load_balancer_backend_set" {
			noResourceFound = false
			request := oci_network_load_balancer.GetBackendSetRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.BackendSetName = &value
			}

			if value, ok := rs.Primary.Attributes["network_load_balancer_id"]; ok {
				request.NetworkLoadBalancerId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_load_balancer")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("NetworkLoadBalancerBackendSet") {
		resource.AddTestSweepers("NetworkLoadBalancerBackendSet", &resource.Sweeper{
			Name:         "NetworkLoadBalancerBackendSet",
			Dependencies: acctest.DependencyGraph["backendSet"],
			F:            sweepNetworkLoadBalancerBackendSetResource,
		})
	}
}

func sweepNetworkLoadBalancerBackendSetResource(compartment string) error {
	networkLoadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkLoadBalancerClient()
	backendSetIds, err := getNetworkLoadBalancerBackendSetIds(compartment)
	if err != nil {
		return err
	}
	for _, backendSetId := range backendSetIds {
		if ok := acctest.SweeperDefaultResourceId[backendSetId]; !ok {
			deleteBackendSetRequest := oci_network_load_balancer.DeleteBackendSetRequest{}

			deleteBackendSetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_load_balancer")
			_, error := networkLoadBalancerClient.DeleteBackendSet(context.Background(), deleteBackendSetRequest)
			if error != nil {
				fmt.Printf("Error deleting BackendSet %s %s, It is possible that the resource is already deleted. Please verify manually \n", backendSetId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkLoadBalancerBackendSetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BackendSetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkLoadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkLoadBalancerClient()

	listBackendSetsRequest := oci_network_load_balancer.ListBackendSetsRequest{}

	networkLoadBalancerIds, error := getNetworkLoadBalancerNetworkLoadBalancerIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkLoadBalancerId required for BackendSet resource requests \n")
	}
	for _, networkLoadBalancerId := range networkLoadBalancerIds {
		listBackendSetsRequest.NetworkLoadBalancerId = &networkLoadBalancerId

		listBackendSetsResponse, err := networkLoadBalancerClient.ListBackendSets(context.Background(), listBackendSetsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting BackendSet list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, backendSet := range listBackendSetsResponse.Items {
			id := *backendSet.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BackendSetId", id)
		}

	}
	return resourceIds, nil
}
