// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data Plane Integration
//
// <b>Microsoft Azure:</b> <br>
// <b>Oracle Azure Connector Resource:</b>:&nbsp;&nbsp;The Oracle Azure Connector Resource is used to install the Azure Arc Server on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
//  The supported method to install the Azure Arc Server (Azure Identity) on the Exadata VM cluster:
// <ul>
//  <li>Using a Bearer Access Token</li>
// </ul>
// <b>Oracle Azure Blob Container Resource:</b>&nbsp;&nbsp;The Oracle Azure Blob Container Resource is used to capture the details of an Azure Blob Container.
// This resource can then be reused across multiple Exadata VM clusters in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D) to mount the Azure container.
// <b>Oracle Azure Blob Mount Resource:</b>&nbsp;&nbsp;The Oracle Azure Blob Mount Resource is used to mount an Azure Blob Container on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// It relies on both the Oracle Azure Connector and the Oracle Azure Blob Container Resource to perform the mount operation.
// <b>Discover Azure Vaults and Keys Resource:</b>&nbsp;&nbsp;The Discover Oracle Azure Vaults and Azure Keys Resource is used to discover Azure Vaults and the associated encryption keys available in your Azure project.
// <b>Oracle Azure Vault:</b>&nbsp;&nbsp;The Oracle Azure Vault Resource is used to manage Azure Vaults within Oracle Cloud Infrastructure (OCI) for use with services such as Oracle Exadata Database Service on Dedicated Infrastructure.
// <b>Oracle Azure Key:</b>&nbsp;&nbsp;Oracle Azure Key Resource is used to register and manage a Oracle Azure Key Key within Oracle Cloud Infrastructure (OCI) under an associated Azure Vault.
// <br>
// <b>Google Cloud:</b><br>
// <b>Oracle Google Cloud Connector Resource:</b>&nbsp;&nbsp;The Oracle Google Cloud Connector Resource is used to install the Google Cloud Identity Connector on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// <b>Discover Google Key Rings and Keys Resource:</b>&nbsp;&nbsp;The Discover Google Key Rings and Keys Resource is used to discover Google Cloud Key Rings and the associated encryption keys available in your Google Cloud project.
// <b>Google Key Rings Resource:</b>&nbsp;&nbsp;The Google Key Rings Resource is used to register and manage Google Cloud Key Rings within Oracle Cloud Infrastructure (OCI) for use with services such as Oracle Exadata Database Service on Dedicated Infrastructure.
// <b>Google Key Resource:</b>&nbsp;&nbsp;The Google Key Resource is used to register and manage a Google Cloud Key within Oracle Cloud Infrastructure (OCI) under an associated Google Key Ring.
// <br>
// <b>AWS</b>:<br>
// <b>Oracle AWS Connector Resource:</b>&nbsp;&nbsp;The Oracle AWS Connector Resource is used to install the AWS Identity Connector on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// <b>Google AWS Key Resource:</b>&nbsp;&nbsp;The Oracle AWS Key Resource is used to register and manage a AWS Key within Oracle Cloud Infrastructure (OCI).
//

package dbmulticloud

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateAzureConnector           OperationTypeEnum = "CREATE_AZURE_CONNECTOR"
	OperationTypeDeleteAzureConnector           OperationTypeEnum = "DELETE_AZURE_CONNECTOR"
	OperationTypeUpdateAzureConnector           OperationTypeEnum = "UPDATE_AZURE_CONNECTOR"
	OperationTypeMoveAzureConnector             OperationTypeEnum = "MOVE_AZURE_CONNECTOR"
	OperationTypeRefreshAzureConnector          OperationTypeEnum = "REFRESH_AZURE_CONNECTOR"
	OperationTypeCreateAzureBlobContainer       OperationTypeEnum = "CREATE_AZURE_BLOB_CONTAINER"
	OperationTypeDeleteAzureBlobContainer       OperationTypeEnum = "DELETE_AZURE_BLOB_CONTAINER"
	OperationTypeUpdateAzureBlobContainer       OperationTypeEnum = "UPDATE_AZURE_BLOB_CONTAINER"
	OperationTypeMoveAzureBlobContainer         OperationTypeEnum = "MOVE_AZURE_BLOB_CONTAINER"
	OperationTypeCreateAzureBlobMount           OperationTypeEnum = "CREATE_AZURE_BLOB_MOUNT"
	OperationTypeMoveAzureBlobMount             OperationTypeEnum = "MOVE_AZURE_BLOB_MOUNT"
	OperationTypeUpdateAzureBlobMount           OperationTypeEnum = "UPDATE_AZURE_BLOB_MOUNT"
	OperationTypeDeleteAzureBlobMount           OperationTypeEnum = "DELETE_AZURE_BLOB_MOUNT"
	OperationTypeCreateMulticloudDiscovery      OperationTypeEnum = "CREATE_MULTICLOUD_DISCOVERY"
	OperationTypeDeleteMulticloudDiscovery      OperationTypeEnum = "DELETE_MULTICLOUD_DISCOVERY"
	OperationTypeUpdateMulticloudDiscovery      OperationTypeEnum = "UPDATE_MULTICLOUD_DISCOVERY"
	OperationTypeMoveMulticloudDiscovery        OperationTypeEnum = "MOVE_MULTICLOUD_DISCOVERY"
	OperationTypeCreateAzureVault               OperationTypeEnum = "CREATE_AZURE_VAULT"
	OperationTypeDeleteAzureVault               OperationTypeEnum = "DELETE_AZURE_VAULT"
	OperationTypeUpdateAzureVault               OperationTypeEnum = "UPDATE_AZURE_VAULT"
	OperationTypeMoveAzureVault                 OperationTypeEnum = "MOVE_AZURE_VAULT"
	OperationTypeRefreshAzureVault              OperationTypeEnum = "REFRESH_AZURE_VAULT"
	OperationTypeCreateAzureVaultAssociation    OperationTypeEnum = "CREATE_AZURE_VAULT_ASSOCIATION"
	OperationTypeDeleteAzureVaultAssociation    OperationTypeEnum = "DELETE_AZURE_VAULT_ASSOCIATION"
	OperationTypeUpdateAzureVaultAssociation    OperationTypeEnum = "UPDATE_AZURE_VAULT_ASSOCIATION"
	OperationTypeMoveAzureVaultAssociation      OperationTypeEnum = "MOVE_AZURE_VAULT_ASSOCIATION"
	OperationTypePatchDbResource                OperationTypeEnum = "PATCH_DB_RESOURCE"
	OperationTypeCreateGcpIdentityConfiguration OperationTypeEnum = "CREATE_GCP_IDENTITY_CONFIGURATION"
	OperationTypeDeleteGcpIdentityConfiguration OperationTypeEnum = "DELETE_GCP_IDENTITY_CONFIGURATION"
	OperationTypeCreateGcpConnector             OperationTypeEnum = "CREATE_GCP_CONNECTOR"
	OperationTypeDeleteGcpConnector             OperationTypeEnum = "DELETE_GCP_CONNECTOR"
	OperationTypeUpdateGcpConnector             OperationTypeEnum = "UPDATE_GCP_CONNECTOR"
	OperationTypeMoveGcpConnector               OperationTypeEnum = "MOVE_GCP_CONNECTOR"
	OperationTypeRefreshGcpConnector            OperationTypeEnum = "REFRESH_GCP_CONNECTOR"
	OperationTypeGcpDiscovery                   OperationTypeEnum = "GCP_DISCOVERY"
	OperationTypeCreateGcpKeyRing               OperationTypeEnum = "CREATE_GCP_KEY_RING"
	OperationTypeDeleteGcpKeyRing               OperationTypeEnum = "DELETE_GCP_KEY_RING"
	OperationTypeUpdateGcpKeyRing               OperationTypeEnum = "UPDATE_GCP_KEY_RING"
	OperationTypeMoveGcpKeyRing                 OperationTypeEnum = "MOVE_GCP_KEY_RING"
	OperationTypeRefreshGcpKeyRing              OperationTypeEnum = "REFRESH_GCP_KEY_RING"
	OperationTypeCreateAwsConnector             OperationTypeEnum = "CREATE_AWS_CONNECTOR"
	OperationTypeDeleteAwsConnector             OperationTypeEnum = "DELETE_AWS_CONNECTOR"
	OperationTypeUpdateAwsConnector             OperationTypeEnum = "UPDATE_AWS_CONNECTOR"
	OperationTypeMoveAwsConnector               OperationTypeEnum = "MOVE_AWS_CONNECTOR"
	OperationTypeRefreshAwsConnector            OperationTypeEnum = "REFRESH_AWS_CONNECTOR"
	OperationTypeCreateAwsKey                   OperationTypeEnum = "CREATE_AWS_KEY"
	OperationTypeDeleteAwsKey                   OperationTypeEnum = "DELETE_AWS_KEY"
	OperationTypeUpdateAwsKey                   OperationTypeEnum = "UPDATE_AWS_KEY"
	OperationTypeMoveAwsKey                     OperationTypeEnum = "MOVE_AWS_KEY"
	OperationTypeRefreshAwsKey                  OperationTypeEnum = "REFRESH_AWS_KEY"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_AZURE_CONNECTOR":            OperationTypeCreateAzureConnector,
	"DELETE_AZURE_CONNECTOR":            OperationTypeDeleteAzureConnector,
	"UPDATE_AZURE_CONNECTOR":            OperationTypeUpdateAzureConnector,
	"MOVE_AZURE_CONNECTOR":              OperationTypeMoveAzureConnector,
	"REFRESH_AZURE_CONNECTOR":           OperationTypeRefreshAzureConnector,
	"CREATE_AZURE_BLOB_CONTAINER":       OperationTypeCreateAzureBlobContainer,
	"DELETE_AZURE_BLOB_CONTAINER":       OperationTypeDeleteAzureBlobContainer,
	"UPDATE_AZURE_BLOB_CONTAINER":       OperationTypeUpdateAzureBlobContainer,
	"MOVE_AZURE_BLOB_CONTAINER":         OperationTypeMoveAzureBlobContainer,
	"CREATE_AZURE_BLOB_MOUNT":           OperationTypeCreateAzureBlobMount,
	"MOVE_AZURE_BLOB_MOUNT":             OperationTypeMoveAzureBlobMount,
	"UPDATE_AZURE_BLOB_MOUNT":           OperationTypeUpdateAzureBlobMount,
	"DELETE_AZURE_BLOB_MOUNT":           OperationTypeDeleteAzureBlobMount,
	"CREATE_MULTICLOUD_DISCOVERY":       OperationTypeCreateMulticloudDiscovery,
	"DELETE_MULTICLOUD_DISCOVERY":       OperationTypeDeleteMulticloudDiscovery,
	"UPDATE_MULTICLOUD_DISCOVERY":       OperationTypeUpdateMulticloudDiscovery,
	"MOVE_MULTICLOUD_DISCOVERY":         OperationTypeMoveMulticloudDiscovery,
	"CREATE_AZURE_VAULT":                OperationTypeCreateAzureVault,
	"DELETE_AZURE_VAULT":                OperationTypeDeleteAzureVault,
	"UPDATE_AZURE_VAULT":                OperationTypeUpdateAzureVault,
	"MOVE_AZURE_VAULT":                  OperationTypeMoveAzureVault,
	"REFRESH_AZURE_VAULT":               OperationTypeRefreshAzureVault,
	"CREATE_AZURE_VAULT_ASSOCIATION":    OperationTypeCreateAzureVaultAssociation,
	"DELETE_AZURE_VAULT_ASSOCIATION":    OperationTypeDeleteAzureVaultAssociation,
	"UPDATE_AZURE_VAULT_ASSOCIATION":    OperationTypeUpdateAzureVaultAssociation,
	"MOVE_AZURE_VAULT_ASSOCIATION":      OperationTypeMoveAzureVaultAssociation,
	"PATCH_DB_RESOURCE":                 OperationTypePatchDbResource,
	"CREATE_GCP_IDENTITY_CONFIGURATION": OperationTypeCreateGcpIdentityConfiguration,
	"DELETE_GCP_IDENTITY_CONFIGURATION": OperationTypeDeleteGcpIdentityConfiguration,
	"CREATE_GCP_CONNECTOR":              OperationTypeCreateGcpConnector,
	"DELETE_GCP_CONNECTOR":              OperationTypeDeleteGcpConnector,
	"UPDATE_GCP_CONNECTOR":              OperationTypeUpdateGcpConnector,
	"MOVE_GCP_CONNECTOR":                OperationTypeMoveGcpConnector,
	"REFRESH_GCP_CONNECTOR":             OperationTypeRefreshGcpConnector,
	"GCP_DISCOVERY":                     OperationTypeGcpDiscovery,
	"CREATE_GCP_KEY_RING":               OperationTypeCreateGcpKeyRing,
	"DELETE_GCP_KEY_RING":               OperationTypeDeleteGcpKeyRing,
	"UPDATE_GCP_KEY_RING":               OperationTypeUpdateGcpKeyRing,
	"MOVE_GCP_KEY_RING":                 OperationTypeMoveGcpKeyRing,
	"REFRESH_GCP_KEY_RING":              OperationTypeRefreshGcpKeyRing,
	"CREATE_AWS_CONNECTOR":              OperationTypeCreateAwsConnector,
	"DELETE_AWS_CONNECTOR":              OperationTypeDeleteAwsConnector,
	"UPDATE_AWS_CONNECTOR":              OperationTypeUpdateAwsConnector,
	"MOVE_AWS_CONNECTOR":                OperationTypeMoveAwsConnector,
	"REFRESH_AWS_CONNECTOR":             OperationTypeRefreshAwsConnector,
	"CREATE_AWS_KEY":                    OperationTypeCreateAwsKey,
	"DELETE_AWS_KEY":                    OperationTypeDeleteAwsKey,
	"UPDATE_AWS_KEY":                    OperationTypeUpdateAwsKey,
	"MOVE_AWS_KEY":                      OperationTypeMoveAwsKey,
	"REFRESH_AWS_KEY":                   OperationTypeRefreshAwsKey,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_azure_connector":            OperationTypeCreateAzureConnector,
	"delete_azure_connector":            OperationTypeDeleteAzureConnector,
	"update_azure_connector":            OperationTypeUpdateAzureConnector,
	"move_azure_connector":              OperationTypeMoveAzureConnector,
	"refresh_azure_connector":           OperationTypeRefreshAzureConnector,
	"create_azure_blob_container":       OperationTypeCreateAzureBlobContainer,
	"delete_azure_blob_container":       OperationTypeDeleteAzureBlobContainer,
	"update_azure_blob_container":       OperationTypeUpdateAzureBlobContainer,
	"move_azure_blob_container":         OperationTypeMoveAzureBlobContainer,
	"create_azure_blob_mount":           OperationTypeCreateAzureBlobMount,
	"move_azure_blob_mount":             OperationTypeMoveAzureBlobMount,
	"update_azure_blob_mount":           OperationTypeUpdateAzureBlobMount,
	"delete_azure_blob_mount":           OperationTypeDeleteAzureBlobMount,
	"create_multicloud_discovery":       OperationTypeCreateMulticloudDiscovery,
	"delete_multicloud_discovery":       OperationTypeDeleteMulticloudDiscovery,
	"update_multicloud_discovery":       OperationTypeUpdateMulticloudDiscovery,
	"move_multicloud_discovery":         OperationTypeMoveMulticloudDiscovery,
	"create_azure_vault":                OperationTypeCreateAzureVault,
	"delete_azure_vault":                OperationTypeDeleteAzureVault,
	"update_azure_vault":                OperationTypeUpdateAzureVault,
	"move_azure_vault":                  OperationTypeMoveAzureVault,
	"refresh_azure_vault":               OperationTypeRefreshAzureVault,
	"create_azure_vault_association":    OperationTypeCreateAzureVaultAssociation,
	"delete_azure_vault_association":    OperationTypeDeleteAzureVaultAssociation,
	"update_azure_vault_association":    OperationTypeUpdateAzureVaultAssociation,
	"move_azure_vault_association":      OperationTypeMoveAzureVaultAssociation,
	"patch_db_resource":                 OperationTypePatchDbResource,
	"create_gcp_identity_configuration": OperationTypeCreateGcpIdentityConfiguration,
	"delete_gcp_identity_configuration": OperationTypeDeleteGcpIdentityConfiguration,
	"create_gcp_connector":              OperationTypeCreateGcpConnector,
	"delete_gcp_connector":              OperationTypeDeleteGcpConnector,
	"update_gcp_connector":              OperationTypeUpdateGcpConnector,
	"move_gcp_connector":                OperationTypeMoveGcpConnector,
	"refresh_gcp_connector":             OperationTypeRefreshGcpConnector,
	"gcp_discovery":                     OperationTypeGcpDiscovery,
	"create_gcp_key_ring":               OperationTypeCreateGcpKeyRing,
	"delete_gcp_key_ring":               OperationTypeDeleteGcpKeyRing,
	"update_gcp_key_ring":               OperationTypeUpdateGcpKeyRing,
	"move_gcp_key_ring":                 OperationTypeMoveGcpKeyRing,
	"refresh_gcp_key_ring":              OperationTypeRefreshGcpKeyRing,
	"create_aws_connector":              OperationTypeCreateAwsConnector,
	"delete_aws_connector":              OperationTypeDeleteAwsConnector,
	"update_aws_connector":              OperationTypeUpdateAwsConnector,
	"move_aws_connector":                OperationTypeMoveAwsConnector,
	"refresh_aws_connector":             OperationTypeRefreshAwsConnector,
	"create_aws_key":                    OperationTypeCreateAwsKey,
	"delete_aws_key":                    OperationTypeDeleteAwsKey,
	"update_aws_key":                    OperationTypeUpdateAwsKey,
	"move_aws_key":                      OperationTypeMoveAwsKey,
	"refresh_aws_key":                   OperationTypeRefreshAwsKey,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_AZURE_CONNECTOR",
		"DELETE_AZURE_CONNECTOR",
		"UPDATE_AZURE_CONNECTOR",
		"MOVE_AZURE_CONNECTOR",
		"REFRESH_AZURE_CONNECTOR",
		"CREATE_AZURE_BLOB_CONTAINER",
		"DELETE_AZURE_BLOB_CONTAINER",
		"UPDATE_AZURE_BLOB_CONTAINER",
		"MOVE_AZURE_BLOB_CONTAINER",
		"CREATE_AZURE_BLOB_MOUNT",
		"MOVE_AZURE_BLOB_MOUNT",
		"UPDATE_AZURE_BLOB_MOUNT",
		"DELETE_AZURE_BLOB_MOUNT",
		"CREATE_MULTICLOUD_DISCOVERY",
		"DELETE_MULTICLOUD_DISCOVERY",
		"UPDATE_MULTICLOUD_DISCOVERY",
		"MOVE_MULTICLOUD_DISCOVERY",
		"CREATE_AZURE_VAULT",
		"DELETE_AZURE_VAULT",
		"UPDATE_AZURE_VAULT",
		"MOVE_AZURE_VAULT",
		"REFRESH_AZURE_VAULT",
		"CREATE_AZURE_VAULT_ASSOCIATION",
		"DELETE_AZURE_VAULT_ASSOCIATION",
		"UPDATE_AZURE_VAULT_ASSOCIATION",
		"MOVE_AZURE_VAULT_ASSOCIATION",
		"PATCH_DB_RESOURCE",
		"CREATE_GCP_IDENTITY_CONFIGURATION",
		"DELETE_GCP_IDENTITY_CONFIGURATION",
		"CREATE_GCP_CONNECTOR",
		"DELETE_GCP_CONNECTOR",
		"UPDATE_GCP_CONNECTOR",
		"MOVE_GCP_CONNECTOR",
		"REFRESH_GCP_CONNECTOR",
		"GCP_DISCOVERY",
		"CREATE_GCP_KEY_RING",
		"DELETE_GCP_KEY_RING",
		"UPDATE_GCP_KEY_RING",
		"MOVE_GCP_KEY_RING",
		"REFRESH_GCP_KEY_RING",
		"CREATE_AWS_CONNECTOR",
		"DELETE_AWS_CONNECTOR",
		"UPDATE_AWS_CONNECTOR",
		"MOVE_AWS_CONNECTOR",
		"REFRESH_AWS_CONNECTOR",
		"CREATE_AWS_KEY",
		"DELETE_AWS_KEY",
		"UPDATE_AWS_KEY",
		"MOVE_AWS_KEY",
		"REFRESH_AWS_KEY",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
