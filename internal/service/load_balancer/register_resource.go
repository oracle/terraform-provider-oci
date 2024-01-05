// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_load_balancer_backend", LoadBalancerBackendResource())
	tfresource.RegisterResource("oci_load_balancer_backend_set", LoadBalancerBackendSetResource())
	tfresource.RegisterResource("oci_load_balancer_certificate", LoadBalancerCertificateResource())
	tfresource.RegisterResource("oci_load_balancer_hostname", LoadBalancerHostnameResource())
	tfresource.RegisterResource("oci_load_balancer_listener", LoadBalancerListenerResource())
	tfresource.RegisterResource("oci_load_balancer_load_balancer", LoadBalancerLoadBalancerResource())
	tfresource.RegisterResource("oci_load_balancer_load_balancer_routing_policy", LoadBalancerLoadBalancerRoutingPolicyResource())
	tfresource.RegisterResource("oci_load_balancer_path_route_set", LoadBalancerPathRouteSetResource())
	tfresource.RegisterResource("oci_load_balancer_rule_set", LoadBalancerRuleSetResource())
	tfresource.RegisterResource("oci_load_balancer_ssl_cipher_suite", LoadBalancerSslCipherSuiteResource())
}
