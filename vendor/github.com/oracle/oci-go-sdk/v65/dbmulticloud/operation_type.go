// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
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
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_AZURE_CONNECTOR":            OperationTypeCreateAzureConnector,
	"DELETE_AZURE_CONNECTOR":            OperationTypeDeleteAzureConnector,
	"UPDATE_AZURE_CONNECTOR":            OperationTypeUpdateAzureConnector,
	"MOVE_AZURE_CONNECTOR":              OperationTypeMoveAzureConnector,
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
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_azure_connector":            OperationTypeCreateAzureConnector,
	"delete_azure_connector":            OperationTypeDeleteAzureConnector,
	"update_azure_connector":            OperationTypeUpdateAzureConnector,
	"move_azure_connector":              OperationTypeMoveAzureConnector,
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
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
