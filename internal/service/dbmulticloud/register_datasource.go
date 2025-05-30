// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_dbmulticloud_multi_cloud_resource_discoveries", DbmulticloudMultiCloudResourceDiscoveriesDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_multi_cloud_resource_discovery", DbmulticloudMultiCloudResourceDiscoveryDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_blob_container", DbmulticloudOracleDbAzureBlobContainerDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_blob_containers", DbmulticloudOracleDbAzureBlobContainersDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_blob_mount", DbmulticloudOracleDbAzureBlobMountDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_blob_mounts", DbmulticloudOracleDbAzureBlobMountsDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_connector", DbmulticloudOracleDbAzureConnectorDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_connectors", DbmulticloudOracleDbAzureConnectorsDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_key", DbmulticloudOracleDbAzureKeyDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_keys", DbmulticloudOracleDbAzureKeysDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_vault", DbmulticloudOracleDbAzureVaultDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_vault_association", DbmulticloudOracleDbAzureVaultAssociationDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_vault_associations", DbmulticloudOracleDbAzureVaultAssociationsDataSource())
	tfresource.RegisterDatasource("oci_dbmulticloud_oracle_db_azure_vaults", DbmulticloudOracleDbAzureVaultsDataSource())
}
