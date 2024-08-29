// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_bds_auto_scaling_configuration", BdsAutoScalingConfigurationDataSource())
	tfresource.RegisterDatasource("oci_bds_auto_scaling_configurations", BdsAutoScalingConfigurationsDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance", BdsBdsInstanceDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance_api_key", BdsBdsInstanceApiKeyDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance_api_keys", BdsBdsInstanceApiKeysDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance_get_os_patch", BdsBdsInstanceGetOsPatchDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance_list_os_patches", BdsBdsInstanceListOsPatchesDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance_metastore_config", BdsBdsInstanceMetastoreConfigDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance_metastore_configs", BdsBdsInstanceMetastoreConfigsDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance_patch_histories", BdsBdsInstancePatchHistoriesDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance_patches", BdsBdsInstancePatchesDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance_resource_principal_configuration", BdsBdsInstanceResourcePrincipalConfigurationDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instance_resource_principal_configurations", BdsBdsInstanceResourcePrincipalConfigurationsDataSource())
	tfresource.RegisterDatasource("oci_bds_bds_instances", BdsBdsInstancesDataSource())
}
