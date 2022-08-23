// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_containerengine_cluster_kube_config", ContainerengineClusterKubeConfigDataSource())
	tfresource.RegisterDatasource("oci_containerengine_cluster_option", ContainerengineClusterOptionDataSource())
	tfresource.RegisterDatasource("oci_containerengine_clusters", ContainerengineClustersDataSource())
	tfresource.RegisterDatasource("oci_containerengine_migrate_to_native_vcn_status", ContainerengineMigrateToNativeVcnStatusDataSource())
	tfresource.RegisterDatasource("oci_containerengine_node_pool", ContainerengineNodePoolDataSource())
	tfresource.RegisterDatasource("oci_containerengine_node_pool_option", ContainerengineNodePoolOptionDataSource())
	tfresource.RegisterDatasource("oci_containerengine_node_pools", ContainerengineNodePoolsDataSource())
	tfresource.RegisterDatasource("oci_containerengine_work_request_errors", ContainerengineWorkRequestErrorsDataSource())
	tfresource.RegisterDatasource("oci_containerengine_work_request_log_entries", ContainerengineWorkRequestLogEntriesDataSource())
	tfresource.RegisterDatasource("oci_containerengine_work_requests", ContainerengineWorkRequestsDataSource())
}
