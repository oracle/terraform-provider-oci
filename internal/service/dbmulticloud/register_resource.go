// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_dbmulticloud_multi_cloud_resource_discovery", DbmulticloudMultiCloudResourceDiscoveryResource())
	tfresource.RegisterResource("oci_dbmulticloud_oracle_db_aws_identity_connector", DbmulticloudOracleDbAwsIdentityConnectorResource())
	tfresource.RegisterResource("oci_dbmulticloud_oracle_db_aws_key", DbmulticloudOracleDbAwsKeyResource())
	tfresource.RegisterResource("oci_dbmulticloud_oracle_db_azure_blob_container", DbmulticloudOracleDbAzureBlobContainerResource())
	tfresource.RegisterResource("oci_dbmulticloud_oracle_db_azure_blob_mount", DbmulticloudOracleDbAzureBlobMountResource())
	tfresource.RegisterResource("oci_dbmulticloud_oracle_db_azure_connector", DbmulticloudOracleDbAzureConnectorResource())
	tfresource.RegisterResource("oci_dbmulticloud_oracle_db_azure_vault", DbmulticloudOracleDbAzureVaultResource())
	tfresource.RegisterResource("oci_dbmulticloud_oracle_db_azure_vault_association", DbmulticloudOracleDbAzureVaultAssociationResource())
	tfresource.RegisterResource("oci_dbmulticloud_oracle_db_gcp_identity_connector", DbmulticloudOracleDbGcpIdentityConnectorResource())
	tfresource.RegisterResource("oci_dbmulticloud_oracle_db_gcp_key_ring", DbmulticloudOracleDbGcpKeyRingResource())
}
