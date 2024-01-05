// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_containerengine_addon", ContainerengineAddonResource())
	tfresource.RegisterResource("oci_containerengine_cluster", ContainerengineClusterResource())
	tfresource.RegisterResource("oci_containerengine_cluster_workload_mapping", ContainerengineClusterWorkloadMappingResource())
	tfresource.RegisterResource("oci_containerengine_cluster_complete_credential_rotation_management", ContainerengineClusterCompleteCredentialRotationManagementResource())
	tfresource.RegisterResource("oci_containerengine_cluster_start_credential_rotation_management", ContainerengineClusterStartCredentialRotationManagementResource())
	tfresource.RegisterResource("oci_containerengine_node_pool", ContainerengineNodePoolResource())
	tfresource.RegisterResource("oci_containerengine_virtual_node_pool", ContainerengineVirtualNodePoolResource())
}
