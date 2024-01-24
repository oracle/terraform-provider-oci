// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_load_balancer

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_network_load_balancer_backend_health", NetworkLoadBalancerBackendHealthDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_backend_set", NetworkLoadBalancerBackendSetDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_backend_sets", NetworkLoadBalancerBackendSetsDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_backend_set_health", NetworkLoadBalancerBackendSetHealthDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_backends", NetworkLoadBalancerBackendsDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_listener", NetworkLoadBalancerListenerDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_listeners", NetworkLoadBalancerListenersDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_network_load_balancer", NetworkLoadBalancerNetworkLoadBalancerDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_network_load_balancer_health", NetworkLoadBalancerNetworkLoadBalancerHealthDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_network_load_balancers", NetworkLoadBalancerNetworkLoadBalancersDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_network_load_balancers_policies", NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSource())
	tfresource.RegisterDatasource("oci_network_load_balancer_network_load_balancers_protocols", NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSource())
}
