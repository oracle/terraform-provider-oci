// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_resourcemanager_private_endpoint", ResourcemanagerPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_resourcemanager_private_endpoint_reachable_ip", ResourcemanagerPrivateEndpointReachableIpDataSource())
	tfresource.RegisterDatasource("oci_resourcemanager_private_endpoints", ResourcemanagerPrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_resourcemanager_stack", ResourcemanagerStackDataSource())
	tfresource.RegisterDatasource("oci_resourcemanager_stack_tf_state", ResourcemanagerStackTfStateDataSource())
	tfresource.RegisterDatasource("oci_resourcemanager_stacks", ResourcemanagerStacksDataSource())
}
