// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_load_balancer_backend_health", LoadBalancerBackendHealthDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_backend_set_health", LoadBalancerBackendSetHealthDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_backend_sets", LoadBalancerBackendSetsDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_backends", LoadBalancerBackendsDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_certificates", LoadBalancerCertificatesDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_hostnames", LoadBalancerHostnamesDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_listener_rules", LoadBalancerListenerRulesDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_health", LoadBalancerLoadBalancerHealthDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_policies", LoadBalancerLoadBalancerPoliciesDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_protocols", LoadBalancerLoadBalancerProtocolsDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_load_balancer_routing_policies", LoadBalancerLoadBalancerRoutingPoliciesDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_load_balancer_routing_policy", LoadBalancerLoadBalancerRoutingPolicyDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_shapes", LoadBalancerLoadBalancerShapesDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_load_balancers", LoadBalancerLoadBalancersDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_path_route_sets", LoadBalancerPathRouteSetsDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_rule_set", LoadBalancerRuleSetDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_rule_sets", LoadBalancerRuleSetsDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_ssl_cipher_suite", LoadBalancerSslCipherSuiteDataSource())
	tfresource.RegisterDatasource("oci_load_balancer_ssl_cipher_suites", LoadBalancerSslCipherSuitesDataSource())
}
