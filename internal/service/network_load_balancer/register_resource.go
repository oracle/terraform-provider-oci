// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_load_balancer

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_network_load_balancer_backend", NetworkLoadBalancerBackendResource())
	tfresource.RegisterResource("oci_network_load_balancer_backend_set", NetworkLoadBalancerBackendSetResource())
	tfresource.RegisterResource("oci_network_load_balancer_listener", NetworkLoadBalancerListenerResource())
	tfresource.RegisterResource("oci_network_load_balancer_network_load_balancer", NetworkLoadBalancerNetworkLoadBalancerResource())
	tfresource.RegisterResource("oci_network_load_balancer_network_load_balancers_backend_sets_unified", NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResource())
}
