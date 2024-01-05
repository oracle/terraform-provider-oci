// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_containerengine_addon", ContainerengineAddonDataSource())
	tfresource.RegisterDatasource("oci_containerengine_addon_options", ContainerengineAddonOptionsDataSource())
	tfresource.RegisterDatasource("oci_containerengine_addons", ContainerengineAddonsDataSource())
	tfresource.RegisterDatasource("oci_containerengine_cluster_credential_rotation_status", ContainerengineClusterCredentialRotationStatusDataSource())
	tfresource.RegisterDatasource("oci_containerengine_cluster_kube_config", ContainerengineClusterKubeConfigDataSource())
	tfresource.RegisterDatasource("oci_containerengine_cluster_option", ContainerengineClusterOptionDataSource())
	tfresource.RegisterDatasource("oci_containerengine_cluster_workload_mapping", ContainerengineClusterWorkloadMappingDataSource())
	tfresource.RegisterDatasource("oci_containerengine_cluster_workload_mappings", ContainerengineClusterWorkloadMappingsDataSource())
	tfresource.RegisterDatasource("oci_containerengine_clusters", ContainerengineClustersDataSource())
	tfresource.RegisterDatasource("oci_containerengine_migrate_to_native_vcn_status", ContainerengineMigrateToNativeVcnStatusDataSource())
	tfresource.RegisterDatasource("oci_containerengine_node_pool", ContainerengineNodePoolDataSource())
	tfresource.RegisterDatasource("oci_containerengine_node_pool_option", ContainerengineNodePoolOptionDataSource())
	tfresource.RegisterDatasource("oci_containerengine_node_pools", ContainerengineNodePoolsDataSource())
	tfresource.RegisterDatasource("oci_containerengine_pod_shapes", ContainerenginePodShapesDataSource())
	tfresource.RegisterDatasource("oci_containerengine_virtual_node_pool", ContainerengineVirtualNodePoolDataSource())
	tfresource.RegisterDatasource("oci_containerengine_virtual_node_pools", ContainerengineVirtualNodePoolsDataSource())
	tfresource.RegisterDatasource("oci_containerengine_work_request_errors", ContainerengineWorkRequestErrorsDataSource())
	tfresource.RegisterDatasource("oci_containerengine_work_request_log_entries", ContainerengineWorkRequestLogEntriesDataSource())
	tfresource.RegisterDatasource("oci_containerengine_work_requests", ContainerengineWorkRequestsDataSource())
}
