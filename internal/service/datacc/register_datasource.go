// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacc

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_datacc_infrastructure", DataccInfrastructureDataSource())
	tfresource.RegisterDatasource("oci_datacc_infrastructure_scale_option", DataccInfrastructureScaleOptionDataSource())
	tfresource.RegisterDatasource("oci_datacc_infrastructures", DataccInfrastructuresDataSource())
	tfresource.RegisterDatasource("oci_datacc_maintenance_execution", DataccMaintenanceExecutionDataSource())
	tfresource.RegisterDatasource("oci_datacc_maintenance_executions", DataccMaintenanceExecutionsDataSource())
	tfresource.RegisterDatasource("oci_datacc_vm_cluster_network", DataccVmClusterNetworkDataSource())
	tfresource.RegisterDatasource("oci_datacc_vm_cluster_networks", DataccVmClusterNetworksDataSource())
	tfresource.RegisterDatasource("oci_datacc_vm_instance", DataccVmInstanceDataSource())
	tfresource.RegisterDatasource("oci_datacc_vm_instances", DataccVmInstancesDataSource())
}
