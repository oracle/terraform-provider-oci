package dbmulticloud

import (
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("dbmulticloud", dbmulticloudResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDbmulticloudOracleDbAzureVaultAssociationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dbmulticloud_oracle_db_azure_vault_association",
	DatasourceClass:        "oci_dbmulticloud_oracle_db_azure_vault_associations",
	DatasourceItemsAttr:    "oracle_db_azure_vault_association_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oracle_db_azure_vault_association",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dbmulticloud.OracleDbAzureVaultAssociationLifecycleStateActive),
	},
}

var exportDbmulticloudMultiCloudResourceDiscoveryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dbmulticloud_multi_cloud_resource_discovery",
	DatasourceClass:        "oci_dbmulticloud_multi_cloud_resource_discoveries",
	DatasourceItemsAttr:    "multi_cloud_resource_discovery_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "multi_cloud_resource_discovery",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateSucceeded),
		string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateNeedsAttention),
	},
}

var exportDbmulticloudOracleDbAzureBlobContainerHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dbmulticloud_oracle_db_azure_blob_container",
	DatasourceClass:        "oci_dbmulticloud_oracle_db_azure_blob_containers",
	DatasourceItemsAttr:    "oracle_db_azure_blob_container_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oracle_db_azure_blob_container",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dbmulticloud.OracleDbAzureBlobContainerLifecycleStateActive),
	},
}

var exportDbmulticloudOracleDbAzureBlobMountHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dbmulticloud_oracle_db_azure_blob_mount",
	DatasourceClass:        "oci_dbmulticloud_oracle_db_azure_blob_mounts",
	DatasourceItemsAttr:    "oracle_db_azure_blob_mount_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oracle_db_azure_blob_mount",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dbmulticloud.OracleDbAzureBlobMountLifecycleStateActive),
	},
}

var exportDbmulticloudOracleDbAzureConnectorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dbmulticloud_oracle_db_azure_connector",
	DatasourceClass:        "oci_dbmulticloud_oracle_db_azure_connectors",
	DatasourceItemsAttr:    "oracle_db_azure_connector_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oracle_db_azure_connector",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dbmulticloud.OracleDbAzureConnectorLifecycleStateActive),
	},
}

var exportDbmulticloudOracleDbAzureVaultHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dbmulticloud_oracle_db_azure_vault",
	DatasourceClass:        "oci_dbmulticloud_oracle_db_azure_vaults",
	DatasourceItemsAttr:    "oracle_db_azure_vault_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oracle_db_azure_vault",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dbmulticloud.OracleDbAzureVaultLifecycleStateActive),
	},
}

var exportDbmulticloudOracleDbGcpIdentityConnectorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dbmulticloud_oracle_db_gcp_identity_connector",
	DatasourceClass:        "oci_dbmulticloud_oracle_db_gcp_identity_connectors",
	DatasourceItemsAttr:    "oracle_db_gcp_identity_connector_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oracle_db_gcp_identity_connector",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dbmulticloud.OracleDbGcpIdentityConnectorLifecycleStateActive),
	},
}

var exportDbmulticloudOracleDbGcpKeyRingHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dbmulticloud_oracle_db_gcp_key_ring",
	DatasourceClass:        "oci_dbmulticloud_oracle_db_gcp_key_rings",
	DatasourceItemsAttr:    "oracle_db_gcp_key_ring_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oracle_db_gcp_key_ring",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dbmulticloud.OracleDbGcpKeyRingLifecycleStateActive),
	},
}

var exportDbmulticloudOracleDbAwsIdentityConnectorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dbmulticloud_oracle_db_aws_identity_connector",
	DatasourceClass:        "oci_dbmulticloud_oracle_db_aws_identity_connectors",
	DatasourceItemsAttr:    "oracle_db_aws_identity_connector_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oracle_db_aws_identity_connector",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dbmulticloud.OracleDbAwsIdentityConnectorLifecycleStateActive),
	},
}

var exportDbmulticloudOracleDbAwsKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dbmulticloud_oracle_db_aws_key",
	DatasourceClass:        "oci_dbmulticloud_oracle_db_aws_keys",
	DatasourceItemsAttr:    "oracle_db_aws_key_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oracle_db_aws_key",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dbmulticloud.OracleDbAwsKeyLifecycleStateActive),
	},
}

var dbmulticloudResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDbmulticloudOracleDbAzureVaultAssociationHints},
		{TerraformResourceHints: exportDbmulticloudMultiCloudResourceDiscoveryHints},
		{TerraformResourceHints: exportDbmulticloudOracleDbAzureBlobContainerHints},
		{TerraformResourceHints: exportDbmulticloudOracleDbAzureBlobMountHints},
		{TerraformResourceHints: exportDbmulticloudOracleDbAzureConnectorHints},
		{TerraformResourceHints: exportDbmulticloudOracleDbAzureVaultHints},
		{TerraformResourceHints: exportDbmulticloudOracleDbGcpIdentityConnectorHints},
		{TerraformResourceHints: exportDbmulticloudOracleDbGcpKeyRingHints},
		{TerraformResourceHints: exportDbmulticloudOracleDbAwsIdentityConnectorHints},
		{TerraformResourceHints: exportDbmulticloudOracleDbAwsKeyHints},
	},
}
